package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
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
	FrameCaptureSettings() TfChannel_FrameCaptureSettingsPropertyOutputReference
	// Experimental.
	FrameCaptureSettingsInput() *TfChannel_FrameCaptureSettingsProperty
	// Experimental.
	H264Settings() TfChannel_H264SettingsPropertyOutputReference
	// Experimental.
	H264SettingsInput() *TfChannel_H264SettingsProperty
	// Experimental.
	H265Settings() TfChannel_H265SettingsPropertyOutputReference
	// Experimental.
	H265SettingsInput() *TfChannel_H265SettingsProperty
	// Experimental.
	InternalValue() *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty)
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
	PutFrameCaptureSettings(value *TfChannel_FrameCaptureSettingsProperty)
	// Experimental.
	PutH264Settings(value *TfChannel_H264SettingsProperty)
	// Experimental.
	PutH265Settings(value *TfChannel_H265SettingsProperty)
	// Experimental.
	ResetFrameCaptureSettings()
	// Experimental.
	ResetH264Settings()
	// Experimental.
	ResetH265Settings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference
type jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) FrameCaptureSettings() TfChannel_FrameCaptureSettingsPropertyOutputReference {
	var returns TfChannel_FrameCaptureSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"frameCaptureSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) FrameCaptureSettingsInput() *TfChannel_FrameCaptureSettingsProperty {
	var returns *TfChannel_FrameCaptureSettingsProperty
	_jsii_.Get(
		j,
		"frameCaptureSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) H264Settings() TfChannel_H264SettingsPropertyOutputReference {
	var returns TfChannel_H264SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"h264Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) H264SettingsInput() *TfChannel_H264SettingsProperty {
	var returns *TfChannel_H264SettingsProperty
	_jsii_.Get(
		j,
		"h264SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) H265Settings() TfChannel_H265SettingsPropertyOutputReference {
	var returns TfChannel_H265SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"h265Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) H265SettingsInput() *TfChannel_H265SettingsProperty {
	var returns *TfChannel_H265SettingsProperty
	_jsii_.Get(
		j,
		"h265SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) InternalValue() *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty {
	var returns *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference_Override(t TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) PutFrameCaptureSettings(value *TfChannel_FrameCaptureSettingsProperty) {
	if err := t.validatePutFrameCaptureSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFrameCaptureSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) PutH264Settings(value *TfChannel_H264SettingsProperty) {
	if err := t.validatePutH264SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putH264Settings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) PutH265Settings(value *TfChannel_H265SettingsProperty) {
	if err := t.validatePutH265SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putH265Settings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) ResetFrameCaptureSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetFrameCaptureSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) ResetH264Settings() {
	_jsii_.InvokeVoid(
		t,
		"resetH264Settings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) ResetH265Settings() {
	_jsii_.InvokeVoid(
		t,
		"resetH265Settings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

