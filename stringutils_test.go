// Copyright 2015 Mathieu MAST. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package goloc

import (
	"testing"
)

var levTests = []struct {
	source     string
	target     string
	ignoreCase bool
	distance   int
}{
	{"", "", false, 0},
	{"PARIS", "", false, 5},
	{"", "PARIS", false, 5},
	{"PARIS", "PARIS", false, 0},
	{"PaRIS", "PARiS", false, 2},
	{"PaRIS", "PARiS", true, 0},
	{"PARIS", "PARI", false, 1},
	{"PARS", "PARIS", false, 1},
	{"PAR", "PARIS", false, 2},
	{"PR", "PARIS", false, 3},
	{"PARIS", "FRANCE", false, 5},
	{"PĂRIS", "PARIŞ", false, 2},
	{"CAR", "Carpeaux", true, 5},
	{"Élysées", "elysees", false, 2},
	{"Élysées", "elysees", true, 0},
}

func TestLevenshteinDistance(t *testing.T) {
	for _, tt := range levTests {
		d := LevenshteinDistance(tt.source, tt.target, tt.ignoreCase)
		if d != tt.distance {
			t.Logf("Distance of %v and %v should be %v but was %v.",
				tt.source, tt.target, tt.distance, d)
			t.Fail()
		}
	}
}

var partialphoneTests = []struct {
	source string
	target string
}{
	{"", ""},
	{"Brest", "BRST"},
	{"Avenue des Champs-Élysées 75008 Paris France", "AVN DS SPS ALSS 7 PRS FRNS"},
	{"Rue du Square Carpeaux 75018 Paris France", "R D SKR KRPS 7 PRS FRNS"},
	{"Place Carnot 69002 Lyon", "PLS KRNT 6 LN"},
	{"a23 3423u 78045R 345TE", "A 2 3 A 7 R 3 T"},
	{" ", ""},
	{"!", ""},
	{"? !", ""},
	{"t", "T"},
	{"i", "A"},
	{"cacécîcocücy", "KSSKKS"},
}

func TestPartialphone(t *testing.T) {
	for _, tt := range partialphoneTests {
		target := Partialphone(tt.source)
		if target != tt.target {
			t.Logf("Partialphone of %v should be %v but was %v.",
				tt.source, tt.target, target)
			t.Fail()
		}
	}
}

func TestScore(t *testing.T) {
	score1 := Score("champs elysees paris", "Avenue des Champs-Élysées 75008 Paris France")
	score2 := Score("paris champs elysees", "Avenue des Champs-Élysées 75008 Paris France")
	score3 := Score("champs elyse paris", "Avenue des Champs-Élysées 75008 Paris France")
	if score1 <= score2 || score1 <= score3 || score1 != 1000 {
		t.Logf("score1=%v score2=%v score3=%v", score1, score2, score3)
		t.Fail()
	}
}
