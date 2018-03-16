package memory

import (
	"github.com/gomods/athens/pkg/storage"
)

func (v *storageImpl) Get(module, vsn string) (*storage.Version, error) {
	v.RLock()
	defer v.RUnlock()
	key := module
	versions := v.versions[key]
	for _, version := range versions {
		if version.RevInfo.Version == vsn {
			return version, nil
		}
	}
	return nil, &storage.ErrVersionNotFound{
		Module:  module,
		Version: vsn,
	}

}
