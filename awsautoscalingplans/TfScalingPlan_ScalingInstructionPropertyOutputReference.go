package awsautoscalingplans

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsautoscalingplans/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsautoscalingplans/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfScalingPlan_ScalingInstructionPropertyOutputReference interface {
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
	CustomizedLoadMetricSpecification() TfScalingPlan_CustomizedLoadMetricSpecificationPropertyOutputReference
	// Experimental.
	CustomizedLoadMetricSpecificationInput() *TfScalingPlan_CustomizedLoadMetricSpecificationProperty
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
	PredefinedLoadMetricSpecification() TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference
	// Experimental.
	PredefinedLoadMetricSpecificationInput() *TfScalingPlan_PredefinedLoadMetricSpecificationProperty
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
	TargetTrackingConfiguration() TfScalingPlan_TargetTrackingConfigurationPropertyList
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
	PutCustomizedLoadMetricSpecification(value *TfScalingPlan_CustomizedLoadMetricSpecificationProperty)
	// Experimental.
	PutPredefinedLoadMetricSpecification(value *TfScalingPlan_PredefinedLoadMetricSpecificationProperty)
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

// The jsii proxy struct for TfScalingPlan_ScalingInstructionPropertyOutputReference
type jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) CustomizedLoadMetricSpecification() TfScalingPlan_CustomizedLoadMetricSpecificationPropertyOutputReference {
	var returns TfScalingPlan_CustomizedLoadMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"customizedLoadMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) CustomizedLoadMetricSpecificationInput() *TfScalingPlan_CustomizedLoadMetricSpecificationProperty {
	var returns *TfScalingPlan_CustomizedLoadMetricSpecificationProperty
	_jsii_.Get(
		j,
		"customizedLoadMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) DisableDynamicScaling() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDynamicScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) DisableDynamicScalingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDynamicScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) MaxCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) MinCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) MinCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) PredefinedLoadMetricSpecification() TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference {
	var returns TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"predefinedLoadMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) PredefinedLoadMetricSpecificationInput() *TfScalingPlan_PredefinedLoadMetricSpecificationProperty {
	var returns *TfScalingPlan_PredefinedLoadMetricSpecificationProperty
	_jsii_.Get(
		j,
		"predefinedLoadMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) PredictiveScalingMaxCapacityBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictiveScalingMaxCapacityBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) PredictiveScalingMaxCapacityBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictiveScalingMaxCapacityBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) PredictiveScalingMaxCapacityBuffer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"predictiveScalingMaxCapacityBuffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) PredictiveScalingMaxCapacityBufferInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"predictiveScalingMaxCapacityBufferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) PredictiveScalingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictiveScalingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) PredictiveScalingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predictiveScalingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ScalableDimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableDimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ScalableDimensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalableDimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ScalingPolicyUpdateBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyUpdateBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ScalingPolicyUpdateBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scalingPolicyUpdateBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ScheduledActionBufferTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduledActionBufferTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ScheduledActionBufferTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scheduledActionBufferTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ServiceNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ServiceNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) TargetTrackingConfiguration() TfScalingPlan_TargetTrackingConfigurationPropertyList {
	var returns TfScalingPlan_TargetTrackingConfigurationPropertyList
	_jsii_.Get(
		j,
		"targetTrackingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) TargetTrackingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetTrackingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfScalingPlan_ScalingInstructionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfScalingPlan_ScalingInstructionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfScalingPlan_ScalingInstructionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.ScalingInstructionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfScalingPlan_ScalingInstructionPropertyOutputReference_Override(t TfScalingPlan_ScalingInstructionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.ScalingInstructionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference)SetDisableDynamicScaling(val interface{}) {
	if err := j.validateSetDisableDynamicScalingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableDynamicScaling",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference)SetMaxCapacity(val *float64) {
	if err := j.validateSetMaxCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference)SetMinCapacity(val *float64) {
	if err := j.validateSetMinCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCapacity",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference)SetPredictiveScalingMaxCapacityBehavior(val *string) {
	if err := j.validateSetPredictiveScalingMaxCapacityBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictiveScalingMaxCapacityBehavior",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference)SetPredictiveScalingMaxCapacityBuffer(val *float64) {
	if err := j.validateSetPredictiveScalingMaxCapacityBufferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictiveScalingMaxCapacityBuffer",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference)SetPredictiveScalingMode(val *string) {
	if err := j.validateSetPredictiveScalingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predictiveScalingMode",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference)SetScalableDimension(val *string) {
	if err := j.validateSetScalableDimensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalableDimension",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference)SetScalingPolicyUpdateBehavior(val *string) {
	if err := j.validateSetScalingPolicyUpdateBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scalingPolicyUpdateBehavior",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference)SetScheduledActionBufferTime(val *float64) {
	if err := j.validateSetScheduledActionBufferTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduledActionBufferTime",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference)SetServiceNamespace(val *string) {
	if err := j.validateSetServiceNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceNamespace",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) PutCustomizedLoadMetricSpecification(value *TfScalingPlan_CustomizedLoadMetricSpecificationProperty) {
	if err := t.validatePutCustomizedLoadMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomizedLoadMetricSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) PutPredefinedLoadMetricSpecification(value *TfScalingPlan_PredefinedLoadMetricSpecificationProperty) {
	if err := t.validatePutPredefinedLoadMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPredefinedLoadMetricSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) PutTargetTrackingConfiguration(value interface{}) {
	if err := t.validatePutTargetTrackingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTargetTrackingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ResetCustomizedLoadMetricSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomizedLoadMetricSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ResetDisableDynamicScaling() {
	_jsii_.InvokeVoid(
		t,
		"resetDisableDynamicScaling",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ResetPredefinedLoadMetricSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetPredefinedLoadMetricSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ResetPredictiveScalingMaxCapacityBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetPredictiveScalingMaxCapacityBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ResetPredictiveScalingMaxCapacityBuffer() {
	_jsii_.InvokeVoid(
		t,
		"resetPredictiveScalingMaxCapacityBuffer",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ResetPredictiveScalingMode() {
	_jsii_.InvokeVoid(
		t,
		"resetPredictiveScalingMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ResetScalingPolicyUpdateBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetScalingPolicyUpdateBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ResetScheduledActionBufferTime() {
	_jsii_.InvokeVoid(
		t,
		"resetScheduledActionBufferTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

