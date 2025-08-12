package handlers

import (
	"github.com/alexplisov/go-tg-card-bot/internal/bot/conversations"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"github.com/mymmrac/telego/telegoutil"
)

type OrderHandler struct {
	states map[int64]*conversations.OrderConversationState
}

func NewOrderHandler(states map[int64]*conversations.OrderConversationState) *OrderHandler {
	return &OrderHandler{
		states: states,
	}
}

func (h *OrderHandler) Handle(ctx *telegohandler.Context, query telego.CallbackQuery) error {
	h.states[query.From.ID] = conversations.NewOrderConversationState()
	if _, err := ctx.Bot().SendMessage(ctx, telegoutil.Message(telegoutil.ID(query.From.ID), "Как к тебе обращаться?")); err != nil {
		return err
	}
	return ctx.Bot().AnswerCallbackQuery(ctx, telegoutil.CallbackQuery(query.ID))
}
