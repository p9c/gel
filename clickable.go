// SPDX-License-Identifier: Unlicense OR MIT

package gel

import (
	"image"
	"time"
	
	"gioui.org/f32"
	"gioui.org/gesture"
	"gioui.org/io/key"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
)

type clickEvents struct {
	Click, Cancel, Press func()
}

// click represents a click.
type click struct {
	Modifiers key.Modifiers
	NumClicks int
}

// Clickable represents a clickable area.
type Clickable struct {
	*Window
	click  gesture.Click
	clicks []click
	// prevClicks is the index into clicks that marks the clicks from the most recent Fn call. prevClicks is used to
	// keep clicks bounded.
	prevClicks int
	history    []Press
	Events     clickEvents
}

func (w *Window) Clickable() (c *Clickable) {
	c = &Clickable{
		Window:     w,
		click:      gesture.Click{},
		clicks:     nil,
		prevClicks: 0,
		history:    nil,
		Events: clickEvents{
			Click: func() {
				D.Ln("click event")
			},
			Cancel: func() {
				D.Ln("cancel event")
			},
			Press: func() {
				D.Ln("press event")
			},
		},
	}
	return
}

func (c *Clickable) SetClick(fn func()) *Clickable {
	c.Events.Click = fn
	return c
}

func (c *Clickable) SetCancel(fn func()) *Clickable {
	c.Events.Cancel = fn
	return c
}

func (c *Clickable) SetPress(fn func()) *Clickable {
	c.Events.Press = fn
	return c
}

// Click represents a click.
type Click struct {
	Modifiers key.Modifiers
	NumClicks int
}

// Press represents a past pointer press.
type Press struct {
	// Position of the press.
	Position f32.Point
	// Start is when the press began.
	Start time.Time
	// End is when the press was ended by a release or cancel.
	// A zero End means it hasn't ended yet.
	End time.Time
	// Cancelled is true for cancelled presses.
	Cancelled bool
}

// Click executes a simple programmatic click
func (c *Clickable) Click() {
	c.clicks = append(c.clicks, click{
		Modifiers: 0,
		NumClicks: 1,
	})
}

// Clicked reports whether there are pending clicks as would be
// reported by Clicks. If so, Clicked removes the earliest click.
func (c *Clickable) Clicked() bool {
	if len(c.clicks) == 0 {
		return false
	}
	n := copy(c.clicks, c.clicks[1:])
	c.clicks = c.clicks[:n]
	if c.prevClicks > 0 {
		c.prevClicks--
	}
	return true
}

// Hovered returns whether pointer is over the element.
func (c *Clickable) Hovered() bool {
	return c.click.Hovered()
}

// Pressed returns whether pointer is pressing the element.
func (c *Clickable) Pressed() bool {
	return c.click.Pressed()
}

// Clicks returns and clear the clicks since the last call to Clicks.
func (c *Clickable) Clicks() []click {
	clicks := c.clicks
	c.clicks = nil
	c.prevClicks = 0
	return clicks
}

// History is the past pointer presses useful for drawing markers.
// History is retained for a short duration (about a second).
func (c *Clickable) History() []Press {
	return c.history
}

// Layout and update the button state
func (c *Clickable) Layout(gtx layout.Context) layout.Dimensions {
	c.update(gtx)
	stack := op.Save(gtx.Ops)
	pointer.Rect(image.Rectangle{Max: gtx.Constraints.Min}).Add(gtx.Ops)
	c.click.Add(gtx.Ops)
	stack.Load()
	for len(c.history) > 0 {
		cl := c.history[0]
		if cl.End.IsZero() || gtx.Now.Sub(cl.End) < 1*time.Second {
			break
		}
		n := copy(c.history, c.history[1:])
		c.history = c.history[:n]
	}
	return layout.Dimensions{Size: gtx.Constraints.Min}
}

// update the button state by processing events.
func (c *Clickable) update(gtx layout.Context) {
	// Flush clicks from before the last update.
	n := copy(c.clicks, c.clicks[c.prevClicks:])
	c.clicks = c.clicks[:n]
	c.prevClicks = n
	
	for _, e := range c.click.Events(gtx) {
		switch e.Type {
		case gesture.TypeClick:
			c.clicks = append(c.clicks, click{
				Modifiers: e.Modifiers,
				NumClicks: e.NumClicks,
			})
			if l := len(c.history); l > 0 {
				c.history[l-1].End = gtx.Now
			}
		case gesture.TypeCancel:
			for i := range c.history {
				c.history[i].Cancelled = true
				if c.history[i].End.IsZero() {
					c.history[i].End = gtx.Now
				}
			}
		case gesture.TypePress:
			c.history = append(c.history, Press{
				Position: e.Position,
				Start:    gtx.Now,
			})
		}
	}
}
