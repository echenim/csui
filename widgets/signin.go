package widgets

import (
	"github.com/echenim/csui/components"
	"github.com/echenim/csui/core"
	"github.com/echenim/csui/window"
	"github.com/gdamore/tcell/v2"
)

func SignInInput(envMsg, usernameMsg, passwordMsg string, options []string) (string, string, string, error) {

	win, err := window.New()
	if err != nil {
		return "", "", "", err
	}

	minLength := 50
	maxLength := minLength
	if len(usernameMsg) > maxLength {
		maxLength = len(usernameMsg)
	}

	minSize := core.SizeStrategyMultiple(
		core.SizeStrategyPercentage(80, 0),
		core.SizeStrategyAtLeast(core.Size{W: minLength, H: 1}),
		core.SizeStrategyAtMost(core.Size{W: maxLength + 8, H: 100}),
	)

	//setup the enviroment
	environmentList := components.NewListSelect(options)
	environmentList.SetSizeStrategy(minSize)
	environmentFrame := components.NewFrame(environmentList)

	usernameBox := components.NewInput()
	usernameBox.SetSizeStrategy(minSize)
	usernameFrame := components.NewFrame(usernameBox)

	passwordBox := components.NewPasswordInput()
	passwordBox.SetSizeStrategy(minSize)
	passwordFrame := components.NewFrame(passwordBox)

	environmentText := components.NewText(envMsg)
	environmentText.PadText(1)
	environmentText.SetSizeStrategy(minSize)

	usernameText := components.NewText(usernameMsg)
	usernameText.PadText(1)
	usernameText.SetSizeStrategy(minSize)

	passwordText := components.NewText(passwordMsg)
	passwordText.PadText(1)
	passwordText.SetSizeStrategy(minSize)

	strip := components.NewColumnLayout()
	strip.SetSizeStrategy(minSize)

	var selected bool

	environmentList.OnKeypress(func(key *tcell.EventKey) bool {
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
	rows.Add(environmentText)
	rows.Add(environmentFrame)
	rows.Add(usernameText)
	rows.Add(usernameFrame)
	rows.Add(passwordText)
	rows.Add(passwordFrame)
	rows.Add(components.NewSpacer(core.Size{H: 1}))
	rows.Add(strip)
	rows.SetAlignment(core.AlignCenter)
	win.SetAlignment(core.AlignCenter)
	win.Add(rows)

	if err := win.Show(); err != nil {
		return "", "", "", err
	}

	if !selected {
		return "", "", "", ErrInputCancelled
	}

	username := usernameBox.GetInput()
	password := passwordBox.GetInput()
	_, enviroment := environmentList.GetSelection()

	return enviroment, username, password, nil
}
