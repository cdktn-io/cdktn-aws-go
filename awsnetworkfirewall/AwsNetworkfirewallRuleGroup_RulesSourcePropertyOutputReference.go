package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference interface {
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
	InternalValue() *AwsNetworkfirewallRuleGroup_RulesSourceProperty
	// Experimental.
	SetInternalValue(val *AwsNetworkfirewallRuleGroup_RulesSourceProperty)
	// Experimental.
	RulesSourceList() AwsNetworkfirewallRuleGroup_RulesSourceListPropertyOutputReference
	// Experimental.
	RulesSourceListInput() *AwsNetworkfirewallRuleGroup_RulesSourceListProperty
	// Experimental.
	RulesString() *string
	// Experimental.
	SetRulesString(val *string)
	// Experimental.
	RulesStringInput() *string
	// Experimental.
	StatefulRule() AwsNetworkfirewallRuleGroup_StatefulRulePropertyList
	// Experimental.
	StatefulRuleInput() interface{}
	// Experimental.
	StatelessRulesAndCustomActions() AwsNetworkfirewallRuleGroup_StatelessRulesAndCustomActionsPropertyOutputReference
	// Experimental.
	StatelessRulesAndCustomActionsInput() *AwsNetworkfirewallRuleGroup_StatelessRulesAndCustomActionsProperty
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
	PutRulesSourceList(value *AwsNetworkfirewallRuleGroup_RulesSourceListProperty)
	// Experimental.
	PutStatefulRule(value interface{})
	// Experimental.
	PutStatelessRulesAndCustomActions(value *AwsNetworkfirewallRuleGroup_StatelessRulesAndCustomActionsProperty)
	// Experimental.
	ResetRulesSourceList()
	// Experimental.
	ResetRulesString()
	// Experimental.
	ResetStatefulRule()
	// Experimental.
	ResetStatelessRulesAndCustomActions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference
type jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) InternalValue() *AwsNetworkfirewallRuleGroup_RulesSourceProperty {
	var returns *AwsNetworkfirewallRuleGroup_RulesSourceProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) RulesSourceList() AwsNetworkfirewallRuleGroup_RulesSourceListPropertyOutputReference {
	var returns AwsNetworkfirewallRuleGroup_RulesSourceListPropertyOutputReference
	_jsii_.Get(
		j,
		"rulesSourceList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) RulesSourceListInput() *AwsNetworkfirewallRuleGroup_RulesSourceListProperty {
	var returns *AwsNetworkfirewallRuleGroup_RulesSourceListProperty
	_jsii_.Get(
		j,
		"rulesSourceListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) RulesString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rulesString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) RulesStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rulesStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) StatefulRule() AwsNetworkfirewallRuleGroup_StatefulRulePropertyList {
	var returns AwsNetworkfirewallRuleGroup_StatefulRulePropertyList
	_jsii_.Get(
		j,
		"statefulRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) StatefulRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) StatelessRulesAndCustomActions() AwsNetworkfirewallRuleGroup_StatelessRulesAndCustomActionsPropertyOutputReference {
	var returns AwsNetworkfirewallRuleGroup_StatelessRulesAndCustomActionsPropertyOutputReference
	_jsii_.Get(
		j,
		"statelessRulesAndCustomActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) StatelessRulesAndCustomActionsInput() *AwsNetworkfirewallRuleGroup_StatelessRulesAndCustomActionsProperty {
	var returns *AwsNetworkfirewallRuleGroup_StatelessRulesAndCustomActionsProperty
	_jsii_.Get(
		j,
		"statelessRulesAndCustomActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsNetworkfirewallRuleGroup.RulesSourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference_Override(a AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsNetworkfirewallRuleGroup.RulesSourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference)SetInternalValue(val *AwsNetworkfirewallRuleGroup_RulesSourceProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference)SetRulesString(val *string) {
	if err := j.validateSetRulesStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rulesString",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) PutRulesSourceList(value *AwsNetworkfirewallRuleGroup_RulesSourceListProperty) {
	if err := a.validatePutRulesSourceListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRulesSourceList",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) PutStatefulRule(value interface{}) {
	if err := a.validatePutStatefulRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStatefulRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) PutStatelessRulesAndCustomActions(value *AwsNetworkfirewallRuleGroup_StatelessRulesAndCustomActionsProperty) {
	if err := a.validatePutStatelessRulesAndCustomActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStatelessRulesAndCustomActions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) ResetRulesSourceList() {
	_jsii_.InvokeVoid(
		a,
		"resetRulesSourceList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) ResetRulesString() {
	_jsii_.InvokeVoid(
		a,
		"resetRulesString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) ResetStatefulRule() {
	_jsii_.InvokeVoid(
		a,
		"resetStatefulRule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) ResetStatelessRulesAndCustomActions() {
	_jsii_.InvokeVoid(
		a,
		"resetStatelessRulesAndCustomActions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsNetworkfirewallRuleGroup_RulesSourcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

