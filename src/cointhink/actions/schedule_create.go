package actions

import (
	"log"

	"cointhink/model/account"
	"cointhink/model/credit_journal"
	"cointhink/model/schedule"
	"cointhink/proto"

	gproto "github.com/golang/protobuf/proto"
)

func DoScheduleCreate(scheduleCreate *proto.ScheduleCreate, accountId string) []gproto.Message {
	var responses []gproto.Message

	_account, err := account.Find(accountId)
	if err != nil {
		responses = append(responses, &proto.ScheduleCreateResponse{Ok: false, Message: err.Error()})
	} else {
		if _account.ScheduleCredits > 0 {

			log.Printf("creating schedule for algorithm %s", scheduleCreate.Schedule.AlgorithmId)
			//	if algorithm.Owns(scheduleCreate.Schedule.AlgorithmId, accountId) {
			_schedule := proto.Schedule{AccountId: accountId,
				AlgorithmId:  scheduleCreate.Schedule.AlgorithmId,
				Status:       proto.Schedule_disabled,
				InitialState: scheduleCreate.Schedule.InitialState}
			log.Printf("inserting sched state %v", _schedule.Status)
			err = schedule.Insert(&_schedule)
			if err != nil {
				responses = append(responses, &proto.ScheduleCreateResponse{Ok: false, Message: err.Error()})
			} else {
				c_err := credit_journal.Debit(&_account, &_schedule, 1)
				if c_err != nil {
					log.Printf("DoScheduleCreate credit_journal Debit err %+v", c_err)
				}
				responses = append(responses, &proto.ScheduleCreateResponse{Ok: true})
			}
		} else {
			responses = append(responses, &proto.ScheduleCreateResponse{Ok: false, Message: "No remaining schedule credits."})
		}
	}
	return responses
}
