package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Message() AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyList
	// Experimental.
	MessageInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Variation() AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyList
	// Experimental.
	VariationInput() interface{}
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
	PutMessage(value interface{})
	// Experimental.
	PutVariation(value interface{})
	// Experimental.
	ResetMessage()
	// Experimental.
	ResetVariation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference
type jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) Message() AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyList {
	var returns AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupMessagePropertyList
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) MessageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) Variation() AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyList {
	var returns AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyList
	_jsii_.Get(
		j,
		"variation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) VariationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsLexv2ModelsSlot.ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference_Override(a AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsLexv2ModelsSlot.ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) PutMessage(value interface{}) {
	if err := a.validatePutMessageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMessage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) PutVariation(value interface{}) {
	if err := a.validatePutVariationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVariation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) ResetMessage() {
	_jsii_.InvokeVoid(
		a,
		"resetMessage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) ResetVariation() {
	_jsii_.InvokeVoid(
		a,
		"resetVariation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLexv2ModelsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

