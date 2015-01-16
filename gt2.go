package gt2

import (
	"code.google.com/p/go-uuid/uuid"
)

func Out() string {
	return "B: " + uuid.NewRandom().String()
}
