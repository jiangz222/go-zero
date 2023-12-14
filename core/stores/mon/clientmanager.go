package mon

import (
	"context"
	"io"

	"github.com/zeromicro/go-zero/core/syncx"
	"go.mongodb.org/mongo-driver/mongo"
	mopt "go.mongodb.org/mongo-driver/mongo/options"
)

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
		opt := mopt.Client().ApplyURI(url)
		if len(clientOpts) > 0 {
			// If options has conflict, opt from URL win.
			opt = mopt.MergeClientOptions(clientOpts[0], opt)
		}
		cli, err := mongo.Connect(context.Background(), opt)
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
