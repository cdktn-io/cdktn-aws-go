package awslightsail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslightsail/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslightsail/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference interface {
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
	ForwardedCookies() AwsLightsailDistribution_ForwardedCookiesPropertyOutputReference
	// Experimental.
	ForwardedCookiesInput() *AwsLightsailDistribution_ForwardedCookiesProperty
	// Experimental.
	ForwardedHeaders() AwsLightsailDistribution_ForwardedHeadersPropertyOutputReference
	// Experimental.
	ForwardedHeadersInput() *AwsLightsailDistribution_ForwardedHeadersProperty
	// Experimental.
	ForwardedQueryStrings() AwsLightsailDistribution_ForwardedQueryStringsPropertyOutputReference
	// Experimental.
	ForwardedQueryStringsInput() *AwsLightsailDistribution_ForwardedQueryStringsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsLightsailDistribution_CacheBehaviorSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsLightsailDistribution_CacheBehaviorSettingsProperty)
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
	PutForwardedCookies(value *AwsLightsailDistribution_ForwardedCookiesProperty)
	// Experimental.
	PutForwardedHeaders(value *AwsLightsailDistribution_ForwardedHeadersProperty)
	// Experimental.
	PutForwardedQueryStrings(value *AwsLightsailDistribution_ForwardedQueryStringsProperty)
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

// The jsii proxy struct for AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference
type jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) AllowedHttpMethods() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedHttpMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) AllowedHttpMethodsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedHttpMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) CachedHttpMethods() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachedHttpMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) CachedHttpMethodsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachedHttpMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) DefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ForwardedCookies() AwsLightsailDistribution_ForwardedCookiesPropertyOutputReference {
	var returns AwsLightsailDistribution_ForwardedCookiesPropertyOutputReference
	_jsii_.Get(
		j,
		"forwardedCookies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ForwardedCookiesInput() *AwsLightsailDistribution_ForwardedCookiesProperty {
	var returns *AwsLightsailDistribution_ForwardedCookiesProperty
	_jsii_.Get(
		j,
		"forwardedCookiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ForwardedHeaders() AwsLightsailDistribution_ForwardedHeadersPropertyOutputReference {
	var returns AwsLightsailDistribution_ForwardedHeadersPropertyOutputReference
	_jsii_.Get(
		j,
		"forwardedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ForwardedHeadersInput() *AwsLightsailDistribution_ForwardedHeadersProperty {
	var returns *AwsLightsailDistribution_ForwardedHeadersProperty
	_jsii_.Get(
		j,
		"forwardedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ForwardedQueryStrings() AwsLightsailDistribution_ForwardedQueryStringsPropertyOutputReference {
	var returns AwsLightsailDistribution_ForwardedQueryStringsPropertyOutputReference
	_jsii_.Get(
		j,
		"forwardedQueryStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ForwardedQueryStringsInput() *AwsLightsailDistribution_ForwardedQueryStringsProperty {
	var returns *AwsLightsailDistribution_ForwardedQueryStringsProperty
	_jsii_.Get(
		j,
		"forwardedQueryStringsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) InternalValue() *AwsLightsailDistribution_CacheBehaviorSettingsProperty {
	var returns *AwsLightsailDistribution_CacheBehaviorSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) MaximumTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) MaximumTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) MinimumTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) MinimumTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lightsail.AwsLightsailDistribution.CacheBehaviorSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference_Override(a AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lightsail.AwsLightsailDistribution.CacheBehaviorSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference)SetAllowedHttpMethods(val *string) {
	if err := j.validateSetAllowedHttpMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedHttpMethods",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference)SetCachedHttpMethods(val *string) {
	if err := j.validateSetCachedHttpMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cachedHttpMethods",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference)SetDefaultTtl(val *float64) {
	if err := j.validateSetDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference)SetInternalValue(val *AwsLightsailDistribution_CacheBehaviorSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference)SetMaximumTtl(val *float64) {
	if err := j.validateSetMaximumTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumTtl",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference)SetMinimumTtl(val *float64) {
	if err := j.validateSetMinimumTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumTtl",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) PutForwardedCookies(value *AwsLightsailDistribution_ForwardedCookiesProperty) {
	if err := a.validatePutForwardedCookiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putForwardedCookies",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) PutForwardedHeaders(value *AwsLightsailDistribution_ForwardedHeadersProperty) {
	if err := a.validatePutForwardedHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putForwardedHeaders",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) PutForwardedQueryStrings(value *AwsLightsailDistribution_ForwardedQueryStringsProperty) {
	if err := a.validatePutForwardedQueryStringsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putForwardedQueryStrings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ResetAllowedHttpMethods() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedHttpMethods",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ResetCachedHttpMethods() {
	_jsii_.InvokeVoid(
		a,
		"resetCachedHttpMethods",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ResetForwardedCookies() {
	_jsii_.InvokeVoid(
		a,
		"resetForwardedCookies",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ResetForwardedHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetForwardedHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ResetForwardedQueryStrings() {
	_jsii_.InvokeVoid(
		a,
		"resetForwardedQueryStrings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ResetMaximumTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ResetMinimumTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetMinimumTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLightsailDistribution_CacheBehaviorSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

