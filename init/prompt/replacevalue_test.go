package prompt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"git.bytecode.nl/bytecode/genesis/init/prompt"
)

var unsorted = prompt.ReplaceValues{
	{OldValue: "b", Description: "b", ReplaceOrder: 2},
	{OldValue: "e", Description: "e", ReplaceOrder: 5},
	{OldValue: "a", Description: "a", ReplaceOrder: 1},
	{OldValue: "d", Description: "d", ReplaceOrder: 4},
	{OldValue: "c", Description: "c", ReplaceOrder: 3},
}

var sorted = prompt.ReplaceValues{
	{OldValue: "a", Description: "a", ReplaceOrder: 1},
	{OldValue: "b", Description: "b", ReplaceOrder: 2},
	{OldValue: "c", Description: "c", ReplaceOrder: 3},
	{OldValue: "d", Description: "d", ReplaceOrder: 4},
	{OldValue: "e", Description: "e", ReplaceOrder: 5},
}

func TestReplaceValueSort(t *testing.T) {
	got := unsorted.Sort()
	assert.Equal(t, sorted, got)
}
