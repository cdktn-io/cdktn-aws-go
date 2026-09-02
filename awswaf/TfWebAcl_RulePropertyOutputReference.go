package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAcl_RulePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Action() TfWebAcl_ActionPropertyOutputReference
	// Experimental.
	ActionInput() *TfWebAcl_ActionProperty
	// Experimental.
	CaptchaConfig() TfWebAcl_RuleCaptchaConfigPropertyOutputReference
	// Experimental.
	CaptchaConfigInput() *TfWebAcl_RuleCaptchaConfigProperty
	// Experimental.
	ChallengeConfig() TfWebAcl_RuleChallengeConfigPropertyOutputReference
	// Experimental.
	ChallengeConfigInput() *TfWebAcl_RuleChallengeConfigProperty
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	OverrideAction() TfWebAcl_OverrideActionPropertyOutputReference
	// Experimental.
	OverrideActionInput() *TfWebAcl_OverrideActionProperty
	// Experimental.
	Priority() *float64
	// Experimental.
	SetPriority(val *float64)
	// Experimental.
	PriorityInput() *float64
	// Experimental.
	RuleLabel() TfWebAcl_RuleLabelPropertyList
	// Experimental.
	RuleLabelInput() interface{}
	// Experimental.
	Statement() interface{}
	// Experimental.
	SetStatement(val interface{})
	// Experimental.
	StatementInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VisibilityConfig() TfWebAcl_RuleVisibilityConfigPropertyOutputReference
	// Experimental.
	VisibilityConfigInput() *TfWebAcl_RuleVisibilityConfigProperty
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
	PutAction(value *TfWebAcl_ActionProperty)
	// Experimental.
	PutCaptchaConfig(value *TfWebAcl_RuleCaptchaConfigProperty)
	// Experimental.
	PutChallengeConfig(value *TfWebAcl_RuleChallengeConfigProperty)
	// Experimental.
	PutOverrideAction(value *TfWebAcl_OverrideActionProperty)
	// Experimental.
	PutRuleLabel(value interface{})
	// Experimental.
	PutVisibilityConfig(value *TfWebAcl_RuleVisibilityConfigProperty)
	// Experimental.
	ResetAction()
	// Experimental.
	ResetCaptchaConfig()
	// Experimental.
	ResetChallengeConfig()
	// Experimental.
	ResetOverrideAction()
	// Experimental.
	ResetRuleLabel()
	// Experimental.
	ResetStatement()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWebAcl_RulePropertyOutputReference
type jsiiProxy_TfWebAcl_RulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) Action() TfWebAcl_ActionPropertyOutputReference {
	var returns TfWebAcl_ActionPropertyOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) ActionInput() *TfWebAcl_ActionProperty {
	var returns *TfWebAcl_ActionProperty
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) CaptchaConfig() TfWebAcl_RuleCaptchaConfigPropertyOutputReference {
	var returns TfWebAcl_RuleCaptchaConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"captchaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) CaptchaConfigInput() *TfWebAcl_RuleCaptchaConfigProperty {
	var returns *TfWebAcl_RuleCaptchaConfigProperty
	_jsii_.Get(
		j,
		"captchaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) ChallengeConfig() TfWebAcl_RuleChallengeConfigPropertyOutputReference {
	var returns TfWebAcl_RuleChallengeConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"challengeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) ChallengeConfigInput() *TfWebAcl_RuleChallengeConfigProperty {
	var returns *TfWebAcl_RuleChallengeConfigProperty
	_jsii_.Get(
		j,
		"challengeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) OverrideAction() TfWebAcl_OverrideActionPropertyOutputReference {
	var returns TfWebAcl_OverrideActionPropertyOutputReference
	_jsii_.Get(
		j,
		"overrideAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) OverrideActionInput() *TfWebAcl_OverrideActionProperty {
	var returns *TfWebAcl_OverrideActionProperty
	_jsii_.Get(
		j,
		"overrideActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) RuleLabel() TfWebAcl_RuleLabelPropertyList {
	var returns TfWebAcl_RuleLabelPropertyList
	_jsii_.Get(
		j,
		"ruleLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) RuleLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) Statement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) StatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) VisibilityConfig() TfWebAcl_RuleVisibilityConfigPropertyOutputReference {
	var returns TfWebAcl_RuleVisibilityConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"visibilityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference) VisibilityConfigInput() *TfWebAcl_RuleVisibilityConfigProperty {
	var returns *TfWebAcl_RuleVisibilityConfigProperty
	_jsii_.Get(
		j,
		"visibilityConfigInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAcl_RulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAcl_RulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAcl_RulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAcl_RulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAcl.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAcl_RulePropertyOutputReference_Override(t TfWebAcl_RulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAcl.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference)SetStatement(val interface{}) {
	if err := j.validateSetStatementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statement",
		val,
	)
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAcl_RulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) PutAction(value *TfWebAcl_ActionProperty) {
	if err := t.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAction",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) PutCaptchaConfig(value *TfWebAcl_RuleCaptchaConfigProperty) {
	if err := t.validatePutCaptchaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCaptchaConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) PutChallengeConfig(value *TfWebAcl_RuleChallengeConfigProperty) {
	if err := t.validatePutChallengeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putChallengeConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) PutOverrideAction(value *TfWebAcl_OverrideActionProperty) {
	if err := t.validatePutOverrideActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOverrideAction",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) PutRuleLabel(value interface{}) {
	if err := t.validatePutRuleLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRuleLabel",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) PutVisibilityConfig(value *TfWebAcl_RuleVisibilityConfigProperty) {
	if err := t.validatePutVisibilityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVisibilityConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		t,
		"resetAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) ResetCaptchaConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCaptchaConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) ResetChallengeConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetChallengeConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) ResetOverrideAction() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) ResetRuleLabel() {
	_jsii_.InvokeVoid(
		t,
		"resetRuleLabel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) ResetStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAcl_RulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

