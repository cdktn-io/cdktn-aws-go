package lexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/lexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/lexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList interface {
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
	Get(index *float64) AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList
type jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList {
	_init_.Initialize()

	if err := validateNewAwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsIntent.ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList_Override(a AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsIntent.ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList) Get(index *float64) AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsIntent_ConfirmationSettingConfirmationConditionalDefaultBranchResponseMessageGroupVariationPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

