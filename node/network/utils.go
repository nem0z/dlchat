package network

func parseFixedSize(data []byte) string {
	for i, c := range data {
		if c == 0 {
			return string(data[:i])
		}
	}
	return string(data)
}
