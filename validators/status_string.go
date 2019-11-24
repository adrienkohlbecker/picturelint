// Code generated by "stringer -type=Status"; DO NOT EDIT.

package validators

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StatusUndefined-0]
	_ = x[StatusSuccess-1]
	_ = x[StatusFailed-2]
	_ = x[StatusSkipped-3]
}

const _Status_name = "StatusUndefinedStatusSuccessStatusFailedStatusSkipped"

var _Status_index = [...]uint8{0, 15, 28, 40, 53}

func (i Status) String() string {
	if i < 0 || i >= Status(len(_Status_index)-1) {
		return "Status(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}
