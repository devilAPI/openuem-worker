package models

import (
	"context"
	"testing"

	"github.com/open-uem/ent/enttest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

func TestIsLibraryPackage(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:winget?mode=memory&_fk=1")
	defer client.Close()
	m := Model{Client: client}

	client.SoftwarePackage.Create().SetPackageID("Example.Community").SetName("Example").SetSource("winget").ExecX(ctx)
	client.SoftwarePackage.Create().SetPackageID("ExampleCorp.Tool").SetName("Example Tool").SetSource(LibraryPackageSource).ExecX(ctx)

	isLibrary, err := m.IsLibraryPackage("ExampleCorp.Tool")
	require.NoError(t, err)
	assert.True(t, isLibrary)

	isLibrary, err = m.IsLibraryPackage("Example.Community")
	require.NoError(t, err)
	assert.False(t, isLibrary)

	isLibrary, err = m.IsLibraryPackage("Unknown.Package")
	require.NoError(t, err)
	assert.False(t, isLibrary)
}
