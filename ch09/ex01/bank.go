package bank

type withdrawResult struct {
	amount int
	result bool
}

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan withdrawResult)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) (int, bool) {
	withdraws <- withdrawResult{amount, true}
	result := <-withdraws
	return result.amount, result.result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case result := <-withdraws:
			isWithdraw := balance > result.amount
			if isWithdraw {
				println(balance, result.amount)
				balance -= result.amount
			}
			withdraws <- withdrawResult{balance, isWithdraw}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
