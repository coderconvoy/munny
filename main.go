package munny

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type M int

func ParseM(s string) (M, error) {
	s = strings.Trim(s, "\n\r \"'\t")
	var minus M = 1
	if s[0] == '-' {
		s = s[1:]
		minus = -1
	}

	ss := strings.Split(s, ".")
	if len(ss) > 2 {
		return 0, errors.New("Too many dots")

	}
	pnds, err := strconv.Atoi(ss[0])
	if err != nil {
		return 0, err
	}
	if len(ss) == 1 {
		return minus * 100 * M(pnds), nil
	}

	if len(ss[1]) == 0 {
		return minus * 100 * M(pnds), nil
	}

	pennies, err := strconv.Atoi(ss[1])
	if err != nil {
		return 0, errors.Wrap(err, "Could not get pennies")
	}
	switch len(ss[1]) {
	case 1:
		return minus * M(100*pnds+10*pennies), nil
	case 2:
		return minus * M(100*pnds+pennies), nil
	}

	return 0, errors.New("Pennies too many dp")
}

func SafeParseM(s string, def M) M {
	m, err := ParseM(s)
	if err != nil {
		return def
	}
	return m
}

func (m M) String() string {
	s := strconv.Itoa(int(m))

	switch len(s) {
	case 1:
		return "0.0" + s
	case 2:
		if s[0] == '-' {
			return "-0.0" + s[1:]
		}
		return "0." + s
	case 3:
		if s[0] == '-' {
			return "-0." + s[1:]
		}
	}
	return s[:len(s)-2] + "." + s[len(s)-2:]

}

func (m M) MarshalJSON() ([]byte, error) {
	s := m.String()
	return []byte(`"` + s + `"`), nil
}

func (m *M) UnmarshalJSON(b []byte) error {
	v, err := ParseM(string(b))
	if err != nil {
		return err
	}
	*m = v
	return nil
}
