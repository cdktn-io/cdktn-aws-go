package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataSet_JoinInstructionPropertyOutputReference interface {
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
	InternalValue() *TfDataSet_JoinInstructionProperty
	// Experimental.
	SetInternalValue(val *TfDataSet_JoinInstructionProperty)
	// Experimental.
	LeftJoinKeyProperties() TfDataSet_LeftJoinKeyPropertiesPropertyOutputReference
	// Experimental.
	LeftJoinKeyPropertiesInput() *TfDataSet_LeftJoinKeyPropertiesProperty
	// Experimental.
	LeftOperand() *string
	// Experimental.
	SetLeftOperand(val *string)
	// Experimental.
	LeftOperandInput() *string
	// Experimental.
	OnClause() *string
	// Experimental.
	SetOnClause(val *string)
	// Experimental.
	OnClauseInput() *string
	// Experimental.
	RightJoinKeyProperties() TfDataSet_RightJoinKeyPropertiesPropertyOutputReference
	// Experimental.
	RightJoinKeyPropertiesInput() *TfDataSet_RightJoinKeyPropertiesProperty
	// Experimental.
	RightOperand() *string
	// Experimental.
	SetRightOperand(val *string)
	// Experimental.
	RightOperandInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	PutLeftJoinKeyProperties(value *TfDataSet_LeftJoinKeyPropertiesProperty)
	// Experimental.
	PutRightJoinKeyProperties(value *TfDataSet_RightJoinKeyPropertiesProperty)
	// Experimental.
	ResetLeftJoinKeyProperties()
	// Experimental.
	ResetRightJoinKeyProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDataSet_JoinInstructionPropertyOutputReference
type jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) InternalValue() *TfDataSet_JoinInstructionProperty {
	var returns *TfDataSet_JoinInstructionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) LeftJoinKeyProperties() TfDataSet_LeftJoinKeyPropertiesPropertyOutputReference {
	var returns TfDataSet_LeftJoinKeyPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"leftJoinKeyProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) LeftJoinKeyPropertiesInput() *TfDataSet_LeftJoinKeyPropertiesProperty {
	var returns *TfDataSet_LeftJoinKeyPropertiesProperty
	_jsii_.Get(
		j,
		"leftJoinKeyPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) LeftOperand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leftOperand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) LeftOperandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leftOperandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) OnClause() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onClause",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) OnClauseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onClauseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) RightJoinKeyProperties() TfDataSet_RightJoinKeyPropertiesPropertyOutputReference {
	var returns TfDataSet_RightJoinKeyPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"rightJoinKeyProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) RightJoinKeyPropertiesInput() *TfDataSet_RightJoinKeyPropertiesProperty {
	var returns *TfDataSet_RightJoinKeyPropertiesProperty
	_jsii_.Get(
		j,
		"rightJoinKeyPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) RightOperand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rightOperand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) RightOperandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rightOperandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataSet_JoinInstructionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDataSet_JoinInstructionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataSet_JoinInstructionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfDataSet.JoinInstructionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataSet_JoinInstructionPropertyOutputReference_Override(t TfDataSet_JoinInstructionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfDataSet.JoinInstructionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference)SetInternalValue(val *TfDataSet_JoinInstructionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference)SetLeftOperand(val *string) {
	if err := j.validateSetLeftOperandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leftOperand",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference)SetOnClause(val *string) {
	if err := j.validateSetOnClauseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onClause",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference)SetRightOperand(val *string) {
	if err := j.validateSetRightOperandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rightOperand",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) PutLeftJoinKeyProperties(value *TfDataSet_LeftJoinKeyPropertiesProperty) {
	if err := t.validatePutLeftJoinKeyPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLeftJoinKeyProperties",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) PutRightJoinKeyProperties(value *TfDataSet_RightJoinKeyPropertiesProperty) {
	if err := t.validatePutRightJoinKeyPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRightJoinKeyProperties",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) ResetLeftJoinKeyProperties() {
	_jsii_.InvokeVoid(
		t,
		"resetLeftJoinKeyProperties",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) ResetRightJoinKeyProperties() {
	_jsii_.InvokeVoid(
		t,
		"resetRightJoinKeyProperties",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataSet_JoinInstructionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

