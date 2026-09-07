package timestreamquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/timestreamquery/jsii"

	"github.com/cdktn-io/cdktn-aws-go/timestreamquery/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	OutputBytes() *float64
	// Experimental.
	OutputRows() *float64
	// Experimental.
	QuerySpatialCoverage() AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponseQuerySpatialCoveragePropertyList
	// Experimental.
	QuerySpatialCoverageInput() interface{}
	// Experimental.
	QueryTableCount() *float64
	// Experimental.
	QueryTemporalRange() AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponseQueryTemporalRangePropertyList
	// Experimental.
	QueryTemporalRangeInput() interface{}
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
	PutQuerySpatialCoverage(value interface{})
	// Experimental.
	PutQueryTemporalRange(value interface{})
	// Experimental.
	ResetQuerySpatialCoverage()
	// Experimental.
	ResetQueryTemporalRange()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference
type jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) OutputBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outputBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) OutputRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"outputRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) QuerySpatialCoverage() AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponseQuerySpatialCoveragePropertyList {
	var returns AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponseQuerySpatialCoveragePropertyList
	_jsii_.Get(
		j,
		"querySpatialCoverage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) QuerySpatialCoverageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"querySpatialCoverageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) QueryTableCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryTableCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) QueryTemporalRange() AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponseQueryTemporalRangePropertyList {
	var returns AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponseQueryTemporalRangePropertyList
	_jsii_.Get(
		j,
		"queryTemporalRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) QueryTemporalRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryTemporalRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-timestream-query.AwsScheduledQuery.RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference_Override(a AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-timestream-query.AwsScheduledQuery.RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) PutQuerySpatialCoverage(value interface{}) {
	if err := a.validatePutQuerySpatialCoverageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQuerySpatialCoverage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) PutQueryTemporalRange(value interface{}) {
	if err := a.validatePutQueryTemporalRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQueryTemporalRange",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) ResetQuerySpatialCoverage() {
	_jsii_.InvokeVoid(
		a,
		"resetQuerySpatialCoverage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) ResetQueryTemporalRange() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryTemporalRange",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsScheduledQuery_RecentlyFailedRunsQueryInsightsResponsePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

