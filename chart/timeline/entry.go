package timeline

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/makeitchaccha/rendering/layout"
)

type Entry struct {
	color      color.Color
	labelColor color.Color
	sections   []section
}

type section struct {
	start, end float64
	label      string
}

var _ layout.Renderer = Entry{}.Render

func (e Entry) Render(dc *gg.Context, x, y, w, h float64) error {
	dc.Push()
	dc.SetColor(e.color)
	for _, s := range e.sections {
		x0 := x + s.start*w
		x1 := x + s.end*w
		dc.DrawRectangle(x0, y, x1-x0, h)
		dc.Fill()
		if s.label != "" {
			dc.Push()
			dc.SetColor(e.labelColor)
			dc.DrawStringAnchored(s.label, (x0+x1)/2, y+h/2, 0.5, 0.5)
			dc.Pop()
		}
	}
	dc.Pop()
	return nil
}
