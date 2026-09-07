package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_EncoderSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioDescriptions() AwsChannel_AudioDescriptionsPropertyList
	// Experimental.
	AudioDescriptionsInput() interface{}
	// Experimental.
	AvailBlanking() AwsChannel_AvailBlankingPropertyOutputReference
	// Experimental.
	AvailBlankingInput() *AwsChannel_AvailBlankingProperty
	// Experimental.
	CaptionDescriptions() AwsChannel_CaptionDescriptionsPropertyList
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
	GlobalConfiguration() AwsChannel_GlobalConfigurationPropertyOutputReference
	// Experimental.
	GlobalConfigurationInput() *AwsChannel_GlobalConfigurationProperty
	// Experimental.
	InternalValue() *AwsChannel_EncoderSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_EncoderSettingsProperty)
	// Experimental.
	MotionGraphicsConfiguration() AwsChannel_MotionGraphicsConfigurationPropertyOutputReference
	// Experimental.
	MotionGraphicsConfigurationInput() *AwsChannel_MotionGraphicsConfigurationProperty
	// Experimental.
	NielsenConfiguration() AwsChannel_NielsenConfigurationPropertyOutputReference
	// Experimental.
	NielsenConfigurationInput() *AwsChannel_NielsenConfigurationProperty
	// Experimental.
	OutputGroups() AwsChannel_OutputGroupsPropertyList
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
	TimecodeConfig() AwsChannel_TimecodeConfigPropertyOutputReference
	// Experimental.
	TimecodeConfigInput() *AwsChannel_TimecodeConfigProperty
	// Experimental.
	VideoDescriptions() AwsChannel_VideoDescriptionsPropertyList
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
	PutAvailBlanking(value *AwsChannel_AvailBlankingProperty)
	// Experimental.
	PutCaptionDescriptions(value interface{})
	// Experimental.
	PutGlobalConfiguration(value *AwsChannel_GlobalConfigurationProperty)
	// Experimental.
	PutMotionGraphicsConfiguration(value *AwsChannel_MotionGraphicsConfigurationProperty)
	// Experimental.
	PutNielsenConfiguration(value *AwsChannel_NielsenConfigurationProperty)
	// Experimental.
	PutOutputGroups(value interface{})
	// Experimental.
	PutTimecodeConfig(value *AwsChannel_TimecodeConfigProperty)
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

// The jsii proxy struct for AwsChannel_EncoderSettingsPropertyOutputReference
type jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) AudioDescriptions() AwsChannel_AudioDescriptionsPropertyList {
	var returns AwsChannel_AudioDescriptionsPropertyList
	_jsii_.Get(
		j,
		"audioDescriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) AudioDescriptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"audioDescriptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) AvailBlanking() AwsChannel_AvailBlankingPropertyOutputReference {
	var returns AwsChannel_AvailBlankingPropertyOutputReference
	_jsii_.Get(
		j,
		"availBlanking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) AvailBlankingInput() *AwsChannel_AvailBlankingProperty {
	var returns *AwsChannel_AvailBlankingProperty
	_jsii_.Get(
		j,
		"availBlankingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) CaptionDescriptions() AwsChannel_CaptionDescriptionsPropertyList {
	var returns AwsChannel_CaptionDescriptionsPropertyList
	_jsii_.Get(
		j,
		"captionDescriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) CaptionDescriptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captionDescriptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) GlobalConfiguration() AwsChannel_GlobalConfigurationPropertyOutputReference {
	var returns AwsChannel_GlobalConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"globalConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) GlobalConfigurationInput() *AwsChannel_GlobalConfigurationProperty {
	var returns *AwsChannel_GlobalConfigurationProperty
	_jsii_.Get(
		j,
		"globalConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) InternalValue() *AwsChannel_EncoderSettingsProperty {
	var returns *AwsChannel_EncoderSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) MotionGraphicsConfiguration() AwsChannel_MotionGraphicsConfigurationPropertyOutputReference {
	var returns AwsChannel_MotionGraphicsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"motionGraphicsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) MotionGraphicsConfigurationInput() *AwsChannel_MotionGraphicsConfigurationProperty {
	var returns *AwsChannel_MotionGraphicsConfigurationProperty
	_jsii_.Get(
		j,
		"motionGraphicsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) NielsenConfiguration() AwsChannel_NielsenConfigurationPropertyOutputReference {
	var returns AwsChannel_NielsenConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"nielsenConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) NielsenConfigurationInput() *AwsChannel_NielsenConfigurationProperty {
	var returns *AwsChannel_NielsenConfigurationProperty
	_jsii_.Get(
		j,
		"nielsenConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) OutputGroups() AwsChannel_OutputGroupsPropertyList {
	var returns AwsChannel_OutputGroupsPropertyList
	_jsii_.Get(
		j,
		"outputGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) OutputGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) TimecodeConfig() AwsChannel_TimecodeConfigPropertyOutputReference {
	var returns AwsChannel_TimecodeConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"timecodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) TimecodeConfigInput() *AwsChannel_TimecodeConfigProperty {
	var returns *AwsChannel_TimecodeConfigProperty
	_jsii_.Get(
		j,
		"timecodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) VideoDescriptions() AwsChannel_VideoDescriptionsPropertyList {
	var returns AwsChannel_VideoDescriptionsPropertyList
	_jsii_.Get(
		j,
		"videoDescriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) VideoDescriptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"videoDescriptionsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_EncoderSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_EncoderSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_EncoderSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.EncoderSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_EncoderSettingsPropertyOutputReference_Override(a AwsChannel_EncoderSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.EncoderSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference)SetInternalValue(val *AwsChannel_EncoderSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) PutAudioDescriptions(value interface{}) {
	if err := a.validatePutAudioDescriptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAudioDescriptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) PutAvailBlanking(value *AwsChannel_AvailBlankingProperty) {
	if err := a.validatePutAvailBlankingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAvailBlanking",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) PutCaptionDescriptions(value interface{}) {
	if err := a.validatePutCaptionDescriptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCaptionDescriptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) PutGlobalConfiguration(value *AwsChannel_GlobalConfigurationProperty) {
	if err := a.validatePutGlobalConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGlobalConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) PutMotionGraphicsConfiguration(value *AwsChannel_MotionGraphicsConfigurationProperty) {
	if err := a.validatePutMotionGraphicsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMotionGraphicsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) PutNielsenConfiguration(value *AwsChannel_NielsenConfigurationProperty) {
	if err := a.validatePutNielsenConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNielsenConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) PutOutputGroups(value interface{}) {
	if err := a.validatePutOutputGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutputGroups",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) PutTimecodeConfig(value *AwsChannel_TimecodeConfigProperty) {
	if err := a.validatePutTimecodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimecodeConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) PutVideoDescriptions(value interface{}) {
	if err := a.validatePutVideoDescriptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVideoDescriptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) ResetAudioDescriptions() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioDescriptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) ResetAvailBlanking() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailBlanking",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) ResetCaptionDescriptions() {
	_jsii_.InvokeVoid(
		a,
		"resetCaptionDescriptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) ResetGlobalConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetGlobalConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) ResetMotionGraphicsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetMotionGraphicsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) ResetNielsenConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetNielsenConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) ResetVideoDescriptions() {
	_jsii_.InvokeVoid(
		a,
		"resetVideoDescriptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

