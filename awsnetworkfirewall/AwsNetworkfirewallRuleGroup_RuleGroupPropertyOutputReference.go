package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference interface {
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
	InternalValue() *AwsNetworkfirewallRuleGroup_RuleGroupProperty
	// Experimental.
	SetInternalValue(val *AwsNetworkfirewallRuleGroup_RuleGroupProperty)
	// Experimental.
	ReferenceSets() AwsNetworkfirewallRuleGroup_ReferenceSetsPropertyOutputReference
	// Experimental.
	ReferenceSetsInput() *AwsNetworkfirewallRuleGroup_ReferenceSetsProperty
	// Experimental.
	RulesSource() AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference
	// Experimental.
	RulesSourceInput() *AwsNetworkfirewallRuleGroup_RulesSourceProperty
	// Experimental.
	RuleVariables() AwsNetworkfirewallRuleGroup_RuleVariablesPropertyOutputReference
	// Experimental.
	RuleVariablesInput() *AwsNetworkfirewallRuleGroup_RuleVariablesProperty
	// Experimental.
	StatefulRuleOptions() AwsNetworkfirewallRuleGroup_StatefulRuleOptionsPropertyOutputReference
	// Experimental.
	StatefulRuleOptionsInput() *AwsNetworkfirewallRuleGroup_StatefulRuleOptionsProperty
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
	PutReferenceSets(value *AwsNetworkfirewallRuleGroup_ReferenceSetsProperty)
	// Experimental.
	PutRulesSource(value *AwsNetworkfirewallRuleGroup_RulesSourceProperty)
	// Experimental.
	PutRuleVariables(value *AwsNetworkfirewallRuleGroup_RuleVariablesProperty)
	// Experimental.
	PutStatefulRuleOptions(value *AwsNetworkfirewallRuleGroup_StatefulRuleOptionsProperty)
	// Experimental.
	ResetReferenceSets()
	// Experimental.
	ResetRuleVariables()
	// Experimental.
	ResetStatefulRuleOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference
type jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) InternalValue() *AwsNetworkfirewallRuleGroup_RuleGroupProperty {
	var returns *AwsNetworkfirewallRuleGroup_RuleGroupProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) ReferenceSets() AwsNetworkfirewallRuleGroup_ReferenceSetsPropertyOutputReference {
	var returns AwsNetworkfirewallRuleGroup_ReferenceSetsPropertyOutputReference
	_jsii_.Get(
		j,
		"referenceSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) ReferenceSetsInput() *AwsNetworkfirewallRuleGroup_ReferenceSetsProperty {
	var returns *AwsNetworkfirewallRuleGroup_ReferenceSetsProperty
	_jsii_.Get(
		j,
		"referenceSetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) RulesSource() AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference {
	var returns AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference
	_jsii_.Get(
		j,
		"rulesSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) RulesSourceInput() *AwsNetworkfirewallRuleGroup_RulesSourceProperty {
	var returns *AwsNetworkfirewallRuleGroup_RulesSourceProperty
	_jsii_.Get(
		j,
		"rulesSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) RuleVariables() AwsNetworkfirewallRuleGroup_RuleVariablesPropertyOutputReference {
	var returns AwsNetworkfirewallRuleGroup_RuleVariablesPropertyOutputReference
	_jsii_.Get(
		j,
		"ruleVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) RuleVariablesInput() *AwsNetworkfirewallRuleGroup_RuleVariablesProperty {
	var returns *AwsNetworkfirewallRuleGroup_RuleVariablesProperty
	_jsii_.Get(
		j,
		"ruleVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) StatefulRuleOptions() AwsNetworkfirewallRuleGroup_StatefulRuleOptionsPropertyOutputReference {
	var returns AwsNetworkfirewallRuleGroup_StatefulRuleOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"statefulRuleOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) StatefulRuleOptionsInput() *AwsNetworkfirewallRuleGroup_StatefulRuleOptionsProperty {
	var returns *AwsNetworkfirewallRuleGroup_StatefulRuleOptionsProperty
	_jsii_.Get(
		j,
		"statefulRuleOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsNetworkfirewallRuleGroup.RuleGroupPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference_Override(a AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsNetworkfirewallRuleGroup.RuleGroupPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference)SetInternalValue(val *AwsNetworkfirewallRuleGroup_RuleGroupProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) PutReferenceSets(value *AwsNetworkfirewallRuleGroup_ReferenceSetsProperty) {
	if err := a.validatePutReferenceSetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReferenceSets",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) PutRulesSource(value *AwsNetworkfirewallRuleGroup_RulesSourceProperty) {
	if err := a.validatePutRulesSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRulesSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) PutRuleVariables(value *AwsNetworkfirewallRuleGroup_RuleVariablesProperty) {
	if err := a.validatePutRuleVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRuleVariables",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) PutStatefulRuleOptions(value *AwsNetworkfirewallRuleGroup_StatefulRuleOptionsProperty) {
	if err := a.validatePutStatefulRuleOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStatefulRuleOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) ResetReferenceSets() {
	_jsii_.InvokeVoid(
		a,
		"resetReferenceSets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) ResetRuleVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetRuleVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) ResetStatefulRuleOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetStatefulRuleOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RuleGroupPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

