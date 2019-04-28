package canvas

import (
	"github.com/mfield4/178_projects/pkg/ui"
)

type Canvas interface {
	ui.Ui
	Update()
	Ref() Canvas
}
