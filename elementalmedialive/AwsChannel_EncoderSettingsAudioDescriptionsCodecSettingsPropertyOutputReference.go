package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AacSettings() AwsChannel_AacSettingsPropertyOutputReference
	// Experimental.
	AacSettingsInput() *AwsChannel_AacSettingsProperty
	// Experimental.
	Ac3Settings() AwsChannel_Ac3SettingsPropertyOutputReference
	// Experimental.
	Ac3SettingsInput() *AwsChannel_Ac3SettingsProperty
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
	Eac3AtmosSettings() AwsChannel_Eac3AtmosSettingsPropertyOutputReference
	// Experimental.
	Eac3AtmosSettingsInput() *AwsChannel_Eac3AtmosSettingsProperty
	// Experimental.
	Eac3Settings() AwsChannel_Eac3SettingsPropertyOutputReference
	// Experimental.
	Eac3SettingsInput() *AwsChannel_Eac3SettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty)
	// Experimental.
	Mp2Settings() AwsChannel_Mp2SettingsPropertyOutputReference
	// Experimental.
	Mp2SettingsInput() *AwsChannel_Mp2SettingsProperty
	// Experimental.
	PassThroughSettings() AwsChannel_PassThroughSettingsPropertyOutputReference
	// Experimental.
	PassThroughSettingsInput() *AwsChannel_PassThroughSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WavSettings() AwsChannel_WavSettingsPropertyOutputReference
	// Experimental.
	WavSettingsInput() *AwsChannel_WavSettingsProperty
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
	PutAacSettings(value *AwsChannel_AacSettingsProperty)
	// Experimental.
	PutAc3Settings(value *AwsChannel_Ac3SettingsProperty)
	// Experimental.
	PutEac3AtmosSettings(value *AwsChannel_Eac3AtmosSettingsProperty)
	// Experimental.
	PutEac3Settings(value *AwsChannel_Eac3SettingsProperty)
	// Experimental.
	PutMp2Settings(value *AwsChannel_Mp2SettingsProperty)
	// Experimental.
	PutPassThroughSettings(value *AwsChannel_PassThroughSettingsProperty)
	// Experimental.
	PutWavSettings(value *AwsChannel_WavSettingsProperty)
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

// The jsii proxy struct for AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference
type jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) AacSettings() AwsChannel_AacSettingsPropertyOutputReference {
	var returns AwsChannel_AacSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"aacSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) AacSettingsInput() *AwsChannel_AacSettingsProperty {
	var returns *AwsChannel_AacSettingsProperty
	_jsii_.Get(
		j,
		"aacSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Ac3Settings() AwsChannel_Ac3SettingsPropertyOutputReference {
	var returns AwsChannel_Ac3SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"ac3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Ac3SettingsInput() *AwsChannel_Ac3SettingsProperty {
	var returns *AwsChannel_Ac3SettingsProperty
	_jsii_.Get(
		j,
		"ac3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Eac3AtmosSettings() AwsChannel_Eac3AtmosSettingsPropertyOutputReference {
	var returns AwsChannel_Eac3AtmosSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"eac3AtmosSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Eac3AtmosSettingsInput() *AwsChannel_Eac3AtmosSettingsProperty {
	var returns *AwsChannel_Eac3AtmosSettingsProperty
	_jsii_.Get(
		j,
		"eac3AtmosSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Eac3Settings() AwsChannel_Eac3SettingsPropertyOutputReference {
	var returns AwsChannel_Eac3SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"eac3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Eac3SettingsInput() *AwsChannel_Eac3SettingsProperty {
	var returns *AwsChannel_Eac3SettingsProperty
	_jsii_.Get(
		j,
		"eac3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) InternalValue() *AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty {
	var returns *AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Mp2Settings() AwsChannel_Mp2SettingsPropertyOutputReference {
	var returns AwsChannel_Mp2SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"mp2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Mp2SettingsInput() *AwsChannel_Mp2SettingsProperty {
	var returns *AwsChannel_Mp2SettingsProperty
	_jsii_.Get(
		j,
		"mp2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PassThroughSettings() AwsChannel_PassThroughSettingsPropertyOutputReference {
	var returns AwsChannel_PassThroughSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"passThroughSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PassThroughSettingsInput() *AwsChannel_PassThroughSettingsProperty {
	var returns *AwsChannel_PassThroughSettingsProperty
	_jsii_.Get(
		j,
		"passThroughSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) WavSettings() AwsChannel_WavSettingsPropertyOutputReference {
	var returns AwsChannel_WavSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"wavSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) WavSettingsInput() *AwsChannel_WavSettingsProperty {
	var returns *AwsChannel_WavSettingsProperty
	_jsii_.Get(
		j,
		"wavSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference_Override(a AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference)SetInternalValue(val *AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutAacSettings(value *AwsChannel_AacSettingsProperty) {
	if err := a.validatePutAacSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAacSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutAc3Settings(value *AwsChannel_Ac3SettingsProperty) {
	if err := a.validatePutAc3SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAc3Settings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutEac3AtmosSettings(value *AwsChannel_Eac3AtmosSettingsProperty) {
	if err := a.validatePutEac3AtmosSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEac3AtmosSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutEac3Settings(value *AwsChannel_Eac3SettingsProperty) {
	if err := a.validatePutEac3SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEac3Settings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutMp2Settings(value *AwsChannel_Mp2SettingsProperty) {
	if err := a.validatePutMp2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMp2Settings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutPassThroughSettings(value *AwsChannel_PassThroughSettingsProperty) {
	if err := a.validatePutPassThroughSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPassThroughSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutWavSettings(value *AwsChannel_WavSettingsProperty) {
	if err := a.validatePutWavSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWavSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetAacSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetAacSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetAc3Settings() {
	_jsii_.InvokeVoid(
		a,
		"resetAc3Settings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetEac3AtmosSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetEac3AtmosSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetEac3Settings() {
	_jsii_.InvokeVoid(
		a,
		"resetEac3Settings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetMp2Settings() {
	_jsii_.InvokeVoid(
		a,
		"resetMp2Settings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetPassThroughSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetPassThroughSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetWavSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetWavSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

