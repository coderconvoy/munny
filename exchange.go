package munny

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const (
	USD = "USD"
	GBP = "GBP"
	CAD = "CAD"
)

type Account struct {
	Owner string
	Curr  string
}

type Exchanger interface {
	Exchange(M, string, string, time.Time) (M, error)
	Uid() string
}

type rate struct {
	cfrom, cto string
	conv       float32
}

type BasicExchange struct {
	rates []rate
	id    string
}

type Transaction struct {
	From   string
	To     string
	TrId   string
	LinkID string
	Amount M
}

var lastTransactionID = 0

func NewTransactionID() string {
	lastTransactionID++
	return strconv.Itoa(lastTransactionID)
}

func (be BasicExchange) Exchange(am M, fcur, tcur string, t time.Time) (M, error) {
	for _, v := range be.rates {
		if v.cfrom == fcur && v.cto == tcur {
			return M(float32(am) * v.conv), nil
		}
	}
	return 0, errors.Errorf("No Rate Available from %s to %s", fcur, tcur)
}

func (be BasicExchange) Uid() string {
	return be.id
}

func Exchange(am M, frid, toid string, aclist map[string]Account, ex Exchanger, t time.Time) (Transaction, Transaction, error) {
	bad := Transaction{}

	frac, ok := aclist[frid]
	if !ok {
		return bad, bad, errors.Errorf("From Account not found %s", frid)
	}
	toac, ok := aclist[toid]
	if !ok {
		return bad, bad, errors.Errorf("To Account not found: %s", toid)
	}

	resAmount, err := ex.Exchange(am, frac.Curr, toac.Curr, t)

	if err != nil {
		return bad, bad, errors.Wrap(err, "Exchange refused to Exchange")
	}

	exid := ex.Uid()
	intoid := ""
	outofid := ""
	for k, v := range aclist {
		if v.Owner == exid && v.Curr == frac.Curr {
			intoid = k
		}
		if v.Owner == exid && v.Curr == toac.Curr {
			outofid = k
		}
	}

	if intoid == "" {
		return bad, bad, errors.Errorf("Exchange cannot recieve Currency : %s", frac.Curr)
	}
	if outofid == "" {
		return bad, bad, errors.Errorf("Exchange cannot pay out Currency : %s", toac.Curr)
	}
	trid1 := NewTransactionID()
	trid2 := NewTransactionID()

	intran := Transaction{
		From:   frid,
		To:     intoid,
		Amount: am,
		TrId:   trid1,
		LinkID: trid2,
	}
	outtran := Transaction{
		From:   toid,
		To:     outofid,
		Amount: resAmount,
		TrId:   trid2,
		LinkID: trid1,
	}

	return intran, outtran, nil //TODO
}
