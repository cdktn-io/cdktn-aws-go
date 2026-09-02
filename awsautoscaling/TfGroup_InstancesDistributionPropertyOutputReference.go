package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsautoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsautoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGroup_InstancesDistributionPropertyOutputReference interface {
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
	InternalValue() *TfGroup_InstancesDistributionProperty
	// Experimental.
	SetInternalValue(val *TfGroup_InstancesDistributionProperty)
	// Experimental.
	OnDemandAllocationStrategy() *string
	// Experimental.
	SetOnDemandAllocationStrategy(val *string)
	// Experimental.
	OnDemandAllocationStrategyInput() *string
	// Experimental.
	OnDemandBaseCapacity() *float64
	// Experimental.
	SetOnDemandBaseCapacity(val *float64)
	// Experimental.
	OnDemandBaseCapacityInput() *float64
	// Experimental.
	OnDemandPercentageAboveBaseCapacity() *float64
	// Experimental.
	SetOnDemandPercentageAboveBaseCapacity(val *float64)
	// Experimental.
	OnDemandPercentageAboveBaseCapacityInput() *float64
	// Experimental.
	SpotAllocationStrategy() *string
	// Experimental.
	SetSpotAllocationStrategy(val *string)
	// Experimental.
	SpotAllocationStrategyInput() *string
	// Experimental.
	SpotInstancePools() *float64
	// Experimental.
	SetSpotInstancePools(val *float64)
	// Experimental.
	SpotInstancePoolsInput() *float64
	// Experimental.
	SpotMaxPrice() *string
	// Experimental.
	SetSpotMaxPrice(val *string)
	// Experimental.
	SpotMaxPriceInput() *string
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
	ResetOnDemandAllocationStrategy()
	// Experimental.
	ResetOnDemandBaseCapacity()
	// Experimental.
	ResetOnDemandPercentageAboveBaseCapacity()
	// Experimental.
	ResetSpotAllocationStrategy()
	// Experimental.
	ResetSpotInstancePools()
	// Experimental.
	ResetSpotMaxPrice()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfGroup_InstancesDistributionPropertyOutputReference
type jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) InternalValue() *TfGroup_InstancesDistributionProperty {
	var returns *TfGroup_InstancesDistributionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) OnDemandAllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandAllocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) OnDemandAllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandAllocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) OnDemandBaseCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandBaseCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) OnDemandBaseCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandBaseCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) OnDemandPercentageAboveBaseCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandPercentageAboveBaseCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) OnDemandPercentageAboveBaseCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandPercentageAboveBaseCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) SpotAllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotAllocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) SpotAllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotAllocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) SpotInstancePools() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotInstancePools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) SpotInstancePoolsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotInstancePoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) SpotMaxPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotMaxPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) SpotMaxPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotMaxPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGroup_InstancesDistributionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfGroup_InstancesDistributionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGroup_InstancesDistributionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.TfGroup.InstancesDistributionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGroup_InstancesDistributionPropertyOutputReference_Override(t TfGroup_InstancesDistributionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.TfGroup.InstancesDistributionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference)SetInternalValue(val *TfGroup_InstancesDistributionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference)SetOnDemandAllocationStrategy(val *string) {
	if err := j.validateSetOnDemandAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandAllocationStrategy",
		val,
	)
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference)SetOnDemandBaseCapacity(val *float64) {
	if err := j.validateSetOnDemandBaseCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandBaseCapacity",
		val,
	)
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference)SetOnDemandPercentageAboveBaseCapacity(val *float64) {
	if err := j.validateSetOnDemandPercentageAboveBaseCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandPercentageAboveBaseCapacity",
		val,
	)
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference)SetSpotAllocationStrategy(val *string) {
	if err := j.validateSetSpotAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotAllocationStrategy",
		val,
	)
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference)SetSpotInstancePools(val *float64) {
	if err := j.validateSetSpotInstancePoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotInstancePools",
		val,
	)
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference)SetSpotMaxPrice(val *string) {
	if err := j.validateSetSpotMaxPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotMaxPrice",
		val,
	)
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) ResetOnDemandAllocationStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetOnDemandAllocationStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) ResetOnDemandBaseCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetOnDemandBaseCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) ResetOnDemandPercentageAboveBaseCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetOnDemandPercentageAboveBaseCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) ResetSpotAllocationStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetSpotAllocationStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) ResetSpotInstancePools() {
	_jsii_.InvokeVoid(
		t,
		"resetSpotInstancePools",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) ResetSpotMaxPrice() {
	_jsii_.InvokeVoid(
		t,
		"resetSpotMaxPrice",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGroup_InstancesDistributionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

