package elastictranscoder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elastictranscoder/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elastictranscoder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPreset_AudioPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioPackingMode() *string
	// Experimental.
	SetAudioPackingMode(val *string)
	// Experimental.
	AudioPackingModeInput() *string
	// Experimental.
	BitRate() *string
	// Experimental.
	SetBitRate(val *string)
	// Experimental.
	BitRateInput() *string
	// Experimental.
	Channels() *string
	// Experimental.
	SetChannels(val *string)
	// Experimental.
	ChannelsInput() *string
	// Experimental.
	Codec() *string
	// Experimental.
	SetCodec(val *string)
	// Experimental.
	CodecInput() *string
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
	InternalValue() *AwsPreset_AudioProperty
	// Experimental.
	SetInternalValue(val *AwsPreset_AudioProperty)
	// Experimental.
	SampleRate() *string
	// Experimental.
	SetSampleRate(val *string)
	// Experimental.
	SampleRateInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetAudioPackingMode()
	// Experimental.
	ResetBitRate()
	// Experimental.
	ResetChannels()
	// Experimental.
	ResetCodec()
	// Experimental.
	ResetSampleRate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsPreset_AudioPropertyOutputReference
type jsiiProxy_AwsPreset_AudioPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) AudioPackingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioPackingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) AudioPackingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioPackingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) BitRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) BitRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) Channels() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) ChannelsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) Codec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) CodecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) InternalValue() *AwsPreset_AudioProperty {
	var returns *AwsPreset_AudioProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) SampleRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) SampleRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPreset_AudioPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPreset_AudioPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPreset_AudioPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPreset_AudioPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elastic-transcoder.AwsPreset.AudioPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPreset_AudioPropertyOutputReference_Override(a AwsPreset_AudioPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elastic-transcoder.AwsPreset.AudioPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference)SetAudioPackingMode(val *string) {
	if err := j.validateSetAudioPackingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioPackingMode",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference)SetBitRate(val *string) {
	if err := j.validateSetBitRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitRate",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference)SetChannels(val *string) {
	if err := j.validateSetChannelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channels",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference)SetCodec(val *string) {
	if err := j.validateSetCodecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codec",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference)SetInternalValue(val *AwsPreset_AudioProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference)SetSampleRate(val *string) {
	if err := j.validateSetSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleRate",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) ResetAudioPackingMode() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioPackingMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) ResetBitRate() {
	_jsii_.InvokeVoid(
		a,
		"resetBitRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) ResetChannels() {
	_jsii_.InvokeVoid(
		a,
		"resetChannels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) ResetCodec() {
	_jsii_.InvokeVoid(
		a,
		"resetCodec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) ResetSampleRate() {
	_jsii_.InvokeVoid(
		a,
		"resetSampleRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPreset_AudioPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

