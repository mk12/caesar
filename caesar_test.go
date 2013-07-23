// Copyright 2013 Mitchell Kember. Subject to the MIT License.

package main

import "testing"

var encryptTests = []struct {
	shift  int
	input  string
	output string
}{
	{0, "The quick brown fox...", "The quick brown fox..."},
	{1, "abcdefghijklmnopqrstuvwxyz", "bcdefghijklmnopqrstuvwxyza"},
	{13, "This is ROT-13.", "Guvf vf EBG-13."},
	{567, "1234567890-=/!@#$%^&*()_+?", "1234567890-=/!@#$%^&*()_+?"},
}

func TestEncrypt(t *testing.T) {
	for i, test := range encryptTests {
		output := []byte(test.input)
		encrypt(output, test.shift)
		strOutput := string(output)
		if strOutput != test.output {
			t.Errorf("%d. encrypt([]byte(%q), %d)\n"+
				"returned []byte(%q)\nexpected []byte(%q)",
				i, test.input, test.shift, strOutput, test.output)
		}
	}
}

var freqTests = []struct {
	text  string
	freqs []float64
}{
	{"", []float64{25: 0}},
	{"Aa", []float64{1.0, 25: 0}},
	{"cjkZ", []float64{2: 0.25, 9: 0.25, 10: 0.25, 25: 0.25}},
	{"The quick brown fox jumps over the lazy dog.", []float64{1.0 / 35,
		1.0 / 35, 1.0 / 35, 1.0 / 35, 3.0 / 35, 1.0 / 35, 1.0 / 35, 2.0 / 35,
		1.0 / 35, 1.0 / 35, 1.0 / 35, 1.0 / 35, 1.0 / 35, 1.0 / 35, 4.0 / 35,
		1.0 / 35, 1.0 / 35, 2.0 / 35, 1.0 / 35, 2.0 / 35, 2.0 / 35, 1.0 / 35,
		1.0 / 35, 1.0 / 35, 1.0 / 35, 1.0 / 35}},
}

func TestFrequencies(t *testing.T) {
	for i, test := range freqTests {
		freqs := frequencies([]byte(test.text))
		for j, freq := range freqs {
			if j >= len(test.freqs) || freq != test.freqs[j] {
				t.Errorf("%d. frequencies([]byte(%q))\nreturned %v\n"+
					"expected %v", i, test.text, freqs, test.freqs)
				break
			}
		}
	}
}

var chisqrTests = []struct {
	freqs []float64
	rot   int
	min   float64
	max   float64
}{
	{[]float64{25: 0}, 0, 0.9, 1.0},
	{frequencies([]byte("abcxyz")), 0, 59.6, 59.7},
	{frequencies([]byte("qz")), 1, 24.6, 24.7},
}

func TestChisqr(t *testing.T) {
	for i, test := range chisqrTests {
		chi := chisqr(test.freqs, test.rot)
		if chi < test.min || chi > test.max {
			t.Errorf("%d. chisqr(%#v, %d)\nreturned %f\nexpected %f to %f",
				i, test.freqs, test.rot, chi, test.min, test.max)
		}
	}
}

var crackTests = []struct {
	msg   string
	shift int
}{
	{"The quick brown fox jumps over the lazy dog.", 0},
	{"guvf vf cresrpgyl angheny Ratyvfu", 13},
	{"uif Dbftbs djqifs dbo cf dsbdlfe vtjoh gsfrvfodz bobmztjt", 1},
}

func TestCrack(t *testing.T) {
	for i, test := range crackTests {
		shift := crack([]byte(test.msg))
		if shift != test.shift {
			t.Errorf("%d. crack([]byte(%q))\nreturned %d\nexpected %d",
				i, test.msg, shift, test.shift)
		}
	}
}
