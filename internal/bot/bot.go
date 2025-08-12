package bot

import (
	"context"

	"github.com/alexplisov/go-tg-card-bot/internal/bot/conversations"
	"github.com/alexplisov/go-tg-card-bot/internal/bot/handlers"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
)

func SetupBotHandler(token string, ownerChatID int64) (*telegohandler.BotHandler, error) {
	bot, err := telego.NewBot(token, telego.WithDefaultDebugLogger())
	if err != nil {
		return nil, err
	}
	updates, _ := bot.UpdatesViaLongPolling(context.Background(), nil)
	handler, err := telegohandler.NewBotHandler(bot, updates)
	if err != nil {
		return nil, err
	}

	conversationStates := make(map[int64]*conversations.OrderConversationState)

	handler.Handle(handlers.StartHandler, telegohandler.CommandEqual("start"))
	handler.HandleCallbackQuery(handlers.HandleInlineMenuEducation, telegohandler.CallbackDataEqual("education"))

	orderHandler := handlers.NewOrderHandler(conversationStates)
	handler.HandleCallbackQuery(orderHandler.Handle, telegohandler.CallbackDataEqual("order"))

	handler.HandleCallbackQuery(handlers.HandleInlineMenuBack, telegohandler.CallbackDataEqual("back"))

	messageHandler := handlers.NewMessageHandler(ownerChatID, conversationStates)
	handler.Handle(messageHandler.Handle, telegohandler.AnyMessage())
	return handler, nil
}
