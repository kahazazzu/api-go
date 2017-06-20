package lxd

import "cointhink/proto"
import "cointhink/q"
import "cointhink/httpclients"
import "cointhink/model/schedule"

import "cointhink/model/algorun"

import "log"

import "github.com/google/uuid"

var op_q []*q.AccountOperation

func AddOp(msg *q.AccountOperation) {
	log.Printf("lxd ADD Type %v Status %v Operation %v",
		msg.Operation.Type, msg.Operation.Status, msg.Operation.Operation)
	op_q = append(op_q, msg)
	WatchOp(msg)
}

func WatchOp(msg *q.AccountOperation) {
	op, err := lxdCallOperation("GET", msg.Operation.Operation+"/wait")
	if err != nil {
		log.Printf("lxd.WatchOp err: %v", err)
	}
	log.Printf("lxd.WatchOp finished %s %s", msg.Algorun.Id, op.Status)

	if op.Status == "error" {
		log.Printf("WatchOp got error, skipping status check")
	} else {
		algoRun, err := algorun.Find(msg.Algorun.Id)
		lxdStatus, err := Status(msg.Algorun.Id)
		log.Printf("lxd.WatchOp lxd status: id:%s status:%v err:%v", msg.Algorun.Id,
			lxdStatus.Metadata.Status, err)

		var algorun_state proto.Algorun_States
		if lxdStatus.ErrorCode == 404 {
			algorun_state = proto.Algorun_deleted
		} else {
			if algoRun.Status == proto.Algorun_States_name[int32(proto.Algorun_building)] &&
				lxdStatus.Metadata.Status == "Stopped" {
				algorun_state = proto.Algorun_stopped
			}
		}
		algorun.UpdateStatus(algoRun, algorun_state)

		// alert client
		s, _ := schedule.Find(msg.Algorun.ScheduleId)
		sr := proto.ScheduleRun{Schedule: &s, Run: algoRun}
		if algoRun.Status == proto.Algorun_States_name[int32(proto.Algorun_deleted)] {
			sr.Run = nil
		}
		g := proto.ScheduleListPartial{ScheduleRun: &sr}

		socket := httpclients.AccountIdToSocket(msg.Algorun.AccountId)
		if socket == nil {
			log.Printf("Watchop client socket lookup fail #s", msg.Algorun.AccountId)
		} else {
			q.OUTq <- q.RpcOut{Socket: socket,
				Response: &q.RpcResponse{Msg: &g, Id: RpcId()}}
		}

	}
}

func RpcId() string {
	uuid, _ := uuid.NewRandom()
	uuidStr := uuid.String()
	return uuidStr[19:35]
}