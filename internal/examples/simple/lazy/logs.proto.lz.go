package simple

import (
	"sync"

	lazyproto "github.com/tigrannajaryan/exp-lazyproto"
	"github.com/tigrannajaryan/molecule"
	"github.com/tigrannajaryan/molecule/src/codec"
)

// ====================== Generated for message LogsData ======================

// LogsData contains all log data
type LogsData struct {
	protoMessage lazyproto.ProtoMessage
	// List of ResourceLogs
	resourceLogs []*ResourceLogs
}

func NewLogsData(bytes []byte) *LogsData {
	m := &LogsData{protoMessage: lazyproto.ProtoMessage{Bytes: bytes}}
	m.decode()
	return m
}

// Bitmasks that indicate that the particular nested message is decoded.
const flagLogsDataResourceLogsDecoded = 0x0000000000000002

func (m *LogsData) GetResourceLogs() []*ResourceLogs {
	if m.protoMessage.Flags&flagLogsDataResourceLogsDecoded == 0 {
		// Decode nested message(s).
		for i := range m.resourceLogs {
			m.resourceLogs[i].decode()
		}
		m.protoMessage.Flags |= flagLogsDataResourceLogsDecoded
	}
	return m.resourceLogs
}

func (m *LogsData) decode() {
	buf := codec.NewBuffer(m.protoMessage.Bytes)

	// Count all repeated fields. We need one counter per field.
	resourceLogsCount := 0
	molecule.MessageFieldNums(
		buf, func(fieldNum int32) {
			if fieldNum == 1 {
				resourceLogsCount++
			}
		},
	)

	// Pre-allocate slices for repeated fields.
	m.resourceLogs = make([]*ResourceLogs, 0, resourceLogsCount)

	// Reset the buffer to start iterating over the fields again
	buf.Reset(m.protoMessage.Bytes)

	// Set slice indexes to 0 to begin iterating over repeated fields.
	resourceLogsCount = 0
	// Iterate and decode the fields.
	molecule.MessageEach(
		buf, func(fieldNum int32, value molecule.Value) (bool, error) {
			switch fieldNum {
			case 1:
				// Decode resourceLogs.
				v, err := value.AsBytesUnsafe()
				if err != nil {
					return false, err
				}
				// The slice is pre-allocated, assign to the appropriate index.
				m.resourceLogs[resourceLogsCount] = &ResourceLogs{
					protoMessage: lazyproto.ProtoMessage{
						Parent: &m.protoMessage, Bytes: v,
					},
				}
				resourceLogsCount++
			}
			return true, nil
		},
	)
}

func (m *LogsData) Marshal(ps *molecule.ProtoStream) error {
	if m.protoMessage.Flags&lazyproto.FlagsMessageModified != 0 {
		// Marshal resourceLogs
		for _, elem := range m.resourceLogs {
			token := ps.BeginEmbedded()
			elem.Marshal(ps)
			ps.EndEmbedded(token, 1)
		}
	} else {
		// Message is unchanged. Used original bytes.
		ps.Raw(m.protoMessage.Bytes)
	}
	return nil
}

// Pool of LogsData structs.
type logsDataPoolType struct {
	pool []*LogsData
	mux  sync.Mutex
}

var logsDataPool = logsDataPoolType{}

// Get one element from the pool. Creates a new element if the pool is empty.
func (p *logsDataPoolType) Get() *LogsData {
	p.mux.Lock()
	defer p.mux.Unlock()

	// Have elements in the pool?
	if len(p.pool) >= 1 {
		// Get the last element.
		r := p.pool[len(p.pool)-1]
		// Shrink the pool.
		p.pool = p.pool[:len(p.pool)-1]
		return r
	}

	// Pool is empty, create a new element.
	return &LogsData{}
}

