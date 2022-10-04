package mime_type

func ResolveMimeType(typ MimeType) MimeType {
	if resolved, ok := m[typ]; ok {
		return resolved
	}

	return typ
}
