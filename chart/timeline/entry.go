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
		fillingFactor += s.FillingFactor
	}
	return fillingFactor
}

var _ layout.Renderer = Entry{}.Render

func (e Entry) Render(dc *gg.Context, x, y, w, h float64) error {

	yAnchor := y
	totalFillingFactor := e.TotalFillingFactor()
	dc.Push()
	for _, series := range e.Series {
		r, g, b, _ := series.Color.RGBA()
		dc.SetColor(color.NRGBA64{
			R: uint16(r),
			G: uint16(g),
			B: uint16(b),
			A: uint16(0xffff),
		})

		seriesH := h * series.FillingFactor / totalFillingFactor
		for _, section := range series.Sections {

			x0 := x + section.Start*w
			x1 := x + section.End*w
			dc.DrawRectangle(x0, yAnchor, x1-x0, seriesH)
			dc.Fill()
			if section.Label != "" {
				dc.Push()
				dc.SetColor(series.LabelColor)
				dc.DrawStringAnchored(section.Label, (x0+x1)/2, y+h/2, 0.5, 0.5)
				dc.Pop()
			}
		}

		yAnchor += seriesH

	}

	dc.Pop()
	return nil
}
