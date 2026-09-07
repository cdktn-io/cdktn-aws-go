package applicationautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/applicationautoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/applicationautoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference interface {
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
	CustomizedMetricSpecification() AwsPolicy_CustomizedMetricSpecificationPropertyOutputReference
	// Experimental.
	CustomizedMetricSpecificationInput() *AwsPolicy_CustomizedMetricSpecificationProperty
	// Experimental.
	DisableScaleIn() interface{}
	// Experimental.
	SetDisableScaleIn(val interface{})
	// Experimental.
	DisableScaleInInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsPolicy_TargetTrackingScalingPolicyConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsPolicy_TargetTrackingScalingPolicyConfigurationProperty)
	// Experimental.
	PredefinedMetricSpecification() AwsPolicy_PredefinedMetricSpecificationPropertyOutputReference
	// Experimental.
	PredefinedMetricSpecificationInput() *AwsPolicy_PredefinedMetricSpecificationProperty
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
	PutCustomizedMetricSpecification(value *AwsPolicy_CustomizedMetricSpecificationProperty)
	// Experimental.
	PutPredefinedMetricSpecification(value *AwsPolicy_PredefinedMetricSpecificationProperty)
	// Experimental.
	ResetCustomizedMetricSpecification()
	// Experimental.
	ResetDisableScaleIn()
	// Experimental.
	ResetPredefinedMetricSpecification()
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

// The jsii proxy struct for AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference
type jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) CustomizedMetricSpecification() AwsPolicy_CustomizedMetricSpecificationPropertyOutputReference {
	var returns AwsPolicy_CustomizedMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"customizedMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) CustomizedMetricSpecificationInput() *AwsPolicy_CustomizedMetricSpecificationProperty {
	var returns *AwsPolicy_CustomizedMetricSpecificationProperty
	_jsii_.Get(
		j,
		"customizedMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) DisableScaleIn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) DisableScaleInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableScaleInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) InternalValue() *AwsPolicy_TargetTrackingScalingPolicyConfigurationProperty {
	var returns *AwsPolicy_TargetTrackingScalingPolicyConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) PredefinedMetricSpecification() AwsPolicy_PredefinedMetricSpecificationPropertyOutputReference {
	var returns AwsPolicy_PredefinedMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"predefinedMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) PredefinedMetricSpecificationInput() *AwsPolicy_PredefinedMetricSpecificationProperty {
	var returns *AwsPolicy_PredefinedMetricSpecificationProperty
	_jsii_.Get(
		j,
		"predefinedMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ScaleInCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ScaleInCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ScaleOutCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ScaleOutCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) TargetValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) TargetValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-application-auto-scaling.AwsPolicy.TargetTrackingScalingPolicyConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference_Override(a AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-application-auto-scaling.AwsPolicy.TargetTrackingScalingPolicyConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetDisableScaleIn(val interface{}) {
	if err := j.validateSetDisableScaleInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableScaleIn",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetInternalValue(val *AwsPolicy_TargetTrackingScalingPolicyConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetScaleInCooldown(val *float64) {
	if err := j.validateSetScaleInCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleInCooldown",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetScaleOutCooldown(val *float64) {
	if err := j.validateSetScaleOutCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleOutCooldown",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetTargetValue(val *float64) {
	if err := j.validateSetTargetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetValue",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) PutCustomizedMetricSpecification(value *AwsPolicy_CustomizedMetricSpecificationProperty) {
	if err := a.validatePutCustomizedMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomizedMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) PutPredefinedMetricSpecification(value *AwsPolicy_PredefinedMetricSpecificationProperty) {
	if err := a.validatePutPredefinedMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPredefinedMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ResetCustomizedMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizedMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ResetDisableScaleIn() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableScaleIn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ResetPredefinedMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ResetScaleInCooldown() {
	_jsii_.InvokeVoid(
		a,
		"resetScaleInCooldown",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ResetScaleOutCooldown() {
	_jsii_.InvokeVoid(
		a,
		"resetScaleOutCooldown",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

