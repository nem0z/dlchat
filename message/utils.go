package message

import "crypto/sha256"

func toFixLength(buf []byte, length int) []byte {
	if len(buf) >= length {
		return nil
	}

	toAdd := make([]byte, length-len(buf))
	return append(buf, toAdd...)
}

func checksum(data []byte) []byte {
	hash := sha256.Sum256(data)
	hash = sha256.Sum256(hash[:])
	return hash[:4]
}
