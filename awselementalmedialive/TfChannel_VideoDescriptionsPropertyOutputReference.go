package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_VideoDescriptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CodecSettings() TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference
	// Experimental.
	CodecSettingsInput() *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty
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
	Height() *float64
	// Experimental.
	SetHeight(val *float64)
	// Experimental.
	HeightInput() *float64
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	RespondToAfd() *string
	// Experimental.
	SetRespondToAfd(val *string)
	// Experimental.
	RespondToAfdInput() *string
	// Experimental.
	ScalingBehavior() *string
	// Experimental.
	SetScalingBehavior(val *string)
	// Experimental.
	ScalingBehaviorInput() *string
	// Experimental.
	Sharpness() *float64
	// Experimental.
	SetSharpness(val *float64)
	// Experimental.
	SharpnessInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Width() *float64
	// Experimental.
	SetWidth(val *float64)
	// Experimental.
	WidthInput() *float64
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
	PutCodecSettings(value *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty)
	// Experimental.
	ResetCodecSettings()
	// Experimental.
	ResetHeight()
	// Experimental.
	ResetRespondToAfd()
	// Experimental.
	ResetScalingBehavior()
	// Experimental.
	ResetSharpness()
	// Experimental.
	ResetWidth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_VideoDescriptionsPropertyOutputReference
type jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) CodecSettings() TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference {
	var returns TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"codecSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) CodecSettingsInput() *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty {
	var returns *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty
	_jsii_.Get(
		j,
		"codecSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) HeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) RespondToAfd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"respondToAfd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) RespondToAfdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"respondToAfdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) ScalingBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) ScalingBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) Sharpness() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharpness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) SharpnessInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharpnessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) WidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widthInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_VideoDescriptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfChannel_VideoDescriptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_VideoDescriptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.VideoDescriptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_VideoDescriptionsPropertyOutputReference_Override(t TfChannel_VideoDescriptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.VideoDescriptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference)SetHeight(val *float64) {
	if err := j.validateSetHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"height",
		val,
	)
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference)SetRespondToAfd(val *string) {
	if err := j.validateSetRespondToAfdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"respondToAfd",
		val,
	)
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference)SetScalingBehavior(val *string) {
	if err := j.validateSetScalingBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingBehavior",
		val,
	)
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference)SetSharpness(val *float64) {
	if err := j.validateSetSharpnessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharpness",
		val,
	)
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference)SetWidth(val *float64) {
	if err := j.validateSetWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"width",
		val,
	)
}

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) PutCodecSettings(value *TfChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty) {
	if err := t.validatePutCodecSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCodecSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) ResetCodecSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetCodecSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) ResetHeight() {
	_jsii_.InvokeVoid(
		t,
		"resetHeight",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) ResetRespondToAfd() {
	_jsii_.InvokeVoid(
		t,
		"resetRespondToAfd",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) ResetScalingBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetScalingBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) ResetSharpness() {
	_jsii_.InvokeVoid(
		t,
		"resetSharpness",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) ResetWidth() {
	_jsii_.InvokeVoid(
		t,
		"resetWidth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_VideoDescriptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

