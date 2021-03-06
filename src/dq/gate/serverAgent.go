package gate

import (
	"dq/datamsg"
	"dq/log"
	"dq/network"
	"dq/protobuf"
	"net"

	"github.com/golang/protobuf/proto"
)

//其他服务器连接上来的代理
type ServersAgent struct {
	conn     network.Conn
	gate     *Gate
	ModeType string
	handles  map[string]func(data *protomsg.MsgBase)

	gateHandles map[string]func(data *protomsg.MsgBase)
}

func (a *ServersAgent) GetConnectId() int32 {

	return 0
}
func (a *ServersAgent) GetModeType() string {
	return a.ModeType
}

//func (a *ServersAgent) registerDataHandle(msgtype string,f func(data *datamsg.MsgBase)) {

//	a.handles[msgtype] = f

//}

func (a *ServersAgent) Init() {
	a.handles = make(map[string]func(data *protomsg.MsgBase))
	a.handles[datamsg.GateMode] = a.DoGateData
	a.handles[datamsg.ClientMode] = a.DoClientData
	a.handles[datamsg.HallMode] = a.DoHallData
	a.handles[datamsg.LoginMode] = a.DoLoginMode
	//a.handles[datamsg.Game5GMode] = a.DoGame5GData
	a.handles[datamsg.GameScene1] = a.DoGameScene1

	a.gateHandles = make(map[string]func(data *protomsg.MsgBase))
	a.gateHandles["Register"] = a.DoGateRegisterData
	a.gateHandles["UserLogin"] = a.DoGateUserLoginData
	//重新登录 旧连接强制断开
	a.gateHandles["ReLoginForceDisconnect"] = a.ReLoginForceDisconnect
}

func (a *ServersAgent) DoGateRegisterData(data *protomsg.MsgBase) {
	h1 := data
	h2 := &protomsg.MsgRegisterToGate{}
	err := proto.Unmarshal(h1.Datas, h2)
	if err == nil {
		a.ModeType = h2.ModeType
		log.Info("--register modetype:" + h2.ModeType)
		a.gate.Models.Set(h2.ModeType, a)
	}
	//}
}

func (a *ServersAgent) ReLoginForceDisconnect(data *protomsg.MsgBase) {
	h1 := data
	if h1.Uid <= 0 {
		log.Info("--h1.Uid <= 0--")
		return
	}
	ag1 := (a.gate.GetAgent((h1.ConnectId)))
	if ag1 != nil {
		ag := ag1.(*agent)
		//异地登录 强制下线
		log.Info("--异地登录 强制下线--:%v", h1.Uid)
		ag.Close()
		return
	}

	log.Info("--not connect %d", (h1.ConnectId))
}

func (a *ServersAgent) DoGateUserLoginData(data *protomsg.MsgBase) {
	h1 := data
	if h1.Uid <= 0 {
		log.Info("--h1.Uid <= 0--")
		return
	}

	ag1 := (a.gate.GetAgent((h1.ConnectId)))
	if ag1 != nil {
		ag := ag1.(*agent)

		ag.UserData.Set("uid", h1.Uid)

		log.Info("--login--:%v", h1.Uid)

		return

	}
	log.Info("--not connect %d", (h1.ConnectId))

}

func (a *ServersAgent) DoGateData(data *protomsg.MsgBase) {
	h1 := data
	if f, ok := a.gateHandles[h1.MsgType]; ok {
		f(h1)
	}

}

func (a *ServersAgent) SendToAll(data1 []byte) {

	//	allAgents := a.gate.GetAgent()
	//	items := allAgents.Items()
	//	for _, v := range items {

	//		ag := v
	//		if ag != nil {
	//			ag.(*agent).WriteMsgBytes(data1)

	//		}
	//	}
}

func (a *ServersAgent) DoClientData(data *protomsg.MsgBase) {
	h1 := data
	connectid := (h1.ConnectId)
	uid := (h1.Uid)
	//向客户端隐藏connectid 和 uid
	h1.ConnectId = 0
	h1.Uid = 0
	data1, err1 := proto.Marshal(h1)
	if err1 == nil {

		//给所有玩家发消息
		if connectid == -2 && uid == -2 {
			go a.SendToAll(data1)
		} else {
			ag := a.gate.GetAgent(connectid)
			//			if ag == nil {

			//				con := a.gate.TcpServer.GetLoginedConnect().Get(uid)
			//				if con != nil {
			//					connectid = (con).(int32)
			//					ag = a.gate.TcpServer.GetAgents().Get(connectid)
			//				}

			//			}

			if ag != nil {
				ag.(*agent).WriteMsgBytes(data1)
				//log.Info("send:%s", data.JsonData)
			}
		}

	} else {
		log.Info("--err:%s", err1.Error())
	}
}

//
func (a *ServersAgent) DoLoginMode(data *protomsg.MsgBase) {
	h1 := data

	if model := a.gate.Models.Get(h1.ModeType); model != nil {
		data1, err1 := proto.Marshal(h1)
		if err1 == nil {
			model.(*ServersAgent).WriteMsgBytes(data1)
		}

	} else {
		log.Info("not find ModeType:%s", h1.ModeType)
	}

}

func (a *ServersAgent) DoHallData(data *protomsg.MsgBase) {
	h1 := data

	if model := a.gate.Models.Get(h1.ModeType); model != nil {
		data1, err1 := proto.Marshal(h1)
		if err1 == nil {
			model.(*ServersAgent).WriteMsgBytes(data1)
		}

	} else {
		log.Info("not find ModeType:%s", h1.ModeType)
	}

}

func (a *ServersAgent) DoGame5GData(data *protomsg.MsgBase) {

	h1 := data

	if model := a.gate.Models.Get(h1.ModeType); model != nil {
		data1, err1 := proto.Marshal(h1)
		if err1 == nil {
			model.(*ServersAgent).WriteMsgBytes(data1)
		}

	} else {
		log.Info("not find ModeType:%s", h1.ModeType)
	}

}

func (a *ServersAgent) DoGameScene1(data *protomsg.MsgBase) {

	h1 := data

	if model := a.gate.Models.Get(h1.ModeType); model != nil {
		data1, err1 := proto.Marshal(h1)
		if err1 == nil {
			model.(*ServersAgent).WriteMsgBytes(data1)
		}

	} else {
		log.Info("not find ModeType:%s", h1.ModeType)
	}

}

func (a *ServersAgent) Run() {
	a.Init()

	for {
		data, err := a.conn.ReadMsg()
		if err != nil {
			log.Debug("read message: %v", err)
			break
		}
		//log.Info("------readmsg:" + string(data))

		h1 := &protomsg.MsgBase{}
		err = proto.Unmarshal(data, h1)
		if err != nil {
			log.Info("--error")
		} else {
			if f, ok := a.handles[h1.ModeType]; ok {
				//go f(h1)
				f(h1)
			}

		}

	}
}

func (a *ServersAgent) OnClose() {

	log.Info("--delete modetype:%s", a.ModeType)
	a.gate.Models.Delete(a.ModeType)
}

func (a *ServersAgent) WriteMsg(msg interface{}) {

}
func (a *ServersAgent) WriteMsgBytes(msg []byte) {

	err := a.conn.WriteMsg(msg)

	if err != nil {
		log.Error("write message  error: %v", err)
	}
}

func (a *ServersAgent) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *ServersAgent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *ServersAgent) Close() {
	a.conn.Close()
}

func (a *ServersAgent) Destroy() {
	a.conn.Destroy()
}
