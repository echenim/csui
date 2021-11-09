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
	if len(passwordMsg) > maxLength {
		maxLength = len(passwordMsg)
	}

	minSize := core.SizeStrategyMultiple(
		core.SizeStrategyPercentage(80, 0),
		core.SizeStrategyAtLeast(core.Size{W: minLength, H: 1}),
		core.SizeStrategyAtMost(core.Size{W: maxLength + 8, H: 100}),
	)

	env := components.NewListSelect(options)
	env.SetSizeStrategy(minSize)
	envtext := components.NewText(envMsg)
	usernamebox := components.NewInput()
	passwordbox := components.NewPasswordInput()

	envFrame := components.NewFrame(env)
	usernameFrame := components.NewFrame(usernamebox)
	passwordFrame := components.NewFrame(passwordbox)

	envtext.PadText(1)
	envtext.SetSizeStrategy(minSize)
	usernametext := components.NewText(usernameMsg)
	usernametext.PadText(1)
	usernametext.SetSizeStrategy(minSize)
	passwordtext := components.NewText(passwordMsg)
	passwordtext.PadText(1)
	passwordtext.SetSizeStrategy(minSize)

	strip := components.NewColumnLayout()
	strip.SetSizeStrategy(minSize)

	var selected bool

	passwordbox.OnKeypress(func(key *tcell.EventKey) bool {
		switch key.Key() {
		case tcell.KeyEnter:
			selected = true
			win.Close()
			return true
		}
		return false
	})

	help := components.NewText("ENTER to confirm, ESC to cancel")
	help.SetAlignment(core.AlignRight)
	help.SetSizeStrategy(core.SizeStrategyMaximumWidth())
	help.SetStyle(core.StyleFaint)
	strip.Add(help)

	rows := components.NewRowLayout()
	rows.Add(envtext)
	rows.Add(envFrame)
	rows.Add(usernametext)
	rows.Add(usernameFrame)
	rows.Add(passwordtext)
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

	username := usernamebox.GetInput()
	password := passwordbox.GetInput()
	_, enviroment := env.GetSelection()

	return enviroment, username, password, nil
}
