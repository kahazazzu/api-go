package actions

import (
	"io/ioutil"
	"log"

	"cointhink/lxd"
	"cointhink/model/schedule"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
)

func DoScheduleStop(scheduleStop *proto.ScheduleStop, accountId string) []gproto.Message {
	var responses []gproto.Message

	log.Printf("ScheduleStop lookup %s %s", scheduleStop.ScheduleId, accountId)
	schedule, err := schedule.Find(scheduleStop.ScheduleId, accountId)
	if err != nil {
		responses = append(responses, &proto.ScheduleStopResponse{Ok: false, Message: "unknown schedule id"})
	} else {
		log.Printf("%v", schedule)
		if schedule.AccountId != accountId {
			return []gproto.Message{&proto.ScheduleStopResponse{Ok: false, Message: "not owner"}}
		}

		resp, err := lxd.Status(schedule.Id)
		if err != nil {
			log.Print("LxdStatus: ", err)
			responses = append(responses, &proto.ScheduleStartResponse{Ok: false, Message: "unknown status"})
		}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		log.Printf("LxdStatus %v %v", resp.Status, string(bodyBytes))
		resp.Body.Close()
	}

	return responses
}
