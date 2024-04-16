package commands

import (
	"github.com/FedorKowarnow/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message, productService *product.Service) {
	outputMSGText := "Here all the products: \n\n"
	products := c.productService.List()
	for _, p := range products {
		outputMSGText += p.Title
		outputMSGText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMSGText)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(

			tgbotapi.NewInlineKeyboardButtonData("Next page", "some data"),
		),
	)

	c.bot.Send(msg)
}
