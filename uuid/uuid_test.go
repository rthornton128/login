package uuid_test

import (
	"strconv"
	"testing"

	"github.com/rthornton128/login/uuid"
)

func TestValidUUID(t *testing.T) {
	for i := 0; i < 1000; i++ {
		id := uuid.NewVersion4()
		t.Log(id)

		b6, err := strconv.ParseUint(id[14:16], 16, 8)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		if byte(b6)&uuid.Version4 != uuid.Version4 {
			t.Log("Version4 bits not set properly")
			t.FailNow()
		}
		b8, err := strconv.ParseUint(id[19:21], 16, 8)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		if byte(b8)&uuid.ReservedRFC4122 != uuid.ReservedRFC4122 {
			t.Log("RFC4122 bits not set properly")
			t.FailNow()
		}
	}
}