func (p *logsDataPoolType) GetSlice(count int) []*LogsData {
	p.mux.Lock()
	defer p.mux.Unlock()

	// Have enough elements in the pool?
	if len(p.pool) >= count {
		// Cut the required slice from the end of the pool.
		r := p.pool[len(p.pool)-count:]
		// Shrink the pool.
		p.pool = p.pool[:len(p.pool)-count]
		return r
	}

	// Create a new slice.
	r := make([]*LogsData, count)

	// Initialize with what remains in the pool.
	i := 0
	for ; i < len(p.pool); i++ {
		r[i] = p.pool[i]
	}
	p.pool = nil

	if i < count {
		// Create remaining elements.
		storage := make([]LogsData, count-i)
		j := 0
		for ; i < count; i++ {
			r[i] = &storage[j]
			j++
		}
	}

	return r
}

// ReleaseSlice releases a slice of elements back to the pool.
func (p *logsDataPoolType) ReleaseSlice(slice []*LogsData) {
	for _, elem := range slice {
		// Release nested resourceLogs recursively to their pool.
		resourceLogsPool.ReleaseSlice(elem.resourceLogs)

		// Zero-initialize the released element.
		*elem = LogsData{}
	}

	p.mux.Lock()
	defer p.mux.Unlock()

	// Add the slice to the end of the pool.
	p.pool = append(p.pool, slice...)
}

// Release an element back to the pool.
func (p *logsDataPoolType) Release(elem *LogsData) {
	// Release nested resourceLogs recursively to their pool.
	resourceLogsPool.ReleaseSlice(elem.resourceLogs)

	// Zero-initialize the released element.
	*elem = LogsData{}

	p.mux.Lock()
	defer p.mux.Unlock()

	// Add the slice to the end of the pool.
	p.pool = append(p.pool, elem)
}

// ====================== Generated for message ResourceLogs ======================

type ResourceLogs struct {
	protoMessage lazyproto.ProtoMessage
	// The Resource
	resource *Resource
	// List of ScopeLogs
	scopeLogs []*ScopeLogs
}

func NewResourceLogs(bytes []byte) *ResourceLogs {
	m := &ResourceLogs{protoMessage: lazyproto.ProtoMessage{Bytes: bytes}}
	m.decode()
	return m
}

// Bitmasks that indicate that the particular nested message is decoded.
const flagResourceLogsResourceDecoded = 0x0000000000000002
const flagResourceLogsScopeLogsDecoded = 0x0000000000000004

func (m *ResourceLogs) GetResource() *Resource {
	if m.protoMessage.Flags&flagResourceLogsResourceDecoded == 0 {
		// Decode nested message(s).
		m.resource.decode()
		m.protoMessage.Flags |= flagResourceLogsResourceDecoded
	}
	return m.resource
}

func (m *ResourceLogs) GetScopeLogs() []*ScopeLogs {
	if m.protoMessage.Flags&flagResourceLogsScopeLogsDecoded == 0 {
		// Decode nested message(s).
		for i := range m.scopeLogs {
			m.scopeLogs[i].decode()
		}
		m.protoMessage.Flags |= flagResourceLogsScopeLogsDecoded
	}
	return m.scopeLogs
}

func (m *ResourceLogs) decode() {
	buf := codec.NewBuffer(m.protoMessage.Bytes)

	// Count all repeated fields. We need one counter per field.
	scopeLogsCount := 0
	molecule.MessageFieldNums(
		buf, func(fieldNum int32) {
			if fieldNum == 2 {
				scopeLogsCount++
			}
		},
	)

	// Pre-allocate slices for repeated fields.
	m.scopeLogs = make([]*ScopeLogs, 0, scopeLogsCount)

	// Reset the buffer to start iterating over the fields again
	buf.Reset(m.protoMessage.Bytes)

	// Set slice indexes to 0 to begin iterating over repeated fields.
	scopeLogsCount = 0
	// Iterate and decode the fields.
	molecule.MessageEach(
		buf, func(fieldNum int32, value molecule.Value) (bool, error) {
			switch fieldNum {
			case 1:
				// Decode resource.
				v, err := value.AsBytesUnsafe()
				if err != nil {
					return false, err
				}
				m.resource = &Resource{
					protoMessage: lazyproto.ProtoMessage{
						Parent: &m.protoMessage, Bytes: v,
					},
				}
			case 2:
				// Decode scopeLogs.
				v, err := value.AsBytesUnsafe()
				if err != nil {
					return false, err
				}
				// The slice is pre-allocated, assign to the appropriate index.
				m.scopeLogs[scopeLogsCount] = &ScopeLogs{
					protoMessage: lazyproto.ProtoMessage{
						Parent: &m.protoMessage, Bytes: v,
					},
				}
				scopeLogsCount++
			}
			return true, nil
		},
	)
}

