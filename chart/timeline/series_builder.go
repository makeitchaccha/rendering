package timeline

import "image/color"

type SeriesBuilder struct {
	Series
}

func NewSeriesBuilder(fillingFactor float64, color color.Color) *SeriesBuilder {
	return &SeriesBuilder{
		Series: Series{
			FillingFactor: fillingFactor,
			Color:         color,
			Sections:      make([]Section, 0),
		},
	}
}

func (b *SeriesBuilder) SetColor(color color.Color) *SeriesBuilder {
	b.Color = color
	return b
}

func (b *SeriesBuilder) SetLabelColor(color color.Color) *SeriesBuilder {
	b.LabelColor = color
	return b
}

func (b *SeriesBuilder) SetFillingFactor(fillingFactor float64) *SeriesBuilder {
	b.FillingFactor = fillingFactor
	return b
}

func (b *SeriesBuilder) SetSections(sections []Section) *SeriesBuilder {
	b.Sections = sections
	return b
}

func (b *SeriesBuilder) AddSection(start, end float64, opts ...SectionOpt) *SeriesBuilder {
	section := Section{start, end, 1, ""}

	section.apply(opts...)

	b.Sections = append(b.Sections, section)
	return b
}

func (b *SeriesBuilder) Build() Series {
	return b.Series
}
