package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemr/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsemr/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MaximumCapacityUnits() *float64
	// Experimental.
	SetMaximumCapacityUnits(val *float64)
	// Experimental.
	MaximumCapacityUnitsInput() *float64
	// Experimental.
	MaximumCoreCapacityUnits() *float64
	// Experimental.
	SetMaximumCoreCapacityUnits(val *float64)
	// Experimental.
	MaximumCoreCapacityUnitsInput() *float64
	// Experimental.
	MaximumOndemandCapacityUnits() *float64
	// Experimental.
	SetMaximumOndemandCapacityUnits(val *float64)
	// Experimental.
	MaximumOndemandCapacityUnitsInput() *float64
	// Experimental.
	MinimumCapacityUnits() *float64
	// Experimental.
	SetMinimumCapacityUnits(val *float64)
	// Experimental.
	MinimumCapacityUnitsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UnitType() *string
	// Experimental.
	SetUnitType(val *string)
	// Experimental.
	UnitTypeInput() *string
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
	ResetMaximumCoreCapacityUnits()
	// Experimental.
	ResetMaximumOndemandCapacityUnits()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference
type jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) MaximumCapacityUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumCapacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) MaximumCapacityUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumCapacityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) MaximumCoreCapacityUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumCoreCapacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) MaximumCoreCapacityUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumCoreCapacityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) MaximumOndemandCapacityUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumOndemandCapacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) MaximumOndemandCapacityUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumOndemandCapacityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) MinimumCapacityUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCapacityUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) MinimumCapacityUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumCapacityUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) UnitType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) UnitTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitTypeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr.AwsEmrManagedScalingPolicy.ComputeLimitsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference_Override(a AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr.AwsEmrManagedScalingPolicy.ComputeLimitsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference)SetMaximumCapacityUnits(val *float64) {
	if err := j.validateSetMaximumCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumCapacityUnits",
		val,
	)
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference)SetMaximumCoreCapacityUnits(val *float64) {
	if err := j.validateSetMaximumCoreCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumCoreCapacityUnits",
		val,
	)
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference)SetMaximumOndemandCapacityUnits(val *float64) {
	if err := j.validateSetMaximumOndemandCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumOndemandCapacityUnits",
		val,
	)
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference)SetMinimumCapacityUnits(val *float64) {
	if err := j.validateSetMinimumCapacityUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumCapacityUnits",
		val,
	)
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference)SetUnitType(val *string) {
	if err := j.validateSetUnitTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unitType",
		val,
	)
}

func (a *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) ResetMaximumCoreCapacityUnits() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumCoreCapacityUnits",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) ResetMaximumOndemandCapacityUnits() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumOndemandCapacityUnits",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEmrManagedScalingPolicy_ComputeLimitsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

