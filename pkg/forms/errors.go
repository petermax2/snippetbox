package forms

type errors map[string][]string

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e errors) Get(field string) string {
	x := e[field]
	if len(x) == 0 {
		return ""
	}
	return x[0]
}
