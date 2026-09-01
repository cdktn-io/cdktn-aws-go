package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList interface {
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
	Get(index *float64) AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList
type jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList {
	_init_.Initialize()

	if err := validateNewAwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsLexv2ModelsIntent.FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList_Override(a AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsLexv2ModelsIntent.FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := a.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		a,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList) Get(index *float64) AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLexv2ModelsIntent_FulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroupVariationPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

