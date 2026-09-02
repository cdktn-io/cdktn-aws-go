package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCapacityProvider_ManagedScalingPropertyOutputReference interface {
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
	InstanceWarmupPeriod() *float64
	// Experimental.
	SetInstanceWarmupPeriod(val *float64)
	// Experimental.
	InstanceWarmupPeriodInput() *float64
	// Experimental.
	InternalValue() *TfCapacityProvider_ManagedScalingProperty
	// Experimental.
	SetInternalValue(val *TfCapacityProvider_ManagedScalingProperty)
	// Experimental.
	MaximumScalingStepSize() *float64
	// Experimental.
	SetMaximumScalingStepSize(val *float64)
	// Experimental.
	MaximumScalingStepSizeInput() *float64
	// Experimental.
	MinimumScalingStepSize() *float64
	// Experimental.
	SetMinimumScalingStepSize(val *float64)
	// Experimental.
	MinimumScalingStepSizeInput() *float64
	// Experimental.
	Status() *string
	// Experimental.
	SetStatus(val *string)
	// Experimental.
	StatusInput() *string
	// Experimental.
	TargetCapacity() *float64
	// Experimental.
	SetTargetCapacity(val *float64)
	// Experimental.
	TargetCapacityInput() *float64
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
	ResetInstanceWarmupPeriod()
	// Experimental.
	ResetMaximumScalingStepSize()
	// Experimental.
	ResetMinimumScalingStepSize()
	// Experimental.
	ResetStatus()
	// Experimental.
	ResetTargetCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCapacityProvider_ManagedScalingPropertyOutputReference
type jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) InstanceWarmupPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceWarmupPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) InstanceWarmupPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceWarmupPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) InternalValue() *TfCapacityProvider_ManagedScalingProperty {
	var returns *TfCapacityProvider_ManagedScalingProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) MaximumScalingStepSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumScalingStepSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) MaximumScalingStepSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumScalingStepSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) MinimumScalingStepSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumScalingStepSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) MinimumScalingStepSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumScalingStepSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) TargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) TargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCapacityProvider_ManagedScalingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCapacityProvider_ManagedScalingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCapacityProvider_ManagedScalingPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.TfCapacityProvider.ManagedScalingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCapacityProvider_ManagedScalingPropertyOutputReference_Override(t TfCapacityProvider_ManagedScalingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.TfCapacityProvider.ManagedScalingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference)SetInstanceWarmupPeriod(val *float64) {
	if err := j.validateSetInstanceWarmupPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceWarmupPeriod",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference)SetInternalValue(val *TfCapacityProvider_ManagedScalingProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference)SetMaximumScalingStepSize(val *float64) {
	if err := j.validateSetMaximumScalingStepSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumScalingStepSize",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference)SetMinimumScalingStepSize(val *float64) {
	if err := j.validateSetMinimumScalingStepSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumScalingStepSize",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference)SetTargetCapacity(val *float64) {
	if err := j.validateSetTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCapacity",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) ResetInstanceWarmupPeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceWarmupPeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) ResetMaximumScalingStepSize() {
	_jsii_.InvokeVoid(
		t,
		"resetMaximumScalingStepSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) ResetMinimumScalingStepSize() {
	_jsii_.InvokeVoid(
		t,
		"resetMinimumScalingStepSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		t,
		"resetStatus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) ResetTargetCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedScalingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

