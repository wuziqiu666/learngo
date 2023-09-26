package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	var balance float64
	for _, v := range transactions {
		if name == v.From {
			balance -= v.Sum
		}
		if name == v.To {
			balance += v.Sum
		}
	}
	return balance
}
