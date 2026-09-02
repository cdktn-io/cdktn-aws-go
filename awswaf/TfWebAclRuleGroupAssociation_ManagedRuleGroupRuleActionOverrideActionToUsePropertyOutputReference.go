package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswaf/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswaf/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Allow() TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseAllowPropertyList
	// Experimental.
	AllowInput() interface{}
	// Experimental.
	Block() TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockPropertyList
	// Experimental.
	BlockInput() interface{}
	// Experimental.
	Captcha() TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseCaptchaPropertyList
	// Experimental.
	CaptchaInput() interface{}
	// Experimental.
	Challenge() TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseChallengePropertyList
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
	Count() TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseCountPropertyList
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

// The jsii proxy struct for TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference
type jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) Allow() TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseAllowPropertyList {
	var returns TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseAllowPropertyList
	_jsii_.Get(
		j,
		"allow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) AllowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) Block() TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockPropertyList {
	var returns TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseBlockPropertyList
	_jsii_.Get(
		j,
		"block",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) BlockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) Captcha() TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseCaptchaPropertyList {
	var returns TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseCaptchaPropertyList
	_jsii_.Get(
		j,
		"captcha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) CaptchaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"captchaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) Challenge() TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseChallengePropertyList {
	var returns TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseChallengePropertyList
	_jsii_.Get(
		j,
		"challenge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) ChallengeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"challengeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) Count() TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseCountPropertyList {
	var returns TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUseCountPropertyList
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) CountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRuleGroupAssociation.ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference_Override(t TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf.TfWebAclRuleGroupAssociation.ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) PutAllow(value interface{}) {
	if err := t.validatePutAllowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAllow",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) PutBlock(value interface{}) {
	if err := t.validatePutBlockParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBlock",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) PutCaptcha(value interface{}) {
	if err := t.validatePutCaptchaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCaptcha",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) PutChallenge(value interface{}) {
	if err := t.validatePutChallengeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putChallenge",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) PutCount(value interface{}) {
	if err := t.validatePutCountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCount",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) ResetAllow() {
	_jsii_.InvokeVoid(
		t,
		"resetAllow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) ResetBlock() {
	_jsii_.InvokeVoid(
		t,
		"resetBlock",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) ResetCaptcha() {
	_jsii_.InvokeVoid(
		t,
		"resetCaptcha",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) ResetChallenge() {
	_jsii_.InvokeVoid(
		t,
		"resetChallenge",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) ResetCount() {
	_jsii_.InvokeVoid(
		t,
		"resetCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_ManagedRuleGroupRuleActionOverrideActionToUsePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

