package mapper

import (
	"github.com/payam1986128/go-fiber-sms-firewall/internal/entity"
	"github.com/payam1986128/go-fiber-sms-firewall/internal/presentation"
	"time"
)

func ToLimiterConditionResponse(limiterCondition *entity.LimiterCondition) presentation.LimiterConditionResponse {
	return presentation.LimiterConditionResponse{
		Name:       limiterCondition.Name,
		Priority:   limiterCondition.Priority,
		Active:     limiterCondition.Active,
		Action:     limiterCondition.Action,
		Filters:    toFiltersDto(limiterCondition.Filters),
		TimeLimits: toTimeLimitsDto(limiterCondition.TimeLimits),
	}
}

func ToBriefLimiterConditionDtos(limiterConditions []entity.LimiterCondition) []presentation.BriefLimiterConditionDto {
	if limiterConditions == nil {
		return nil
	}
	result := make([]presentation.BriefLimiterConditionDto, len(limiterConditions))
	for _, limiterCondition := range limiterConditions {
		result = append(result, presentation.BriefLimiterConditionDto{
			ID:          limiterCondition.ID.String(),
			Name:        limiterCondition.Name,
			Priority:    limiterCondition.Priority,
			CreatedTime: limiterCondition.CreatedTime.String(),
			TimeLimits:  toTimeLimitsDto(limiterCondition.TimeLimits),
		})
	}
	return result
}

func ToLimiterCondition(request presentation.LimiterConditionRequest, target *entity.LimiterCondition) {
	target.Name = *request.Name
	target.Priority = *request.Priority
	target.Action = *request.Action
	target.Filters = toFilters(*request.Filters)
	target.TimeLimits = toTimeLimits(request.TimeLimits)
	target.Active = true
}

func toFiltersDto(filters entity.Filters) presentation.FiltersDto {
	return presentation.FiltersDto{
		Keyword:   toKeywordFiltersDto(filters.Keyword),
		Sender:    toSendersFilterDto(filters.Sender),
		Receivers: filters.Receivers,
	}
}

func toFilters(dto presentation.FiltersDto) entity.Filters {
	return entity.Filters{
		Keyword:   toKeywordsFilter(dto.Keyword),
		Sender:    toSendersFilter(dto.Sender),
		Receivers: dto.Receivers,
	}
}

func toSendersFilterDto(sender *entity.SendersFilter) *presentation.SendersFilterDto {
	if sender == nil {
		return nil
	}
	return &presentation.SendersFilterDto{
		Senders: sender.Senders,
		Rate:    toRateFilterDto(sender.Rate),
	}
}

func toSendersFilter(dto *presentation.SendersFilterDto) *entity.SendersFilter {
	if dto == nil {
		return nil
	}
	return &entity.SendersFilter{
		Senders: dto.Senders,
		Rate:    toRateFilter(dto.Rate),
	}
}

func toRateFilterDto(filter *entity.RateFilter) *presentation.RateFilterDto {
	if filter == nil {
		return nil
	}
	return &presentation.RateFilterDto{
		IntervalMinutes: &filter.IntervalMinutes,
		Threshold:       &filter.Threshold,
	}
}

func toRateFilter(filter *presentation.RateFilterDto) *entity.RateFilter {
	if filter == nil {
		return nil
	}
	return &entity.RateFilter{
		IntervalMinutes: *filter.IntervalMinutes,
		Threshold:       *filter.Threshold,
	}
}

func toKeywordFiltersDto(filter *entity.KeywordsFilter) *presentation.KeywordsFilterDto {
	if filter == nil {
		return nil
	}
	return &presentation.KeywordsFilterDto{
		Keywords:   filter.Keywords,
		Categories: filter.Categories,
		Regexes:    filter.Regexes,
	}
}

func toKeywordsFilter(dto *presentation.KeywordsFilterDto) *entity.KeywordsFilter {
	if dto == nil {
		return nil
	}
	return &entity.KeywordsFilter{
		Keywords:   dto.Keywords,
		Categories: dto.Categories,
		Regexes:    dto.Regexes,
	}
}

func toTimeLimitsDto(limits []entity.TimeLimit) []presentation.TimeLimitDto {
	if limits == nil {
		return nil
	}
	var result []presentation.TimeLimitDto
	for _, limit := range limits {
		to := limit.To.String()
		limitDto := presentation.TimeLimitDto{
			To: &to,
		}
		from := limit.From.String()
		if limit.From != nil {
			limitDto.From = &from
		}
		result = append(result, limitDto)
	}
	return result
}

func toTimeLimits(dtos []presentation.TimeLimitDto) []entity.TimeLimit {
	if dtos == nil {
		return nil
	}
	layout := "1979-02-01 09:00:00"
	result := make([]entity.TimeLimit, len(dtos))
	for _, dto := range dtos {
		var limit = entity.TimeLimit{}
		from, _ := time.Parse(layout, *dto.From)
		limit.From = &from
		to, _ := time.Parse(layout, *dto.To)
		limit.To = to
		result = append(result, limit)
	}
	return result
}
