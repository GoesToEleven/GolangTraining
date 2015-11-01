package payment

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
)

const secretKey = "sk_test_C4cVqiMxarOi7drYWgUsXmLr"

func chargeAccount(ctx context.Context, stripeToken string) error {
	b := stripe.BackendConfiguration{stripe.APIBackend, "https://api.stripe.com/v1", urlfetch.Client(ctx)}
	c := &charge.Client{Key: secretKey, B: b}
	chargeParams := &stripe.ChargeParams{
		Amount:   20000000,
		Currency: "usd",
		Desc:     "Charge for test@example.com",
	}
	chargeParams.SetSource(stripeToken)
	_, err := c.New(chargeParams)
	if err != nil {
		return err
	}
	return nil
}
