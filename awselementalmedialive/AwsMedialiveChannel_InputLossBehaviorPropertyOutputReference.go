package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BlackFrameMsec() *float64
	// Experimental.
	SetBlackFrameMsec(val *float64)
	// Experimental.
	BlackFrameMsecInput() *float64
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
	InputLossImageColor() *string
	// Experimental.
	SetInputLossImageColor(val *string)
	// Experimental.
	InputLossImageColorInput() *string
	// Experimental.
	InputLossImageSlate() AwsMedialiveChannel_InputLossImageSlatePropertyOutputReference
	// Experimental.
	InputLossImageSlateInput() *AwsMedialiveChannel_InputLossImageSlateProperty
	// Experimental.
	InputLossImageType() *string
	// Experimental.
	SetInputLossImageType(val *string)
	// Experimental.
	InputLossImageTypeInput() *string
	// Experimental.
	InternalValue() *AwsMedialiveChannel_InputLossBehaviorProperty
	// Experimental.
	SetInternalValue(val *AwsMedialiveChannel_InputLossBehaviorProperty)
	// Experimental.
	RepeatFrameMsec() *float64
	// Experimental.
	SetRepeatFrameMsec(val *float64)
	// Experimental.
	RepeatFrameMsecInput() *float64
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
	PutInputLossImageSlate(value *AwsMedialiveChannel_InputLossImageSlateProperty)
	// Experimental.
	ResetBlackFrameMsec()
	// Experimental.
	ResetInputLossImageColor()
	// Experimental.
	ResetInputLossImageSlate()
	// Experimental.
	ResetInputLossImageType()
	// Experimental.
	ResetRepeatFrameMsec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference
type jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) BlackFrameMsec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blackFrameMsec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) BlackFrameMsecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blackFrameMsecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) InputLossImageColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossImageColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) InputLossImageColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossImageColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) InputLossImageSlate() AwsMedialiveChannel_InputLossImageSlatePropertyOutputReference {
	var returns AwsMedialiveChannel_InputLossImageSlatePropertyOutputReference
	_jsii_.Get(
		j,
		"inputLossImageSlate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) InputLossImageSlateInput() *AwsMedialiveChannel_InputLossImageSlateProperty {
	var returns *AwsMedialiveChannel_InputLossImageSlateProperty
	_jsii_.Get(
		j,
		"inputLossImageSlateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) InputLossImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossImageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) InputLossImageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossImageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) InternalValue() *AwsMedialiveChannel_InputLossBehaviorProperty {
	var returns *AwsMedialiveChannel_InputLossBehaviorProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) RepeatFrameMsec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"repeatFrameMsec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) RepeatFrameMsecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"repeatFrameMsecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMedialiveChannel_InputLossBehaviorPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMedialiveChannel_InputLossBehaviorPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.InputLossBehaviorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMedialiveChannel_InputLossBehaviorPropertyOutputReference_Override(a AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.InputLossBehaviorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference)SetBlackFrameMsec(val *float64) {
	if err := j.validateSetBlackFrameMsecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blackFrameMsec",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference)SetInputLossImageColor(val *string) {
	if err := j.validateSetInputLossImageColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputLossImageColor",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference)SetInputLossImageType(val *string) {
	if err := j.validateSetInputLossImageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputLossImageType",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference)SetInternalValue(val *AwsMedialiveChannel_InputLossBehaviorProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference)SetRepeatFrameMsec(val *float64) {
	if err := j.validateSetRepeatFrameMsecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repeatFrameMsec",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) PutInputLossImageSlate(value *AwsMedialiveChannel_InputLossImageSlateProperty) {
	if err := a.validatePutInputLossImageSlateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInputLossImageSlate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) ResetBlackFrameMsec() {
	_jsii_.InvokeVoid(
		a,
		"resetBlackFrameMsec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) ResetInputLossImageColor() {
	_jsii_.InvokeVoid(
		a,
		"resetInputLossImageColor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) ResetInputLossImageSlate() {
	_jsii_.InvokeVoid(
		a,
		"resetInputLossImageSlate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) ResetInputLossImageType() {
	_jsii_.InvokeVoid(
		a,
		"resetInputLossImageType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) ResetRepeatFrameMsec() {
	_jsii_.InvokeVoid(
		a,
		"resetRepeatFrameMsec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_InputLossBehaviorPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

