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

func testHarness(problem func() int, expected int, t *testing.T) {
	answer := problem()
	if answer != expected {
		t.Errorf("Expected: %d Got: %d", expected, answer)
	}
}

func TestEuler001(t *testing.T) { testHarness(problems.Euler001, problems.ANSWER_001, t) }
func TestEuler002(t *testing.T) { testHarness(problems.Euler002, problems.ANSWER_002, t) }
func TestEuler003(t *testing.T) { testHarness(problems.Euler003, problems.ANSWER_003, t) }
func TestEuler004(t *testing.T) { testHarness(problems.Euler004, problems.ANSWER_004, t) }
func TestEuler005(t *testing.T) { testHarness(problems.Euler005, problems.ANSWER_005, t) }
func TestEuler006(t *testing.T) { testHarness(problems.Euler006, problems.ANSWER_006, t) }
func TestEuler007(t *testing.T) { testHarness(problems.Euler007, problems.ANSWER_007, t) }
func TestEuler008(t *testing.T) { testHarness(problems.Euler008, problems.ANSWER_008, t) }
func TestEuler009(t *testing.T) { testHarness(problems.Euler009, problems.ANSWER_009, t) }
func TestEuler010(t *testing.T) { testHarness(problems.Euler010, problems.ANSWER_010, t) }

func TestEuler011(t *testing.T) { testHarness(problems.Euler011, problems.ANSWER_011, t) }
func TestEuler012(t *testing.T) { testHarness(problems.Euler012, problems.ANSWER_012, t) }
func TestEuler013(t *testing.T) { testHarness(problems.Euler013, problems.ANSWER_013, t) }
func TestEuler014(t *testing.T) { testHarness(problems.Euler014, problems.ANSWER_014, t) }
func TestEuler015(t *testing.T) { testHarness(problems.Euler015, problems.ANSWER_015, t) }
func TestEuler016(t *testing.T) { testHarness(problems.Euler016, problems.ANSWER_016, t) }
func TestEuler017(t *testing.T) { testHarness(problems.Euler017, problems.ANSWER_017, t) }
func TestEuler018(t *testing.T) { testHarness(problems.Euler018, problems.ANSWER_018, t) }
func TestEuler019(t *testing.T) { testHarness(problems.Euler019, problems.ANSWER_019, t) }

func TestEuler020(t *testing.T) { testHarness(problems.Euler020, problems.ANSWER_020, t) }
func TestEuler021(t *testing.T) { testHarness(problems.Euler021, problems.ANSWER_021, t) }
func TestEuler022(t *testing.T) { testHarness(problems.Euler022, problems.ANSWER_022, t) }
func TestEuler023(t *testing.T) { testHarness(problems.Euler023, problems.ANSWER_023, t) }
func TestEuler024(t *testing.T) { testHarness(problems.Euler024, problems.ANSWER_024, t) }
func TestEuler025(t *testing.T) { testHarness(problems.Euler025, problems.ANSWER_025, t) }
func TestEuler026(t *testing.T) { testHarness(problems.Euler026, problems.ANSWER_026, t) }
func TestEuler027(t *testing.T) { testHarness(problems.Euler027, problems.ANSWER_027, t) }
func TestEuler028(t *testing.T) { testHarness(problems.Euler028, problems.ANSWER_028, t) }
func TestEuler029(t *testing.T) { testHarness(problems.Euler029, problems.ANSWER_029, t) }

