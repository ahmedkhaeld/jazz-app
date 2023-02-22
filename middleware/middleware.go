package middleware

import (
	"github.com/ahmedkhaeld/jazz"
	"myapp/data"
)

type Middleware struct {
	*jazz.Jazz
	Models data.Models
}
