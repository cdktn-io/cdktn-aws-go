package waf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/waf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/waf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWebAcl_RulePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Action() AwsWebAcl_ActionPropertyOutputReference
	// Experimental.
	ActionInput() *AwsWebAcl_ActionProperty
	// Experimental.
	CaptchaConfig() AwsWebAcl_RuleCaptchaConfigPropertyOutputReference
	// Experimental.
	CaptchaConfigInput() *AwsWebAcl_RuleCaptchaConfigProperty
	// Experimental.
	ChallengeConfig() AwsWebAcl_RuleChallengeConfigPropertyOutputReference
	// Experimental.
	ChallengeConfigInput() *AwsWebAcl_RuleChallengeConfigProperty
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
	OverrideAction() AwsWebAcl_OverrideActionPropertyOutputReference
	// Experimental.
	OverrideActionInput() *AwsWebAcl_OverrideActionProperty
	// Experimental.
	Priority() *float64
	// Experimental.
	SetPriority(val *float64)
	// Experimental.
	PriorityInput() *float64
	// Experimental.
	RuleLabel() AwsWebAcl_RuleLabelPropertyList
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
	VisibilityConfig() AwsWebAcl_RuleVisibilityConfigPropertyOutputReference
	// Experimental.
	VisibilityConfigInput() *AwsWebAcl_RuleVisibilityConfigProperty
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
	PutAction(value *AwsWebAcl_ActionProperty)
	// Experimental.
	PutCaptchaConfig(value *AwsWebAcl_RuleCaptchaConfigProperty)
	// Experimental.
	PutChallengeConfig(value *AwsWebAcl_RuleChallengeConfigProperty)
	// Experimental.
	PutOverrideAction(value *AwsWebAcl_OverrideActionProperty)
	// Experimental.
	PutRuleLabel(value interface{})
	// Experimental.
	PutVisibilityConfig(value *AwsWebAcl_RuleVisibilityConfigProperty)
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

// The jsii proxy struct for AwsWebAcl_RulePropertyOutputReference
type jsiiProxy_AwsWebAcl_RulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) Action() AwsWebAcl_ActionPropertyOutputReference {
	var returns AwsWebAcl_ActionPropertyOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) ActionInput() *AwsWebAcl_ActionProperty {
	var returns *AwsWebAcl_ActionProperty
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) CaptchaConfig() AwsWebAcl_RuleCaptchaConfigPropertyOutputReference {
	var returns AwsWebAcl_RuleCaptchaConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"captchaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) CaptchaConfigInput() *AwsWebAcl_RuleCaptchaConfigProperty {
	var returns *AwsWebAcl_RuleCaptchaConfigProperty
	_jsii_.Get(
		j,
		"captchaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) ChallengeConfig() AwsWebAcl_RuleChallengeConfigPropertyOutputReference {
	var returns AwsWebAcl_RuleChallengeConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"challengeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) ChallengeConfigInput() *AwsWebAcl_RuleChallengeConfigProperty {
	var returns *AwsWebAcl_RuleChallengeConfigProperty
	_jsii_.Get(
		j,
		"challengeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) OverrideAction() AwsWebAcl_OverrideActionPropertyOutputReference {
	var returns AwsWebAcl_OverrideActionPropertyOutputReference
	_jsii_.Get(
		j,
		"overrideAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) OverrideActionInput() *AwsWebAcl_OverrideActionProperty {
	var returns *AwsWebAcl_OverrideActionProperty
	_jsii_.Get(
		j,
		"overrideActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) RuleLabel() AwsWebAcl_RuleLabelPropertyList {
	var returns AwsWebAcl_RuleLabelPropertyList
	_jsii_.Get(
		j,
		"ruleLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) RuleLabelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) Statement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) StatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) VisibilityConfig() AwsWebAcl_RuleVisibilityConfigPropertyOutputReference {
	var returns AwsWebAcl_RuleVisibilityConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"visibilityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) VisibilityConfigInput() *AwsWebAcl_RuleVisibilityConfigProperty {
	var returns *AwsWebAcl_RuleVisibilityConfigProperty
	_jsii_.Get(
		j,
		"visibilityConfigInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWebAcl_RulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsWebAcl_RulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWebAcl_RulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWebAcl_RulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAcl.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWebAcl_RulePropertyOutputReference_Override(a AwsWebAcl_RulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAcl.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference)SetStatement(val interface{}) {
	if err := j.validateSetStatementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statement",
		val,
	)
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWebAcl_RulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) PutAction(value *AwsWebAcl_ActionProperty) {
	if err := a.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) PutCaptchaConfig(value *AwsWebAcl_RuleCaptchaConfigProperty) {
	if err := a.validatePutCaptchaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCaptchaConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) PutChallengeConfig(value *AwsWebAcl_RuleChallengeConfigProperty) {
	if err := a.validatePutChallengeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putChallengeConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) PutOverrideAction(value *AwsWebAcl_OverrideActionProperty) {
	if err := a.validatePutOverrideActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOverrideAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) PutRuleLabel(value interface{}) {
	if err := a.validatePutRuleLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRuleLabel",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) PutVisibilityConfig(value *AwsWebAcl_RuleVisibilityConfigProperty) {
	if err := a.validatePutVisibilityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVisibilityConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		a,
		"resetAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) ResetCaptchaConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCaptchaConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) ResetChallengeConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetChallengeConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) ResetOverrideAction() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) ResetRuleLabel() {
	_jsii_.InvokeVoid(
		a,
		"resetRuleLabel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) ResetStatement() {
	_jsii_.InvokeVoid(
		a,
		"resetStatement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWebAcl_RulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

