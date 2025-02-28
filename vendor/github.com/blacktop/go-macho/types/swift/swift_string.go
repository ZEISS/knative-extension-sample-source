// Code generated by "stringer -type SpecialPointerAuthDiscriminators -trimprefix=Disc -output swift_string.go"; DO NOT EDIT.

package swift

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DiscHeapDestructor-48063]
	_ = x[DiscTypeDescriptor-44678]
	_ = x[DiscRuntimeFunctionEntry-25179]
	_ = x[DiscProtocolConformanceDescriptor-50923]
	_ = x[DiscValueWitnessTable-11839]
	_ = x[DiscExtendedExistentialTypeShape-23101]
	_ = x[DiscNonUniqueExtendedExistentialTypeShape-59288]
	_ = x[DiscInitializeBufferWithCopyOfBuffer-55882]
	_ = x[DiscDestroy-1272]
	_ = x[DiscInitializeWithCopy-58298]
	_ = x[DiscAssignWithCopy-34641]
	_ = x[DiscInitializeWithTake-18648]
	_ = x[DiscAssignWithTake-61402]
	_ = x[DiscDestroyArray-9112]
	_ = x[DiscInitializeArrayWithCopy-41052]
	_ = x[DiscInitializeArrayWithTakeFrontToBack-7230]
	_ = x[DiscInitializeArrayWithTakeBackToFront-36307]
	_ = x[DiscStoreExtraInhabitant-31173]
	_ = x[DiscGetExtraInhabitantIndex-11432]
	_ = x[DiscGetEnumTag-41909]
	_ = x[DiscDestructiveProjectEnumData-1053]
	_ = x[DiscDestructiveInjectEnumTag-45796]
	_ = x[DiscGetEnumTagSinglePayload-24816]
	_ = x[DiscStoreEnumTagSinglePayload-41169]
	_ = x[DiscKeyPathDestroy-28786]
	_ = x[DiscKeyPathCopy-28518]
	_ = x[DiscKeyPathEquals-30062]
	_ = x[DiscKeyPathHash-25460]
	_ = x[DiscKeyPathGetter-28530]
	_ = x[DiscKeyPathNonmutatingSetter-28528]
	_ = x[DiscKeyPathMutatingSetter-29801]
	_ = x[DiscKeyPathGetLayout-25459]
	_ = x[DiscKeyPathInitializer-25205]
	_ = x[DiscKeyPathMetadataAccessor-29812]
	_ = x[DiscObjectiveCTypeDiscriminator-12739]
	_ = x[DiscbridgeToObjectiveCDiscriminator-48288]
	_ = x[DiscforceBridgeFromObjectiveCDiscriminator-8955]
	_ = x[DiscconditionallyBridgeFromObjectiveCDiscriminator-39579]
	_ = x[DiscDynamicReplacementScope-18672]
	_ = x[DiscDynamicReplacementKey-11389]
	_ = x[DiscOpaqueReadResumeFunction-56769]
	_ = x[DiscOpaqueModifyResumeFunction-3909]
	_ = x[DiscObjCISA-27361]
	_ = x[DiscObjCSuperclass-46507]
	_ = x[DiscResilientClassStubInitCallback-50801]
	_ = x[DiscJobInvokeFunction-52324]
	_ = x[DiscTaskResumeFunction-11330]
	_ = x[DiscTaskResumeContext-30010]
	_ = x[DiscAsyncRunAndBlockFunction-3848]
	_ = x[DiscAsyncContextParent-48546]
	_ = x[DiscAsyncContextResume-55047]
	_ = x[DiscAsyncContextYield-57863]
	_ = x[DiscCancellationNotificationFunction-6451]
	_ = x[DiscEscalationNotificationFunction-23524]
	_ = x[DiscAsyncThinNullaryFunction-3848]
	_ = x[DiscAsyncFutureFunction-29199]
	_ = x[DiscSwiftAsyncContextExtendedFrameEntry-49946]
	_ = x[DiscClangTypeTaskContinuationFunction-10942]
	_ = x[DiscDispatchInvokeFunction-62611]
	_ = x[DiscAccessibleFunctionRecord-17292]
}

const _SpecialPointerAuthDiscriminators_name = "DestructiveProjectEnumDataDestroyAsyncRunAndBlockFunctionOpaqueModifyResumeFunctionCancellationNotificationFunctionInitializeArrayWithTakeFrontToBackforceBridgeFromObjectiveCDiscriminatorDestroyArrayClangTypeTaskContinuationFunctionTaskResumeFunctionDynamicReplacementKeyGetExtraInhabitantIndexValueWitnessTableObjectiveCTypeDiscriminatorAccessibleFunctionRecordInitializeWithTakeDynamicReplacementScopeExtendedExistentialTypeShapeEscalationNotificationFunctionGetEnumTagSinglePayloadRuntimeFunctionEntryKeyPathInitializerKeyPathGetLayoutKeyPathHashObjCISAKeyPathCopyKeyPathNonmutatingSetterKeyPathGetterKeyPathDestroyAsyncFutureFunctionKeyPathMutatingSetterKeyPathMetadataAccessorTaskResumeContextKeyPathEqualsStoreExtraInhabitantAssignWithCopyInitializeArrayWithTakeBackToFrontconditionallyBridgeFromObjectiveCDiscriminatorInitializeArrayWithCopyStoreEnumTagSinglePayloadGetEnumTagTypeDescriptorDestructiveInjectEnumTagObjCSuperclassHeapDestructorbridgeToObjectiveCDiscriminatorAsyncContextParentSwiftAsyncContextExtendedFrameEntryResilientClassStubInitCallbackProtocolConformanceDescriptorJobInvokeFunctionAsyncContextResumeInitializeBufferWithCopyOfBufferOpaqueReadResumeFunctionAsyncContextYieldInitializeWithCopyNonUniqueExtendedExistentialTypeShapeAssignWithTakeDispatchInvokeFunction"

