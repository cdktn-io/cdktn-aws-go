package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AacSettings() TfChannel_AacSettingsPropertyOutputReference
	// Experimental.
	AacSettingsInput() *TfChannel_AacSettingsProperty
	// Experimental.
	Ac3Settings() TfChannel_Ac3SettingsPropertyOutputReference
	// Experimental.
	Ac3SettingsInput() *TfChannel_Ac3SettingsProperty
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
	Eac3AtmosSettings() TfChannel_Eac3AtmosSettingsPropertyOutputReference
	// Experimental.
	Eac3AtmosSettingsInput() *TfChannel_Eac3AtmosSettingsProperty
	// Experimental.
	Eac3Settings() TfChannel_Eac3SettingsPropertyOutputReference
	// Experimental.
	Eac3SettingsInput() *TfChannel_Eac3SettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty)
	// Experimental.
	Mp2Settings() TfChannel_Mp2SettingsPropertyOutputReference
	// Experimental.
	Mp2SettingsInput() *TfChannel_Mp2SettingsProperty
	// Experimental.
	PassThroughSettings() TfChannel_PassThroughSettingsPropertyOutputReference
	// Experimental.
	PassThroughSettingsInput() *TfChannel_PassThroughSettingsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WavSettings() TfChannel_WavSettingsPropertyOutputReference
	// Experimental.
	WavSettingsInput() *TfChannel_WavSettingsProperty
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
	PutAacSettings(value *TfChannel_AacSettingsProperty)
	// Experimental.
	PutAc3Settings(value *TfChannel_Ac3SettingsProperty)
	// Experimental.
	PutEac3AtmosSettings(value *TfChannel_Eac3AtmosSettingsProperty)
	// Experimental.
	PutEac3Settings(value *TfChannel_Eac3SettingsProperty)
	// Experimental.
	PutMp2Settings(value *TfChannel_Mp2SettingsProperty)
	// Experimental.
	PutPassThroughSettings(value *TfChannel_PassThroughSettingsProperty)
	// Experimental.
	PutWavSettings(value *TfChannel_WavSettingsProperty)
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

// The jsii proxy struct for TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference
type jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) AacSettings() TfChannel_AacSettingsPropertyOutputReference {
	var returns TfChannel_AacSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"aacSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) AacSettingsInput() *TfChannel_AacSettingsProperty {
	var returns *TfChannel_AacSettingsProperty
	_jsii_.Get(
		j,
		"aacSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Ac3Settings() TfChannel_Ac3SettingsPropertyOutputReference {
	var returns TfChannel_Ac3SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"ac3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Ac3SettingsInput() *TfChannel_Ac3SettingsProperty {
	var returns *TfChannel_Ac3SettingsProperty
	_jsii_.Get(
		j,
		"ac3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Eac3AtmosSettings() TfChannel_Eac3AtmosSettingsPropertyOutputReference {
	var returns TfChannel_Eac3AtmosSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"eac3AtmosSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Eac3AtmosSettingsInput() *TfChannel_Eac3AtmosSettingsProperty {
	var returns *TfChannel_Eac3AtmosSettingsProperty
	_jsii_.Get(
		j,
		"eac3AtmosSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Eac3Settings() TfChannel_Eac3SettingsPropertyOutputReference {
	var returns TfChannel_Eac3SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"eac3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Eac3SettingsInput() *TfChannel_Eac3SettingsProperty {
	var returns *TfChannel_Eac3SettingsProperty
	_jsii_.Get(
		j,
		"eac3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) InternalValue() *TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty {
	var returns *TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Mp2Settings() TfChannel_Mp2SettingsPropertyOutputReference {
	var returns TfChannel_Mp2SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"mp2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Mp2SettingsInput() *TfChannel_Mp2SettingsProperty {
	var returns *TfChannel_Mp2SettingsProperty
	_jsii_.Get(
		j,
		"mp2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PassThroughSettings() TfChannel_PassThroughSettingsPropertyOutputReference {
	var returns TfChannel_PassThroughSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"passThroughSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PassThroughSettingsInput() *TfChannel_PassThroughSettingsProperty {
	var returns *TfChannel_PassThroughSettingsProperty
	_jsii_.Get(
		j,
		"passThroughSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) WavSettings() TfChannel_WavSettingsPropertyOutputReference {
	var returns TfChannel_WavSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"wavSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) WavSettingsInput() *TfChannel_WavSettingsProperty {
	var returns *TfChannel_WavSettingsProperty
	_jsii_.Get(
		j,
		"wavSettingsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference_Override(t TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutAacSettings(value *TfChannel_AacSettingsProperty) {
	if err := t.validatePutAacSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAacSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutAc3Settings(value *TfChannel_Ac3SettingsProperty) {
	if err := t.validatePutAc3SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAc3Settings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutEac3AtmosSettings(value *TfChannel_Eac3AtmosSettingsProperty) {
	if err := t.validatePutEac3AtmosSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEac3AtmosSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutEac3Settings(value *TfChannel_Eac3SettingsProperty) {
	if err := t.validatePutEac3SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEac3Settings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutMp2Settings(value *TfChannel_Mp2SettingsProperty) {
	if err := t.validatePutMp2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMp2Settings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutPassThroughSettings(value *TfChannel_PassThroughSettingsProperty) {
	if err := t.validatePutPassThroughSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPassThroughSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) PutWavSettings(value *TfChannel_WavSettingsProperty) {
	if err := t.validatePutWavSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWavSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetAacSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetAacSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetAc3Settings() {
	_jsii_.InvokeVoid(
		t,
		"resetAc3Settings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetEac3AtmosSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetEac3AtmosSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetEac3Settings() {
	_jsii_.InvokeVoid(
		t,
		"resetEac3Settings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetMp2Settings() {
	_jsii_.InvokeVoid(
		t,
		"resetMp2Settings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetPassThroughSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetPassThroughSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ResetWavSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetWavSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

