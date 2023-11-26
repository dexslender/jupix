package util

import (
	"strings"

	"github.com/disgoorg/disgo/discord"
)

type Order func(ModalCtx)

func Send(msg string) Order {
	return func(ctx ModalCtx) {
		ctx.CreateMessage(discord.NewMessageCreateBuilder().
			SetContent(msg).
			Build(),
		)
		// ctx.Client().Rest().CreateMessage(ctx.Channel().ID(), discord.NewMessageCreateBuilder().
		// 	SetContent(msg).
		// 	Build(),
		// )
	}
}

func Parse(input string) (o []Order) {
	var (
		b             strings.Builder
		msg_read      bool
		not_fst_space bool
	)
	for i, v := range input {
		if msg_read {
			if !not_fst_space && v == ' ' || v == '	' {
				not_fst_space = true
			} else {
				not_fst_space = true
				b.WriteRune(v)
			}

			if v == '\n' || v == '.' || len(input)-1 == i {
				o = append(o, Send(b.String()))
				msg_read = false
				not_fst_space = false
				b.Reset()
			}
			continue
		}
		switch v {
		case '>':
			msg_read = true
		}
	}
	return
}
