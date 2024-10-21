package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

// Метод для вычисления Манхэттенского расстояния между двумя точками
func (p1 Point) manhattanDistance(p2 Point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func getPoints(x, y, d int) map[Point]struct{} {
	posiblePoints := make(map[Point]struct{})
	for i := -d; i <= d; i++ {
		for j := -d; j <= d; j++ {
			if abs(i)+abs(j) <= d {
				curPoint := Point{x: x + i, y: y + j}
				posiblePoints[curPoint] = struct{}{}
			}
		}
	}
	return posiblePoints
}

func collision(fS, sS map[Point]struct{}, t int) map[Point]struct{} {
	res := make(map[Point]struct{})
	for p1 := range fS {
		for p2 := range sS {
			if p1.manhattanDistance(p2) <= t {
				res[p2] = struct{}{}
			}
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var t, d, n int
	fmt.Scan(&t, &d, &n)

	var x, y int
	fmt.Scan(&x, &y)
	oldPoints := getPoints(x, y, d)

	for i := 1; i < n; i++ {
		fmt.Scan(&x, &y)
		newPoints := getPoints(x, y, d)
		oldPoints = collision(oldPoints, newPoints, t)
	}

	fmt.Println(len(oldPoints))
	for p := range oldPoints {
		fmt.Printf("%d %d\n", p.x, p.y)
	}
}
