package gel

import (
	"image"
	"unicode/utf8"
	
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	
	"golang.org/x/image/math/fixed"
)

// Text is a widget for laying out and drawing text.
type Text struct {
	*Window
	// alignment specify the text alignment.
	alignment text.Alignment
	// maxLines limits the number of lines. Zero means no limit.
	maxLines int
}

func (w *Window) Text() *Text {
	return &Text{Window: w}
}

// Alignment sets the alignment for the text
func (t *Text) Alignment(alignment text.Alignment) *Text {
	t.alignment = alignment
	return t
}

// MaxLines sets the alignment for the text
func (t *Text) MaxLines(maxLines int) *Text {
	t.maxLines = maxLines
	return t
}

type lineIterator struct {
	Lines     []text.Line
	Clip      image.Rectangle
	Alignment text.Alignment
	Width     int
	Offset    image.Point
	
	y, prevDesc fixed.Int26_6
	txtOff      int
}

func (l *lineIterator) Next() (text.Layout, image.Point, bool) {
	for len(l.Lines) > 0 {
		line := l.Lines[0]
		l.Lines = l.Lines[1:]
		x := align(l.Alignment, line.Width, l.Width) + fixed.I(l.Offset.X)
		l.y += l.prevDesc + line.Ascent
		l.prevDesc = line.Descent
		// Align baseline and line start to the pixel grid.
		off := fixed.Point26_6{X: fixed.I(x.Floor()), Y: fixed.I(l.y.Ceil())}
		l.y = off.Y
		off.Y += fixed.I(l.Offset.Y)
		if (off.Y + line.Bounds.Min.Y).Floor() > l.Clip.Max.Y {
			break
		}
		lo := line.Layout
		start := l.txtOff
		l.txtOff += len(line.Layout.Text)
		if (off.Y + line.Bounds.Max.Y).Ceil() < l.Clip.Min.Y {
			continue
		}
		for len(lo.Advances) > 0 {
			_, n := utf8.DecodeRuneInString(lo.Text)
			adv := lo.Advances[0]
			if (off.X + adv + line.Bounds.Max.X - line.Width).Ceil() >= l.Clip.Min.X {
				break
			}
			off.X += adv
			lo.Text = lo.Text[n:]
			lo.Advances = lo.Advances[1:]
			start += n
		}
		end := start
		endx := off.X
		rn := 0
		for n, r := range lo.Text {
			if (endx + line.Bounds.Min.X).Floor() > l.Clip.Max.X {
				lo.Advances = lo.Advances[:rn]
				lo.Text = lo.Text[:n]
				break
			}
			end += utf8.RuneLen(r)
			endx += lo.Advances[rn]
			rn++
		}
		offf := image.Point{X: off.X.Floor(), Y: off.Y.Floor()}
		return lo, offf, true
	}
	return text.Layout{}, image.Point{}, false
}

func (t *Text) Fn(gtx layout.Context, s text.Shaper, font text.Font, size unit.Value, txt string) layout.Dimensions {
	cs := gtx.Constraints
	textSize := fixed.I(gtx.Px(size))
	lines := s.LayoutString(font, textSize, cs.Max.X, txt)
	if max := t.maxLines; max > 0 && len(lines) > max {
		lines = lines[:max]
	}
	dims := linesDimens(lines)
	dims.Size = cs.Constrain(dims.Size)
	cl := textPadding(lines)
	cl.Max = cl.Max.Add(dims.Size)
	it := segmentIterator{
		Lines:     lines,
		Clip:      cl,
		Alignment: t.alignment,
		Width:     dims.Size.X,
	}
	for {
		l, off, _, _, _, ok := it.Next()
		if !ok {
			break
		}
		stack := op.Save(gtx.Ops)
		op.Offset(layout.FPt(off)).Add(gtx.Ops)
		s.Shape(font, textSize, l).Add(gtx.Ops)
		clip.Rect(cl.Sub(off)).Add(gtx.Ops)
		paint.PaintOp{}.Add(gtx.Ops)
		stack.Load()
	}
	return dims
}
