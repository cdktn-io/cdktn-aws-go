package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AacSettings() AwsMedialiveChannel_AacSettingsPropertyOutputReference
	// Experimental.
	AacSettingsInput() *AwsMedialiveChannel_AacSettingsProperty
	// Experimental.
	Ac3Settings() AwsMedialiveChannel_Ac3SettingsPropertyOutputReference
	// Experimental.
	Ac3SettingsInput() *AwsMedialiveChannel_Ac3SettingsProperty
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
	Eac3AtmosSettings() AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference
	// Experimental.
	Eac3AtmosSettingsInput() *AwsMedialiveChannel_Eac3AtmosSettingsProperty
	// Experimental.
	Eac3Settings() AwsMedialiveChannel_Eac3SettingsPropertyOutputReference
	// Experimental.
	Eac3SettingsInput() *AwsMedialiveChannel_Eac3SettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty)
	// Experimental.
	Mp2Settings() AwsMedialiveChannel_Mp2SettingsPropertyOutputReference
	// Experimental.
	Mp2SettingsInput() *AwsMedialiveChannel_Mp2SettingsProperty
	// Experimental.
	PassThroughSettings() AwsMedialiveChannel_PassThroughSettingsPropertyOutputReference
	// Experimental.
	PassThroughSettingsInput() *AwsMedialiveChannel_PassThroughSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WavSettings() AwsMedialiveChannel_WavSettingsPropertyOutputReference
	// Experimental.
	WavSettingsInput() *AwsMedialiveChannel_WavSettingsProperty
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
	PutAacSettings(value *AwsMedialiveChannel_AacSettingsProperty)
	// Experimental.
	PutAc3Settings(value *AwsMedialiveChannel_Ac3SettingsProperty)
	// Experimental.
	PutEac3AtmosSettings(value *AwsMedialiveChannel_Eac3AtmosSettingsProperty)
	// Experimental.
	PutEac3Settings(value *AwsMedialiveChannel_Eac3SettingsProperty)
	// Experimental.
	PutMp2Settings(value *AwsMedialiveChannel_Mp2SettingsProperty)
	// Experimental.
	PutPassThroughSettings(value *AwsMedialiveChannel_PassThroughSettingsProperty)
	// Experimental.
	PutWavSettings(value *AwsMedialiveChannel_WavSettingsProperty)
	// Experimental.
	ResetAacSettings()
	// Experimental.
	ResetAc3Settings()
	// Experimental.
	ResetEac3AtmosSettings()
	// Experimental.
	ResetEac3Settings()
	// Experimental.
	ResetMp2Settings()
	// Experimental.
	ResetPassThroughSettings()
	// Experimental.
	ResetWavSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference
type jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) AacSettings() AwsMedialiveChannel_AacSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_AacSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"aacSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) AacSettingsInput() *AwsMedialiveChannel_AacSettingsProperty {
	var returns *AwsMedialiveChannel_AacSettingsProperty
	_jsii_.Get(
		j,
		"aacSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Ac3Settings() AwsMedialiveChannel_Ac3SettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_Ac3SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"ac3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Ac3SettingsInput() *AwsMedialiveChannel_Ac3SettingsProperty {
	var returns *AwsMedialiveChannel_Ac3SettingsProperty
	_jsii_.Get(
		j,
		"ac3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Eac3AtmosSettings() AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_Eac3AtmosSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"eac3AtmosSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Eac3AtmosSettingsInput() *AwsMedialiveChannel_Eac3AtmosSettingsProperty {
	var returns *AwsMedialiveChannel_Eac3AtmosSettingsProperty
	_jsii_.Get(
		j,
		"eac3AtmosSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Eac3Settings() AwsMedialiveChannel_Eac3SettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_Eac3SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"eac3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Eac3SettingsInput() *AwsMedialiveChannel_Eac3SettingsProperty {
	var returns *AwsMedialiveChannel_Eac3SettingsProperty
	_jsii_.Get(
		j,
		"eac3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) InternalValue() *AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty {
	var returns *AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Mp2Settings() AwsMedialiveChannel_Mp2SettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_Mp2SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"mp2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Mp2SettingsInput() *AwsMedialiveChannel_Mp2SettingsProperty {
	var returns *AwsMedialiveChannel_Mp2SettingsProperty
	_jsii_.Get(
		j,
		"mp2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PassThroughSettings() AwsMedialiveChannel_PassThroughSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_PassThroughSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"passThroughSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PassThroughSettingsInput() *AwsMedialiveChannel_PassThroughSettingsProperty {
	var returns *AwsMedialiveChannel_PassThroughSettingsProperty
	_jsii_.Get(
		j,
		"passThroughSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) WavSettings() AwsMedialiveChannel_WavSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_WavSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"wavSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) WavSettingsInput() *AwsMedialiveChannel_WavSettingsProperty {
	var returns *AwsMedialiveChannel_WavSettingsProperty
	_jsii_.Get(
		j,
		"wavSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference_Override(a AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference)SetInternalValue(val *AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutAacSettings(value *AwsMedialiveChannel_AacSettingsProperty) {
	if err := a.validatePutAacSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAacSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutAc3Settings(value *AwsMedialiveChannel_Ac3SettingsProperty) {
	if err := a.validatePutAc3SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAc3Settings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutEac3AtmosSettings(value *AwsMedialiveChannel_Eac3AtmosSettingsProperty) {
	if err := a.validatePutEac3AtmosSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEac3AtmosSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutEac3Settings(value *AwsMedialiveChannel_Eac3SettingsProperty) {
	if err := a.validatePutEac3SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEac3Settings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutMp2Settings(value *AwsMedialiveChannel_Mp2SettingsProperty) {
	if err := a.validatePutMp2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMp2Settings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutPassThroughSettings(value *AwsMedialiveChannel_PassThroughSettingsProperty) {
	if err := a.validatePutPassThroughSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPassThroughSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutWavSettings(value *AwsMedialiveChannel_WavSettingsProperty) {
	if err := a.validatePutWavSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWavSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetAacSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetAacSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetAc3Settings() {
	_jsii_.InvokeVoid(
		a,
		"resetAc3Settings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetEac3AtmosSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetEac3AtmosSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetEac3Settings() {
	_jsii_.InvokeVoid(
		a,
		"resetEac3Settings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetMp2Settings() {
	_jsii_.InvokeVoid(
		a,
		"resetMp2Settings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetPassThroughSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetPassThroughSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetWavSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetWavSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

