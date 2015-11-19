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
	// this is being set globally, however
	// if we wanted to do it this way, we'd have to use a lock
	// https://youtu.be/KT4ki_ClX2A?t=1018
	hc := urlfetch.Client(ctx)
	stripe.SetHTTPClient(hc)

	id, _ := uuid.NewV4()

	for {
		chargeParams := &stripe.ChargeParams{
			Amount:   100 * 200000,
			Currency: "usd",
			Desc:     "Charge for test@example.com",
		}
		chargeParams.IdempotencyKey = id.String()
		chargeParams.SetSource(stripeToken)
		ch, err := charge.New(chargeParams)
		// https://youtu.be/KT4ki_ClX2A?t=1310
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				time.Sleep(time.Second)
				continue
			}
			return err
		}
		log.Infof(ctx, "CHARGE: %v", ch)
		log.Infof(ctx, "IDEMPOTENCY: %v", chargeParams.IdempotencyKey)
		return nil
	}
}
