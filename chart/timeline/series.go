package timeline

import "image/color"

type Series struct {
	fillingFactor float64
	color         color.Color
	labelColor    color.Color
	sections      []section
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
