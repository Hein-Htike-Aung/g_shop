package util

import (
	"errors"
	"sync"
	"time"
)

// snowflake: distributed unique IDs include cluster and node
const (
	workerBits uint8 = 10 // worker ID bits (max 1024 nodes)
	numberBits uint8 = 12 // sequence bits per ms (4096 IDs/ms max)
	// max value via bit ops，-1 two's complement; verify:  -1 ^ (-1 << nodeBits) equals 1023
	workerMax   int64 = -1 ^ (-1 << workerBits) // max worker ID
	numberMax   int64 = -1 ^ (-1 << numberBits) // max sequence number
	timeShift         = workerBits + numberBits // timestamp shift
	workerShift       = numberBits              // worker ID shift
	// ~68 years with 41-bit timestamp
	// subtract custom epoch to avoid wasting timestamp range
	// do not change epoch after IDs are generated
	epoch int64 = 1525705533000 // epoch ms when this was written
)

// snowflake worker node state
type Worker struct {
	mu        sync.Mutex // mutex for concurrency
	timestamp int64      // last timestamp
	workerId  int64      // worker ID
	number    int64      // sequence in current ms
}

// create worker node
func NewWorker(workerId int64) (*Worker, error) {
	// validate workerId range
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID excess of quantity")
	}
	// allocate worker
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

// generate ID
// Generate must be called on a worker
func (w *Worker) GetId() int64 {
	// lock before generating ID
	w.mu.Lock()
	defer w.mu.Unlock() // unlock after generate

	// current timestamp ms
	now := time.Now().UnixNano() / 1e6 // nanoseconds to ms
	if w.timestamp == now {
		w.number++

		// check sequence limit per ms
		if w.number > numberMax {
			// wait 1ms if sequence exhausted
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		// reset sequence on new millisecond
		w.number = 0
		w.timestamp = now // update last timestamp
	}

	// timestamp component: now - epoch
	// changing epoch later may duplicate IDs
	ID := (now-epoch)<<timeShift | (w.workerId << workerShift) | (w.number)
	return ID
}