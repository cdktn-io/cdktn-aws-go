package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMedialiveChannel_M3u8SettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioFramesPerPes() *float64
	// Experimental.
	SetAudioFramesPerPes(val *float64)
	// Experimental.
	AudioFramesPerPesInput() *float64
	// Experimental.
	AudioPids() *string
	// Experimental.
	SetAudioPids(val *string)
	// Experimental.
	AudioPidsInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EcmPid() *string
	// Experimental.
	SetEcmPid(val *string)
	// Experimental.
	EcmPidInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsMedialiveChannel_M3u8SettingsProperty
	// Experimental.
	SetInternalValue(val *AwsMedialiveChannel_M3u8SettingsProperty)
	// Experimental.
	NielsenId3Behavior() *string
	// Experimental.
	SetNielsenId3Behavior(val *string)
	// Experimental.
	NielsenId3BehaviorInput() *string
	// Experimental.
	PatInterval() *float64
	// Experimental.
	SetPatInterval(val *float64)
	// Experimental.
	PatIntervalInput() *float64
	// Experimental.
	PcrControl() *string
	// Experimental.
	SetPcrControl(val *string)
	// Experimental.
	PcrControlInput() *string
	// Experimental.
	PcrPeriod() *float64
	// Experimental.
	SetPcrPeriod(val *float64)
	// Experimental.
	PcrPeriodInput() *float64
	// Experimental.
	PcrPid() *string
	// Experimental.
	SetPcrPid(val *string)
	// Experimental.
	PcrPidInput() *string
	// Experimental.
	PmtInterval() *float64
	// Experimental.
	SetPmtInterval(val *float64)
	// Experimental.
	PmtIntervalInput() *float64
	// Experimental.
	PmtPid() *string
	// Experimental.
	SetPmtPid(val *string)
	// Experimental.
	PmtPidInput() *string
	// Experimental.
	ProgramNum() *float64
	// Experimental.
	SetProgramNum(val *float64)
	// Experimental.
	ProgramNumInput() *float64
	// Experimental.
	Scte35Behavior() *string
	// Experimental.
	SetScte35Behavior(val *string)
	// Experimental.
	Scte35BehaviorInput() *string
	// Experimental.
	Scte35Pid() *string
	// Experimental.
	SetScte35Pid(val *string)
	// Experimental.
	Scte35PidInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimedMetadataBehavior() *string
	// Experimental.
	SetTimedMetadataBehavior(val *string)
	// Experimental.
	TimedMetadataBehaviorInput() *string
	// Experimental.
	TimedMetadataPid() *string
	// Experimental.
	SetTimedMetadataPid(val *string)
	// Experimental.
	TimedMetadataPidInput() *string
	// Experimental.
	TransportStreamId() *float64
	// Experimental.
	SetTransportStreamId(val *float64)
	// Experimental.
	TransportStreamIdInput() *float64
	// Experimental.
	VideoPid() *string
	// Experimental.
	SetVideoPid(val *string)
	// Experimental.
	VideoPidInput() *string
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
	ResetAudioFramesPerPes()
	// Experimental.
	ResetAudioPids()
	// Experimental.
	ResetEcmPid()
	// Experimental.
	ResetNielsenId3Behavior()
	// Experimental.
	ResetPatInterval()
	// Experimental.
	ResetPcrControl()
	// Experimental.
	ResetPcrPeriod()
	// Experimental.
	ResetPcrPid()
	// Experimental.
	ResetPmtInterval()
	// Experimental.
	ResetPmtPid()
	// Experimental.
	ResetProgramNum()
	// Experimental.
	ResetScte35Behavior()
	// Experimental.
	ResetScte35Pid()
	// Experimental.
	ResetTimedMetadataBehavior()
	// Experimental.
	ResetTimedMetadataPid()
	// Experimental.
	ResetTransportStreamId()
	// Experimental.
	ResetVideoPid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMedialiveChannel_M3u8SettingsPropertyOutputReference
type jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) AudioFramesPerPes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"audioFramesPerPes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) AudioFramesPerPesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"audioFramesPerPesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) AudioPids() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioPids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) AudioPidsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioPidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) EcmPid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecmPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) EcmPidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecmPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) InternalValue() *AwsMedialiveChannel_M3u8SettingsProperty {
	var returns *AwsMedialiveChannel_M3u8SettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) NielsenId3Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nielsenId3Behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) NielsenId3BehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nielsenId3BehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) PatInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"patInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) PatIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"patIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) PcrControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pcrControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) PcrControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pcrControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) PcrPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pcrPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) PcrPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pcrPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) PcrPid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pcrPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) PcrPidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pcrPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) PmtInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pmtInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) PmtIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pmtIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) PmtPid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pmtPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) PmtPidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pmtPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ProgramNum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programNum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ProgramNumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"programNumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) Scte35Behavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scte35Behavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) Scte35BehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scte35BehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) Scte35Pid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scte35Pid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) Scte35PidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scte35PidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) TimedMetadataBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timedMetadataBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) TimedMetadataBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timedMetadataBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) TimedMetadataPid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timedMetadataPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) TimedMetadataPidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timedMetadataPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) TransportStreamId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) TransportStreamIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transportStreamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) VideoPid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) VideoPidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"videoPidInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMedialiveChannel_M3u8SettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMedialiveChannel_M3u8SettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMedialiveChannel_M3u8SettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.M3u8SettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMedialiveChannel_M3u8SettingsPropertyOutputReference_Override(a AwsMedialiveChannel_M3u8SettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.M3u8SettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetAudioFramesPerPes(val *float64) {
	if err := j.validateSetAudioFramesPerPesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioFramesPerPes",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetAudioPids(val *string) {
	if err := j.validateSetAudioPidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioPids",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetEcmPid(val *string) {
	if err := j.validateSetEcmPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecmPid",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetInternalValue(val *AwsMedialiveChannel_M3u8SettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetNielsenId3Behavior(val *string) {
	if err := j.validateSetNielsenId3BehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nielsenId3Behavior",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetPatInterval(val *float64) {
	if err := j.validateSetPatIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patInterval",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetPcrControl(val *string) {
	if err := j.validateSetPcrControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pcrControl",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetPcrPeriod(val *float64) {
	if err := j.validateSetPcrPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pcrPeriod",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetPcrPid(val *string) {
	if err := j.validateSetPcrPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pcrPid",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetPmtInterval(val *float64) {
	if err := j.validateSetPmtIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pmtInterval",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetPmtPid(val *string) {
	if err := j.validateSetPmtPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pmtPid",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetProgramNum(val *float64) {
	if err := j.validateSetProgramNumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"programNum",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetScte35Behavior(val *string) {
	if err := j.validateSetScte35BehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scte35Behavior",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetScte35Pid(val *string) {
	if err := j.validateSetScte35PidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scte35Pid",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetTimedMetadataBehavior(val *string) {
	if err := j.validateSetTimedMetadataBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timedMetadataBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetTimedMetadataPid(val *string) {
	if err := j.validateSetTimedMetadataPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timedMetadataPid",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetTransportStreamId(val *float64) {
	if err := j.validateSetTransportStreamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportStreamId",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference)SetVideoPid(val *string) {
	if err := j.validateSetVideoPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"videoPid",
		val,
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetAudioFramesPerPes() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioFramesPerPes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetAudioPids() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioPids",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetEcmPid() {
	_jsii_.InvokeVoid(
		a,
		"resetEcmPid",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetNielsenId3Behavior() {
	_jsii_.InvokeVoid(
		a,
		"resetNielsenId3Behavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetPatInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetPatInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetPcrControl() {
	_jsii_.InvokeVoid(
		a,
		"resetPcrControl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetPcrPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetPcrPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetPcrPid() {
	_jsii_.InvokeVoid(
		a,
		"resetPcrPid",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetPmtInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetPmtInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetPmtPid() {
	_jsii_.InvokeVoid(
		a,
		"resetPmtPid",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetProgramNum() {
	_jsii_.InvokeVoid(
		a,
		"resetProgramNum",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetScte35Behavior() {
	_jsii_.InvokeVoid(
		a,
		"resetScte35Behavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetScte35Pid() {
	_jsii_.InvokeVoid(
		a,
		"resetScte35Pid",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetTimedMetadataBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetTimedMetadataBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetTimedMetadataPid() {
	_jsii_.InvokeVoid(
		a,
		"resetTimedMetadataPid",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetTransportStreamId() {
	_jsii_.InvokeVoid(
		a,
		"resetTransportStreamId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ResetVideoPid() {
	_jsii_.InvokeVoid(
		a,
		"resetVideoPid",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_M3u8SettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

