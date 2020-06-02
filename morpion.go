package main

import "github.com/tadvi/winc"

func main() {

	var tableau [10]string
	var joueur string = "X"
	var gagnant string = "N"
	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(400, 300)
	mainWindow.SetText("Morpion")
	btn := winc.NewPushButton(mainWindow)
	btn.SetText("Fermer")
	btn.SetPos(150, 200)
	btn.SetSize(100, 40)
	btn1 := winc.NewPushButton(mainWindow)
	btn1.SetText("1")
	btn1.SetPos(125, 50)
	btn1.SetSize(50, 30)
	btn1.OnClick().Bind(func(e *winc.Event) {
		tableau[0] = joueur
		btn1.SetText(joueur)
		if joueur == "X" {
			joueur = "O"
		} else {
			joueur = "X"
		}
		gagnant = verifGagnant(tableau)
	})
	btn2 := winc.NewPushButton(mainWindow)
	btn2.SetText("2")
	btn2.SetPos(175, 50)
	btn2.SetSize(50, 30)
	btn2.OnClick().Bind(func(e *winc.Event) {
		tableau[1] = joueur
		btn2.SetText(joueur)
		if joueur == "X" {
			joueur = "O"
		} else {
			joueur = "X"
		}
		gagnant = verifGagnant(tableau)
	})
	btn3 := winc.NewPushButton(mainWindow)
	btn3.SetText("3")
	btn3.SetPos(225, 50)
	btn3.SetSize(50, 30)
	btn3.OnClick().Bind(func(e *winc.Event) {
		tableau[2] = joueur
		btn3.SetText(joueur)
		if joueur == "X" {
			joueur = "O"
		} else {
			joueur = "X"
		}
		gagnant = verifGagnant(tableau)
	})
	btn4 := winc.NewPushButton(mainWindow)
	btn4.SetText("4")
	btn4.SetPos(125, 80)
	btn4.SetSize(50, 30)
	btn4.OnClick().Bind(func(e *winc.Event) {
		tableau[3] = joueur
		btn4.SetText(joueur)
		if joueur == "X" {
			joueur = "O"
		} else {
			joueur = "X"
		}
		gagnant = verifGagnant(tableau)
	})
	btn5 := winc.NewPushButton(mainWindow)
	btn5.SetText("5")
	btn5.SetPos(175, 80)
	btn5.SetSize(50, 30)
	btn5.OnClick().Bind(func(e *winc.Event) {
		tableau[4] = joueur
		btn5.SetText(joueur)
		if joueur == "X" {
			joueur = "O"
		} else {
			joueur = "X"
		}
		gagnant = verifGagnant(tableau)
	})
	btn6 := winc.NewPushButton(mainWindow)
	btn6.SetText("6")
	btn6.SetPos(225, 80)
	btn6.SetSize(50, 30)
	btn6.OnClick().Bind(func(e *winc.Event) {
		tableau[5] = joueur
		btn6.SetText(joueur)
		if joueur == "X" {
			joueur = "O"
		} else {
			joueur = "X"
		}
		gagnant = verifGagnant(tableau)
	})
	btn7 := winc.NewPushButton(mainWindow)
	btn7.SetText("7")
	btn7.SetPos(125, 110)
	btn7.SetSize(50, 30)
	btn7.OnClick().Bind(func(e *winc.Event) {
		tableau[6] = joueur
		btn7.SetText(joueur)
		if joueur == "X" {
			joueur = "O"
		} else {
			joueur = "X"
		}
		gagnant = verifGagnant(tableau)
	})
	btn8 := winc.NewPushButton(mainWindow)
	btn8.SetText("8")
	btn8.SetPos(175, 110)
	btn8.SetSize(50, 30)
	btn8.OnClick().Bind(func(e *winc.Event) {
		tableau[7] = joueur
		btn8.SetText(joueur)
		if joueur == "X" {
			joueur = "O"
		} else {
			joueur = "X"
		}
		gagnant = verifGagnant(tableau)
	})
	btn9 := winc.NewPushButton(mainWindow)
	btn9.SetText("9")
	btn9.SetPos(225, 110)
	btn9.SetSize(50, 30)
	btn9.OnClick().Bind(func(e *winc.Event) {
		tableau[8] = joueur
		btn9.SetText(joueur)
		if joueur == "X" {
			joueur = "O"
		} else {
			joueur = "X"
		}
		gagnant = verifGagnant(tableau)
	})
	btn.OnClick().Bind(wndOnClose)
	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)
	winc.RunMainLoop()

}

func verifGagnant(tableau [10]string) string {
	if tableau[0] == "X" && tableau[1] == "X" && tableau[2] == "X" ||
		tableau[3] == "X" && tableau[4] == "X" && tableau[5] == "X" ||
		tableau[6] == "X" && tableau[7] == "X" && tableau[8] == "X" ||

		tableau[0] == "X" && tableau[3] == "X" && tableau[6] == "X" ||
		tableau[1] == "X" && tableau[4] == "X" && tableau[7] == "X" ||
		tableau[2] == "X" && tableau[5] == "X" && tableau[8] == "X" ||

		tableau[0] == "X" && tableau[4] == "X" && tableau[8] == "X" ||
		tableau[5] == "X" && tableau[4] == "X" && tableau[6] == "X" {
		return "X"
	}

	if tableau[0] == "O" && tableau[1] == "O" && tableau[2] == "O" ||
		tableau[3] == "O" && tableau[4] == "O" && tableau[5] == "O" ||
		tableau[6] == "O" && tableau[7] == "O" && tableau[8] == "O" ||

		tableau[0] == "O" && tableau[3] == "O" && tableau[6] == "O" ||
		tableau[1] == "O" && tableau[4] == "O" && tableau[7] == "O" ||
		tableau[2] == "O" && tableau[5] == "O" && tableau[8] == "O" ||

		tableau[0] == "O" && tableau[4] == "O" && tableau[8] == "O" ||
		tableau[5] == "O" && tableau[4] == "O" && tableau[6] == "O" {
		return "O"
	}
	return "N"
}

func wndOnClose(arg *winc.Event) {
	winc.Exit()
}
