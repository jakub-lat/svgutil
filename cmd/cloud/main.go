package main

import (
	"fmt"
	"github.com/ajstarks/svgo"
	"math/rand"
	"os"
	"svgutil/pkg/utils"
	"time"
)

const filename = "cloud.svg"
const w, h = 900, 500

func main() {
	rand.Seed(time.Now().Unix())

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	canvas := svg.New(f)
	canvas.Start(w, h)

	//canvas.Rect(0, 0, w, h, "fill:white")

	DrawEllipses(canvas, 500)
	DrawLines(canvas, 600, 100)

	canvas.End()
}

func DrawEllipses(canvas *svg.SVG, count int) {
	for i := 0; i < count; i++ {
		offsetX, offsetY, ew, eh :=
			utils.TruncatedNormal(0, 240, -350, 350),
			utils.TruncatedNormal(0, 70, -100, 100),
			utils.TruncatedNormal(0, 40, 80, 150),
			utils.TruncatedNormal(0, 40, 80, 150)

		color, alpha, stroke := 0, utils.RandRange(0.4, 1.0), utils.RandRange(2, 4)
		canvas.Ellipse(w/2+offsetX, h/2+offsetY, ew, eh, fmt.Sprintf("fill:none; stroke:#%02x%02x%02x; stroke-width:%d; stroke-opacity:%.2f;", color, color, color, stroke, alpha))
	}
}

func DrawLines(canvas *svg.SVG, count int, boundsOffset int) {
	for i := 0; i < count; i++ {
		c := utils.RandRange(0, 170)
		canvas.Line(
			utils.TruncatedNormal(float64(w/2), float64(w/2), float64(boundsOffset), float64(w-boundsOffset)),
			utils.TruncatedNormal(float64(h/2), float64(h/2), float64(boundsOffset), float64(h-boundsOffset)),
			utils.TruncatedNormal(float64(w/2), float64(w/2), float64(boundsOffset), float64(w-boundsOffset)),
			utils.TruncatedNormal(float64(h/2), float64(h/2), float64(boundsOffset), float64(h-boundsOffset)),
			fmt.Sprintf("stroke:#%02x%02x%02x; stroke-width:%d; stroke-opacity: %.2f", c, c, c, utils.RandRange(1, 2), utils.RandRange(0.2, 0.5)),
		)
	}
}
