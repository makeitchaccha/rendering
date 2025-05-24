package timeline

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/makeitchaccha/rendering/layout"
)

type Entry struct {
	Series []Series
}

func (e *Entry) TotalFillingFactor() float64 {
	if len(e.Series) == 0 {
		return 0
	}
	fillingFactor := 0.0
	for _, s := range e.Series {
		fillingFactor += s.fillingFactor
	}
	return fillingFactor
}

var _ layout.Renderer = Entry{}.Render

func (e Entry) Render(dc *gg.Context, x, y, w, h float64) error {

	yAnchor := y
	dc.Push()
	for _, series := range e.Series {
		r, g, b, _ := series.color.RGBA()
		dc.SetColor(color.NRGBA64{
			R: uint16(r),
			G: uint16(g),
			B: uint16(b),
			A: uint16(0xffff),
		})

		for _, section := range series.sections {

			x0 := x + section.start*w
			x1 := x + section.end*w
			dc.DrawRectangle(x0, yAnchor, x1-x0, h)
			dc.Fill()
			if section.label != "" {
				dc.Push()
				dc.SetColor(series.labelColor)
				dc.DrawStringAnchored(section.label, (x0+x1)/2, y+h/2, 0.5, 0.5)
				dc.Pop()
			}
		}

		yAnchor += h * series.fillingFactor

	}

	dc.Pop()
	return nil
}
