package discordbot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session *discordgo.Session
	guildID string
	OnReady func(s *discordgo.Session)
	OnStop  func()
}

func New(token string, guildId string) (*Bot, error) {
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		return &Bot{}, err
	}

	return &Bot{
		session: session,
		guildID: guildId,
		OnReady: func(s *discordgo.Session) {
			log.Println(fmt.Sprintf("%s is ready", s.State.User.Username))
		},
		OnStop: func() {
			log.Println("Shutting down.")
		},
	}, nil
}

func (b *Bot) Start() {
	log.Println("Start()")

	b.session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		b.OnReady(s)
	})

	err := b.session.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer b.session.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	b.OnStop()
}
