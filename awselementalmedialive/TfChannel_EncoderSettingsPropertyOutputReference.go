package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_EncoderSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioDescriptions() TfChannel_AudioDescriptionsPropertyList
	// Experimental.
	AudioDescriptionsInput() interface{}
	// Experimental.
	AvailBlanking() TfChannel_AvailBlankingPropertyOutputReference
	// Experimental.
	AvailBlankingInput() *TfChannel_AvailBlankingProperty
	// Experimental.
	CaptionDescriptions() TfChannel_CaptionDescriptionsPropertyList
	// Experimental.
	CaptionDescriptionsInput() interface{}
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
	GlobalConfiguration() TfChannel_GlobalConfigurationPropertyOutputReference
	// Experimental.
	GlobalConfigurationInput() *TfChannel_GlobalConfigurationProperty
	// Experimental.
	InternalValue() *TfChannel_EncoderSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_EncoderSettingsProperty)
	// Experimental.
	MotionGraphicsConfiguration() TfChannel_MotionGraphicsConfigurationPropertyOutputReference
	// Experimental.
	MotionGraphicsConfigurationInput() *TfChannel_MotionGraphicsConfigurationProperty
	// Experimental.
	NielsenConfiguration() TfChannel_NielsenConfigurationPropertyOutputReference
	// Experimental.
	NielsenConfigurationInput() *TfChannel_NielsenConfigurationProperty
	// Experimental.
	OutputGroups() TfChannel_OutputGroupsPropertyList
	// Experimental.
	OutputGroupsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimecodeConfig() TfChannel_TimecodeConfigPropertyOutputReference
	// Experimental.
	TimecodeConfigInput() *TfChannel_TimecodeConfigProperty
	// Experimental.
	VideoDescriptions() TfChannel_VideoDescriptionsPropertyList
	// Experimental.
	VideoDescriptionsInput() interface{}
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
	PutAudioDescriptions(value interface{})
	// Experimental.
	PutAvailBlanking(value *TfChannel_AvailBlankingProperty)
	// Experimental.
	PutCaptionDescriptions(value interface{})
	// Experimental.
	PutGlobalConfiguration(value *TfChannel_GlobalConfigurationProperty)
	// Experimental.
	PutMotionGraphicsConfiguration(value *TfChannel_MotionGraphicsConfigurationProperty)
	// Experimental.
	PutNielsenConfiguration(value *TfChannel_NielsenConfigurationProperty)
	// Experimental.
	PutOutputGroups(value interface{})
	// Experimental.
	PutTimecodeConfig(value *TfChannel_TimecodeConfigProperty)
	// Experimental.
	PutVideoDescriptions(value interface{})
	// Experimental.
	ResetAudioDescriptions()
	// Experimental.
	ResetAvailBlanking()
	// Experimental.
	ResetCaptionDescriptions()
	// Experimental.
	ResetGlobalConfiguration()
	// Experimental.
	ResetMotionGraphicsConfiguration()
	// Experimental.
	ResetNielsenConfiguration()
	// Experimental.
	ResetVideoDescriptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_EncoderSettingsPropertyOutputReference
type jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) AudioDescriptions() TfChannel_AudioDescriptionsPropertyList {
	var returns TfChannel_AudioDescriptionsPropertyList
	_jsii_.Get(
		j,
		"audioDescriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) AudioDescriptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"audioDescriptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) AvailBlanking() TfChannel_AvailBlankingPropertyOutputReference {
	var returns TfChannel_AvailBlankingPropertyOutputReference
	_jsii_.Get(
		j,
		"availBlanking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) AvailBlankingInput() *TfChannel_AvailBlankingProperty {
	var returns *TfChannel_AvailBlankingProperty
	_jsii_.Get(
		j,
		"availBlankingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) CaptionDescriptions() TfChannel_CaptionDescriptionsPropertyList {
	var returns TfChannel_CaptionDescriptionsPropertyList
	_jsii_.Get(
		j,
		"captionDescriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) CaptionDescriptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captionDescriptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) GlobalConfiguration() TfChannel_GlobalConfigurationPropertyOutputReference {
	var returns TfChannel_GlobalConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"globalConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) GlobalConfigurationInput() *TfChannel_GlobalConfigurationProperty {
	var returns *TfChannel_GlobalConfigurationProperty
	_jsii_.Get(
		j,
		"globalConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) InternalValue() *TfChannel_EncoderSettingsProperty {
	var returns *TfChannel_EncoderSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) MotionGraphicsConfiguration() TfChannel_MotionGraphicsConfigurationPropertyOutputReference {
	var returns TfChannel_MotionGraphicsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"motionGraphicsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) MotionGraphicsConfigurationInput() *TfChannel_MotionGraphicsConfigurationProperty {
	var returns *TfChannel_MotionGraphicsConfigurationProperty
	_jsii_.Get(
		j,
		"motionGraphicsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) NielsenConfiguration() TfChannel_NielsenConfigurationPropertyOutputReference {
	var returns TfChannel_NielsenConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"nielsenConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) NielsenConfigurationInput() *TfChannel_NielsenConfigurationProperty {
	var returns *TfChannel_NielsenConfigurationProperty
	_jsii_.Get(
		j,
		"nielsenConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) OutputGroups() TfChannel_OutputGroupsPropertyList {
	var returns TfChannel_OutputGroupsPropertyList
	_jsii_.Get(
		j,
		"outputGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) OutputGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) TimecodeConfig() TfChannel_TimecodeConfigPropertyOutputReference {
	var returns TfChannel_TimecodeConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"timecodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) TimecodeConfigInput() *TfChannel_TimecodeConfigProperty {
	var returns *TfChannel_TimecodeConfigProperty
	_jsii_.Get(
		j,
		"timecodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) VideoDescriptions() TfChannel_VideoDescriptionsPropertyList {
	var returns TfChannel_VideoDescriptionsPropertyList
	_jsii_.Get(
		j,
		"videoDescriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) VideoDescriptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"videoDescriptionsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_EncoderSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_EncoderSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_EncoderSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.EncoderSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_EncoderSettingsPropertyOutputReference_Override(t TfChannel_EncoderSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.EncoderSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_EncoderSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) PutAudioDescriptions(value interface{}) {
	if err := t.validatePutAudioDescriptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAudioDescriptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) PutAvailBlanking(value *TfChannel_AvailBlankingProperty) {
	if err := t.validatePutAvailBlankingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAvailBlanking",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) PutCaptionDescriptions(value interface{}) {
	if err := t.validatePutCaptionDescriptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCaptionDescriptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) PutGlobalConfiguration(value *TfChannel_GlobalConfigurationProperty) {
	if err := t.validatePutGlobalConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGlobalConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) PutMotionGraphicsConfiguration(value *TfChannel_MotionGraphicsConfigurationProperty) {
	if err := t.validatePutMotionGraphicsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMotionGraphicsConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) PutNielsenConfiguration(value *TfChannel_NielsenConfigurationProperty) {
	if err := t.validatePutNielsenConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNielsenConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) PutOutputGroups(value interface{}) {
	if err := t.validatePutOutputGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOutputGroups",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) PutTimecodeConfig(value *TfChannel_TimecodeConfigProperty) {
	if err := t.validatePutTimecodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimecodeConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) PutVideoDescriptions(value interface{}) {
	if err := t.validatePutVideoDescriptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVideoDescriptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) ResetAudioDescriptions() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioDescriptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) ResetAvailBlanking() {
	_jsii_.InvokeVoid(
		t,
		"resetAvailBlanking",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) ResetCaptionDescriptions() {
	_jsii_.InvokeVoid(
		t,
		"resetCaptionDescriptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) ResetGlobalConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetGlobalConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) ResetMotionGraphicsConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetMotionGraphicsConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) ResetNielsenConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetNielsenConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) ResetVideoDescriptions() {
	_jsii_.InvokeVoid(
		t,
		"resetVideoDescriptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

