package munny

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type M int

func SafeParseM(s string, def M) M {
	m, err := ParseM(s)
	if err != nil {
		return def
	}
	return m
}

func parsePos(s string) (M, error) {
	ss := strings.Split(s, ".")

	if len(ss) > 2 {
		return 0, errors.New("Too many dots")
	}

	if len(ss) == 1 {
		i, err := strconv.Atoi(s)
		return M(i) * 100, err
	}

	i, err := strconv.Atoi(ss[0])
	if err != nil {
		return 0, err
	}
	res := M(i) * 100

	if len(ss[1]) == 0 {
		return res, nil
	}
	n, err := strconv.Atoi(ss[1])

	if err != nil {
		return 0, err
	}

	switch len(ss[1]) {
	case 0:
		return res, nil
	case 1:
		return res + M(n)*10, nil
	case 2:
		return res + M(n), nil
	}
	return 0, errors.New("Too much after the dot")
}

func ParseM(s string) (M, error) {
	s = strings.TrimSpace(s)
	minus := strings.HasPrefix(s, "-")
	if minus {
		s = s[1:]
	}

	m, err := parsePos(s)
	if minus {
		return -m, err
	}
	return m, err

}

func (m M) MarshalJSON() ([]byte, error) {
	return []byte(`"` + m.String() + `"`), nil
}

func (m *M) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = strings.Trim(s, "\"\t \n\r")
	v, err := ParseM(s)
	if err != nil {
		return err
	}
	*m = v
	return nil
}

func (m M) String() string {
	pm := m
	if m < 0 {
		pm = -m
	}
	s := strconv.Itoa(int(pm))

	res := ""

	switch len(s) {
	case 1:
		res = "0.0" + s
	case 2:
		res = "0." + s
	default:
		res = s[:len(s)-2] + "." + s[len(s)-2:]
	}

	if m < 0 {
		res = "-" + res
	}

	return res
}
