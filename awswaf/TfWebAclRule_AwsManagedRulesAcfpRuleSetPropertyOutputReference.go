package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference interface {
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
	// Experimental.
	CreationPath() *string
	// Experimental.
	SetCreationPath(val *string)
	// Experimental.
	CreationPathInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EnableRegexInPath() interface{}
	// Experimental.
	SetEnableRegexInPath(val interface{})
	// Experimental.
	EnableRegexInPathInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	RegistrationPagePath() *string
	// Experimental.
	SetRegistrationPagePath(val *string)
	// Experimental.
	RegistrationPagePathInput() *string
	// Experimental.
	RequestInspection() TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetRequestInspectionPropertyList
	// Experimental.
	RequestInspectionInput() interface{}
	// Experimental.
	ResponseInspection() TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetResponseInspectionPropertyList
	// Experimental.
	ResponseInspectionInput() interface{}
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
	PutRequestInspection(value interface{})
	// Experimental.
	PutResponseInspection(value interface{})
	// Experimental.
	ResetEnableRegexInPath()
	// Experimental.
	ResetRequestInspection()
	// Experimental.
	ResetResponseInspection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference
type jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) CreationPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) CreationPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) EnableRegexInPath() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRegexInPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) EnableRegexInPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRegexInPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) RegistrationPagePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registrationPagePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) RegistrationPagePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registrationPagePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) RequestInspection() TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetRequestInspectionPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetRequestInspectionPropertyList
	_jsii_.Get(
		j,
		"requestInspection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) RequestInspectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestInspectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) ResponseInspection() TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetResponseInspectionPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAcfpRuleSetResponseInspectionPropertyList
	_jsii_.Get(
		j,
		"responseInspection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) ResponseInspectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseInspectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.AwsManagedRulesAcfpRuleSetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference_Override(t TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.AwsManagedRulesAcfpRuleSetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference)SetCreationPath(val *string) {
	if err := j.validateSetCreationPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creationPath",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference)SetEnableRegexInPath(val interface{}) {
	if err := j.validateSetEnableRegexInPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableRegexInPath",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference)SetRegistrationPagePath(val *string) {
	if err := j.validateSetRegistrationPagePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registrationPagePath",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) PutRequestInspection(value interface{}) {
	if err := t.validatePutRequestInspectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRequestInspection",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) PutResponseInspection(value interface{}) {
	if err := t.validatePutResponseInspectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResponseInspection",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) ResetEnableRegexInPath() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableRegexInPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) ResetRequestInspection() {
	_jsii_.InvokeVoid(
		t,
		"resetRequestInspection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) ResetResponseInspection() {
	_jsii_.InvokeVoid(
		t,
		"resetResponseInspection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_AwsManagedRulesAcfpRuleSetPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

