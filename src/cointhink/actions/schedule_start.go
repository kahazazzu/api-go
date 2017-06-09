package actions

import (
	"io/ioutil"
	"log"

	"cointhink/container"
	"cointhink/lxd"
	"cointhink/model/schedule"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
)

func DoScheduleStart(scheduleStart *proto.ScheduleStart, accountId string) []gproto.Message {
	var responses []gproto.Message

	log.Printf("ScheduleStart lookup %s %s", scheduleStart.ScheduleId, accountId)
	_schedule, err := schedule.Find(scheduleStart.ScheduleId, accountId)
	if err != nil {
		responses = append(responses, &proto.ScheduleStartResponse{Ok: false, Message: "unknown schedule id"})
	} else {
		log.Printf("schedule found %v", _schedule)
		if _schedule.AccountId != accountId {
			return []gproto.Message{&proto.ScheduleStopResponse{Ok: false, Message: "not owner"}}
		}

		schedule.UpdateStatus(_schedule, proto.Schedule_running)

		// push this somewhere else
		resp, err := lxd.Status(_schedule.Id)
		if err != nil {
			log.Print("LxdStatus: ", err)
			responses = append(responses, &proto.ScheduleStartResponse{Ok: false, Message: "unknown status"})
		}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		log.Printf("LxdStatus %v %v", resp.Status, string(bodyBytes))
		resp.Body.Close()
		if resp.StatusCode == 404 {
			err = container.Start(accountId, _schedule)
			if err != nil {
				responses = append(responses, &proto.ScheduleStartResponse{Ok: false, Message: err.Error()})
			}
		} else {
			responses = append(responses, &proto.ScheduleStartResponse{Ok: false, Message: "already running"})
			log.Printf("container not launched: exists")
		}
	}

	return responses
}
