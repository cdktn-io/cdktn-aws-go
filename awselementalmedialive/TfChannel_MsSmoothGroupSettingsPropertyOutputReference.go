package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_MsSmoothGroupSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AcquisitionPointId() *string
	// Experimental.
	SetAcquisitionPointId(val *string)
	// Experimental.
	AcquisitionPointIdInput() *string
	// Experimental.
	AudioOnlyTimecodeControl() *string
	// Experimental.
	SetAudioOnlyTimecodeControl(val *string)
	// Experimental.
	AudioOnlyTimecodeControlInput() *string
	// Experimental.
	CertificateMode() *string
	// Experimental.
	SetCertificateMode(val *string)
	// Experimental.
	CertificateModeInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// Experimental.
	ConnectionRetryInterval() *float64
	// Experimental.
	SetConnectionRetryInterval(val *float64)
	// Experimental.
	ConnectionRetryIntervalInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Destination() TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference
	// Experimental.
	DestinationInput() *TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty
	// Experimental.
	EventId() *string
	// Experimental.
	SetEventId(val *string)
	// Experimental.
	EventIdInput() *string
	// Experimental.
	EventIdMode() *string
	// Experimental.
	SetEventIdMode(val *string)
	// Experimental.
	EventIdModeInput() *string
	// Experimental.
	EventStopBehavior() *string
	// Experimental.
	SetEventStopBehavior(val *string)
	// Experimental.
	EventStopBehaviorInput() *string
	// Experimental.
	FilecacheDuration() *float64
	// Experimental.
	SetFilecacheDuration(val *float64)
	// Experimental.
	FilecacheDurationInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	FragmentLength() *float64
	// Experimental.
	SetFragmentLength(val *float64)
	// Experimental.
	FragmentLengthInput() *float64
	// Experimental.
	InputLossAction() *string
	// Experimental.
	SetInputLossAction(val *string)
	// Experimental.
	InputLossActionInput() *string
	// Experimental.
	InternalValue() *TfChannel_MsSmoothGroupSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_MsSmoothGroupSettingsProperty)
	// Experimental.
	NumRetries() *float64
	// Experimental.
	SetNumRetries(val *float64)
	// Experimental.
	NumRetriesInput() *float64
	// Experimental.
	RestartDelay() *float64
	// Experimental.
	SetRestartDelay(val *float64)
	// Experimental.
	RestartDelayInput() *float64
	// Experimental.
	SegmentationMode() *string
	// Experimental.
	SetSegmentationMode(val *string)
	// Experimental.
	SegmentationModeInput() *string
	// Experimental.
	SendDelayMs() *float64
	// Experimental.
	SetSendDelayMs(val *float64)
	// Experimental.
	SendDelayMsInput() *float64
	// Experimental.
	SparseTrackType() *string
	// Experimental.
	SetSparseTrackType(val *string)
	// Experimental.
	SparseTrackTypeInput() *string
	// Experimental.
	StreamManifestBehavior() *string
	// Experimental.
	SetStreamManifestBehavior(val *string)
	// Experimental.
	StreamManifestBehaviorInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimestampOffset() *string
	// Experimental.
	SetTimestampOffset(val *string)
	// Experimental.
	TimestampOffsetInput() *string
	// Experimental.
	TimestampOffsetMode() *string
	// Experimental.
	SetTimestampOffsetMode(val *string)
	// Experimental.
	TimestampOffsetModeInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	PutDestination(value *TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty)
	// Experimental.
	ResetAcquisitionPointId()
	// Experimental.
	ResetAudioOnlyTimecodeControl()
	// Experimental.
	ResetCertificateMode()
	// Experimental.
	ResetConnectionRetryInterval()
	// Experimental.
	ResetEventId()
	// Experimental.
	ResetEventIdMode()
	// Experimental.
	ResetEventStopBehavior()
	// Experimental.
	ResetFilecacheDuration()
	// Experimental.
	ResetFragmentLength()
	// Experimental.
	ResetInputLossAction()
	// Experimental.
	ResetNumRetries()
	// Experimental.
	ResetRestartDelay()
	// Experimental.
	ResetSegmentationMode()
	// Experimental.
	ResetSendDelayMs()
	// Experimental.
	ResetSparseTrackType()
	// Experimental.
	ResetStreamManifestBehavior()
	// Experimental.
	ResetTimestampOffset()
	// Experimental.
	ResetTimestampOffsetMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_MsSmoothGroupSettingsPropertyOutputReference
type jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) AcquisitionPointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acquisitionPointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) AcquisitionPointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acquisitionPointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) AudioOnlyTimecodeControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioOnlyTimecodeControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) AudioOnlyTimecodeControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioOnlyTimecodeControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) CertificateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) CertificateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ConnectionRetryInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionRetryInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ConnectionRetryIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionRetryIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) Destination() TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference {
	var returns TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) DestinationInput() *TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty {
	var returns *TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) EventId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) EventIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) EventIdMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventIdMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) EventIdModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventIdModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) EventStopBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventStopBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) EventStopBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventStopBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) FilecacheDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"filecacheDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) FilecacheDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"filecacheDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) FragmentLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fragmentLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) FragmentLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fragmentLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) InputLossAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) InputLossActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) InternalValue() *TfChannel_MsSmoothGroupSettingsProperty {
	var returns *TfChannel_MsSmoothGroupSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) NumRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) NumRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) RestartDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restartDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) RestartDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restartDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) SegmentationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) SegmentationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) SendDelayMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sendDelayMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) SendDelayMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sendDelayMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) SparseTrackType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparseTrackType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) SparseTrackTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparseTrackTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) StreamManifestBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamManifestBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) StreamManifestBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamManifestBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) TimestampOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) TimestampOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) TimestampOffsetMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOffsetMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) TimestampOffsetModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOffsetModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_MsSmoothGroupSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_MsSmoothGroupSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_MsSmoothGroupSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.MsSmoothGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_MsSmoothGroupSettingsPropertyOutputReference_Override(t TfChannel_MsSmoothGroupSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.MsSmoothGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetAcquisitionPointId(val *string) {
	if err := j.validateSetAcquisitionPointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acquisitionPointId",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetAudioOnlyTimecodeControl(val *string) {
	if err := j.validateSetAudioOnlyTimecodeControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioOnlyTimecodeControl",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetCertificateMode(val *string) {
	if err := j.validateSetCertificateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateMode",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetConnectionRetryInterval(val *float64) {
	if err := j.validateSetConnectionRetryIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionRetryInterval",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetEventId(val *string) {
	if err := j.validateSetEventIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventId",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetEventIdMode(val *string) {
	if err := j.validateSetEventIdModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventIdMode",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetEventStopBehavior(val *string) {
	if err := j.validateSetEventStopBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventStopBehavior",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetFilecacheDuration(val *float64) {
	if err := j.validateSetFilecacheDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filecacheDuration",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetFragmentLength(val *float64) {
	if err := j.validateSetFragmentLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fragmentLength",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetInputLossAction(val *string) {
	if err := j.validateSetInputLossActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputLossAction",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_MsSmoothGroupSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetNumRetries(val *float64) {
	if err := j.validateSetNumRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numRetries",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetRestartDelay(val *float64) {
	if err := j.validateSetRestartDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartDelay",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetSegmentationMode(val *string) {
	if err := j.validateSetSegmentationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentationMode",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetSendDelayMs(val *float64) {
	if err := j.validateSetSendDelayMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendDelayMs",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetSparseTrackType(val *string) {
	if err := j.validateSetSparseTrackTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparseTrackType",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetStreamManifestBehavior(val *string) {
	if err := j.validateSetStreamManifestBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamManifestBehavior",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetTimestampOffset(val *string) {
	if err := j.validateSetTimestampOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampOffset",
		val,
	)
}

func (j *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference)SetTimestampOffsetMode(val *string) {
	if err := j.validateSetTimestampOffsetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampOffsetMode",
		val,
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) PutDestination(value *TfChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty) {
	if err := t.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDestination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetAcquisitionPointId() {
	_jsii_.InvokeVoid(
		t,
		"resetAcquisitionPointId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetAudioOnlyTimecodeControl() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioOnlyTimecodeControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetCertificateMode() {
	_jsii_.InvokeVoid(
		t,
		"resetCertificateMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetConnectionRetryInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectionRetryInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetEventId() {
	_jsii_.InvokeVoid(
		t,
		"resetEventId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetEventIdMode() {
	_jsii_.InvokeVoid(
		t,
		"resetEventIdMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetEventStopBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetEventStopBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetFilecacheDuration() {
	_jsii_.InvokeVoid(
		t,
		"resetFilecacheDuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetFragmentLength() {
	_jsii_.InvokeVoid(
		t,
		"resetFragmentLength",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetInputLossAction() {
	_jsii_.InvokeVoid(
		t,
		"resetInputLossAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetNumRetries() {
	_jsii_.InvokeVoid(
		t,
		"resetNumRetries",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetRestartDelay() {
	_jsii_.InvokeVoid(
		t,
		"resetRestartDelay",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetSegmentationMode() {
	_jsii_.InvokeVoid(
		t,
		"resetSegmentationMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetSendDelayMs() {
	_jsii_.InvokeVoid(
		t,
		"resetSendDelayMs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetSparseTrackType() {
	_jsii_.InvokeVoid(
		t,
		"resetSparseTrackType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetStreamManifestBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetStreamManifestBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetTimestampOffset() {
	_jsii_.InvokeVoid(
		t,
		"resetTimestampOffset",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetTimestampOffsetMode() {
	_jsii_.InvokeVoid(
		t,
		"resetTimestampOffsetMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_MsSmoothGroupSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

