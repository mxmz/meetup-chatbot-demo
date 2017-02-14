package tskbroker

import (
	"log"
	"sync"
	"time"

	"golang.org/x/net/context"
)

type MqMessage struct {
	ID      string
	Type    string
	Delayed time.Time
	Payload interface{}
}

type MqInMessage struct {
	Msg  *MqMessage
	ErrC chan<- error
	Ctx  context.Context
}
type MqOutMessage struct {
	Msg  *MqMessage
	ErrC chan<- error
	Ctx  context.Context
}

type MqOutChannel chan MqOutMessage

type MqInChannel chan MqInMessage

type MqBroker struct {
	inC         MqInChannel
	outC        MqOutChannel
	cancelCtx   context.Context
	cancelFunc  func()
	wg          sync.WaitGroup
	concurrency int
	bufferSize  int
}

func NewMqBroker(ctx context.Context, concurrency int, bufferSize int) *MqBroker {
	ctx, f := context.WithCancel(ctx)
	return &MqBroker{
		inC:         make(MqInChannel),
		outC:        make(MqOutChannel),
		cancelCtx:   ctx,
		cancelFunc:  f,
		concurrency: concurrency,
		bufferSize:  bufferSize,
	}
}

func (b *MqBroker) InC() MqInChannel {
	return b.inC
}

func (b *MqBroker) Cancel() context.Context {
	return b.cancelCtx
}

func (b *MqBroker) OutC() MqOutChannel {
	return b.outC
}
func (b *MqBroker) Shutdown() {
	b.cancelFunc()
	b.wg.Wait()
}

func (b *MqBroker) Start() {
	var bufferC = make(chan *MqMessage, b.bufferSize)
	var delayedC = make(chan *MqMessage, b.bufferSize)
	done := b.cancelCtx.Done()
	delayer := func() {
		defer b.wg.Done()
		defer log.Println("quit delayer")
		jobMap := map[int64][]*MqMessage{}
		next := time.Now()
		for {
			//			log.Println(next)
			select {
			case <-done:
				{
					return
				}
			case im := <-delayedC:
				{
					log.Println("delaying msg ", im.Delayed)
					jobs := jobMap[im.Delayed.Unix()]
					if jobs == nil {
						jobs = make([]*MqMessage, 0)
					}
					jobs = append(jobs, im)
					jobMap[im.Delayed.Unix()] = jobs
				}

			case <-time.After(next.Sub(time.Now())):
				{
					next = next.Add(10 * time.Second)
					for k, v := range jobMap {
						if k <= time.Now().Unix() {
							jobs := v
							log.Println("retry", jobs)
							go func() {
								for _, v := range jobs {
									select {
									case bufferC <- v:
									case <-done:
									}
								}
							}()
							delete(jobMap, k)
						} else {
							if k < next.Unix() {
								next = time.Unix(k, 0)
							}
						}
					}

				}
			}

		}
	}

	reader := func() {
		defer b.wg.Done()
		defer log.Println("quit reader")
		for {
			select {
			case <-done:
				{
					return
				}
			case im := <-b.inC:
				{
					routeC := bufferC
					if im.Msg.Delayed.After(time.Now()) {
						routeC = delayedC
					}

					select {
					case im.ErrC <- nil:
					case <-done:
						{
							return
						}
					}
					select {
					case routeC <- im.Msg:
					case <-done:
						{
							return
						}
					}

				}
			}
		}
	}

	writer := func() {
		defer b.wg.Done()
		defer log.Println("quit writer")
		errC := make(chan error, 1)
		for {
			select {
			case <-done:
				{
					return
				}
			case m, ok := <-bufferC:
				{
					if !ok {
						return
					}
					om := MqOutMessage{m, errC, b.cancelCtx}
					select {
					case b.outC <- om:
					case <-done:
						{
							return
						}
					}
					select {
					case err := <-errC:
						{
							if err != nil {
								log.Println(err)
								go func() {
									om.Msg.Delayed =
										om.Msg.Delayed.Add(20 * time.Second)
									log.Println("delaying", om)
									select {
									case delayedC <- om.Msg:
									case <-done:
									}

								}()
							}
						}
					case <-done:
						{
							return
						}
					}
				}
			}
		}
	}
	for i := 0; i < b.concurrency; i++ {
		b.wg.Add(3)
		go reader()
		go writer()
		go delayer()
	}
}

type MqBrokerMap struct {
	mtx         sync.Mutex
	brokers     map[string]*MqBroker
	cancelCtx   context.Context
	concurrency int
	bufferSize  int
}

func NewMqBrokerMap(ctx context.Context, concurrency int, bufferSize int) *MqBrokerMap {
	return &MqBrokerMap{
		brokers:     make(map[string]*MqBroker),
		cancelCtx:   ctx,
		concurrency: concurrency,
		bufferSize:  bufferSize,
	}
}

func (bs *MqBrokerMap) GetBrokerByName(name string) *MqBroker {
	bs.mtx.Lock()
	defer bs.mtx.Unlock()
	b := bs.brokers[name]
	if b == nil {
		b = NewMqBroker(bs.cancelCtx, bs.concurrency, bs.bufferSize)
		bs.brokers[name] = b
		b.Start()
	}
	return b
}
func (bs *MqBrokerMap) Shutdown() {
	bs.mtx.Lock()
	brokers := make([]*MqBroker, 0)
	for _, v := range bs.brokers {
		brokers = append(brokers, v)
	}
	bs.mtx.Unlock()
	for _, b := range brokers {
		b.Shutdown()
	}
}

func (bs *MqBrokerMap) Inject(ctx context.Context, queue string, m *MqMessage) error {
	b := bs.GetBrokerByName(queue)
	errC := make(chan error, 1)
	select {
	case b.InC() <- MqInMessage{m, errC, ctx}:
		{
			select {
			case err := <-errC:
				{
					return err
				}
			case <-ctx.Done():
				{
					return ctx.Err()
				}
			}

		}
	case <-b.Cancel().Done():
		{
			return b.Cancel().Err()

		}
	case <-ctx.Done():
		{
			return ctx.Err()
		}
	}
	return nil
}

func (bs *MqBrokerMap) Listen(ctx context.Context, f func(m *MqMessage) error, queues ...string) {
	read := make(chan MqOutMessage)
	ctx, cancel := context.WithCancel(ctx)
	brokers := make([]*MqBroker, 0)
	for _, n := range queues {
		brokers = append(brokers, bs.GetBrokerByName(n))
	}
	for _, b := range brokers {
		go func(b *MqBroker) {
			defer log.Println("quit listener")
			defer cancel()
			done := b.Cancel().Done()
			for {
				select {
				case <-done:
					{
						return
					}
				case m := <-b.OutC():
					{
						select {
						case <-done:
							{
								return
							}
						case read <- m:
						}
					}
				}

			}
		}(b)
	}
	for {
		select {
		case <-ctx.Done():
			{
				return
			}

		case m := <-read:
			{
				err := f(m.Msg)
				select {
				case <-ctx.Done():
					{
						return
					}
				case <-m.Ctx.Done():
					{
						return
					}
				case m.ErrC <- err:
				}

			}
		}
	}

}
