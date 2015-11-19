package stripeexample

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	stripe.Key = "sk_test_8bZtX27RlHVMwe0YexxgBB1s"
}

func chargeAccount(ctx context.Context, stripeToken string) error {
	// because we're on app engine, use a custom http client
	// this is being set globally, however
	// if we wanted to do it this way, we'd have to use a lock
	// https://youtu.be/KT4ki_ClX2A?t=1018
	hc := urlfetch.Client(ctx)
	stripe.SetHTTPClient(hc)
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