func (m *ResourceLogs) Marshal(ps *molecule.ProtoStream) error {
	if m.protoMessage.Flags&lazyproto.FlagsMessageModified != 0 {
		// Marshal resource
		if m.resource != nil {
			token := ps.BeginEmbedded()
			m.resource.Marshal(ps)
			ps.EndEmbedded(token, 1)
		}
		// Marshal scopeLogs
		for _, elem := range m.scopeLogs {
			token := ps.BeginEmbedded()
			elem.Marshal(ps)
			ps.EndEmbedded(token, 2)
		}
	} else {
		// Message is unchanged. Used original bytes.
		ps.Raw(m.protoMessage.Bytes)
	}
	return nil
}

// Pool of ResourceLogs structs.
type resourceLogsPoolType struct {
	pool []*ResourceLogs
	mux  sync.Mutex
}

var resourceLogsPool = resourceLogsPoolType{}

// Get one element from the pool. Creates a new element if the pool is empty.
func (p *resourceLogsPoolType) Get() *ResourceLogs {
	p.mux.Lock()
	defer p.mux.Unlock()

	// Have elements in the pool?
	if len(p.pool) >= 1 {
		// Get the last element.
		r := p.pool[len(p.pool)-1]
		// Shrink the pool.
		p.pool = p.pool[:len(p.pool)-1]
		return r
	}

	// Pool is empty, create a new element.
	return &ResourceLogs{}
}

func (p *resourceLogsPoolType) GetSlice(count int) []*ResourceLogs {
	p.mux.Lock()
	defer p.mux.Unlock()

	// Have enough elements in the pool?
	if len(p.pool) >= count {
		// Cut the required slice from the end of the pool.
		r := p.pool[len(p.pool)-count:]
		// Shrink the pool.
		p.pool = p.pool[:len(p.pool)-count]
		return r
	}

	// Create a new slice.
	r := make([]*ResourceLogs, count)

	// Initialize with what remains in the pool.
	i := 0
	for ; i < len(p.pool); i++ {
		r[i] = p.pool[i]
	}
	p.pool = nil

	if i < count {
		// Create remaining elements.
		storage := make([]ResourceLogs, count-i)
		j := 0
		for ; i < count; i++ {
			r[i] = &storage[j]
			j++
		}
	}

	return r
}

// ReleaseSlice releases a slice of elements back to the pool.
func (p *resourceLogsPoolType) ReleaseSlice(slice []*ResourceLogs) {
	for _, elem := range slice {
		// Release nested resource recursively to their pool.
		if elem.resource != nil {
			resourcePool.Release(elem.resource)
		}
		// Release nested scopeLogs recursively to their pool.
		scopeLogsPool.ReleaseSlice(elem.scopeLogs)

		// Zero-initialize the released element.
		*elem = ResourceLogs{}
	}

	p.mux.Lock()
	defer p.mux.Unlock()

	// Add the slice to the end of the pool.
	p.pool = append(p.pool, slice...)
}

// Release an element back to the pool.
func (p *resourceLogsPoolType) Release(elem *ResourceLogs) {
	// Release nested resource recursively to their pool.
	if elem.resource != nil {
		resourcePool.Release(elem.resource)
	}
	// Release nested scopeLogs recursively to their pool.
	scopeLogsPool.ReleaseSlice(elem.scopeLogs)

	// Zero-initialize the released element.
	*elem = ResourceLogs{}

	p.mux.Lock()
	defer p.mux.Unlock()

	// Add the slice to the end of the pool.
	p.pool = append(p.pool, elem)
}

// ====================== Generated for message Resource ======================

type Resource struct {
	protoMessage           lazyproto.ProtoMessage
	attributes             []*KeyValue
	droppedAttributesCount uint32
}

func NewResource(bytes []byte) *Resource {
	m := &Resource{protoMessage: lazyproto.ProtoMessage{Bytes: bytes}}
	m.decode()
	return m
}

