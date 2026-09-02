package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappsync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappsync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference interface {
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
	DataSourceLevelMetricsBehavior() *string
	// Experimental.
	SetDataSourceLevelMetricsBehavior(val *string)
	// Experimental.
	DataSourceLevelMetricsBehaviorInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfGraphqlApi_EnhancedMetricsConfigProperty
	// Experimental.
	SetInternalValue(val *TfGraphqlApi_EnhancedMetricsConfigProperty)
	// Experimental.
	OperationLevelMetricsConfig() *string
	// Experimental.
	SetOperationLevelMetricsConfig(val *string)
	// Experimental.
	OperationLevelMetricsConfigInput() *string
	// Experimental.
	ResolverLevelMetricsBehavior() *string
	// Experimental.
	SetResolverLevelMetricsBehavior(val *string)
	// Experimental.
	ResolverLevelMetricsBehaviorInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference
type jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) DataSourceLevelMetricsBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceLevelMetricsBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) DataSourceLevelMetricsBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceLevelMetricsBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) InternalValue() *TfGraphqlApi_EnhancedMetricsConfigProperty {
	var returns *TfGraphqlApi_EnhancedMetricsConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) OperationLevelMetricsConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationLevelMetricsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) OperationLevelMetricsConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationLevelMetricsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) ResolverLevelMetricsBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolverLevelMetricsBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) ResolverLevelMetricsBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolverLevelMetricsBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGraphqlApi_EnhancedMetricsConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appsync.TfGraphqlApi.EnhancedMetricsConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference_Override(t TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appsync.TfGraphqlApi.EnhancedMetricsConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference)SetDataSourceLevelMetricsBehavior(val *string) {
	if err := j.validateSetDataSourceLevelMetricsBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceLevelMetricsBehavior",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference)SetInternalValue(val *TfGraphqlApi_EnhancedMetricsConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference)SetOperationLevelMetricsConfig(val *string) {
	if err := j.validateSetOperationLevelMetricsConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationLevelMetricsConfig",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference)SetResolverLevelMetricsBehavior(val *string) {
	if err := j.validateSetResolverLevelMetricsBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolverLevelMetricsBehavior",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

