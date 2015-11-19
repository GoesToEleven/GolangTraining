package stripeexample

import (
	"net"
	"time"

	"github.com/nu7hatch/gouuid"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	stripe.Key = "sk_test_oh4N9TLop9gVlVQ1N05IGamg"
}

func chargeAccount(ctx context.Context, stripeToken string) error {
	// because we're on app engine, use a custom http client
	b := stripe.BackendConfiguration{
		stripe.APIBackend,
		"https://api.stripe.com/v1",
		urlfetch.Client(ctx),
	}

	id, _ := uuid.NewV4()

	for {
		chargeParams := &stripe.ChargeParams{
			Amount:   100 * 200000,
			Currency: "usd",
			Desc:     "Charge for test@example.com",
		}
		chargeParams.IdempotencyKey = id.String()
		chargeParams.SetSource(stripeToken)
		chargeClient := &charge.Client{
			Key: stripe.Key,
			B:   b,
		}
		ch, err := chargeClient.New(chargeParams)
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				time.Sleep(time.Second)
				continue
			}
			return err
		}
		log.Infof(ctx, "CHARGE: %v", ch)
		return nil
	}
}
