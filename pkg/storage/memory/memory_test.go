package memory

func (m *MemoryTests) TestNew() {
	r := m.Require()
	iface := New()
	storage, ok := iface.(*storageImpl)
	r.True(ok)
	r.NotNil(storage.versions)
	r.NotNil(storage.RWMutex)
}
