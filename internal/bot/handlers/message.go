package handlers

import (
	"fmt"
	"strings"

	"github.com/alexplisov/go-tg-card-bot/internal/bot/conversations"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"github.com/mymmrac/telego/telegoutil"
)

type MessageHandler struct {
	ownerChatID int64
	states      map[int64]*conversations.OrderConversationState
}

func NewMessageHandler(ownerChatID int64, states map[int64]*conversations.OrderConversationState) *MessageHandler {
	return &MessageHandler{
		ownerChatID: ownerChatID,
		states:      states,
	}
}

func (h *MessageHandler) Handle(ctx *telegohandler.Context, update telego.Update) error {
	state, ok := h.states[update.Message.From.ID]
	if !ok {
		return nil
	}
	switch state.OrderStep {
	case conversations.OrderStepName:
		if update.Message.Text != "" {
			state.OrderStep = conversations.OrderStepRequirements
			state.UserName = update.Message.Text
			if _, err := ctx.Bot().SendMessage(ctx, telegoutil.Message(update.Message.Chat.ChatID(), "Отлично! А теперь опиши какого бота ты хочешь сделать?")); err != nil {
				return err
			}
		}
	case conversations.OrderStepRequirements:
		if update.Message.Text != "" {
			defer delete(h.states, update.Message.From.ID)
			state.Requirements = update.Message.Text
			if _, err := ctx.Bot().SendMessage(ctx, telegoutil.Message(update.Message.Chat.ChatID(), "Спасибо за заявку! Я отвечу на столько быстро, на сколько смогу! 😉")); err != nil {
				return err
			}
			// send order to owner
			var sb strings.Builder
			sb.WriteString("🛒 <b>Новый заказ</b>\n\n")
			fmt.Fprintf(&sb, "👤 <b>От</b>\n\n%s (@%s)\n\n", state.UserName, update.Message.Chat.Username)
			sb.WriteString("📝 <b>Описание</b>\n\n")
			sb.WriteString(state.Requirements)
			if _, err := ctx.Bot().SendMessage(ctx, telegoutil.Message(telegoutil.ID(h.ownerChatID), sb.String()).WithParseMode(telego.ModeHTML)); err != nil {
				return err
			}
		}
	}
	return nil
}
