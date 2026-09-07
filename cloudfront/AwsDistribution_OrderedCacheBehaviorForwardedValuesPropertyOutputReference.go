package cloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference interface {
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
	Cookies() AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference
	// Experimental.
	CookiesInput() *AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty
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
	InternalValue() *AwsDistribution_OrderedCacheBehaviorForwardedValuesProperty
	// Experimental.
	SetInternalValue(val *AwsDistribution_OrderedCacheBehaviorForwardedValuesProperty)
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
	PutCookies(value *AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty)
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

// The jsii proxy struct for AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference
type jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) Cookies() AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference {
	var returns AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesPropertyOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) CookiesInput() *AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty {
	var returns *AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) Headers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) HeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) InternalValue() *AwsDistribution_OrderedCacheBehaviorForwardedValuesProperty {
	var returns *AwsDistribution_OrderedCacheBehaviorForwardedValuesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) QueryString() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) QueryStringCacheKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringCacheKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) QueryStringCacheKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringCacheKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsDistribution.OrderedCacheBehaviorForwardedValuesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference_Override(a AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsDistribution.OrderedCacheBehaviorForwardedValuesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference)SetHeaders(val *[]*string) {
	if err := j.validateSetHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference)SetInternalValue(val *AwsDistribution_OrderedCacheBehaviorForwardedValuesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference)SetQueryString(val interface{}) {
	if err := j.validateSetQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference)SetQueryStringCacheKeys(val *[]*string) {
	if err := j.validateSetQueryStringCacheKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringCacheKeys",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) PutCookies(value *AwsDistribution_OrderedCacheBehaviorForwardedValuesCookiesProperty) {
	if err := a.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCookies",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) ResetQueryStringCacheKeys() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryStringCacheKeys",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDistribution_OrderedCacheBehaviorForwardedValuesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

