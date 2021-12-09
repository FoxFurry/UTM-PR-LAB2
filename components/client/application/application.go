package application

import (
	"context"
	"github.com/foxfurry/university/PRLab2/components/client/infrastructure/word"
	"github.com/foxfurry/university/PRLab2/components/client/internal/net"
	"time"
)

type Application interface {
	Start(ctx context.Context)
}

type serverApp struct {
}

func New() Application {
	return &serverApp{}
}

func (a *serverApp) Start(ctx context.Context) {
	clientTick := time.Tick(time.Second)

	for {
		select {
		case <-clientTick:
			net.UDPSend(word.Say())
		case <-ctx.Done():
			return
		}
	}
}
