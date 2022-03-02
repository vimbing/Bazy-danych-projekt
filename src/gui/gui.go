package gui

import (
	"fmt"
	"runtime"

	"main/src/db"

	"github.com/rodrigocfd/windigo/ui"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

func Render() {
	runtime.LockOSThread()

	myWindow := NewMyWindow()
	myWindow.wnd.RunAsMain()
}

func validateUser(users *[]db.User, login string, password string) bool {
	for _, user := range *users {
		fmt.Println(user.Login, user.Password)
		if user.Login == login && user.Password == password {
			fmt.Println(login, password)
			return true
		}
	}
	return false
}

func NewMyWindow() *MyWindow {
	wnd := ui.NewWindowMain(
		ui.WindowMainOpts().
			Title("Logowanie").
			ClientArea(win.SIZE{Cx: 270, Cy: 180}),
	)

	me := &MyWindow{
		wnd: wnd,
		lblName: ui.NewStatic(wnd,
			ui.StaticOpts().
				Text("Login").
				Position(win.POINT{X: 10, Y: 22}),
		),
		txtName: ui.NewEdit(wnd,
			ui.EditOpts().
				Position(win.POINT{X: 80, Y: 20}).
				Size(win.SIZE{Cx: 150}),
		),
		lblPassword: ui.NewStatic(wnd,
			ui.StaticOpts().
				Text("Password").
				Position(win.POINT{X: 10, Y: 64}),
		),
		txtPassword: ui.NewEdit(wnd,
			ui.EditOpts().
				Position(win.POINT{X: 80, Y: 64}).
				Size(win.SIZE{Cx: 150}),
		),
		btnShow: ui.NewButton(wnd,
			ui.ButtonOpts().
				Text("&Login").
				Position(win.POINT{X: 80, Y: 120}),
		),
	}

	usersInDb := db.GetData()

	me.btnShow.On().BnClicked(func() {
		password := me.txtPassword.Text()
		login := me.txtName.Text()

		ifValid := validateUser(&usersInDb, login, password)

		if ifValid {
			me.wnd.Hwnd().MessageBox(fmt.Sprintf("Hello %s", login), "Success!", co.MB_ICONINFORMATION)
		} else {
			me.wnd.Hwnd().MessageBox(fmt.Sprintf("No such credentials in db (%s)", login), "Failed!", co.MB_ICONINFORMATION)
		}

	})

	return me
}
