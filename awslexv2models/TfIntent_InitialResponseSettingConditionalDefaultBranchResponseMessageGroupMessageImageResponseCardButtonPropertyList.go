package awslexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexv2models/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList interface {
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
	Get(index *float64) TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList
type jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList {
	_init_.Initialize()

	if err := validateNewTfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfIntent.InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList_Override(t TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.TfIntent.InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		t,
	)
}

func (j *jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (t *jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (t *jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) Get(index *float64) TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyOutputReference {
	if err := t.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyOutputReference

	_jsii_.Invoke(
		t,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfIntent_InitialResponseSettingConditionalDefaultBranchResponseMessageGroupMessageImageResponseCardButtonPropertyList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

