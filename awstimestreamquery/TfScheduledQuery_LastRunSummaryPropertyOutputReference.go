package awstimestreamquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstimestreamquery/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstimestreamquery/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfScheduledQuery_LastRunSummaryPropertyOutputReference interface {
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
	ErrorReportLocation() TfScheduledQuery_LastRunSummaryErrorReportLocationPropertyList
	// Experimental.
	ErrorReportLocationInput() interface{}
	// Experimental.
	ExecutionStats() TfScheduledQuery_LastRunSummaryExecutionStatsPropertyList
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
	QueryInsightsResponse() TfScheduledQuery_LastRunSummaryQueryInsightsResponsePropertyList
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

// The jsii proxy struct for TfScheduledQuery_LastRunSummaryPropertyOutputReference
type jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) ErrorReportLocation() TfScheduledQuery_LastRunSummaryErrorReportLocationPropertyList {
	var returns TfScheduledQuery_LastRunSummaryErrorReportLocationPropertyList
	_jsii_.Get(
		j,
		"errorReportLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) ErrorReportLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"errorReportLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) ExecutionStats() TfScheduledQuery_LastRunSummaryExecutionStatsPropertyList {
	var returns TfScheduledQuery_LastRunSummaryExecutionStatsPropertyList
	_jsii_.Get(
		j,
		"executionStats",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) ExecutionStatsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"executionStatsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) FailureReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) InvocationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invocationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) QueryInsightsResponse() TfScheduledQuery_LastRunSummaryQueryInsightsResponsePropertyList {
	var returns TfScheduledQuery_LastRunSummaryQueryInsightsResponsePropertyList
	_jsii_.Get(
		j,
		"queryInsightsResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) QueryInsightsResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryInsightsResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) RunStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) TriggerTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerTime",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfScheduledQuery_LastRunSummaryPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfScheduledQuery_LastRunSummaryPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfScheduledQuery_LastRunSummaryPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-timestream-query.TfScheduledQuery.LastRunSummaryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfScheduledQuery_LastRunSummaryPropertyOutputReference_Override(t TfScheduledQuery_LastRunSummaryPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-timestream-query.TfScheduledQuery.LastRunSummaryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) PutErrorReportLocation(value interface{}) {
	if err := t.validatePutErrorReportLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putErrorReportLocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) PutExecutionStats(value interface{}) {
	if err := t.validatePutExecutionStatsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExecutionStats",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) PutQueryInsightsResponse(value interface{}) {
	if err := t.validatePutQueryInsightsResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putQueryInsightsResponse",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) ResetErrorReportLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetErrorReportLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) ResetExecutionStats() {
	_jsii_.InvokeVoid(
		t,
		"resetExecutionStats",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) ResetQueryInsightsResponse() {
	_jsii_.InvokeVoid(
		t,
		"resetQueryInsightsResponse",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfScheduledQuery_LastRunSummaryPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

