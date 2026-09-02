package awslightsail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslightsail/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslightsail/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDistribution_CacheBehaviorSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllowedHttpMethods() *string
	// Experimental.
	SetAllowedHttpMethods(val *string)
	// Experimental.
	AllowedHttpMethodsInput() *string
	// Experimental.
	CachedHttpMethods() *string
	// Experimental.
	SetCachedHttpMethods(val *string)
	// Experimental.
	CachedHttpMethodsInput() *string
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
	DefaultTtl() *float64
	// Experimental.
	SetDefaultTtl(val *float64)
	// Experimental.
	DefaultTtlInput() *float64
	// Experimental.
	ForwardedCookies() TfDistribution_ForwardedCookiesPropertyOutputReference
	// Experimental.
	ForwardedCookiesInput() *TfDistribution_ForwardedCookiesProperty
	// Experimental.
	ForwardedHeaders() TfDistribution_ForwardedHeadersPropertyOutputReference
	// Experimental.
	ForwardedHeadersInput() *TfDistribution_ForwardedHeadersProperty
	// Experimental.
	ForwardedQueryStrings() TfDistribution_ForwardedQueryStringsPropertyOutputReference
	// Experimental.
	ForwardedQueryStringsInput() *TfDistribution_ForwardedQueryStringsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDistribution_CacheBehaviorSettingsProperty
	// Experimental.
	SetInternalValue(val *TfDistribution_CacheBehaviorSettingsProperty)
	// Experimental.
	MaximumTtl() *float64
	// Experimental.
	SetMaximumTtl(val *float64)
	// Experimental.
	MaximumTtlInput() *float64
	// Experimental.
	MinimumTtl() *float64
	// Experimental.
	SetMinimumTtl(val *float64)
	// Experimental.
	MinimumTtlInput() *float64
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
	PutForwardedCookies(value *TfDistribution_ForwardedCookiesProperty)
	// Experimental.
	PutForwardedHeaders(value *TfDistribution_ForwardedHeadersProperty)
	// Experimental.
	PutForwardedQueryStrings(value *TfDistribution_ForwardedQueryStringsProperty)
	// Experimental.
	ResetAllowedHttpMethods()
	// Experimental.
	ResetCachedHttpMethods()
	// Experimental.
	ResetDefaultTtl()
	// Experimental.
	ResetForwardedCookies()
	// Experimental.
	ResetForwardedHeaders()
	// Experimental.
	ResetForwardedQueryStrings()
	// Experimental.
	ResetMaximumTtl()
	// Experimental.
	ResetMinimumTtl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDistribution_CacheBehaviorSettingsPropertyOutputReference
type jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) AllowedHttpMethods() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedHttpMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) AllowedHttpMethodsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedHttpMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) CachedHttpMethods() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachedHttpMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) CachedHttpMethodsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachedHttpMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) DefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ForwardedCookies() TfDistribution_ForwardedCookiesPropertyOutputReference {
	var returns TfDistribution_ForwardedCookiesPropertyOutputReference
	_jsii_.Get(
		j,
		"forwardedCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ForwardedCookiesInput() *TfDistribution_ForwardedCookiesProperty {
	var returns *TfDistribution_ForwardedCookiesProperty
	_jsii_.Get(
		j,
		"forwardedCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ForwardedHeaders() TfDistribution_ForwardedHeadersPropertyOutputReference {
	var returns TfDistribution_ForwardedHeadersPropertyOutputReference
	_jsii_.Get(
		j,
		"forwardedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ForwardedHeadersInput() *TfDistribution_ForwardedHeadersProperty {
	var returns *TfDistribution_ForwardedHeadersProperty
	_jsii_.Get(
		j,
		"forwardedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ForwardedQueryStrings() TfDistribution_ForwardedQueryStringsPropertyOutputReference {
	var returns TfDistribution_ForwardedQueryStringsPropertyOutputReference
	_jsii_.Get(
		j,
		"forwardedQueryStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ForwardedQueryStringsInput() *TfDistribution_ForwardedQueryStringsProperty {
	var returns *TfDistribution_ForwardedQueryStringsProperty
	_jsii_.Get(
		j,
		"forwardedQueryStringsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) InternalValue() *TfDistribution_CacheBehaviorSettingsProperty {
	var returns *TfDistribution_CacheBehaviorSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) MaximumTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) MaximumTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) MinimumTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) MinimumTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDistribution_CacheBehaviorSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDistribution_CacheBehaviorSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDistribution_CacheBehaviorSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lightsail.TfDistribution.CacheBehaviorSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDistribution_CacheBehaviorSettingsPropertyOutputReference_Override(t TfDistribution_CacheBehaviorSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lightsail.TfDistribution.CacheBehaviorSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference)SetAllowedHttpMethods(val *string) {
	if err := j.validateSetAllowedHttpMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedHttpMethods",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference)SetCachedHttpMethods(val *string) {
	if err := j.validateSetCachedHttpMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cachedHttpMethods",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference)SetDefaultTtl(val *float64) {
	if err := j.validateSetDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference)SetInternalValue(val *TfDistribution_CacheBehaviorSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference)SetMaximumTtl(val *float64) {
	if err := j.validateSetMaximumTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumTtl",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference)SetMinimumTtl(val *float64) {
	if err := j.validateSetMinimumTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumTtl",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) PutForwardedCookies(value *TfDistribution_ForwardedCookiesProperty) {
	if err := t.validatePutForwardedCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putForwardedCookies",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) PutForwardedHeaders(value *TfDistribution_ForwardedHeadersProperty) {
	if err := t.validatePutForwardedHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putForwardedHeaders",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) PutForwardedQueryStrings(value *TfDistribution_ForwardedQueryStringsProperty) {
	if err := t.validatePutForwardedQueryStringsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putForwardedQueryStrings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ResetAllowedHttpMethods() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowedHttpMethods",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ResetCachedHttpMethods() {
	_jsii_.InvokeVoid(
		t,
		"resetCachedHttpMethods",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ResetForwardedCookies() {
	_jsii_.InvokeVoid(
		t,
		"resetForwardedCookies",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ResetForwardedHeaders() {
	_jsii_.InvokeVoid(
		t,
		"resetForwardedHeaders",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ResetForwardedQueryStrings() {
	_jsii_.InvokeVoid(
		t,
		"resetForwardedQueryStrings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ResetMaximumTtl() {
	_jsii_.InvokeVoid(
		t,
		"resetMaximumTtl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ResetMinimumTtl() {
	_jsii_.InvokeVoid(
		t,
		"resetMinimumTtl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDistribution_CacheBehaviorSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

