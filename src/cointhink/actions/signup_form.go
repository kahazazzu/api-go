package actions

import (
	"log"
	"strings"

	"cointhink/db"
	"cointhink/mailer"
	"cointhink/model/account"
	"cointhink/model/credit_journal"
	"cointhink/proto"
	"cointhink/token"
	"cointhink/validate"

	gproto "github.com/golang/protobuf/proto"
)

func DoSignupform(form *proto.SignupForm) []gproto.Message {
	if form.Account != nil {
		rows, _ := db.D.Handle.Query(
			"select count(*) from accounts where email = $1",
			form.Account.Email)

		emailGood, emailFailReason := validate.Email(form.Account.Email)
		if emailGood == false {
			return []gproto.Message{&proto.SignupFormResponse{Ok: false,
				Reason:  proto.SignupFormResponse_EMAIL_ALERT,
				Message: emailFailReason}}
		}

		if rows.Next() {
			var count int
			rows.Scan(&count)
			if count > 0 {
				log.Printf("email check %d", count)
				return []gproto.Message{&proto.SignupFormResponse{Ok: false,
					Reason:  proto.SignupFormResponse_EMAIL_ALERT,
					Message: "email already in use"}}
			}
		}

		if len(strings.TrimSpace(form.Account.Username)) > 0 {
			rows, _ = db.D.Handle.Query(
				"select count(*) from accounts where username = $1",
				form.Account.Username)
			if rows.Next() {
				var count int
				rows.Scan(&count)
				if count > 0 {
					log.Printf("username check %d", count)
					return []gproto.Message{&proto.SignupFormResponse{Ok: false,
						Reason:  proto.SignupFormResponse_USERNAME_ALERT,
						Message: "email already in use"}}
				}
			}
		}

		err := account.Insert(form.Account)
		if err != nil {
			log.Printf("insert %+v", err)
			return []gproto.Message{&proto.SignupFormResponse{Ok: false}}
		} else {
			// Account success
			token := token.InsertToken(form.Account.Id)
			mailer.MailToken(token, form.Account.Email)
			c_err := credit_journal.Credit(form.Account, 2)
			if c_err != nil {
				log.Printf("credit_journal.Credit %+v", c_err)
			}
			return []gproto.Message{&proto.SignupFormResponse{Ok: true, Token: token}}
		}
	} else {
		return []gproto.Message{&proto.SignupFormResponse{Ok: false,
			Reason:  proto.SignupFormResponse_EMAIL_ALERT,
			Message: "missing email"}}
	}
}
