package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFleet_TargetCapacitySpecificationPropertyOutputReference interface {
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
	DefaultTargetCapacityType() *string
	// Experimental.
	SetDefaultTargetCapacityType(val *string)
	// Experimental.
	DefaultTargetCapacityTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfFleet_TargetCapacitySpecificationProperty
	// Experimental.
	SetInternalValue(val *TfFleet_TargetCapacitySpecificationProperty)
	// Experimental.
	OnDemandTargetCapacity() *float64
	// Experimental.
	SetOnDemandTargetCapacity(val *float64)
	// Experimental.
	OnDemandTargetCapacityInput() *float64
	// Experimental.
	SpotTargetCapacity() *float64
	// Experimental.
	SetSpotTargetCapacity(val *float64)
	// Experimental.
	SpotTargetCapacityInput() *float64
	// Experimental.
	TargetCapacityUnitType() *string
	// Experimental.
	SetTargetCapacityUnitType(val *string)
	// Experimental.
	TargetCapacityUnitTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TotalTargetCapacity() *float64
	// Experimental.
	SetTotalTargetCapacity(val *float64)
	// Experimental.
	TotalTargetCapacityInput() *float64
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
	ResetOnDemandTargetCapacity()
	// Experimental.
	ResetSpotTargetCapacity()
	// Experimental.
	ResetTargetCapacityUnitType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFleet_TargetCapacitySpecificationPropertyOutputReference
type jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) DefaultTargetCapacityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTargetCapacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) DefaultTargetCapacityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTargetCapacityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) InternalValue() *TfFleet_TargetCapacitySpecificationProperty {
	var returns *TfFleet_TargetCapacitySpecificationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) OnDemandTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) OnDemandTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) SpotTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) SpotTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) TargetCapacityUnitType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacityUnitType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) TargetCapacityUnitTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacityUnitTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) TotalTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) TotalTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalTargetCapacityInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFleet_TargetCapacitySpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFleet_TargetCapacitySpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFleet_TargetCapacitySpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2.TfFleet.TargetCapacitySpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFleet_TargetCapacitySpecificationPropertyOutputReference_Override(t TfFleet_TargetCapacitySpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.TfFleet.TargetCapacitySpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference)SetDefaultTargetCapacityType(val *string) {
	if err := j.validateSetDefaultTargetCapacityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTargetCapacityType",
		val,
	)
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference)SetInternalValue(val *TfFleet_TargetCapacitySpecificationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference)SetOnDemandTargetCapacity(val *float64) {
	if err := j.validateSetOnDemandTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference)SetSpotTargetCapacity(val *float64) {
	if err := j.validateSetSpotTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference)SetTargetCapacityUnitType(val *string) {
	if err := j.validateSetTargetCapacityUnitTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCapacityUnitType",
		val,
	)
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference)SetTotalTargetCapacity(val *float64) {
	if err := j.validateSetTotalTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalTargetCapacity",
		val,
	)
}

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) ResetOnDemandTargetCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetOnDemandTargetCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) ResetSpotTargetCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetSpotTargetCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) ResetTargetCapacityUnitType() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetCapacityUnitType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFleet_TargetCapacitySpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

