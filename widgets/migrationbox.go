package widgets

import (
	"github.com/echenim/csui/components"
	"github.com/echenim/csui/core"
	"github.com/echenim/csui/window"
	"github.com/gdamore/tcell/v2"
)

func MigrationComponent() (string, string, error) {

	win, err := window.New()
	if err != nil {
		return "", "", err
	}

	minLength := 70
	maxLength := minLength

	minSize := core.SizeStrategyMultiple(
		core.SizeStrategyPercentage(80, 0),
		core.SizeStrategyAtLeast(core.Size{W: minLength, H: 1}),
		core.SizeStrategyAtMost(core.Size{W: maxLength + 8, H: 100}),
	)

	oldBox := components.NewInput()
	oldBox.SetSizeStrategy(minSize)
	oldFrame := components.NewFrame(oldBox)

	newBox := components.NewInput()
	newBox.SetSizeStrategy(minSize)
	newFrame := components.NewFrame(newBox)

	oldText := components.NewText("Enter organization Old ID ......")
	oldText.PadText(1)
	oldText.SetSizeStrategy(minSize)

	newText := components.NewText("Enter organization New ID ......")
	newText.PadText(1)
	newText.SetSizeStrategy(minSize)

	strip := components.NewColumnLayout()
	strip.SetSizeStrategy(minSize)

	var selected bool

	newBox.OnKeypress(func(key *tcell.EventKey) bool {
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
	rows.Add(oldText)
	rows.Add(oldFrame)
	rows.Add(newText)
	rows.Add(newFrame)
	rows.Add(components.NewSpacer(core.Size{H: 1}))
	rows.Add(strip)
	rows.SetAlignment(core.AlignCenter)
	win.SetAlignment(core.AlignCenter)
	win.Add(rows)

	if err := win.Show(); err != nil {
		return "", "", err
	}

	if !selected {
		return "", "", ErrInputCancelled
	}

	oldID := oldBox.GetInput()
	newID := newBox.GetInput()

	return oldID, newID, nil
}
