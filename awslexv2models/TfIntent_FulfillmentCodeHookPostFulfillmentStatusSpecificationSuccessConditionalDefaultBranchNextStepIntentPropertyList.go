package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList interface {
	cdktn.ComplexList
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
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	Get(index *float64) TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList
type jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList {
	_init_.Initialize()

	if err := validateNewTfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfIntent.FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList_Override(t TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfIntent.FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		t,
	)
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := t.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		t,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList) Get(index *float64) TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyOutputReference {
	if err := t.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyOutputReference

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalDefaultBranchNextStepIntentPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

