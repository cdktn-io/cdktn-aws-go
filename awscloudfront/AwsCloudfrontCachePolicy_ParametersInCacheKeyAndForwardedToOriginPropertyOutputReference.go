package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference interface {
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
	CookiesConfig() AwsCloudfrontCachePolicy_CookiesConfigPropertyOutputReference
	// Experimental.
	CookiesConfigInput() *AwsCloudfrontCachePolicy_CookiesConfigProperty
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
	HeadersConfig() AwsCloudfrontCachePolicy_HeadersConfigPropertyOutputReference
	// Experimental.
	HeadersConfigInput() *AwsCloudfrontCachePolicy_HeadersConfigProperty
	// Experimental.
	InternalValue() *AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty
	// Experimental.
	SetInternalValue(val *AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty)
	// Experimental.
	QueryStringsConfig() AwsCloudfrontCachePolicy_QueryStringsConfigPropertyOutputReference
	// Experimental.
	QueryStringsConfigInput() *AwsCloudfrontCachePolicy_QueryStringsConfigProperty
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
	PutCookiesConfig(value *AwsCloudfrontCachePolicy_CookiesConfigProperty)
	// Experimental.
	PutHeadersConfig(value *AwsCloudfrontCachePolicy_HeadersConfigProperty)
	// Experimental.
	PutQueryStringsConfig(value *AwsCloudfrontCachePolicy_QueryStringsConfigProperty)
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

// The jsii proxy struct for AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference
type jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) CookiesConfig() AwsCloudfrontCachePolicy_CookiesConfigPropertyOutputReference {
	var returns AwsCloudfrontCachePolicy_CookiesConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"cookiesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) CookiesConfigInput() *AwsCloudfrontCachePolicy_CookiesConfigProperty {
	var returns *AwsCloudfrontCachePolicy_CookiesConfigProperty
	_jsii_.Get(
		j,
		"cookiesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) EnableAcceptEncodingBrotli() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingBrotli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) EnableAcceptEncodingBrotliInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingBrotliInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) EnableAcceptEncodingGzip() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingGzip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) EnableAcceptEncodingGzipInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingGzipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) HeadersConfig() AwsCloudfrontCachePolicy_HeadersConfigPropertyOutputReference {
	var returns AwsCloudfrontCachePolicy_HeadersConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"headersConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) HeadersConfigInput() *AwsCloudfrontCachePolicy_HeadersConfigProperty {
	var returns *AwsCloudfrontCachePolicy_HeadersConfigProperty
	_jsii_.Get(
		j,
		"headersConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) InternalValue() *AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty {
	var returns *AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) QueryStringsConfig() AwsCloudfrontCachePolicy_QueryStringsConfigPropertyOutputReference {
	var returns AwsCloudfrontCachePolicy_QueryStringsConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"queryStringsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) QueryStringsConfigInput() *AwsCloudfrontCachePolicy_QueryStringsConfigProperty {
	var returns *AwsCloudfrontCachePolicy_QueryStringsConfigProperty
	_jsii_.Get(
		j,
		"queryStringsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsCloudfrontCachePolicy.ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference_Override(a AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.AwsCloudfrontCachePolicy.ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetEnableAcceptEncodingBrotli(val interface{}) {
	if err := j.validateSetEnableAcceptEncodingBrotliParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAcceptEncodingBrotli",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetEnableAcceptEncodingGzip(val interface{}) {
	if err := j.validateSetEnableAcceptEncodingGzipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAcceptEncodingGzip",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetInternalValue(val *AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) PutCookiesConfig(value *AwsCloudfrontCachePolicy_CookiesConfigProperty) {
	if err := a.validatePutCookiesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCookiesConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) PutHeadersConfig(value *AwsCloudfrontCachePolicy_HeadersConfigProperty) {
	if err := a.validatePutHeadersConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHeadersConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) PutQueryStringsConfig(value *AwsCloudfrontCachePolicy_QueryStringsConfigProperty) {
	if err := a.validatePutQueryStringsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueryStringsConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ResetEnableAcceptEncodingBrotli() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableAcceptEncodingBrotli",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ResetEnableAcceptEncodingGzip() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableAcceptEncodingGzip",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCloudfrontCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

