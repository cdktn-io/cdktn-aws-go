package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BodyContains() TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionBodyContainsPropertyList
	// Experimental.
	BodyContainsInput() interface{}
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
	Header() TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionHeaderPropertyList
	// Experimental.
	HeaderInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Json() TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionJsonPropertyList
	// Experimental.
	JsonInput() interface{}
	// Experimental.
	StatusCode() TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionStatusCodePropertyList
	// Experimental.
	StatusCodeInput() interface{}
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
	PutBodyContains(value interface{})
	// Experimental.
	PutHeader(value interface{})
	// Experimental.
	PutJson(value interface{})
	// Experimental.
	PutStatusCode(value interface{})
	// Experimental.
	ResetBodyContains()
	// Experimental.
	ResetHeader()
	// Experimental.
	ResetJson()
	// Experimental.
	ResetStatusCode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference
type jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) BodyContains() TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionBodyContainsPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionBodyContainsPropertyList
	_jsii_.Get(
		j,
		"bodyContains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) BodyContainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bodyContainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) Header() TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionHeaderPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionHeaderPropertyList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) Json() TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionJsonPropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionJsonPropertyList
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) JsonInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) StatusCode() TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionStatusCodePropertyList {
	var returns TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionStatusCodePropertyList
	_jsii_.Get(
		j,
		"statusCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) StatusCodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statusCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference_Override(t TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRule.StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) PutBodyContains(value interface{}) {
	if err := t.validatePutBodyContainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBodyContains",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) PutHeader(value interface{}) {
	if err := t.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) PutJson(value interface{}) {
	if err := t.validatePutJsonParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJson",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) PutStatusCode(value interface{}) {
	if err := t.validatePutStatusCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStatusCode",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) ResetBodyContains() {
	_jsii_.InvokeVoid(
		t,
		"resetBodyContains",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) ResetJson() {
	_jsii_.InvokeVoid(
		t,
		"resetJson",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) ResetStatusCode() {
	_jsii_.InvokeVoid(
		t,
		"resetStatusCode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

