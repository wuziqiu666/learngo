package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func applyTransaction(transaction Transaction, a Account) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}

func NewBalanceFor(account Account, transaction []Transaction) Account {
	return Reduce(transaction, applyTransaction, account)
}
func Reduce[A, B any](collections []A, accumlator func(A, B) B, initialValue B) B {
	result := initialValue
	for _, v := range collections {
		result = accumlator(v, result)
	}
	return result
}
