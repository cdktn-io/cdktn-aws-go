package cloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference interface {
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
	Cookies() AwsDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference
	// Experimental.
	CookiesInput() *AwsDistribution_DefaultCacheBehaviorForwardedValuesCookiesProperty
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
	InternalValue() *AwsDistribution_DefaultCacheBehaviorForwardedValuesProperty
	// Experimental.
	SetInternalValue(val *AwsDistribution_DefaultCacheBehaviorForwardedValuesProperty)
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
	PutCookies(value *AwsDistribution_DefaultCacheBehaviorForwardedValuesCookiesProperty)
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

// The jsii proxy struct for AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference
type jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) Cookies() AwsDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference {
	var returns AwsDistribution_DefaultCacheBehaviorForwardedValuesCookiesPropertyOutputReference
	_jsii_.Get(
		j,
		"cookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) CookiesInput() *AwsDistribution_DefaultCacheBehaviorForwardedValuesCookiesProperty {
	var returns *AwsDistribution_DefaultCacheBehaviorForwardedValuesCookiesProperty
	_jsii_.Get(
		j,
		"cookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) Headers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) HeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) InternalValue() *AwsDistribution_DefaultCacheBehaviorForwardedValuesProperty {
	var returns *AwsDistribution_DefaultCacheBehaviorForwardedValuesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) QueryString() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) QueryStringCacheKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringCacheKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) QueryStringCacheKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringCacheKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) QueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsDistribution.DefaultCacheBehaviorForwardedValuesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference_Override(a AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsDistribution.DefaultCacheBehaviorForwardedValuesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference)SetHeaders(val *[]*string) {
	if err := j.validateSetHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference)SetInternalValue(val *AwsDistribution_DefaultCacheBehaviorForwardedValuesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference)SetQueryString(val interface{}) {
	if err := j.validateSetQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryString",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference)SetQueryStringCacheKeys(val *[]*string) {
	if err := j.validateSetQueryStringCacheKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringCacheKeys",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) PutCookies(value *AwsDistribution_DefaultCacheBehaviorForwardedValuesCookiesProperty) {
	if err := a.validatePutCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCookies",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) ResetQueryStringCacheKeys() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryStringCacheKeys",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDistribution_DefaultCacheBehaviorForwardedValuesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

