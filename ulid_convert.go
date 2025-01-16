package ulidconv

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

// see https://asecuritysite.com/coding/ulid

func ULIDToUUID(u ulid.ULID) uuid.UUID {
	var uuidBytes [16]byte
	copy(uuidBytes[:], u[:16]) // Copy the first 16 bytes of ULID to the UUID
	return uuid.UUID(uuidBytes)
}

func UlidStrToUuidStr(ulidStr string) string {
	u, _ := ulid.Parse(ulidStr)
	uu := ULIDToUUID(u)
	return uu.String()
}

func UUIDToULID(u uuid.UUID) ulid.ULID {
	var ulidBytes [16]byte
	copy(ulidBytes[:], u[:16]) // Copy the UUID bytes into the ULID array
	return ulid.ULID(ulidBytes[:])
}

func UuidStrToUlidStr(uuidStr string) (string, error) {
	uu, err := uuid.Parse(uuidStr)
	if err != nil {
		return "", fmt.Errorf("uuidStr parse: %s", err.Error())
	}
	ul := UUIDToULID(uu)
	return ul.String(), nil
}
