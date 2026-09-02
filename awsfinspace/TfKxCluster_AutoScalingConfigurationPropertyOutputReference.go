package awsfinspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfinspace/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfinspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfKxCluster_AutoScalingConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutoScalingMetric() *string
	// Experimental.
	SetAutoScalingMetric(val *string)
	// Experimental.
	AutoScalingMetricInput() *string
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
	InternalValue() *TfKxCluster_AutoScalingConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfKxCluster_AutoScalingConfigurationProperty)
	// Experimental.
	MaxNodeCount() *float64
	// Experimental.
	SetMaxNodeCount(val *float64)
	// Experimental.
	MaxNodeCountInput() *float64
	// Experimental.
	MetricTarget() *float64
	// Experimental.
	SetMetricTarget(val *float64)
	// Experimental.
	MetricTargetInput() *float64
	// Experimental.
	MinNodeCount() *float64
	// Experimental.
	SetMinNodeCount(val *float64)
	// Experimental.
	MinNodeCountInput() *float64
	// Experimental.
	ScaleInCooldownSeconds() *float64
	// Experimental.
	SetScaleInCooldownSeconds(val *float64)
	// Experimental.
	ScaleInCooldownSecondsInput() *float64
	// Experimental.
	ScaleOutCooldownSeconds() *float64
	// Experimental.
	SetScaleOutCooldownSeconds(val *float64)
	// Experimental.
	ScaleOutCooldownSecondsInput() *float64
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

// The jsii proxy struct for TfKxCluster_AutoScalingConfigurationPropertyOutputReference
type jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) AutoScalingMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) AutoScalingMetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) InternalValue() *TfKxCluster_AutoScalingConfigurationProperty {
	var returns *TfKxCluster_AutoScalingConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) MaxNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) MaxNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) MetricTarget() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) MetricTargetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"metricTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) MinNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) MinNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) ScaleInCooldownSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInCooldownSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) ScaleInCooldownSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInCooldownSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) ScaleOutCooldownSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutCooldownSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) ScaleOutCooldownSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutCooldownSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfKxCluster_AutoScalingConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfKxCluster_AutoScalingConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfKxCluster_AutoScalingConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-finspace.TfKxCluster.AutoScalingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfKxCluster_AutoScalingConfigurationPropertyOutputReference_Override(t TfKxCluster_AutoScalingConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-finspace.TfKxCluster.AutoScalingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference)SetAutoScalingMetric(val *string) {
	if err := j.validateSetAutoScalingMetricParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingMetric",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference)SetInternalValue(val *TfKxCluster_AutoScalingConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference)SetMaxNodeCount(val *float64) {
	if err := j.validateSetMaxNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxNodeCount",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference)SetMetricTarget(val *float64) {
	if err := j.validateSetMetricTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricTarget",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference)SetMinNodeCount(val *float64) {
	if err := j.validateSetMinNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNodeCount",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference)SetScaleInCooldownSeconds(val *float64) {
	if err := j.validateSetScaleInCooldownSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleInCooldownSeconds",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference)SetScaleOutCooldownSeconds(val *float64) {
	if err := j.validateSetScaleOutCooldownSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleOutCooldownSeconds",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfKxCluster_AutoScalingConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

