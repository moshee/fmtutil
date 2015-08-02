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
