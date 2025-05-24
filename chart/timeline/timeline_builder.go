package timeline

type TimelineBuilder struct {
	Timeline
}

func NewTimelineBuilder() *TimelineBuilder {
	return &TimelineBuilder{
		Timeline: Timeline{
			entries: make([]Entry, 0),
		},
	}
}

func (b *TimelineBuilder) AddEntry(e Entry) *TimelineBuilder {
	b.entries = append(b.entries, e)
	return b
}

func (b *TimelineBuilder) Build() Timeline {
	return b.Timeline
}
