package awsapplicationautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapplicationautoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsapplicationautoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference interface {
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
	CustomizedMetricSpecification() TfPolicy_CustomizedMetricSpecificationPropertyOutputReference
	// Experimental.
	CustomizedMetricSpecificationInput() *TfPolicy_CustomizedMetricSpecificationProperty
	// Experimental.
	DisableScaleIn() interface{}
	// Experimental.
	SetDisableScaleIn(val interface{})
	// Experimental.
	DisableScaleInInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfPolicy_TargetTrackingScalingPolicyConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfPolicy_TargetTrackingScalingPolicyConfigurationProperty)
	// Experimental.
	PredefinedMetricSpecification() TfPolicy_PredefinedMetricSpecificationPropertyOutputReference
	// Experimental.
	PredefinedMetricSpecificationInput() *TfPolicy_PredefinedMetricSpecificationProperty
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
	PutCustomizedMetricSpecification(value *TfPolicy_CustomizedMetricSpecificationProperty)
	// Experimental.
	PutPredefinedMetricSpecification(value *TfPolicy_PredefinedMetricSpecificationProperty)
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

// The jsii proxy struct for TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference
type jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) CustomizedMetricSpecification() TfPolicy_CustomizedMetricSpecificationPropertyOutputReference {
	var returns TfPolicy_CustomizedMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"customizedMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) CustomizedMetricSpecificationInput() *TfPolicy_CustomizedMetricSpecificationProperty {
	var returns *TfPolicy_CustomizedMetricSpecificationProperty
	_jsii_.Get(
		j,
		"customizedMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) DisableScaleIn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) DisableScaleInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableScaleInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) InternalValue() *TfPolicy_TargetTrackingScalingPolicyConfigurationProperty {
	var returns *TfPolicy_TargetTrackingScalingPolicyConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) PredefinedMetricSpecification() TfPolicy_PredefinedMetricSpecificationPropertyOutputReference {
	var returns TfPolicy_PredefinedMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"predefinedMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) PredefinedMetricSpecificationInput() *TfPolicy_PredefinedMetricSpecificationProperty {
	var returns *TfPolicy_PredefinedMetricSpecificationProperty
	_jsii_.Get(
		j,
		"predefinedMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ScaleInCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ScaleInCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ScaleOutCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ScaleOutCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) TargetValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) TargetValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-application-auto-scaling.TfPolicy.TargetTrackingScalingPolicyConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference_Override(t TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-application-auto-scaling.TfPolicy.TargetTrackingScalingPolicyConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetDisableScaleIn(val interface{}) {
	if err := j.validateSetDisableScaleInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableScaleIn",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetInternalValue(val *TfPolicy_TargetTrackingScalingPolicyConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetScaleInCooldown(val *float64) {
	if err := j.validateSetScaleInCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleInCooldown",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetScaleOutCooldown(val *float64) {
	if err := j.validateSetScaleOutCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleOutCooldown",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetTargetValue(val *float64) {
	if err := j.validateSetTargetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetValue",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) PutCustomizedMetricSpecification(value *TfPolicy_CustomizedMetricSpecificationProperty) {
	if err := t.validatePutCustomizedMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomizedMetricSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) PutPredefinedMetricSpecification(value *TfPolicy_PredefinedMetricSpecificationProperty) {
	if err := t.validatePutPredefinedMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPredefinedMetricSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ResetCustomizedMetricSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomizedMetricSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ResetDisableScaleIn() {
	_jsii_.InvokeVoid(
		t,
		"resetDisableScaleIn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ResetPredefinedMetricSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetPredefinedMetricSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ResetScaleInCooldown() {
	_jsii_.InvokeVoid(
		t,
		"resetScaleInCooldown",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ResetScaleOutCooldown() {
	_jsii_.InvokeVoid(
		t,
		"resetScaleOutCooldown",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPolicy_TargetTrackingScalingPolicyConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

