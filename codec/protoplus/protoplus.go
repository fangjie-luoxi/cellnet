package protoplus

import (
	"github.com/davyxu/protoplus/proto"
	"github.com/fangjie-luoxi/cellnet"
	"github.com/fangjie-luoxi/cellnet/codec"
)

type protoplus struct {
}

func (self *protoplus) Name() string {
	return "protoplus"
}

func (self *protoplus) MimeType() string {
	return "application/binary"
}

func (self *protoplus) Encode(msgObj interface{}, ctx cellnet.ContextSet) (data interface{}, err error) {

	return proto.Marshal(msgObj)

}

func (self *protoplus) Decode(data interface{}, msgObj interface{}) error {

	return proto.Unmarshal(data.([]byte), msgObj)
}

func init() {

	codec.RegisterCodec(new(protoplus))
}
