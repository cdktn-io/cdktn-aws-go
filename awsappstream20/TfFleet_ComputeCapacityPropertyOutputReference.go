package awsappstream20

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappstream20/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappstream20/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFleet_ComputeCapacityPropertyOutputReference interface {
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
	InternalValue() *TfFleet_ComputeCapacityProperty
	// Experimental.
	SetInternalValue(val *TfFleet_ComputeCapacityProperty)
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

// The jsii proxy struct for TfFleet_ComputeCapacityPropertyOutputReference
type jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) Available() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"available",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) DesiredInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) DesiredInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) DesiredSessions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredSessions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) DesiredSessionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredSessionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) InternalValue() *TfFleet_ComputeCapacityProperty {
	var returns *TfFleet_ComputeCapacityProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) InUse() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inUse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) Running() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"running",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFleet_ComputeCapacityPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFleet_ComputeCapacityPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFleet_ComputeCapacityPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appstream-2-0.TfFleet.ComputeCapacityPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFleet_ComputeCapacityPropertyOutputReference_Override(t TfFleet_ComputeCapacityPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appstream-2-0.TfFleet.ComputeCapacityPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference)SetDesiredInstances(val *float64) {
	if err := j.validateSetDesiredInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredInstances",
		val,
	)
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference)SetDesiredSessions(val *float64) {
	if err := j.validateSetDesiredSessionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredSessions",
		val,
	)
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference)SetInternalValue(val *TfFleet_ComputeCapacityProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) ResetDesiredInstances() {
	_jsii_.InvokeVoid(
		t,
		"resetDesiredInstances",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) ResetDesiredSessions() {
	_jsii_.InvokeVoid(
		t,
		"resetDesiredSessions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFleet_ComputeCapacityPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

