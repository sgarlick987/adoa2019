package main

import "testing"

func TestSevenish(t *testing.T) {
	s := sevenishNumber(1)
	if s != 1 {
		t.Errorf("expected 1 for sevenish number 1 but got %d", s)
	}
	s = sevenishNumber(2)
	if s != 7 {
		t.Errorf("expected 7 for sevenish number 2 but got %d", s)
	}
	s = sevenishNumber(3)
	if s != 8 {
		t.Errorf("expected 8 for sevenish number 3 but got %d", s)
	}
	s = sevenishNumber(4)
	if s != 49 {
		t.Errorf("expected 49 for sevenish number 4 but got %d", s)
	}
	s = sevenishNumber(5)
	if s != 50 {
		t.Errorf("expected 50 for sevenish number 5 but got %d", s)
	}
	s = sevenishNumber(6)
	if s != 56 {
		t.Errorf("expected 56 for sevenish number 6 but got %d", s)
	}
	s = sevenishNumber(7)
	if s != 343 {
		t.Errorf("expected 343 for sevenish number 7 but got %d", s)
	}
	s = sevenishNumber(8)
	if s != 344 {
		t.Errorf("expected 344 for sevenish number 8 but got %d", s)
	}
	s = sevenishNumber(9)
	if s != 350 {
		t.Errorf("expected 350 for sevenish number 9 but got %d", s)
	}
	s = sevenishNumber(10)
	if s != 392 {
		t.Errorf("expected 392 for sevenish number 10 but got %d", s)
	}
	s = sevenishNumber(100)
	if s != 96889833950 {
		t.Errorf("expected 96889833950 for sevenish number 100 but got %d", s)
	}
	s = sevenishNumber(9223372036860540609)
	if s != 9223372036860540609 {
		t.Errorf("expected 9223372036860540609 for sevenish number 100 but got %d", s)
	}
}
