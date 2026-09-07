package waf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/waf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/waf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRuleGroup_ActionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Allow() AwsRuleGroup_AllowPropertyOutputReference
	// Experimental.
	AllowInput() *AwsRuleGroup_AllowProperty
	// Experimental.
	Block() AwsRuleGroup_BlockPropertyOutputReference
	// Experimental.
	BlockInput() *AwsRuleGroup_BlockProperty
	// Experimental.
	Captcha() AwsRuleGroup_CaptchaPropertyOutputReference
	// Experimental.
	CaptchaInput() *AwsRuleGroup_CaptchaProperty
	// Experimental.
	Challenge() AwsRuleGroup_ChallengePropertyOutputReference
	// Experimental.
	ChallengeInput() *AwsRuleGroup_ChallengeProperty
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
	Count() AwsRuleGroup_CountPropertyOutputReference
	// Experimental.
	CountInput() *AwsRuleGroup_CountProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsRuleGroup_ActionProperty
	// Experimental.
	SetInternalValue(val *AwsRuleGroup_ActionProperty)
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
	PutAllow(value *AwsRuleGroup_AllowProperty)
	// Experimental.
	PutBlock(value *AwsRuleGroup_BlockProperty)
	// Experimental.
	PutCaptcha(value *AwsRuleGroup_CaptchaProperty)
	// Experimental.
	PutChallenge(value *AwsRuleGroup_ChallengeProperty)
	// Experimental.
	PutCount(value *AwsRuleGroup_CountProperty)
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

// The jsii proxy struct for AwsRuleGroup_ActionPropertyOutputReference
type jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) Allow() AwsRuleGroup_AllowPropertyOutputReference {
	var returns AwsRuleGroup_AllowPropertyOutputReference
	_jsii_.Get(
		j,
		"allow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) AllowInput() *AwsRuleGroup_AllowProperty {
	var returns *AwsRuleGroup_AllowProperty
	_jsii_.Get(
		j,
		"allowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) Block() AwsRuleGroup_BlockPropertyOutputReference {
	var returns AwsRuleGroup_BlockPropertyOutputReference
	_jsii_.Get(
		j,
		"block",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) BlockInput() *AwsRuleGroup_BlockProperty {
	var returns *AwsRuleGroup_BlockProperty
	_jsii_.Get(
		j,
		"blockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) Captcha() AwsRuleGroup_CaptchaPropertyOutputReference {
	var returns AwsRuleGroup_CaptchaPropertyOutputReference
	_jsii_.Get(
		j,
		"captcha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) CaptchaInput() *AwsRuleGroup_CaptchaProperty {
	var returns *AwsRuleGroup_CaptchaProperty
	_jsii_.Get(
		j,
		"captchaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) Challenge() AwsRuleGroup_ChallengePropertyOutputReference {
	var returns AwsRuleGroup_ChallengePropertyOutputReference
	_jsii_.Get(
		j,
		"challenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) ChallengeInput() *AwsRuleGroup_ChallengeProperty {
	var returns *AwsRuleGroup_ChallengeProperty
	_jsii_.Get(
		j,
		"challengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) Count() AwsRuleGroup_CountPropertyOutputReference {
	var returns AwsRuleGroup_CountPropertyOutputReference
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) CountInput() *AwsRuleGroup_CountProperty {
	var returns *AwsRuleGroup_CountProperty
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) InternalValue() *AwsRuleGroup_ActionProperty {
	var returns *AwsRuleGroup_ActionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRuleGroup_ActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsRuleGroup_ActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRuleGroup_ActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.AwsRuleGroup.ActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRuleGroup_ActionPropertyOutputReference_Override(a AwsRuleGroup_ActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.AwsRuleGroup.ActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference)SetInternalValue(val *AwsRuleGroup_ActionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) PutAllow(value *AwsRuleGroup_AllowProperty) {
	if err := a.validatePutAllowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAllow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) PutBlock(value *AwsRuleGroup_BlockProperty) {
	if err := a.validatePutBlockParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBlock",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) PutCaptcha(value *AwsRuleGroup_CaptchaProperty) {
	if err := a.validatePutCaptchaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCaptcha",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) PutChallenge(value *AwsRuleGroup_ChallengeProperty) {
	if err := a.validatePutChallengeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putChallenge",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) PutCount(value *AwsRuleGroup_CountProperty) {
	if err := a.validatePutCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCount",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) ResetAllow() {
	_jsii_.InvokeVoid(
		a,
		"resetAllow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) ResetBlock() {
	_jsii_.InvokeVoid(
		a,
		"resetBlock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) ResetCaptcha() {
	_jsii_.InvokeVoid(
		a,
		"resetCaptcha",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) ResetChallenge() {
	_jsii_.InvokeVoid(
		a,
		"resetChallenge",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) ResetCount() {
	_jsii_.InvokeVoid(
		a,
		"resetCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRuleGroup_ActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

