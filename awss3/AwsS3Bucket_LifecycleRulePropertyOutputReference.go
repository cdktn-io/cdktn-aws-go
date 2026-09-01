package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsS3Bucket_LifecycleRulePropertyOutputReference interface {
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
	Expiration() AwsS3Bucket_ExpirationPropertyOutputReference
	// Experimental.
	ExpirationInput() *AwsS3Bucket_ExpirationProperty
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
	NoncurrentVersionExpiration() AwsS3Bucket_NoncurrentVersionExpirationPropertyOutputReference
	// Experimental.
	NoncurrentVersionExpirationInput() *AwsS3Bucket_NoncurrentVersionExpirationProperty
	// Experimental.
	NoncurrentVersionTransition() AwsS3Bucket_NoncurrentVersionTransitionPropertyList
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
	Transition() AwsS3Bucket_TransitionPropertyList
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
	PutExpiration(value *AwsS3Bucket_ExpirationProperty)
	// Experimental.
	PutNoncurrentVersionExpiration(value *AwsS3Bucket_NoncurrentVersionExpirationProperty)
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

// The jsii proxy struct for AwsS3Bucket_LifecycleRulePropertyOutputReference
type jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) AbortIncompleteMultipartUploadDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUploadDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) AbortIncompleteMultipartUploadDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUploadDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) Expiration() AwsS3Bucket_ExpirationPropertyOutputReference {
	var returns AwsS3Bucket_ExpirationPropertyOutputReference
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) ExpirationInput() *AwsS3Bucket_ExpirationProperty {
	var returns *AwsS3Bucket_ExpirationProperty
	_jsii_.Get(
		j,
		"expirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) NoncurrentVersionExpiration() AwsS3Bucket_NoncurrentVersionExpirationPropertyOutputReference {
	var returns AwsS3Bucket_NoncurrentVersionExpirationPropertyOutputReference
	_jsii_.Get(
		j,
		"noncurrentVersionExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) NoncurrentVersionExpirationInput() *AwsS3Bucket_NoncurrentVersionExpirationProperty {
	var returns *AwsS3Bucket_NoncurrentVersionExpirationProperty
	_jsii_.Get(
		j,
		"noncurrentVersionExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) NoncurrentVersionTransition() AwsS3Bucket_NoncurrentVersionTransitionPropertyList {
	var returns AwsS3Bucket_NoncurrentVersionTransitionPropertyList
	_jsii_.Get(
		j,
		"noncurrentVersionTransition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) NoncurrentVersionTransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noncurrentVersionTransitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) Transition() AwsS3Bucket_TransitionPropertyList {
	var returns AwsS3Bucket_TransitionPropertyList
	_jsii_.Get(
		j,
		"transition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) TransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsS3Bucket_LifecycleRulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsS3Bucket_LifecycleRulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsS3Bucket_LifecycleRulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.AwsS3Bucket.LifecycleRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsS3Bucket_LifecycleRulePropertyOutputReference_Override(a AwsS3Bucket_LifecycleRulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.AwsS3Bucket.LifecycleRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference)SetAbortIncompleteMultipartUploadDays(val *float64) {
	if err := j.validateSetAbortIncompleteMultipartUploadDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"abortIncompleteMultipartUploadDays",
		val,
	)
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) PutExpiration(value *AwsS3Bucket_ExpirationProperty) {
	if err := a.validatePutExpirationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExpiration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) PutNoncurrentVersionExpiration(value *AwsS3Bucket_NoncurrentVersionExpirationProperty) {
	if err := a.validatePutNoncurrentVersionExpirationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNoncurrentVersionExpiration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) PutNoncurrentVersionTransition(value interface{}) {
	if err := a.validatePutNoncurrentVersionTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNoncurrentVersionTransition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) PutTransition(value interface{}) {
	if err := a.validatePutTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTransition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) ResetAbortIncompleteMultipartUploadDays() {
	_jsii_.InvokeVoid(
		a,
		"resetAbortIncompleteMultipartUploadDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) ResetExpiration() {
	_jsii_.InvokeVoid(
		a,
		"resetExpiration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) ResetNoncurrentVersionExpiration() {
	_jsii_.InvokeVoid(
		a,
		"resetNoncurrentVersionExpiration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) ResetNoncurrentVersionTransition() {
	_jsii_.InvokeVoid(
		a,
		"resetNoncurrentVersionTransition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) ResetTransition() {
	_jsii_.InvokeVoid(
		a,
		"resetTransition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsS3Bucket_LifecycleRulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

