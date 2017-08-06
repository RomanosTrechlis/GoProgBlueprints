package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPath(t *testing.T) {

	var p *Path

	p = NewPath("/polls")
	require.False(t, p.HasID())
	require.Equal(t, "polls", p.Path)

	p = NewPath("/polls/")
	require.False(t, p.HasID())
	require.Equal(t, "polls", p.Path)

	p = NewPath("polls/")
	require.False(t, p.HasID())
	require.Equal(t, "polls", p.Path)

	p = NewPath("polls")
	require.False(t, p.HasID())
	require.Equal(t, "polls", p.Path)

	p = NewPath("/polls/1")
	require.True(t, p.HasID())
	require.Equal(t, "polls", p.Path)
	require.Equal(t, "1", p.ID)

	p = NewPath("/polls/1/")
	require.True(t, p.HasID())
	require.Equal(t, "polls", p.Path)
	require.Equal(t, "1", p.ID)

	p = NewPath("polls/1/")
	require.True(t, p.HasID())
	require.Equal(t, "polls", p.Path)
	require.Equal(t, "1", p.ID)

	p = NewPath("polls/1")
	require.True(t, p.HasID())
	require.Equal(t, "polls", p.Path)
	require.Equal(t, "1", p.ID)

}
