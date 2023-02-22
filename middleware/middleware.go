package middleware

import (
	"github.com/ahmedkhaled/jazz"
	"myapp/data"
)

type Middleware struct {
	*jazz.Jazz
	Models data.Models
}
