package awsapplicationautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapplicationautoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsapplicationautoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference interface {
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
	CustomizedCapacityMetricSpecification() AwsAppautoscalingPolicy_CustomizedCapacityMetricSpecificationPropertyOutputReference
	// Experimental.
	CustomizedCapacityMetricSpecificationInput() *AwsAppautoscalingPolicy_CustomizedCapacityMetricSpecificationProperty
	// Experimental.
	CustomizedLoadMetricSpecification() AwsAppautoscalingPolicy_CustomizedLoadMetricSpecificationPropertyOutputReference
	// Experimental.
	CustomizedLoadMetricSpecificationInput() *AwsAppautoscalingPolicy_CustomizedLoadMetricSpecificationProperty
	// Experimental.
	CustomizedScalingMetricSpecification() AwsAppautoscalingPolicy_CustomizedScalingMetricSpecificationPropertyOutputReference
	// Experimental.
	CustomizedScalingMetricSpecificationInput() *AwsAppautoscalingPolicy_CustomizedScalingMetricSpecificationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PredefinedLoadMetricSpecification() AwsAppautoscalingPolicy_PredefinedLoadMetricSpecificationPropertyOutputReference
	// Experimental.
	PredefinedLoadMetricSpecificationInput() *AwsAppautoscalingPolicy_PredefinedLoadMetricSpecificationProperty
	// Experimental.
	PredefinedMetricPairSpecification() AwsAppautoscalingPolicy_PredefinedMetricPairSpecificationPropertyOutputReference
	// Experimental.
	PredefinedMetricPairSpecificationInput() *AwsAppautoscalingPolicy_PredefinedMetricPairSpecificationProperty
	// Experimental.
	PredefinedScalingMetricSpecification() AwsAppautoscalingPolicy_PredefinedScalingMetricSpecificationPropertyOutputReference
	// Experimental.
	PredefinedScalingMetricSpecificationInput() *AwsAppautoscalingPolicy_PredefinedScalingMetricSpecificationProperty
	// Experimental.
	TargetValue() *string
	// Experimental.
	SetTargetValue(val *string)
	// Experimental.
	TargetValueInput() *string
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
	PutCustomizedCapacityMetricSpecification(value *AwsAppautoscalingPolicy_CustomizedCapacityMetricSpecificationProperty)
	// Experimental.
	PutCustomizedLoadMetricSpecification(value *AwsAppautoscalingPolicy_CustomizedLoadMetricSpecificationProperty)
	// Experimental.
	PutCustomizedScalingMetricSpecification(value *AwsAppautoscalingPolicy_CustomizedScalingMetricSpecificationProperty)
	// Experimental.
	PutPredefinedLoadMetricSpecification(value *AwsAppautoscalingPolicy_PredefinedLoadMetricSpecificationProperty)
	// Experimental.
	PutPredefinedMetricPairSpecification(value *AwsAppautoscalingPolicy_PredefinedMetricPairSpecificationProperty)
	// Experimental.
	PutPredefinedScalingMetricSpecification(value *AwsAppautoscalingPolicy_PredefinedScalingMetricSpecificationProperty)
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

// The jsii proxy struct for AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference
type jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) CustomizedCapacityMetricSpecification() AwsAppautoscalingPolicy_CustomizedCapacityMetricSpecificationPropertyOutputReference {
	var returns AwsAppautoscalingPolicy_CustomizedCapacityMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"customizedCapacityMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) CustomizedCapacityMetricSpecificationInput() *AwsAppautoscalingPolicy_CustomizedCapacityMetricSpecificationProperty {
	var returns *AwsAppautoscalingPolicy_CustomizedCapacityMetricSpecificationProperty
	_jsii_.Get(
		j,
		"customizedCapacityMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) CustomizedLoadMetricSpecification() AwsAppautoscalingPolicy_CustomizedLoadMetricSpecificationPropertyOutputReference {
	var returns AwsAppautoscalingPolicy_CustomizedLoadMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"customizedLoadMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) CustomizedLoadMetricSpecificationInput() *AwsAppautoscalingPolicy_CustomizedLoadMetricSpecificationProperty {
	var returns *AwsAppautoscalingPolicy_CustomizedLoadMetricSpecificationProperty
	_jsii_.Get(
		j,
		"customizedLoadMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) CustomizedScalingMetricSpecification() AwsAppautoscalingPolicy_CustomizedScalingMetricSpecificationPropertyOutputReference {
	var returns AwsAppautoscalingPolicy_CustomizedScalingMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"customizedScalingMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) CustomizedScalingMetricSpecificationInput() *AwsAppautoscalingPolicy_CustomizedScalingMetricSpecificationProperty {
	var returns *AwsAppautoscalingPolicy_CustomizedScalingMetricSpecificationProperty
	_jsii_.Get(
		j,
		"customizedScalingMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) PredefinedLoadMetricSpecification() AwsAppautoscalingPolicy_PredefinedLoadMetricSpecificationPropertyOutputReference {
	var returns AwsAppautoscalingPolicy_PredefinedLoadMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"predefinedLoadMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) PredefinedLoadMetricSpecificationInput() *AwsAppautoscalingPolicy_PredefinedLoadMetricSpecificationProperty {
	var returns *AwsAppautoscalingPolicy_PredefinedLoadMetricSpecificationProperty
	_jsii_.Get(
		j,
		"predefinedLoadMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) PredefinedMetricPairSpecification() AwsAppautoscalingPolicy_PredefinedMetricPairSpecificationPropertyOutputReference {
	var returns AwsAppautoscalingPolicy_PredefinedMetricPairSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"predefinedMetricPairSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) PredefinedMetricPairSpecificationInput() *AwsAppautoscalingPolicy_PredefinedMetricPairSpecificationProperty {
	var returns *AwsAppautoscalingPolicy_PredefinedMetricPairSpecificationProperty
	_jsii_.Get(
		j,
		"predefinedMetricPairSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) PredefinedScalingMetricSpecification() AwsAppautoscalingPolicy_PredefinedScalingMetricSpecificationPropertyOutputReference {
	var returns AwsAppautoscalingPolicy_PredefinedScalingMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"predefinedScalingMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) PredefinedScalingMetricSpecificationInput() *AwsAppautoscalingPolicy_PredefinedScalingMetricSpecificationProperty {
	var returns *AwsAppautoscalingPolicy_PredefinedScalingMetricSpecificationProperty
	_jsii_.Get(
		j,
		"predefinedScalingMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) TargetValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) TargetValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-application-auto-scaling.AwsAppautoscalingPolicy.MetricSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference_Override(a AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-application-auto-scaling.AwsAppautoscalingPolicy.MetricSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference)SetTargetValue(val *string) {
	if err := j.validateSetTargetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) PutCustomizedCapacityMetricSpecification(value *AwsAppautoscalingPolicy_CustomizedCapacityMetricSpecificationProperty) {
	if err := a.validatePutCustomizedCapacityMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomizedCapacityMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) PutCustomizedLoadMetricSpecification(value *AwsAppautoscalingPolicy_CustomizedLoadMetricSpecificationProperty) {
	if err := a.validatePutCustomizedLoadMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomizedLoadMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) PutCustomizedScalingMetricSpecification(value *AwsAppautoscalingPolicy_CustomizedScalingMetricSpecificationProperty) {
	if err := a.validatePutCustomizedScalingMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomizedScalingMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) PutPredefinedLoadMetricSpecification(value *AwsAppautoscalingPolicy_PredefinedLoadMetricSpecificationProperty) {
	if err := a.validatePutPredefinedLoadMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPredefinedLoadMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) PutPredefinedMetricPairSpecification(value *AwsAppautoscalingPolicy_PredefinedMetricPairSpecificationProperty) {
	if err := a.validatePutPredefinedMetricPairSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPredefinedMetricPairSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) PutPredefinedScalingMetricSpecification(value *AwsAppautoscalingPolicy_PredefinedScalingMetricSpecificationProperty) {
	if err := a.validatePutPredefinedScalingMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPredefinedScalingMetricSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) ResetCustomizedCapacityMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizedCapacityMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) ResetCustomizedLoadMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizedLoadMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) ResetCustomizedScalingMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizedScalingMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) ResetPredefinedLoadMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedLoadMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) ResetPredefinedMetricPairSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedMetricPairSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) ResetPredefinedScalingMetricSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetPredefinedScalingMetricSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppautoscalingPolicy_MetricSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

