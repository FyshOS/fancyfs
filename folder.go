package fancyfs

import (
	"errors"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

var ErrNoMetadata = errors.New("no metadata for requested folder")

type FancyFolder struct {
	BackgroundURI      fyne.URI
	BackgroundResource fyne.Resource
	BackgroundFill     canvas.ImageFill
}

func DetailsForFolder(dir fyne.URI) (*FancyFolder, error) {
	err := ErrNoMetadata

	bg, err1 := checkBGImage(dir, ".background.png")
	if err1 == nil {
		return bg, nil
	} else if err1 != ErrNoMetadata {
		err = err1
	}
	bg, err2 := checkBGImage(dir, ".background.jpg")
	if err2 == nil {
		return bg, nil
	} else if err2 != ErrNoMetadata {
		err = err2
	}
	bg, err3 := checkBGImage(dir, ".background.jpeg")
	if err3 == nil {
		return bg, nil
	} else if err3 != ErrNoMetadata {
		err = err3
	}

	return nil, err
}

func checkBGImage(dir fyne.URI, name string) (*FancyFolder, error) {
	bgFile, _ := storage.Child(dir, name)

	if yes, err := storage.Exists(bgFile); !yes || err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNoMetadata
		}

		return nil, err
	}

	return &FancyFolder{
		BackgroundURI:  bgFile,
		BackgroundFill: canvas.ImageFillCover,
	}, nil
}
