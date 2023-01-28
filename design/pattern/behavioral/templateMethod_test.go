package behavioral

import "testing"

func TestTemplate(t *testing.T) {
	dp := &DepositBusinessHandler{userVip: true}
	executor := NewBankBusinessExecutor(dp)
	executor.ExecutorBankBusiness()
}
