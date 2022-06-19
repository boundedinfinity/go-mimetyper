package mime_type

func ResolveMimeType(typ MimeType) MimeType {
	if ptyp, ok := m[typ]; ok {
		return ptyp
	}

	return typ
}
