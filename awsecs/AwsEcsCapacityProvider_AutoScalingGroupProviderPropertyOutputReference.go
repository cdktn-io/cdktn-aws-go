package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutoScalingGroupArn() *string
	// Experimental.
	SetAutoScalingGroupArn(val *string)
	// Experimental.
	AutoScalingGroupArnInput() *string
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
	InternalValue() *AwsEcsCapacityProvider_AutoScalingGroupProviderProperty
	// Experimental.
	SetInternalValue(val *AwsEcsCapacityProvider_AutoScalingGroupProviderProperty)
	// Experimental.
	ManagedDraining() *string
	// Experimental.
	SetManagedDraining(val *string)
	// Experimental.
	ManagedDrainingInput() *string
	// Experimental.
	ManagedScaling() AwsEcsCapacityProvider_ManagedScalingPropertyOutputReference
	// Experimental.
	ManagedScalingInput() *AwsEcsCapacityProvider_ManagedScalingProperty
	// Experimental.
	ManagedTerminationProtection() *string
	// Experimental.
	SetManagedTerminationProtection(val *string)
	// Experimental.
	ManagedTerminationProtectionInput() *string
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
	PutManagedScaling(value *AwsEcsCapacityProvider_ManagedScalingProperty)
	// Experimental.
	ResetManagedDraining()
	// Experimental.
	ResetManagedScaling()
	// Experimental.
	ResetManagedTerminationProtection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference
type jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) AutoScalingGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) AutoScalingGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) InternalValue() *AwsEcsCapacityProvider_AutoScalingGroupProviderProperty {
	var returns *AwsEcsCapacityProvider_AutoScalingGroupProviderProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ManagedDraining() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedDraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ManagedDrainingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedDrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ManagedScaling() AwsEcsCapacityProvider_ManagedScalingPropertyOutputReference {
	var returns AwsEcsCapacityProvider_ManagedScalingPropertyOutputReference
	_jsii_.Get(
		j,
		"managedScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ManagedScalingInput() *AwsEcsCapacityProvider_ManagedScalingProperty {
	var returns *AwsEcsCapacityProvider_ManagedScalingProperty
	_jsii_.Get(
		j,
		"managedScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ManagedTerminationProtection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedTerminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ManagedTerminationProtectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedTerminationProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsEcsCapacityProvider.AutoScalingGroupProviderPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference_Override(a AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsEcsCapacityProvider.AutoScalingGroupProviderPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference)SetAutoScalingGroupArn(val *string) {
	if err := j.validateSetAutoScalingGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingGroupArn",
		val,
	)
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference)SetInternalValue(val *AwsEcsCapacityProvider_AutoScalingGroupProviderProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference)SetManagedDraining(val *string) {
	if err := j.validateSetManagedDrainingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedDraining",
		val,
	)
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference)SetManagedTerminationProtection(val *string) {
	if err := j.validateSetManagedTerminationProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedTerminationProtection",
		val,
	)
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) PutManagedScaling(value *AwsEcsCapacityProvider_ManagedScalingProperty) {
	if err := a.validatePutManagedScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedScaling",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ResetManagedDraining() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedDraining",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ResetManagedScaling() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedScaling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ResetManagedTerminationProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedTerminationProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEcsCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

