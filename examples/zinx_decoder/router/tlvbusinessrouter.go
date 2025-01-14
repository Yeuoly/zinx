package router

import (
	"github.com/Yeuoly/zinx/zdecoder"
	"github.com/Yeuoly/zinx/ziface"
	"github.com/Yeuoly/zinx/zlog"
	"github.com/Yeuoly/zinx/znet"
)

type TLVBusinessRouter struct {
	znet.BaseRouter
}

func (this *TLVBusinessRouter) Handle(request ziface.IRequest) {
	zlog.Ins().DebugF("Call TLVRouter Handle %d %+v\n", request.GetMessage().GetMsgID(), request.GetMessage().GetData())
	msgID := request.GetMessage().GetMsgID()
	if msgID == 0x00000001 {
		_response := request.GetResponse()
		if _response != nil {
			switch _response.(type) {
			case zdecoder.TLVDecoder:
				tlvData := _response.(zdecoder.TLVDecoder)
				zlog.Ins().DebugF("do msgid=0x00000001 data business %+v\n", tlvData)
			}
		}
	}

}
