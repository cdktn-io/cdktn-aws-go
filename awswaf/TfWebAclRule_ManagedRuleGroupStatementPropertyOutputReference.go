package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ManagedRuleGroupConfigs() TfWebAclRule_ManagedRuleGroupConfigsPropertyList
	// Experimental.
	ManagedRuleGroupConfigsInput() interface{}
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	RuleActionOverride() TfWebAclRule_StatementManagedRuleGroupStatementRuleActionOverridePropertyList
	// Experimental.
	RuleActionOverrideInput() interface{}
	// Experimental.
	ScopeDownStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyList
	// Experimental.
	ScopeDownStatementInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VendorName() *string
	// Experimental.
	SetVendorName(val *string)
	// Experimental.
	VendorNameInput() *string
	// Experimental.
	Version() *string
	// Experimental.
	SetVersion(val *string)
	// Experimental.
	VersionInput() *string
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
	PutManagedRuleGroupConfigs(value interface{})
	// Experimental.
	PutRuleActionOverride(value interface{})
	// Experimental.
	PutScopeDownStatement(value interface{})
	// Experimental.
	ResetManagedRuleGroupConfigs()
	// Experimental.
	ResetRuleActionOverride()
	// Experimental.
	ResetScopeDownStatement()
	// Experimental.
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference
type jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) ManagedRuleGroupConfigs() TfWebAclRule_ManagedRuleGroupConfigsPropertyList {
	var returns TfWebAclRule_ManagedRuleGroupConfigsPropertyList
	_jsii_.Get(
		j,
		"managedRuleGroupConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) ManagedRuleGroupConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedRuleGroupConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) RuleActionOverride() TfWebAclRule_StatementManagedRuleGroupStatementRuleActionOverridePropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementRuleActionOverridePropertyList
	_jsii_.Get(
		j,
		"ruleActionOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) RuleActionOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleActionOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) ScopeDownStatement() TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementScopeDownStatementPropertyList
	_jsii_.Get(
		j,
		"scopeDownStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) ScopeDownStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scopeDownStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) VendorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vendorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) VendorNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vendorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_ManagedRuleGroupStatementPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.ManagedRuleGroupStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference_Override(t TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.ManagedRuleGroupStatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference)SetVendorName(val *string) {
	if err := j.validateSetVendorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vendorName",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) PutManagedRuleGroupConfigs(value interface{}) {
	if err := t.validatePutManagedRuleGroupConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putManagedRuleGroupConfigs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) PutRuleActionOverride(value interface{}) {
	if err := t.validatePutRuleActionOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRuleActionOverride",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) PutScopeDownStatement(value interface{}) {
	if err := t.validatePutScopeDownStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putScopeDownStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) ResetManagedRuleGroupConfigs() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedRuleGroupConfigs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) ResetRuleActionOverride() {
	_jsii_.InvokeVoid(
		t,
		"resetRuleActionOverride",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) ResetScopeDownStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetScopeDownStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_ManagedRuleGroupStatementPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

