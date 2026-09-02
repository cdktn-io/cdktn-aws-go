package awssesmailmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssesmailmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssesmailmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRuleSet_UnlessPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BooleanExpression() TfRuleSet_RuleUnlessBooleanExpressionPropertyList
	// Experimental.
	BooleanExpressionInput() interface{}
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
	DmarcExpression() TfRuleSet_RuleUnlessDmarcExpressionPropertyList
	// Experimental.
	DmarcExpressionInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IpExpression() TfRuleSet_RuleUnlessIpExpressionPropertyList
	// Experimental.
	IpExpressionInput() interface{}
	// Experimental.
	NumberExpression() TfRuleSet_RuleUnlessNumberExpressionPropertyList
	// Experimental.
	NumberExpressionInput() interface{}
	// Experimental.
	StringExpression() TfRuleSet_RuleUnlessStringExpressionPropertyList
	// Experimental.
	StringExpressionInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VerdictExpression() TfRuleSet_RuleUnlessVerdictExpressionPropertyList
	// Experimental.
	VerdictExpressionInput() interface{}
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
	PutBooleanExpression(value interface{})
	// Experimental.
	PutDmarcExpression(value interface{})
	// Experimental.
	PutIpExpression(value interface{})
	// Experimental.
	PutNumberExpression(value interface{})
	// Experimental.
	PutStringExpression(value interface{})
	// Experimental.
	PutVerdictExpression(value interface{})
	// Experimental.
	ResetBooleanExpression()
	// Experimental.
	ResetDmarcExpression()
	// Experimental.
	ResetIpExpression()
	// Experimental.
	ResetNumberExpression()
	// Experimental.
	ResetStringExpression()
	// Experimental.
	ResetVerdictExpression()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfRuleSet_UnlessPropertyOutputReference
type jsiiProxy_TfRuleSet_UnlessPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) BooleanExpression() TfRuleSet_RuleUnlessBooleanExpressionPropertyList {
	var returns TfRuleSet_RuleUnlessBooleanExpressionPropertyList
	_jsii_.Get(
		j,
		"booleanExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) BooleanExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"booleanExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) DmarcExpression() TfRuleSet_RuleUnlessDmarcExpressionPropertyList {
	var returns TfRuleSet_RuleUnlessDmarcExpressionPropertyList
	_jsii_.Get(
		j,
		"dmarcExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) DmarcExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dmarcExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) IpExpression() TfRuleSet_RuleUnlessIpExpressionPropertyList {
	var returns TfRuleSet_RuleUnlessIpExpressionPropertyList
	_jsii_.Get(
		j,
		"ipExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) IpExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) NumberExpression() TfRuleSet_RuleUnlessNumberExpressionPropertyList {
	var returns TfRuleSet_RuleUnlessNumberExpressionPropertyList
	_jsii_.Get(
		j,
		"numberExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) NumberExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) StringExpression() TfRuleSet_RuleUnlessStringExpressionPropertyList {
	var returns TfRuleSet_RuleUnlessStringExpressionPropertyList
	_jsii_.Get(
		j,
		"stringExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) StringExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) VerdictExpression() TfRuleSet_RuleUnlessVerdictExpressionPropertyList {
	var returns TfRuleSet_RuleUnlessVerdictExpressionPropertyList
	_jsii_.Get(
		j,
		"verdictExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) VerdictExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verdictExpressionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRuleSet_UnlessPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfRuleSet_UnlessPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRuleSet_UnlessPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRuleSet_UnlessPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.TfRuleSet.UnlessPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRuleSet_UnlessPropertyOutputReference_Override(t TfRuleSet_UnlessPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.TfRuleSet.UnlessPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) PutBooleanExpression(value interface{}) {
	if err := t.validatePutBooleanExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBooleanExpression",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) PutDmarcExpression(value interface{}) {
	if err := t.validatePutDmarcExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDmarcExpression",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) PutIpExpression(value interface{}) {
	if err := t.validatePutIpExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIpExpression",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) PutNumberExpression(value interface{}) {
	if err := t.validatePutNumberExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNumberExpression",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) PutStringExpression(value interface{}) {
	if err := t.validatePutStringExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStringExpression",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) PutVerdictExpression(value interface{}) {
	if err := t.validatePutVerdictExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVerdictExpression",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) ResetBooleanExpression() {
	_jsii_.InvokeVoid(
		t,
		"resetBooleanExpression",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) ResetDmarcExpression() {
	_jsii_.InvokeVoid(
		t,
		"resetDmarcExpression",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) ResetIpExpression() {
	_jsii_.InvokeVoid(
		t,
		"resetIpExpression",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) ResetNumberExpression() {
	_jsii_.InvokeVoid(
		t,
		"resetNumberExpression",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) ResetStringExpression() {
	_jsii_.InvokeVoid(
		t,
		"resetStringExpression",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) ResetVerdictExpression() {
	_jsii_.InvokeVoid(
		t,
		"resetVerdictExpression",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRuleSet_UnlessPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

