package s3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/s3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/s3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBucketLifecycleConfiguration_RulePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AbortIncompleteMultipartUpload() AwsBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyList
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
	Expiration() AwsBucketLifecycleConfiguration_ExpirationPropertyList
	// Experimental.
	ExpirationInput() interface{}
	// Experimental.
	Filter() AwsBucketLifecycleConfiguration_FilterPropertyList
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
	NoncurrentVersionExpiration() AwsBucketLifecycleConfiguration_NoncurrentVersionExpirationPropertyList
	// Experimental.
	NoncurrentVersionExpirationInput() interface{}
	// Experimental.
	NoncurrentVersionTransition() AwsBucketLifecycleConfiguration_NoncurrentVersionTransitionPropertyList
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
	Transition() AwsBucketLifecycleConfiguration_TransitionPropertyList
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

// The jsii proxy struct for AwsBucketLifecycleConfiguration_RulePropertyOutputReference
type jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) AbortIncompleteMultipartUpload() AwsBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyList {
	var returns AwsBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyList
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUpload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) AbortIncompleteMultipartUploadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortIncompleteMultipartUploadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) Expiration() AwsBucketLifecycleConfiguration_ExpirationPropertyList {
	var returns AwsBucketLifecycleConfiguration_ExpirationPropertyList
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) ExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) Filter() AwsBucketLifecycleConfiguration_FilterPropertyList {
	var returns AwsBucketLifecycleConfiguration_FilterPropertyList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) NoncurrentVersionExpiration() AwsBucketLifecycleConfiguration_NoncurrentVersionExpirationPropertyList {
	var returns AwsBucketLifecycleConfiguration_NoncurrentVersionExpirationPropertyList
	_jsii_.Get(
		j,
		"noncurrentVersionExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) NoncurrentVersionExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noncurrentVersionExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) NoncurrentVersionTransition() AwsBucketLifecycleConfiguration_NoncurrentVersionTransitionPropertyList {
	var returns AwsBucketLifecycleConfiguration_NoncurrentVersionTransitionPropertyList
	_jsii_.Get(
		j,
		"noncurrentVersionTransition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) NoncurrentVersionTransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noncurrentVersionTransitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) Transition() AwsBucketLifecycleConfiguration_TransitionPropertyList {
	var returns AwsBucketLifecycleConfiguration_TransitionPropertyList
	_jsii_.Get(
		j,
		"transition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) TransitionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBucketLifecycleConfiguration_RulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBucketLifecycleConfiguration_RulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBucketLifecycleConfiguration_RulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.AwsBucketLifecycleConfiguration.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBucketLifecycleConfiguration_RulePropertyOutputReference_Override(a AwsBucketLifecycleConfiguration_RulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.AwsBucketLifecycleConfiguration.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) PutAbortIncompleteMultipartUpload(value interface{}) {
	if err := a.validatePutAbortIncompleteMultipartUploadParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAbortIncompleteMultipartUpload",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) PutExpiration(value interface{}) {
	if err := a.validatePutExpirationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExpiration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) PutFilter(value interface{}) {
	if err := a.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) PutNoncurrentVersionExpiration(value interface{}) {
	if err := a.validatePutNoncurrentVersionExpirationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNoncurrentVersionExpiration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) PutNoncurrentVersionTransition(value interface{}) {
	if err := a.validatePutNoncurrentVersionTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNoncurrentVersionTransition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) PutTransition(value interface{}) {
	if err := a.validatePutTransitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTransition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) ResetAbortIncompleteMultipartUpload() {
	_jsii_.InvokeVoid(
		a,
		"resetAbortIncompleteMultipartUpload",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) ResetExpiration() {
	_jsii_.InvokeVoid(
		a,
		"resetExpiration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) ResetNoncurrentVersionExpiration() {
	_jsii_.InvokeVoid(
		a,
		"resetNoncurrentVersionExpiration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) ResetNoncurrentVersionTransition() {
	_jsii_.InvokeVoid(
		a,
		"resetNoncurrentVersionTransition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) ResetTransition() {
	_jsii_.InvokeVoid(
		a,
		"resetTransition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_RulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

