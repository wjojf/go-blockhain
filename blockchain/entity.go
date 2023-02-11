package blockchain

type TransactionData struct {
	from   string
	to     string
	amount int64
}

func NewTransaction(from, to string, amount int64) TransactionData {
	return TransactionData{
		from:   from,
		to:     to,
		amount: amount,
	}
}
