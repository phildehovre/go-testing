package generics2

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	adjustBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			return currentBalance - t.Sum
		}
		if t.To == name {
			return currentBalance + t.Sum

		}
		return currentBalance
	}
	return Reduce(transactions, adjustBalance, 0.0)
}

func Reduce[A, B any](a []A, fn func(B, A) B, initialValue B) B {
	result := initialValue
	for _, i := range a {
		result = fn(result, i)
	}
	return result
}
