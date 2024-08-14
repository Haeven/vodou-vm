// Code generated by "stringer -type=OpCode"; DO NOT EDIT.

package vm

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OpReturn-0]
	_ = x[OpConst-1]
	_ = x[OpAdd-2]
	_ = x[OpSub-3]
	_ = x[OpMul-4]
	_ = x[OpDiv-5]
	_ = x[OpNeg-6]
}

const _OpCode_name = "OpReturnOpConstOpAddOpSubOpMulOpDivOpNeg"

var _OpCode_index = [...]uint8{0, 8, 15, 20, 25, 30, 35, 40}

func (i OpCode) String() string {
	if i >= OpCode(len(_OpCode_index)-1) {
		return "OpCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OpCode_name[_OpCode_index[i]:_OpCode_index[i+1]]
}