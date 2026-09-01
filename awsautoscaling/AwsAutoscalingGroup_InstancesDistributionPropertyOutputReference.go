package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsautoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsautoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference interface {
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
	InternalValue() *AwsAutoscalingGroup_InstancesDistributionProperty
	// Experimental.
	SetInternalValue(val *AwsAutoscalingGroup_InstancesDistributionProperty)
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

// The jsii proxy struct for AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference
type jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) InternalValue() *AwsAutoscalingGroup_InstancesDistributionProperty {
	var returns *AwsAutoscalingGroup_InstancesDistributionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) OnDemandAllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandAllocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) OnDemandAllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandAllocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) OnDemandBaseCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandBaseCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) OnDemandBaseCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandBaseCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) OnDemandPercentageAboveBaseCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandPercentageAboveBaseCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) OnDemandPercentageAboveBaseCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandPercentageAboveBaseCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) SpotAllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotAllocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) SpotAllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotAllocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) SpotInstancePools() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotInstancePools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) SpotInstancePoolsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotInstancePoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) SpotMaxPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotMaxPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) SpotMaxPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotMaxPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAutoscalingGroup_InstancesDistributionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAutoscalingGroup_InstancesDistributionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.AwsAutoscalingGroup.InstancesDistributionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAutoscalingGroup_InstancesDistributionPropertyOutputReference_Override(a AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.AwsAutoscalingGroup.InstancesDistributionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference)SetInternalValue(val *AwsAutoscalingGroup_InstancesDistributionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference)SetOnDemandAllocationStrategy(val *string) {
	if err := j.validateSetOnDemandAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandAllocationStrategy",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference)SetOnDemandBaseCapacity(val *float64) {
	if err := j.validateSetOnDemandBaseCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandBaseCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference)SetOnDemandPercentageAboveBaseCapacity(val *float64) {
	if err := j.validateSetOnDemandPercentageAboveBaseCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandPercentageAboveBaseCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference)SetSpotAllocationStrategy(val *string) {
	if err := j.validateSetSpotAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotAllocationStrategy",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference)SetSpotInstancePools(val *float64) {
	if err := j.validateSetSpotInstancePoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotInstancePools",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference)SetSpotMaxPrice(val *string) {
	if err := j.validateSetSpotMaxPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotMaxPrice",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) ResetOnDemandAllocationStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandAllocationStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) ResetOnDemandBaseCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandBaseCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) ResetOnDemandPercentageAboveBaseCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandPercentageAboveBaseCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) ResetSpotAllocationStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotAllocationStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) ResetSpotInstancePools() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotInstancePools",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) ResetSpotMaxPrice() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotMaxPrice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAutoscalingGroup_InstancesDistributionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

