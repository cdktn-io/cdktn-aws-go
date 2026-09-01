package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference interface {
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
	FlowTimeouts() AwsNetworkfirewallFirewallPolicy_FlowTimeoutsPropertyOutputReference
	// Experimental.
	FlowTimeoutsInput() *AwsNetworkfirewallFirewallPolicy_FlowTimeoutsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsProperty
	// Experimental.
	SetInternalValue(val *AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsProperty)
	// Experimental.
	RuleOrder() *string
	// Experimental.
	SetRuleOrder(val *string)
	// Experimental.
	RuleOrderInput() *string
	// Experimental.
	StreamExceptionPolicy() *string
	// Experimental.
	SetStreamExceptionPolicy(val *string)
	// Experimental.
	StreamExceptionPolicyInput() *string
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
	PutFlowTimeouts(value *AwsNetworkfirewallFirewallPolicy_FlowTimeoutsProperty)
	// Experimental.
	ResetFlowTimeouts()
	// Experimental.
	ResetRuleOrder()
	// Experimental.
	ResetStreamExceptionPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference
type jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) FlowTimeouts() AwsNetworkfirewallFirewallPolicy_FlowTimeoutsPropertyOutputReference {
	var returns AwsNetworkfirewallFirewallPolicy_FlowTimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"flowTimeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) FlowTimeoutsInput() *AwsNetworkfirewallFirewallPolicy_FlowTimeoutsProperty {
	var returns *AwsNetworkfirewallFirewallPolicy_FlowTimeoutsProperty
	_jsii_.Get(
		j,
		"flowTimeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) InternalValue() *AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsProperty {
	var returns *AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) RuleOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) RuleOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) StreamExceptionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamExceptionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) StreamExceptionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamExceptionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsNetworkfirewallFirewallPolicy.StatefulEngineOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference_Override(a AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsNetworkfirewallFirewallPolicy.StatefulEngineOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference)SetInternalValue(val *AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference)SetRuleOrder(val *string) {
	if err := j.validateSetRuleOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleOrder",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference)SetStreamExceptionPolicy(val *string) {
	if err := j.validateSetStreamExceptionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamExceptionPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) PutFlowTimeouts(value *AwsNetworkfirewallFirewallPolicy_FlowTimeoutsProperty) {
	if err := a.validatePutFlowTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFlowTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) ResetFlowTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetFlowTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) ResetRuleOrder() {
	_jsii_.InvokeVoid(
		a,
		"resetRuleOrder",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) ResetStreamExceptionPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetStreamExceptionPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsNetworkfirewallFirewallPolicy_StatefulEngineOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

