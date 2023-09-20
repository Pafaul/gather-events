package logsource

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

type (
	LogBroadcaster interface {
		Subscribe() <-chan types.Log
		Unsubscribe(<-chan types.Log)
		StartBroadcasting(ctx context.Context)
	}

	logBroadcaster struct {
		source    <-chan types.Log
		listeners []chan types.Log
	}
)

func (b *logBroadcaster) Subscribe() <-chan types.Log {
	logSub := make(chan types.Log, 1)
	b.listeners = append(b.listeners, logSub)
	return logSub
}

func (b *logBroadcaster) Unsubscribe(chanToUnsubscribe <-chan types.Log) {
	for id, listener := range b.listeners {
		if listener == chanToUnsubscribe {
			b.listeners[id] = b.listeners[len(b.listeners)-1]
			b.listeners = b.listeners[:len(b.listeners)-1]
			close(listener)
		}
	}
}

func (b *logBroadcaster) serve(ctx context.Context) {
	defer (func() {
		for _, listener := range b.listeners {
			close(listener)
		}
	})()

	for {
		select {
		case <-ctx.Done():
			return
		case event := <-b.source:
			for id := range b.listeners {
				go func(l types.Log, listener chan types.Log) {
					listener <- event
				}(event, b.listeners[id])
			}
		}
	}
}

func (b *logBroadcaster) StartBroadcasting(ctx context.Context) {
	go b.serve(ctx)
}

func NewLogBroadcaster(logSource <-chan types.Log) LogBroadcaster {
	broadcaster := new(logBroadcaster)
	broadcaster.source = logSource
	broadcaster.listeners = make([]chan types.Log, 0)

	return broadcaster
}

var _ LogBroadcaster = (*logBroadcaster)(nil)
