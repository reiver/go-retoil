package retoil


import (
	"math/rand"
	"time"
)


type internalLimitedExponentialBackoffStrategy struct{
	randomness    *rand.Rand
	retoilCount    uint
	maxRetoils     uint
	timeResolution time.Duration
}


// LimitedExponentialBackoffStrategy returns an initialized retoil.Strategizer which will only
// cause a retoil on a panic() (and not on a return) at most 'maxRetoils' times with
// using exponential backoff with time resolution 'timeResolution' to decide the delay between
// each retoil.
func LimitedExponentialBackoffStrategy(maxRetoils uint, timeResolution time.Duration) Strategizer {
	randomness := rand.New( rand.NewSource( time.Now().UTC().UnixNano() ) )

	strategy := internalLimitedExponentialBackoffStrategy{
		randomness:randomness,
		retoilCount:0,
		maxRetoils:maxRetoils,
		timeResolution:timeResolution,
	}

	return &strategy
}

func (strategy *internalLimitedExponentialBackoffStrategy) MayPanickedRetoil(interface{}) bool {
	if strategy.retoilCount >= strategy.maxRetoils {
		return false
	}

	delay := time.Duration( strategy.randomness.Intn(2 << strategy.retoilCount) )

	time.Sleep(delay)
	strategy.retoilCount++

	return true
}

func (strategy *internalLimitedExponentialBackoffStrategy) MayReturnedRetoil() bool {
	return false
}
