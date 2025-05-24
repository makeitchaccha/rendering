package timeline

import "image/color"

type Series struct {
	FillingFactor float64
	Color         color.Color
	LabelColor    color.Color
	Sections      []Section
}

type Section struct {
	Start, End float64
	Alpha      float64 // 0-1, 0 is transparent, 1 is opaque
	Label      string
}

func (s *Section) apply(opts ...SectionOpt) {
	for _, opt := range opts {
		opt(s)
	}
}

type SectionOpt func(*Section)

func WithAlpha(alpha float64) SectionOpt {
	return func(s *Section) {
		s.Alpha = alpha
	}
}

func WithLabel(label string) SectionOpt {
	return func(s *Section) {
		s.Label = label
	}
}
