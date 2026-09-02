package awsapplicationautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapplicationautoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsapplicationautoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference interface {
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
	InternalValue() *TfPolicy_PredictiveScalingPolicyConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfPolicy_PredictiveScalingPolicyConfigurationProperty)
	// Experimental.
	MaxCapacityBreachBehavior() *string
	// Experimental.
	SetMaxCapacityBreachBehavior(val *string)
	// Experimental.
	MaxCapacityBreachBehaviorInput() *string
	// Experimental.
	MaxCapacityBuffer() *float64
	// Experimental.
	SetMaxCapacityBuffer(val *float64)
	// Experimental.
	MaxCapacityBufferInput() *float64
	// Experimental.
	MetricSpecification() TfPolicy_MetricSpecificationPropertyList
	// Experimental.
	MetricSpecificationInput() interface{}
	// Experimental.
	Mode() *string
	// Experimental.
	SetMode(val *string)
	// Experimental.
	ModeInput() *string
	// Experimental.
	SchedulingBufferTime() *float64
	// Experimental.
	SetSchedulingBufferTime(val *float64)
	// Experimental.
	SchedulingBufferTimeInput() *float64
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
	PutMetricSpecification(value interface{})
	// Experimental.
	ResetMaxCapacityBreachBehavior()
	// Experimental.
	ResetMaxCapacityBuffer()
	// Experimental.
	ResetMode()
	// Experimental.
	ResetSchedulingBufferTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference
type jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) InternalValue() *TfPolicy_PredictiveScalingPolicyConfigurationProperty {
	var returns *TfPolicy_PredictiveScalingPolicyConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) MaxCapacityBreachBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxCapacityBreachBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) MaxCapacityBreachBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxCapacityBreachBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) MaxCapacityBuffer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityBuffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) MaxCapacityBufferInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityBufferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) MetricSpecification() TfPolicy_MetricSpecificationPropertyList {
	var returns TfPolicy_MetricSpecificationPropertyList
	_jsii_.Get(
		j,
		"metricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) MetricSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) SchedulingBufferTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulingBufferTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) SchedulingBufferTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulingBufferTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-application-auto-scaling.TfPolicy.PredictiveScalingPolicyConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference_Override(t TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-application-auto-scaling.TfPolicy.PredictiveScalingPolicyConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference)SetInternalValue(val *TfPolicy_PredictiveScalingPolicyConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference)SetMaxCapacityBreachBehavior(val *string) {
	if err := j.validateSetMaxCapacityBreachBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacityBreachBehavior",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference)SetMaxCapacityBuffer(val *float64) {
	if err := j.validateSetMaxCapacityBufferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacityBuffer",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference)SetSchedulingBufferTime(val *float64) {
	if err := j.validateSetSchedulingBufferTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulingBufferTime",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) PutMetricSpecification(value interface{}) {
	if err := t.validatePutMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMetricSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) ResetMaxCapacityBreachBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxCapacityBreachBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) ResetMaxCapacityBuffer() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxCapacityBuffer",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		t,
		"resetMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) ResetSchedulingBufferTime() {
	_jsii_.InvokeVoid(
		t,
		"resetSchedulingBufferTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPolicy_PredictiveScalingPolicyConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

