package main

import "fmt"

/*
	Provides a simple interface to complex subsystems with many active parts.
*/

var (
	creditType = "credit"
	debitType  = "debit"
)

//facade wraps several components and provides add/deduct operations for clients
type walletFacade struct {
	account      *account
	securityCode *securityCode
	wallet       *wallet
	ledger       *ledger
	notification *notification
}

func newWalletFacade(accountID string, code int) *walletFacade {
	fmt.Printf("Create account\n")
	return &walletFacade{
		account: &account{
			name: accountID,
		},
		securityCode: &securityCode{
			code: code,
		},
		wallet:       &wallet{},
		ledger:       &ledger{},
		notification: &notification{},
	}
}

func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.ledger.makeEntry(accountID, creditType, amount)
	return nil
}

func (w *walletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.debitBalance(amount)
	w.ledger.makeEntry(accountID, debitType, amount)
	return nil
}

type account struct {
	name string
}

func (a *account) checkAccount(name string) error {
	if a.name != name {
		return fmt.Errorf(" account name incorrect")
	} else {
		fmt.Printf("Account Verified\n")
	}
	return nil
}

type securityCode struct {
	code int
}

func (s *securityCode) checkCode(code int) error {
	if s.code != code {
		return fmt.Errorf(" security code incorrect")
	} else {
		fmt.Printf("SecurityCode Verified\n")
	}
	return nil
}

type wallet struct {
	balance int
}

func (w *wallet) creditBalance(amount int) {
	w.balance += amount
	fmt.Printf("Wallet balance added successful for $ %d\n", amount)
}
func (w *wallet) debitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("balance is insufficient")
	}
	w.balance -= amount
	fmt.Printf("Wallet balance minused successful for $ %d\n", amount)
	return nil
}

type ledger struct{}

func (l *ledger) makeEntry(account, taxType string, amount int) {
	fmt.Printf("Making ledger record : account %s with taxType %s for amount $%d\n", account, taxType, amount)
}

type notification struct{}

func (n *notification) sendWalletCreditNotification(account string) {
	fmt.Printf("Sending wallet credit notification to account %s\n", account)
}

func (n *notification) sendWalletDebitNotification(account string) {
	fmt.Printf("Sending wallet debit notification to account %s\n", account)
}

func RunFacade() {
	PersonalServer := newWalletFacade("xiao ming", 123456)
	PersonalServer.addMoneyToWallet("xiao ming", 123456, 10)

	PersonalServer.deductMoneyFromWallet("xiao ming", 123456, 5)
}
