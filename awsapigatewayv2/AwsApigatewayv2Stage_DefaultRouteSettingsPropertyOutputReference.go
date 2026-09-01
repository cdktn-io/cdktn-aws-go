package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapigatewayv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsapigatewayv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference interface {
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
	DetailedMetricsEnabled() interface{}
	// Experimental.
	SetDetailedMetricsEnabled(val interface{})
	// Experimental.
	DetailedMetricsEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsApigatewayv2Stage_DefaultRouteSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsApigatewayv2Stage_DefaultRouteSettingsProperty)
	// Experimental.
	LoggingLevel() *string
	// Experimental.
	SetLoggingLevel(val *string)
	// Experimental.
	LoggingLevelInput() *string
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
	ResetDataTraceEnabled()
	// Experimental.
	ResetDetailedMetricsEnabled()
	// Experimental.
	ResetLoggingLevel()
	// Experimental.
	ResetThrottlingBurstLimit()
	// Experimental.
	ResetThrottlingRateLimit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference
type jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) DataTraceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTraceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) DataTraceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTraceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) DetailedMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detailedMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) DetailedMetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detailedMetricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) InternalValue() *AwsApigatewayv2Stage_DefaultRouteSettingsProperty {
	var returns *AwsApigatewayv2Stage_DefaultRouteSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) LoggingLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) LoggingLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) ThrottlingBurstLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingBurstLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) ThrottlingBurstLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingBurstLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) ThrottlingRateLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingRateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) ThrottlingRateLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingRateLimitInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-api-gateway-v2.AwsApigatewayv2Stage.DefaultRouteSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference_Override(a AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-api-gateway-v2.AwsApigatewayv2Stage.DefaultRouteSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference)SetDataTraceEnabled(val interface{}) {
	if err := j.validateSetDataTraceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTraceEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference)SetDetailedMetricsEnabled(val interface{}) {
	if err := j.validateSetDetailedMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detailedMetricsEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference)SetInternalValue(val *AwsApigatewayv2Stage_DefaultRouteSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference)SetLoggingLevel(val *string) {
	if err := j.validateSetLoggingLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingLevel",
		val,
	)
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference)SetThrottlingBurstLimit(val *float64) {
	if err := j.validateSetThrottlingBurstLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttlingBurstLimit",
		val,
	)
}

func (j *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference)SetThrottlingRateLimit(val *float64) {
	if err := j.validateSetThrottlingRateLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttlingRateLimit",
		val,
	)
}

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) ResetDataTraceEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDataTraceEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) ResetDetailedMetricsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDetailedMetricsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) ResetLoggingLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetLoggingLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) ResetThrottlingBurstLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottlingBurstLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) ResetThrottlingRateLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetThrottlingRateLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsApigatewayv2Stage_DefaultRouteSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

