package auth

import (
	"context"

	brevo "github.com/getbrevo/brevo-go/lib"
	"github.com/gofiber/fiber/v2"
)

func SendEmailWithCode(ctx *fiber.Ctx, code string, username string, email string) (bool, error) {
	var vanillaContext context.Context
	br := ctx.Locals("BrevoClient").(*brevo.APIClient)

	_, _, err := br.TransactionalEmailsApi.SendTransacEmail(vanillaContext, brevo.SendSmtpEmail{
		Sender: &brevo.SendSmtpEmailSender{
			Name:  "GoAlign",
			Email: "goalign.app@gmail.com",
		},
		To: []brevo.SendSmtpEmailTo{
			{
				Name:  username,
				Email: email,
			},
		},
		Subject:     "GoAlign Authorization Code",
		TextContent: CreateNonHtmlStringForUserAuth(username, code),
		HtmlContent: CreateEmailForUserAuth(username, code),
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
