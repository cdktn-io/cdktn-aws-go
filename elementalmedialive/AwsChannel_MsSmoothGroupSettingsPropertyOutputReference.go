package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_MsSmoothGroupSettingsPropertyOutputReference interface {
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
	Destination() AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference
	// Experimental.
	DestinationInput() *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty
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
	InternalValue() *AwsChannel_MsSmoothGroupSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_MsSmoothGroupSettingsProperty)
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
	PutDestination(value *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty)
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

// The jsii proxy struct for AwsChannel_MsSmoothGroupSettingsPropertyOutputReference
type jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) AcquisitionPointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acquisitionPointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) AcquisitionPointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acquisitionPointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) AudioOnlyTimecodeControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioOnlyTimecodeControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) AudioOnlyTimecodeControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioOnlyTimecodeControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) CertificateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) CertificateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ConnectionRetryInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionRetryInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ConnectionRetryIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionRetryIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) Destination() AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference {
	var returns AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) DestinationInput() *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty {
	var returns *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) EventId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) EventIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) EventIdMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventIdMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) EventIdModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventIdModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) EventStopBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventStopBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) EventStopBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventStopBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) FilecacheDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"filecacheDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) FilecacheDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"filecacheDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) FragmentLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fragmentLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) FragmentLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fragmentLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) InputLossAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) InputLossActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) InternalValue() *AwsChannel_MsSmoothGroupSettingsProperty {
	var returns *AwsChannel_MsSmoothGroupSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) NumRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) NumRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) RestartDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restartDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) RestartDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restartDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) SegmentationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) SegmentationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) SendDelayMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sendDelayMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) SendDelayMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sendDelayMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) SparseTrackType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparseTrackType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) SparseTrackTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparseTrackTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) StreamManifestBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamManifestBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) StreamManifestBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamManifestBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) TimestampOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) TimestampOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) TimestampOffsetMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOffsetMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) TimestampOffsetModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampOffsetModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_MsSmoothGroupSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_MsSmoothGroupSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_MsSmoothGroupSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.MsSmoothGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_MsSmoothGroupSettingsPropertyOutputReference_Override(a AwsChannel_MsSmoothGroupSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.MsSmoothGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetAcquisitionPointId(val *string) {
	if err := j.validateSetAcquisitionPointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acquisitionPointId",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetAudioOnlyTimecodeControl(val *string) {
	if err := j.validateSetAudioOnlyTimecodeControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioOnlyTimecodeControl",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetCertificateMode(val *string) {
	if err := j.validateSetCertificateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateMode",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetConnectionRetryInterval(val *float64) {
	if err := j.validateSetConnectionRetryIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionRetryInterval",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetEventId(val *string) {
	if err := j.validateSetEventIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventId",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetEventIdMode(val *string) {
	if err := j.validateSetEventIdModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventIdMode",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetEventStopBehavior(val *string) {
	if err := j.validateSetEventStopBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventStopBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetFilecacheDuration(val *float64) {
	if err := j.validateSetFilecacheDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filecacheDuration",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetFragmentLength(val *float64) {
	if err := j.validateSetFragmentLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fragmentLength",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetInputLossAction(val *string) {
	if err := j.validateSetInputLossActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputLossAction",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetInternalValue(val *AwsChannel_MsSmoothGroupSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetNumRetries(val *float64) {
	if err := j.validateSetNumRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numRetries",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetRestartDelay(val *float64) {
	if err := j.validateSetRestartDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartDelay",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetSegmentationMode(val *string) {
	if err := j.validateSetSegmentationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentationMode",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetSendDelayMs(val *float64) {
	if err := j.validateSetSendDelayMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendDelayMs",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetSparseTrackType(val *string) {
	if err := j.validateSetSparseTrackTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparseTrackType",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetStreamManifestBehavior(val *string) {
	if err := j.validateSetStreamManifestBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamManifestBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetTimestampOffset(val *string) {
	if err := j.validateSetTimestampOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampOffset",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference)SetTimestampOffsetMode(val *string) {
	if err := j.validateSetTimestampOffsetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampOffsetMode",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) PutDestination(value *AwsChannel_EncoderSettingsOutputGroupsOutputGroupSettingsMsSmoothGroupSettingsDestinationProperty) {
	if err := a.validatePutDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDestination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetAcquisitionPointId() {
	_jsii_.InvokeVoid(
		a,
		"resetAcquisitionPointId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetAudioOnlyTimecodeControl() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioOnlyTimecodeControl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetCertificateMode() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetConnectionRetryInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionRetryInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetEventId() {
	_jsii_.InvokeVoid(
		a,
		"resetEventId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetEventIdMode() {
	_jsii_.InvokeVoid(
		a,
		"resetEventIdMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetEventStopBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetEventStopBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetFilecacheDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetFilecacheDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetFragmentLength() {
	_jsii_.InvokeVoid(
		a,
		"resetFragmentLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetInputLossAction() {
	_jsii_.InvokeVoid(
		a,
		"resetInputLossAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetNumRetries() {
	_jsii_.InvokeVoid(
		a,
		"resetNumRetries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetRestartDelay() {
	_jsii_.InvokeVoid(
		a,
		"resetRestartDelay",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetSegmentationMode() {
	_jsii_.InvokeVoid(
		a,
		"resetSegmentationMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetSendDelayMs() {
	_jsii_.InvokeVoid(
		a,
		"resetSendDelayMs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetSparseTrackType() {
	_jsii_.InvokeVoid(
		a,
		"resetSparseTrackType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetStreamManifestBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetStreamManifestBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetTimestampOffset() {
	_jsii_.InvokeVoid(
		a,
		"resetTimestampOffset",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ResetTimestampOffsetMode() {
	_jsii_.InvokeVoid(
		a,
		"resetTimestampOffsetMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_MsSmoothGroupSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

