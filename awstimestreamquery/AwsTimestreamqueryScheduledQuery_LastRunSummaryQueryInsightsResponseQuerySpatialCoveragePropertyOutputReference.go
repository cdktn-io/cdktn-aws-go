package awstimestreamquery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstimestreamquery/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstimestreamquery/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference interface {
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
	Max() AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxPropertyList
	// Experimental.
	MaxInput() interface{}
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
	PutMax(value interface{})
	// Experimental.
	ResetMax()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference
type jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) Max() AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxPropertyList {
	var returns AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxPropertyList
	_jsii_.Get(
		j,
		"max",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) MaxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-timestream-query.AwsTimestreamqueryScheduledQuery.LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference_Override(a AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-timestream-query.AwsTimestreamqueryScheduledQuery.LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) PutMax(value interface{}) {
	if err := a.validatePutMaxParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMax",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) ResetMax() {
	_jsii_.InvokeVoid(
		a,
		"resetMax",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoveragePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

