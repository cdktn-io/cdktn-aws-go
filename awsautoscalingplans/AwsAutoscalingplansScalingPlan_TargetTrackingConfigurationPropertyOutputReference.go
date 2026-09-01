package awsautoscalingplans

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsautoscalingplans/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsautoscalingplans/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference interface {
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
	CustomizedScalingMetricSpecification() AwsAutoscalingplansScalingPlan_CustomizedScalingMetricSpecificationPropertyOutputReference
	// Experimental.
	CustomizedScalingMetricSpecificationInput() *AwsAutoscalingplansScalingPlan_CustomizedScalingMetricSpecificationProperty
	// Experimental.
	DisableScaleIn() interface{}
	// Experimental.
	SetDisableScaleIn(val interface{})
	// Experimental.
	DisableScaleInInput() interface{}
	// Experimental.
	EstimatedInstanceWarmup() *float64
	// Experimental.
	SetEstimatedInstanceWarmup(val *float64)
	// Experimental.
	EstimatedInstanceWarmupInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PredefinedScalingMetricSpecification() AwsAutoscalingplansScalingPlan_PredefinedScalingMetricSpecificationPropertyOutputReference
	// Experimental.
	PredefinedScalingMetricSpecificationInput() *AwsAutoscalingplansScalingPlan_PredefinedScalingMetricSpecificationProperty
	// Experimental.
	ScaleInCooldown() *float64
	// Experimental.
	SetScaleInCooldown(val *float64)
	// Experimental.
	ScaleInCooldownInput() *float64
	// Experimental.
	ScaleOutCooldown() *float64
	// Experimental.
	SetScaleOutCooldown(val *float64)
	// Experimental.
	ScaleOutCooldownInput() *float64
	// Experimental.
	TargetValue() *float64
	// Experimental.
	SetTargetValue(val *float64)
	// Experimental.
	TargetValueInput() *float64
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
	PutCustomizedScalingMetricSpecification(value *AwsAutoscalingplansScalingPlan_CustomizedScalingMetricSpecificationProperty)
	// Experimental.
	PutPredefinedScalingMetricSpecification(value *AwsAutoscalingplansScalingPlan_PredefinedScalingMetricSpecificationProperty)
	// Experimental.
	ResetCustomizedScalingMetricSpecification()
	// Experimental.
	ResetDisableScaleIn()
	// Experimental.
	ResetEstimatedInstanceWarmup()
	// Experimental.
	ResetPredefinedScalingMetricSpecification()
	// Experimental.
	ResetScaleInCooldown()
	// Experimental.
	ResetScaleOutCooldown()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference
type jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) CustomizedScalingMetricSpecification() AwsAutoscalingplansScalingPlan_CustomizedScalingMetricSpecificationPropertyOutputReference {
	var returns AwsAutoscalingplansScalingPlan_CustomizedScalingMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"customizedScalingMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) CustomizedScalingMetricSpecificationInput() *AwsAutoscalingplansScalingPlan_CustomizedScalingMetricSpecificationProperty {
	var returns *AwsAutoscalingplansScalingPlan_CustomizedScalingMetricSpecificationProperty
	_jsii_.Get(
		j,
		"customizedScalingMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) DisableScaleIn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) DisableScaleInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableScaleInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) EstimatedInstanceWarmup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"estimatedInstanceWarmup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) EstimatedInstanceWarmupInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"estimatedInstanceWarmupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) PredefinedScalingMetricSpecification() AwsAutoscalingplansScalingPlan_PredefinedScalingMetricSpecificationPropertyOutputReference {
	var returns AwsAutoscalingplansScalingPlan_PredefinedScalingMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"predefinedScalingMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) PredefinedScalingMetricSpecificationInput() *AwsAutoscalingplansScalingPlan_PredefinedScalingMetricSpecificationProperty {
	var returns *AwsAutoscalingplansScalingPlan_PredefinedScalingMetricSpecificationProperty
	_jsii_.Get(
		j,
		"predefinedScalingMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) ScaleInCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) ScaleInCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) ScaleOutCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) ScaleOutCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) TargetValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) TargetValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling-plans.AwsAutoscalingplansScalingPlan.TargetTrackingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference_Override(a AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling-plans.AwsAutoscalingplansScalingPlan.TargetTrackingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference)SetDisableScaleIn(val interface{}) {
	if err := j.validateSetDisableScaleInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableScaleIn",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference)SetEstimatedInstanceWarmup(val *float64) {
	if err := j.validateSetEstimatedInstanceWarmupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"estimatedInstanceWarmup",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference)SetScaleInCooldown(val *float64) {
	if err := j.validateSetScaleInCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleInCooldown",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference)SetScaleOutCooldown(val *float64) {
	if err := j.validateSetScaleOutCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleOutCooldown",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference)SetTargetValue(val *float64) {
	if err := j.validateSetTargetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetValue",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) PutCustomizedScalingMetricSpecification(value *AwsAutoscalingplansScalingPlan_CustomizedScalingMetricSpecificationProperty) {
	if err := a.validatePutCustomizedScalingMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomizedScalingMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) PutPredefinedScalingMetricSpecification(value *AwsAutoscalingplansScalingPlan_PredefinedScalingMetricSpecificationProperty) {
	if err := a.validatePutPredefinedScalingMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPredefinedScalingMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) ResetCustomizedScalingMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizedScalingMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) ResetDisableScaleIn() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableScaleIn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) ResetEstimatedInstanceWarmup() {
	_jsii_.InvokeVoid(
		a,
		"resetEstimatedInstanceWarmup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) ResetPredefinedScalingMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedScalingMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) ResetScaleInCooldown() {
	_jsii_.InvokeVoid(
		a,
		"resetScaleInCooldown",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) ResetScaleOutCooldown() {
	_jsii_.InvokeVoid(
		a,
		"resetScaleOutCooldown",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAutoscalingplansScalingPlan_TargetTrackingConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

