package strings

func (s Strings) ExtractTwoPointArgument(option string, field *string) (err string) {
	parts := s.Split(option, ":")
	if len(parts) >= 2 {
		*field = s.Join(parts[1:], ":")
		return ""
	}
	return "delimitador ':' no encontrado en la cadena " + option
}

// remover \cmd ej: "mi_directorio\cmd"  "\\cmd"
func (s Strings) RemoveSuffixIfPresent(path *string, suffix string) {
	if path != nil {
		if s.Contains(*path, suffix) != 0 {
			// Si el path tiene la extensión que queremos eliminar, la quitamos
			*path = (*path)[:len(*path)-len(suffix)]
		}
	}
}
