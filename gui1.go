package main

import (
	"database/sql"
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	_ "github.com/go-sql-driver/mysql"

	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Registeratin Form")
	entry := widget.NewEntry()
	textArea := widget.NewMultiLineEntry()
	entry2 := widget.NewEntry()
	entry3 := widget.NewEntry()
	textArea3 := widget.NewMultiLineEntry()
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Entry your Name", Widget: entry}},
		OnSubmit: func() {

			enttxt := entry.Text
			phontxt := textArea.Text

			db, err := sql.Open("mysql", "Margo:12345@tcp(192.168.65.84:3306)/hotel")
			insert, err := db.Query("INSERT INTO guest(ID , first_name,Phone) values(24," + enttxt + "," + phontxt + ")")

			if err != nil {
				panic(err.Error())
			}
			defer insert.Close()
			log.Println("Form submitted:", entry.Text)
			log.Println("multiline:", textArea.Text)
		},
	}
	form2 := &widget.Form{

		Items: []*widget.FormItem{
			{Text: "Your Name", Widget: entry2}},
		OnSubmit: func() {
			enttxt := entry2.Text
			db, err := sql.Open("mysql", "Margo:12345@tcp(192.168.124.84:3306)/hotel")
			delete, err := db.Query("DELETE FROM guest WHERE (Name = " + enttxt + ")")

			if err != nil {
				panic(err.Error())
			}
			defer delete.Close()
			log.Println("Form submitted:", entry2.Text)
		},
	}
	form3 := &widget.Form{

		Items: []*widget.FormItem{
			{Text: "Enter your name", Widget: entry3}},

		OnSubmit: func() { // optional, handle form submission
			enttxt := entry3.Text
			phontxt := textArea3.Text

			db, err := sql.Open("mysql", "Margo:12345@tcp(192.168.65.84:3306)/hotel")

			update, err := db.Query("UPDATE guest SET Phone =" + phontxt + " WHERE Name =" + enttxt + "")

			if err != nil {
				panic(err.Error())
			}
			defer update.Close()
			log.Println("Form submitted:", entry3.Text)
			log.Println("multiline:", textArea3.Text)
		},
	}
	form.Append("Phone", textArea)
	form3.Append("Enter your correct phone", textArea3)

	tabs := container.NewAppTabs()
	tabs.Append(container.NewTabItemWithIcon("Book", theme.ConfirmIcon(), form))
	tabs.Append(container.NewTabItemWithIcon("Modify", theme.HistoryIcon(), form3))
	tabs.Append(container.NewTabItemWithIcon("Cancel", theme.CancelIcon(), form2))

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
