package timestreamquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/timestreamquery/jsii"

	"github.com/cdktn-io/cdktn-aws-go/timestreamquery/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference interface {
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
	ErrorReportLocation() AwsScheduledQuery_RecentlyFailedRunsErrorReportLocationPropertyList
	// Experimental.
	ErrorReportLocationInput() interface{}
	// Experimental.
	ExecutionStats() AwsScheduledQuery_RecentlyFailedRunsExecutionStatsPropertyList
	// Experimental.
	ExecutionStatsInput() interface{}
	// Experimental.
	FailureReason() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	InvocationTime() *string
	// Experimental.
	QueryInsightsResponse() AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyList
	// Experimental.
	QueryInsightsResponseInput() interface{}
	// Experimental.
	RunStatus() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TriggerTime() *string
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
	PutErrorReportLocation(value interface{})
	// Experimental.
	PutExecutionStats(value interface{})
	// Experimental.
	PutQueryInsightsResponse(value interface{})
	// Experimental.
	ResetErrorReportLocation()
	// Experimental.
	ResetExecutionStats()
	// Experimental.
	ResetQueryInsightsResponse()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference
type jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) ErrorReportLocation() AwsScheduledQuery_RecentlyFailedRunsErrorReportLocationPropertyList {
	var returns AwsScheduledQuery_RecentlyFailedRunsErrorReportLocationPropertyList
	_jsii_.Get(
		j,
		"errorReportLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) ErrorReportLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorReportLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) ExecutionStats() AwsScheduledQuery_RecentlyFailedRunsExecutionStatsPropertyList {
	var returns AwsScheduledQuery_RecentlyFailedRunsExecutionStatsPropertyList
	_jsii_.Get(
		j,
		"executionStats",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) ExecutionStatsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executionStatsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) FailureReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) InvocationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invocationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) QueryInsightsResponse() AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyList {
	var returns AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyList
	_jsii_.Get(
		j,
		"queryInsightsResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) QueryInsightsResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInsightsResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) RunStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) TriggerTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerTime",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsScheduledQuery_RecentlyFailedRunsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-timestream-query.AwsScheduledQuery.RecentlyFailedRunsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference_Override(a AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-timestream-query.AwsScheduledQuery.RecentlyFailedRunsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) PutErrorReportLocation(value interface{}) {
	if err := a.validatePutErrorReportLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putErrorReportLocation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) PutExecutionStats(value interface{}) {
	if err := a.validatePutExecutionStatsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExecutionStats",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) PutQueryInsightsResponse(value interface{}) {
	if err := a.validatePutQueryInsightsResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueryInsightsResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) ResetErrorReportLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetErrorReportLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) ResetExecutionStats() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionStats",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) ResetQueryInsightsResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryInsightsResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

