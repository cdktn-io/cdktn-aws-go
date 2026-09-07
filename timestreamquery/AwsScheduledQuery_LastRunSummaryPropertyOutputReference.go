package timestreamquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/timestreamquery/jsii"

	"github.com/cdktn-io/cdktn-aws-go/timestreamquery/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsScheduledQuery_LastRunSummaryPropertyOutputReference interface {
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
	ErrorReportLocation() AwsScheduledQuery_LastRunSummaryErrorReportLocationPropertyList
	// Experimental.
	ErrorReportLocationInput() interface{}
	// Experimental.
	ExecutionStats() AwsScheduledQuery_LastRunSummaryExecutionStatsPropertyList
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
	QueryInsightsResponse() AwsScheduledQuery_LastRunSummaryQueryInsightsResponsePropertyList
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

// The jsii proxy struct for AwsScheduledQuery_LastRunSummaryPropertyOutputReference
type jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) ErrorReportLocation() AwsScheduledQuery_LastRunSummaryErrorReportLocationPropertyList {
	var returns AwsScheduledQuery_LastRunSummaryErrorReportLocationPropertyList
	_jsii_.Get(
		j,
		"errorReportLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) ErrorReportLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorReportLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) ExecutionStats() AwsScheduledQuery_LastRunSummaryExecutionStatsPropertyList {
	var returns AwsScheduledQuery_LastRunSummaryExecutionStatsPropertyList
	_jsii_.Get(
		j,
		"executionStats",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) ExecutionStatsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executionStatsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) FailureReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) InvocationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invocationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) QueryInsightsResponse() AwsScheduledQuery_LastRunSummaryQueryInsightsResponsePropertyList {
	var returns AwsScheduledQuery_LastRunSummaryQueryInsightsResponsePropertyList
	_jsii_.Get(
		j,
		"queryInsightsResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) QueryInsightsResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInsightsResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) RunStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) TriggerTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerTime",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsScheduledQuery_LastRunSummaryPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsScheduledQuery_LastRunSummaryPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsScheduledQuery_LastRunSummaryPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-timestream-query.AwsScheduledQuery.LastRunSummaryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsScheduledQuery_LastRunSummaryPropertyOutputReference_Override(a AwsScheduledQuery_LastRunSummaryPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-timestream-query.AwsScheduledQuery.LastRunSummaryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) PutErrorReportLocation(value interface{}) {
	if err := a.validatePutErrorReportLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putErrorReportLocation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) PutExecutionStats(value interface{}) {
	if err := a.validatePutExecutionStatsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExecutionStats",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) PutQueryInsightsResponse(value interface{}) {
	if err := a.validatePutQueryInsightsResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueryInsightsResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) ResetErrorReportLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetErrorReportLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) ResetExecutionStats() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionStats",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) ResetQueryInsightsResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryInsightsResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsScheduledQuery_LastRunSummaryPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

