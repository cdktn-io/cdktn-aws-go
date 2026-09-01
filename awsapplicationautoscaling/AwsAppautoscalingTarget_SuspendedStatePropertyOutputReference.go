package awsapplicationautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapplicationautoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsapplicationautoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference interface {
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
	InternalValue() *AwsAppautoscalingTarget_SuspendedStateProperty
	// Experimental.
	SetInternalValue(val *AwsAppautoscalingTarget_SuspendedStateProperty)
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

// The jsii proxy struct for AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference
type jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) DynamicScalingInSuspended() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicScalingInSuspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) DynamicScalingInSuspendedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicScalingInSuspendedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) DynamicScalingOutSuspended() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicScalingOutSuspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) DynamicScalingOutSuspendedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicScalingOutSuspendedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) InternalValue() *AwsAppautoscalingTarget_SuspendedStateProperty {
	var returns *AwsAppautoscalingTarget_SuspendedStateProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) ScheduledScalingSuspended() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledScalingSuspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) ScheduledScalingSuspendedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledScalingSuspendedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppautoscalingTarget_SuspendedStatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppautoscalingTarget_SuspendedStatePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-application-auto-scaling.AwsAppautoscalingTarget.SuspendedStatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppautoscalingTarget_SuspendedStatePropertyOutputReference_Override(a AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-application-auto-scaling.AwsAppautoscalingTarget.SuspendedStatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference)SetDynamicScalingInSuspended(val interface{}) {
	if err := j.validateSetDynamicScalingInSuspendedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicScalingInSuspended",
		val,
	)
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference)SetDynamicScalingOutSuspended(val interface{}) {
	if err := j.validateSetDynamicScalingOutSuspendedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamicScalingOutSuspended",
		val,
	)
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference)SetInternalValue(val *AwsAppautoscalingTarget_SuspendedStateProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference)SetScheduledScalingSuspended(val interface{}) {
	if err := j.validateSetScheduledScalingSuspendedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduledScalingSuspended",
		val,
	)
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) ResetDynamicScalingInSuspended() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamicScalingInSuspended",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) ResetDynamicScalingOutSuspended() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamicScalingOutSuspended",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) ResetScheduledScalingSuspended() {
	_jsii_.InvokeVoid(
		a,
		"resetScheduledScalingSuspended",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppautoscalingTarget_SuspendedStatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

