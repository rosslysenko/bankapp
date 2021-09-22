package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

var testCard = types.Card{
	Balance:    10_000_00,
	MinBalance: 10_000_00,
	Active:     true,
}

var testCard2 = types.Card{
	Balance:    0,
	MinBalance: 0,
}

var testCard3 = types.Card{
	Balance:    -10,
	MinBalance: -12,
	Active:     true,
}

var testCard4 = types.Card{
	Balance:    10_000_00000000,
	MinBalance: 10_000_00000000,
	Active:     true,
}

func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 1000000
}

func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 0
}

func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 0, Active: false}, 10_000_00)
	fmt.Println(result.Balance)
	// Output: 0
}

func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 25_000_00, Active: true}, 20_000_00)
	fmt.Println(result.Balance)
	// Output: 500000
}

func ExampleDeposit_positive() {
	Deposit(&testCard, 10_000_00)
	fmt.Println(testCard.Balance)
	// Output: 2000000
}

func ExampleDeposit_limit() {
	Deposit(&testCard, 50_000_01)
	fmt.Println(testCard.Balance)
	// Output: 2000000
}

func ExampleDeposit_inactive() {
	Deposit(&testCard2, 10_000_00)
	fmt.Println(testCard2.Balance)
	// Output: 0
}

func ExampleAddBonus_positive() {
	AddBonus(&testCard, 3, 30, 365)
	fmt.Println(testCard.Balance)
	// Output: 2002465
}

func ExampleAddBonus_inactive() {
	AddBonus(&testCard2, 3, 30, 365)
	fmt.Println(testCard2.Balance)
	// Output: 0
}

func ExampleAddBonus_negative() {
	AddBonus(&testCard3, 3, 30, 365)
	fmt.Println(testCard3.Balance)
	// Output: -10
}

func ExampleAddBonus_overLimit() {
	AddBonus(&testCard4, 3, 30, 365)
	fmt.Println(testCard4.Balance)
	// Output: 1000000005000
}

func ExampleTotal() {
	fmt.Println(Total([]types.Card{
		{Balance: 100000,
			Active: true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{Balance: 100000,
			Active: true,
		},
		{Balance: 200000,
			Active: true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 100000,
			Active:  false,
		},
	}))
	fmt.Println(Total([]types.Card{
		{Balance: -100000,
			Active: true,
		},
	}))
	// Output:
	// 100000
	// 300000
	// 0
	// 0

}
