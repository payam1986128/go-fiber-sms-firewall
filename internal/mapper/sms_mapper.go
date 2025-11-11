package mapper

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
)

func toBriefSmsDto(sms entity.Sms) presentation.BriefSmsDto {
	return presentation.BriefSmsDto{
		ID:              sms.ID.String(),
		Sender:          sms.Sender,
		Receiver:        sms.Receiver,
		Message:         sms.Message,
		Action:          sms.Action,
		ReceivedTime:    sms.ReceivedTime,
		EvaluatedTime:   sms.EvaluatedTime,
		AppliedFilterID: sms.AppliedFilterId.String(),
		SendTime:        sms.SendTime,
	}
}

func ToBriefSmsDtos(sms []entity.Sms) []presentation.BriefSmsDto {
	if sms == nil {
		return nil
	}
	result := make([]presentation.BriefSmsDto, len(sms))
	for _, s := range sms {
		result = append(result, toBriefSmsDto(s))
	}
	return result
}
