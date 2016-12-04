package main

import (
    "github.com/ajstarks/svgo"
    "os"
)

func main() {
    width := 800
    height := 500
    canvas := svg.New(os.Stdout)
    canvas.Start(width, height)
    canvas.Rect(0, 0, width, height, "fill:none;stroke:blue")
    canvas.Text(width/2, 10, "First Header", "text-anchor:middle;font-size:10px;fill:black")
    canvas.Text(width/2, 30, "Second Header", "text-anchor:middle;font-size:10px;fill:black") 
    canvas.Text(width/2, 50, "Third Header", "text-anchor:middle;font-size:10px;fill:black")     
    canvas.Text(width/20, (height - 10), "First Footer", "text-anchor:left;font-size:10px;fill:black")
    canvas.Text(width/20, (height - 30), "Second Footer", "text-anchor:left;font-size:10px;fill:black") 
    canvas.Text(width/20, (height - 50), "Third Footer", "text-anchor:left;font-size:10px;fill:black")
    canvas.Line(100, 100, 100, 400, "stroke:black")
    canvas.Line(100, 400, 700, 400, "stroke:black")    
	canvas.Text(50, 100, "Axis Text", "rotate:90;text-anchor:left;font-size:10px;fill:black")

      
    canvas.End()
}
