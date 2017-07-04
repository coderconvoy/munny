package munny

import "testing"

func TestStringer(t *testing.T) {
	ts := []struct {
		m M
		s string
	}{
		{23, "0.23"},
		{-23, "-0.23"},
		{1, "0.01"},
		{-1, "-0.01"},
		{0, "0.00"},
		{403, "4.03"},
		{-403, "-4.03"},
	}

	for _, v := range ts {
		ss := v.m.String()
		if ss != v.s {
			t.Logf("with : %d, expected %s, got %s", int(v.m), v.s, ss)
			t.Fail()
		}
	}
}

func TestParse(t *testing.T) {
	ts := []struct {
		s string
		m M
		e bool
	}{
		{"hello", 0, true},
		{"0.01", 1, false},
		{"0.23", 23, false},
		{"-0.23", -23, false},
		{"-34", -3400, false},
		{"34.", 3400, false},
		{"34.2", 3420, false},
	}

	for k, v := range ts {
		n, err := ParseM(v.s)
		if n != v.m {
			t.Logf("LN%d:With %s, Expected %d, Got %d", k, v.s, int(v.m), int(n))
			t.Fail()
		}
		if (err != nil) != v.e {
			t.Logf("LN%d:With Errs, expected %b, got %s", k, v.e, err)
			t.Fail()
		}
	}

}
