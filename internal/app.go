package internal

import (
	"fmt"
	"sync"

	"github.com/alexplisov/go-tg-card-bot/internal/bot"
	"github.com/alexplisov/go-tg-card-bot/internal/config"
)

func Run() error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading config: %w", err)
	}

	botHandler, err := bot.SetupBotHandler(cfg.TelegramBotToken, cfg.OwnerChatID)
	if err != nil {
		return fmt.Errorf("error setting up bot: %w", err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer botHandler.Stop()
		botHandler.Start()
	}()

	wg.Wait()

	return nil
}
