package awsssmincidentmanagerincidents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsssmincidentmanagerincidents/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsssmincidentmanagerincidents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference interface {
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
	Parameter() AwsSsmincidentsResponsePlan_ParameterPropertyList
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

// The jsii proxy struct for AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference
type jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) DocumentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) DocumentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) DocumentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) DocumentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) DynamicParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dynamicParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) DynamicParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dynamicParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) Parameter() AwsSsmincidentsResponsePlan_ParameterPropertyList {
	var returns AwsSsmincidentsResponsePlan_ParameterPropertyList
	_jsii_.Get(
		j,
		"parameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) ParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) TargetAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) TargetAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ssm-incident-manager-incidents.AwsSsmincidentsResponsePlan.SsmAutomationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference_Override(a AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ssm-incident-manager-incidents.AwsSsmincidentsResponsePlan.SsmAutomationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference)SetDocumentName(val *string) {
	if err := j.validateSetDocumentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentName",
		val,
	)
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference)SetDocumentVersion(val *string) {
	if err := j.validateSetDocumentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentVersion",
		val,
	)
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference)SetDynamicParameters(val *map[string]*string) {
	if err := j.validateSetDynamicParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicParameters",
		val,
	)
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference)SetTargetAccount(val *string) {
	if err := j.validateSetTargetAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetAccount",
		val,
	)
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) PutParameter(value interface{}) {
	if err := a.validatePutParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParameter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) ResetDocumentVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetDocumentVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) ResetDynamicParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamicParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) ResetParameter() {
	_jsii_.InvokeVoid(
		a,
		"resetParameter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) ResetTargetAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSsmincidentsResponsePlan_SsmAutomationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

