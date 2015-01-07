// Package uuid implements the barest minimum requirements for a Version 4 UUID
// as specified by RFC4122 found at: http://tools.ietf.org/html/rfc4122
package uuid

import (
	"crypto/rand"
	"fmt"
)

const (
	ReservedRFC4122 byte = 0x80
	ReservedMask    byte = 0x1f
	Version4        byte = 0x40
	VersionMask     byte = 0x0f
)

// NewVersion4 returns a string containing a new, randomly generated UUID
func NewVersion4() string {
	b := make([]byte, 16)
	rand.Read(b)
	b[6] = b[6]&VersionMask | Version4
	b[8] = b[8]&ReservedMask | ReservedRFC4122
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
