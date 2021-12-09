package application

import (
	"context"
	"github.com/foxfurry/university/PRLab2/components/server/internal/net"
)

type Application interface {
	Start(ctx context.Context)
	Shutdown()
}

type serverApp struct {
	servers []net.Server
}

func New() Application {
	return &serverApp{
		servers: []net.Server{
			net.NewUDP(),
		},
	}
}

func (a *serverApp) Start(ctx context.Context) {
	for idx, _ := range a.servers {
		go a.servers[idx].Listen(ctx)
	}
}

func (a *serverApp) Shutdown() {
	for idx, _ := range a.servers {
		a.servers[idx].Shutdown()
	}
}
