package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRuleGroup_RulesSourcePropertyOutputReference interface {
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
	InternalValue() *TfRuleGroup_RulesSourceProperty
	// Experimental.
	SetInternalValue(val *TfRuleGroup_RulesSourceProperty)
	// Experimental.
	RulesSourceList() TfRuleGroup_RulesSourceListPropertyOutputReference
	// Experimental.
	RulesSourceListInput() *TfRuleGroup_RulesSourceListProperty
	// Experimental.
	RulesString() *string
	// Experimental.
	SetRulesString(val *string)
	// Experimental.
	RulesStringInput() *string
	// Experimental.
	StatefulRule() TfRuleGroup_StatefulRulePropertyList
	// Experimental.
	StatefulRuleInput() interface{}
	// Experimental.
	StatelessRulesAndCustomActions() TfRuleGroup_StatelessRulesAndCustomActionsPropertyOutputReference
	// Experimental.
	StatelessRulesAndCustomActionsInput() *TfRuleGroup_StatelessRulesAndCustomActionsProperty
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
	PutRulesSourceList(value *TfRuleGroup_RulesSourceListProperty)
	// Experimental.
	PutStatefulRule(value interface{})
	// Experimental.
	PutStatelessRulesAndCustomActions(value *TfRuleGroup_StatelessRulesAndCustomActionsProperty)
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

// The jsii proxy struct for TfRuleGroup_RulesSourcePropertyOutputReference
type jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) InternalValue() *TfRuleGroup_RulesSourceProperty {
	var returns *TfRuleGroup_RulesSourceProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) RulesSourceList() TfRuleGroup_RulesSourceListPropertyOutputReference {
	var returns TfRuleGroup_RulesSourceListPropertyOutputReference
	_jsii_.Get(
		j,
		"rulesSourceList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) RulesSourceListInput() *TfRuleGroup_RulesSourceListProperty {
	var returns *TfRuleGroup_RulesSourceListProperty
	_jsii_.Get(
		j,
		"rulesSourceListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) RulesString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rulesString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) RulesStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rulesStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) StatefulRule() TfRuleGroup_StatefulRulePropertyList {
	var returns TfRuleGroup_StatefulRulePropertyList
	_jsii_.Get(
		j,
		"statefulRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) StatefulRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) StatelessRulesAndCustomActions() TfRuleGroup_StatelessRulesAndCustomActionsPropertyOutputReference {
	var returns TfRuleGroup_StatelessRulesAndCustomActionsPropertyOutputReference
	_jsii_.Get(
		j,
		"statelessRulesAndCustomActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) StatelessRulesAndCustomActionsInput() *TfRuleGroup_StatelessRulesAndCustomActionsProperty {
	var returns *TfRuleGroup_StatelessRulesAndCustomActionsProperty
	_jsii_.Get(
		j,
		"statelessRulesAndCustomActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRuleGroup_RulesSourcePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfRuleGroup_RulesSourcePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRuleGroup_RulesSourcePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-firewall.TfRuleGroup.RulesSourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRuleGroup_RulesSourcePropertyOutputReference_Override(t TfRuleGroup_RulesSourcePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-firewall.TfRuleGroup.RulesSourcePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference)SetInternalValue(val *TfRuleGroup_RulesSourceProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference)SetRulesString(val *string) {
	if err := j.validateSetRulesStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rulesString",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) PutRulesSourceList(value *TfRuleGroup_RulesSourceListProperty) {
	if err := t.validatePutRulesSourceListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRulesSourceList",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) PutStatefulRule(value interface{}) {
	if err := t.validatePutStatefulRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStatefulRule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) PutStatelessRulesAndCustomActions(value *TfRuleGroup_StatelessRulesAndCustomActionsProperty) {
	if err := t.validatePutStatelessRulesAndCustomActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStatelessRulesAndCustomActions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) ResetRulesSourceList() {
	_jsii_.InvokeVoid(
		t,
		"resetRulesSourceList",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) ResetRulesString() {
	_jsii_.InvokeVoid(
		t,
		"resetRulesString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) ResetStatefulRule() {
	_jsii_.InvokeVoid(
		t,
		"resetStatefulRule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) ResetStatelessRulesAndCustomActions() {
	_jsii_.InvokeVoid(
		t,
		"resetStatelessRulesAndCustomActions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRuleGroup_RulesSourcePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

