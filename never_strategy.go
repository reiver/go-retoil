package retoil


type internalNeverStrategy struct{}


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
