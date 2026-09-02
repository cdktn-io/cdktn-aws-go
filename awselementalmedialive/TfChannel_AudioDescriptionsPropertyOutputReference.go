package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_AudioDescriptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioNormalizationSettings() TfChannel_AudioNormalizationSettingsPropertyOutputReference
	// Experimental.
	AudioNormalizationSettingsInput() *TfChannel_AudioNormalizationSettingsProperty
	// Experimental.
	AudioSelectorName() *string
	// Experimental.
	SetAudioSelectorName(val *string)
	// Experimental.
	AudioSelectorNameInput() *string
	// Experimental.
	AudioType() *string
	// Experimental.
	SetAudioType(val *string)
	// Experimental.
	AudioTypeControl() *string
	// Experimental.
	SetAudioTypeControl(val *string)
	// Experimental.
	AudioTypeControlInput() *string
	// Experimental.
	AudioTypeInput() *string
	// Experimental.
	AudioWatermarkSettings() TfChannel_AudioWatermarkSettingsPropertyOutputReference
	// Experimental.
	AudioWatermarkSettingsInput() *TfChannel_AudioWatermarkSettingsProperty
	// Experimental.
	CodecSettings() TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference
	// Experimental.
	CodecSettingsInput() *TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LanguageCode() *string
	// Experimental.
	SetLanguageCode(val *string)
	// Experimental.
	LanguageCodeControl() *string
	// Experimental.
	SetLanguageCodeControl(val *string)
	// Experimental.
	LanguageCodeControlInput() *string
	// Experimental.
	LanguageCodeInput() *string
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	RemixSettings() TfChannel_RemixSettingsPropertyOutputReference
	// Experimental.
	RemixSettingsInput() *TfChannel_RemixSettingsProperty
	// Experimental.
	StreamName() *string
	// Experimental.
	SetStreamName(val *string)
	// Experimental.
	StreamNameInput() *string
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
	PutAudioNormalizationSettings(value *TfChannel_AudioNormalizationSettingsProperty)
	// Experimental.
	PutAudioWatermarkSettings(value *TfChannel_AudioWatermarkSettingsProperty)
	// Experimental.
	PutCodecSettings(value *TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty)
	// Experimental.
	PutRemixSettings(value *TfChannel_RemixSettingsProperty)
	// Experimental.
	ResetAudioNormalizationSettings()
	// Experimental.
	ResetAudioType()
	// Experimental.
	ResetAudioTypeControl()
	// Experimental.
	ResetAudioWatermarkSettings()
	// Experimental.
	ResetCodecSettings()
	// Experimental.
	ResetLanguageCode()
	// Experimental.
	ResetLanguageCodeControl()
	// Experimental.
	ResetRemixSettings()
	// Experimental.
	ResetStreamName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_AudioDescriptionsPropertyOutputReference
type jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) AudioNormalizationSettings() TfChannel_AudioNormalizationSettingsPropertyOutputReference {
	var returns TfChannel_AudioNormalizationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"audioNormalizationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) AudioNormalizationSettingsInput() *TfChannel_AudioNormalizationSettingsProperty {
	var returns *TfChannel_AudioNormalizationSettingsProperty
	_jsii_.Get(
		j,
		"audioNormalizationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) AudioSelectorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioSelectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) AudioSelectorNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioSelectorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) AudioType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) AudioTypeControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioTypeControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) AudioTypeControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioTypeControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) AudioTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) AudioWatermarkSettings() TfChannel_AudioWatermarkSettingsPropertyOutputReference {
	var returns TfChannel_AudioWatermarkSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"audioWatermarkSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) AudioWatermarkSettingsInput() *TfChannel_AudioWatermarkSettingsProperty {
	var returns *TfChannel_AudioWatermarkSettingsProperty
	_jsii_.Get(
		j,
		"audioWatermarkSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) CodecSettings() TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference {
	var returns TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"codecSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) CodecSettingsInput() *TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty {
	var returns *TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty
	_jsii_.Get(
		j,
		"codecSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) LanguageCodeControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) LanguageCodeControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) RemixSettings() TfChannel_RemixSettingsPropertyOutputReference {
	var returns TfChannel_RemixSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"remixSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) RemixSettingsInput() *TfChannel_RemixSettingsProperty {
	var returns *TfChannel_RemixSettingsProperty
	_jsii_.Get(
		j,
		"remixSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) StreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) StreamNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_AudioDescriptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfChannel_AudioDescriptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_AudioDescriptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.AudioDescriptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_AudioDescriptionsPropertyOutputReference_Override(t TfChannel_AudioDescriptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.AudioDescriptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference)SetAudioSelectorName(val *string) {
	if err := j.validateSetAudioSelectorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioSelectorName",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference)SetAudioType(val *string) {
	if err := j.validateSetAudioTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioType",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference)SetAudioTypeControl(val *string) {
	if err := j.validateSetAudioTypeControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioTypeControl",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference)SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference)SetLanguageCodeControl(val *string) {
	if err := j.validateSetLanguageCodeControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCodeControl",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference)SetStreamName(val *string) {
	if err := j.validateSetStreamNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamName",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) PutAudioNormalizationSettings(value *TfChannel_AudioNormalizationSettingsProperty) {
	if err := t.validatePutAudioNormalizationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAudioNormalizationSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) PutAudioWatermarkSettings(value *TfChannel_AudioWatermarkSettingsProperty) {
	if err := t.validatePutAudioWatermarkSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAudioWatermarkSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) PutCodecSettings(value *TfChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty) {
	if err := t.validatePutCodecSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCodecSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) PutRemixSettings(value *TfChannel_RemixSettingsProperty) {
	if err := t.validatePutRemixSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRemixSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) ResetAudioNormalizationSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioNormalizationSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) ResetAudioType() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) ResetAudioTypeControl() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioTypeControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) ResetAudioWatermarkSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioWatermarkSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) ResetCodecSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetCodecSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) ResetLanguageCode() {
	_jsii_.InvokeVoid(
		t,
		"resetLanguageCode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) ResetLanguageCodeControl() {
	_jsii_.InvokeVoid(
		t,
		"resetLanguageCodeControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) ResetRemixSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetRemixSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) ResetStreamName() {
	_jsii_.InvokeVoid(
		t,
		"resetStreamName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_AudioDescriptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

