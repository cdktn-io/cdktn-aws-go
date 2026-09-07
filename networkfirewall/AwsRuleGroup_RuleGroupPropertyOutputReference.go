package networkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/networkfirewall/jsii"

	"github.com/cdktn-io/cdktn-aws-go/networkfirewall/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRuleGroup_RuleGroupPropertyOutputReference interface {
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
	InternalValue() *AwsRuleGroup_RuleGroupProperty
	// Experimental.
	SetInternalValue(val *AwsRuleGroup_RuleGroupProperty)
	// Experimental.
	ReferenceSets() AwsRuleGroup_ReferenceSetsPropertyOutputReference
	// Experimental.
	ReferenceSetsInput() *AwsRuleGroup_ReferenceSetsProperty
	// Experimental.
	RulesSource() AwsRuleGroup_RulesSourcePropertyOutputReference
	// Experimental.
	RulesSourceInput() *AwsRuleGroup_RulesSourceProperty
	// Experimental.
	RuleVariables() AwsRuleGroup_RuleVariablesPropertyOutputReference
	// Experimental.
	RuleVariablesInput() *AwsRuleGroup_RuleVariablesProperty
	// Experimental.
	StatefulRuleOptions() AwsRuleGroup_StatefulRuleOptionsPropertyOutputReference
	// Experimental.
	StatefulRuleOptionsInput() *AwsRuleGroup_StatefulRuleOptionsProperty
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
	PutReferenceSets(value *AwsRuleGroup_ReferenceSetsProperty)
	// Experimental.
	PutRulesSource(value *AwsRuleGroup_RulesSourceProperty)
	// Experimental.
	PutRuleVariables(value *AwsRuleGroup_RuleVariablesProperty)
	// Experimental.
	PutStatefulRuleOptions(value *AwsRuleGroup_StatefulRuleOptionsProperty)
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

// The jsii proxy struct for AwsRuleGroup_RuleGroupPropertyOutputReference
type jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) InternalValue() *AwsRuleGroup_RuleGroupProperty {
	var returns *AwsRuleGroup_RuleGroupProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) ReferenceSets() AwsRuleGroup_ReferenceSetsPropertyOutputReference {
	var returns AwsRuleGroup_ReferenceSetsPropertyOutputReference
	_jsii_.Get(
		j,
		"referenceSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) ReferenceSetsInput() *AwsRuleGroup_ReferenceSetsProperty {
	var returns *AwsRuleGroup_ReferenceSetsProperty
	_jsii_.Get(
		j,
		"referenceSetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) RulesSource() AwsRuleGroup_RulesSourcePropertyOutputReference {
	var returns AwsRuleGroup_RulesSourcePropertyOutputReference
	_jsii_.Get(
		j,
		"rulesSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) RulesSourceInput() *AwsRuleGroup_RulesSourceProperty {
	var returns *AwsRuleGroup_RulesSourceProperty
	_jsii_.Get(
		j,
		"rulesSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) RuleVariables() AwsRuleGroup_RuleVariablesPropertyOutputReference {
	var returns AwsRuleGroup_RuleVariablesPropertyOutputReference
	_jsii_.Get(
		j,
		"ruleVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) RuleVariablesInput() *AwsRuleGroup_RuleVariablesProperty {
	var returns *AwsRuleGroup_RuleVariablesProperty
	_jsii_.Get(
		j,
		"ruleVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) StatefulRuleOptions() AwsRuleGroup_StatefulRuleOptionsPropertyOutputReference {
	var returns AwsRuleGroup_StatefulRuleOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"statefulRuleOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) StatefulRuleOptionsInput() *AwsRuleGroup_StatefulRuleOptionsProperty {
	var returns *AwsRuleGroup_StatefulRuleOptionsProperty
	_jsii_.Get(
		j,
		"statefulRuleOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRuleGroup_RuleGroupPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsRuleGroup_RuleGroupPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRuleGroup_RuleGroupPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsRuleGroup.RuleGroupPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRuleGroup_RuleGroupPropertyOutputReference_Override(a AwsRuleGroup_RuleGroupPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsRuleGroup.RuleGroupPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference)SetInternalValue(val *AwsRuleGroup_RuleGroupProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) PutReferenceSets(value *AwsRuleGroup_ReferenceSetsProperty) {
	if err := a.validatePutReferenceSetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReferenceSets",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) PutRulesSource(value *AwsRuleGroup_RulesSourceProperty) {
	if err := a.validatePutRulesSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRulesSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) PutRuleVariables(value *AwsRuleGroup_RuleVariablesProperty) {
	if err := a.validatePutRuleVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRuleVariables",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) PutStatefulRuleOptions(value *AwsRuleGroup_StatefulRuleOptionsProperty) {
	if err := a.validatePutStatefulRuleOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStatefulRuleOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) ResetReferenceSets() {
	_jsii_.InvokeVoid(
		a,
		"resetReferenceSets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) ResetRuleVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetRuleVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) ResetStatefulRuleOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetStatefulRuleOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRuleGroup_RuleGroupPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