// Bitmasks that indicate that the particular nested message is decoded.
const flagResourceAttributesDecoded = 0x0000000000000002

func (m *Resource) GetAttributes() []*KeyValue {
	if m.protoMessage.Flags&flagResourceAttributesDecoded == 0 {
		// Decode nested message(s).
		for i := range m.attributes {
			m.attributes[i].decode()
		}
		m.protoMessage.Flags |= flagResourceAttributesDecoded
	}
	return m.attributes
}

func (m *Resource) GetDroppedAttributesCount() uint32 {
	return m.droppedAttributesCount
}

func (m *Resource) decode() {
	buf := codec.NewBuffer(m.protoMessage.Bytes)

	// Count all repeated fields. We need one counter per field.
	attributesCount := 0
	molecule.MessageFieldNums(
		buf, func(fieldNum int32) {
			if fieldNum == 1 {
				attributesCount++
			}
		},
	)

	// Pre-allocate slices for repeated fields.
	m.attributes = make([]*KeyValue, 0, attributesCount)

	// Reset the buffer to start iterating over the fields again
	buf.Reset(m.protoMessage.Bytes)

	// Set slice indexes to 0 to begin iterating over repeated fields.
	attributesCount = 0
	// Iterate and decode the fields.
	molecule.MessageEach(
		buf, func(fieldNum int32, value molecule.Value) (bool, error) {
			switch fieldNum {
			case 1:
				// Decode attributes.
				v, err := value.AsBytesUnsafe()
				if err != nil {
					return false, err
				}
				// The slice is pre-allocated, assign to the appropriate index.
				m.attributes[attributesCount] = &KeyValue{
					protoMessage: lazyproto.ProtoMessage{
						Parent: &m.protoMessage, Bytes: v,
					},
				}
				attributesCount++
			case 2:
				// Decode droppedAttributesCount.
				v, err := value.AsUint32()
				if err != nil {
					return false, err
				}
				m.droppedAttributesCount = v
			}
			return true, nil
		},
	)
}

var preparedResourceDroppedAttributesCount = molecule.PrepareUint32Field(2)

func (m *Resource) Marshal(ps *molecule.ProtoStream) error {
	if m.protoMessage.Flags&lazyproto.FlagsMessageModified != 0 {
		// Marshal attributes
		for _, elem := range m.attributes {
			token := ps.BeginEmbedded()
			elem.Marshal(ps)
			ps.EndEmbedded(token, 1)
		}
		// Marshal droppedAttributesCount
		ps.Uint32Prepared(preparedResourceDroppedAttributesCount, m.droppedAttributesCount)
	} else {
		// Message is unchanged. Used original bytes.
		ps.Raw(m.protoMessage.Bytes)
	}
	return nil
}

// Pool of Resource structs.
type resourcePoolType struct {
	pool []*Resource
	mux  sync.Mutex
}

var resourcePool = resourcePoolType{}

// Get one element from the pool. Creates a new element if the pool is empty.
func (p *resourcePoolType) Get() *Resource {
	p.mux.Lock()
	defer p.mux.Unlock()

	// Have elements in the pool?
	if len(p.pool) >= 1 {
		// Get the last element.
		r := p.pool[len(p.pool)-1]
		// Shrink the pool.
		p.pool = p.pool[:len(p.pool)-1]
		return r
	}

	// Pool is empty, create a new element.
	return &Resource{}
}

func (p *resourcePoolType) GetSlice(count int) []*Resource {
	p.mux.Lock()
	defer p.mux.Unlock()

	// Have enough elements in the pool?
	if len(p.pool) >= count {
		// Cut the required slice from the end of the pool.
		r := p.pool[len(p.pool)-count:]
		// Shrink the pool.
		p.pool = p.pool[:len(p.pool)-count]
		return r
	}

	// Create a new slice.
	r := make([]*Resource, count)

	// Initialize with what remains in the pool.
	i := 0
	for ; i < len(p.pool); i++ {
		r[i] = p.pool[i]
	}
	p.pool = nil

	if i < count {
		// Create remaining elements.
		storage := make([]Resource, count-i)
		j := 0
		for ; i < count; i++ {
			r[i] = &storage[j]
			j++
		}
	}

	return r
}

