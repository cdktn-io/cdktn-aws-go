package awssesmailmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssesmailmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssesmailmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMailmanagerRuleSet_UnlessPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BooleanExpression() AwsMailmanagerRuleSet_RuleUnlessBooleanExpressionPropertyList
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
	DmarcExpression() AwsMailmanagerRuleSet_RuleUnlessDmarcExpressionPropertyList
	// Experimental.
	DmarcExpressionInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IpExpression() AwsMailmanagerRuleSet_RuleUnlessIpExpressionPropertyList
	// Experimental.
	IpExpressionInput() interface{}
	// Experimental.
	NumberExpression() AwsMailmanagerRuleSet_RuleUnlessNumberExpressionPropertyList
	// Experimental.
	NumberExpressionInput() interface{}
	// Experimental.
	StringExpression() AwsMailmanagerRuleSet_RuleUnlessStringExpressionPropertyList
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
	VerdictExpression() AwsMailmanagerRuleSet_RuleUnlessVerdictExpressionPropertyList
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

// The jsii proxy struct for AwsMailmanagerRuleSet_UnlessPropertyOutputReference
type jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) BooleanExpression() AwsMailmanagerRuleSet_RuleUnlessBooleanExpressionPropertyList {
	var returns AwsMailmanagerRuleSet_RuleUnlessBooleanExpressionPropertyList
	_jsii_.Get(
		j,
		"booleanExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) BooleanExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"booleanExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) DmarcExpression() AwsMailmanagerRuleSet_RuleUnlessDmarcExpressionPropertyList {
	var returns AwsMailmanagerRuleSet_RuleUnlessDmarcExpressionPropertyList
	_jsii_.Get(
		j,
		"dmarcExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) DmarcExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dmarcExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) IpExpression() AwsMailmanagerRuleSet_RuleUnlessIpExpressionPropertyList {
	var returns AwsMailmanagerRuleSet_RuleUnlessIpExpressionPropertyList
	_jsii_.Get(
		j,
		"ipExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) IpExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) NumberExpression() AwsMailmanagerRuleSet_RuleUnlessNumberExpressionPropertyList {
	var returns AwsMailmanagerRuleSet_RuleUnlessNumberExpressionPropertyList
	_jsii_.Get(
		j,
		"numberExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) NumberExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) StringExpression() AwsMailmanagerRuleSet_RuleUnlessStringExpressionPropertyList {
	var returns AwsMailmanagerRuleSet_RuleUnlessStringExpressionPropertyList
	_jsii_.Get(
		j,
		"stringExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) StringExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) VerdictExpression() AwsMailmanagerRuleSet_RuleUnlessVerdictExpressionPropertyList {
	var returns AwsMailmanagerRuleSet_RuleUnlessVerdictExpressionPropertyList
	_jsii_.Get(
		j,
		"verdictExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) VerdictExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verdictExpressionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMailmanagerRuleSet_UnlessPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsMailmanagerRuleSet_UnlessPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMailmanagerRuleSet_UnlessPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.AwsMailmanagerRuleSet.UnlessPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMailmanagerRuleSet_UnlessPropertyOutputReference_Override(a AwsMailmanagerRuleSet_UnlessPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ses-mail-manager.AwsMailmanagerRuleSet.UnlessPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) PutBooleanExpression(value interface{}) {
	if err := a.validatePutBooleanExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBooleanExpression",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) PutDmarcExpression(value interface{}) {
	if err := a.validatePutDmarcExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDmarcExpression",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) PutIpExpression(value interface{}) {
	if err := a.validatePutIpExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIpExpression",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) PutNumberExpression(value interface{}) {
	if err := a.validatePutNumberExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNumberExpression",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) PutStringExpression(value interface{}) {
	if err := a.validatePutStringExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStringExpression",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) PutVerdictExpression(value interface{}) {
	if err := a.validatePutVerdictExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVerdictExpression",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) ResetBooleanExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetBooleanExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) ResetDmarcExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetDmarcExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) ResetIpExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetIpExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) ResetNumberExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetNumberExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) ResetStringExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetStringExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) ResetVerdictExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetVerdictExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMailmanagerRuleSet_UnlessPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

