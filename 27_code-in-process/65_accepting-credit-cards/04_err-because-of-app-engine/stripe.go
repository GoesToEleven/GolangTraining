package stripeexample

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

func init() {
	stripe.Key = "sk_test_8bZtX27RlHVMwe0YexxgBB1s"
}

func chargeAccount(ctx context.Context, stripeToken string) error {
	chargeParams := &stripe.ChargeParams{
		Amount:   100 * 200000,
		Currency: "usd",
		Desc:     "Charge for test@example.com",
	}
	chargeParams.SetSource(stripeToken)
	ch, err := charge.New(chargeParams)
	if err != nil {
		return err
	}
	log.Infof(ctx, "CHARGE: %v", ch)
	return nil
}
