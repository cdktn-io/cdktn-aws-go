package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference interface {
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
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	Slot() TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentSlotPropertyList
	// Experimental.
	SlotInput() interface{}
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
	PutSlot(value interface{})
	// Experimental.
	ResetName()
	// Experimental.
	ResetSlot()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference
type jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) Slot() TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentSlotPropertyList {
	var returns TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentSlotPropertyList
	_jsii_.Get(
		j,
		"slot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) SlotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfIntent.FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference_Override(t TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfIntent.FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) PutSlot(value interface{}) {
	if err := t.validatePutSlotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSlot",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) ResetSlot() {
	_jsii_.InvokeVoid(
		t,
		"resetSlot",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalConditionalBranchNextStepIntentPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

