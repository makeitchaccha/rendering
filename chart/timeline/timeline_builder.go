package timeline

type TimelineBuilder struct {
	Timeline
}

func NewTimelineBuilder() *TimelineBuilder {
	return &TimelineBuilder{
		Timeline: Timeline{
			fillingFactor: 0.8,
		},
	}
}

func (b *TimelineBuilder) SetFillingFactor(f float64) *TimelineBuilder {
	b.fillingFactor = f
	return b
}

func (b *TimelineBuilder) AddEntry(e Entry) *TimelineBuilder {
	b.entries = append(b.entries, e)
	return b
}

func (b *TimelineBuilder) Build() Timeline {
	return b.Timeline
}
