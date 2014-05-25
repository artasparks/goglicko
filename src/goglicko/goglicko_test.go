package goglicko

import "testing"

// Much of this data comes from the paper:
// http://en.wikipedia.org/wiki/Glicko_rating_system
var pl = NewRating(1500, 200, DefaultVol)
var opps = []*Rating{
	NewRating(1400, 30, DefaultVol),
	NewRating(1550, 100, DefaultVol),
	NewRating(1700, 300, DefaultVol),
}

func TestEquivTransfOpps(t *testing.T) {
	for i := range opps {
		o := opps[i]
		o2 := opps[i].ToGlicko2().FromGlicko2()
		if !o.MostlyEquals(o2, 0.0001) {
			t.Errorf("o %v != o2 %v", o, o2)
		}
	}
}

func TestToGlicko2(t *testing.T) {
	p2 := pl.ToGlicko2()
	exp := NewRating(0, 1.1513, DefaultVol)
	if !p2.MostlyEquals(exp, 0.0001) {
		t.Errorf("p2 %v != expected %v", p2, exp)
	}
}

func TestOppToGlicko2(t *testing.T) {
	exp := []*Rating{
		NewRating(-0.5756, 0.1727, DefaultVol),
		NewRating(0.2878, 0.5756, DefaultVol),
		NewRating(1.1513, 1.7269, DefaultVol),
	}
	for i := range exp {
		g2 := opps[i].ToGlicko2()
		if !g2.MostlyEquals(exp[i], 0.0001) {
			t.Errorf("For i=%v: Glicko2 scaled opp %v != expected %v\n", i, g2, exp[i])
		}
	}
}

func TestEeGeeValues(t *testing.T) {
	expGee := []float64{0.9955,0.9531,0.7242}
	expEe := []float64{0.639,0.432,0.303}
	p2 := pl.ToGlicko2()
	for i := range opps {
		o := opps[i].ToGlicko2()
		geeVal := gee(o.Deviation)
		if !FloatsMostlyEqual(geeVal, expGee[i], 0.0001) {
			t.Errorf("Floats not mostly equal. g=%v exp_g=%v", geeVal, expGee[i])
		}

		eeVal := ee(p2.Rating, o.Rating, o.Deviation)
		if !FloatsMostlyEqual(eeVal, expEe[i], 0.001) {
			t.Errorf("Floats not mostly equal. ee=%v exp_ee=%v", eeVal, expEe[i])
		}
	}
}
