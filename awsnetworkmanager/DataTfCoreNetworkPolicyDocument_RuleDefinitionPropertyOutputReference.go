package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsnetworkmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Action() DataTfCoreNetworkPolicyDocument_RoutingPoliciesRoutingPolicyRulesRuleDefinitionActionPropertyOutputReference
	// Experimental.
	ActionInput() *DataTfCoreNetworkPolicyDocument_RoutingPoliciesRoutingPolicyRulesRuleDefinitionActionProperty
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
	ConditionLogic() *string
	// Experimental.
	SetConditionLogic(val *string)
	// Experimental.
	ConditionLogicInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *DataTfCoreNetworkPolicyDocument_RuleDefinitionProperty
	// Experimental.
	SetInternalValue(val *DataTfCoreNetworkPolicyDocument_RuleDefinitionProperty)
	// Experimental.
	MatchConditions() DataTfCoreNetworkPolicyDocument_MatchConditionsPropertyList
	// Experimental.
	MatchConditionsInput() interface{}
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
	PutAction(value *DataTfCoreNetworkPolicyDocument_RoutingPoliciesRoutingPolicyRulesRuleDefinitionActionProperty)
	// Experimental.
	PutMatchConditions(value interface{})
	// Experimental.
	ResetConditionLogic()
	// Experimental.
	ResetMatchConditions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference
type jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) Action() DataTfCoreNetworkPolicyDocument_RoutingPoliciesRoutingPolicyRulesRuleDefinitionActionPropertyOutputReference {
	var returns DataTfCoreNetworkPolicyDocument_RoutingPoliciesRoutingPolicyRulesRuleDefinitionActionPropertyOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) ActionInput() *DataTfCoreNetworkPolicyDocument_RoutingPoliciesRoutingPolicyRulesRuleDefinitionActionProperty {
	var returns *DataTfCoreNetworkPolicyDocument_RoutingPoliciesRoutingPolicyRulesRuleDefinitionActionProperty
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) ConditionLogic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionLogic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) ConditionLogicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"conditionLogicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) InternalValue() *DataTfCoreNetworkPolicyDocument_RuleDefinitionProperty {
	var returns *DataTfCoreNetworkPolicyDocument_RuleDefinitionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) MatchConditions() DataTfCoreNetworkPolicyDocument_MatchConditionsPropertyList {
	var returns DataTfCoreNetworkPolicyDocument_MatchConditionsPropertyList
	_jsii_.Get(
		j,
		"matchConditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) MatchConditionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchConditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-manager.DataTfCoreNetworkPolicyDocument.RuleDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference_Override(d DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-manager.DataTfCoreNetworkPolicyDocument.RuleDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference)SetConditionLogic(val *string) {
	if err := j.validateSetConditionLogicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conditionLogic",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference)SetInternalValue(val *DataTfCoreNetworkPolicyDocument_RuleDefinitionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) PutAction(value *DataTfCoreNetworkPolicyDocument_RoutingPoliciesRoutingPolicyRulesRuleDefinitionActionProperty) {
	if err := d.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAction",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) PutMatchConditions(value interface{}) {
	if err := d.validatePutMatchConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMatchConditions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) ResetConditionLogic() {
	_jsii_.InvokeVoid(
		d,
		"resetConditionLogic",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) ResetMatchConditions() {
	_jsii_.InvokeVoid(
		d,
		"resetMatchConditions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_RuleDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

