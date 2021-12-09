package net

import "context"

type Server interface {
	Listen(ctx context.Context)
	Shutdown() error
}
