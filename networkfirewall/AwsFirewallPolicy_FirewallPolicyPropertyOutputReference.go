package networkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/networkfirewall/jsii"

	"github.com/cdktn-io/cdktn-aws-go/networkfirewall/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFirewallPolicy_FirewallPolicyPropertyOutputReference interface {
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
	EnableTlsSessionHolding() interface{}
	// Experimental.
	SetEnableTlsSessionHolding(val interface{})
	// Experimental.
	EnableTlsSessionHoldingInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsFirewallPolicy_FirewallPolicyProperty
	// Experimental.
	SetInternalValue(val *AwsFirewallPolicy_FirewallPolicyProperty)
	// Experimental.
	PolicyVariables() AwsFirewallPolicy_PolicyVariablesPropertyOutputReference
	// Experimental.
	PolicyVariablesInput() *AwsFirewallPolicy_PolicyVariablesProperty
	// Experimental.
	StatefulDefaultActions() *[]*string
	// Experimental.
	SetStatefulDefaultActions(val *[]*string)
	// Experimental.
	StatefulDefaultActionsInput() *[]*string
	// Experimental.
	StatefulEngineOptions() AwsFirewallPolicy_StatefulEngineOptionsPropertyOutputReference
	// Experimental.
	StatefulEngineOptionsInput() *AwsFirewallPolicy_StatefulEngineOptionsProperty
	// Experimental.
	StatefulRuleGroupReference() AwsFirewallPolicy_StatefulRuleGroupReferencePropertyList
	// Experimental.
	StatefulRuleGroupReferenceInput() interface{}
	// Experimental.
	StatelessCustomAction() AwsFirewallPolicy_StatelessCustomActionPropertyList
	// Experimental.
	StatelessCustomActionInput() interface{}
	// Experimental.
	StatelessDefaultActions() *[]*string
	// Experimental.
	SetStatelessDefaultActions(val *[]*string)
	// Experimental.
	StatelessDefaultActionsInput() *[]*string
	// Experimental.
	StatelessFragmentDefaultActions() *[]*string
	// Experimental.
	SetStatelessFragmentDefaultActions(val *[]*string)
	// Experimental.
	StatelessFragmentDefaultActionsInput() *[]*string
	// Experimental.
	StatelessRuleGroupReference() AwsFirewallPolicy_StatelessRuleGroupReferencePropertyList
	// Experimental.
	StatelessRuleGroupReferenceInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TlsInspectionConfigurationArn() *string
	// Experimental.
	SetTlsInspectionConfigurationArn(val *string)
	// Experimental.
	TlsInspectionConfigurationArnInput() *string
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
	PutPolicyVariables(value *AwsFirewallPolicy_PolicyVariablesProperty)
	// Experimental.
	PutStatefulEngineOptions(value *AwsFirewallPolicy_StatefulEngineOptionsProperty)
	// Experimental.
	PutStatefulRuleGroupReference(value interface{})
	// Experimental.
	PutStatelessCustomAction(value interface{})
	// Experimental.
	PutStatelessRuleGroupReference(value interface{})
	// Experimental.
	ResetEnableTlsSessionHolding()
	// Experimental.
	ResetPolicyVariables()
	// Experimental.
	ResetStatefulDefaultActions()
	// Experimental.
	ResetStatefulEngineOptions()
	// Experimental.
	ResetStatefulRuleGroupReference()
	// Experimental.
	ResetStatelessCustomAction()
	// Experimental.
	ResetStatelessRuleGroupReference()
	// Experimental.
	ResetTlsInspectionConfigurationArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsFirewallPolicy_FirewallPolicyPropertyOutputReference
type jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) EnableTlsSessionHolding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTlsSessionHolding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) EnableTlsSessionHoldingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTlsSessionHoldingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) InternalValue() *AwsFirewallPolicy_FirewallPolicyProperty {
	var returns *AwsFirewallPolicy_FirewallPolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) PolicyVariables() AwsFirewallPolicy_PolicyVariablesPropertyOutputReference {
	var returns AwsFirewallPolicy_PolicyVariablesPropertyOutputReference
	_jsii_.Get(
		j,
		"policyVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) PolicyVariablesInput() *AwsFirewallPolicy_PolicyVariablesProperty {
	var returns *AwsFirewallPolicy_PolicyVariablesProperty
	_jsii_.Get(
		j,
		"policyVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) StatefulDefaultActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statefulDefaultActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) StatefulDefaultActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statefulDefaultActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) StatefulEngineOptions() AwsFirewallPolicy_StatefulEngineOptionsPropertyOutputReference {
	var returns AwsFirewallPolicy_StatefulEngineOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"statefulEngineOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) StatefulEngineOptionsInput() *AwsFirewallPolicy_StatefulEngineOptionsProperty {
	var returns *AwsFirewallPolicy_StatefulEngineOptionsProperty
	_jsii_.Get(
		j,
		"statefulEngineOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) StatefulRuleGroupReference() AwsFirewallPolicy_StatefulRuleGroupReferencePropertyList {
	var returns AwsFirewallPolicy_StatefulRuleGroupReferencePropertyList
	_jsii_.Get(
		j,
		"statefulRuleGroupReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) StatefulRuleGroupReferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulRuleGroupReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) StatelessCustomAction() AwsFirewallPolicy_StatelessCustomActionPropertyList {
	var returns AwsFirewallPolicy_StatelessCustomActionPropertyList
	_jsii_.Get(
		j,
		"statelessCustomAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) StatelessCustomActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statelessCustomActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) StatelessDefaultActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statelessDefaultActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) StatelessDefaultActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statelessDefaultActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) StatelessFragmentDefaultActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statelessFragmentDefaultActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) StatelessFragmentDefaultActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statelessFragmentDefaultActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) StatelessRuleGroupReference() AwsFirewallPolicy_StatelessRuleGroupReferencePropertyList {
	var returns AwsFirewallPolicy_StatelessRuleGroupReferencePropertyList
	_jsii_.Get(
		j,
		"statelessRuleGroupReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) StatelessRuleGroupReferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statelessRuleGroupReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) TlsInspectionConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsInspectionConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) TlsInspectionConfigurationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsInspectionConfigurationArnInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFirewallPolicy_FirewallPolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFirewallPolicy_FirewallPolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFirewallPolicy_FirewallPolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsFirewallPolicy.FirewallPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFirewallPolicy_FirewallPolicyPropertyOutputReference_Override(a AwsFirewallPolicy_FirewallPolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsFirewallPolicy.FirewallPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference)SetEnableTlsSessionHolding(val interface{}) {
	if err := j.validateSetEnableTlsSessionHoldingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTlsSessionHolding",
		val,
	)
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference)SetInternalValue(val *AwsFirewallPolicy_FirewallPolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference)SetStatefulDefaultActions(val *[]*string) {
	if err := j.validateSetStatefulDefaultActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statefulDefaultActions",
		val,
	)
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference)SetStatelessDefaultActions(val *[]*string) {
	if err := j.validateSetStatelessDefaultActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statelessDefaultActions",
		val,
	)
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference)SetStatelessFragmentDefaultActions(val *[]*string) {
	if err := j.validateSetStatelessFragmentDefaultActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statelessFragmentDefaultActions",
		val,
	)
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference)SetTlsInspectionConfigurationArn(val *string) {
	if err := j.validateSetTlsInspectionConfigurationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsInspectionConfigurationArn",
		val,
	)
}

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) PutPolicyVariables(value *AwsFirewallPolicy_PolicyVariablesProperty) {
	if err := a.validatePutPolicyVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPolicyVariables",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) PutStatefulEngineOptions(value *AwsFirewallPolicy_StatefulEngineOptionsProperty) {
	if err := a.validatePutStatefulEngineOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStatefulEngineOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) PutStatefulRuleGroupReference(value interface{}) {
	if err := a.validatePutStatefulRuleGroupReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStatefulRuleGroupReference",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) PutStatelessCustomAction(value interface{}) {
	if err := a.validatePutStatelessCustomActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStatelessCustomAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) PutStatelessRuleGroupReference(value interface{}) {
	if err := a.validatePutStatelessRuleGroupReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStatelessRuleGroupReference",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) ResetEnableTlsSessionHolding() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableTlsSessionHolding",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) ResetPolicyVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) ResetStatefulDefaultActions() {
	_jsii_.InvokeVoid(
		a,
		"resetStatefulDefaultActions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) ResetStatefulEngineOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetStatefulEngineOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) ResetStatefulRuleGroupReference() {
	_jsii_.InvokeVoid(
		a,
		"resetStatefulRuleGroupReference",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) ResetStatelessCustomAction() {
	_jsii_.InvokeVoid(
		a,
		"resetStatelessCustomAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) ResetStatelessRuleGroupReference() {
	_jsii_.InvokeVoid(
		a,
		"resetStatelessRuleGroupReference",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) ResetTlsInspectionConfigurationArn() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsInspectionConfigurationArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFirewallPolicy_FirewallPolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

