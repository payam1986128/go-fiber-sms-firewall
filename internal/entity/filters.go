package entity

type Filters struct {
	Keyword   KeywordFilter
	Sender    SenderFilter
	Receivers []string
}
