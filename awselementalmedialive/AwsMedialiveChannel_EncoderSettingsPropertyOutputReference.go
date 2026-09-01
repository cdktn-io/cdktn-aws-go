package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMedialiveChannel_EncoderSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioDescriptions() AwsMedialiveChannel_AudioDescriptionsPropertyList
	// Experimental.
	AudioDescriptionsInput() interface{}
	// Experimental.
	AvailBlanking() AwsMedialiveChannel_AvailBlankingPropertyOutputReference
	// Experimental.
	AvailBlankingInput() *AwsMedialiveChannel_AvailBlankingProperty
	// Experimental.
	CaptionDescriptions() AwsMedialiveChannel_CaptionDescriptionsPropertyList
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
	GlobalConfiguration() AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference
	// Experimental.
	GlobalConfigurationInput() *AwsMedialiveChannel_GlobalConfigurationProperty
	// Experimental.
	InternalValue() *AwsMedialiveChannel_EncoderSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsMedialiveChannel_EncoderSettingsProperty)
	// Experimental.
	MotionGraphicsConfiguration() AwsMedialiveChannel_MotionGraphicsConfigurationPropertyOutputReference
	// Experimental.
	MotionGraphicsConfigurationInput() *AwsMedialiveChannel_MotionGraphicsConfigurationProperty
	// Experimental.
	NielsenConfiguration() AwsMedialiveChannel_NielsenConfigurationPropertyOutputReference
	// Experimental.
	NielsenConfigurationInput() *AwsMedialiveChannel_NielsenConfigurationProperty
	// Experimental.
	OutputGroups() AwsMedialiveChannel_OutputGroupsPropertyList
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
	TimecodeConfig() AwsMedialiveChannel_TimecodeConfigPropertyOutputReference
	// Experimental.
	TimecodeConfigInput() *AwsMedialiveChannel_TimecodeConfigProperty
	// Experimental.
	VideoDescriptions() AwsMedialiveChannel_VideoDescriptionsPropertyList
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
	PutAvailBlanking(value *AwsMedialiveChannel_AvailBlankingProperty)
	// Experimental.
	PutCaptionDescriptions(value interface{})
	// Experimental.
	PutGlobalConfiguration(value *AwsMedialiveChannel_GlobalConfigurationProperty)
	// Experimental.
	PutMotionGraphicsConfiguration(value *AwsMedialiveChannel_MotionGraphicsConfigurationProperty)
	// Experimental.
	PutNielsenConfiguration(value *AwsMedialiveChannel_NielsenConfigurationProperty)
	// Experimental.
	PutOutputGroups(value interface{})
	// Experimental.
	PutTimecodeConfig(value *AwsMedialiveChannel_TimecodeConfigProperty)
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

