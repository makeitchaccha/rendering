package timeline

type EntryBuilder struct {
	Entry
}

func NewEntryBuilder() *EntryBuilder {
	return &EntryBuilder{Entry: Entry{
		Series: make([]Series, 0),
	}}
}

func (b *EntryBuilder) AddSeries(s Series) *EntryBuilder {
	b.Series = append(b.Series, s)
	return b
}

func (b *EntryBuilder) Build() Entry {
	return b.Entry
}
