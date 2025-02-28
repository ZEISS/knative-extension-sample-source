// Code generated by "stringer -type ContextDescriptorKind,TypeReferenceKind,MetadataInitializationKind,SpecialKind,GenericParamKind,GenericRequirementLayoutKind -linecomment -output types_string.go"; DO NOT EDIT.

package swift

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CDKindModule-0]
	_ = x[CDKindExtension-1]
	_ = x[CDKindAnonymous-2]
	_ = x[CDKindProtocol-3]
	_ = x[CDKindOpaqueType-4]
	_ = x[CDKindClass-16]
	_ = x[CDKindStruct-17]
	_ = x[CDKindEnum-18]
}

const (
	_ContextDescriptorKind_name_0 = "moduleextensionanonymousprotocolopaque_type"
	_ContextDescriptorKind_name_1 = "classstructenum"
)

var (
	_ContextDescriptorKind_index_0 = [...]uint8{0, 6, 15, 24, 32, 43}
	_ContextDescriptorKind_index_1 = [...]uint8{0, 5, 11, 15}
)

func (i ContextDescriptorKind) String() string {
	switch {
	case i <= 4:
		return _ContextDescriptorKind_name_0[_ContextDescriptorKind_index_0[i]:_ContextDescriptorKind_index_0[i+1]]
	case 16 <= i && i <= 18:
		i -= 16
		return _ContextDescriptorKind_name_1[_ContextDescriptorKind_index_1[i]:_ContextDescriptorKind_index_1[i+1]]
	default:
		return "ContextDescriptorKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DirectTypeDescriptor-0]
	_ = x[IndirectTypeDescriptor-1]
	_ = x[DirectObjCClassName-2]
	_ = x[IndirectObjCClass-3]
}

const _TypeReferenceKind_name = "directindirectdirect_objc_classindirect_objc_class"

var _TypeReferenceKind_index = [...]uint8{0, 6, 14, 31, 50}

func (i TypeReferenceKind) String() string {
	if i >= TypeReferenceKind(len(_TypeReferenceKind_index)-1) {
		return "TypeReferenceKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TypeReferenceKind_name[_TypeReferenceKind_index[i]:_TypeReferenceKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MetadataInitNone-0]
	_ = x[MetadataInitSingleton-1]
	_ = x[MetadataInitForeign-2]
}

const _MetadataInitializationKind_name = "nonesingletonforeign"

var _MetadataInitializationKind_index = [...]uint8{0, 4, 13, 20}

func (i MetadataInitializationKind) String() string {
	if i >= MetadataInitializationKind(len(_MetadataInitializationKind_index)-1) {
		return "MetadataInitializationKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MetadataInitializationKind_name[_MetadataInitializationKind_index[i]:_MetadataInitializationKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SKNone-0]
	_ = x[SKClass-1]
	_ = x[SKMetatype-2]
	_ = x[SKExplicitLayout-3]
}

const _SpecialKind_name = "noneclassmetatypeexplicit layout"

var _SpecialKind_index = [...]uint8{0, 4, 9, 17, 32}

func (i SpecialKind) String() string {
	if i >= SpecialKind(len(_SpecialKind_index)-1) {
		return "SpecialKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SpecialKind_name[_SpecialKind_index[i]:_SpecialKind_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[GPKType-0]
	_ = x[GPKTypePack-1]
	_ = x[GPKMax-63]
}

const (
	_GenericParamKind_name_0 = "typetype_pack"
	_GenericParamKind_name_1 = "max"
)

var (
	_GenericParamKind_index_0 = [...]uint8{0, 4, 13}
)

func (i GenericParamKind) String() string {
	switch {
	case i <= 1:
		return _GenericParamKind_name_0[_GenericParamKind_index_0[i]:_GenericParamKind_index_0[i+1]]
	case i == 63:
		return _GenericParamKind_name_1
	default:
		return "GenericParamKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[GRLKClass-0]
}

const _GenericRequirementLayoutKind_name = "class"

var _GenericRequirementLayoutKind_index = [...]uint8{0, 5}

func (i GenericRequirementLayoutKind) String() string {
	if i >= GenericRequirementLayoutKind(len(_GenericRequirementLayoutKind_index)-1) {
		return "GenericRequirementLayoutKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _GenericRequirementLayoutKind_name[_GenericRequirementLayoutKind_index[i]:_GenericRequirementLayoutKind_index[i+1]]
}
