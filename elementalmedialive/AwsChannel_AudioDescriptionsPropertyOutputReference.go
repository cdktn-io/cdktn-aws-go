package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_AudioDescriptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioNormalizationSettings() AwsChannel_AudioNormalizationSettingsPropertyOutputReference
	// Experimental.
	AudioNormalizationSettingsInput() *AwsChannel_AudioNormalizationSettingsProperty
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
	AudioWatermarkSettings() AwsChannel_AudioWatermarkSettingsPropertyOutputReference
	// Experimental.
	AudioWatermarkSettingsInput() *AwsChannel_AudioWatermarkSettingsProperty
	// Experimental.
	CodecSettings() AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference
	// Experimental.
	CodecSettingsInput() *AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty
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
	RemixSettings() AwsChannel_RemixSettingsPropertyOutputReference
	// Experimental.
	RemixSettingsInput() *AwsChannel_RemixSettingsProperty
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
	PutAudioNormalizationSettings(value *AwsChannel_AudioNormalizationSettingsProperty)
	// Experimental.
	PutAudioWatermarkSettings(value *AwsChannel_AudioWatermarkSettingsProperty)
	// Experimental.
	PutCodecSettings(value *AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty)
	// Experimental.
	PutRemixSettings(value *AwsChannel_RemixSettingsProperty)
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

// The jsii proxy struct for AwsChannel_AudioDescriptionsPropertyOutputReference
type jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) AudioNormalizationSettings() AwsChannel_AudioNormalizationSettingsPropertyOutputReference {
	var returns AwsChannel_AudioNormalizationSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"audioNormalizationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) AudioNormalizationSettingsInput() *AwsChannel_AudioNormalizationSettingsProperty {
	var returns *AwsChannel_AudioNormalizationSettingsProperty
	_jsii_.Get(
		j,
		"audioNormalizationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) AudioSelectorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioSelectorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) AudioSelectorNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioSelectorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) AudioType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) AudioTypeControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioTypeControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) AudioTypeControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioTypeControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) AudioTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) AudioWatermarkSettings() AwsChannel_AudioWatermarkSettingsPropertyOutputReference {
	var returns AwsChannel_AudioWatermarkSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"audioWatermarkSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) AudioWatermarkSettingsInput() *AwsChannel_AudioWatermarkSettingsProperty {
	var returns *AwsChannel_AudioWatermarkSettingsProperty
	_jsii_.Get(
		j,
		"audioWatermarkSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) CodecSettings() AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference {
	var returns AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"codecSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) CodecSettingsInput() *AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty {
	var returns *AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty
	_jsii_.Get(
		j,
		"codecSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) LanguageCodeControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) LanguageCodeControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) RemixSettings() AwsChannel_RemixSettingsPropertyOutputReference {
	var returns AwsChannel_RemixSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"remixSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) RemixSettingsInput() *AwsChannel_RemixSettingsProperty {
	var returns *AwsChannel_RemixSettingsProperty
	_jsii_.Get(
		j,
		"remixSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) StreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) StreamNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_AudioDescriptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsChannel_AudioDescriptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_AudioDescriptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.AudioDescriptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_AudioDescriptionsPropertyOutputReference_Override(a AwsChannel_AudioDescriptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.AudioDescriptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference)SetAudioSelectorName(val *string) {
	if err := j.validateSetAudioSelectorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioSelectorName",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference)SetAudioType(val *string) {
	if err := j.validateSetAudioTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioType",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference)SetAudioTypeControl(val *string) {
	if err := j.validateSetAudioTypeControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioTypeControl",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference)SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference)SetLanguageCodeControl(val *string) {
	if err := j.validateSetLanguageCodeControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCodeControl",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference)SetStreamName(val *string) {
	if err := j.validateSetStreamNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamName",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) PutAudioNormalizationSettings(value *AwsChannel_AudioNormalizationSettingsProperty) {
	if err := a.validatePutAudioNormalizationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAudioNormalizationSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) PutAudioWatermarkSettings(value *AwsChannel_AudioWatermarkSettingsProperty) {
	if err := a.validatePutAudioWatermarkSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAudioWatermarkSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) PutCodecSettings(value *AwsChannel_EncoderSettingsAudioDescriptionsCodecSettingsProperty) {
	if err := a.validatePutCodecSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodecSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) PutRemixSettings(value *AwsChannel_RemixSettingsProperty) {
	if err := a.validatePutRemixSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRemixSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) ResetAudioNormalizationSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioNormalizationSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) ResetAudioType() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) ResetAudioTypeControl() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioTypeControl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) ResetAudioWatermarkSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioWatermarkSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) ResetCodecSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetCodecSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) ResetLanguageCode() {
	_jsii_.InvokeVoid(
		a,
		"resetLanguageCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) ResetLanguageCodeControl() {
	_jsii_.InvokeVoid(
		a,
		"resetLanguageCodeControl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) ResetRemixSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRemixSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) ResetStreamName() {
	_jsii_.InvokeVoid(
		a,
		"resetStreamName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_AudioDescriptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

