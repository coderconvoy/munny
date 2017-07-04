package munny

import (
	"encoding/json"
	"testing"
)

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

func TestMarshal(t *testing.T) {
	td := []M{
		23, 34, 121, -604, 0, -1,
	}

	md, err := json.Marshal(td)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	var sd []M

	err = json.Unmarshal(md, &sd)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	for k, v := range td {
		if sd[k] != v {
			t.Log("Expected %d, got %d", int(v), int(sd[k]))
			t.Fail()

		}
	}
}
