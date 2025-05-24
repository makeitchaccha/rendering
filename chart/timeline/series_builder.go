package timeline

import "image/color"

type SeriesBuilder struct {
	Series
}

func NewSeriesBuilder(fillingFactor float64, color color.Color) *SeriesBuilder {
	return &SeriesBuilder{
		Series: Series{
			fillingFactor: fillingFactor,
			color:         color,
			sections:      make([]section, 0),
		},
	}
}

func (b *SeriesBuilder) SetColor(color color.Color) *SeriesBuilder {
	b.color = color
	return b
}

func (b *SeriesBuilder) SetLabelColor(color color.Color) *SeriesBuilder {
	b.labelColor = color
	return b
}

func (b *SeriesBuilder) SetFillingFactor(fillingFactor float64) *SeriesBuilder {
	b.fillingFactor = fillingFactor
	return b
}

func (b *SeriesBuilder) SetSections(sections []section) *SeriesBuilder {
	b.sections = sections
	return b
}

func (b *SeriesBuilder) AddSection(start, end float64, opts ...SectionOpt) *SeriesBuilder {
	section := section{start, end, 1, ""}

	section.apply(opts...)

	b.sections = append(b.sections, section)
	return b
}

func (b *SeriesBuilder) Build() Series {
	return b.Series
}
