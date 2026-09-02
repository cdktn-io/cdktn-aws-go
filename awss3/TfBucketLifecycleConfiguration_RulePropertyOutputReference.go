package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfBucketLifecycleConfiguration_RulePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AbortIncompleteMultipartUpload() TfBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyList
	// Experimental.
	AbortIncompleteMultipartUploadInput() interface{}
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
	Expiration() TfBucketLifecycleConfiguration_ExpirationPropertyList
	// Experimental.
	ExpirationInput() interface{}
	// Experimental.
	Filter() TfBucketLifecycleConfiguration_FilterPropertyList
	// Experimental.
	FilterInput() interface{}
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
	NoncurrentVersionExpiration() TfBucketLifecycleConfiguration_NoncurrentVersionExpirationPropertyList
	// Experimental.
	NoncurrentVersionExpirationInput() interface{}
	// Experimental.
	NoncurrentVersionTransition() TfBucketLifecycleConfiguration_NoncurrentVersionTransitionPropertyList
	// Experimental.
	NoncurrentVersionTransitionInput() interface{}
	// Experimental.
	Prefix() *string
	// Experimental.
	SetPrefix(val *string)
	// Experimental.
	PrefixInput() *string
	// Experimental.
	Status() *string
	// Experimental.
	SetStatus(val *string)
	// Experimental.
	StatusInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Transition() TfBucketLifecycleConfiguration_TransitionPropertyList
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
	PutAbortIncompleteMultipartUpload(value interface{})
	// Experimental.
	PutExpiration(value interface{})
	// Experimental.
	PutFilter(value interface{})
	// Experimental.
	PutNoncurrentVersionExpiration(value interface{})
	// Experimental.
	PutNoncurrentVersionTransition(value interface{})
	// Experimental.
	PutTransition(value interface{})
	// Experimental.
	ResetAbortIncompleteMultipartUpload()
	// Experimental.
	ResetExpiration()
	// Experimental.
	ResetFilter()
	// Experimental.
	ResetNoncurrentVersionExpiration()
	// Experimental.
	ResetNoncurrentVersionTransition()
	// Experimental.
	ResetPrefix()
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

// The jsii proxy struct for TfBucketLifecycleConfiguration_RulePropertyOutputReference
type jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) AbortIncompleteMultipartUpload() TfBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyList {
	var returns TfBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyList
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUpload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) AbortIncompleteMultipartUploadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUploadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) Expiration() TfBucketLifecycleConfiguration_ExpirationPropertyList {
	var returns TfBucketLifecycleConfiguration_ExpirationPropertyList
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) ExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) Filter() TfBucketLifecycleConfiguration_FilterPropertyList {
	var returns TfBucketLifecycleConfiguration_FilterPropertyList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) NoncurrentVersionExpiration() TfBucketLifecycleConfiguration_NoncurrentVersionExpirationPropertyList {
	var returns TfBucketLifecycleConfiguration_NoncurrentVersionExpirationPropertyList
	_jsii_.Get(
		j,
		"noncurrentVersionExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) NoncurrentVersionExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noncurrentVersionExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) NoncurrentVersionTransition() TfBucketLifecycleConfiguration_NoncurrentVersionTransitionPropertyList {
	var returns TfBucketLifecycleConfiguration_NoncurrentVersionTransitionPropertyList
	_jsii_.Get(
		j,
		"noncurrentVersionTransition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) NoncurrentVersionTransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noncurrentVersionTransitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) Transition() TfBucketLifecycleConfiguration_TransitionPropertyList {
	var returns TfBucketLifecycleConfiguration_TransitionPropertyList
	_jsii_.Get(
		j,
		"transition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) TransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfBucketLifecycleConfiguration_RulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfBucketLifecycleConfiguration_RulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfBucketLifecycleConfiguration_RulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.TfBucketLifecycleConfiguration.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfBucketLifecycleConfiguration_RulePropertyOutputReference_Override(t TfBucketLifecycleConfiguration_RulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.TfBucketLifecycleConfiguration.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) PutAbortIncompleteMultipartUpload(value interface{}) {
	if err := t.validatePutAbortIncompleteMultipartUploadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAbortIncompleteMultipartUpload",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) PutExpiration(value interface{}) {
	if err := t.validatePutExpirationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExpiration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) PutFilter(value interface{}) {
	if err := t.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFilter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) PutNoncurrentVersionExpiration(value interface{}) {
	if err := t.validatePutNoncurrentVersionExpirationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNoncurrentVersionExpiration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) PutNoncurrentVersionTransition(value interface{}) {
	if err := t.validatePutNoncurrentVersionTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNoncurrentVersionTransition",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) PutTransition(value interface{}) {
	if err := t.validatePutTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTransition",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) ResetAbortIncompleteMultipartUpload() {
	_jsii_.InvokeVoid(
		t,
		"resetAbortIncompleteMultipartUpload",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) ResetExpiration() {
	_jsii_.InvokeVoid(
		t,
		"resetExpiration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		t,
		"resetFilter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) ResetNoncurrentVersionExpiration() {
	_jsii_.InvokeVoid(
		t,
		"resetNoncurrentVersionExpiration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) ResetNoncurrentVersionTransition() {
	_jsii_.InvokeVoid(
		t,
		"resetNoncurrentVersionTransition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) ResetTransition() {
	_jsii_.InvokeVoid(
		t,
		"resetTransition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfBucketLifecycleConfiguration_RulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

