package gt2

import (
	"code.google.com/p/go-uuid/uuid"
)

func Out() string {
	return "A: " + uuid.NewRandom().String()
}
