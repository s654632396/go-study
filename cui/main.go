package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"log"
	"strings"
)

var (
	hx int = 1
	hy int = 1
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	g.SetManagerFunc(myLayout)
	g.BgColor = gocui.ColorDefault
	g.FgColor = gocui.ColorCyan

	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("help")
		if err != nil {
			// handle error
		}
		v.Clear()
		outputH(v)
		return nil
	})

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func myLayout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	v, err := g.SetView("help", 0, 0, maxX-1, maxY-1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		outputH(v)
	}
	return nil

}

func outputH(v *gocui.View) {
	fmt.Fprintln(v, strings.Repeat("x", 10))
}
