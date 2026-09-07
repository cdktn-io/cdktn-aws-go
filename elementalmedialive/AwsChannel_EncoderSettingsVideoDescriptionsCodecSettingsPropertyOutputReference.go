package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference interface {
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
	FrameCaptureSettings() AwsChannel_FrameCaptureSettingsPropertyOutputReference
	// Experimental.
	FrameCaptureSettingsInput() *AwsChannel_FrameCaptureSettingsProperty
	// Experimental.
	H264Settings() AwsChannel_H264SettingsPropertyOutputReference
	// Experimental.
	H264SettingsInput() *AwsChannel_H264SettingsProperty
	// Experimental.
	H265Settings() AwsChannel_H265SettingsPropertyOutputReference
	// Experimental.
	H265SettingsInput() *AwsChannel_H265SettingsProperty
	// Experimental.
	InternalValue() *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty)
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
	PutFrameCaptureSettings(value *AwsChannel_FrameCaptureSettingsProperty)
	// Experimental.
	PutH264Settings(value *AwsChannel_H264SettingsProperty)
	// Experimental.
	PutH265Settings(value *AwsChannel_H265SettingsProperty)
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

// The jsii proxy struct for AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference
type jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) FrameCaptureSettings() AwsChannel_FrameCaptureSettingsPropertyOutputReference {
	var returns AwsChannel_FrameCaptureSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"frameCaptureSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) FrameCaptureSettingsInput() *AwsChannel_FrameCaptureSettingsProperty {
	var returns *AwsChannel_FrameCaptureSettingsProperty
	_jsii_.Get(
		j,
		"frameCaptureSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) H264Settings() AwsChannel_H264SettingsPropertyOutputReference {
	var returns AwsChannel_H264SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"h264Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) H264SettingsInput() *AwsChannel_H264SettingsProperty {
	var returns *AwsChannel_H264SettingsProperty
	_jsii_.Get(
		j,
		"h264SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) H265Settings() AwsChannel_H265SettingsPropertyOutputReference {
	var returns AwsChannel_H265SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"h265Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) H265SettingsInput() *AwsChannel_H265SettingsProperty {
	var returns *AwsChannel_H265SettingsProperty
	_jsii_.Get(
		j,
		"h265SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) InternalValue() *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty {
	var returns *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference_Override(a AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference)SetInternalValue(val *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) PutFrameCaptureSettings(value *AwsChannel_FrameCaptureSettingsProperty) {
	if err := a.validatePutFrameCaptureSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFrameCaptureSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) PutH264Settings(value *AwsChannel_H264SettingsProperty) {
	if err := a.validatePutH264SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putH264Settings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) PutH265Settings(value *AwsChannel_H265SettingsProperty) {
	if err := a.validatePutH265SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putH265Settings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) ResetFrameCaptureSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetFrameCaptureSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) ResetH264Settings() {
	_jsii_.InvokeVoid(
		a,
		"resetH264Settings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) ResetH265Settings() {
	_jsii_.InvokeVoid(
		a,
		"resetH265Settings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

