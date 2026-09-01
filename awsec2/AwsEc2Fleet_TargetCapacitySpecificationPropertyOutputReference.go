package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference interface {
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
	InternalValue() *AwsEc2Fleet_TargetCapacitySpecificationProperty
	// Experimental.
	SetInternalValue(val *AwsEc2Fleet_TargetCapacitySpecificationProperty)
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

// The jsii proxy struct for AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference
type jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) DefaultTargetCapacityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTargetCapacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) DefaultTargetCapacityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTargetCapacityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) InternalValue() *AwsEc2Fleet_TargetCapacitySpecificationProperty {
	var returns *AwsEc2Fleet_TargetCapacitySpecificationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) OnDemandTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) OnDemandTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) SpotTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) SpotTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"spotTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) TargetCapacityUnitType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacityUnitType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) TargetCapacityUnitTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacityUnitTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) TotalTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) TotalTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalTargetCapacityInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsEc2Fleet.TargetCapacitySpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference_Override(a AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsEc2Fleet.TargetCapacitySpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference)SetDefaultTargetCapacityType(val *string) {
	if err := j.validateSetDefaultTargetCapacityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTargetCapacityType",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference)SetInternalValue(val *AwsEc2Fleet_TargetCapacitySpecificationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference)SetOnDemandTargetCapacity(val *float64) {
	if err := j.validateSetOnDemandTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference)SetSpotTargetCapacity(val *float64) {
	if err := j.validateSetSpotTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference)SetTargetCapacityUnitType(val *string) {
	if err := j.validateSetTargetCapacityUnitTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCapacityUnitType",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference)SetTotalTargetCapacity(val *float64) {
	if err := j.validateSetTotalTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalTargetCapacity",
		val,
	)
}

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) ResetOnDemandTargetCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandTargetCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) ResetSpotTargetCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotTargetCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) ResetTargetCapacityUnitType() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetCapacityUnitType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

