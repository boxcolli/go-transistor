package simplecore

import "github.com/boxcolli/go-transistor/core"

type simpleCore struct {

}

func NewSimpleCore() core.Core {
	return &simpleCore{}
}
