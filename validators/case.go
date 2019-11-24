package validators

type Case struct {
	Legend string
	Status Status
}

func (v *Case) Success() {

	if v.Status != StatusUndefined {
		panic("Attempting to mark succeeded a case already judged")
	}
	v.Status = StatusSuccess

}

func (v *Case) Fail() {

	if v.Status != StatusUndefined {
		panic("Attempting to fail a case already judged")
	}
	v.Status = StatusFailed

}

func (v *Case) Skip() {

	if v.Status != StatusUndefined {
		panic("Attempting to skip a case already judged")
	}
	v.Status = StatusSkipped

}

func (v *Case) Undefined() bool {
	return v.Status == StatusUndefined
}
