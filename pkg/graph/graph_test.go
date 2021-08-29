package graph_test

import (
	"github.com/paynejacob/hotcereal/pkg/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := graph.New()

	assert.NotNil(t, g)
}

func TestGraph_Write(t *testing.T) {
	var g *graph.Graph

	// Write to empty graph
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")

		assert.Equal(t, []string{"z"}, g.Search("abc"))
		assert.Equal(t, []string{"z"}, g.Search("ab"))
		assert.Equal(t, []string{"z"}, g.Search("a"))
		assert.Equal(t, []string{"z"}, g.Search(""))
	}

	// Write with no overlap
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")
		g.Write([]string{"def"}, "y")

		assert.Equal(t, []string{"z"}, g.Search("abc"))
		assert.Equal(t, []string{"z"}, g.Search("ab"))
		assert.Equal(t, []string{"z"}, g.Search("a"))

		assert.Equal(t, []string{"y"}, g.Search("def"))
		assert.Equal(t, []string{"y"}, g.Search("de"))
		assert.Equal(t, []string{"y"}, g.Search("f"))
		assert.Equal(t, []string{"z", "y"}, g.Search(""))
	}

	// Write with overlap
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")
		g.Write([]string{"abcd"}, "y")

		assert.Equal(t, []string{"y"}, g.Search("abcd"))
		assert.Equal(t, []string{"z", "y"}, g.Search("abc"))
		assert.Equal(t, []string{"z", "y"}, g.Search("ab"))
		assert.Equal(t, []string{"z", "y"}, g.Search("a"))
		assert.Equal(t, []string{"z", "y"}, g.Search(""))
	}

	// Write twice
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")
		g.Write([]string{"abc"}, "z")

		assert.Equal(t, []string{"z"}, g.Search("abc"))
		assert.Equal(t, []string{"z"}, g.Search("ab"))
		assert.Equal(t, []string{"z"}, g.Search("a"))
		assert.Equal(t, []string{"z"}, g.Search(""))
	}
}

func TestGraph_Search(t *testing.T) {
	var g *graph.Graph

	// search empty graph
	{
		g = graph.New()

		assert.Equal(t, []string{}, g.Search("a"))
		assert.Equal(t, []string{}, g.Search("ab"))
		assert.Equal(t, []string{}, g.Search(""))
	}

	// search empty query
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")

		assert.Equal(t, []string{"z"}, g.Search(""))
	}

	// search partial match
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")

		assert.Equal(t, []string{"z"}, g.Search("ab"))
		assert.Equal(t, []string{"z"}, g.Search("a"))
	}

	// search full match
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")

		assert.Equal(t, []string{"z"}, g.Search("abc"))
	}

	// search no match
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")

		assert.Equal(t, []string{}, g.Search("def"))
	}
}

func TestGraph_Delete(t *testing.T) {
	var g *graph.Graph

	// empty graph
	{
		g = graph.New()

		g.Delete("a")
	}

	// missing value
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")

		g.Delete("y")

		assert.Equal(t, []string{"z"}, g.Search("abc"))
		assert.Equal(t, []string{"z"}, g.Search("ab"))
		assert.Equal(t, []string{"z"}, g.Search("a"))
		assert.Equal(t, []string{"z"}, g.Search(""))
	}

	// match on all nodes
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")
		g.Write([]string{"abc"}, "y")

		g.Delete("y")

		assert.Equal(t, []string{"z"}, g.Search("abc"))
		assert.Equal(t, []string{"z"}, g.Search("ab"))
		assert.Equal(t, []string{"z"}, g.Search("a"))
		assert.Equal(t, []string{"z"}, g.Search(""))
	}

	// match head nodes
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")
		g.Write([]string{"ab"}, "y")

		g.Delete("y")

		assert.Equal(t, []string{"z"}, g.Search("abc"))
		assert.Equal(t, []string{"z"}, g.Search("ab"))
		assert.Equal(t, []string{"z"}, g.Search("a"))
		assert.Equal(t, []string{"z"}, g.Search(""))
	}

	// match tail nodes
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")
		g.Write([]string{"abcd"}, "y")

		g.Delete("y")

		assert.Equal(t, []string{}, g.Search("abcd"))
		assert.Equal(t, []string{"z"}, g.Search("abc"))
		assert.Equal(t, []string{"z"}, g.Search("ab"))
		assert.Equal(t, []string{"z"}, g.Search("a"))
		assert.Equal(t, []string{"z"}, g.Search(""))
	}
}
