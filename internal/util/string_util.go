package util

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func JoinQuotedUUIDs(ids []uuid.UUID, separator string) string {
	str := make([]string, len(ids))
	for i, u := range ids {
		str[i] = fmt.Sprintf("'%s'", u.String())
	}
	return strings.Join(str, separator)
}
