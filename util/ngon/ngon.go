package ngon

import (
	"math"

	"github.com/zolrath/euler/util/sink"
)

func IsPentagon(n int) bool {
	pent := (math.Sqrt(24*float64(n)+1) + 1) / 6
	return math.Floor(pent) == pent
}

func NthPentagon(n int) int {
	return n * (3*n - 1) / 2
}

func NPentagons(limit int) []int {
	pents := make([]int, limit)
	for i := 0; i < limit; i++ {
		pents[i] = NthPentagon(i + 1)
	}
	return pents
}

func PentagonGenerator(pos int) func() int {
	pos--
	return func() int {
		pos++
		return NthPentagon(pos)
	}
}

func IsTriangle(n int) bool {
	if sink.IsPerfectSquare(8*n + 1) {
		return true
	}
	return false
}

func NthTriangle(n int) int {
	return (n * (n + 1)) / 2
}

func NTriangles(limit int) []int {
	tris := make([]int, limit)
	curTri := 0
	for i := 1; i <= limit; i++ {
		curTri += i
		tris[i-1] = curTri
	}
	return tris
}

func TriangleGenerator(start int) func() int {
	var curTri, pos int
	if start == 0 {
		curTri, pos = 0, 1
	} else {
		curTri, pos = NthTriangle(start-1), start
	}
	return func() int {
		curTri += pos
		pos++
		return curTri
	}
}

func IsHexagon(n int) bool {
	hex := (math.Sqrt(8*float64(n)+1) + 1) / 4
	return hex == math.Floor(hex)
}

func NthHexagon(n int) int {
	return n * (2*n - 1)
}

func NHexagons(limit int) []int {
	hexes := make([]int, limit)
	for i := 0; i < limit; i++ {
		hexes[i] = NthHexagon(i + 1)
	}
	return hexes
}

func HexagonGenerator(pos int) func() int {
	pos--
	return func() int {
		pos++
		return NthHexagon(pos)
	}
}
