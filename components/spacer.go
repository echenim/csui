package components

import (
	"github.com/echenim/csui/core"
)

type spacer struct {
	size core.Size
}

func NewSpacer(size core.Size) *spacer {
	return &spacer{
		size: size,
	}
}

func (t *spacer) Render(_ core.Canvas) {

}

func (t *spacer) MinimumSize() core.Size {
	return t.size
}
