package main

import (
	"code.google.com/p/x-go-binding/xgb"
	"os"
	"path/filepath"
)

type Config struct {
	Instance string
	Class    string

	BackgroundColor    uint32
	NormalBorderColor  uint32
	FocusedBorderColor uint32
	BorderWidth        int16
	StatusLogger       StatusLogger

	DefaultCursor    xgb.Id
	MoveCursor       xgb.Id
	MultiClickTime   xgb.Timestamp
	MovedClickRadius int

	ModMask uint16
	Keys    map[byte]Cmd

	Ignore List
	Float  List
}

func configure() {
	// Configuration variables
	cfg = &Config{
		Instance: filepath.Base(os.Args[0]),
		Class:    "Mdtwm",

		NormalBorderColor:  rgbColor(0x8888, 0x8888, 0x8888),
		FocusedBorderColor: rgbColor(0x4444, 0x0000, 0xffff),
		BorderWidth:        1,

		StatusLogger: &Dzen2Logger{
			Writer:     os.Stdout,
			FgColor:	"#ddddcc",
			BgColor:	"#555588",
			TimeFormat: "Mon, Jan _2 15:04:05",
			TimePos:    1286,
		},

		DefaultCursor:    stdCursor(68),
		MoveCursor:       stdCursor(52),
		MultiClickTime:   300, // maximum interfal for multiclick [ms]
		MovedClickRadius: 5,   // minimal radius for moved click [pixel]

		ModMask: xgb.ModMask4,
		Keys: map[byte]Cmd{
			Key1:     {chDesk, 1},
			Key2:     {chDesk, 2},
			Key3:     {chDesk, 3},
			KeyEnter: {spawn, "gnome-terminal"},
			KeyQ:     {exit, 0},
		},

		Ignore: List{},
		Float:  List{"Mplayer", "Gimp"},
	}
	cfg.MovedClickRadius *= cfg.MovedClickRadius // We need square of radius
	if cfg.StatusLogger != nil {
		cfg.StatusLogger.Start()
	}

	// Layout
	root = NewRootPanel()
	// Setup list of desk
	desk1 := NewPanel(Horizontal, 1.97)
	desk2 := NewPanel(Horizontal, 1)
	desk3 := NewPanel(Horizontal, 1)
	root.Insert(desk1)
	root.Insert(desk2)
	root.Insert(desk3)
	// Setup two main panels on first desk
	desk1.Insert(NewPanel(Vertical, 1))
	desk1.Insert(NewPanel(Vertical, 0.3))
	// Setup one main panel on second and thrid desk
	desk2.Insert(NewPanel(Horizontal, 1))
	desk3.Insert(NewPanel(Vertical, 1))
	// Set current desk and current box
	currentDesk = desk1
	currentDesk.Raise()
	// In this box all existing windows will be placed
	currentBox = currentDesk.Children().Front()
}
