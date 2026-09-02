package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference interface {
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
	CookiesConfig() TfCachePolicy_CookiesConfigPropertyOutputReference
	// Experimental.
	CookiesConfigInput() *TfCachePolicy_CookiesConfigProperty
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
	HeadersConfig() TfCachePolicy_HeadersConfigPropertyOutputReference
	// Experimental.
	HeadersConfigInput() *TfCachePolicy_HeadersConfigProperty
	// Experimental.
	InternalValue() *TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty
	// Experimental.
	SetInternalValue(val *TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty)
	// Experimental.
	QueryStringsConfig() TfCachePolicy_QueryStringsConfigPropertyOutputReference
	// Experimental.
	QueryStringsConfigInput() *TfCachePolicy_QueryStringsConfigProperty
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
	PutCookiesConfig(value *TfCachePolicy_CookiesConfigProperty)
	// Experimental.
	PutHeadersConfig(value *TfCachePolicy_HeadersConfigProperty)
	// Experimental.
	PutQueryStringsConfig(value *TfCachePolicy_QueryStringsConfigProperty)
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

// The jsii proxy struct for TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference
type jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) CookiesConfig() TfCachePolicy_CookiesConfigPropertyOutputReference {
	var returns TfCachePolicy_CookiesConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"cookiesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) CookiesConfigInput() *TfCachePolicy_CookiesConfigProperty {
	var returns *TfCachePolicy_CookiesConfigProperty
	_jsii_.Get(
		j,
		"cookiesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) EnableAcceptEncodingBrotli() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingBrotli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) EnableAcceptEncodingBrotliInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingBrotliInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) EnableAcceptEncodingGzip() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingGzip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) EnableAcceptEncodingGzipInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceptEncodingGzipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) HeadersConfig() TfCachePolicy_HeadersConfigPropertyOutputReference {
	var returns TfCachePolicy_HeadersConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"headersConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) HeadersConfigInput() *TfCachePolicy_HeadersConfigProperty {
	var returns *TfCachePolicy_HeadersConfigProperty
	_jsii_.Get(
		j,
		"headersConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) InternalValue() *TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty {
	var returns *TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) QueryStringsConfig() TfCachePolicy_QueryStringsConfigPropertyOutputReference {
	var returns TfCachePolicy_QueryStringsConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"queryStringsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) QueryStringsConfigInput() *TfCachePolicy_QueryStringsConfigProperty {
	var returns *TfCachePolicy_QueryStringsConfigProperty
	_jsii_.Get(
		j,
		"queryStringsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfCachePolicy.ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference_Override(t TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfCachePolicy.ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetEnableAcceptEncodingBrotli(val interface{}) {
	if err := j.validateSetEnableAcceptEncodingBrotliParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAcceptEncodingBrotli",
		val,
	)
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetEnableAcceptEncodingGzip(val interface{}) {
	if err := j.validateSetEnableAcceptEncodingGzipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAcceptEncodingGzip",
		val,
	)
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetInternalValue(val *TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) PutCookiesConfig(value *TfCachePolicy_CookiesConfigProperty) {
	if err := t.validatePutCookiesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCookiesConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) PutHeadersConfig(value *TfCachePolicy_HeadersConfigProperty) {
	if err := t.validatePutHeadersConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHeadersConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) PutQueryStringsConfig(value *TfCachePolicy_QueryStringsConfigProperty) {
	if err := t.validatePutQueryStringsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putQueryStringsConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ResetEnableAcceptEncodingBrotli() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableAcceptEncodingBrotli",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ResetEnableAcceptEncodingGzip() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableAcceptEncodingGzip",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCachePolicy_ParametersInCacheKeyAndForwardedToOriginPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

