package rid

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"hash/fnv"
	"os"
)

// Salt calculated to generate an uint64 id from machine id
func Salt() uint64 {
	// Use FNV-a1 algo to calculate hash value
	hasher := fnv.New64a()
	hasher.Write(ReadMachineID())

	// Convert hash value to uint64-salt
	hashValue := hasher.Sum64()
	return hashValue
}

// ReadMachineID to fetch machine ID if no result then generate a random ID
func ReadMachineID() []byte {
	id := make([]byte, 3)
	machineID, err := readPlatformMachineID()
	if err != nil || len(machineID) == 0 {
		machineID, err = os.Hostname()
	}

	if err == nil || len(machineID) == 0 {
		hasher := sha256.New()
		hasher.Write([]byte(machineID))
		copy(id, hasher.Sum(nil))
	} else {
		// if the machine id cannot be found, then generate a random number
		if _, randErr := rand.Read(id); randErr != nil {
			panic(fmt.Errorf("id: cannot get hostname nor generate a random number: %w; %w", err, randErr))
		}
	}
	return id
}

func readPlatformMachineID() (string, error) {
	data, err := os.ReadFile("/etc/machine-id")
	if err != nil || len(data) == 0 {
		data, err = os.ReadFile("/sys/class/dmi/id/product_uuid")
	}
	return string(data), err
}
