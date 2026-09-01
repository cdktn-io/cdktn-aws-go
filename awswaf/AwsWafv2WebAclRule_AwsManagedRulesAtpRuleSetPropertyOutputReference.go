package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference interface {
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
	LoginPath() *string
	// Experimental.
	SetLoginPath(val *string)
	// Experimental.
	LoginPathInput() *string
	// Experimental.
	RequestInspection() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetRequestInspectionPropertyList
	// Experimental.
	RequestInspectionInput() interface{}
	// Experimental.
	ResponseInspection() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyList
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

// The jsii proxy struct for AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference
type jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) EnableRegexInPath() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRegexInPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) EnableRegexInPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableRegexInPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) LoginPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) LoginPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) RequestInspection() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetRequestInspectionPropertyList {
	var returns AwsWafv2WebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetRequestInspectionPropertyList
	_jsii_.Get(
		j,
		"requestInspection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) RequestInspectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestInspectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) ResponseInspection() AwsWafv2WebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyList {
	var returns AwsWafv2WebAclRule_StatementManagedRuleGroupStatementManagedRuleGroupConfigsAwsManagedRulesAtpRuleSetResponseInspectionPropertyList
	_jsii_.Get(
		j,
		"responseInspection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) ResponseInspectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseInspectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWafv2WebAclRule.AwsManagedRulesAtpRuleSetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference_Override(a AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWafv2WebAclRule.AwsManagedRulesAtpRuleSetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference)SetEnableRegexInPath(val interface{}) {
	if err := j.validateSetEnableRegexInPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableRegexInPath",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference)SetLoginPath(val *string) {
	if err := j.validateSetLoginPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginPath",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) PutRequestInspection(value interface{}) {
	if err := a.validatePutRequestInspectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRequestInspection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) PutResponseInspection(value interface{}) {
	if err := a.validatePutResponseInspectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResponseInspection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) ResetEnableRegexInPath() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableRegexInPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) ResetRequestInspection() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestInspection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) ResetResponseInspection() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseInspection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWafv2WebAclRule_AwsManagedRulesAtpRuleSetPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

