package cloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/cloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/cloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference interface {
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
	CookiesConfig() AwsCachePolicy_CookiesConfigPropertyOutputReference
	// Experimental.
	CookiesConfigInput() *AwsCachePolicy_CookiesConfigProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EnableAcceptEncodingBrotli() interface{}
	// Experimental.
	SetEnableAcceptEncodingBrotli(val interface{})
	// Experimental.
	EnableAcceptEncodingBrotliInput() interface{}
	// Experimental.
	EnableAcceptEncodingGzip() interface{}
	// Experimental.
	SetEnableAcceptEncodingGzip(val interface{})
	// Experimental.
	EnableAcceptEncodingGzipInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	HeadersConfig() AwsCachePolicy_HeadersConfigPropertyOutputReference
	// Experimental.
	HeadersConfigInput() *AwsCachePolicy_HeadersConfigProperty
	// Experimental.
	InternalValue() *AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty
	// Experimental.
	SetInternalValue(val *AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty)
	// Experimental.
	QueryStringsConfig() AwsCachePolicy_QueryStringsConfigPropertyOutputReference
	// Experimental.
	QueryStringsConfigInput() *AwsCachePolicy_QueryStringsConfigProperty
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
	PutCookiesConfig(value *AwsCachePolicy_CookiesConfigProperty)
	// Experimental.
	PutHeadersConfig(value *AwsCachePolicy_HeadersConfigProperty)
	// Experimental.
	PutQueryStringsConfig(value *AwsCachePolicy_QueryStringsConfigProperty)
	// Experimental.
	ResetEnableAcceptEncodingBrotli()
	// Experimental.
	ResetEnableAcceptEncodingGzip()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference
type jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) CookiesConfig() AwsCachePolicy_CookiesConfigPropertyOutputReference {
	var returns AwsCachePolicy_CookiesConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"cookiesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) CookiesConfigInput() *AwsCachePolicy_CookiesConfigProperty {
	var returns *AwsCachePolicy_CookiesConfigProperty
	_jsii_.Get(
		j,
		"cookiesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) EnableAcceptEncodingBrotli() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingBrotli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) EnableAcceptEncodingBrotliInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingBrotliInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) EnableAcceptEncodingGzip() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingGzip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) EnableAcceptEncodingGzipInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingGzipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) HeadersConfig() AwsCachePolicy_HeadersConfigPropertyOutputReference {
	var returns AwsCachePolicy_HeadersConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"headersConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) HeadersConfigInput() *AwsCachePolicy_HeadersConfigProperty {
	var returns *AwsCachePolicy_HeadersConfigProperty
	_jsii_.Get(
		j,
		"headersConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) InternalValue() *AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty {
	var returns *AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) QueryStringsConfig() AwsCachePolicy_QueryStringsConfigPropertyOutputReference {
	var returns AwsCachePolicy_QueryStringsConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"queryStringsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) QueryStringsConfigInput() *AwsCachePolicy_QueryStringsConfigProperty {
	var returns *AwsCachePolicy_QueryStringsConfigProperty
	_jsii_.Get(
		j,
		"queryStringsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsCachePolicy.ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference_Override(a AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsCachePolicy.ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetEnableAcceptEncodingBrotli(val interface{}) {
	if err := j.validateSetEnableAcceptEncodingBrotliParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAcceptEncodingBrotli",
		val,
	)
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetEnableAcceptEncodingGzip(val interface{}) {
	if err := j.validateSetEnableAcceptEncodingGzipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAcceptEncodingGzip",
		val,
	)
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetInternalValue(val *AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) PutCookiesConfig(value *AwsCachePolicy_CookiesConfigProperty) {
	if err := a.validatePutCookiesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCookiesConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) PutHeadersConfig(value *AwsCachePolicy_HeadersConfigProperty) {
	if err := a.validatePutHeadersConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHeadersConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) PutQueryStringsConfig(value *AwsCachePolicy_QueryStringsConfigProperty) {
	if err := a.validatePutQueryStringsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueryStringsConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ResetEnableAcceptEncodingBrotli() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableAcceptEncodingBrotli",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ResetEnableAcceptEncodingGzip() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableAcceptEncodingGzip",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

