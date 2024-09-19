package id

import (
	"fmt"
	"sync"
	"time"
)

const (
	timestampBits   = 41
	datacenterBits  = 5
	machineBits     = 5
	sequenceBits    = 12
	maxDatacenterID = (1 << datacenterBits) - 1
	maxMachineID    = (1 << machineBits) - 1
	maxSequence     = (1 << sequenceBits) - 1
	epoch           = int64(1288834974657)
	timestampShift  = sequenceBits + machineBits + datacenterBits
	datacenterShift = sequenceBits + machineBits
	machineShift    = sequenceBits
)

type Snowflake struct {
	sync.Mutex
	timestamp    int64
	datacenterID int64
	machineID    int64
	sequence     int64
}

func NewSnowflake(datacenterID, machineID int64) (*Snowflake, error) {
	if datacenterID > maxDatacenterID || datacenterID < 0 {
		return nil, fmt.Errorf("datacenter ID must be between 0 and %d", maxDatacenterID)
	}

	if machineID > maxMachineID || machineID < 0 {
		return nil, fmt.Errorf("machine ID must be between 0 and %d", maxMachineID)
	}

	return &Snowflake{
		timestamp:    0,
		datacenterID: datacenterID,
		machineID:    machineID,
		sequence:     0,
	}, nil
}

func (snowflake *Snowflake) GenerateID() int64 {
	snowflake.Lock()
	defer snowflake.Unlock()
	now := time.Now().UnixMilli()
	if now == snowflake.timestamp {
		snowflake.sequence = (snowflake.sequence + 1) & maxSequence
		if snowflake.sequence == 0 {
			for now <= snowflake.timestamp {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		snowflake.sequence = 0
	}

	snowflake.timestamp = now
	id := ((now - epoch) << timestampShift) |
		(snowflake.datacenterID << datacenterShift) |
		(snowflake.machineID << machineShift) |
		snowflake.sequence

	return id
}
