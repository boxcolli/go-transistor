package basicindex

import (
	"fmt"
	"testing"

	"github.com/boxcolli/go-transistor/index"
	"github.com/boxcolli/go-transistor/io"
	"github.com/boxcolli/go-transistor/io/bus/channelbus"
	"github.com/boxcolli/go-transistor/types"
	"github.com/stretchr/testify/assert"
)

var entries = []io.Bus{
	channelbus.NewChannelBus(1),
	channelbus.NewChannelBus(1),
}

var addTopic = []([]types.Topic){
	{
		types.Topic{"A0", "B0"},
		types.Topic{"A0", "B1"},
		types.Topic{"A0", "B2"},
	},
	{
		types.Topic{"A0"},
		types.Topic{"A0", "B1"},
		types.EmptyTopic,
	},
}

func TestBasicIndex(t *testing.T) {
	index := newBasicIndex()

	// Test one entry
	{
		fmt.Println("test one entry")

		i := 1

		// Add e[0]
		for _, topic := range addTopic[i] {
			index.Add(entries[i], topic)
			printInode(index.i)
			printVset(index.v)
		}


		// Del e[0]
		for _, topic := range addTopic[i] {
			index.Del(entries[i], topic)
			printInode(index.i)
			printVset(index.v)
		}

		printInode(index.i)
		printVset(index.v)

		assert.Zero(t, len(index.v))
	}

	// Test two entries
	{
		fmt.Println("test two entries")

		// Add
		for i := 0; i < len(entries); i++ {
			for _, topic := range addTopic[i] {
				index.Add(entries[i], topic)
			}
		}

		printInode(index.i)
		printVset(index.v)

		// Del
		for i := 0; i < len(entries); i++ {
			for _, topic := range addTopic[i] {
				index.Del(entries[i], topic)
			}
		}

		printInode(index.i)
		printVset(index.v)

		assert.Zero(t, len(index.v))
	}
}
func newBasicIndex() *basicIndex {
	return &basicIndex{
		i: index.NewInode(nil, ""),
		v: make(map[index.Entry]*index.Vnode),
	}
}
func _printIndent(in int) {
	for i := 0; i < in; i++ {
		fmt.Print("  ")
	}
}
func printInode(i *index.Inode) {
	fmt.Println("inode")
	_printInode(i, 0)
}
func _printInode(i *index.Inode, level int) {
	_printIndent(level); fmt.Printf("[%s]\n", i.Token)

	for e := range i.Eset {
		_printIndent(level); fmt.Printf("- %p\n", e)
	}
	for _, inext := range i.Next {
		_printInode(inext, level + 1)
	}
}
func printVset(vset map[index.Entry]*index.Vnode) {
	fmt.Println("vset")
	for e, v := range vset {
		printVnode(e, v)
	}
}
func printVnode(e index.Entry, v *index.Vnode) {
	fmt.Printf("vnode %p\n", e)
	_printVnode(v, 0)
}
func _printVnode(v *index.Vnode, level int) {
	_printIndent(level); fmt.Printf("[%s]\n", v.Token)

	for _, vnext := range v.Next {
		_printVnode(vnext, level + 1)
	}
}