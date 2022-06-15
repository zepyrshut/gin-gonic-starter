package routes

import (
	"testing"

	"github.com/zepyrshut/gin-gonic-starter/internal/config"
)

func TestNewRoutes(t *testing.T) {
	a := &config.Application{}
	NewRoutes(a)

	if app == nil {
		t.Error("app should not be nil")
	}
}
