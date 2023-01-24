package behavioral

import "testing"

func TestStrategy(t *testing.T) {
	wxPay  := &WXPay{}
	payCtx := NewPayCtx(wxPay)
	payCtx.Pay()
	// change strategy
	payCtx.setPayBehavior(&AliPay{})
	payCtx.Pay()
}
