package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsautoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsautoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPolicy_MetricSpecificationPropertyOutputReference interface {
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
	CustomizedCapacityMetricSpecification() TfPolicy_CustomizedCapacityMetricSpecificationPropertyOutputReference
	// Experimental.
	CustomizedCapacityMetricSpecificationInput() *TfPolicy_CustomizedCapacityMetricSpecificationProperty
	// Experimental.
	CustomizedLoadMetricSpecification() TfPolicy_CustomizedLoadMetricSpecificationPropertyOutputReference
	// Experimental.
	CustomizedLoadMetricSpecificationInput() *TfPolicy_CustomizedLoadMetricSpecificationProperty
	// Experimental.
	CustomizedScalingMetricSpecification() TfPolicy_CustomizedScalingMetricSpecificationPropertyOutputReference
	// Experimental.
	CustomizedScalingMetricSpecificationInput() *TfPolicy_CustomizedScalingMetricSpecificationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfPolicy_MetricSpecificationProperty
	// Experimental.
	SetInternalValue(val *TfPolicy_MetricSpecificationProperty)
	// Experimental.
	PredefinedLoadMetricSpecification() TfPolicy_PredefinedLoadMetricSpecificationPropertyOutputReference
	// Experimental.
	PredefinedLoadMetricSpecificationInput() *TfPolicy_PredefinedLoadMetricSpecificationProperty
	// Experimental.
	PredefinedMetricPairSpecification() TfPolicy_PredefinedMetricPairSpecificationPropertyOutputReference
	// Experimental.
	PredefinedMetricPairSpecificationInput() *TfPolicy_PredefinedMetricPairSpecificationProperty
	// Experimental.
	PredefinedScalingMetricSpecification() TfPolicy_PredefinedScalingMetricSpecificationPropertyOutputReference
	// Experimental.
	PredefinedScalingMetricSpecificationInput() *TfPolicy_PredefinedScalingMetricSpecificationProperty
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
	PutCustomizedCapacityMetricSpecification(value *TfPolicy_CustomizedCapacityMetricSpecificationProperty)
	// Experimental.
	PutCustomizedLoadMetricSpecification(value *TfPolicy_CustomizedLoadMetricSpecificationProperty)
	// Experimental.
	PutCustomizedScalingMetricSpecification(value *TfPolicy_CustomizedScalingMetricSpecificationProperty)
	// Experimental.
	PutPredefinedLoadMetricSpecification(value *TfPolicy_PredefinedLoadMetricSpecificationProperty)
	// Experimental.
	PutPredefinedMetricPairSpecification(value *TfPolicy_PredefinedMetricPairSpecificationProperty)
	// Experimental.
	PutPredefinedScalingMetricSpecification(value *TfPolicy_PredefinedScalingMetricSpecificationProperty)
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

// The jsii proxy struct for TfPolicy_MetricSpecificationPropertyOutputReference
type jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) CustomizedCapacityMetricSpecification() TfPolicy_CustomizedCapacityMetricSpecificationPropertyOutputReference {
	var returns TfPolicy_CustomizedCapacityMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"customizedCapacityMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) CustomizedCapacityMetricSpecificationInput() *TfPolicy_CustomizedCapacityMetricSpecificationProperty {
	var returns *TfPolicy_CustomizedCapacityMetricSpecificationProperty
	_jsii_.Get(
		j,
		"customizedCapacityMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) CustomizedLoadMetricSpecification() TfPolicy_CustomizedLoadMetricSpecificationPropertyOutputReference {
	var returns TfPolicy_CustomizedLoadMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"customizedLoadMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) CustomizedLoadMetricSpecificationInput() *TfPolicy_CustomizedLoadMetricSpecificationProperty {
	var returns *TfPolicy_CustomizedLoadMetricSpecificationProperty
	_jsii_.Get(
		j,
		"customizedLoadMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) CustomizedScalingMetricSpecification() TfPolicy_CustomizedScalingMetricSpecificationPropertyOutputReference {
	var returns TfPolicy_CustomizedScalingMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"customizedScalingMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) CustomizedScalingMetricSpecificationInput() *TfPolicy_CustomizedScalingMetricSpecificationProperty {
	var returns *TfPolicy_CustomizedScalingMetricSpecificationProperty
	_jsii_.Get(
		j,
		"customizedScalingMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) InternalValue() *TfPolicy_MetricSpecificationProperty {
	var returns *TfPolicy_MetricSpecificationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) PredefinedLoadMetricSpecification() TfPolicy_PredefinedLoadMetricSpecificationPropertyOutputReference {
	var returns TfPolicy_PredefinedLoadMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"predefinedLoadMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) PredefinedLoadMetricSpecificationInput() *TfPolicy_PredefinedLoadMetricSpecificationProperty {
	var returns *TfPolicy_PredefinedLoadMetricSpecificationProperty
	_jsii_.Get(
		j,
		"predefinedLoadMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) PredefinedMetricPairSpecification() TfPolicy_PredefinedMetricPairSpecificationPropertyOutputReference {
	var returns TfPolicy_PredefinedMetricPairSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"predefinedMetricPairSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) PredefinedMetricPairSpecificationInput() *TfPolicy_PredefinedMetricPairSpecificationProperty {
	var returns *TfPolicy_PredefinedMetricPairSpecificationProperty
	_jsii_.Get(
		j,
		"predefinedMetricPairSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) PredefinedScalingMetricSpecification() TfPolicy_PredefinedScalingMetricSpecificationPropertyOutputReference {
	var returns TfPolicy_PredefinedScalingMetricSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"predefinedScalingMetricSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) PredefinedScalingMetricSpecificationInput() *TfPolicy_PredefinedScalingMetricSpecificationProperty {
	var returns *TfPolicy_PredefinedScalingMetricSpecificationProperty
	_jsii_.Get(
		j,
		"predefinedScalingMetricSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) TargetValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) TargetValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPolicy_MetricSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPolicy_MetricSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPolicy_MetricSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.TfPolicy.MetricSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPolicy_MetricSpecificationPropertyOutputReference_Override(t TfPolicy_MetricSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.TfPolicy.MetricSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference)SetInternalValue(val *TfPolicy_MetricSpecificationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference)SetTargetValue(val *float64) {
	if err := j.validateSetTargetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetValue",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) PutCustomizedCapacityMetricSpecification(value *TfPolicy_CustomizedCapacityMetricSpecificationProperty) {
	if err := t.validatePutCustomizedCapacityMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomizedCapacityMetricSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) PutCustomizedLoadMetricSpecification(value *TfPolicy_CustomizedLoadMetricSpecificationProperty) {
	if err := t.validatePutCustomizedLoadMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomizedLoadMetricSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) PutCustomizedScalingMetricSpecification(value *TfPolicy_CustomizedScalingMetricSpecificationProperty) {
	if err := t.validatePutCustomizedScalingMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomizedScalingMetricSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) PutPredefinedLoadMetricSpecification(value *TfPolicy_PredefinedLoadMetricSpecificationProperty) {
	if err := t.validatePutPredefinedLoadMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPredefinedLoadMetricSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) PutPredefinedMetricPairSpecification(value *TfPolicy_PredefinedMetricPairSpecificationProperty) {
	if err := t.validatePutPredefinedMetricPairSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPredefinedMetricPairSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) PutPredefinedScalingMetricSpecification(value *TfPolicy_PredefinedScalingMetricSpecificationProperty) {
	if err := t.validatePutPredefinedScalingMetricSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPredefinedScalingMetricSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) ResetCustomizedCapacityMetricSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomizedCapacityMetricSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) ResetCustomizedLoadMetricSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomizedLoadMetricSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) ResetCustomizedScalingMetricSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomizedScalingMetricSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) ResetPredefinedLoadMetricSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetPredefinedLoadMetricSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) ResetPredefinedMetricPairSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetPredefinedMetricPairSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) ResetPredefinedScalingMetricSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetPredefinedScalingMetricSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPolicy_MetricSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

