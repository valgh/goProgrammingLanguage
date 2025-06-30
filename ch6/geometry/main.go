package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

// composing a type
type ColoredPoint struct {
	Point
	Color color.RGBA
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) ScaleByNoPointer(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(p.Distance(q)) // 5
	p.ScaleByNoPointer(2)
	fmt.Println(p.X, p.Y) // still 1,2
	p.ScaleBy(2)          // compiler performs an IMPLICIT &p on the variable if the receiver is a pointer!!!
	fmt.Println(p.X, p.Y) // now 2, 4

	// composition: struct embedding
	var cp ColoredPoint
	cp.X = 1
	cp.Y = 2
	cp.Color = color.RGBA{255, 0, 0, 255}
	fmt.Println(cp.X, cp.Y, cp.Color) // 1 2 {255, 0, 0, 255}
	// we can also call methods of the embedded structure!!
	// In a way, this is similar to inheritance in e.g. Java. Composition
	// allows to "inherit" bottom-up the methods of an embedded struct
	// into the structure that embeds it.
	cp.ScaleBy(2)
	fmt.Println(cp.X, cp.Y, cp.Color) // 2 4 {255, 0, 0, 255}
	// However, this is not real inheritance, and it's made clear by the Distance method:
	fmt.Println(q.Distance(cp.Point))
	// In such case, we must EXPLICITLY pass cp.Point in the funciont, as the parameter is declared
	// to be a Point, and NOT ColoredPoint, so it's clear that ColoredPoint is not a "class"
	// that extends the "base class" Point. So we say that in composition, the methods are PROMOTED
	// when a struct is embedded into another, and it's a slightly different
	// concept than inheritance! ColoredPoint HAS a Point, but IS NOT a Point.
}
