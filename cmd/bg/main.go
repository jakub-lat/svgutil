package main

import (
	"encoding/xml"
	"fmt"
	svg "github.com/ajstarks/svgo"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"svgutil/pkg/utils"
	"time"
)

const filename = "bg.svg"
const w, h = 1620, 2160
const stainsDir = "stains"

func main() {
	rand.Seed(time.Now().Unix())

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	canvas := svg.New(f)
	canvas.Start(w, h)

	canvas.Rect(0, 0, w, h, "fill:#bfbebd")
	DrawLines(canvas, 150, "#a1a19a")
	if err := DrawStains(canvas, 100); err != nil {
		panic(err)
	}

	canvas.End()
}

func DrawLines(canvas *svg.SVG, count int, color string) {
	outside := 50
	for i := 0; i < count; i++ {
		canvas.Line(
			utils.RandRange(-outside, w+outside),
			utils.RandRange(-outside, h+outside),
			utils.RandRange(-outside, w+outside),
			utils.RandRange(-outside, h+outside),
			fmt.Sprintf("stroke:%v; stroke-width:%d; stroke-opacity: %.2f", color, utils.RandRange(3, 5), utils.RandRange(0.1, 0.4)),
		)
	}
}

type SVG struct {
	Width  float64 `xml:"width,attr"`
	Height float64 `xml:"height,attr"`
	Doc    string  `xml:",innerxml"`
}

func LoadSvg(filename string) (SVG, error) {
	s := &SVG{}
	f, err := os.Open(filename)
	defer f.Close()

	err = xml.NewDecoder(f).Decode(&s)
	if err != nil {
		return SVG{}, err
	}

	return *s, nil
}

func LoadStains() (res []SVG, err error) {
	dir, err := os.ReadDir(stainsDir)
	if err != nil {
		return
	}

	for _, x := range dir {
		s, err := LoadSvg(filepath.Join(stainsDir, x.Name()))
		if err != nil {
			return nil, err
		}
		res = append(res, s)
	}

	return
}

func DrawStains(canvas *svg.SVG, count int) error {
	stains, err := LoadStains()
	if err != nil {
		return err
	}

	boundsOffset := 50

	for i, x := range stains {
		canvas.Group(
			fmt.Sprintf("id='stain%d'", i),
			fmt.Sprintf("transform='scale(0.12) translate(-800, -800)'"),
			"fill:#85857f")
		io.WriteString(canvas.Writer, x.Doc)
		canvas.Gend()
	}

	for i := 0; i < count; i++ {
		stainId := utils.RandRange(0, len(stains)-1)

		canvas.Use(
			utils.RandRange(boundsOffset, w-boundsOffset),
			utils.RandRange(boundsOffset, h-boundsOffset),
			fmt.Sprintf("#stain%d", stainId),
			fmt.Sprintf("opacity:%.2f", utils.RandRange(0.1, 0.5)),
		)
	}

	return nil
}
