package memory

import (
	"github.com/gomods/athens/pkg/storage"
)

func (v *storageImpl) List(module string) ([]string, error) {
	key := v.key(module)
	v.RLock()
	defer v.RUnlock()
	versions, ok := v.versions[key]
	if !ok {
		return nil, storage.NotFoundErr{Module: module}
	}
	ret := make([]string, len(versions))
	for i, version := range versions {
		ret[i] = version.RevInfo.Version
	}
	return ret, nil
}
