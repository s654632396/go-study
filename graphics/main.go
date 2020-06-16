package main

import (
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/math/f64"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"image"
	"image/draw"
	"log"
)

func main() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(nil)
		if err != nil {
			handleError(err)
			return
		}
		defer w.Release()

		aff3 := f64.Aff3{1,1,1,1,1,1}
		var t screen.Texture
		t, err = s.NewTexture(image.Pt(100, 200))
		rect := image.Rect(0, 0, 1000, 1000)


		for {
			switch e := w.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
				log.Println(e.From, e.To)
			case mouse.Event:
				// e.Direction : press or release
				log.Println("pointer is", e.Direction, e.X, e.Y)
			case key.Event:
				// e.Direction : press or release
				log.Println("key is", e.String(), e.Direction, e.Code, e.Modifiers)
			default:
				log.Printf("e.T = %T\n", e)
				w.Draw(aff3, t,  rect, draw.Op(1),nil)

			}
		}
	})
}

func handleError(err error) {

}