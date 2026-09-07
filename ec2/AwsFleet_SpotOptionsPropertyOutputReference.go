package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ec2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/ec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFleet_SpotOptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllocationStrategy() *string
	// Experimental.
	SetAllocationStrategy(val *string)
	// Experimental.
	AllocationStrategyInput() *string
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
	InstanceInterruptionBehavior() *string
	// Experimental.
	SetInstanceInterruptionBehavior(val *string)
	// Experimental.
	InstanceInterruptionBehaviorInput() *string
	// Experimental.
	InstancePoolsToUseCount() *float64
	// Experimental.
	SetInstancePoolsToUseCount(val *float64)
	// Experimental.
	InstancePoolsToUseCountInput() *float64
	// Experimental.
	InternalValue() *AwsFleet_SpotOptionsProperty
	// Experimental.
	SetInternalValue(val *AwsFleet_SpotOptionsProperty)
	// Experimental.
	MaintenanceStrategies() AwsFleet_MaintenanceStrategiesPropertyOutputReference
	// Experimental.
	MaintenanceStrategiesInput() *AwsFleet_MaintenanceStrategiesProperty
	// Experimental.
	MaxTotalPrice() *string
	// Experimental.
	SetMaxTotalPrice(val *string)
	// Experimental.
	MaxTotalPriceInput() *string
	// Experimental.
	MinTargetCapacity() *float64
	// Experimental.
	SetMinTargetCapacity(val *float64)
	// Experimental.
	MinTargetCapacityInput() *float64
	// Experimental.
	SingleAvailabilityZone() interface{}
	// Experimental.
	SetSingleAvailabilityZone(val interface{})
	// Experimental.
	SingleAvailabilityZoneInput() interface{}
	// Experimental.
	SingleInstanceType() interface{}
	// Experimental.
	SetSingleInstanceType(val interface{})
	// Experimental.
	SingleInstanceTypeInput() interface{}
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
	PutMaintenanceStrategies(value *AwsFleet_MaintenanceStrategiesProperty)
	// Experimental.
	ResetAllocationStrategy()
	// Experimental.
	ResetInstanceInterruptionBehavior()
	// Experimental.
	ResetInstancePoolsToUseCount()
	// Experimental.
	ResetMaintenanceStrategies()
	// Experimental.
	ResetMaxTotalPrice()
	// Experimental.
	ResetMinTargetCapacity()
	// Experimental.
	ResetSingleAvailabilityZone()
	// Experimental.
	ResetSingleInstanceType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsFleet_SpotOptionsPropertyOutputReference
type jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) AllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) AllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) InstanceInterruptionBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInterruptionBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) InstanceInterruptionBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInterruptionBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) InstancePoolsToUseCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancePoolsToUseCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) InstancePoolsToUseCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancePoolsToUseCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) InternalValue() *AwsFleet_SpotOptionsProperty {
	var returns *AwsFleet_SpotOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) MaintenanceStrategies() AwsFleet_MaintenanceStrategiesPropertyOutputReference {
	var returns AwsFleet_MaintenanceStrategiesPropertyOutputReference
	_jsii_.Get(
		j,
		"maintenanceStrategies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) MaintenanceStrategiesInput() *AwsFleet_MaintenanceStrategiesProperty {
	var returns *AwsFleet_MaintenanceStrategiesProperty
	_jsii_.Get(
		j,
		"maintenanceStrategiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) MaxTotalPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTotalPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) MaxTotalPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTotalPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) MinTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) MinTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) SingleAvailabilityZone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) SingleAvailabilityZoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleAvailabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) SingleInstanceType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) SingleInstanceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFleet_SpotOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFleet_SpotOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFleet_SpotOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsFleet.SpotOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFleet_SpotOptionsPropertyOutputReference_Override(a AwsFleet_SpotOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsFleet.SpotOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference)SetAllocationStrategy(val *string) {
	if err := j.validateSetAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationStrategy",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference)SetInstanceInterruptionBehavior(val *string) {
	if err := j.validateSetInstanceInterruptionBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInterruptionBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference)SetInstancePoolsToUseCount(val *float64) {
	if err := j.validateSetInstancePoolsToUseCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolsToUseCount",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference)SetInternalValue(val *AwsFleet_SpotOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference)SetMaxTotalPrice(val *string) {
	if err := j.validateSetMaxTotalPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTotalPrice",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference)SetMinTargetCapacity(val *float64) {
	if err := j.validateSetMinTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference)SetSingleAvailabilityZone(val interface{}) {
	if err := j.validateSetSingleAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference)SetSingleInstanceType(val interface{}) {
	if err := j.validateSetSingleInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleInstanceType",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) PutMaintenanceStrategies(value *AwsFleet_MaintenanceStrategiesProperty) {
	if err := a.validatePutMaintenanceStrategiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMaintenanceStrategies",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) ResetAllocationStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetAllocationStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) ResetInstanceInterruptionBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceInterruptionBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) ResetInstancePoolsToUseCount() {
	_jsii_.InvokeVoid(
		a,
		"resetInstancePoolsToUseCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) ResetMaintenanceStrategies() {
	_jsii_.InvokeVoid(
		a,
		"resetMaintenanceStrategies",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) ResetMaxTotalPrice() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxTotalPrice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) ResetMinTargetCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetMinTargetCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) ResetSingleAvailabilityZone() {
	_jsii_.InvokeVoid(
		a,
		"resetSingleAvailabilityZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) ResetSingleInstanceType() {
	_jsii_.InvokeVoid(
		a,
		"resetSingleInstanceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFleet_SpotOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

