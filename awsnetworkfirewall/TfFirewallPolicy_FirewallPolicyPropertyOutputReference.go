package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFirewallPolicy_FirewallPolicyPropertyOutputReference interface {
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
	InternalValue() *TfFirewallPolicy_FirewallPolicyProperty
	// Experimental.
	SetInternalValue(val *TfFirewallPolicy_FirewallPolicyProperty)
	// Experimental.
	PolicyVariables() TfFirewallPolicy_PolicyVariablesPropertyOutputReference
	// Experimental.
	PolicyVariablesInput() *TfFirewallPolicy_PolicyVariablesProperty
	// Experimental.
	StatefulDefaultActions() *[]*string
	// Experimental.
	SetStatefulDefaultActions(val *[]*string)
	// Experimental.
	StatefulDefaultActionsInput() *[]*string
	// Experimental.
	StatefulEngineOptions() TfFirewallPolicy_StatefulEngineOptionsPropertyOutputReference
	// Experimental.
	StatefulEngineOptionsInput() *TfFirewallPolicy_StatefulEngineOptionsProperty
	// Experimental.
	StatefulRuleGroupReference() TfFirewallPolicy_StatefulRuleGroupReferencePropertyList
	// Experimental.
	StatefulRuleGroupReferenceInput() interface{}
	// Experimental.
	StatelessCustomAction() TfFirewallPolicy_StatelessCustomActionPropertyList
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
	StatelessRuleGroupReference() TfFirewallPolicy_StatelessRuleGroupReferencePropertyList
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
	PutPolicyVariables(value *TfFirewallPolicy_PolicyVariablesProperty)
	// Experimental.
	PutStatefulEngineOptions(value *TfFirewallPolicy_StatefulEngineOptionsProperty)
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

// The jsii proxy struct for TfFirewallPolicy_FirewallPolicyPropertyOutputReference
type jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) EnableTlsSessionHolding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTlsSessionHolding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) EnableTlsSessionHoldingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTlsSessionHoldingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) InternalValue() *TfFirewallPolicy_FirewallPolicyProperty {
	var returns *TfFirewallPolicy_FirewallPolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) PolicyVariables() TfFirewallPolicy_PolicyVariablesPropertyOutputReference {
	var returns TfFirewallPolicy_PolicyVariablesPropertyOutputReference
	_jsii_.Get(
		j,
		"policyVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) PolicyVariablesInput() *TfFirewallPolicy_PolicyVariablesProperty {
	var returns *TfFirewallPolicy_PolicyVariablesProperty
	_jsii_.Get(
		j,
		"policyVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) StatefulDefaultActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statefulDefaultActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) StatefulDefaultActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statefulDefaultActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) StatefulEngineOptions() TfFirewallPolicy_StatefulEngineOptionsPropertyOutputReference {
	var returns TfFirewallPolicy_StatefulEngineOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"statefulEngineOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) StatefulEngineOptionsInput() *TfFirewallPolicy_StatefulEngineOptionsProperty {
	var returns *TfFirewallPolicy_StatefulEngineOptionsProperty
	_jsii_.Get(
		j,
		"statefulEngineOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) StatefulRuleGroupReference() TfFirewallPolicy_StatefulRuleGroupReferencePropertyList {
	var returns TfFirewallPolicy_StatefulRuleGroupReferencePropertyList
	_jsii_.Get(
		j,
		"statefulRuleGroupReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) StatefulRuleGroupReferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulRuleGroupReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) StatelessCustomAction() TfFirewallPolicy_StatelessCustomActionPropertyList {
	var returns TfFirewallPolicy_StatelessCustomActionPropertyList
	_jsii_.Get(
		j,
		"statelessCustomAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) StatelessCustomActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statelessCustomActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) StatelessDefaultActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statelessDefaultActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) StatelessDefaultActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statelessDefaultActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) StatelessFragmentDefaultActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statelessFragmentDefaultActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) StatelessFragmentDefaultActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statelessFragmentDefaultActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) StatelessRuleGroupReference() TfFirewallPolicy_StatelessRuleGroupReferencePropertyList {
	var returns TfFirewallPolicy_StatelessRuleGroupReferencePropertyList
	_jsii_.Get(
		j,
		"statelessRuleGroupReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) StatelessRuleGroupReferenceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statelessRuleGroupReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) TlsInspectionConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsInspectionConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) TlsInspectionConfigurationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsInspectionConfigurationArnInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFirewallPolicy_FirewallPolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFirewallPolicy_FirewallPolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFirewallPolicy_FirewallPolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-firewall.TfFirewallPolicy.FirewallPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFirewallPolicy_FirewallPolicyPropertyOutputReference_Override(t TfFirewallPolicy_FirewallPolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-firewall.TfFirewallPolicy.FirewallPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference)SetEnableTlsSessionHolding(val interface{}) {
	if err := j.validateSetEnableTlsSessionHoldingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTlsSessionHolding",
		val,
	)
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference)SetInternalValue(val *TfFirewallPolicy_FirewallPolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference)SetStatefulDefaultActions(val *[]*string) {
	if err := j.validateSetStatefulDefaultActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statefulDefaultActions",
		val,
	)
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference)SetStatelessDefaultActions(val *[]*string) {
	if err := j.validateSetStatelessDefaultActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statelessDefaultActions",
		val,
	)
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference)SetStatelessFragmentDefaultActions(val *[]*string) {
	if err := j.validateSetStatelessFragmentDefaultActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statelessFragmentDefaultActions",
		val,
	)
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference)SetTlsInspectionConfigurationArn(val *string) {
	if err := j.validateSetTlsInspectionConfigurationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsInspectionConfigurationArn",
		val,
	)
}

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) PutPolicyVariables(value *TfFirewallPolicy_PolicyVariablesProperty) {
	if err := t.validatePutPolicyVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPolicyVariables",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) PutStatefulEngineOptions(value *TfFirewallPolicy_StatefulEngineOptionsProperty) {
	if err := t.validatePutStatefulEngineOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStatefulEngineOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) PutStatefulRuleGroupReference(value interface{}) {
	if err := t.validatePutStatefulRuleGroupReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStatefulRuleGroupReference",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) PutStatelessCustomAction(value interface{}) {
	if err := t.validatePutStatelessCustomActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStatelessCustomAction",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) PutStatelessRuleGroupReference(value interface{}) {
	if err := t.validatePutStatelessRuleGroupReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStatelessRuleGroupReference",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) ResetEnableTlsSessionHolding() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableTlsSessionHolding",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) ResetPolicyVariables() {
	_jsii_.InvokeVoid(
		t,
		"resetPolicyVariables",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) ResetStatefulDefaultActions() {
	_jsii_.InvokeVoid(
		t,
		"resetStatefulDefaultActions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) ResetStatefulEngineOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetStatefulEngineOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) ResetStatefulRuleGroupReference() {
	_jsii_.InvokeVoid(
		t,
		"resetStatefulRuleGroupReference",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) ResetStatelessCustomAction() {
	_jsii_.InvokeVoid(
		t,
		"resetStatelessCustomAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) ResetStatelessRuleGroupReference() {
	_jsii_.InvokeVoid(
		t,
		"resetStatelessRuleGroupReference",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) ResetTlsInspectionConfigurationArn() {
	_jsii_.InvokeVoid(
		t,
		"resetTlsInspectionConfigurationArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFirewallPolicy_FirewallPolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

