package task5kyu

func Rot13(msg string) string {
	var result []byte
	for _, b := range []byte(msg) {
		switch {
		case b >= 'A' && b <= 'Z':
			result = append(result, 'A'+(b-'A'+13)%26)
		case b >= 'a' && b <= 'z':
			result = append(result, 'a'+(b-'a'+13)%26)
		default:
			result = append(result, b)
		}
	}
	return string(result)
}
