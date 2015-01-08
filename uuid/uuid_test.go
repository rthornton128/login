package uuid_test

import (
	"strconv"
	"testing"

	"github.com/rthornton128/login/uuid"
)

func TestValidUUID(t *testing.T) {
	for i := 0; i < 1000; i++ {
		id, err := uuid.NewVersion4()
		t.Log(id)
		if err != nil {
			t.Fatal(err)
		}

		b6, err := strconv.ParseUint(id[14:16], 16, 8)
		if err != nil {
			t.Fatal(err)
		}
		if byte(b6)&uuid.Version4 != uuid.Version4 {
			t.Fatal("Version4 bits not set properly")
		}
		b8, err := strconv.ParseUint(id[19:21], 16, 8)
		if err != nil {
			t.Fatal(err)
		}
		if byte(b8)&uuid.ReservedRFC4122 != uuid.ReservedRFC4122 {
			t.Fatal("RFC4122 bits not set properly")
		}
	}
}