// ReleaseSlice releases a slice of elements back to the pool.
func (p *resourcePoolType) ReleaseSlice(slice []*Resource) {
	for _, elem := range slice {
		// Release nested attributes recursively to their pool.
		keyValuePool.ReleaseSlice(elem.attributes)

		// Zero-initialize the released element.
		*elem = Resource{}
	}

	p.mux.Lock()
	defer p.mux.Unlock()

	// Add the slice to the end of the pool.
	p.pool = append(p.pool, slice...)
}

// Release an element back to the pool.
func (p *resourcePoolType) Release(elem *Resource) {
	// Release nested attributes recursively to their pool.
	keyValuePool.ReleaseSlice(elem.attributes)

	// Zero-initialize the released element.
	*elem = Resource{}

	p.mux.Lock()
	defer p.mux.Unlock()

	// Add the slice to the end of the pool.
	p.pool = append(p.pool, elem)
}

// ====================== Generated for message ScopeLogs ======================

// A collection of Logs produced by a Scope.
type ScopeLogs struct {
	protoMessage lazyproto.ProtoMessage
	// A list of log records.
	logRecords []*LogRecord
}

func NewScopeLogs(bytes []byte) *ScopeLogs {
	m := &ScopeLogs{protoMessage: lazyproto.ProtoMessage{Bytes: bytes}}
	m.decode()
	return m
}

// Bitmasks that indicate that the particular nested message is decoded.
const flagScopeLogsLogRecordsDecoded = 0x0000000000000002

func (m *ScopeLogs) GetLogRecords() []*LogRecord {
	if m.protoMessage.Flags&flagScopeLogsLogRecordsDecoded == 0 {
		// Decode nested message(s).
		for i := range m.logRecords {
			m.logRecords[i].decode()
		}
		m.protoMessage.Flags |= flagScopeLogsLogRecordsDecoded
	}
	return m.logRecords
}

func (m *ScopeLogs) decode() {
	buf := codec.NewBuffer(m.protoMessage.Bytes)

	// Count all repeated fields. We need one counter per field.
	logRecordsCount := 0
	molecule.MessageFieldNums(
		buf, func(fieldNum int32) {
			if fieldNum == 1 {
				logRecordsCount++
			}
		},
	)

	// Pre-allocate slices for repeated fields.
	m.logRecords = make([]*LogRecord, 0, logRecordsCount)

	// Reset the buffer to start iterating over the fields again
	buf.Reset(m.protoMessage.Bytes)

	// Set slice indexes to 0 to begin iterating over repeated fields.
	logRecordsCount = 0
	// Iterate and decode the fields.
	molecule.MessageEach(
		buf, func(fieldNum int32, value molecule.Value) (bool, error) {
			switch fieldNum {
			case 1:
				// Decode logRecords.
				v, err := value.AsBytesUnsafe()
				if err != nil {
					return false, err
				}
				// The slice is pre-allocated, assign to the appropriate index.
				m.logRecords[logRecordsCount] = &LogRecord{
					protoMessage: lazyproto.ProtoMessage{
						Parent: &m.protoMessage, Bytes: v,
					},
				}
				logRecordsCount++
			}
			return true, nil
		},
	)
}

func (m *ScopeLogs) Marshal(ps *molecule.ProtoStream) error {
	if m.protoMessage.Flags&lazyproto.FlagsMessageModified != 0 {
		// Marshal logRecords
		for _, elem := range m.logRecords {
			token := ps.BeginEmbedded()
			elem.Marshal(ps)
			ps.EndEmbedded(token, 1)
		}
	} else {
		// Message is unchanged. Used original bytes.
		ps.Raw(m.protoMessage.Bytes)
	}
	return nil
}

// Pool of ScopeLogs structs.
type scopeLogsPoolType struct {
	pool []*ScopeLogs
	mux  sync.Mutex
}

var scopeLogsPool = scopeLogsPoolType{}

