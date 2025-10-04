package storage_test

import (
	"monitoring-demo/internal/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapStorage_SaveAndLoad(t *testing.T) {
	st := storage.NewMapStorage()
	code := "testCode"
	url := "https://google.com"

	st.Save(code, url)

	got, ok := st.Load(code)
	assert.True(t, ok)
	assert.Equal(t, url, got)

	_, ok = st.Load("DimaLox")
	assert.False(t, ok)
}
