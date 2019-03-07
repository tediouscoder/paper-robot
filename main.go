package main

import (
	"context"
	"github.com/tediouscoder/paper-robot/utils"
	"net/http"
	"os"

	"gopkg.in/go-playground/webhooks.v5/github"

	"github.com/tediouscoder/paper-robot/internal/log"
	"github.com/tediouscoder/paper-robot/paper"
)

func main() {
	log.SetLevel(log.LevelDebug)

	hook, err := github.New(github.Options.Secret(os.Getenv("WEBHOOK_SECRET")))
	if err != nil {
		log.Error("Init github webhook failed", "error", err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		payload, err := hook.Parse(r, github.IssuesEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				log.Info("Event is not an issue event, ignore.")
				return
			}

			log.Error("Parse event failed", "error", err)
			return
		}

		switch v := payload.(type) {
		case github.IssuesPayload:
			ctx := utils.NewEventMetadataContext(ctx, &utils.EventMetadata{
				Owner:       v.Repository.Owner.Login,
				Repo:        v.Repository.Name,
				IssueNumber: v.Issue.Number,
			})

			err = paper.Handler(ctx, &v)
			if err != nil {
				log.Error("Robot failed to handle this event", "error", err)
				return
			}
		default:
			log.Error("Payload should be issue, but it not", "payload", v)
		}

		return
	})

	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Error("Server server failed", "error", err)
	}
}
