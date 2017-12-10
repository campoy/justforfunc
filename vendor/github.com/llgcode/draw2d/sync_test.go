// go test -race -test.v sync_test.go

package draw2d_test


import (
	"fmt"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"image"
	"testing"
)

func TestSync(t *testing.T) {
	ch := make(chan int)
	limit := 2
	for i := 0; i < limit; i++ {
		go Draw(i, ch)
	}

	for i := 0; i < limit; i++ {
		counter := <-ch
		t.Logf("Goroutine %d returned\n", counter)
	}
}

func Draw(i int, ch chan<- int) {
	draw2d.SetFontFolder("./resource/font")
	// Draw a rounded rectangle using default colors
	dest := image.NewRGBA(image.Rect(0, 0, 297, 210.0))
	gc := draw2dimg.NewGraphicContext(dest)

	draw2dkit.RoundedRectangle(gc, 5, 5, 135, 95, 10, 10)
	gc.FillStroke()

	// Set the fill text color to black
	gc.SetFillColor(image.Black)
	gc.SetFontSize(14)

	// Display Hello World dimensions
	x1, y1, x2, y2 := gc.GetStringBounds("Hello world")
	gc.FillStringAt(fmt.Sprintf("%.2f %.2f %.2f %.2f", x1, y1, x2, y2), 0, 0)

	ch <- i
}
