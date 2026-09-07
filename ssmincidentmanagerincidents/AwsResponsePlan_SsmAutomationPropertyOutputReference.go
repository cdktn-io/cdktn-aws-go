package ssmincidentmanagerincidents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ssmincidentmanagerincidents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/ssmincidentmanagerincidents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsResponsePlan_SsmAutomationPropertyOutputReference interface {
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
	DocumentName() *string
	// Experimental.
	SetDocumentName(val *string)
	// Experimental.
	DocumentNameInput() *string
	// Experimental.
	DocumentVersion() *string
	// Experimental.
	SetDocumentVersion(val *string)
	// Experimental.
	DocumentVersionInput() *string
	// Experimental.
	DynamicParameters() *map[string]*string
	// Experimental.
	SetDynamicParameters(val *map[string]*string)
	// Experimental.
	DynamicParametersInput() *map[string]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Parameter() AwsResponsePlan_ParameterPropertyList
	// Experimental.
	ParameterInput() interface{}
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	TargetAccount() *string
	// Experimental.
	SetTargetAccount(val *string)
	// Experimental.
	TargetAccountInput() *string
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
	PutParameter(value interface{})
	// Experimental.
	ResetDocumentVersion()
	// Experimental.
	ResetDynamicParameters()
	// Experimental.
	ResetParameter()
	// Experimental.
	ResetTargetAccount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsResponsePlan_SsmAutomationPropertyOutputReference
type jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) DocumentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) DocumentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) DocumentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) DocumentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) DynamicParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dynamicParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) DynamicParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dynamicParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) Parameter() AwsResponsePlan_ParameterPropertyList {
	var returns AwsResponsePlan_ParameterPropertyList
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) ParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) TargetAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) TargetAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsResponsePlan_SsmAutomationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsResponsePlan_SsmAutomationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsResponsePlan_SsmAutomationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ssm-incident-manager-incidents.AwsResponsePlan.SsmAutomationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsResponsePlan_SsmAutomationPropertyOutputReference_Override(a AwsResponsePlan_SsmAutomationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ssm-incident-manager-incidents.AwsResponsePlan.SsmAutomationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference)SetDocumentName(val *string) {
	if err := j.validateSetDocumentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentName",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference)SetDocumentVersion(val *string) {
	if err := j.validateSetDocumentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentVersion",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference)SetDynamicParameters(val *map[string]*string) {
	if err := j.validateSetDynamicParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicParameters",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference)SetTargetAccount(val *string) {
	if err := j.validateSetTargetAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetAccount",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) PutParameter(value interface{}) {
	if err := a.validatePutParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParameter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) ResetDocumentVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetDocumentVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) ResetDynamicParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamicParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) ResetParameter() {
	_jsii_.InvokeVoid(
		a,
		"resetParameter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) ResetTargetAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsResponsePlan_SsmAutomationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

