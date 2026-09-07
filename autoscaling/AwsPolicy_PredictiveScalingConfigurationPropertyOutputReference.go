package autoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/autoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/autoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference interface {
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
	InternalValue() *AwsPolicy_PredictiveScalingConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsPolicy_PredictiveScalingConfigurationProperty)
	// Experimental.
	MaxCapacityBreachBehavior() *string
	// Experimental.
	SetMaxCapacityBreachBehavior(val *string)
	// Experimental.
	MaxCapacityBreachBehaviorInput() *string
	// Experimental.
	MaxCapacityBuffer() *string
	// Experimental.
	SetMaxCapacityBuffer(val *string)
	// Experimental.
	MaxCapacityBufferInput() *string
	// Experimental.
	MetricSpecification() AwsPolicy_MetricSpecificationPropertyOutputReference
	// Experimental.
	MetricSpecificationInput() *AwsPolicy_MetricSpecificationProperty
	// Experimental.
	Mode() *string
	// Experimental.
	SetMode(val *string)
	// Experimental.
	ModeInput() *string
	// Experimental.
	SchedulingBufferTime() *string
	// Experimental.
	SetSchedulingBufferTime(val *string)
	// Experimental.
	SchedulingBufferTimeInput() *string
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
	PutMetricSpecification(value *AwsPolicy_MetricSpecificationProperty)
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

// The jsii proxy struct for AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference
type jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) InternalValue() *AwsPolicy_PredictiveScalingConfigurationProperty {
	var returns *AwsPolicy_PredictiveScalingConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) MaxCapacityBreachBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxCapacityBreachBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) MaxCapacityBreachBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxCapacityBreachBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) MaxCapacityBuffer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxCapacityBuffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) MaxCapacityBufferInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxCapacityBufferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) MetricSpecification() AwsPolicy_MetricSpecificationPropertyOutputReference {
	var returns AwsPolicy_MetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"metricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) MetricSpecificationInput() *AwsPolicy_MetricSpecificationProperty {
	var returns *AwsPolicy_MetricSpecificationProperty
	_jsii_.Get(
		j,
		"metricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) SchedulingBufferTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingBufferTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) SchedulingBufferTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingBufferTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPolicy_PredictiveScalingConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPolicy_PredictiveScalingConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.AwsPolicy.PredictiveScalingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPolicy_PredictiveScalingConfigurationPropertyOutputReference_Override(a AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.AwsPolicy.PredictiveScalingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference)SetInternalValue(val *AwsPolicy_PredictiveScalingConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference)SetMaxCapacityBreachBehavior(val *string) {
	if err := j.validateSetMaxCapacityBreachBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacityBreachBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference)SetMaxCapacityBuffer(val *string) {
	if err := j.validateSetMaxCapacityBufferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacityBuffer",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference)SetSchedulingBufferTime(val *string) {
	if err := j.validateSetSchedulingBufferTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulingBufferTime",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) PutMetricSpecification(value *AwsPolicy_MetricSpecificationProperty) {
	if err := a.validatePutMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) ResetMaxCapacityBreachBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxCapacityBreachBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) ResetMaxCapacityBuffer() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxCapacityBuffer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		a,
		"resetMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) ResetSchedulingBufferTime() {
	_jsii_.InvokeVoid(
		a,
		"resetSchedulingBufferTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPolicy_PredictiveScalingConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