var _SpecialPointerAuthDiscriminators_map = map[SpecialPointerAuthDiscriminators]string{
	1053:  _SpecialPointerAuthDiscriminators_name[0:26],
	1272:  _SpecialPointerAuthDiscriminators_name[26:33],
	3848:  _SpecialPointerAuthDiscriminators_name[33:57],
	3909:  _SpecialPointerAuthDiscriminators_name[57:83],
	6451:  _SpecialPointerAuthDiscriminators_name[83:115],
	7230:  _SpecialPointerAuthDiscriminators_name[115:149],
	8955:  _SpecialPointerAuthDiscriminators_name[149:187],
	9112:  _SpecialPointerAuthDiscriminators_name[187:199],
	10942: _SpecialPointerAuthDiscriminators_name[199:232],
	11330: _SpecialPointerAuthDiscriminators_name[232:250],
	11389: _SpecialPointerAuthDiscriminators_name[250:271],
	11432: _SpecialPointerAuthDiscriminators_name[271:294],
	11839: _SpecialPointerAuthDiscriminators_name[294:311],
	12739: _SpecialPointerAuthDiscriminators_name[311:338],
	17292: _SpecialPointerAuthDiscriminators_name[338:362],
	18648: _SpecialPointerAuthDiscriminators_name[362:380],
	18672: _SpecialPointerAuthDiscriminators_name[380:403],
	23101: _SpecialPointerAuthDiscriminators_name[403:431],
	23524: _SpecialPointerAuthDiscriminators_name[431:461],
	24816: _SpecialPointerAuthDiscriminators_name[461:484],
	25179: _SpecialPointerAuthDiscriminators_name[484:504],
	25205: _SpecialPointerAuthDiscriminators_name[504:522],
	25459: _SpecialPointerAuthDiscriminators_name[522:538],
	25460: _SpecialPointerAuthDiscriminators_name[538:549],
	27361: _SpecialPointerAuthDiscriminators_name[549:556],
	28518: _SpecialPointerAuthDiscriminators_name[556:567],
	28528: _SpecialPointerAuthDiscriminators_name[567:591],
	28530: _SpecialPointerAuthDiscriminators_name[591:604],
	28786: _SpecialPointerAuthDiscriminators_name[604:618],
	29199: _SpecialPointerAuthDiscriminators_name[618:637],
	29801: _SpecialPointerAuthDiscriminators_name[637:658],
	29812: _SpecialPointerAuthDiscriminators_name[658:681],
	30010: _SpecialPointerAuthDiscriminators_name[681:698],
	30062: _SpecialPointerAuthDiscriminators_name[698:711],
	31173: _SpecialPointerAuthDiscriminators_name[711:731],
	34641: _SpecialPointerAuthDiscriminators_name[731:745],
	36307: _SpecialPointerAuthDiscriminators_name[745:779],
	39579: _SpecialPointerAuthDiscriminators_name[779:825],
	41052: _SpecialPointerAuthDiscriminators_name[825:848],
	41169: _SpecialPointerAuthDiscriminators_name[848:873],
	41909: _SpecialPointerAuthDiscriminators_name[873:883],
	44678: _SpecialPointerAuthDiscriminators_name[883:897],
	45796: _SpecialPointerAuthDiscriminators_name[897:921],
	46507: _SpecialPointerAuthDiscriminators_name[921:935],
	48063: _SpecialPointerAuthDiscriminators_name[935:949],
	48288: _SpecialPointerAuthDiscriminators_name[949:980],
	48546: _SpecialPointerAuthDiscriminators_name[980:998],
	49946: _SpecialPointerAuthDiscriminators_name[998:1033],
	50801: _SpecialPointerAuthDiscriminators_name[1033:1063],
	50923: _SpecialPointerAuthDiscriminators_name[1063:1092],
	52324: _SpecialPointerAuthDiscriminators_name[1092:1109],
	55047: _SpecialPointerAuthDiscriminators_name[1109:1127],
	55882: _SpecialPointerAuthDiscriminators_name[1127:1159],
	56769: _SpecialPointerAuthDiscriminators_name[1159:1183],
	57863: _SpecialPointerAuthDiscriminators_name[1183:1200],
	58298: _SpecialPointerAuthDiscriminators_name[1200:1218],
	59288: _SpecialPointerAuthDiscriminators_name[1218:1255],
	61402: _SpecialPointerAuthDiscriminators_name[1255:1269],
	62611: _SpecialPointerAuthDiscriminators_name[1269:1291],
}

func (i SpecialPointerAuthDiscriminators) String() string {
	if str, ok := _SpecialPointerAuthDiscriminators_map[i]; ok {
		return str
	}
	return "SpecialPointerAuthDiscriminators(" + strconv.FormatInt(int64(i), 10) + ")"
}
