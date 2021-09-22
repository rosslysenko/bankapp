package types

// Money is a sum of money in minimum units (cents, kopecks, dirams)
type Money int64

// Currency represents the currency code
type Currency string

// The currency codes
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN represents the payment card number
type PAN string

// Card represents the bank card
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

// Payment represents info about payment oparations
type Payment struct {
	ID     int
	Amount Money
}

type PaymentSource struct {
	Type    string // 'card'
	Number  string // number '5058 xxxx xxxx 8888'
	Balance Money  // balance in dirams
}
