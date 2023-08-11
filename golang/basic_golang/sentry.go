package main

import (
	"log"
	"time"
	"errors"
	"github.com/getsentry/sentry-go"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://1745618f6cbe4c22b58d1f5ba1ce433a@o4504376101765120.ingest.sentry.io/4504376104321024",
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	sentry.WithScope(func(scope *sentry.Scope) {

		// set context character
		scope.SetContext("character", map[string]interface{}{
			"name":        "Mighty Fighter",
			"age":         19,
			"attack_type": "melee",
		})

		sentry.AddBreadcrumb(&sentry.Breadcrumb{
			Category: "auth",
			Message: "Authenticated user " + "test@gmail.com",
			Level: sentry.LevelInfo,
		})

		scope.AddEventProcessor(func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			/* -> custom title exeption tidak bisa digunakan di capture message*/
			event.Exception[0].Type = "Testing change name sentry title"
			return event
		})

		//add extra data additional
		scope.SetExtra("extra_key", "extra_value")
		//set user id
		scope.SetUser(sentry.User{ID: "user_id_example"})
		//exemple error, fake err
		err = errors.New("Example Error 5")
		//capture error and level error
		scope.SetLevel(sentry.LevelError)
		sentry.CaptureException(err)
	})	
}

