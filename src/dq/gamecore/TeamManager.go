package gamecore

import (
	"dq/log"
	"dq/protobuf"
	"dq/utils"
	"sync"
)

var (
	TeamManagerObj  = &TeamManager{}
	TeamCount       = int32(100)
	TeamCountLock   = new(sync.RWMutex)
	OneTeamMaxCount = (5)
)

type ServerInterface interface {
	GetPlayerByID(id int32) *Player
}

func PopTeamID() int32 {
	TeamCountLock.Lock()
	defer TeamCountLock.Unlock()
	TeamCount++
	return TeamCount
}

type TeamInfo struct {
	ID          int32         //唯一id
	MainPlayer  *Player       //队长的uid
	Players     *utils.BeeMap //所有的成员
	OperateLock *sync.RWMutex //同步操作锁
	IsDismiss   bool          //是否已经解散
}

func (this *TeamInfo) Init() {
	this.Players = utils.NewBeeMap()
	this.OperateLock = new(sync.RWMutex)
}
func (this *TeamInfo) Dismiss() {
	this.OperateLock.Lock()
	defer this.OperateLock.Unlock()
	playeritem := this.Players.Items()
	for _, v := range playeritem {
		v.(*Player).TeamID = 0
	}
	this.Players.DeleteAll()
	this.IsDismiss = true
	log.Info("   解散队伍  ")
}

//删除成员
func (this *TeamInfo) RemoveMember(player *Player) {
	this.OperateLock.Lock()
	defer this.OperateLock.Unlock()
	if player == nil {
		return
	}

	player.TeamID = 0
	this.Players.Delete(player.Uid)
	log.Info("   成员离开队伍  ")
	//如果是队长
	if this.MainPlayer == player {
		this.MainPlayer = nil
		//随机选一个成员为队长
		players := this.Players.Items()
		for _, v := range players {
			this.MainPlayer = v.(*Player)
			break
		}
	}
}

//添加成员
func (this *TeamInfo) AddMember(player *Player) {

	this.OperateLock.Lock()
	defer this.OperateLock.Unlock()
	if player == nil {
		return
	}

	if player.TeamID > 0 {
		//已经在别的队伍里面了
		return
	}

	if this.IsDismiss == true {
		//已经解散了
		return
	}

	if this.Players.Size() >= OneTeamMaxCount {
		//满员了
		return
	}

	//成功入队
	player.TeamID = this.ID
	this.Players.Set(player.Uid, player)

	log.Info("   添加成员  ")

}

type TeamManager struct {
	Teams       *utils.BeeMap //当前服务器组队信息
	OperateLock *sync.RWMutex //同步操作锁
	Server      ServerInterface
}

//离开队伍
func (this *TeamManager) LeaveTeam(player *Player) {
	if player == nil || player.TeamID <= 0 {
		return
	}
	lastteam := this.Teams.Get(player.TeamID)
	if lastteam == nil {
		return
	}

	lastteam.(*TeamInfo).RemoveMember(player)

	if lastteam.(*TeamInfo).Players.Size() <= 1 {
		this.DismissTeam(lastteam.(*TeamInfo).ID, lastteam.(*TeamInfo).MainPlayer)
	}

}

//解散team
func (this *TeamManager) DismissTeam(teamid int32, player *Player) {

	this.OperateLock.Lock()
	defer this.OperateLock.Unlock()

	if player == nil {
		return
	}

	team := this.Teams.Get(teamid)
	if team == nil {
		return
	}

	//只有队长可以解散队伍
	if player != team.(*TeamInfo).MainPlayer {
		return
	}

	//
	team.(*TeamInfo).Dismiss()
	this.Teams.Delete(teamid)
}

