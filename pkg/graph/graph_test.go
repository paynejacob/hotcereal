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

		assert.ElementsMatch(t, []string{"z"}, g.Search("abc"))
		assert.ElementsMatch(t, []string{"z"}, g.Search("ab"))
		assert.ElementsMatch(t, []string{"z"}, g.Search("a"))
		assert.ElementsMatch(t, []string{"z"}, g.Search(""))
	}

	// Write with no overlap
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")
		g.Write([]string{"def"}, "y")

		assert.ElementsMatch(t, []string{"z"}, g.Search("abc"))
		assert.ElementsMatch(t, []string{"z"}, g.Search("ab"))
		assert.ElementsMatch(t, []string{"z"}, g.Search("a"))

		assert.ElementsMatch(t, []string{"y"}, g.Search("def"))
		assert.ElementsMatch(t, []string{"y"}, g.Search("de"))
		assert.ElementsMatch(t, []string{"y"}, g.Search("d"))
		assert.ElementsMatch(t, []string{"z", "y"}, g.Search(""))
	}

	// Write with overlap
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")
		g.Write([]string{"abcd"}, "y")

		assert.ElementsMatch(t, []string{"y"}, g.Search("abcd"))
		assert.ElementsMatch(t, []string{"z", "y"}, g.Search("abc"))
		assert.ElementsMatch(t, []string{"z", "y"}, g.Search("ab"))
		assert.ElementsMatch(t, []string{"y", "z"}, g.Search("a"))
		assert.ElementsMatch(t, []string{"z", "y"}, g.Search(""))
	}

	// Write twice
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")
		g.Write([]string{"abc"}, "z")

		assert.ElementsMatch(t, []string{"z"}, g.Search("abc"))
		assert.ElementsMatch(t, []string{"z"}, g.Search("ab"))
		assert.ElementsMatch(t, []string{"z"}, g.Search("a"))
		assert.ElementsMatch(t, []string{"z"}, g.Search(""))
	}
}

func TestGraph_Search(t *testing.T) {
	var g *graph.Graph

	// search empty graph
	{
		g = graph.New()

		assert.ElementsMatch(t, []string{}, g.Search("a"))
		assert.ElementsMatch(t, []string{}, g.Search("ab"))
		assert.ElementsMatch(t, []string{}, g.Search(""))
	}

	// search empty query
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")

		assert.ElementsMatch(t, []string{"z"}, g.Search(""))
	}

	// search partial match
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")

		assert.ElementsMatch(t, []string{"z"}, g.Search("ab"))
		assert.ElementsMatch(t, []string{"z"}, g.Search("a"))
	}

	// search full match
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")

		assert.ElementsMatch(t, []string{"z"}, g.Search("abc"))
	}

	// search no match
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")

		assert.ElementsMatch(t, []string{}, g.Search("def"))
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

		assert.ElementsMatch(t, []string{"z"}, g.Search("abc"))
		assert.ElementsMatch(t, []string{"z"}, g.Search("ab"))
		assert.ElementsMatch(t, []string{"z"}, g.Search("a"))
		assert.ElementsMatch(t, []string{"z"}, g.Search(""))
	}

	// match on all nodes
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")
		g.Write([]string{"abc"}, "y")

		g.Delete("y")

		assert.ElementsMatch(t, []string{"z"}, g.Search("abc"))
		assert.ElementsMatch(t, []string{"z"}, g.Search("ab"))
		assert.ElementsMatch(t, []string{"z"}, g.Search("a"))
		assert.ElementsMatch(t, []string{"z"}, g.Search(""))
	}

	// match head nodes
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")
		g.Write([]string{"ab"}, "y")

		g.Delete("y")

		assert.ElementsMatch(t, []string{"z"}, g.Search("abc"))
		assert.ElementsMatch(t, []string{"z"}, g.Search("ab"))
		assert.ElementsMatch(t, []string{"z"}, g.Search("a"))
		assert.ElementsMatch(t, []string{"z"}, g.Search(""))
	}

	// match tail nodes
	{
		g = graph.New()
		g.Write([]string{"abc"}, "z")
		g.Write([]string{"abcd"}, "y")

		g.Delete("y")

		assert.ElementsMatch(t, []string{}, g.Search("abcd"))
		assert.ElementsMatch(t, []string{"z"}, g.Search("abc"))
		assert.ElementsMatch(t, []string{"z"}, g.Search("ab"))
		assert.ElementsMatch(t, []string{"z"}, g.Search("a"))
		assert.ElementsMatch(t, []string{"z"}, g.Search(""))
	}
}
