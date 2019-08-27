// Package schema implements a parser and generator of a user-provided schema.
// A big shoutout to Patrick Stephen - @200sc - https://github.com/200sc - for his help in providing the sample code that served as the basis of this package.
// https://play.golang.org/p/2fPPlBPiYux
package schema

import "encoding/json"

// Schema interface defines the behavior of types which can generate a valid instance of a schema
type Schema interface {
	Generate() map[string]interface{}
}

func Parse(schema string) (Schema, error) {
	mp := map[string]string{}

	err := json.Unmarshal([]byte(schema), &mp)
	if err != nil {
		return nil, err
	}

	sch := &mapSchema{
		fields: make(map[string]fieldType, len(mp)),
	}

	for k, v := range mp {
		sch.fields[k] = fieldType(v)
		if _, ok := knownFields[sch.fields[k]]; !ok {
			return nil, UnknownField{}
		}
	}

	return sch, nil
}
