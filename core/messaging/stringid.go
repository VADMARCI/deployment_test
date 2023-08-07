package messaging

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

type StringID int

func MarshalStringID(id StringID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(strconv.Itoa(int(id))))
	})
}

func UnmarshalStringID(v interface{}) (StringID, error) {
	if tmpStr, ok := v.(string); ok {
		if tmpInt, err := strconv.Atoi(tmpStr); err == nil {
			return StringID(tmpInt), nil
		}
	}
	return 0, fmt.Errorf("StringID must be a string formed of digits")
}
