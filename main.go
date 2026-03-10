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

type UmaPaths struct {
	Target     string
	GlobalPath string
	JpPath     string
	UmaPath    string
	TempPath   string
}

func NewUmaPaths(appdata string) (UmaPaths, error) {
	if appdata == "" {
		return UmaPaths{}, fmt.Errorf("appdata environment variable not set")
	}
	target := filepath.Join(filepath.Dir(appdata), "LocalLow", "Cygames")
	return UmaPaths{
		Target:     target,
		GlobalPath: filepath.Join(target, "Umamusumegb"),
		JpPath:     filepath.Join(target, "Umamusumejp"),
		UmaPath:    filepath.Join(target, "Umamusume"),
		TempPath:   filepath.Join(target, "Umamusume_temp"),
	}, nil
}

func (u UmaPaths) checkCurrent() (string, error) {
	if _, err := os.Stat(u.JpPath); err == nil {
		return "current uma version: JP", nil
	} else if _, err := os.Stat(u.GlobalPath); err == nil {
		return "current uma version: Global", nil
	} else {
		return "", fmt.Errorf("no uma folder found")
	}
}

func main() {
	u, err := NewUmaPaths(os.Getenv("APPDATA"))

	fmt.Println(u)
	a := app.New()
	w := a.NewWindow("Uma Folder Changer")

	w.Resize(fyne.NewSize(200, 200))
	w.SetFixedSize(true)

	resultString := widget.NewLabel("")
	resultString.Wrapping = fyne.TextWrapWord

	if err != nil {
		resultString.SetText(fmt.Sprintf("Error: %v", err))
	}
	msg, err := u.checkCurrent()
	if err != nil {
		resultString.SetText(fmt.Sprintf("Error: %v", err))
	} else {
		resultString.SetText(msg)
	}

	w.SetContent(container.NewVBox(
		widget.NewButton("Click here to change folders", func() {
			msg, err := u.changeFolder()
			if err != nil {
				resultString.SetText(fmt.Sprintf("Error: %v", err))
			} else {
				resultString.SetText(msg)
			}
		}),
		resultString,
	))

	w.ShowAndRun()
}

func (u UmaPaths) changeFolder() (string, error) {

	if _, err := os.Stat(u.Target); err != nil {
		return "", err
	}

	if _, err := os.Stat(u.JpPath); err == nil {
		// this if goes from global to jp
		if err := os.Rename(u.UmaPath, u.TempPath); err != nil {
			return "", err
		}
		if err := os.Rename(u.JpPath, u.UmaPath); err != nil {
			return "", err
		}
		if err := os.Rename(u.TempPath, u.GlobalPath); err != nil {
			return "", err
		}
		return "Swapped from global to jp. you can now launch jp uma.", nil

	} else if _, err := os.Stat(u.GlobalPath); err == nil {
		if err := os.Rename(u.UmaPath, u.TempPath); err != nil {
			return "", err
		}
		if err := os.Rename(u.GlobalPath, u.UmaPath); err != nil {
			return "", err
		}
		if err := os.Rename(u.TempPath, u.JpPath); err != nil {
			return "", err
		}
		return "Swapped from jp to global. you can now launch global uma.", nil
	}
	return "no folder found to swap", nil
}
