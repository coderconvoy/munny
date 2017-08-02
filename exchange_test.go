package munny

import (
	"testing"
	"time"
)

func Test_Exchange_happy(t *testing.T) {
	rates := []rate{
		{USD, GBP, 1.5},
		{GBP, USD, 0.5},
	}
	be := BasicExchange{
		rates: rates,
		id:    "EX",
	}

	accounts := map[string]Account{

		"EX1": {"EX", "USD"},
		"EX2": {"EX", "GBP"},
		"F":   {"MT", "GBP"},
		"T":   {"RC", "USD"},
	}

	t1, t2, err := Exchange(100, "F", "T", accounts, be, time.Now())

	if err != nil {
		t.Error(err)
	}

	if t1.From != "F" {
		t.Errorf("t1 should be from F , Got : %s", t1.From)
	}

	if t2.Amount != 50 {
		t.Errorf("Expected 50, got %d", t2.Amount)
	}

}
