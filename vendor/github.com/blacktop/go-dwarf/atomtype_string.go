// Code generated by "stringer -type atomType -trimprefix=AtomType"; DO NOT EDIT.

package dwarf

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AtomTypeNULL-0]
	_ = x[AtomTypeDIEOffset-1]
	_ = x[AtomTypeCUOffset-2]
	_ = x[AtomTypeTag-3]
	_ = x[AtomTypeNameFlags-4]
	_ = x[AtomTypeTypeFlags-5]
	_ = x[AtomTypeQualNameHash-6]
}

const _atomType_name = "NULLDIEOffsetCUOffsetTagNameFlagsTypeFlagsQualNameHash"

var _atomType_index = [...]uint8{0, 4, 13, 21, 24, 33, 42, 54}

func (i atomType) String() string {
	if i >= atomType(len(_atomType_index)-1) {
		return "atomType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _atomType_name[_atomType_index[i]:_atomType_index[i+1]]
}
