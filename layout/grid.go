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

type Cell Rectangle

type CellIndex struct {
	Row int
	Col int
}

var _ iter.Seq2[CellIndex, RendererFunc] = Grid{}.ForEachCellRenderFunc

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

func (g Grid) Bounds() Rectangle {
	return Rectangle{
		Min: Point{X: g.cellXs[0], Y: g.cellYs[0]},
		Max: Point{X: g.cellXs[g.Cols], Y: g.cellYs[g.Rows]},
	}
}

func (g Grid) IsInBounds(row, col int) bool {
	return 0 <= row && row < g.Rows && 0 <= col && col < g.Cols
}

func (g Grid) Cell(row, col int) (Rectangle, error) {
	if !g.IsInBounds(row, col) {
		return Rectangle{}, ErrOutOfBounds
	}
	return Rectangle{
		Min: Point{X: g.cellXs[col], Y: g.cellYs[row]},
		Max: Point{X: g.cellXs[col+1], Y: g.cellYs[row+1]},
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

func (g Grid) RowAsSubgrid(row int) (Grid, error) {
	return g.Subgrid(row, 0, row, g.Cols-1)
}

func (g Grid) ColAsSubgrid(col int) (Grid, error) {
	return g.Subgrid(0, col, g.Rows-1, col)
}

func (g Grid) CellRenderFunc(row, col int) (RendererFunc, error) {
	if !g.IsInBounds(row, col) {
		return nil, ErrOutOfBounds
	}

	return func(dc *gg.Context, r Renderer) error {
		x := g.cellXs[col]
		y := g.cellYs[row]
		w := g.cellXs[col+1] - x
		h := g.cellYs[row+1] - y
		return r(dc, x, y, w, h)
	}, nil
}

func (g Grid) JointCellRenderFunc(row1, col1, row2, col2 int) (RendererFunc, error) {
	if !g.IsInBounds(row1, col1) || !g.IsInBounds(row2, col2) {
		return nil, ErrOutOfBounds
	}

	return func(dc *gg.Context, r Renderer) error {
		x := g.cellXs[col1]
		y := g.cellYs[row1]
		w := g.cellXs[col2+1] - x
		h := g.cellYs[row2+1] - y
		return r(dc, x, y, w, h)
	}, nil
}

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
