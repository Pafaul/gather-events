package logsource

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
)

func TestBroadcaster(t *testing.T) {
	t.Log("Starting default broadcaster")

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	source := make(chan types.Log)

	broadcaster := NewLogBroadcaster(source)

	testBlocks := []types.Log{
		{},
		{},
	}

	receiver1 := broadcaster.Subscribe()
	receiver2 := broadcaster.Subscribe()

	broadcaster.StartBroadcasting(ctx)
	go (func() {
		for _, b := range testBlocks {
			source <- b
		}
	})()

	for range testBlocks {
		res1 := <-receiver1
		res2 := <-receiver2

		if res1.TxHash.Hex() != res2.TxHash.Hex() {
			t.Fatal("received values do not match")
		}
	}
}

func TestCancelContext(t *testing.T) {
	t.Log("Starting test with cancelled context")

	ctx, cancel := context.WithCancel(context.Background())

	source := make(chan types.Log)
	broadcaster := NewLogBroadcaster(source)

	receiver1 := broadcaster.Subscribe()

	broadcaster.StartBroadcasting(ctx)
	cancel()

	go (func() {
		source <- types.Log{}
	})()

	if len(receiver1) != 0 {
		t.Errorf("Broadcasting values after ctx cancel")
	}
}
