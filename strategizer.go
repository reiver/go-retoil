package retoil


type Strategizer interface {

	// MayPanickedRetoil will be asked by a retoiler if it may retoil after a the toiler panic()ed.
	MayPanickedRetoil(interface{}) bool

	// MayReturnedRetoil will be asked by a retoiler if it may retoil after a the toiler returned (gracefully).
	MayReturnedRetoil() bool
}
