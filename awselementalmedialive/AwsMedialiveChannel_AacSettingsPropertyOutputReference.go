package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMedialiveChannel_AacSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Bitrate() *float64
	// Experimental.
	SetBitrate(val *float64)
	// Experimental.
	BitrateInput() *float64
	// Experimental.
	CodingMode() *string
	// Experimental.
	SetCodingMode(val *string)
	// Experimental.
	CodingModeInput() *string
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
	Fqn() *string
	// Experimental.
	InputType() *string
	// Experimental.
	SetInputType(val *string)
	// Experimental.
	InputTypeInput() *string
	// Experimental.
	InternalValue() *AwsMedialiveChannel_AacSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsMedialiveChannel_AacSettingsProperty)
	// Experimental.
	Profile() *string
	// Experimental.
	SetProfile(val *string)
	// Experimental.
	ProfileInput() *string
	// Experimental.
	RateControlMode() *string
	// Experimental.
	SetRateControlMode(val *string)
	// Experimental.
	RateControlModeInput() *string
	// Experimental.
	RawFormat() *string
	// Experimental.
	SetRawFormat(val *string)
	// Experimental.
	RawFormatInput() *string
	// Experimental.
	SampleRate() *float64
	// Experimental.
	SetSampleRate(val *float64)
	// Experimental.
	SampleRateInput() *float64
	// Experimental.
	Spec() *string
	// Experimental.
	SetSpec(val *string)
	// Experimental.
	SpecInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VbrQuality() *string
	// Experimental.
	SetVbrQuality(val *string)
	// Experimental.
	VbrQualityInput() *string
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
	ResetBitrate()
	// Experimental.
	ResetCodingMode()
	// Experimental.
	ResetInputType()
	// Experimental.
	ResetProfile()
	// Experimental.
	ResetRateControlMode()
	// Experimental.
	ResetRawFormat()
	// Experimental.
	ResetSampleRate()
	// Experimental.
	ResetSpec()
	// Experimental.
	ResetVbrQuality()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMedialiveChannel_AacSettingsPropertyOutputReference
type jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) Bitrate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) BitrateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) CodingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) CodingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) InputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) InputTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) InternalValue() *AwsMedialiveChannel_AacSettingsProperty {
	var returns *AwsMedialiveChannel_AacSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) RateControlMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateControlMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) RateControlModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rateControlModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) RawFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rawFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) RawFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rawFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) SampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) SampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) Spec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) SpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) VbrQuality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vbrQuality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) VbrQualityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vbrQualityInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMedialiveChannel_AacSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMedialiveChannel_AacSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMedialiveChannel_AacSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.AacSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMedialiveChannel_AacSettingsPropertyOutputReference_Override(a AwsMedialiveChannel_AacSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.AacSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference)SetBitrate(val *float64) {
	if err := j.validateSetBitrateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitrate",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference)SetCodingMode(val *string) {
	if err := j.validateSetCodingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codingMode",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference)SetInputType(val *string) {
	if err := j.validateSetInputTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputType",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference)SetInternalValue(val *AwsMedialiveChannel_AacSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference)SetProfile(val *string) {
	if err := j.validateSetProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference)SetRateControlMode(val *string) {
	if err := j.validateSetRateControlModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rateControlMode",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference)SetRawFormat(val *string) {
	if err := j.validateSetRawFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rawFormat",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference)SetSampleRate(val *float64) {
	if err := j.validateSetSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleRate",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference)SetSpec(val *string) {
	if err := j.validateSetSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spec",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference)SetVbrQuality(val *string) {
	if err := j.validateSetVbrQualityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vbrQuality",
		val,
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) ResetBitrate() {
	_jsii_.InvokeVoid(
		a,
		"resetBitrate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) ResetCodingMode() {
	_jsii_.InvokeVoid(
		a,
		"resetCodingMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) ResetInputType() {
	_jsii_.InvokeVoid(
		a,
		"resetInputType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) ResetProfile() {
	_jsii_.InvokeVoid(
		a,
		"resetProfile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) ResetRateControlMode() {
	_jsii_.InvokeVoid(
		a,
		"resetRateControlMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) ResetRawFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetRawFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) ResetSampleRate() {
	_jsii_.InvokeVoid(
		a,
		"resetSampleRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) ResetSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) ResetVbrQuality() {
	_jsii_.InvokeVoid(
		a,
		"resetVbrQuality",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_AacSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

