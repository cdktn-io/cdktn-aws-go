package awselementalmediaconvert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmediaconvert/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmediaconvert/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Commitment() *string
	// Experimental.
	SetCommitment(val *string)
	// Experimental.
	CommitmentInput() *string
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
	InternalValue() *AwsMediaConvertQueue_ReservationPlanSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsMediaConvertQueue_ReservationPlanSettingsProperty)
	// Experimental.
	RenewalType() *string
	// Experimental.
	SetRenewalType(val *string)
	// Experimental.
	RenewalTypeInput() *string
	// Experimental.
	ReservedSlots() *float64
	// Experimental.
	SetReservedSlots(val *float64)
	// Experimental.
	ReservedSlotsInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference
type jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) Commitment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) CommitmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) InternalValue() *AwsMediaConvertQueue_ReservationPlanSettingsProperty {
	var returns *AwsMediaConvertQueue_ReservationPlanSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) RenewalType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renewalType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) RenewalTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renewalTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) ReservedSlots() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedSlots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) ReservedSlotsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedSlotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-mediaconvert.AwsMediaConvertQueue.ReservationPlanSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference_Override(a AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-mediaconvert.AwsMediaConvertQueue.ReservationPlanSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference)SetCommitment(val *string) {
	if err := j.validateSetCommitmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitment",
		val,
	)
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference)SetInternalValue(val *AwsMediaConvertQueue_ReservationPlanSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference)SetRenewalType(val *string) {
	if err := j.validateSetRenewalTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renewalType",
		val,
	)
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference)SetReservedSlots(val *float64) {
	if err := j.validateSetReservedSlotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedSlots",
		val,
	)
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMediaConvertQueue_ReservationPlanSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