func TestEuler030(t *testing.T) { testHarness(problems.Euler030, problems.ANSWER_030, t) }
func TestEuler031(t *testing.T) { testHarness(problems.Euler031, problems.ANSWER_031, t) }
func TestEuler032(t *testing.T) { testHarness(problems.Euler032, problems.ANSWER_032, t) }
func TestEuler033(t *testing.T) { testHarness(problems.Euler033, problems.ANSWER_033, t) }
func TestEuler034(t *testing.T) { testHarness(problems.Euler034, problems.ANSWER_034, t) }
func TestEuler035(t *testing.T) { testHarness(problems.Euler035, problems.ANSWER_035, t) }
func TestEuler036(t *testing.T) { testHarness(problems.Euler036, problems.ANSWER_036, t) }
func TestEuler037(t *testing.T) { testHarness(problems.Euler037, problems.ANSWER_037, t) }
func TestEuler038(t *testing.T) { testHarness(problems.Euler038, problems.ANSWER_038, t) }
func TestEuler039(t *testing.T) { testHarness(problems.Euler039, problems.ANSWER_039, t) }

func TestEuler040(t *testing.T) { testHarness(problems.Euler040, problems.ANSWER_040, t) }
func TestEuler041(t *testing.T) { testHarness(problems.Euler041, problems.ANSWER_041, t) }
func TestEuler042(t *testing.T) { testHarness(problems.Euler042, problems.ANSWER_042, t) }
func TestEuler043(t *testing.T) { testHarness(problems.Euler043, problems.ANSWER_043, t) }
func TestEuler044(t *testing.T) { testHarness(problems.Euler044, problems.ANSWER_044, t) }
func TestEuler045(t *testing.T) { testHarness(problems.Euler045, problems.ANSWER_045, t) }
func TestEuler046(t *testing.T) { testHarness(problems.Euler046, problems.ANSWER_046, t) }
func TestEuler047(t *testing.T) { testHarness(problems.Euler047, problems.ANSWER_047, t) }
func TestEuler048(t *testing.T) { testHarness(problems.Euler048, problems.ANSWER_048, t) }
func TestEuler049(t *testing.T) { testHarness(problems.Euler049, problems.ANSWER_049, t) }

func TestEuler050(t *testing.T) { testHarness(problems.Euler050, problems.ANSWER_050, t) }
func TestEuler051(t *testing.T) { testHarness(problems.Euler051, problems.ANSWER_051, t) }

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
func BenchmarkEuler033(b *testing.B) { benchmarkHarness(problems.Euler033, b) }
func BenchmarkEuler034(b *testing.B) { benchmarkHarness(problems.Euler034, b) }
func BenchmarkEuler035(b *testing.B) { benchmarkHarness(problems.Euler035, b) }
func BenchmarkEuler036(b *testing.B) { benchmarkHarness(problems.Euler036, b) }
func BenchmarkEuler037(b *testing.B) { benchmarkHarness(problems.Euler037, b) }
func BenchmarkEuler038(b *testing.B) { benchmarkHarness(problems.Euler038, b) }
func BenchmarkEuler039(b *testing.B) { benchmarkHarness(problems.Euler039, b) }

func BenchmarkEuler040(b *testing.B) { benchmarkHarness(problems.Euler040, b) }
func BenchmarkEuler041(b *testing.B) { benchmarkHarness(problems.Euler041, b) }
func BenchmarkEuler042(b *testing.B) { benchmarkHarness(problems.Euler042, b) }
func BenchmarkEuler043(b *testing.B) { benchmarkHarness(problems.Euler043, b) }
func BenchmarkEuler044(b *testing.B) { benchmarkHarness(problems.Euler044, b) }
func BenchmarkEuler045(b *testing.B) { benchmarkHarness(problems.Euler045, b) }
func BenchmarkEuler046(b *testing.B) { benchmarkHarness(problems.Euler046, b) }
func BenchmarkEuler047(b *testing.B) { benchmarkHarness(problems.Euler047, b) }
func BenchmarkEuler048(b *testing.B) { benchmarkHarness(problems.Euler048, b) }
func BenchmarkEuler049(b *testing.B) { benchmarkHarness(problems.Euler049, b) }

func BenchmarkEuler050(b *testing.B) { benchmarkHarness(problems.Euler050, b) }
func BenchmarkEuler051(b *testing.B) { benchmarkHarness(problems.Euler051, b) }
