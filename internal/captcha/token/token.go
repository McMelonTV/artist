package token

import (
	"sync/atomic"
	"time"
)

var InitAmount = uint32(8)
var Limit = uint32(128)
var AvailableTokenCount uint32
var AvailableTokens = make(chan Token, Limit)
var QueuedTokenCount uint32
var NeededDuration = 5 * time.Second

type Token struct {
	Tk string
	Ex time.Time
}

func (t *Token) Valid() bool {
	return !t.IsExpired() && t.RemainingDuration() >= NeededDuration
}

func (t *Token) IsExpired() bool {
	return t.Ex.Before(time.Now())
}

func (t *Token) RemainingDuration() time.Duration {
	return t.Ex.Sub(time.Now())
}

var DefaultDuration = 5 * time.Minute

func CreateExpiration(issueTime time.Time) time.Time {
	return issueTime.Add(DefaultDuration)
}

func GetAvailable() uint32 {
	return atomic.LoadUint32(&AvailableTokenCount)
}

func GetQueued() uint32 {
	return atomic.LoadUint32(&QueuedTokenCount)
}

func GetTotal() uint32 {
	return GetAvailable() + GetQueued()
}

func IncrementAvailable() {
	atomic.AddUint32(&AvailableTokenCount, 1)
}
func DecrementAvailable() {
	atomic.AddUint32(&AvailableTokenCount, ^uint32(0))
}

func IncrementQueued() {
	atomic.AddUint32(&QueuedTokenCount, 1)
}
func DecrementQueued() {
	atomic.AddUint32(&QueuedTokenCount, ^uint32(0))
}
