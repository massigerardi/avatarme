package avatarme

import (
	"code.google.com/p/draw2d/draw2d"
	"image"
	"image/color"
	"image/draw"
)

type Shape struct {
	vertices []image.Point
}

func newShape(verts []image.Point) Shape {
	return Shape{
		vertices: verts,
	}
}

var (
	v0  = image.Point{0, 0}
	v10 = image.Point{0, 2}
	v20 = image.Point{0, 4}
	v6  = image.Point{1, 1}
	v11 = image.Point{1, 2}
	v16 = image.Point{1, 3}

	v2  = image.Point{2, 0}
	v12 = image.Point{2, 2}
	v22 = image.Point{2, 4}
	v8  = image.Point{3, 1}
	v13 = image.Point{3, 2}
	v18 = image.Point{3, 3}
	v4  = image.Point{4, 0}
	v14 = image.Point{4, 2}
	v24 = image.Point{4, 4}

	Shapes = []Shape{
		newShape([]image.Point{v0, v4, v24, v20, v0}),
		newShape([]image.Point{v0, v4, v20, v0}),
		newShape([]image.Point{v2, v24, v20, v2}),
		newShape([]image.Point{v0, v2, v20, v22, v0}),
		newShape([]image.Point{v2, v14, v22, v10, v2}),
		newShape([]image.Point{v0, v14, v24, v22, v0}),
		newShape([]image.Point{v2, v24, v22, v13, v11, v22, v20, v2}),
		newShape([]image.Point{v0, v14, v22, v0}),
		newShape([]image.Point{v6, v8, v18, v16, v6}),
		newShape([]image.Point{v4, v20, v10, v12, v2, v4}),
		newShape([]image.Point{v0, v2, v12, v10, v0}),
		newShape([]image.Point{v10, v14, v22, v10}),
		newShape([]image.Point{v20, v12, v24, v20}),
		newShape([]image.Point{v10, v2, v12, v10}),
		newShape([]image.Point{v0, v2, v10, v0})}

	ShapeHorMirrors = []Shape{
		newShape([]image.Point{v0, v4, v24, v20, v0}),
		newShape([]image.Point{v0, v4, v24, v0}),
		newShape([]image.Point{v2, v24, v20, v2}),
		newShape([]image.Point{v4, v2, v24, v22, v4}),
		newShape([]image.Point{v2, v14, v22, v10, v2}),
		newShape([]image.Point{v4, v22, v20, v10, v4}),
		newShape([]image.Point{v2, v24, v22, v13, v11, v22, v20, v2}),
		newShape([]image.Point{v4, v10, v22, v4}),
		newShape([]image.Point{v6, v8, v18, v16, v6}),
		newShape([]image.Point{v0, v24, v14, v12, v2, v0}),
		newShape([]image.Point{v2, v4, v14, v12, v2}),
		newShape([]image.Point{v10, v14, v22, v10}),
		newShape([]image.Point{v20, v12, v24, v20}),
		newShape([]image.Point{v14, v2, v12, v14}),
		newShape([]image.Point{v4, v2, v14, v4})}

	bg = color.NRGBA{0xf0, 0xf0, 0xf0, 0xff}
)

func DrawShapeOnImage(stroke, fill color.Color, i, size int, img draw.Image, p image.Point, angle float64) {
	i %= len(Shapes)
	shape := Shapes[i]
	gc := draw2d.NewGraphicContext(img)
	if angle != 0.0 {
		gc.Rotate(angle)
		gc.Translate(2, 2)
	}
	gc.SetLineWidth(0.0)
	gc.SetFillColor(fill)
	gc.SetStrokeColor(stroke)

	gc.BeginPath()
	v := shape.vertices
	dx := float64(v[0].X*size/4 + p.X)
	dy := float64(v[0].Y*size/4 + p.Y)
	gc.MoveTo(dx, dy)
	for j := 1; j < len(v); j++ {
		dx = float64(v[j].X*size/4 + p.X)
		dy = float64(v[j].Y*size/4 + p.Y)
		gc.LineTo(dx, dy)
	}
	gc.Close()
	gc.FillStroke()
}

func DrawShape(stroke, fill color.Color, i, size int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	DrawShapeOnImage(stroke, fill, i, size, img, image.ZP, 0.0)
	return img
}

func DrawShapeRotated(stroke, fill color.Color, i, size int, angle float64) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	DrawShapeOnImage(stroke, fill, i, size, img, image.ZP, angle)
	return img
}

func DrawShapeHorMirror(stroke, fill color.Color, i, size int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	DrawShapeHorMirrorOnImage(stroke, fill, i, size, img, image.ZP)
	return img
}

func DrawShapeHorMirrorOnImage(stroke, fill color.Color, i, size int, img draw.Image, p image.Point) {
	i %= len(ShapeHorMirrors)
	shape := ShapeHorMirrors[i]
	gc := draw2d.NewGraphicContext(img)
	gc.SetLineWidth(0.0)
	gc.SetFillColor(fill)
	gc.SetStrokeColor(stroke)
	v := shape.vertices
	dx := float64(v[0].X*size/4 + p.X)
	dy := float64(v[0].Y*size/4 + p.Y)
	gc.MoveTo(dx, dy)
	for j := 1; j < len(v); j++ {
		dx = float64(v[j].X*size/4 + p.X)
		dy = float64(v[j].Y*size/4 + p.Y)
		gc.LineTo(dx, dy)
	}
	gc.Close()
	gc.FillStroke()
}
