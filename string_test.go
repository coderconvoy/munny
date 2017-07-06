package munny

import (
	"encoding/json"
	"testing"
)

func Test_Stringer(t *testing.T) {
	ts := []struct {
		m M
		s string
	}{
		{0, "0.00"},
		{1, "0.01"},
		{-1, "-0.01"},
		{12, "0.12"},
		{-12, "-0.12"},
		{123, "1.23"},
		{-123, "-1.23"},
		{100, "1.00"},
		{1200, "12.00"},
	}

	for _, v := range ts {
		ss := v.m.String()
		if ss != v.s {
			t.Logf("With %d Expected, %s, Got %s", v.m, v.s, ss)
			t.Fail()
		}
	}
}

func Test_Parse(t *testing.T) {
	ts := []struct {
		s string
		m M
		e bool
	}{
		{"0.00", 0, false},
		{"0.01", 1, false},
		{"-0.01", -1, false},
		{"0.12", 12, false},
		{"-0.12", -12, false},
		{"1.23", 123, false},
		{"-1.23", -123, false},
		{"1.00", 100, false},
		{"12.00", 1200, false},
		{"hello", 0, true},
		{"23.", 2300, false},
		{"1.2.3", 0, true},
	}

	for k, v := range ts {
		mm, err := ParseM(v.s)
		if mm != v.m {
			t.Logf("with %s, expected %d, got %d", v.s, v.m, mm)
			t.Fail()
		}

		if (err != nil) != v.e {
			t.Logf("ERRLN %d, With %s, expected %b, got %s", k, v.s, v.e, err)
			t.Fail()
		}
	}

}

func Test_Marshal(t *testing.T) {
	ts := []M{
		0, 1, -1, 12, -12, 123, -123, 1234, -1234,
	}
	mm, err := json.Marshal(ts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	var t2 []M

	err = json.Unmarshal(mm, &t2)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	for k, v := range ts {
		if t2[k] != v {
			t.Log("With %d, got %d", v, t2[k])
			t.Fail()
		}
	}

}
