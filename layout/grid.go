package layout

import (
	"errors"
	"iter"

	"github.com/fogleman/gg"
)

var (
	ErrOutOfBounds = errors.New("out of bounds")
)

type GridType struct {
	Rows, Cols int
}

type Grid struct {
	GridType
	cellYs, cellXs []float64 // cellXs has col+1 elements, cellYs has row+1 elements.
}

type Cell struct {
	X, Y, W, H float64
}

func NewGrid(x, y float64, cellWidths, cellHeights []float64) Grid {
	cellXs := make([]float64, len(cellWidths)+1)
	cellYs := make([]float64, len(cellHeights)+1)
	cellXs[0] = x
	cellYs[0] = y
	for i, w := range cellWidths {
		cellXs[i+1] = cellXs[i] + w
	}
	for i, h := range cellHeights {
		cellYs[i+1] = cellYs[i] + h
	}
	return Grid{
		GridType: GridType{
			Rows: len(cellHeights),
			Cols: len(cellWidths),
		},
		cellXs: cellXs,
		cellYs: cellYs,
	}
}

func (g Grid) IsInBounds(i, j int) bool {
	return 0 <= i && i < g.Rows && 0 <= j && j < g.Cols
}

func (g Grid) Cell(i, j int) (Cell, error) {
	if !g.IsInBounds(i, j) {
		return Cell{}, ErrOutOfBounds
	}
	return Cell{
		X: g.cellXs[j],
		Y: g.cellYs[i],
		W: g.cellXs[j+1] - g.cellXs[j],
		H: g.cellYs[i+1] - g.cellYs[i],
	}, nil
}

func (g Grid) Subgrid(i1, j1, i2, j2 int) (Grid, error) {
	if !g.IsInBounds(i1, j1) || !g.IsInBounds(i2, j2) {
		return Grid{}, ErrOutOfBounds
	}
	return Grid{
		GridType: GridType{
			Rows: i2 - i1 + 1,
			Cols: j2 - j1 + 1,
		},
		cellXs: g.cellXs[j1 : j2+2],
		cellYs: g.cellYs[i1 : i2+2],
	}, nil
}

func (g Grid) RowAsSubgrid(i int) (Grid, error) {
	return g.Subgrid(i, 0, i, g.Cols-1)
}

func (g Grid) ColAsSubgrid(j int) (Grid, error) {
	return g.Subgrid(0, j, g.Rows-1, j)
}

func (g Grid) CellRenderFunc(i, j int) (RendererFunc, error) {
	if !g.IsInBounds(i, j) {
		return nil, ErrOutOfBounds
	}

	return func(dc *gg.Context, r Renderer) error {
		x := g.cellXs[j]
		y := g.cellYs[i]
		w := g.cellXs[j+1] - x
		h := g.cellYs[i+1] - y
		return r(dc, x, y, w, h)
	}, nil
}

func (g Grid) JointCellRenderFunc(i1, j1, i2, j2 int) (RendererFunc, error) {
	if !g.IsInBounds(i1, j1) || !g.IsInBounds(i2, j2) {
		return nil, ErrOutOfBounds
	}

	return func(dc *gg.Context, r Renderer) error {
		x := g.cellXs[j1]
		y := g.cellYs[i1]
		w := g.cellXs[j2+1] - x
		h := g.cellYs[i2+1] - y
		return r(dc, x, y, w, h)
	}, nil
}

type CellIndex struct{ I, J int }

var _ iter.Seq2[CellIndex, RendererFunc] = Grid{}.ForEachCellRenderFunc

func (g Grid) ForEachCellRenderFunc(f func(pos CellIndex, renderFunc RendererFunc) bool) {
	for i := 0; i < g.Rows; i++ {
		for j := 0; j < g.Cols; j++ {
			renderFunc, _ := g.CellRenderFunc(i, j)
			if !f(CellIndex{i, j}, renderFunc) {
				return
			}
		}
	}
}
