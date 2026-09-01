package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioGroupId() *string
	// Experimental.
	SetAudioGroupId(val *string)
	// Experimental.
	AudioGroupIdInput() *string
	// Experimental.
	AudioOnlyImage() AwsMedialiveChannel_AudioOnlyImagePropertyOutputReference
	// Experimental.
	AudioOnlyImageInput() *AwsMedialiveChannel_AudioOnlyImageProperty
	// Experimental.
	AudioTrackType() *string
	// Experimental.
	SetAudioTrackType(val *string)
	// Experimental.
	AudioTrackTypeInput() *string
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
	InternalValue() *AwsMedialiveChannel_AudioOnlyHlsSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsMedialiveChannel_AudioOnlyHlsSettingsProperty)
	// Experimental.
	SegmentType() *string
	// Experimental.
	SetSegmentType(val *string)
	// Experimental.
	SegmentTypeInput() *string
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
	PutAudioOnlyImage(value *AwsMedialiveChannel_AudioOnlyImageProperty)
	// Experimental.
	ResetAudioGroupId()
	// Experimental.
	ResetAudioOnlyImage()
	// Experimental.
	ResetAudioTrackType()
	// Experimental.
	ResetSegmentType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference
type jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) AudioGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) AudioGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) AudioOnlyImage() AwsMedialiveChannel_AudioOnlyImagePropertyOutputReference {
	var returns AwsMedialiveChannel_AudioOnlyImagePropertyOutputReference
	_jsii_.Get(
		j,
		"audioOnlyImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) AudioOnlyImageInput() *AwsMedialiveChannel_AudioOnlyImageProperty {
	var returns *AwsMedialiveChannel_AudioOnlyImageProperty
	_jsii_.Get(
		j,
		"audioOnlyImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) AudioTrackType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioTrackType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) AudioTrackTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioTrackTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) InternalValue() *AwsMedialiveChannel_AudioOnlyHlsSettingsProperty {
	var returns *AwsMedialiveChannel_AudioOnlyHlsSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) SegmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) SegmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.AudioOnlyHlsSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference_Override(a AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.AudioOnlyHlsSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference)SetAudioGroupId(val *string) {
	if err := j.validateSetAudioGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioGroupId",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference)SetAudioTrackType(val *string) {
	if err := j.validateSetAudioTrackTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioTrackType",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference)SetInternalValue(val *AwsMedialiveChannel_AudioOnlyHlsSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference)SetSegmentType(val *string) {
	if err := j.validateSetSegmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentType",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) PutAudioOnlyImage(value *AwsMedialiveChannel_AudioOnlyImageProperty) {
	if err := a.validatePutAudioOnlyImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAudioOnlyImage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) ResetAudioGroupId() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioGroupId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) ResetAudioOnlyImage() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioOnlyImage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) ResetAudioTrackType() {
	_jsii_.InvokeVoid(
		a,
		"resetAudioTrackType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) ResetSegmentType() {
	_jsii_.InvokeVoid(
		a,
		"resetSegmentType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_AudioOnlyHlsSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

