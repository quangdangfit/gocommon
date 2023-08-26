package helpers

import (
	"context"
	"encoding/json"

	"golang.org/x/text/language"
	"google.golang.org/grpc/metadata"
)

// Copy value as json
func Copy(toValue interface{}, fromValue interface{}) error {
	data, err := json.Marshal(fromValue)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, toValue)
	if err != nil {
		return err
	}

	return nil
}

// get lang from context
func GetLangFromContext(ctx context.Context) string {
	l := language.Vietnamese.String()
	md, ok := metadata.FromIncomingContext(ctx)
	if ok && len(md["lang"]) > 0 {
		l = md["lang"][0]
	}
	return l
}
