package memory

import (
	"sync"

	"github.com/gomods/athens/pkg/storage"
)

type storageImpl struct {
	*sync.RWMutex
	versions map[string][]*storage.Version
}

func (e *storageImpl) key(module string) string {
	return module
}

// New creates a new in-memory storage implementation
func New() storage.Storage {
	return &storageImpl{
		RWMutex:  new(sync.RWMutex),
		versions: make(map[string][]*storage.Version),
	}
}
