package mapper

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"time"
)

func ToSuspiciousCategoryDtos(source []entity.SuspiciousCategory) []presentation.SuspiciousCategoryDto {
	if source == nil {
		return nil
	}
	target := make([]presentation.SuspiciousCategoryDto, len(source))
	layout := "1979-02-01 09:00:00"
	for _, suspiciousCategory := range source {
		dto := presentation.SuspiciousCategoryDto{
			ID:    suspiciousCategory.ID.String(),
			Name:  suspiciousCategory.Name,
			Words: suspiciousCategory.Words,
		}
		dateTime, _ := time.Parse(layout, suspiciousCategory.DateTime)
		dto.DateTime = dateTime
		target = append(target, dto)
	}
	return target
}

func ToSuspiciousCategory(source presentation.SuspiciousCategoryWordsRequest) entity.SuspiciousCategory {
	return entity.SuspiciousCategory{
		Name:  source.Name,
		Words: source.Words,
	}
}
