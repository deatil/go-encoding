package encoding

import (
    "sync"
)

type DataName interface {
    ~uint | ~int | ~string
}

/**
 * Data Set
 *
 * @create 2023-3-31
 * @author deatil
 */
type DataSet[N DataName, M any] struct {
    // 锁定
    mu sync.RWMutex

    // 数据
    data map[N]func() M
}

// NewDataSet
func NewDataSet[N DataName, M any]() *DataSet[N, M] {
    return &DataSet[N, M]{
        data: make(map[N]func() M),
    }
}

// Add
func (this *DataSet[N, M]) Add(name N, data func() M) *DataSet[N, M] {
    this.mu.Lock()
    defer this.mu.Unlock()

    this.data[name] = data

    return this
}

// Has
func (this *DataSet[N, M]) Has(name N) bool {
    this.mu.RLock()
    defer this.mu.RUnlock()

    if _, ok := this.data[name]; ok {
        return true
    }

    return false
}

// Get
func (this *DataSet[N, M]) Get(name N) func() M {
    this.mu.RLock()
    defer this.mu.RUnlock()

    if data, ok := this.data[name]; ok {
        return data
    }

    return nil
}

// Remove
func (this *DataSet[N, M]) Remove(name N) *DataSet[N, M] {
    this.mu.Lock()
    defer this.mu.Unlock()

    delete(this.data, name)

    return this
}

// Names
func (this *DataSet[N, M]) Names() []N {
    names := make([]N, 0)
    for name, _ := range this.data {
        names = append(names, name)
    }

    return names
}

// All
func (this *DataSet[N, M]) All() map[N]func() M {
    return this.data
}

// Clean
func (this *DataSet[N, M]) Clean() {
    this.mu.Lock()
    defer this.mu.Unlock()

    for name, _ := range this.data {
        delete(this.data, name)
    }
}

// Len
func (this *DataSet[N, M]) Len() int {
    return len(this.data)
}
