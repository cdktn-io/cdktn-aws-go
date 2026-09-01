package awsautoscalingplans

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsautoscalingplans/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsautoscalingplans/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference interface {
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
	CustomizedLoadMetricSpecification() AwsAutoscalingplansScalingPlan_CustomizedLoadMetricSpecificationPropertyOutputReference
	// Experimental.
	CustomizedLoadMetricSpecificationInput() *AwsAutoscalingplansScalingPlan_CustomizedLoadMetricSpecificationProperty
	// Experimental.
	DisableDynamicScaling() interface{}
	// Experimental.
	SetDisableDynamicScaling(val interface{})
	// Experimental.
	DisableDynamicScalingInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MaxCapacity() *float64
	// Experimental.
	SetMaxCapacity(val *float64)
	// Experimental.
	MaxCapacityInput() *float64
	// Experimental.
	MinCapacity() *float64
	// Experimental.
	SetMinCapacity(val *float64)
	// Experimental.
	MinCapacityInput() *float64
	// Experimental.
	PredefinedLoadMetricSpecification() AwsAutoscalingplansScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference
	// Experimental.
	PredefinedLoadMetricSpecificationInput() *AwsAutoscalingplansScalingPlan_PredefinedLoadMetricSpecificationProperty
	// Experimental.
	PredictiveScalingMaxCapacityBehavior() *string
	// Experimental.
	SetPredictiveScalingMaxCapacityBehavior(val *string)
	// Experimental.
	PredictiveScalingMaxCapacityBehaviorInput() *string
	// Experimental.
	PredictiveScalingMaxCapacityBuffer() *float64
	// Experimental.
	SetPredictiveScalingMaxCapacityBuffer(val *float64)
	// Experimental.
	PredictiveScalingMaxCapacityBufferInput() *float64
	// Experimental.
	PredictiveScalingMode() *string
	// Experimental.
	SetPredictiveScalingMode(val *string)
	// Experimental.
	PredictiveScalingModeInput() *string
	// Experimental.
	ResourceId() *string
	// Experimental.
	SetResourceId(val *string)
	// Experimental.
	ResourceIdInput() *string
	// Experimental.
	ScalableDimension() *string
	// Experimental.
	SetScalableDimension(val *string)
	// Experimental.
	ScalableDimensionInput() *string
	// Experimental.
	ScalingPolicyUpdateBehavior() *string
	// Experimental.
	SetScalingPolicyUpdateBehavior(val *string)
	// Experimental.
	ScalingPolicyUpdateBehaviorInput() *string
	// Experimental.
	ScheduledActionBufferTime() *float64
	// Experimental.
	SetScheduledActionBufferTime(val *float64)
	// Experimental.
	ScheduledActionBufferTimeInput() *float64
	// Experimental.
	ServiceNamespace() *string
	// Experimental.
	SetServiceNamespace(val *string)
	// Experimental.
	ServiceNamespaceInput() *string
	// Experimental.
	TargetTrackingConfiguration() AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyList
	// Experimental.
	TargetTrackingConfigurationInput() interface{}
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
	PutCustomizedLoadMetricSpecification(value *AwsAutoscalingplansScalingPlan_CustomizedLoadMetricSpecificationProperty)
	// Experimental.
	PutPredefinedLoadMetricSpecification(value *AwsAutoscalingplansScalingPlan_PredefinedLoadMetricSpecificationProperty)
	// Experimental.
	PutTargetTrackingConfiguration(value interface{})
	// Experimental.
	ResetCustomizedLoadMetricSpecification()
	// Experimental.
	ResetDisableDynamicScaling()
	// Experimental.
	ResetPredefinedLoadMetricSpecification()
	// Experimental.
	ResetPredictiveScalingMaxCapacityBehavior()
	// Experimental.
	ResetPredictiveScalingMaxCapacityBuffer()
	// Experimental.
	ResetPredictiveScalingMode()
	// Experimental.
	ResetScalingPolicyUpdateBehavior()
	// Experimental.
	ResetScheduledActionBufferTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference
type jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) CustomizedLoadMetricSpecification() AwsAutoscalingplansScalingPlan_CustomizedLoadMetricSpecificationPropertyOutputReference {
	var returns AwsAutoscalingplansScalingPlan_CustomizedLoadMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"customizedLoadMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) CustomizedLoadMetricSpecificationInput() *AwsAutoscalingplansScalingPlan_CustomizedLoadMetricSpecificationProperty {
	var returns *AwsAutoscalingplansScalingPlan_CustomizedLoadMetricSpecificationProperty
	_jsii_.Get(
		j,
		"customizedLoadMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) DisableDynamicScaling() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDynamicScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) DisableDynamicScalingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDynamicScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) MaxCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) MinCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) MinCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) PredefinedLoadMetricSpecification() AwsAutoscalingplansScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference {
	var returns AwsAutoscalingplansScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"predefinedLoadMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) PredefinedLoadMetricSpecificationInput() *AwsAutoscalingplansScalingPlan_PredefinedLoadMetricSpecificationProperty {
	var returns *AwsAutoscalingplansScalingPlan_PredefinedLoadMetricSpecificationProperty
	_jsii_.Get(
		j,
		"predefinedLoadMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) PredictiveScalingMaxCapacityBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictiveScalingMaxCapacityBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) PredictiveScalingMaxCapacityBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictiveScalingMaxCapacityBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) PredictiveScalingMaxCapacityBuffer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"predictiveScalingMaxCapacityBuffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) PredictiveScalingMaxCapacityBufferInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"predictiveScalingMaxCapacityBufferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) PredictiveScalingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictiveScalingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) PredictiveScalingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictiveScalingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ScalableDimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableDimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ScalableDimensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableDimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ScalingPolicyUpdateBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyUpdateBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ScalingPolicyUpdateBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyUpdateBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ScheduledActionBufferTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduledActionBufferTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ScheduledActionBufferTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduledActionBufferTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ServiceNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ServiceNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) TargetTrackingConfiguration() AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyList {
	var returns AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyList
	_jsii_.Get(
		j,
		"targetTrackingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) TargetTrackingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetTrackingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling-plans.AwsAutoscalingplansScalingPlan.ScalingInstructionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference_Override(a AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling-plans.AwsAutoscalingplansScalingPlan.ScalingInstructionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference)SetDisableDynamicScaling(val interface{}) {
	if err := j.validateSetDisableDynamicScalingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableDynamicScaling",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference)SetMaxCapacity(val *float64) {
	if err := j.validateSetMaxCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference)SetMinCapacity(val *float64) {
	if err := j.validateSetMinCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference)SetPredictiveScalingMaxCapacityBehavior(val *string) {
	if err := j.validateSetPredictiveScalingMaxCapacityBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictiveScalingMaxCapacityBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference)SetPredictiveScalingMaxCapacityBuffer(val *float64) {
	if err := j.validateSetPredictiveScalingMaxCapacityBufferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictiveScalingMaxCapacityBuffer",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference)SetPredictiveScalingMode(val *string) {
	if err := j.validateSetPredictiveScalingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictiveScalingMode",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference)SetScalableDimension(val *string) {
	if err := j.validateSetScalableDimensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalableDimension",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference)SetScalingPolicyUpdateBehavior(val *string) {
	if err := j.validateSetScalingPolicyUpdateBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingPolicyUpdateBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference)SetScheduledActionBufferTime(val *float64) {
	if err := j.validateSetScheduledActionBufferTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduledActionBufferTime",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference)SetServiceNamespace(val *string) {
	if err := j.validateSetServiceNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceNamespace",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) PutCustomizedLoadMetricSpecification(value *AwsAutoscalingplansScalingPlan_CustomizedLoadMetricSpecificationProperty) {
	if err := a.validatePutCustomizedLoadMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomizedLoadMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) PutPredefinedLoadMetricSpecification(value *AwsAutoscalingplansScalingPlan_PredefinedLoadMetricSpecificationProperty) {
	if err := a.validatePutPredefinedLoadMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPredefinedLoadMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) PutTargetTrackingConfiguration(value interface{}) {
	if err := a.validatePutTargetTrackingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTargetTrackingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ResetCustomizedLoadMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizedLoadMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ResetDisableDynamicScaling() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableDynamicScaling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ResetPredefinedLoadMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedLoadMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ResetPredictiveScalingMaxCapacityBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetPredictiveScalingMaxCapacityBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ResetPredictiveScalingMaxCapacityBuffer() {
	_jsii_.InvokeVoid(
		a,
		"resetPredictiveScalingMaxCapacityBuffer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ResetPredictiveScalingMode() {
	_jsii_.InvokeVoid(
		a,
		"resetPredictiveScalingMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ResetScalingPolicyUpdateBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetScalingPolicyUpdateBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ResetScheduledActionBufferTime() {
	_jsii_.InvokeVoid(
		a,
		"resetScheduledActionBufferTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_ScalingInstructionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

