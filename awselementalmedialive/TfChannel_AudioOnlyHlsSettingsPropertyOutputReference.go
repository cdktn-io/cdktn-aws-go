package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_AudioOnlyHlsSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AudioGroupId() *string
	// Experimental.
	SetAudioGroupId(val *string)
	// Experimental.
	AudioGroupIdInput() *string
	// Experimental.
	AudioOnlyImage() TfChannel_AudioOnlyImagePropertyOutputReference
	// Experimental.
	AudioOnlyImageInput() *TfChannel_AudioOnlyImageProperty
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
	InternalValue() *TfChannel_AudioOnlyHlsSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_AudioOnlyHlsSettingsProperty)
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
	PutAudioOnlyImage(value *TfChannel_AudioOnlyImageProperty)
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

// The jsii proxy struct for TfChannel_AudioOnlyHlsSettingsPropertyOutputReference
type jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) AudioGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) AudioGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) AudioOnlyImage() TfChannel_AudioOnlyImagePropertyOutputReference {
	var returns TfChannel_AudioOnlyImagePropertyOutputReference
	_jsii_.Get(
		j,
		"audioOnlyImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) AudioOnlyImageInput() *TfChannel_AudioOnlyImageProperty {
	var returns *TfChannel_AudioOnlyImageProperty
	_jsii_.Get(
		j,
		"audioOnlyImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) AudioTrackType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioTrackType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) AudioTrackTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioTrackTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) InternalValue() *TfChannel_AudioOnlyHlsSettingsProperty {
	var returns *TfChannel_AudioOnlyHlsSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) SegmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) SegmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_AudioOnlyHlsSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_AudioOnlyHlsSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_AudioOnlyHlsSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.AudioOnlyHlsSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_AudioOnlyHlsSettingsPropertyOutputReference_Override(t TfChannel_AudioOnlyHlsSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.AudioOnlyHlsSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference)SetAudioGroupId(val *string) {
	if err := j.validateSetAudioGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioGroupId",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference)SetAudioTrackType(val *string) {
	if err := j.validateSetAudioTrackTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioTrackType",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_AudioOnlyHlsSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference)SetSegmentType(val *string) {
	if err := j.validateSetSegmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segmentType",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) PutAudioOnlyImage(value *TfChannel_AudioOnlyImageProperty) {
	if err := t.validatePutAudioOnlyImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAudioOnlyImage",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) ResetAudioGroupId() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioGroupId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) ResetAudioOnlyImage() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioOnlyImage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) ResetAudioTrackType() {
	_jsii_.InvokeVoid(
		t,
		"resetAudioTrackType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) ResetSegmentType() {
	_jsii_.InvokeVoid(
		t,
		"resetSegmentType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_AudioOnlyHlsSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

