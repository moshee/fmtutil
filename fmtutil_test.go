package fmtutil

import (
	"testing"
	"time"
)

var hmsTest = []struct {
	d time.Duration
	s string
}{
	{0, "0:00"},
	{time.Millisecond, "0:00"},
	{time.Second, "0:01"},
	{3*time.Second + 500*time.Millisecond, "0:03"},
	{59 * time.Second, "0:59"},
	{time.Minute, "1:00"},
	{time.Minute + 24*time.Second, "1:24"},
	{10 * time.Minute, "10:00"},
	{24*time.Minute + 42*time.Second, "24:42"},
	{59 * time.Minute, "59:00"},
	{60 * time.Minute, "1:00:00"},
	{65 * time.Minute, "1:05:00"},
	{70 * time.Minute, "1:10:00"},
	{80*time.Minute + 7*time.Second, "1:20:07"},
	{120*time.Minute + 24*time.Second, "2:00:24"},
	{3600 * time.Minute, "60:00:00"},
	{6000 * time.Minute, "100:00:00"},
}

func TestHMS(t *testing.T) {
	for _, pair := range hmsTest {
		if s := HMS(pair.d); s != pair.s {
			t.Errorf("HMS: have '%s', want '%s'", s, pair.s)
		}
	}
}

func TestFormatRoman(t *testing.T) {
	romanTest := []struct {
		n int
		s string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{12, "XII"},
		{123, "CXXIII"},
		{1234, "MCCXXXIV"},
		{403, "CDIII"},
		{1823, "MDCCCXXIII"},
		{983, "CMLXXXIII"},
		{391, "CCCXCI"},
		{882, "DCCCLXXXII"},
		{533, "DXXXIII"},
		{24, "XXIV"},
		{95, "XCV"},
		{38, "XXXVIII"},
		{50, "L"},
		{51, "LI"},
		{69, "LXIX"},
	}

	for _, test := range romanTest {
		res := FormatRoman(test.n)
		if res != test.s {
			t.Errorf("FormatRoman(%d): want %q, got %q", test.n, test.s, res)
		}
	}
}
