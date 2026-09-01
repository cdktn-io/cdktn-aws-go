package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsautoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsautoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference interface {
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
	CustomizedCapacityMetricSpecification() AwsAutoscalingPolicy_CustomizedCapacityMetricSpecificationPropertyOutputReference
	// Experimental.
	CustomizedCapacityMetricSpecificationInput() *AwsAutoscalingPolicy_CustomizedCapacityMetricSpecificationProperty
	// Experimental.
	CustomizedLoadMetricSpecification() AwsAutoscalingPolicy_CustomizedLoadMetricSpecificationPropertyOutputReference
	// Experimental.
	CustomizedLoadMetricSpecificationInput() *AwsAutoscalingPolicy_CustomizedLoadMetricSpecificationProperty
	// Experimental.
	CustomizedScalingMetricSpecification() AwsAutoscalingPolicy_CustomizedScalingMetricSpecificationPropertyOutputReference
	// Experimental.
	CustomizedScalingMetricSpecificationInput() *AwsAutoscalingPolicy_CustomizedScalingMetricSpecificationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsAutoscalingPolicy_MetricSpecificationProperty
	// Experimental.
	SetInternalValue(val *AwsAutoscalingPolicy_MetricSpecificationProperty)
	// Experimental.
	PredefinedLoadMetricSpecification() AwsAutoscalingPolicy_PredefinedLoadMetricSpecificationPropertyOutputReference
	// Experimental.
	PredefinedLoadMetricSpecificationInput() *AwsAutoscalingPolicy_PredefinedLoadMetricSpecificationProperty
	// Experimental.
	PredefinedMetricPairSpecification() AwsAutoscalingPolicy_PredefinedMetricPairSpecificationPropertyOutputReference
	// Experimental.
	PredefinedMetricPairSpecificationInput() *AwsAutoscalingPolicy_PredefinedMetricPairSpecificationProperty
	// Experimental.
	PredefinedScalingMetricSpecification() AwsAutoscalingPolicy_PredefinedScalingMetricSpecificationPropertyOutputReference
	// Experimental.
	PredefinedScalingMetricSpecificationInput() *AwsAutoscalingPolicy_PredefinedScalingMetricSpecificationProperty
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
	PutCustomizedCapacityMetricSpecification(value *AwsAutoscalingPolicy_CustomizedCapacityMetricSpecificationProperty)
	// Experimental.
	PutCustomizedLoadMetricSpecification(value *AwsAutoscalingPolicy_CustomizedLoadMetricSpecificationProperty)
	// Experimental.
	PutCustomizedScalingMetricSpecification(value *AwsAutoscalingPolicy_CustomizedScalingMetricSpecificationProperty)
	// Experimental.
	PutPredefinedLoadMetricSpecification(value *AwsAutoscalingPolicy_PredefinedLoadMetricSpecificationProperty)
	// Experimental.
	PutPredefinedMetricPairSpecification(value *AwsAutoscalingPolicy_PredefinedMetricPairSpecificationProperty)
	// Experimental.
	PutPredefinedScalingMetricSpecification(value *AwsAutoscalingPolicy_PredefinedScalingMetricSpecificationProperty)
	// Experimental.
	ResetCustomizedCapacityMetricSpecification()
	// Experimental.
	ResetCustomizedLoadMetricSpecification()
	// Experimental.
	ResetCustomizedScalingMetricSpecification()
	// Experimental.
	ResetPredefinedLoadMetricSpecification()
	// Experimental.
	ResetPredefinedMetricPairSpecification()
	// Experimental.
	ResetPredefinedScalingMetricSpecification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference
type jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) CustomizedCapacityMetricSpecification() AwsAutoscalingPolicy_CustomizedCapacityMetricSpecificationPropertyOutputReference {
	var returns AwsAutoscalingPolicy_CustomizedCapacityMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"customizedCapacityMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) CustomizedCapacityMetricSpecificationInput() *AwsAutoscalingPolicy_CustomizedCapacityMetricSpecificationProperty {
	var returns *AwsAutoscalingPolicy_CustomizedCapacityMetricSpecificationProperty
	_jsii_.Get(
		j,
		"customizedCapacityMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) CustomizedLoadMetricSpecification() AwsAutoscalingPolicy_CustomizedLoadMetricSpecificationPropertyOutputReference {
	var returns AwsAutoscalingPolicy_CustomizedLoadMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"customizedLoadMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) CustomizedLoadMetricSpecificationInput() *AwsAutoscalingPolicy_CustomizedLoadMetricSpecificationProperty {
	var returns *AwsAutoscalingPolicy_CustomizedLoadMetricSpecificationProperty
	_jsii_.Get(
		j,
		"customizedLoadMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) CustomizedScalingMetricSpecification() AwsAutoscalingPolicy_CustomizedScalingMetricSpecificationPropertyOutputReference {
	var returns AwsAutoscalingPolicy_CustomizedScalingMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"customizedScalingMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) CustomizedScalingMetricSpecificationInput() *AwsAutoscalingPolicy_CustomizedScalingMetricSpecificationProperty {
	var returns *AwsAutoscalingPolicy_CustomizedScalingMetricSpecificationProperty
	_jsii_.Get(
		j,
		"customizedScalingMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) InternalValue() *AwsAutoscalingPolicy_MetricSpecificationProperty {
	var returns *AwsAutoscalingPolicy_MetricSpecificationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) PredefinedLoadMetricSpecification() AwsAutoscalingPolicy_PredefinedLoadMetricSpecificationPropertyOutputReference {
	var returns AwsAutoscalingPolicy_PredefinedLoadMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"predefinedLoadMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) PredefinedLoadMetricSpecificationInput() *AwsAutoscalingPolicy_PredefinedLoadMetricSpecificationProperty {
	var returns *AwsAutoscalingPolicy_PredefinedLoadMetricSpecificationProperty
	_jsii_.Get(
		j,
		"predefinedLoadMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) PredefinedMetricPairSpecification() AwsAutoscalingPolicy_PredefinedMetricPairSpecificationPropertyOutputReference {
	var returns AwsAutoscalingPolicy_PredefinedMetricPairSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"predefinedMetricPairSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) PredefinedMetricPairSpecificationInput() *AwsAutoscalingPolicy_PredefinedMetricPairSpecificationProperty {
	var returns *AwsAutoscalingPolicy_PredefinedMetricPairSpecificationProperty
	_jsii_.Get(
		j,
		"predefinedMetricPairSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) PredefinedScalingMetricSpecification() AwsAutoscalingPolicy_PredefinedScalingMetricSpecificationPropertyOutputReference {
	var returns AwsAutoscalingPolicy_PredefinedScalingMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"predefinedScalingMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) PredefinedScalingMetricSpecificationInput() *AwsAutoscalingPolicy_PredefinedScalingMetricSpecificationProperty {
	var returns *AwsAutoscalingPolicy_PredefinedScalingMetricSpecificationProperty
	_jsii_.Get(
		j,
		"predefinedScalingMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) TargetValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) TargetValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAutoscalingPolicy_MetricSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.AwsAutoscalingPolicy.MetricSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference_Override(a AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.AwsAutoscalingPolicy.MetricSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference)SetInternalValue(val *AwsAutoscalingPolicy_MetricSpecificationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference)SetTargetValue(val *float64) {
	if err := j.validateSetTargetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetValue",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) PutCustomizedCapacityMetricSpecification(value *AwsAutoscalingPolicy_CustomizedCapacityMetricSpecificationProperty) {
	if err := a.validatePutCustomizedCapacityMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomizedCapacityMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) PutCustomizedLoadMetricSpecification(value *AwsAutoscalingPolicy_CustomizedLoadMetricSpecificationProperty) {
	if err := a.validatePutCustomizedLoadMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomizedLoadMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) PutCustomizedScalingMetricSpecification(value *AwsAutoscalingPolicy_CustomizedScalingMetricSpecificationProperty) {
	if err := a.validatePutCustomizedScalingMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomizedScalingMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) PutPredefinedLoadMetricSpecification(value *AwsAutoscalingPolicy_PredefinedLoadMetricSpecificationProperty) {
	if err := a.validatePutPredefinedLoadMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPredefinedLoadMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) PutPredefinedMetricPairSpecification(value *AwsAutoscalingPolicy_PredefinedMetricPairSpecificationProperty) {
	if err := a.validatePutPredefinedMetricPairSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPredefinedMetricPairSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) PutPredefinedScalingMetricSpecification(value *AwsAutoscalingPolicy_PredefinedScalingMetricSpecificationProperty) {
	if err := a.validatePutPredefinedScalingMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPredefinedScalingMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) ResetCustomizedCapacityMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizedCapacityMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) ResetCustomizedLoadMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizedLoadMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) ResetCustomizedScalingMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizedScalingMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) ResetPredefinedLoadMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedLoadMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) ResetPredefinedMetricPairSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedMetricPairSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) ResetPredefinedScalingMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedScalingMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAutoscalingPolicy_MetricSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

