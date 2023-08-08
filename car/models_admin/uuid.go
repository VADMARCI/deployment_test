package models_admin

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

type UUID string

func MarshalUUID(uuid uuid.UUID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(uuid.String()))
	})
}

func UnmarshalUUID(v interface{}) (uuid.UUID, error) {
	str, ok := v.(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("UUID must be a string")
	}

	parsedUUID, err := uuid.Parse(str)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to parse UUID: %w", err)
	}

	return parsedUUID, nil
}

func (u UUID) MarshalGQL(w io.Writer) {
	q := fmt.Sprintf("%q", u)
	w.Write([]byte(q))
}

func (u *UUID) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("UUID must be a string")
	}

	parsedUUID, err := uuid.Parse(str)
	if err != nil {
		return fmt.Errorf("failed to parse UUID: %w", err)
	}

	*u = UUID(parsedUUID.String())
	return nil
}
