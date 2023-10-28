package main

import (
	"github.com/boxcolli/go-transistor/base"
	"github.com/boxcolli/go-transistor/base/basicbase"
	"github.com/boxcolli/go-transistor/collector"
	"github.com/boxcolli/go-transistor/collector/basiccollector"
)

func newCollector(b base.Base) collector.Collector {
	return basiccollector.NewBasicCollector(b)
}

func newBase() base.Base {
	return basicbase.NewBasicBase()
}

