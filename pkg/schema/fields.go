package schema

import "math/rand"

type fieldType string

func (ft fieldType) Generate() interface{} {
	switch ft {
	case intField:
		return rand.Int()
	case f32Field:
		return float32(rand.NormFloat64())
	case strField:
		// Bit more work
		const lenMax = 100
		ln := rand.Intn(lenMax)
		out := make([]byte, ln)
		for i := range out {
			// readable ascii
			out[i] = byte(rand.Intn(96) + 32)
		}
		return string(out)
	case f64Field:
		return rand.NormFloat64()
	case boolField:
		if rand.Float64() < .5 {
			return true
		}
		return false
	}
	return nil
}

const (
	intField  fieldType = "int"
	f32Field  fieldType = "float32"
	f64Field  fieldType = "float64"
	strField  fieldType = "string"
	boolField fieldType = "bool"
)

var knownFields = map[fieldType]struct{}{
	intField:  struct{}{},
	f32Field:  struct{}{},
	f64Field:  struct{}{},
	strField:  struct{}{},
	boolField: struct{}{},
}

type UnknownField struct{}

func (UnknownField) Error() string {
	return "Couldn't parse field type"
}
