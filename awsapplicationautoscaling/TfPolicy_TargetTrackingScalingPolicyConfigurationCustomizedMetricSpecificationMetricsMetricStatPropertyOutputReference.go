package awsapplicationautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapplicationautoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsapplicationautoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference interface {
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
	InternalValue() *TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatProperty
	// Experimental.
	SetInternalValue(val *TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatProperty)
	// Experimental.
	Metric() TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricPropertyOutputReference
	// Experimental.
	MetricInput() *TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricProperty
	// Experimental.
	Stat() *string
	// Experimental.
	SetStat(val *string)
	// Experimental.
	StatInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Unit() *string
	// Experimental.
	SetUnit(val *string)
	// Experimental.
	UnitInput() *string
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
	PutMetric(value *TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricProperty)
	// Experimental.
	ResetUnit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference
type jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) InternalValue() *TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatProperty {
	var returns *TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) Metric() TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricPropertyOutputReference {
	var returns TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricPropertyOutputReference
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) MetricInput() *TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricProperty {
	var returns *TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricProperty
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) Stat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) StatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-application-auto-scaling.TfPolicy.TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference_Override(t TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-application-auto-scaling.TfPolicy.TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference)SetInternalValue(val *TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference)SetStat(val *string) {
	if err := j.validateSetStatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stat",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference)SetUnit(val *string) {
	if err := j.validateSetUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) PutMetric(value *TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricProperty) {
	if err := t.validatePutMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMetric",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) ResetUnit() {
	_jsii_.InvokeVoid(
		t,
		"resetUnit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

