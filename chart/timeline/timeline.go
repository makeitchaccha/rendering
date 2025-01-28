package timeline

import (
	"errors"

	"github.com/fogleman/gg"
	"github.com/makeitchaccha/rendering/layout"
)

var (
	ErrNoEntries = errors.New("no entries")
)

type Timeline struct {
	fillingFactor float64
	entries       []Entry
}

var _ layout.Renderer = Timeline{}.Render

func (t Timeline) Render(dc *gg.Context, x, y, w, h float64) error {
	if len(t.entries) == 0 {
		return ErrNoEntries
	}
	l := len(t.entries)
	hEntry := h / float64(l)
	for i, e := range t.entries {
		yEntry := y + hEntry*float64(i) + (1-t.fillingFactor)*hEntry/2
		hFilling := t.fillingFactor * hEntry
		e.Render(dc, x, yEntry, w, hFilling)
	}
	return nil
}

func (t Timeline) GridType() layout.GridType {
	return layout.GridType{Rows: len(t.entries), Cols: 1}
}

func (t Timeline) RenderInGrid(dc *gg.Context, grid layout.Grid) error {
	if t.GridType() != grid.GridType {
		return errors.New("size mismatch")
	}

	for i, e := range t.entries {
		cell, err := grid.Cell(i, 0)
		if err != nil {
			return err
		}
		hFilling := t.fillingFactor * cell.H
		e.Render(dc, cell.X, cell.Y+(cell.H-hFilling)/2, cell.W, hFilling)
	}
	return nil
}
