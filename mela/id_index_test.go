package mela_test

import (
	"fmt"
	"testing"

	"github.com/kdisneur/obsidiantomela/mela"
	"github.com/kdisneur/obsidiantomela/testutil/require"
)

func TestIDIndexAddGet(t *testing.T) {
	idIndex := make(mela.IDIndex)

	idIndex.AddName("name-1")
	idIndex.AddName("name-2")
	idIndex.AddName("name-3")

	id1, found1 := idIndex.ID("name-1")
	id1Bis, found1Bis := idIndex.ID("name-1")
	require.True(t, found1, "id of name-1 not found")
	require.True(t, found1Bis, "id of name-1 not found")
	require.NotEmpty(t, id1, "id of name-1 should be present")
	require.Equal(t, id1, id1Bis, "id of name-1 should be consistent")

	id2, found2 := idIndex.ID("name-2")
	id2Bis, found2Bis := idIndex.ID("name-2")
	require.True(t, found2, "id of name-2 not found")
	require.True(t, found2Bis, "id of name-2 not found")
	require.NotEmpty(t, id2, "id of name-2 should be present")
	require.Equal(t, id2, id2Bis, "id of name-2 should be consistent")

	id3, found3 := idIndex.ID("name-3")
	id3Bis, found3Bis := idIndex.ID("name-3")
	require.True(t, found3, "id of name-3 not found")
	require.True(t, found3Bis, "id of name-3 not found")
	require.NotEmpty(t, id3, "id of name-3 should be present")
	require.Equal(t, id3, id3Bis, "id of name-3 should be consistent")

	_, found4 := idIndex.ID("name-4")
	require.False(t, found4, "id of name-4 should not exist")
}

func TestIDIndexLink(t *testing.T) {
	idIndex := make(mela.IDIndex)

	idIndex.AddName("name-1")
	idIndex.AddName("name-2")
	idIndex.AddName("name-3")

	id2, found2 := idIndex.ID("name-2")
	require.True(t, found2, "id of name-2 not found")
	require.NotEmpty(t, id2, "id of name-2 should be present")

	link2, ok2 := idIndex.LinkName("name-2")
	require.True(t, ok2, "link to name-2 not found")
	require.Equal(
		t,
		fmt.Sprintf("[name-2](mela://recipe/%s)", id2),
		link2,
		"link to name-2 should be valid",
	)

	_, ok4 := idIndex.LinkName("name-4")
	require.False(t, ok4, "link to name-4 should not exist")
}
