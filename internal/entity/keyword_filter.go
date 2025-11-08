package entity

import "github.com/google/uuid"

type KeywordFilter struct {
	Keywords         []string
	Categories       []uuid.UUID
	CategoryKeywords []string
	Regexes          []string
}
