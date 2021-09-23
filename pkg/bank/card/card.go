package card

import (
	"bank/pkg/bank/types"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	newCard := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return newCard
}

func Withdraw(card types.Card, amount types.Money) types.Card {
	if amount > card.Balance || card.Balance == 0 || amount < 0 || amount > 2000000 || !card.Active {
		return card
	}
	card.Balance -= amount
	return card
}

func Deposit(card *types.Card, amount types.Money) {
	if amount > 5000000 || amount < 0 || !card.Active {
		return
	}
	card.Balance += amount
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active || card.Balance <= 0 {
		return
	}
	bonus := types.Money(int(card.MinBalance) * percent / 100 * daysInMonth / daysInYear)
	if bonus > 5000 {
		bonus = 5000
	}
	card.Balance += bonus
}

func Total(cards []types.Card) types.Money {
	var total types.Money
	for _, card := range cards {
		if card.Balance >= 0 && card.Active {
			total += card.Balance
		}
	}
	return total
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var paySource []types.PaymentSource
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			paySource = append(paySource, types.PaymentSource{Type: "card", Number: string(card.PAN), Balance: card.Balance})
		}
	}
	return paySource
}
