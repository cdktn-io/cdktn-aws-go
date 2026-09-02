package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference interface {
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
	InternalValue() *TfCapacityProvider_AutoScalingGroupProviderProperty
	// Experimental.
	SetInternalValue(val *TfCapacityProvider_AutoScalingGroupProviderProperty)
	// Experimental.
	ManagedDraining() *string
	// Experimental.
	SetManagedDraining(val *string)
	// Experimental.
	ManagedDrainingInput() *string
	// Experimental.
	ManagedScaling() TfCapacityProvider_ManagedScalingPropertyOutputReference
	// Experimental.
	ManagedScalingInput() *TfCapacityProvider_ManagedScalingProperty
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
	PutManagedScaling(value *TfCapacityProvider_ManagedScalingProperty)
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

// The jsii proxy struct for TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference
type jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) AutoScalingGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) AutoScalingGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) InternalValue() *TfCapacityProvider_AutoScalingGroupProviderProperty {
	var returns *TfCapacityProvider_AutoScalingGroupProviderProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ManagedDraining() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedDraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ManagedDrainingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedDrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ManagedScaling() TfCapacityProvider_ManagedScalingPropertyOutputReference {
	var returns TfCapacityProvider_ManagedScalingPropertyOutputReference
	_jsii_.Get(
		j,
		"managedScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ManagedScalingInput() *TfCapacityProvider_ManagedScalingProperty {
	var returns *TfCapacityProvider_ManagedScalingProperty
	_jsii_.Get(
		j,
		"managedScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ManagedTerminationProtection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedTerminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ManagedTerminationProtectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedTerminationProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCapacityProvider_AutoScalingGroupProviderPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.TfCapacityProvider.AutoScalingGroupProviderPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference_Override(t TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.TfCapacityProvider.AutoScalingGroupProviderPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference)SetAutoScalingGroupArn(val *string) {
	if err := j.validateSetAutoScalingGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingGroupArn",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference)SetInternalValue(val *TfCapacityProvider_AutoScalingGroupProviderProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference)SetManagedDraining(val *string) {
	if err := j.validateSetManagedDrainingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedDraining",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference)SetManagedTerminationProtection(val *string) {
	if err := j.validateSetManagedTerminationProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedTerminationProtection",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) PutManagedScaling(value *TfCapacityProvider_ManagedScalingProperty) {
	if err := t.validatePutManagedScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putManagedScaling",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ResetManagedDraining() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedDraining",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ResetManagedScaling() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedScaling",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ResetManagedTerminationProtection() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedTerminationProtection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCapacityProvider_AutoScalingGroupProviderPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

