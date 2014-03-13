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
func BenchmarkEuler016(b *testing.B) { benchmarkHarness(problems.Euler016, b) }
func BenchmarkEuler017(b *testing.B) { benchmarkHarness(problems.Euler017, b) }
func BenchmarkEuler018(b *testing.B) { benchmarkHarness(problems.Euler018, b) }
func BenchmarkEuler019(b *testing.B) { benchmarkHarness(problems.Euler019, b) }
func BenchmarkEuler020(b *testing.B) { benchmarkHarness(problems.Euler020, b) }
func BenchmarkEuler021(b *testing.B) { benchmarkHarness(problems.Euler021, b) }
func BenchmarkEuler022(b *testing.B) { benchmarkHarness(problems.Euler022, b) }
func BenchmarkEuler023(b *testing.B) { benchmarkHarness(problems.Euler023, b) }
func BenchmarkEuler024(b *testing.B) { benchmarkHarness(problems.Euler024, b) }
func BenchmarkEuler025(b *testing.B) { benchmarkHarness(problems.Euler025, b) }
func BenchmarkEuler026(b *testing.B) { benchmarkHarness(problems.Euler026, b) }
func BenchmarkEuler027(b *testing.B) { benchmarkHarness(problems.Euler027, b) }
func BenchmarkEuler028(b *testing.B) { benchmarkHarness(problems.Euler028, b) }
func BenchmarkEuler029(b *testing.B) { benchmarkHarness(problems.Euler029, b) }
func BenchmarkEuler030(b *testing.B) { benchmarkHarness(problems.Euler030, b) }
func BenchmarkEuler031(b *testing.B) { benchmarkHarness(problems.Euler031, b) }
func BenchmarkEuler032(b *testing.B) { benchmarkHarness(problems.Euler032, b) }
