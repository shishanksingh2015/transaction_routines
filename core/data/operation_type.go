package data

const (
	NormalPurchase           OperationType = 1
	PurchaseWithInstallments               = 2
	Withdrawal                             = 3
	CreditVoucher                          = 4
)

func IsValidOperationType(ot OperationType) bool {
	switch ot {
	case NormalPurchase, PurchaseWithInstallments, Withdrawal, CreditVoucher:
		return true
	default:
		return false
	}
}

func (ot OperationType) IsCreditVoucher() bool {
	return ot == CreditVoucher
}

func (ot OperationType) IsPurchaseOrWithdraw() bool {
	return ot == PurchaseWithInstallments || ot == Withdrawal || ot == NormalPurchase
}

func (ot OperationType) Int() int {
	return int(ot)
}