//创建team
func (this *TeamManager) CreateTeam(player1 *Player, player2 *Player) {

	this.OperateLock.Lock()
	defer this.OperateLock.Unlock()

	if player1 == nil || player2 == nil || player1.TeamID > 0 || player2.TeamID > 0 {
		return
	}

	ti := &TeamInfo{}
	ti.Init()

	ti.ID = PopTeamID()
	ti.MainPlayer = player1
	ti.Players.Set(player1.Uid, player1)
	ti.Players.Set(player2.Uid, player2)
	player1.TeamID = ti.ID
	player2.TeamID = ti.ID
	this.Teams.Set(ti.ID, ti)

	log.Info("   创建队伍成功  ")

}

//组队回复
func (this *TeamManager) ResponseOrgTeam(data *protomsg.CS_ResponseOrgTeam, player *Player) {
	if data == nil || player == nil || this.Server == nil {
		return
	}
	srcplayer := this.Server.GetPlayerByID(data.SrcPlayerUID)
	if srcplayer == nil {
		return
	}

	//如果不同意
	if data.IsAgree != 1 {
		log.Info("   不同意  ")
		return
	}

	lastteam1 := this.Teams.Get(srcplayer.TeamID)
	lastteam2 := this.Teams.Get(player.TeamID)
	//都没有队伍就创建队伍
	if lastteam1 == nil && lastteam2 == nil {
		this.CreateTeam(srcplayer, player)
	}

	// 组队类型 1:请求加入战队 2:邀请目标加入战队
	if data.RequestType == 1 {
		if lastteam1 == nil && lastteam2 != nil {
			//
			lastteam2.(*TeamInfo).AddMember(srcplayer)
		}
	} else {
		if lastteam1 != nil && lastteam2 == nil {
			//
			lastteam1.(*TeamInfo).AddMember(player)
		}
	}

}

//组队 player1主动向player2发起组队请求
func (this *TeamManager) OrganizeTeam(player1 *Player, player2 *Player) {
	if player1 == nil || player2 == nil || player1.MainUnit == nil {
		return
	}
	lastteam1 := this.Teams.Get(player1.TeamID)
	lastteam2 := this.Teams.Get(player2.TeamID)

	msg := &protomsg.SC_RequestTeam{}
	msg.SrcPlayerUID = player1.Uid
	msg.SrcName = player1.MainUnit.Name
	msg.SrcUnitTypeID = player1.MainUnit.TypeID
	msg.SrcLevel = player1.MainUnit.Level

	if lastteam1 != nil && lastteam2 != nil {
		//2个玩家都有队伍 请先退出当前队伍再来组队

	} else if lastteam1 == nil && lastteam2 != nil {
		//向玩家2的队伍申请组队

		//// 组队类型 1:请求加入战队 2:邀请目标加入战队
		msg.RequestType = 1
		//		if lastteam2.(*TeamInfo).MainPlayer != nil {
		//			lastteam2.(*TeamInfo).MainPlayer.SendMsgToClient("SC_RequestTeam", msg)
		//		}
		player2.SendMsgToClient("SC_RequestTeam", msg)
	} else if lastteam1 != nil && lastteam2 == nil {
		//邀请玩家2进入队伍

		//// 组队类型 1:请求加入战队 2:邀请目标加入战队
		msg.RequestType = 2
		player2.SendMsgToClient("SC_RequestTeam", msg)
	} else {
		//玩家1向玩家2发出组队请求
		//this.CreateTeam(player1, player2)

		//// 组队类型 1:请求加入战队 2:邀请目标加入战队
		msg.RequestType = 2
		player2.SendMsgToClient("SC_RequestTeam", msg)
	}

}

//初始化
func (this *TeamManager) Init(server ServerInterface) {
	log.Info("----------TeamManager Init---------")
	this.Teams = utils.NewBeeMap()
	this.Server = server
	this.OperateLock = new(sync.RWMutex)

}

//更新
func (this *TeamManager) Update() {
	//当队伍只有1人的时候自动解散队伍

	//	teams := this.Teams.Items()
	//	for _, v := range teams {
	//		if v == nil {
	//			continue
	//		}
	//		if v.(*TeamInfo).Players.Size() <= 1 {
	//			this.DismissTeam(v.(*TeamInfo).ID, v.(*TeamInfo).MainPlayer)
	//		}
	//	}
}