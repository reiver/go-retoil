package retoil


import (
	"time"
)


type internalDelayedLimitedStrategy struct{
	retoilCount uint
	maxRetoils  uint
	delay       time.Duration
}


// DelayedLimitedStrategy returns an initialized retoil.Strategizer which will retoil
// at most 'maxRetoils' times with at least a delay of 'delay' between each retoil.
func DelayedLimitedStrategy(maxRetoils uint, delay time.Duration) Strategizer {
	strategy := internalDelayedLimitedStrategy{
		retoilCount:0,
		maxRetoils:maxRetoils,
		delay:delay,
	}

	return &strategy
}

func (strategy *internalDelayedLimitedStrategy) MayPanickedRetoil(interface{}) bool {
	if strategy.retoilCount >= strategy.maxRetoils {
		return false
	}

	time.Sleep(strategy.delay)
	strategy.retoilCount++

	return true
}

func (strategy *internalDelayedLimitedStrategy) MayReturnedRetoil() bool {
	return false
}
