package awsdlm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdlm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdlm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CmkArn() *string
	// Experimental.
	SetCmkArn(val *string)
	// Experimental.
	CmkArnInput() *string
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
	CopyTags() interface{}
	// Experimental.
	SetCopyTags(val interface{})
	// Experimental.
	CopyTagsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DeprecateRule() TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleDeprecateRulePropertyOutputReference
	// Experimental.
	DeprecateRuleInput() *TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleProperty
	// Experimental.
	Encrypted() interface{}
	// Experimental.
	SetEncrypted(val interface{})
	// Experimental.
	EncryptedInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	RetainRule() TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleRetainRulePropertyOutputReference
	// Experimental.
	RetainRuleInput() *TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleRetainRuleProperty
	// Experimental.
	Target() *string
	// Experimental.
	SetTarget(val *string)
	// Experimental.
	TargetInput() *string
	// Experimental.
	TargetRegion() *string
	// Experimental.
	SetTargetRegion(val *string)
	// Experimental.
	TargetRegionInput() *string
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
	PutDeprecateRule(value *TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleProperty)
	// Experimental.
	PutRetainRule(value *TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleRetainRuleProperty)
	// Experimental.
	ResetCmkArn()
	// Experimental.
	ResetCopyTags()
	// Experimental.
	ResetDeprecateRule()
	// Experimental.
	ResetRetainRule()
	// Experimental.
	ResetTarget()
	// Experimental.
	ResetTargetRegion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference
type jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) CmkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) CmkArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmkArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) CopyTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) CopyTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) DeprecateRule() TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleDeprecateRulePropertyOutputReference {
	var returns TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleDeprecateRulePropertyOutputReference
	_jsii_.Get(
		j,
		"deprecateRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) DeprecateRuleInput() *TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleProperty {
	var returns *TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleProperty
	_jsii_.Get(
		j,
		"deprecateRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) Encrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) EncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) RetainRule() TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleRetainRulePropertyOutputReference {
	var returns TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleRetainRulePropertyOutputReference
	_jsii_.Get(
		j,
		"retainRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) RetainRuleInput() *TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleRetainRuleProperty {
	var returns *TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleRetainRuleProperty
	_jsii_.Get(
		j,
		"retainRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) Target() *string {
	var returns *string
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) TargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) TargetRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) TargetRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dlm.TfLifecyclePolicy.CrossRegionCopyRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference_Override(t TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dlm.TfLifecyclePolicy.CrossRegionCopyRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference)SetCmkArn(val *string) {
	if err := j.validateSetCmkArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cmkArn",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference)SetCopyTags(val interface{}) {
	if err := j.validateSetCopyTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTags",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference)SetEncrypted(val interface{}) {
	if err := j.validateSetEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference)SetTarget(val *string) {
	if err := j.validateSetTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"target",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference)SetTargetRegion(val *string) {
	if err := j.validateSetTargetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRegion",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) PutDeprecateRule(value *TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleDeprecateRuleProperty) {
	if err := t.validatePutDeprecateRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeprecateRule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) PutRetainRule(value *TfLifecyclePolicy_PolicyDetailsScheduleCrossRegionCopyRuleRetainRuleProperty) {
	if err := t.validatePutRetainRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRetainRule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) ResetCmkArn() {
	_jsii_.InvokeVoid(
		t,
		"resetCmkArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) ResetCopyTags() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) ResetDeprecateRule() {
	_jsii_.InvokeVoid(
		t,
		"resetDeprecateRule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) ResetRetainRule() {
	_jsii_.InvokeVoid(
		t,
		"resetRetainRule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) ResetTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) ResetTargetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfLifecyclePolicy_CrossRegionCopyRulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

