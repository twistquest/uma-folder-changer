package main

import (
	"fmt"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Uma Folder Changer")

	w.Resize(fyne.NewSize(200, 200))
	w.SetFixedSize(true)

	resultString := widget.NewLabel("")
	resultString.Wrapping = fyne.TextWrapWord
	w.SetContent(container.NewVBox(
		widget.NewButton("Click here to change folders", func() {
			msg, err := changeFolder()
			if err != nil {
				resultString.SetText(fmt.Sprintf("Error %v", err))
			} else {
				resultString.SetText(msg)
			}
		}),
		resultString,
	))

	w.ShowAndRun()
}

func changeFolder() (string, error) {
	appdata := os.Getenv("APPDATA")

	if appdata == "" {
		return "", fmt.Errorf("No appdata found!")
	}

	targetDir := filepath.Join(filepath.Dir(appdata), "LocalLow", "Cygames")
	jpPath := filepath.Join(targetDir, "Umamusumejp")
	gbPath := filepath.Join(targetDir, "Umamusumegb")
	umaPath := filepath.Join(targetDir, "Umamusume")
	tempPath := filepath.Join(targetDir, "Umamusume_temp")

	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		return "", err
	}

	if _, err := os.Stat(jpPath); err == nil {
		// this if goes from global to jp
		if err := os.Rename(umaPath, tempPath); err != nil {
			return "", err
		}
		if err := os.Rename(jpPath, umaPath); err != nil {
			return "", err
		}
		if err := os.Rename(tempPath, gbPath); err != nil {
			return "", err
		}
		return "Swapped from global to jp. you can now launch jp uma.", nil

	} else if _, err := os.Stat(gbPath); err == nil {
		if err := os.Rename(umaPath, tempPath); err != nil {
			return "", err
		}
		if err := os.Rename(gbPath, umaPath); err != nil {
			return "", err
		}
		if err := os.Rename(tempPath, jpPath); err != nil {
			return "", err
		}
		return "Swapped from jp to global. you can now launch global uma.", nil
	}
	return "no folder found to swap", nil
}
