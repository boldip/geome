package geome

import "math"
import "fmt"

type Point struct {
  X, Y float64
}

func Distance(p1, p2 Point) float64 {
  return math.Sqrt((p1.X - p2.X) * (p1.X - p2.X) + (p1.Y - p2.Y) * (p1.Y - p2.Y))
}

func String(p Point) string {
  return fmt.Sprintf("(%.3f,%.3f)", p.X, p.Y)
}

