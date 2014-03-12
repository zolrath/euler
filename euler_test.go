package main

import (
	"testing"

	"github.com/zolrath/euler/problems"
)

func benchmarkHarness(problem func() int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		problem()
	}
}

func BenchmarkEuler001(b *testing.B) { benchmarkHarness(problems.Euler001, b) }
func BenchmarkEuler002(b *testing.B) { benchmarkHarness(problems.Euler002, b) }
func BenchmarkEuler003(b *testing.B) { benchmarkHarness(problems.Euler003, b) }
func BenchmarkEuler004(b *testing.B) { benchmarkHarness(problems.Euler004, b) }
func BenchmarkEuler005(b *testing.B) { benchmarkHarness(problems.Euler005, b) }
func BenchmarkEuler006(b *testing.B) { benchmarkHarness(problems.Euler006, b) }
func BenchmarkEuler007(b *testing.B) { benchmarkHarness(problems.Euler007, b) }
func BenchmarkEuler008(b *testing.B) { benchmarkHarness(problems.Euler008, b) }
func BenchmarkEuler009(b *testing.B) { benchmarkHarness(problems.Euler009, b) }
func BenchmarkEuler010(b *testing.B) { benchmarkHarness(problems.Euler010, b) }
func BenchmarkEuler011(b *testing.B) { benchmarkHarness(problems.Euler011, b) }
func BenchmarkEuler012(b *testing.B) { benchmarkHarness(problems.Euler012, b) }
func BenchmarkEuler013(b *testing.B) { benchmarkHarness(problems.Euler013, b) }
func BenchmarkEuler014(b *testing.B) { benchmarkHarness(problems.Euler014, b) }
func BenchmarkEuler015(b *testing.B) { benchmarkHarness(problems.Euler015, b) }
