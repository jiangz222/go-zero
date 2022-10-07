package mon

import (
	"context"
	"io"
	"time"

	"github.com/jiangz222/go-zero/core/syncx"
	"go.mongodb.org/mongo-driver/mongo"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

const defaultTimeout = time.Second

var clientManager = syncx.NewResourceManager()

// ClosableClient wraps *mongo.Client and provides a Close method.
type ClosableClient struct {
	*mongo.Client
}

// Close disconnects the underlying *mongo.Client.
func (cs *ClosableClient) Close() error {
	return cs.Client.Disconnect(context.Background())
}

// Inject injects a *mongo.Client into the client manager.
// Typically, this is used to inject a *mongo.Client for test purpose.
func Inject(key string, client *mongo.Client) {
	clientManager.Inject(key, &ClosableClient{client})
}

func getClient(url string, clientOpts ...*mopt.ClientOptions) (*mongo.Client, error) {
	val, err := clientManager.GetResource(url, func() (io.Closer, error) {
		var cli *mongo.Client
		var err error
		if len(clientOpts) == 0 {
			cli, err = mongo.Connect(context.Background(), mopt.Client().ApplyURI(url))
		} else {
			// Don't find a good way to do both ApplyURI and mergeOptions
			// so only use what we need here like Registry
			cli, err = mongo.Connect(context.Background(), mopt.Client().ApplyURI(url).SetRegistry(clientOpts[0].Registry))
		}

		if err != nil {
			return nil, err
		}

		concurrentSess := &ClosableClient{
			Client: cli,
		}

		return concurrentSess, nil
	})
	if err != nil {
		return nil, err
	}

	return val.(*ClosableClient).Client, nil
}
