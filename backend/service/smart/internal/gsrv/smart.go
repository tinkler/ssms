// Code generated by github.com/tinkler/mqttadmin; DO NOT EDIT.
package gsrv
import (
	"context"
	mrz "github.com/tinkler/mqttadmin/mrz/v1"
	"github.com/tinkler/ssms/backend/service/smart/internal/model/smart"
	pb_smart_v1 "github.com/tinkler/ssms/backend/service/smart/smart/v1"
	"github.com/tinkler/mqttadmin/pkg/jsonz/sjson"
)


type smartGsrv struct {
	pb_smart_v1.UnimplementedSmartGsrvServer
}

func NewSmartGsrv() *smartGsrv {
	return &smartGsrv{}
}

type SmartSmartHomeEnquireStream struct {
	stream pb_smart_v1.SmartGsrv_SmartHomeEnquireServer
	m      *smart.SmartHome
}
func (s *SmartSmartHomeEnquireStream) Context() context.Context {
	return s.stream.Context()
}
func (s *SmartSmartHomeEnquireStream) Send(v *smart.SmartChat) error {
	res := mrz.NewTypedRes[*pb_smart_v1.SmartHome, *pb_smart_v1.SmartChat]()
	// data
	res.Data = new(pb_smart_v1.SmartHome)
	jd, err := sjson.Marshal(s.m)
	if err != nil {
		return err
	}
	err = sjson.Unmarshal(jd, res.Data)
	if err != nil {
		return err
	}
	// resp
	respByt, _ := sjson.Marshal(v)
	newResp := &pb_smart_v1.SmartChat{}
	err = sjson.Unmarshal(respByt, newResp)
	if err != nil {
		return err
	}
	res.Resp = newResp
	
	return s.stream.Send(res.ToAny())
}
func (s *SmartSmartHomeEnquireStream) Recv() (*smart.SmartChat, error) {
	in, err := s.stream.Recv()
	if err != nil {
		return nil, err
	}
	req := mrz.ToTypedModel[*pb_smart_v1.SmartHome, *pb_smart_v1.SmartChat](in)
	jd, err := sjson.Marshal(req.Data)
	if err != nil {
		return nil, err
	}
	err = sjson.Unmarshal(jd, s.m)
	argsByt, _ := sjson.Marshal(req.Args)
	newArgs := &smart.SmartChat{}
	err = sjson.Unmarshal(argsByt, newArgs)
	return newArgs, err
	
}

func (s *smartGsrv) SmartHomeEnquire(stream pb_smart_v1.SmartGsrv_SmartHomeEnquireServer) error {
	gsStream := &SmartSmartHomeEnquireStream{stream, &smart.SmartHome{} }
	return gsStream.m.Enquire(gsStream)
}
