package smart

import (
	"context"
	"errors"

	"github.com/tinkler/mqttadmin/pkg/gs"
	"github.com/tinkler/mqttadmin/pkg/logger"
	"github.com/tinkler/mqttadmin/pkg/status"
)

type SmartHome struct {
	CurrentModel string
}

type Model string

const (
	_                     Model = "None"
	ModelListCommodity    Model = "ListCommodities"
	ModelListPurchseOrder Model = "ListPurchaseOrders"
)

func (sh *SmartHome) SetModel(model Model) {
	sh.CurrentModel = string(model)
}

func (sh *SmartHome) Enquire(stream gs.Stream[*SmartChat]) error {
	// Recv first time
	// Ignore the messages
	_, err := stream.Recv()
	if errors.Is(status.ErrCanceled, err) {
		return nil
	}
	if err != nil {
		logger.Error(err)
		return err
	}
	chatCtx, chatCancel := context.WithCancel(stream.Context())
	go func() {
		for {
			sc, err := stream.Recv()
			if errors.Is(status.ErrCanceled, err) {
				chatCancel()
				return
			}
			if err != nil {
				logger.Error(err)
				chatCancel()
				return
			}
			if len(sc.Messages)%2 == 0 {
				sh.SetModel(ModelListPurchseOrder)
			} else {
				sh.SetModel("")
			}
			err = stream.Send(&SmartChat{
				Messages: sc.Messages,
			})

			if errors.Is(status.ErrCanceled, err) {
				chatCancel()
				return
			}
			if err != nil {
				logger.Error(err)
				chatCancel()
				return
			}
		}
	}()
	<-chatCtx.Done()
	return nil
}
