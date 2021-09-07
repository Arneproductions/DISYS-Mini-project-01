package fork

type fork struct {
    numberOfTimesUsed int
    isFree bool
	in chan bool
	out chan bool
}