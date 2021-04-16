package main

import (
	l "gioui.org/layout"
	"github.com/p9c/qu"

	"github.com/p9c/gel"
)

type State struct {
	*gel.Window
	evKey       int
	showClicker *gel.Clickable
	showText    *string
}

func NewState(quit qu.C) (s *State) {
	s = &State{
		Window: gel.NewWindowP9(quit),
	}
	s.showText = &s.Window.ClipboardContent
	s.showClicker = s.WidgetPool.GetClickable()
	return
}

func main() {
	quit := qu.T()
	state := NewState(quit)
	*state.showText = "hello world!"
	var e error
	// rootWidget := state.rootWidget()
	if e = state.Window.
		Size(48, 32).
		Title("hello world").
		Open().
		Run(state.rootWidget, quit.Q, quit); E.Chk(e) {
	}
}

func (s *State) rootWidget(gtx l.Context) l.Dimensions {
	return s.Direction().Center().
		Embed(
			s.Inset(0.5,
				s.Border().Color("DocText").Embed(
					s.VFlex().
						Flexed(0.25,
							s.Inset(0.5,
								s.Border().Color("DocText").Embed(
									s.Flex().Flexed(1,
										s.Direction().Center().Embed(
											s.ButtonLayout(
												s.showClicker.
													SetClick(func() {
														I.Ln("user clicked button")
														s.ClipboardReadReqs <- func(cs string) {
															*s.showText = cs
															I.Ln("clipboard contents:", cs)
														}
														// // clipboard.ReadOp{Tag: &s.evKey}.Add(gtx.Ops)
														// // I.Ln(runtime.GOOS)
														// // switch runtime.GOOS {
														// // case "ios", "android":
														// s.Window.ReadClipboard()
														// // I.S(s.Window.Window)
														// I.S(*s.showText, s.Window.ClipboardContent)
														// *s.showText = s.Window.ClipboardContent
														// // default:
														// // 	txt := clipboard.Get()
														// // 	*s.showText = txt
														// // }
														// I.Ln(*s.showText)
													}),
											).CornerRadius(0.25).Corners(^0).
												Embed(
													s.Border().CornerRadius(0.25).Color("DocText").Embed(
														s.Inset(0.5,
															s.H6("display text in clipboard").
																// Alignment(text.Middle).
																Fn,
														).Fn,
													).Fn,
												).Fn,
										).Fn,
									).Fn,
								).Fn,
							).Fn,

						).
						Flexed(0.75,
							s.Inset(0.5,
								s.Border().Color("DocText").Embed(
									s.Flex().Flexed(1,
										s.Direction().Center().Embed(
											s.H2(*s.showText).
												// Alignment(text.Middle).
												Fn,
										).Fn,
									).
										Fn,
								).Fn,
							).Fn,
						).Fn,
				).Fn,
			).Fn,
		).
		Fn(gtx)
}
