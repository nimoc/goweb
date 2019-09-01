package main

import (
	"github.com/og/goweb/action"
	"github.com/og/goweb/core"
)
func main () {
	app := core.App{}
	action.Start(app)
}