// Get one element from the pool. Creates a new element if the pool is empty.
func (p *scopeLogsPoolType) Get() *ScopeLogs {
	p.mux.Lock()
	defer p.mux.Unlock()

	// Have elements in the pool?
	if len(p.pool) >= 1 {
		// Get the last element.
		r := p.pool[len(p.pool)-1]
		// Shrink the pool.
		p.pool = p.pool[:len(p.pool)-1]
		return r
	}

	// Pool is empty, create a new element.
	return &ScopeLogs{}
}

func (p *scopeLogsPoolType) GetSlice(count int) []*ScopeLogs {
	p.mux.Lock()
	defer p.mux.Unlock()

	// Have enough elements in the pool?
	if len(p.pool) >= count {
		// Cut the required slice from the end of the pool.
		r := p.pool[len(p.pool)-count:]
		// Shrink the pool.
		p.pool = p.pool[:len(p.pool)-count]
		return r
	}

	// Create a new slice.
	r := make([]*ScopeLogs, count)

	// Initialize with what remains in the pool.
	i := 0
	for ; i < len(p.pool); i++ {
		r[i] = p.pool[i]
	}
	p.pool = nil

	if i < count {
		// Create remaining elements.
		storage := make([]ScopeLogs, count-i)
		j := 0
		for ; i < count; i++ {
			r[i] = &storage[j]
			j++
		}
	}

	return r
}

// ReleaseSlice releases a slice of elements back to the pool.
func (p *scopeLogsPoolType) ReleaseSlice(slice []*ScopeLogs) {
	for _, elem := range slice {
		// Release nested logRecords recursively to their pool.
		logRecordPool.ReleaseSlice(elem.logRecords)

		// Zero-initialize the released element.
		*elem = ScopeLogs{}
	}

	p.mux.Lock()
	defer p.mux.Unlock()

	// Add the slice to the end of the pool.
	p.pool = append(p.pool, slice...)
}

// Release an element back to the pool.
func (p *scopeLogsPoolType) Release(elem *ScopeLogs) {
	// Release nested logRecords recursively to their pool.
	logRecordPool.ReleaseSlice(elem.logRecords)

	// Zero-initialize the released element.
	*elem = ScopeLogs{}

	p.mux.Lock()
	defer p.mux.Unlock()

	// Add the slice to the end of the pool.
	p.pool = append(p.pool, elem)
}

// ====================== Generated for message LogRecord ======================

type LogRecord struct {
	protoMessage           lazyproto.ProtoMessage
	timeUnixNano           uint64
	attributes             []*KeyValue
	droppedAttributesCount uint32
}

func NewLogRecord(bytes []byte) *LogRecord {
	m := &LogRecord{protoMessage: lazyproto.ProtoMessage{Bytes: bytes}}
	m.decode()
	return m
}

// Bitmasks that indicate that the particular nested message is decoded.
const flagLogRecordAttributesDecoded = 0x0000000000000002

func (m *LogRecord) GetTimeUnixNano() uint64 {
	return m.timeUnixNano
}

func (m *LogRecord) GetAttributes() []*KeyValue {
	if m.protoMessage.Flags&flagLogRecordAttributesDecoded == 0 {
		// Decode nested message(s).
		for i := range m.attributes {
			m.attributes[i].decode()
		}
		m.protoMessage.Flags |= flagLogRecordAttributesDecoded
	}
	return m.attributes
}

func (m *LogRecord) GetDroppedAttributesCount() uint32 {
	return m.droppedAttributesCount
}

