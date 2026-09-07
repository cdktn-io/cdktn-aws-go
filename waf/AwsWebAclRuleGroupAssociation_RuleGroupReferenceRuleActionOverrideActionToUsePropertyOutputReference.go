package waf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/waf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/waf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Allow() AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseAllowPropertyList
	// Experimental.
	AllowInput() interface{}
	// Experimental.
	Block() AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockPropertyList
	// Experimental.
	BlockInput() interface{}
	// Experimental.
	Captcha() AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCaptchaPropertyList
	// Experimental.
	CaptchaInput() interface{}
	// Experimental.
	Challenge() AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseChallengePropertyList
	// Experimental.
	ChallengeInput() interface{}
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
	Count() AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCountPropertyList
	// Experimental.
	CountInput() interface{}
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
	PutAllow(value interface{})
	// Experimental.
	PutBlock(value interface{})
	// Experimental.
	PutCaptcha(value interface{})
	// Experimental.
	PutChallenge(value interface{})
	// Experimental.
	PutCount(value interface{})
	// Experimental.
	ResetAllow()
	// Experimental.
	ResetBlock()
	// Experimental.
	ResetCaptcha()
	// Experimental.
	ResetChallenge()
	// Experimental.
	ResetCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference
type jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) Allow() AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseAllowPropertyList {
	var returns AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseAllowPropertyList
	_jsii_.Get(
		j,
		"allow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) AllowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) Block() AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockPropertyList {
	var returns AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockPropertyList
	_jsii_.Get(
		j,
		"block",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) BlockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) Captcha() AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCaptchaPropertyList {
	var returns AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCaptchaPropertyList
	_jsii_.Get(
		j,
		"captcha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) CaptchaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captchaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) Challenge() AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseChallengePropertyList {
	var returns AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseChallengePropertyList
	_jsii_.Get(
		j,
		"challenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) ChallengeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"challengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) Count() AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCountPropertyList {
	var returns AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCountPropertyList
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) CountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRuleGroupAssociation.RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference_Override(a AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsWebAclRuleGroupAssociation.RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) PutAllow(value interface{}) {
	if err := a.validatePutAllowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAllow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) PutBlock(value interface{}) {
	if err := a.validatePutBlockParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBlock",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) PutCaptcha(value interface{}) {
	if err := a.validatePutCaptchaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCaptcha",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) PutChallenge(value interface{}) {
	if err := a.validatePutChallengeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putChallenge",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) PutCount(value interface{}) {
	if err := a.validatePutCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCount",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) ResetAllow() {
	_jsii_.InvokeVoid(
		a,
		"resetAllow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) ResetBlock() {
	_jsii_.InvokeVoid(
		a,
		"resetBlock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) ResetCaptcha() {
	_jsii_.InvokeVoid(
		a,
		"resetCaptcha",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) ResetChallenge() {
	_jsii_.InvokeVoid(
		a,
		"resetChallenge",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) ResetCount() {
	_jsii_.InvokeVoid(
		a,
		"resetCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

