package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsautoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsautoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference interface {
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
	InternalValue() *TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatProperty
	// Experimental.
	SetInternalValue(val *TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatProperty)
	// Experimental.
	Metric() TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricPropertyOutputReference
	// Experimental.
	MetricInput() *TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricProperty
	// Experimental.
	Period() *float64
	// Experimental.
	SetPeriod(val *float64)
	// Experimental.
	PeriodInput() *float64
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
	PutMetric(value *TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricProperty)
	// Experimental.
	ResetPeriod()
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

// The jsii proxy struct for TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference
type jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) InternalValue() *TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatProperty {
	var returns *TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) Metric() TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricPropertyOutputReference {
	var returns TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricPropertyOutputReference
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) MetricInput() *TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricProperty {
	var returns *TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricProperty
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) PeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) Stat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) StatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) Unit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) UnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.TfPolicy.TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference_Override(t TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.TfPolicy.TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference)SetInternalValue(val *TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference)SetPeriod(val *float64) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference)SetStat(val *string) {
	if err := j.validateSetStatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stat",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference)SetUnit(val *string) {
	if err := j.validateSetUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unit",
		val,
	)
}

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) PutMetric(value *TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricProperty) {
	if err := t.validatePutMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMetric",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) ResetPeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetPeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) ResetUnit() {
	_jsii_.InvokeVoid(
		t,
		"resetUnit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