func (m *LogRecord) decode() {
	buf := codec.NewBuffer(m.protoMessage.Bytes)

	// Count all repeated fields. We need one counter per field.
	attributesCount := 0
	molecule.MessageFieldNums(
		buf, func(fieldNum int32) {
			if fieldNum == 2 {
				attributesCount++
			}
		},
	)

	// Pre-allocate slices for repeated fields.
	m.attributes = make([]*KeyValue, 0, attributesCount)

	// Reset the buffer to start iterating over the fields again
	buf.Reset(m.protoMessage.Bytes)

	// Set slice indexes to 0 to begin iterating over repeated fields.
	attributesCount = 0
	// Iterate and decode the fields.
	molecule.MessageEach(
		buf, func(fieldNum int32, value molecule.Value) (bool, error) {
			switch fieldNum {
			case 1:
				// Decode timeUnixNano.
				v, err := value.AsFixed64()
				if err != nil {
					return false, err
				}
				m.timeUnixNano = v
			case 2:
				// Decode attributes.
				v, err := value.AsBytesUnsafe()
				if err != nil {
					return false, err
				}
				// The slice is pre-allocated, assign to the appropriate index.
				m.attributes[attributesCount] = &KeyValue{
					protoMessage: lazyproto.ProtoMessage{
						Parent: &m.protoMessage, Bytes: v,
					},
				}
				attributesCount++
			case 3:
				// Decode droppedAttributesCount.
				v, err := value.AsUint32()
				if err != nil {
					return false, err
				}
				m.droppedAttributesCount = v
			}
			return true, nil
		},
	)
}

var preparedLogRecordTimeUnixNano = molecule.PrepareFixed64Field(1)
var preparedLogRecordDroppedAttributesCount = molecule.PrepareUint32Field(3)

func (m *LogRecord) Marshal(ps *molecule.ProtoStream) error {
	if m.protoMessage.Flags&lazyproto.FlagsMessageModified != 0 {
		// Marshal timeUnixNano
		ps.Fixed64Prepared(preparedLogRecordTimeUnixNano, m.timeUnixNano)
		// Marshal attributes
		for _, elem := range m.attributes {
			token := ps.BeginEmbedded()
			elem.Marshal(ps)
			ps.EndEmbedded(token, 2)
		}
		// Marshal droppedAttributesCount
		ps.Uint32Prepared(preparedLogRecordDroppedAttributesCount, m.droppedAttributesCount)
	} else {
		// Message is unchanged. Used original bytes.
		ps.Raw(m.protoMessage.Bytes)
	}
	return nil
}

// Pool of LogRecord structs.
type logRecordPoolType struct {
	pool []*LogRecord
	mux  sync.Mutex
}

var logRecordPool = logRecordPoolType{}

// Get one element from the pool. Creates a new element if the pool is empty.
func (p *logRecordPoolType) Get() *LogRecord {
	p.mux.Lock()
	defer p.mux.Unlock()

	// Have elements in the pool?
	if len(p.pool) >= 1 {
		// Get the last element.
		r := p.pool[len(p.pool)-1]
		// Shrink the pool.
		p.pool = p.pool[:len(p.pool)-1]
		return r
	}

	// Pool is empty, create a new element.
	return &LogRecord{}
}

func (p *logRecordPoolType) GetSlice(count int) []*LogRecord {
	p.mux.Lock()
	defer p.mux.Unlock()

	// Have enough elements in the pool?
	if len(p.pool) >= count {
		// Cut the required slice from the end of the pool.
		r := p.pool[len(p.pool)-count:]
		// Shrink the pool.
		p.pool = p.pool[:len(p.pool)-count]
		return r
	}

	// Create a new slice.
	r := make([]*LogRecord, count)

	// Initialize with what remains in the pool.
	i := 0
	for ; i < len(p.pool); i++ {
		r[i] = p.pool[i]
	}
	p.pool = nil

	if i < count {
		// Create remaining elements.
		storage := make([]LogRecord, count-i)
		j := 0
		for ; i < count; i++ {
			r[i] = &storage[j]
			j++
		}
	}

	return r
}

// ReleaseSlice releases a slice of elements back to the pool.
func (p *logRecordPoolType) ReleaseSlice(slice []*LogRecord) {
	for _, elem := range slice {
		// Release nested attributes recursively to their pool.
		keyValuePool.ReleaseSlice(elem.attributes)

		// Zero-initialize the released element.
		*elem = LogRecord{}
	}

	p.mux.Lock()
	defer p.mux.Unlock()

	// Add the slice to the end of the pool.
	p.pool = append(p.pool, slice...)
}

// Release an element back to the pool.
func (p *logRecordPoolType) Release(elem *LogRecord) {
	// Release nested attributes recursively to their pool.
	keyValuePool.ReleaseSlice(elem.attributes)

	// Zero-initialize the released element.
	*elem = LogRecord{}

	p.mux.Lock()
	defer p.mux.Unlock()

	// Add the slice to the end of the pool.
	p.pool = append(p.pool, elem)
}

