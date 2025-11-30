package mapper

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
)

func ToSuspiciousWordDtos(source []entity.SuspiciousWord) []presentation.SuspiciousWordDto {
	if source == nil {
		return nil
	}
	target := make([]presentation.SuspiciousWordDto, len(source))
	for _, suspiciousWord := range source {
		target = append(target, toSuspiciousWordDto(suspiciousWord))
	}
	return target
}

func toSuspiciousWordDto(source entity.SuspiciousWord) presentation.SuspiciousWordDto {
	return presentation.SuspiciousWordDto{
		ID:       source.ID.String(),
		DateTime: source.DateTime.String(),
		Word:     source.Word,
	}
}
