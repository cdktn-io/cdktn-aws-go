package apigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/apigateway/jsii"

	"github.com/cdktn-io/cdktn-aws-go/apigateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMethodSettings_SettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CacheDataEncrypted() interface{}
	// Experimental.
	SetCacheDataEncrypted(val interface{})
	// Experimental.
	CacheDataEncryptedInput() interface{}
	// Experimental.
	CacheTtlInSeconds() *float64
	// Experimental.
	SetCacheTtlInSeconds(val *float64)
	// Experimental.
	CacheTtlInSecondsInput() *float64
	// Experimental.
	CachingEnabled() interface{}
	// Experimental.
	SetCachingEnabled(val interface{})
	// Experimental.
	CachingEnabledInput() interface{}
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
	DataTraceEnabled() interface{}
	// Experimental.
	SetDataTraceEnabled(val interface{})
	// Experimental.
	DataTraceEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsMethodSettings_SettingsProperty
	// Experimental.
	SetInternalValue(val *AwsMethodSettings_SettingsProperty)
	// Experimental.
	LoggingLevel() *string
	// Experimental.
	SetLoggingLevel(val *string)
	// Experimental.
	LoggingLevelInput() *string
	// Experimental.
	MetricsEnabled() interface{}
	// Experimental.
	SetMetricsEnabled(val interface{})
	// Experimental.
	MetricsEnabledInput() interface{}
	// Experimental.
	RequireAuthorizationForCacheControl() interface{}
	// Experimental.
	SetRequireAuthorizationForCacheControl(val interface{})
	// Experimental.
	RequireAuthorizationForCacheControlInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ThrottlingBurstLimit() *float64
	// Experimental.
	SetThrottlingBurstLimit(val *float64)
	// Experimental.
	ThrottlingBurstLimitInput() *float64
	// Experimental.
	ThrottlingRateLimit() *float64
	// Experimental.
	SetThrottlingRateLimit(val *float64)
	// Experimental.
	ThrottlingRateLimitInput() *float64
	// Experimental.
	UnauthorizedCacheControlHeaderStrategy() *string
	// Experimental.
	SetUnauthorizedCacheControlHeaderStrategy(val *string)
	// Experimental.
	UnauthorizedCacheControlHeaderStrategyInput() *string
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
	ResetCacheDataEncrypted()
	// Experimental.
	ResetCacheTtlInSeconds()
	// Experimental.
	ResetCachingEnabled()
	// Experimental.
	ResetDataTraceEnabled()
	// Experimental.
	ResetLoggingLevel()
	// Experimental.
	ResetMetricsEnabled()
	// Experimental.
	ResetRequireAuthorizationForCacheControl()
	// Experimental.
	ResetThrottlingBurstLimit()
	// Experimental.
	ResetThrottlingRateLimit()
	// Experimental.
	ResetUnauthorizedCacheControlHeaderStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMethodSettings_SettingsPropertyOutputReference
type jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) CacheDataEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheDataEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) CacheDataEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheDataEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) CacheTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) CacheTtlInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheTtlInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) CachingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cachingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) CachingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cachingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) DataTraceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTraceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) DataTraceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTraceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) InternalValue() *AwsMethodSettings_SettingsProperty {
	var returns *AwsMethodSettings_SettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) LoggingLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) LoggingLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) MetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) MetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) RequireAuthorizationForCacheControl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAuthorizationForCacheControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) RequireAuthorizationForCacheControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAuthorizationForCacheControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ThrottlingBurstLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingBurstLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ThrottlingBurstLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingBurstLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ThrottlingRateLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingRateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ThrottlingRateLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingRateLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) UnauthorizedCacheControlHeaderStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthorizedCacheControlHeaderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) UnauthorizedCacheControlHeaderStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthorizedCacheControlHeaderStrategyInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMethodSettings_SettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMethodSettings_SettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMethodSettings_SettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-api-gateway.AwsMethodSettings.SettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMethodSettings_SettingsPropertyOutputReference_Override(a AwsMethodSettings_SettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-api-gateway.AwsMethodSettings.SettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference)SetCacheDataEncrypted(val interface{}) {
	if err := j.validateSetCacheDataEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheDataEncrypted",
		val,
	)
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference)SetCacheTtlInSeconds(val *float64) {
	if err := j.validateSetCacheTtlInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference)SetCachingEnabled(val interface{}) {
	if err := j.validateSetCachingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cachingEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference)SetDataTraceEnabled(val interface{}) {
	if err := j.validateSetDataTraceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTraceEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference)SetInternalValue(val *AwsMethodSettings_SettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference)SetLoggingLevel(val *string) {
	if err := j.validateSetLoggingLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingLevel",
		val,
	)
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference)SetMetricsEnabled(val interface{}) {
	if err := j.validateSetMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference)SetRequireAuthorizationForCacheControl(val interface{}) {
	if err := j.validateSetRequireAuthorizationForCacheControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireAuthorizationForCacheControl",
		val,
	)
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference)SetThrottlingBurstLimit(val *float64) {
	if err := j.validateSetThrottlingBurstLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttlingBurstLimit",
		val,
	)
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference)SetThrottlingRateLimit(val *float64) {
	if err := j.validateSetThrottlingRateLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttlingRateLimit",
		val,
	)
}

func (j *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference)SetUnauthorizedCacheControlHeaderStrategy(val *string) {
	if err := j.validateSetUnauthorizedCacheControlHeaderStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unauthorizedCacheControlHeaderStrategy",
		val,
	)
}

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ResetCacheDataEncrypted() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheDataEncrypted",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ResetCacheTtlInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheTtlInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ResetCachingEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetCachingEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ResetDataTraceEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDataTraceEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ResetLoggingLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetLoggingLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ResetMetricsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ResetRequireAuthorizationForCacheControl() {
	_jsii_.InvokeVoid(
		a,
		"resetRequireAuthorizationForCacheControl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ResetThrottlingBurstLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottlingBurstLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ResetThrottlingRateLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottlingRateLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ResetUnauthorizedCacheControlHeaderStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetUnauthorizedCacheControlHeaderStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMethodSettings_SettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

