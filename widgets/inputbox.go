package widgets

import (
	"github.com/echenim/csui/components"
	"github.com/echenim/csui/core"
	"github.com/echenim/csui/window"
	"github.com/gdamore/tcell/v2"
)

func Input(msg string) (string, error) {

	win, err := window.New()
	if err != nil {
		return "", err
	}

	minLength := 50
	maxLength := minLength
	if len(msg) > maxLength {
		maxLength = len(msg)
	}

	minSize := core.SizeStrategyMultiple(
		core.SizeStrategyPercentage(80, 0),
		core.SizeStrategyAtLeast(core.Size{W: minLength, H: 1}),
		core.SizeStrategyAtMost(core.Size{W: maxLength + 8, H: 100}),
	)

	inputbox := components.NewInput()
	inputbox.SetSizeStrategy(minSize)
	listFrame := components.NewFrame(inputbox)

	text := components.NewText(msg)
	text.PadText(1)
	text.SetSizeStrategy(minSize)

	strip := components.NewColumnLayout()
	strip.SetSizeStrategy(minSize)

	var selected bool

	inputbox.OnKeypress(func(key *tcell.EventKey) bool {
		switch key.Key() {
		case tcell.KeyEnter:
			selected = true
			win.Close()
			return true
		}
		return false
	})

	help := components.NewText("ENTER to confirm, ESC to cancel")
	help.SetSizeStrategy(core.SizeStrategyMaximumWidth())
	help.SetAlignment(core.AlignRight)
	help.SetStyle(core.StyleFaint)
	strip.Add(help)

	rows := components.NewRowLayout()
	rows.Add(text)
	rows.Add(listFrame)
	rows.Add(components.NewSpacer(core.Size{H: 1}))
	rows.Add(strip)
	rows.SetAlignment(core.AlignCenter)

	win.SetAlignment(core.AlignCenter)
	win.Add(rows)

	if err := win.Show(); err != nil {
		return "", err
	}

	if !selected {
		return "", ErrInputCancelled
	}

	input := inputbox.GetInput()

	return input, nil
}
