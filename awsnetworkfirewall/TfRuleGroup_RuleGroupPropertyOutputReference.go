package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRuleGroup_RuleGroupPropertyOutputReference interface {
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
	InternalValue() *TfRuleGroup_RuleGroupProperty
	// Experimental.
	SetInternalValue(val *TfRuleGroup_RuleGroupProperty)
	// Experimental.
	ReferenceSets() TfRuleGroup_ReferenceSetsPropertyOutputReference
	// Experimental.
	ReferenceSetsInput() *TfRuleGroup_ReferenceSetsProperty
	// Experimental.
	RulesSource() TfRuleGroup_RulesSourcePropertyOutputReference
	// Experimental.
	RulesSourceInput() *TfRuleGroup_RulesSourceProperty
	// Experimental.
	RuleVariables() TfRuleGroup_RuleVariablesPropertyOutputReference
	// Experimental.
	RuleVariablesInput() *TfRuleGroup_RuleVariablesProperty
	// Experimental.
	StatefulRuleOptions() TfRuleGroup_StatefulRuleOptionsPropertyOutputReference
	// Experimental.
	StatefulRuleOptionsInput() *TfRuleGroup_StatefulRuleOptionsProperty
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
	PutReferenceSets(value *TfRuleGroup_ReferenceSetsProperty)
	// Experimental.
	PutRulesSource(value *TfRuleGroup_RulesSourceProperty)
	// Experimental.
	PutRuleVariables(value *TfRuleGroup_RuleVariablesProperty)
	// Experimental.
	PutStatefulRuleOptions(value *TfRuleGroup_StatefulRuleOptionsProperty)
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

// The jsii proxy struct for TfRuleGroup_RuleGroupPropertyOutputReference
type jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) InternalValue() *TfRuleGroup_RuleGroupProperty {
	var returns *TfRuleGroup_RuleGroupProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) ReferenceSets() TfRuleGroup_ReferenceSetsPropertyOutputReference {
	var returns TfRuleGroup_ReferenceSetsPropertyOutputReference
	_jsii_.Get(
		j,
		"referenceSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) ReferenceSetsInput() *TfRuleGroup_ReferenceSetsProperty {
	var returns *TfRuleGroup_ReferenceSetsProperty
	_jsii_.Get(
		j,
		"referenceSetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) RulesSource() TfRuleGroup_RulesSourcePropertyOutputReference {
	var returns TfRuleGroup_RulesSourcePropertyOutputReference
	_jsii_.Get(
		j,
		"rulesSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) RulesSourceInput() *TfRuleGroup_RulesSourceProperty {
	var returns *TfRuleGroup_RulesSourceProperty
	_jsii_.Get(
		j,
		"rulesSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) RuleVariables() TfRuleGroup_RuleVariablesPropertyOutputReference {
	var returns TfRuleGroup_RuleVariablesPropertyOutputReference
	_jsii_.Get(
		j,
		"ruleVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) RuleVariablesInput() *TfRuleGroup_RuleVariablesProperty {
	var returns *TfRuleGroup_RuleVariablesProperty
	_jsii_.Get(
		j,
		"ruleVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) StatefulRuleOptions() TfRuleGroup_StatefulRuleOptionsPropertyOutputReference {
	var returns TfRuleGroup_StatefulRuleOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"statefulRuleOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) StatefulRuleOptionsInput() *TfRuleGroup_StatefulRuleOptionsProperty {
	var returns *TfRuleGroup_StatefulRuleOptionsProperty
	_jsii_.Get(
		j,
		"statefulRuleOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRuleGroup_RuleGroupPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfRuleGroup_RuleGroupPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRuleGroup_RuleGroupPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-firewall.TfRuleGroup.RuleGroupPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRuleGroup_RuleGroupPropertyOutputReference_Override(t TfRuleGroup_RuleGroupPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-firewall.TfRuleGroup.RuleGroupPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference)SetInternalValue(val *TfRuleGroup_RuleGroupProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) PutReferenceSets(value *TfRuleGroup_ReferenceSetsProperty) {
	if err := t.validatePutReferenceSetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putReferenceSets",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) PutRulesSource(value *TfRuleGroup_RulesSourceProperty) {
	if err := t.validatePutRulesSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRulesSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) PutRuleVariables(value *TfRuleGroup_RuleVariablesProperty) {
	if err := t.validatePutRuleVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRuleVariables",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) PutStatefulRuleOptions(value *TfRuleGroup_StatefulRuleOptionsProperty) {
	if err := t.validatePutStatefulRuleOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStatefulRuleOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) ResetReferenceSets() {
	_jsii_.InvokeVoid(
		t,
		"resetReferenceSets",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) ResetRuleVariables() {
	_jsii_.InvokeVoid(
		t,
		"resetRuleVariables",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) ResetStatefulRuleOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetStatefulRuleOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRuleGroup_RuleGroupPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

