package timeline

import "image/color"

type EntryBuilder struct {
	Entry
}

func NewEntryBuilder(color color.Color) *EntryBuilder {
	return &EntryBuilder{Entry{color: color}}
}

func (b *EntryBuilder) SetColor(color color.Color) *EntryBuilder {
	b.color = color
	return b
}

func (b *EntryBuilder) SetLabelColor(color color.Color) *EntryBuilder {
	b.labelColor = color
	return b
}

func (b *EntryBuilder) AddSection(start, end float64, opts ...SectionOpt) *EntryBuilder {
	section := section{start, end, 1, ""}

	section.apply(opts...)

	b.sections = append(b.sections, section)
	return b
}

func (b *EntryBuilder) Build() Entry {
	return b.Entry
}
