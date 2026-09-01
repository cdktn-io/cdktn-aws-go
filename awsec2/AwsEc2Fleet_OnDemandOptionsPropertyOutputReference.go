package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEc2Fleet_OnDemandOptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllocationStrategy() *string
	// Experimental.
	SetAllocationStrategy(val *string)
	// Experimental.
	AllocationStrategyInput() *string
	// Experimental.
	CapacityReservationOptions() AwsEc2Fleet_CapacityReservationOptionsPropertyOutputReference
	// Experimental.
	CapacityReservationOptionsInput() *AwsEc2Fleet_CapacityReservationOptionsProperty
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
	InternalValue() *AwsEc2Fleet_OnDemandOptionsProperty
	// Experimental.
	SetInternalValue(val *AwsEc2Fleet_OnDemandOptionsProperty)
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
	PutCapacityReservationOptions(value *AwsEc2Fleet_CapacityReservationOptionsProperty)
	// Experimental.
	ResetAllocationStrategy()
	// Experimental.
	ResetCapacityReservationOptions()
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

// The jsii proxy struct for AwsEc2Fleet_OnDemandOptionsPropertyOutputReference
type jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) AllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) AllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) CapacityReservationOptions() AwsEc2Fleet_CapacityReservationOptionsPropertyOutputReference {
	var returns AwsEc2Fleet_CapacityReservationOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"capacityReservationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) CapacityReservationOptionsInput() *AwsEc2Fleet_CapacityReservationOptionsProperty {
	var returns *AwsEc2Fleet_CapacityReservationOptionsProperty
	_jsii_.Get(
		j,
		"capacityReservationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) InternalValue() *AwsEc2Fleet_OnDemandOptionsProperty {
	var returns *AwsEc2Fleet_OnDemandOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) MaxTotalPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTotalPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) MaxTotalPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTotalPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) MinTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) MinTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) SingleAvailabilityZone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) SingleAvailabilityZoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleAvailabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) SingleInstanceType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) SingleInstanceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEc2Fleet_OnDemandOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEc2Fleet_OnDemandOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEc2Fleet_OnDemandOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsEc2Fleet.OnDemandOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEc2Fleet_OnDemandOptionsPropertyOutputReference_Override(a AwsEc2Fleet_OnDemandOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsEc2Fleet.OnDemandOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference)SetAllocationStrategy(val *string) {
	if err := j.validateSetAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationStrategy",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference)SetInternalValue(val *AwsEc2Fleet_OnDemandOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference)SetMaxTotalPrice(val *string) {
	if err := j.validateSetMaxTotalPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTotalPrice",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference)SetMinTargetCapacity(val *float64) {
	if err := j.validateSetMinTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference)SetSingleAvailabilityZone(val interface{}) {
	if err := j.validateSetSingleAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference)SetSingleInstanceType(val interface{}) {
	if err := j.validateSetSingleInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleInstanceType",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) PutCapacityReservationOptions(value *AwsEc2Fleet_CapacityReservationOptionsProperty) {
	if err := a.validatePutCapacityReservationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCapacityReservationOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) ResetAllocationStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetAllocationStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) ResetCapacityReservationOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacityReservationOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) ResetMaxTotalPrice() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxTotalPrice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) ResetMinTargetCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetMinTargetCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) ResetSingleAvailabilityZone() {
	_jsii_.InvokeVoid(
		a,
		"resetSingleAvailabilityZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) ResetSingleInstanceType() {
	_jsii_.InvokeVoid(
		a,
		"resetSingleInstanceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEc2Fleet_OnDemandOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

