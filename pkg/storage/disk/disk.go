package disk

import (
	"path/filepath"

	"github.com/gomods/athens/pkg/storage"
)

// Storage is the only interface defined by the disk storage. Use
// NewStorage to create one of these. Everything is all in one
// because it all has to share the same tree
type Storage interface {
	storage.Lister
	storage.Getter
	storage.Saver
}

type storageImpl struct {
	rootDir string
}

func (s *storageImpl) moduleDiskLocation(module string) string {
	return filepath.Join(s.rootDir, module)
}

func (s *storageImpl) versionDiskLocation(module, version string) string {
	return filepath.Join(s.moduleDiskLocation(module), version)

}

// NewStorage returns a new ListerSaver implementation that stores
// everything under rootDir
func NewStorage(rootDir string) Storage {
	return &storageImpl{rootDir: rootDir}

}
