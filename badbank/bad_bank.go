package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

type Account struct {
	Name    string
	Balance float64
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

func Reduce[A, B any](collections []A, accumlator func(A, B) B, initialValue B) B {
	result := initialValue
	for _, v := range collections {
		result = accumlator(v, result)
	}
	return result
}
