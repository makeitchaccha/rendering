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

func (b *EntryBuilder) AddSection(start, end float64) *EntryBuilder {
	b.sections = append(b.sections, section{start, end, ""})
	return b
}

func (b *EntryBuilder) AddSectionWithLabel(start, end float64, label string) *EntryBuilder {
	b.sections = append(b.sections, section{start, end, label})
	return b
}

func (b *EntryBuilder) Build() Entry {
	return b.Entry
}
