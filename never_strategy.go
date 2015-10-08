package retoil


type internalNeverStrategy struct{}


// NeverStrategy returns an initialized retoil.Strategizer which will never allow a retoil.
func NeverStrategy() Strategizer {
	strategy := internalNeverStrategy{}

	return &strategy
}


func (strategy *internalNeverStrategy) MayPanickedRetoil(interface{}) bool {
	return false
}

func (strategy *internalNeverStrategy) MayReturnedRetoil() bool {
	return false
}
