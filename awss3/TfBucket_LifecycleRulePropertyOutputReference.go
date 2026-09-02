package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfBucket_LifecycleRulePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AbortIncompleteMultipartUploadDays() *float64
	// Experimental.
	SetAbortIncompleteMultipartUploadDays(val *float64)
	// Experimental.
	AbortIncompleteMultipartUploadDaysInput() *float64
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
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	Expiration() TfBucket_ExpirationPropertyOutputReference
	// Experimental.
	ExpirationInput() *TfBucket_ExpirationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	NoncurrentVersionExpiration() TfBucket_NoncurrentVersionExpirationPropertyOutputReference
	// Experimental.
	NoncurrentVersionExpirationInput() *TfBucket_NoncurrentVersionExpirationProperty
	// Experimental.
	NoncurrentVersionTransition() TfBucket_NoncurrentVersionTransitionPropertyList
	// Experimental.
	NoncurrentVersionTransitionInput() interface{}
	// Experimental.
	Prefix() *string
	// Experimental.
	SetPrefix(val *string)
	// Experimental.
	PrefixInput() *string
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Transition() TfBucket_TransitionPropertyList
	// Experimental.
	TransitionInput() interface{}
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
	PutExpiration(value *TfBucket_ExpirationProperty)
	// Experimental.
	PutNoncurrentVersionExpiration(value *TfBucket_NoncurrentVersionExpirationProperty)
	// Experimental.
	PutNoncurrentVersionTransition(value interface{})
	// Experimental.
	PutTransition(value interface{})
	// Experimental.
	ResetAbortIncompleteMultipartUploadDays()
	// Experimental.
	ResetExpiration()
	// Experimental.
	ResetId()
	// Experimental.
	ResetNoncurrentVersionExpiration()
	// Experimental.
	ResetNoncurrentVersionTransition()
	// Experimental.
	ResetPrefix()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTransition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfBucket_LifecycleRulePropertyOutputReference
type jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) AbortIncompleteMultipartUploadDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUploadDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) AbortIncompleteMultipartUploadDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUploadDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) Expiration() TfBucket_ExpirationPropertyOutputReference {
	var returns TfBucket_ExpirationPropertyOutputReference
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) ExpirationInput() *TfBucket_ExpirationProperty {
	var returns *TfBucket_ExpirationProperty
	_jsii_.Get(
		j,
		"expirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) NoncurrentVersionExpiration() TfBucket_NoncurrentVersionExpirationPropertyOutputReference {
	var returns TfBucket_NoncurrentVersionExpirationPropertyOutputReference
	_jsii_.Get(
		j,
		"noncurrentVersionExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) NoncurrentVersionExpirationInput() *TfBucket_NoncurrentVersionExpirationProperty {
	var returns *TfBucket_NoncurrentVersionExpirationProperty
	_jsii_.Get(
		j,
		"noncurrentVersionExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) NoncurrentVersionTransition() TfBucket_NoncurrentVersionTransitionPropertyList {
	var returns TfBucket_NoncurrentVersionTransitionPropertyList
	_jsii_.Get(
		j,
		"noncurrentVersionTransition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) NoncurrentVersionTransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noncurrentVersionTransitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) Transition() TfBucket_TransitionPropertyList {
	var returns TfBucket_TransitionPropertyList
	_jsii_.Get(
		j,
		"transition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) TransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfBucket_LifecycleRulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfBucket_LifecycleRulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfBucket_LifecycleRulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.TfBucket.LifecycleRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfBucket_LifecycleRulePropertyOutputReference_Override(t TfBucket_LifecycleRulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.TfBucket.LifecycleRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference)SetAbortIncompleteMultipartUploadDays(val *float64) {
	if err := j.validateSetAbortIncompleteMultipartUploadDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"abortIncompleteMultipartUploadDays",
		val,
	)
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) PutExpiration(value *TfBucket_ExpirationProperty) {
	if err := t.validatePutExpirationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExpiration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) PutNoncurrentVersionExpiration(value *TfBucket_NoncurrentVersionExpirationProperty) {
	if err := t.validatePutNoncurrentVersionExpirationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNoncurrentVersionExpiration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) PutNoncurrentVersionTransition(value interface{}) {
	if err := t.validatePutNoncurrentVersionTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNoncurrentVersionTransition",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) PutTransition(value interface{}) {
	if err := t.validatePutTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTransition",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) ResetAbortIncompleteMultipartUploadDays() {
	_jsii_.InvokeVoid(
		t,
		"resetAbortIncompleteMultipartUploadDays",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) ResetExpiration() {
	_jsii_.InvokeVoid(
		t,
		"resetExpiration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) ResetNoncurrentVersionExpiration() {
	_jsii_.InvokeVoid(
		t,
		"resetNoncurrentVersionExpiration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) ResetNoncurrentVersionTransition() {
	_jsii_.InvokeVoid(
		t,
		"resetNoncurrentVersionTransition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) ResetTransition() {
	_jsii_.InvokeVoid(
		t,
		"resetTransition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfBucket_LifecycleRulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

