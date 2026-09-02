package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFleet_OnDemandOptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllocationStrategy() *string
	// Experimental.
	SetAllocationStrategy(val *string)
	// Experimental.
	AllocationStrategyInput() *string
	// Experimental.
	CapacityReservationOptions() TfFleet_CapacityReservationOptionsPropertyOutputReference
	// Experimental.
	CapacityReservationOptionsInput() *TfFleet_CapacityReservationOptionsProperty
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
	InternalValue() *TfFleet_OnDemandOptionsProperty
	// Experimental.
	SetInternalValue(val *TfFleet_OnDemandOptionsProperty)
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
	PutCapacityReservationOptions(value *TfFleet_CapacityReservationOptionsProperty)
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

// The jsii proxy struct for TfFleet_OnDemandOptionsPropertyOutputReference
type jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) AllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) AllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) CapacityReservationOptions() TfFleet_CapacityReservationOptionsPropertyOutputReference {
	var returns TfFleet_CapacityReservationOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"capacityReservationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) CapacityReservationOptionsInput() *TfFleet_CapacityReservationOptionsProperty {
	var returns *TfFleet_CapacityReservationOptionsProperty
	_jsii_.Get(
		j,
		"capacityReservationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) InternalValue() *TfFleet_OnDemandOptionsProperty {
	var returns *TfFleet_OnDemandOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) MaxTotalPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTotalPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) MaxTotalPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTotalPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) MinTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) MinTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) SingleAvailabilityZone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) SingleAvailabilityZoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleAvailabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) SingleInstanceType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) SingleInstanceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFleet_OnDemandOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFleet_OnDemandOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFleet_OnDemandOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2.TfFleet.OnDemandOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFleet_OnDemandOptionsPropertyOutputReference_Override(t TfFleet_OnDemandOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.TfFleet.OnDemandOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference)SetAllocationStrategy(val *string) {
	if err := j.validateSetAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationStrategy",
		val,
	)
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference)SetInternalValue(val *TfFleet_OnDemandOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference)SetMaxTotalPrice(val *string) {
	if err := j.validateSetMaxTotalPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTotalPrice",
		val,
	)
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference)SetMinTargetCapacity(val *float64) {
	if err := j.validateSetMinTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference)SetSingleAvailabilityZone(val interface{}) {
	if err := j.validateSetSingleAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference)SetSingleInstanceType(val interface{}) {
	if err := j.validateSetSingleInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleInstanceType",
		val,
	)
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) PutCapacityReservationOptions(value *TfFleet_CapacityReservationOptionsProperty) {
	if err := t.validatePutCapacityReservationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCapacityReservationOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) ResetAllocationStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetAllocationStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) ResetCapacityReservationOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetCapacityReservationOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) ResetMaxTotalPrice() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxTotalPrice",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) ResetMinTargetCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetMinTargetCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) ResetSingleAvailabilityZone() {
	_jsii_.InvokeVoid(
		t,
		"resetSingleAvailabilityZone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) ResetSingleInstanceType() {
	_jsii_.InvokeVoid(
		t,
		"resetSingleInstanceType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFleet_OnDemandOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

