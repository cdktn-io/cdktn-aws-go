package appstream20

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appstream20/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appstream20/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFleet_ComputeCapacityPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Available() *float64
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
	DesiredInstances() *float64
	// Experimental.
	SetDesiredInstances(val *float64)
	// Experimental.
	DesiredInstancesInput() *float64
	// Experimental.
	DesiredSessions() *float64
	// Experimental.
	SetDesiredSessions(val *float64)
	// Experimental.
	DesiredSessionsInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsFleet_ComputeCapacityProperty
	// Experimental.
	SetInternalValue(val *AwsFleet_ComputeCapacityProperty)
	// Experimental.
	InUse() *float64
	// Experimental.
	Running() *float64
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
	ResetDesiredInstances()
	// Experimental.
	ResetDesiredSessions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsFleet_ComputeCapacityPropertyOutputReference
type jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) Available() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"available",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) DesiredInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) DesiredInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) DesiredSessions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredSessions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) DesiredSessionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredSessionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) InternalValue() *AwsFleet_ComputeCapacityProperty {
	var returns *AwsFleet_ComputeCapacityProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) InUse() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inUse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) Running() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"running",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFleet_ComputeCapacityPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFleet_ComputeCapacityPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFleet_ComputeCapacityPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appstream-2-0.AwsFleet.ComputeCapacityPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFleet_ComputeCapacityPropertyOutputReference_Override(a AwsFleet_ComputeCapacityPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appstream-2-0.AwsFleet.ComputeCapacityPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference)SetDesiredInstances(val *float64) {
	if err := j.validateSetDesiredInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredInstances",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference)SetDesiredSessions(val *float64) {
	if err := j.validateSetDesiredSessionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredSessions",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference)SetInternalValue(val *AwsFleet_ComputeCapacityProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) ResetDesiredInstances() {
	_jsii_.InvokeVoid(
		a,
		"resetDesiredInstances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) ResetDesiredSessions() {
	_jsii_.InvokeVoid(
		a,
		"resetDesiredSessions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFleet_ComputeCapacityPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

