package drop

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"golang.org/x/net/context"

	"github.com/zikwall/monit/src/pkg/logger"
)

type Impl struct {
	*Droppable
	context    context.Context
	cancelRoot context.CancelFunc
}

func NewContext(ctx context.Context) *Impl {
	i := &Impl{}
	i.context, i.cancelRoot = context.WithCancel(ctx)
	i.Droppable = &Droppable{}
	return i
}

func (s *Impl) Context() context.Context {
	return s.context
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
	s.cancelRoot()
	logger.Info("root context is canceled")
	s.EachDroppers(func(dropper Drop) {
		if impl, ok := dropper.(Debug); ok {
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
