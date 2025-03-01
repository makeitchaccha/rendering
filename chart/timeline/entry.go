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
	alpha      float64 // 0-1, 0 is transparent, 1 is opaque
	label      string
}

func (s *section) apply(opts ...SectionOpt) {
	for _, opt := range opts {
		opt(s)
	}
}

type SectionOpt func(*section)

func WithAlpha(alpha float64) SectionOpt {
	return func(s *section) {
		s.alpha = alpha
	}
}

func WithLabel(label string) SectionOpt {
	return func(s *section) {
		s.label = label
	}
}

var _ layout.Renderer = Entry{}.Render

func (e Entry) Render(dc *gg.Context, x, y, w, h float64) error {
	dc.Push()
	for _, s := range e.sections {
		r, g, b, _ := e.color.RGBA()

		dc.SetColor(color.NRGBA64{
			R: uint16(r),
			G: uint16(g),
			B: uint16(b),
			A: uint16(s.alpha * 0xffff),
		})

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
