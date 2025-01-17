all: gen-proto test

test: gen-proto
	go test ./...
	cd internal/examples/simple && go test -run=nosuchname -bench . --benchmem -benchtime=1ms

benchmark:
	-rm internal/benchmark/benchmark.log
	-rm internal/benchmark/benchmark.csv
	cd internal/examples/simple && FORREPORT=1 go test -run=nosuchname -bench BenchmarkGoogle --benchmem $(BENCHARGS) | tee ../../benchmark/benchmark-temp.log
	cd internal/benchmark && sed -f patch_results.sed benchmark-temp.log > benchmark.log
	cd internal/benchmark && benchstat -csv ./benchmark.log >> benchmark.csv
	cd internal/examples/simple && FORREPORT=1 go test -run=nosuchname -bench BenchmarkGogo --benchmem $(BENCHARGS) | tee ../../benchmark/benchmark-temp.log
	cd internal/benchmark && sed -f patch_results.sed benchmark-temp.log > benchmark.log
	cd internal/benchmark && benchstat -csv ./benchmark.log >> benchmark.csv
	cd internal/examples/simple && FORREPORT=1 VALIDATE=1 go test -run=nosuchname -bench BenchmarkLazy --benchmem $(BENCHARGS) | tee ../../benchmark/benchmark-temp.log
	cd internal/benchmark && sed -f patch_results.sed benchmark-temp.log > benchmark.log
	cd internal/benchmark && benchstat -csv ./benchmark.log >> benchmark.csv
	cd internal/examples/simple && FORREPORT=1 go test -run=nosuchname -bench BenchmarkLazy --benchmem $(BENCHARGS) | tee ../../benchmark/benchmark-temp.log
	cd internal/benchmark && sed -f patch_results.sed benchmark-temp.log > benchmark.log
	cd internal/benchmark && benchstat -csv ./benchmark.log >> benchmark.csv

.PHONY: gen-proto
gen-proto: gen-gogo gen-google gen-lazy

.PHONY: gen-gogo
gen-gogo: internal/examples/simple/gogo/gen/logs/logs.pb.go

.PHONY: gen-google
gen-google: internal/examples/simple/google/gen/logs/logs.pb.go

.PHONY: gen-lazy
gen-lazy: internal/examples/simple/lazy/logs.pb.go

internal/examples/simple/gogo/gen/logs/logs.pb.go: internal/examples/simple/logs.proto Makefile
	docker run --rm -v${PWD}:${PWD} \
            -w${PWD} otel/build-protobuf:latest --proto_path=${PWD}/internal/examples/simple \
            --gogofaster_out=plugins=grpc:./internal/examples/simple/gogo/ ${PWD}/internal/examples/simple/logs.proto

internal/examples/simple/google/gen/logs/logs.pb.go: internal/examples/simple/logs.proto Makefile
	docker run --rm -v${PWD}:${PWD} \
            -w${PWD} otel/build-protobuf:latest --proto_path=${PWD}/internal/examples/simple \
            --go_out=plugins=grpc:./internal/examples/simple/google/ ${PWD}/internal/examples/simple/logs.proto

internal/examples/simple/lazy/logs.pb.go: internal/examples/simple/logs.proto Makefile
	go run cmd/main.go --proto_path internal/examples/simple --go_out internal/examples/simple/lazy logs.proto
