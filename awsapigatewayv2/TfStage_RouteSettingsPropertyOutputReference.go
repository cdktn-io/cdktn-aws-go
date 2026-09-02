package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapigatewayv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsapigatewayv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfStage_RouteSettingsPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LoggingLevel() *string
	// Experimental.
	SetLoggingLevel(val *string)
	// Experimental.
	LoggingLevelInput() *string
	// Experimental.
	RouteKey() *string
	// Experimental.
	SetRouteKey(val *string)
	// Experimental.
	RouteKeyInput() *string
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

// The jsii proxy struct for TfStage_RouteSettingsPropertyOutputReference
type jsiiProxy_TfStage_RouteSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) DataTraceEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTraceEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) DataTraceEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTraceEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) DetailedMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detailedMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) DetailedMetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detailedMetricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) LoggingLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) LoggingLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) RouteKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) RouteKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) ThrottlingBurstLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingBurstLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) ThrottlingBurstLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingBurstLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) ThrottlingRateLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingRateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) ThrottlingRateLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throttlingRateLimitInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfStage_RouteSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfStage_RouteSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfStage_RouteSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfStage_RouteSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-api-gateway-v2.TfStage.RouteSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfStage_RouteSettingsPropertyOutputReference_Override(t TfStage_RouteSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-api-gateway-v2.TfStage.RouteSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference)SetDataTraceEnabled(val interface{}) {
	if err := j.validateSetDataTraceEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTraceEnabled",
		val,
	)
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference)SetDetailedMetricsEnabled(val interface{}) {
	if err := j.validateSetDetailedMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detailedMetricsEnabled",
		val,
	)
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference)SetLoggingLevel(val *string) {
	if err := j.validateSetLoggingLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingLevel",
		val,
	)
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference)SetRouteKey(val *string) {
	if err := j.validateSetRouteKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routeKey",
		val,
	)
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference)SetThrottlingBurstLimit(val *float64) {
	if err := j.validateSetThrottlingBurstLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttlingBurstLimit",
		val,
	)
}

func (j *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference)SetThrottlingRateLimit(val *float64) {
	if err := j.validateSetThrottlingRateLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throttlingRateLimit",
		val,
	)
}

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) ResetDataTraceEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetDataTraceEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) ResetDetailedMetricsEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetDetailedMetricsEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) ResetLoggingLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetLoggingLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) ResetThrottlingBurstLimit() {
	_jsii_.InvokeVoid(
		t,
		"resetThrottlingBurstLimit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) ResetThrottlingRateLimit() {
	_jsii_.InvokeVoid(
		t,
		"resetThrottlingRateLimit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfStage_RouteSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

