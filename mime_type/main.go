package mime_type

func Resolve(v MimeType) MimeType {
	s := m[v.String()]
	x, _ := Parse(s)
	return x
}

func ParseResolve(v string) (MimeType, error) {
	var o MimeType
	m, err := Parse(v)

	if err != nil {
		return o, err
	}

	return Resolve(m), nil
}
