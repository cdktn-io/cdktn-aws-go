package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference interface {
	cdktn.ComplexObject
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
	Cookies() TfDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference
	// Experimental.
	CookiesInput() *TfDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	Headers() *[]*string
	// Experimental.
	SetHeaders(val *[]*string)
	// Experimental.
	HeadersInput() *[]*string
	// Experimental.
	InternalValue() *TfDistribution_OrderedCacheBehaviorForwardedValuesProperty
	// Experimental.
	SetInternalValue(val *TfDistribution_OrderedCacheBehaviorForwardedValuesProperty)
	// Experimental.
	QueryString() interface{}
	// Experimental.
	SetQueryString(val interface{})
	// Experimental.
	QueryStringCacheKeys() *[]*string
	// Experimental.
	SetQueryStringCacheKeys(val *[]*string)
	// Experimental.
	QueryStringCacheKeysInput() *[]*string
	// Experimental.
	QueryStringInput() interface{}
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
	PutCookies(value *TfDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty)
	// Experimental.
	ResetHeaders()
	// Experimental.
	ResetQueryStringCacheKeys()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference
type jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) Cookies() TfDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference {
	var returns TfDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) CookiesInput() *TfDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty {
	var returns *TfDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) Headers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) HeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) InternalValue() *TfDistribution_OrderedCacheBehaviorForwardedValuesProperty {
	var returns *TfDistribution_OrderedCacheBehaviorForwardedValuesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) QueryString() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) QueryStringCacheKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringCacheKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) QueryStringCacheKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringCacheKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfDistribution.OrderedCacheBehaviorForwardedValuesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference_Override(t TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfDistribution.OrderedCacheBehaviorForwardedValuesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference)SetHeaders(val *[]*string) {
	if err := j.validateSetHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference)SetInternalValue(val *TfDistribution_OrderedCacheBehaviorForwardedValuesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference)SetQueryString(val interface{}) {
	if err := j.validateSetQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference)SetQueryStringCacheKeys(val *[]*string) {
	if err := j.validateSetQueryStringCacheKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringCacheKeys",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) PutCookies(value *TfDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty) {
	if err := t.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCookies",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		t,
		"resetHeaders",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) ResetQueryStringCacheKeys() {
	_jsii_.InvokeVoid(
		t,
		"resetQueryStringCacheKeys",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

