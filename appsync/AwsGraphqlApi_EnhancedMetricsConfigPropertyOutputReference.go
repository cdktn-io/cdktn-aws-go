package appsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appsync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appsync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference interface {
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
	InternalValue() *AwsGraphqlApi_EnhancedMetricsConfigProperty
	// Experimental.
	SetInternalValue(val *AwsGraphqlApi_EnhancedMetricsConfigProperty)
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

// The jsii proxy struct for AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference
type jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) DataSourceLevelMetricsBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceLevelMetricsBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) DataSourceLevelMetricsBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceLevelMetricsBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) InternalValue() *AwsGraphqlApi_EnhancedMetricsConfigProperty {
	var returns *AwsGraphqlApi_EnhancedMetricsConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) OperationLevelMetricsConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationLevelMetricsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) OperationLevelMetricsConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationLevelMetricsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) ResolverLevelMetricsBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolverLevelMetricsBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) ResolverLevelMetricsBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolverLevelMetricsBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsGraphqlApi.EnhancedMetricsConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference_Override(a AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsGraphqlApi.EnhancedMetricsConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference)SetDataSourceLevelMetricsBehavior(val *string) {
	if err := j.validateSetDataSourceLevelMetricsBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceLevelMetricsBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference)SetInternalValue(val *AwsGraphqlApi_EnhancedMetricsConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference)SetOperationLevelMetricsConfig(val *string) {
	if err := j.validateSetOperationLevelMetricsConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationLevelMetricsConfig",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference)SetResolverLevelMetricsBehavior(val *string) {
	if err := j.validateSetResolverLevelMetricsBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolverLevelMetricsBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

