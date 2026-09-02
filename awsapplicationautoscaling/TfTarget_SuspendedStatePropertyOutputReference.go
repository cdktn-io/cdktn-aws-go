package awsapplicationautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapplicationautoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsapplicationautoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTarget_SuspendedStatePropertyOutputReference interface {
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
	DynamicScalingInSuspended() interface{}
	// Experimental.
	SetDynamicScalingInSuspended(val interface{})
	// Experimental.
	DynamicScalingInSuspendedInput() interface{}
	// Experimental.
	DynamicScalingOutSuspended() interface{}
	// Experimental.
	SetDynamicScalingOutSuspended(val interface{})
	// Experimental.
	DynamicScalingOutSuspendedInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfTarget_SuspendedStateProperty
	// Experimental.
	SetInternalValue(val *TfTarget_SuspendedStateProperty)
	// Experimental.
	ScheduledScalingSuspended() interface{}
	// Experimental.
	SetScheduledScalingSuspended(val interface{})
	// Experimental.
	ScheduledScalingSuspendedInput() interface{}
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
	ResetDynamicScalingInSuspended()
	// Experimental.
	ResetDynamicScalingOutSuspended()
	// Experimental.
	ResetScheduledScalingSuspended()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTarget_SuspendedStatePropertyOutputReference
type jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) DynamicScalingInSuspended() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicScalingInSuspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) DynamicScalingInSuspendedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicScalingInSuspendedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) DynamicScalingOutSuspended() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicScalingOutSuspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) DynamicScalingOutSuspendedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicScalingOutSuspendedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) InternalValue() *TfTarget_SuspendedStateProperty {
	var returns *TfTarget_SuspendedStateProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) ScheduledScalingSuspended() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledScalingSuspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) ScheduledScalingSuspendedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledScalingSuspendedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTarget_SuspendedStatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTarget_SuspendedStatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTarget_SuspendedStatePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-application-auto-scaling.TfTarget.SuspendedStatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTarget_SuspendedStatePropertyOutputReference_Override(t TfTarget_SuspendedStatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-application-auto-scaling.TfTarget.SuspendedStatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference)SetDynamicScalingInSuspended(val interface{}) {
	if err := j.validateSetDynamicScalingInSuspendedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicScalingInSuspended",
		val,
	)
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference)SetDynamicScalingOutSuspended(val interface{}) {
	if err := j.validateSetDynamicScalingOutSuspendedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicScalingOutSuspended",
		val,
	)
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference)SetInternalValue(val *TfTarget_SuspendedStateProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference)SetScheduledScalingSuspended(val interface{}) {
	if err := j.validateSetScheduledScalingSuspendedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduledScalingSuspended",
		val,
	)
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) ResetDynamicScalingInSuspended() {
	_jsii_.InvokeVoid(
		t,
		"resetDynamicScalingInSuspended",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) ResetDynamicScalingOutSuspended() {
	_jsii_.InvokeVoid(
		t,
		"resetDynamicScalingOutSuspended",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) ResetScheduledScalingSuspended() {
	_jsii_.InvokeVoid(
		t,
		"resetScheduledScalingSuspended",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTarget_SuspendedStatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

