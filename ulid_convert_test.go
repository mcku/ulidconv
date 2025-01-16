package ulidconv

import (
	"testing"
)

func TestUlidUuid(t *testing.T) {
	type pair struct {
		name    string
		ulidStr string
		uuidStr string
	}
	pairs := []pair{
		// TODO: Add test cases.
		{
			name:    "1",
			ulidStr: "7SFEM1BVJV835VF1XDGKT5GDW5",
			uuidStr: "f97ba815-ee5b-40cb-b787-ad84f4583785",
		},
		{
			name:    "2",
			ulidStr: "7ZZZZZZZZZZZZZZZZZZZZZZZZZ",
			uuidStr: "ffffffff-ffff-ffff-ffff-ffffffffffff",
		},
	}
	for _, tt := range pairs {
		if tt.uuidStr != UlidStrToUuidStr(tt.ulidStr) {
			t.Error(tt.name)
		}
		ulidStr, err := UuidStrToUlidStr(tt.uuidStr)
		if err != nil {
			t.Errorf("%s %s", tt.name, err.Error())
		}
		if tt.ulidStr != ulidStr {
			t.Errorf("%s expected: %s, actual: %s", tt.name, tt.ulidStr, ulidStr)
		}
	}
}
