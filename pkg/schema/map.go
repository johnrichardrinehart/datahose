package schema

type mapSchema struct {
	fields map[string]fieldType
}

func (ms *mapSchema) Generate() map[string]interface{} {
	out := make(map[string]interface{}, len(ms.fields))

	for k, v := range ms.fields {
		out[k] = v.Generate()
	}

	return out
}