// ====================== Generated for message KeyValue ======================

type KeyValue struct {
	protoMessage lazyproto.ProtoMessage
	key          string
	value        string
}

func NewKeyValue(bytes []byte) *KeyValue {
	m := &KeyValue{protoMessage: lazyproto.ProtoMessage{Bytes: bytes}}
	m.decode()
	return m
}

func (m *KeyValue) GetKey() string {
	return m.key
}

func (m *KeyValue) GetValue() string {
	return m.value
}

func (m *KeyValue) decode() {
	buf := codec.NewBuffer(m.protoMessage.Bytes)

	// Iterate and decode the fields.
	molecule.MessageEach(
		buf, func(fieldNum int32, value molecule.Value) (bool, error) {
			switch fieldNum {
			case 1:
				// Decode key.
				v, err := value.AsStringUnsafe()
				if err != nil {
					return false, err
				}
				m.key = v
			case 2:
				// Decode value.
				v, err := value.AsStringUnsafe()
				if err != nil {
					return false, err
				}
				m.value = v
			}
			return true, nil
		},
	)
}

var preparedKeyValueKey = molecule.PrepareStringField(1)
var preparedKeyValueValue = molecule.PrepareStringField(2)

func (m *KeyValue) Marshal(ps *molecule.ProtoStream) error {
	if m.protoMessage.Flags&lazyproto.FlagsMessageModified != 0 {
		// Marshal key
		ps.StringPrepared(preparedKeyValueKey, m.key)
		// Marshal value
		ps.StringPrepared(preparedKeyValueValue, m.value)
	} else {
		// Message is unchanged. Used original bytes.
		ps.Raw(m.protoMessage.Bytes)
	}
	return nil
}

// Pool of KeyValue structs.
type keyValuePoolType struct {
	pool []*KeyValue
	mux  sync.Mutex
}

var keyValuePool = keyValuePoolType{}

// Get one element from the pool. Creates a new element if the pool is empty.
func (p *keyValuePoolType) Get() *KeyValue {
	p.mux.Lock()
	defer p.mux.Unlock()

	// Have elements in the pool?
	if len(p.pool) >= 1 {
		// Get the last element.
		r := p.pool[len(p.pool)-1]
		// Shrink the pool.
		p.pool = p.pool[:len(p.pool)-1]
		return r
	}

	// Pool is empty, create a new element.
	return &KeyValue{}
}

func (p *keyValuePoolType) GetSlice(count int) []*KeyValue {
	p.mux.Lock()
	defer p.mux.Unlock()

	// Have enough elements in the pool?
	if len(p.pool) >= count {
		// Cut the required slice from the end of the pool.
		r := p.pool[len(p.pool)-count:]
		// Shrink the pool.
		p.pool = p.pool[:len(p.pool)-count]
		return r
	}

	// Create a new slice.
	r := make([]*KeyValue, count)

	// Initialize with what remains in the pool.
	i := 0
	for ; i < len(p.pool); i++ {
		r[i] = p.pool[i]
	}
	p.pool = nil

	if i < count {
		// Create remaining elements.
		storage := make([]KeyValue, count-i)
		j := 0
		for ; i < count; i++ {
			r[i] = &storage[j]
			j++
		}
	}

	return r
}

// ReleaseSlice releases a slice of elements back to the pool.
func (p *keyValuePoolType) ReleaseSlice(slice []*KeyValue) {
	for _, elem := range slice {

		// Zero-initialize the released element.
		*elem = KeyValue{}
	}

	p.mux.Lock()
	defer p.mux.Unlock()

	// Add the slice to the end of the pool.
	p.pool = append(p.pool, slice...)
}

// Release an element back to the pool.
func (p *keyValuePoolType) Release(elem *KeyValue) {

	// Zero-initialize the released element.
	*elem = KeyValue{}

	p.mux.Lock()
	defer p.mux.Unlock()

	// Add the slice to the end of the pool.
	p.pool = append(p.pool, elem)
}
