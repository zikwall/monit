package service

import (
	"context"
	"fmt"
	"github.com/zikwall/monit/src/pkg/drop"
	"github.com/zikwall/monit/src/pkg/logger"
	"runtime"
	"strconv"
	"time"
)

type Impl struct {
	StorageClient *StorageClientImp
	Context       context.Context
	cancelRoot    context.CancelFunc
	drop          *drop.Droppable
}

type ImplOptions struct {
	StorageAddress string
}

func New(ctx context.Context, options *ImplOptions) (*Impl, error) {
	s := &Impl{drop: &drop.Droppable{}}
	s.Context, s.cancelRoot = context.WithCancel(ctx)

	storageClientImpl, err := newStorageClient(options.StorageAddress)
	if err != nil {
		return nil, err
	}

	s.StorageClient = storageClientImpl

	s.drop.AddDroppers(
		s.StorageClient,
	)

	return s, nil
}

// Shutdown method of implementing the cleaning of resources and connections when the application is shut down.
// Shutdown method addresses all connected listeners (which implement the corresponding interface)
// and calls the methods of this interface in turn.
// Shutdown accepts a single argument - a callback function with an argument of type error for custom error handling
// ```code
// 	Shutdown(func(err error) {
//		log.Warning(err)
//		bugsnag.Notify(err)
//	})
// ```
func (s *Impl) Shutdown(onError func(error)) {
	// cancel the root context and clear all allocated resources
	s.cancelRoot()
	s.drop.EachDroppers(func(dropper drop.Drop) {
		if impl, ok := dropper.(drop.DropDebug); ok {
			logger.Info(impl.DropMsg())
		}
		if err := dropper.Drop(); err != nil {
			onError(err)
		}
	})
}

// Stacktrace simple output of debugging information on the operation of the application
// and on the status of released resources (gorutin)
func (s *Impl) Stacktrace() {
	logger.Info("waiting for the server completion report to be generated")

	<-time.After(time.Second * 2)

	memory := runtime.MemStats{}
	runtime.ReadMemStats(&memory)

	colored := func(category, context string) string {
		return fmt.Sprintf("%s: %s", logger.Colored(category, logger.Cyan), logger.Colored(context, logger.Green))
	}

	fmt.Printf(
		"%s \n \t - %s \n \t - %s \n",
		logger.Colored("REPORT", logger.Green),
		colored("number of remaining goroutines:", strconv.Itoa(runtime.NumGoroutine())),
		colored("number of operations of the garbage collector:", strconv.Itoa(int(memory.NumGC))),
	)
}
