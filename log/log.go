package log

import (
	"github.com/mgutz/ansi"
	"log"
	"os"
)

var (
	d = ansi.ColorFunc("white+b:magenta")
	i = ansi.ColorFunc("white+b:blue")
	w = ansi.ColorFunc("white+b:red")
	f = ansi.ColorFunc("white+b:red+h")
	v = ansi.ColorFunc("white+b:black")

	infol = log.New(os.Stdout, i("[INFO] "))
	debugl = log.New(os.Stdout, d("[DEBUG] "))
	warnl = log.New(os.Stdout, w("[WARN] "))
	verbl = log.New(os.Stdout, v("[VERB] "))
	fatall = log.New(os.Stdout, f("[FATAL] "))
)

func Debug(i interface{}) {
	debugl.Print(i)
}

func Debugf(s string, i interface{}) {
	debugl.Printf(s, i)
}

func Info(i interface{}) {
	infol.Print(i)
}

func Infof(s string, i interface{}) {
	infol.Printf(s, i)
}

func Warn(i interface{}) {
	warnl.Print(i)
}

func Warnf(s string, i interface{}) {
	warnl.Printf(s, i)
}

func Verb(i interface{}) {
	verbl.Print(i)
}

func Verbf(s string, i interface{}) {
	verbl.Printf(s, i)
}

func Fatal(i interface{}) {
	fatall.Print(i)
}

func Fatalf(s string, i interface{}) {
	fatall.Printf(s, i)
}