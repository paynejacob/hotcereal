package graph_test

import (
	"github.com/paynejacob/hotcereal/pkg/graph"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenize(t *testing.T) {
	// empty string
	{
		in := ""
		out := graph.Tokenize(in)

		assert.Equal(t, []string{}, out)
	}

	// 1 token
	{
		in := "abc"
		out := graph.Tokenize(in)

		assert.Equal(t, []string{"abc"}, out)
	}

	// 2 token
	{
		in := "abc def"
		out := graph.Tokenize(in)

		assert.Equal(t, []string{"abc", "def"}, out)
	}

	// special characters
	{
		in := "abc, def"
		out := graph.Tokenize(in)

		assert.Equal(t, []string{"abc", "def"}, out)
	}

	// capital letters
	{
		in := "aBc"
		out := graph.Tokenize(in)

		assert.Equal(t, []string{"abc"}, out)

		in = "abC DeF"
		out = graph.Tokenize(in)

		assert.Equal(t, []string{"abc", "def"}, out)
	}
}
