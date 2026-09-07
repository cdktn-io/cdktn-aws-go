package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_VideoDescriptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CodecSettings() AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference
	// Experimental.
	CodecSettingsInput() *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty
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
	PutCodecSettings(value *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty)
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

// The jsii proxy struct for AwsChannel_VideoDescriptionsPropertyOutputReference
type jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) CodecSettings() AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference {
	var returns AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"codecSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) CodecSettingsInput() *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty {
	var returns *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty
	_jsii_.Get(
		j,
		"codecSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) HeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"heightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) RespondToAfd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"respondToAfd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) RespondToAfdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"respondToAfdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) ScalingBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) ScalingBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) Sharpness() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharpness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) SharpnessInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sharpnessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) WidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"widthInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_VideoDescriptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsChannel_VideoDescriptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_VideoDescriptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.VideoDescriptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_VideoDescriptionsPropertyOutputReference_Override(a AwsChannel_VideoDescriptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.VideoDescriptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference)SetHeight(val *float64) {
	if err := j.validateSetHeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"height",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference)SetRespondToAfd(val *string) {
	if err := j.validateSetRespondToAfdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"respondToAfd",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference)SetScalingBehavior(val *string) {
	if err := j.validateSetScalingBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference)SetSharpness(val *float64) {
	if err := j.validateSetSharpnessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharpness",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference)SetWidth(val *float64) {
	if err := j.validateSetWidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"width",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) PutCodecSettings(value *AwsChannel_EncoderSettingsVideoDescriptionsCodecSettingsProperty) {
	if err := a.validatePutCodecSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodecSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) ResetCodecSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetCodecSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) ResetHeight() {
	_jsii_.InvokeVoid(
		a,
		"resetHeight",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) ResetRespondToAfd() {
	_jsii_.InvokeVoid(
		a,
		"resetRespondToAfd",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) ResetScalingBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetScalingBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) ResetSharpness() {
	_jsii_.InvokeVoid(
		a,
		"resetSharpness",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) ResetWidth() {
	_jsii_.InvokeVoid(
		a,
		"resetWidth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_VideoDescriptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

