// SPDX-License-Identifier: Unlicense OR MIT

package gel

import (
	"gioui.org/f32"
	"image"
	
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

const defaultScale = float32(160.0 / 72.0)

// Image is a widget that displays an image.
type Image struct {
	// src is the image to display.
	src paint.ImageOp
	// Fit specifies how to scale the image to the constraints.
	// By default it does not do any scaling.
	Fit Fit
	// Position specifies where to position the image within
	// the constraints.
	Position layout.Direction
	// Scale is the ratio of image pixels to
	// dps. If Scale is zero Image falls back to
	// a scale that match a standard 72 DPI.
	scale float32
}

func (th *Theme) Image() *Image {
	return &Image{}
}

func (i *Image) Src(img paint.ImageOp) *Image {
	i.src = img
	return i
}

func (i *Image) Scale(scale float32) *Image {
	i.scale = scale
	return i
}

func (i Image) Fn(gtx layout.Context) layout.Dimensions {
	defer op.Save(gtx.Ops).Load()
	scale := i.scale
	if scale == 0 {
		scale = defaultScale
	}
	size := i.src.Size()
	wf, hf := float32(size.X), float32(size.Y)
	w, h := gtx.Px(unit.Dp(wf*scale)), gtx.Px(unit.Dp(hf*scale))
	dims := i.Fit.scale(gtx, i.Position, layout.Dimensions{Size: image.Pt(w, h)})
	pixelScale := scale * gtx.Metric.PxPerDp
	op.Affine(f32.Affine2D{}.Scale(f32.Point{}, f32.Pt(pixelScale, pixelScale))).Add(gtx.Ops)
	i.src.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
	return dims
}
