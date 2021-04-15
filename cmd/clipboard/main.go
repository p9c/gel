package main

import (
	l "gioui.org/layout"
	"github.com/p9c/qu"

	"github.com/p9c/gel"
)

type State struct {
	*gel.Window
}

func NewState(quit qu.C) *State {
	return &State{
		Window: gel.NewWindowP9(quit),
	}
}

func main() {
	quit := qu.T()
	state := NewState(quit)
	var e error
	rootWidget := state.rootWidget()
	if e = state.Window.
		Size(48, 32).
		Title("hello world").
		Open().
		Run(rootWidget, quit.Q, quit); E.Chk(e) {
	}
}

func (s *State) rootWidget() l.Widget {
	showText := "hello world!"
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
												s.WidgetPool.GetClickable().
													SetClick(func() {
														I.Ln("user clicked button")

													}),
											).CornerRadius(0.25).Corners(^0).
												Embed(
													s.Border().CornerRadius(0.25).Color("DocText").Embed(
														s.Inset(0.5,
															s.H6("click me!").
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
											s.H2(showText).
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
		Fn
}
