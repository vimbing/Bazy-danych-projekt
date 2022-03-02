package gui

import "github.com/rodrigocfd/windigo/ui"

type MyWindow struct {
	wnd         ui.WindowMain
	lblName     ui.Static
	txtName     ui.Edit
	lblPassword ui.Static
	txtPassword ui.Edit
	btnShow     ui.Button
}
