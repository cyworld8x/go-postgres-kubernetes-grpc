package utils

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type PgGuid pgtype.UUID

// ConvertUUIDToPgTypeUUID converts a UUID to pgtype.UUID
func (guid PgGuid) ToString() string {
	s := fmt.Sprintf("%x-%x-%x-%x-%x", guid.Bytes[0:4], guid.Bytes[4:6], guid.Bytes[6:8], guid.Bytes[8:10], guid.Bytes[10:16])
	return s
}

// ConvertUUIDToPgTypeUUID converts a UUID to pgtype.UUID
func (guid PgGuid) ToUUID() uuid.UUID {
	s := guid.ToString()
	guild, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil
	}
	return guild
}

type Guid string

func (guid Guid) ToString() pgtype.UUID {
	pgtypeUUID := pgtype.UUID{}
	pgtypeUUID.Scan(guid)
	return pgtypeUUID
}
