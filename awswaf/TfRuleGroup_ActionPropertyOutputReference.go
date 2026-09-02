package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRuleGroup_ActionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Allow() TfRuleGroup_AllowPropertyOutputReference
	// Experimental.
	AllowInput() *TfRuleGroup_AllowProperty
	// Experimental.
	Block() TfRuleGroup_BlockPropertyOutputReference
	// Experimental.
	BlockInput() *TfRuleGroup_BlockProperty
	// Experimental.
	Captcha() TfRuleGroup_CaptchaPropertyOutputReference
	// Experimental.
	CaptchaInput() *TfRuleGroup_CaptchaProperty
	// Experimental.
	Challenge() TfRuleGroup_ChallengePropertyOutputReference
	// Experimental.
	ChallengeInput() *TfRuleGroup_ChallengeProperty
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
	Count() TfRuleGroup_CountPropertyOutputReference
	// Experimental.
	CountInput() *TfRuleGroup_CountProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfRuleGroup_ActionProperty
	// Experimental.
	SetInternalValue(val *TfRuleGroup_ActionProperty)
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
	PutAllow(value *TfRuleGroup_AllowProperty)
	// Experimental.
	PutBlock(value *TfRuleGroup_BlockProperty)
	// Experimental.
	PutCaptcha(value *TfRuleGroup_CaptchaProperty)
	// Experimental.
	PutChallenge(value *TfRuleGroup_ChallengeProperty)
	// Experimental.
	PutCount(value *TfRuleGroup_CountProperty)
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

// The jsii proxy struct for TfRuleGroup_ActionPropertyOutputReference
type jsiiProxy_TfRuleGroup_ActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) Allow() TfRuleGroup_AllowPropertyOutputReference {
	var returns TfRuleGroup_AllowPropertyOutputReference
	_jsii_.Get(
		j,
		"allow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) AllowInput() *TfRuleGroup_AllowProperty {
	var returns *TfRuleGroup_AllowProperty
	_jsii_.Get(
		j,
		"allowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) Block() TfRuleGroup_BlockPropertyOutputReference {
	var returns TfRuleGroup_BlockPropertyOutputReference
	_jsii_.Get(
		j,
		"block",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) BlockInput() *TfRuleGroup_BlockProperty {
	var returns *TfRuleGroup_BlockProperty
	_jsii_.Get(
		j,
		"blockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) Captcha() TfRuleGroup_CaptchaPropertyOutputReference {
	var returns TfRuleGroup_CaptchaPropertyOutputReference
	_jsii_.Get(
		j,
		"captcha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) CaptchaInput() *TfRuleGroup_CaptchaProperty {
	var returns *TfRuleGroup_CaptchaProperty
	_jsii_.Get(
		j,
		"captchaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) Challenge() TfRuleGroup_ChallengePropertyOutputReference {
	var returns TfRuleGroup_ChallengePropertyOutputReference
	_jsii_.Get(
		j,
		"challenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) ChallengeInput() *TfRuleGroup_ChallengeProperty {
	var returns *TfRuleGroup_ChallengeProperty
	_jsii_.Get(
		j,
		"challengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) Count() TfRuleGroup_CountPropertyOutputReference {
	var returns TfRuleGroup_CountPropertyOutputReference
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) CountInput() *TfRuleGroup_CountProperty {
	var returns *TfRuleGroup_CountProperty
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) InternalValue() *TfRuleGroup_ActionProperty {
	var returns *TfRuleGroup_ActionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRuleGroup_ActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfRuleGroup_ActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRuleGroup_ActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRuleGroup_ActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfRuleGroup.ActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRuleGroup_ActionPropertyOutputReference_Override(t TfRuleGroup_ActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfRuleGroup.ActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference)SetInternalValue(val *TfRuleGroup_ActionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) PutAllow(value *TfRuleGroup_AllowProperty) {
	if err := t.validatePutAllowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAllow",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) PutBlock(value *TfRuleGroup_BlockProperty) {
	if err := t.validatePutBlockParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBlock",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) PutCaptcha(value *TfRuleGroup_CaptchaProperty) {
	if err := t.validatePutCaptchaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCaptcha",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) PutChallenge(value *TfRuleGroup_ChallengeProperty) {
	if err := t.validatePutChallengeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putChallenge",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) PutCount(value *TfRuleGroup_CountProperty) {
	if err := t.validatePutCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCount",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) ResetAllow() {
	_jsii_.InvokeVoid(
		t,
		"resetAllow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) ResetBlock() {
	_jsii_.InvokeVoid(
		t,
		"resetBlock",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) ResetCaptcha() {
	_jsii_.InvokeVoid(
		t,
		"resetCaptcha",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) ResetChallenge() {
	_jsii_.InvokeVoid(
		t,
		"resetChallenge",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) ResetCount() {
	_jsii_.InvokeVoid(
		t,
		"resetCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRuleGroup_ActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