// The jsii proxy struct for AwsMedialiveChannel_EncoderSettingsPropertyOutputReference
type jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) AudioDescriptions() AwsMedialiveChannel_AudioDescriptionsPropertyList {
	var returns AwsMedialiveChannel_AudioDescriptionsPropertyList
	_jsii_.Get(
		j,
		"audioDescriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) AudioDescriptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"audioDescriptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) AvailBlanking() AwsMedialiveChannel_AvailBlankingPropertyOutputReference {
	var returns AwsMedialiveChannel_AvailBlankingPropertyOutputReference
	_jsii_.Get(
		j,
		"availBlanking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) AvailBlankingInput() *AwsMedialiveChannel_AvailBlankingProperty {
	var returns *AwsMedialiveChannel_AvailBlankingProperty
	_jsii_.Get(
		j,
		"availBlankingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) CaptionDescriptions() AwsMedialiveChannel_CaptionDescriptionsPropertyList {
	var returns AwsMedialiveChannel_CaptionDescriptionsPropertyList
	_jsii_.Get(
		j,
		"captionDescriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) CaptionDescriptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captionDescriptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) GlobalConfiguration() AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference {
	var returns AwsMedialiveChannel_GlobalConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"globalConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) GlobalConfigurationInput() *AwsMedialiveChannel_GlobalConfigurationProperty {
	var returns *AwsMedialiveChannel_GlobalConfigurationProperty
	_jsii_.Get(
		j,
		"globalConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) InternalValue() *AwsMedialiveChannel_EncoderSettingsProperty {
	var returns *AwsMedialiveChannel_EncoderSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) MotionGraphicsConfiguration() AwsMedialiveChannel_MotionGraphicsConfigurationPropertyOutputReference {
	var returns AwsMedialiveChannel_MotionGraphicsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"motionGraphicsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) MotionGraphicsConfigurationInput() *AwsMedialiveChannel_MotionGraphicsConfigurationProperty {
	var returns *AwsMedialiveChannel_MotionGraphicsConfigurationProperty
	_jsii_.Get(
		j,
		"motionGraphicsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) NielsenConfiguration() AwsMedialiveChannel_NielsenConfigurationPropertyOutputReference {
	var returns AwsMedialiveChannel_NielsenConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"nielsenConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) NielsenConfigurationInput() *AwsMedialiveChannel_NielsenConfigurationProperty {
	var returns *AwsMedialiveChannel_NielsenConfigurationProperty
	_jsii_.Get(
		j,
		"nielsenConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) OutputGroups() AwsMedialiveChannel_OutputGroupsPropertyList {
	var returns AwsMedialiveChannel_OutputGroupsPropertyList
	_jsii_.Get(
		j,
		"outputGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) OutputGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) TimecodeConfig() AwsMedialiveChannel_TimecodeConfigPropertyOutputReference {
	var returns AwsMedialiveChannel_TimecodeConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"timecodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) TimecodeConfigInput() *AwsMedialiveChannel_TimecodeConfigProperty {
	var returns *AwsMedialiveChannel_TimecodeConfigProperty
	_jsii_.Get(
		j,
		"timecodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) VideoDescriptions() AwsMedialiveChannel_VideoDescriptionsPropertyList {
	var returns AwsMedialiveChannel_VideoDescriptionsPropertyList
	_jsii_.Get(
		j,
		"videoDescriptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) VideoDescriptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"videoDescriptionsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMedialiveChannel_EncoderSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMedialiveChannel_EncoderSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMedialiveChannel_EncoderSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.EncoderSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMedialiveChannel_EncoderSettingsPropertyOutputReference_Override(a AwsMedialiveChannel_EncoderSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.EncoderSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference)SetInternalValue(val *AwsMedialiveChannel_EncoderSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) PutAudioDescriptions(value interface{}) {
	if err := a.validatePutAudioDescriptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAudioDescriptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) PutAvailBlanking(value *AwsMedialiveChannel_AvailBlankingProperty) {
	if err := a.validatePutAvailBlankingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAvailBlanking",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) PutCaptionDescriptions(value interface{}) {
	if err := a.validatePutCaptionDescriptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCaptionDescriptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) PutGlobalConfiguration(value *AwsMedialiveChannel_GlobalConfigurationProperty) {
	if err := a.validatePutGlobalConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGlobalConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) PutMotionGraphicsConfiguration(value *AwsMedialiveChannel_MotionGraphicsConfigurationProperty) {
	if err := a.validatePutMotionGraphicsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMotionGraphicsConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) PutNielsenConfiguration(value *AwsMedialiveChannel_NielsenConfigurationProperty) {
	if err := a.validatePutNielsenConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNielsenConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) PutOutputGroups(value interface{}) {
	if err := a.validatePutOutputGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutputGroups",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) PutTimecodeConfig(value *AwsMedialiveChannel_TimecodeConfigProperty) {
	if err := a.validatePutTimecodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimecodeConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) PutVideoDescriptions(value interface{}) {
	if err := a.validatePutVideoDescriptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVideoDescriptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) ResetAudioDescriptions() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioDescriptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) ResetAvailBlanking() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailBlanking",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) ResetCaptionDescriptions() {
	_jsii_.InvokeVoid(
		a,
		"resetCaptionDescriptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) ResetGlobalConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetGlobalConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) ResetMotionGraphicsConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetMotionGraphicsConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) ResetNielsenConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetNielsenConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) ResetVideoDescriptions() {
	_jsii_.InvokeVoid(
		a,
		"resetVideoDescriptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_EncoderSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

