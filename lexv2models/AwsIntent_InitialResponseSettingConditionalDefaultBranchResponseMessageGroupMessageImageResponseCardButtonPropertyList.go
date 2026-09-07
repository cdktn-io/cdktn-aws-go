package lexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/lexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/lexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList interface {
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
	Get(index *float64) AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList
type jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList {
	_init_.Initialize()

	if err := validateNewAwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsIntent.InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList_Override(a AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsIntent.InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		a,
	)
}

func (j *jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (a *jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (a *jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) Get(index *float64) AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyOutputReference {
	if err := a.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyOutputReference

	_jsii_.Invoke(
		a,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

