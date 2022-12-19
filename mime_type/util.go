package mime_type

func ResolveMimeType(mt MimeType) MimeType {
	if resolved, ok := m[mt]; ok {
		return resolved
	}

	return mt
}
