package layout

import "github.com/fogleman/gg"

type Renderer func(dc *gg.Context, x, y, w, h float64) error

type RendererFunc func(dc *gg.Context, rendering Renderer) error
