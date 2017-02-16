package main

import (
	"fmt"
	"math"
)

// Isometric projection is done by setting the angle of the front point
// lines to the horizontal to 30 degrees
const (
	width, height		=	600, 320 				// canvas size in pixels
	cells				=	5   					// number of grid cells
	xyrange				=	30.0					// axis ranges (-xyrange...+xyrange)
	xyscale				=	width / 2 / xyrange 	// pixels per x or y unit
	zscale				=	height * 0.4			// pixels per z unit
	angle				=	math.Pi / 6				// angle of x and y axes (=30deg)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)	// sin(30), sin(30)

func main() {

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			fmt.Printf("Cell: i=%v, j=%v\n", i, j)
			ax, ay := corner(i+1, j)
			fmt.Printf("\t Corner A: ax=%6.2f, ay=%6.2f\n", ax, ay)
			bx, by := corner(i, j)
			fmt.Printf("\t Corner B: bx=%6.2f, by=%6.2f\n", bx, by)
			cx, cy := corner(i, j+1)
			fmt.Printf("\t Corner C: cx=%6.2f, cy=%6.2f\n", cx, cy)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("\t Corner D: dx=%6.2f, dy=%6.2f\n\n", dx, dy)
		}
	}
}

func corner(i, j int) (float64, float64) {
	// Find point (x, y) at corner of cell (i, j)
	fmt.Printf("Original x=%v, y=%v" ,i,j)
	x := xyrange * (float64(i)/cells - 0.5)	
	y := xyrange * (float64(j)/cells - 0.5)	
	// Compute surface height z
	z := f(x, y)
	fmt.Printf("\t In 3D space x=%6.2f, y=%6.2f, z=%6.3f" ,x,y,z)
	
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy).
	xfact := cos30*xyscale
	sx := width/2 + (x-y)*xfact
	yfact := sin30*xyscale
	sy := height/2 + (x+y)*yfact - z*zscale
	fmt.Printf("\t Projection to X: x-y=%6.2f, xfact=%6.2f" ,(x-y), xfact)
	fmt.Printf("\t Projection to Y: x+y=%6.2f, yfact=%6.2f, zterm=%6.2f" ,(x+y), yfact, (z*zscale))
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
