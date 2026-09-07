package iotcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/iotcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/iotcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference interface {
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
	InternalValue() *AwsTopicRule_ErrorActionCloudwatchMetricProperty
	// Experimental.
	SetInternalValue(val *AwsTopicRule_ErrorActionCloudwatchMetricProperty)
	// Experimental.
	MetricName() *string
	// Experimental.
	SetMetricName(val *string)
	// Experimental.
	MetricNameInput() *string
	// Experimental.
	MetricNamespace() *string
	// Experimental.
	SetMetricNamespace(val *string)
	// Experimental.
	MetricNamespaceInput() *string
	// Experimental.
	MetricTimestamp() *string
	// Experimental.
	SetMetricTimestamp(val *string)
	// Experimental.
	MetricTimestampInput() *string
	// Experimental.
	MetricUnit() *string
	// Experimental.
	SetMetricUnit(val *string)
	// Experimental.
	MetricUnitInput() *string
	// Experimental.
	MetricValue() *string
	// Experimental.
	SetMetricValue(val *string)
	// Experimental.
	MetricValueInput() *string
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
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
	ResetMetricTimestamp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference
type jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) InternalValue() *AwsTopicRule_ErrorActionCloudwatchMetricProperty {
	var returns *AwsTopicRule_ErrorActionCloudwatchMetricProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricTimestampInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) MetricValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-iot-core.AwsTopicRule.ErrorActionCloudwatchMetricPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference_Override(a AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iot-core.AwsTopicRule.ErrorActionCloudwatchMetricPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetInternalValue(val *AwsTopicRule_ErrorActionCloudwatchMetricProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetMetricName(val *string) {
	if err := j.validateSetMetricNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricName",
		val,
	)
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetMetricNamespace(val *string) {
	if err := j.validateSetMetricNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricNamespace",
		val,
	)
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetMetricTimestamp(val *string) {
	if err := j.validateSetMetricTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricTimestamp",
		val,
	)
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetMetricUnit(val *string) {
	if err := j.validateSetMetricUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricUnit",
		val,
	)
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetMetricValue(val *string) {
	if err := j.validateSetMetricValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricValue",
		val,
	)
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) ResetMetricTimestamp() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricTimestamp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

