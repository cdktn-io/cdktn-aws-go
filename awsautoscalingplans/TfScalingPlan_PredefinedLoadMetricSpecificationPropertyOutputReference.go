package awsautoscalingplans

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsautoscalingplans/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsautoscalingplans/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference interface {
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
	InternalValue() *TfScalingPlan_PredefinedLoadMetricSpecificationProperty
	// Experimental.
	SetInternalValue(val *TfScalingPlan_PredefinedLoadMetricSpecificationProperty)
	// Experimental.
	PredefinedLoadMetricType() *string
	// Experimental.
	SetPredefinedLoadMetricType(val *string)
	// Experimental.
	PredefinedLoadMetricTypeInput() *string
	// Experimental.
	ResourceLabel() *string
	// Experimental.
	SetResourceLabel(val *string)
	// Experimental.
	ResourceLabelInput() *string
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
	ResetResourceLabel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference
type jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) InternalValue() *TfScalingPlan_PredefinedLoadMetricSpecificationProperty {
	var returns *TfScalingPlan_PredefinedLoadMetricSpecificationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) PredefinedLoadMetricType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predefinedLoadMetricType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) PredefinedLoadMetricTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"predefinedLoadMetricTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) ResourceLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) ResourceLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.PredefinedLoadMetricSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference_Override(t TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.PredefinedLoadMetricSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference)SetInternalValue(val *TfScalingPlan_PredefinedLoadMetricSpecificationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference)SetPredefinedLoadMetricType(val *string) {
	if err := j.validateSetPredefinedLoadMetricTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"predefinedLoadMetricType",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference)SetResourceLabel(val *string) {
	if err := j.validateSetResourceLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceLabel",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) ResetResourceLabel() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceLabel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

