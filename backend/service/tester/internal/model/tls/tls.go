package tls

import (
	"io"
	"strconv"
	"sync/atomic"

	"github.com/tinkler/mqttadmin/pkg/gs"
	"github.com/tinkler/mqttadmin/pkg/logger"
)

var (
	currentInt uint64
)

func getNextIntVal() string {
	i := atomic.AddUint64(&currentInt, 1)
	return strconv.FormatInt(int64(i), 10)
}

type Chat struct {
	Messages []string
}

func (m *Chat) InRoom(stream gs.NullStream) error {
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			logger.Error(err)
			return gs.ErrCanceled
		}
		m.Messages = append(m.Messages, getNextIntVal())
		err = stream.Send(nil)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			logger.Error(err)
			return gs.ErrCanceled
		}
	}
}
