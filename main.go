package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func rgb(r, g, b int) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

func lolcatify(i int) (int, int, int) {
	var f = 0.1

	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func hslToRGB(h, s, l float64) (r, g, b int) {
	var fR, fG, fB float64

	if s == 0 {
		fR, fG, fB = l, l, l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - s*l
		}
		p := 2*l - q
		fR = hueToRGB(p, q, h+1.0/3)
		fG = hueToRGB(p, q, h)
		fB = hueToRGB(p, q, h-1.0/3)
	}

	r = int((fR * 255) + 0.5)
	g = int((fG * 255) + 0.5)
	b = int((fB * 255) + 0.5)

	return
}

func hueToRGB(p, q, t float64) float64 {
	if t < 0 {
		t++
	}
	if t > 1 {
		t--
	}
	if t < 1.0/6 {
		return p + (q-p)*6*t
	}
	if t < 0.5 {
		return q
	}
	if t < 2.0/3 {
		return p + (q-p)*(2.0/3-t)*6
	}
	return p
}

func rainbow(input string, repeated bool, lolcat bool) {
	steps := strings.Split(input, "")

	for i, step := range steps {
		var h float64

		if repeated {
			h = float64(i) / 10
			h = h - math.Floor(h)
		} else {
			h = float64(i) / float64(len(steps))
		}

		var r, g, b int

		if lolcat {
			r, g, b = lolcatify(i)
		} else {
			r, g, b = hslToRGB(h, 1.0, 0.8)
		}

		color := rgb(r, g, b)

		fmt.Printf("%s%s", color, step)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var isRepeated bool
var isLolcat bool

func input() string {
	var data []byte
	var err error

	flag.BoolVar(&isRepeated, "repeated", false, "Make colors look very repeated, kinda like the lights in temple fairs")

	flag.BoolVar(&isLolcat, "lolcat", false, "Make colors look rainbow-ish like lolcats")

	flag.Parse()

	switch flag.NArg() {
	case 0:
		data, err = ioutil.ReadAll(os.Stdin)
		check(err)
		return string(data)
	case 1:
		data, err = ioutil.ReadFile(flag.Arg(0))
		check(err)
		return string(data)
	default:
		fmt.Printf("input must be from stdin or file\n")
		os.Exit(1)
		return ""
	}
}

func main() {
	words := input()

	rainbow(words, isRepeated, isLolcat)
}
