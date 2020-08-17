package recover

import (
	"fmt"

	"github.com/gofiber/fiber"
)

// Config defines the config for middleware.
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	Next func(c *fiber.Ctx) bool
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next: nil,
}

// New creates a new middleware handler
func New(config ...Config) fiber.Handler {
	// Set default config
	cfg := ConfigDefault

	// Override config if provided
	if len(config) > 0 {
		cfg = config[0]

		if cfg.Next == nil {
			cfg.Next = ConfigDefault.Next
		}
	}

	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		// Catch panics
		defer func() {
			if r := recover(); r != nil {
				var ok bool
				if err, ok = r.(error); !ok {
					// Set error that will call the global error handler
					err = fmt.Errorf("%v", r)
				}
			}
		}()

		// Return err if exist, else move to next handler
		return c.Next()
	}
}