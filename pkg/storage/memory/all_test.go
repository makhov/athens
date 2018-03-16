package memory

import (
	"testing"

	"github.com/gomods/athens/pkg/storage"
	"github.com/stretchr/testify/suite"
)

const (
	module = "my/module"
)

var (
	mod = []byte("123")
	zip = []byte("456")
)

type MemoryTests struct {
	suite.Suite
	mem storage.Storage
}

func (m *MemoryTests) SetupTest() {
	m.mem = New()
}

func TestMemoryStorage(t *testing.T) {
	suite.Run(t, new(MemoryTests))
}
