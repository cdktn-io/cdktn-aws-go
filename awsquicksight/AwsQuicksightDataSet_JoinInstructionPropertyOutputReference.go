package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsQuicksightDataSet_JoinInstructionPropertyOutputReference interface {
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
	InternalValue() *AwsQuicksightDataSet_JoinInstructionProperty
	// Experimental.
	SetInternalValue(val *AwsQuicksightDataSet_JoinInstructionProperty)
	// Experimental.
	LeftJoinKeyProperties() AwsQuicksightDataSet_LeftJoinKeyPropertiesPropertyOutputReference
	// Experimental.
	LeftJoinKeyPropertiesInput() *AwsQuicksightDataSet_LeftJoinKeyPropertiesProperty
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
	RightJoinKeyProperties() AwsQuicksightDataSet_RightJoinKeyPropertiesPropertyOutputReference
	// Experimental.
	RightJoinKeyPropertiesInput() *AwsQuicksightDataSet_RightJoinKeyPropertiesProperty
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
	PutLeftJoinKeyProperties(value *AwsQuicksightDataSet_LeftJoinKeyPropertiesProperty)
	// Experimental.
	PutRightJoinKeyProperties(value *AwsQuicksightDataSet_RightJoinKeyPropertiesProperty)
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

// The jsii proxy struct for AwsQuicksightDataSet_JoinInstructionPropertyOutputReference
type jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) InternalValue() *AwsQuicksightDataSet_JoinInstructionProperty {
	var returns *AwsQuicksightDataSet_JoinInstructionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) LeftJoinKeyProperties() AwsQuicksightDataSet_LeftJoinKeyPropertiesPropertyOutputReference {
	var returns AwsQuicksightDataSet_LeftJoinKeyPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"leftJoinKeyProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) LeftJoinKeyPropertiesInput() *AwsQuicksightDataSet_LeftJoinKeyPropertiesProperty {
	var returns *AwsQuicksightDataSet_LeftJoinKeyPropertiesProperty
	_jsii_.Get(
		j,
		"leftJoinKeyPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) LeftOperand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leftOperand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) LeftOperandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leftOperandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) OnClause() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onClause",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) OnClauseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onClauseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) RightJoinKeyProperties() AwsQuicksightDataSet_RightJoinKeyPropertiesPropertyOutputReference {
	var returns AwsQuicksightDataSet_RightJoinKeyPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"rightJoinKeyProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) RightJoinKeyPropertiesInput() *AwsQuicksightDataSet_RightJoinKeyPropertiesProperty {
	var returns *AwsQuicksightDataSet_RightJoinKeyPropertiesProperty
	_jsii_.Get(
		j,
		"rightJoinKeyPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) RightOperand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rightOperand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) RightOperandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rightOperandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsQuicksightDataSet_JoinInstructionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsQuicksightDataSet_JoinInstructionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsQuicksightDataSet_JoinInstructionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsQuicksightDataSet.JoinInstructionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsQuicksightDataSet_JoinInstructionPropertyOutputReference_Override(a AwsQuicksightDataSet_JoinInstructionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsQuicksightDataSet.JoinInstructionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference)SetInternalValue(val *AwsQuicksightDataSet_JoinInstructionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference)SetLeftOperand(val *string) {
	if err := j.validateSetLeftOperandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leftOperand",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference)SetOnClause(val *string) {
	if err := j.validateSetOnClauseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onClause",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference)SetRightOperand(val *string) {
	if err := j.validateSetRightOperandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rightOperand",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) PutLeftJoinKeyProperties(value *AwsQuicksightDataSet_LeftJoinKeyPropertiesProperty) {
	if err := a.validatePutLeftJoinKeyPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLeftJoinKeyProperties",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) PutRightJoinKeyProperties(value *AwsQuicksightDataSet_RightJoinKeyPropertiesProperty) {
	if err := a.validatePutRightJoinKeyPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRightJoinKeyProperties",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) ResetLeftJoinKeyProperties() {
	_jsii_.InvokeVoid(
		a,
		"resetLeftJoinKeyProperties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) ResetRightJoinKeyProperties() {
	_jsii_.InvokeVoid(
		a,
		"resetRightJoinKeyProperties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsQuicksightDataSet_JoinInstructionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

