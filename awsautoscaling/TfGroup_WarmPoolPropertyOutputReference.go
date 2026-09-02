package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsautoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsautoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGroup_WarmPoolPropertyOutputReference interface {
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
	InstanceReusePolicy() TfGroup_InstanceReusePolicyPropertyOutputReference
	// Experimental.
	InstanceReusePolicyInput() *TfGroup_InstanceReusePolicyProperty
	// Experimental.
	InternalValue() *TfGroup_WarmPoolProperty
	// Experimental.
	SetInternalValue(val *TfGroup_WarmPoolProperty)
	// Experimental.
	MaxGroupPreparedCapacity() *float64
	// Experimental.
	SetMaxGroupPreparedCapacity(val *float64)
	// Experimental.
	MaxGroupPreparedCapacityInput() *float64
	// Experimental.
	MinSize() *float64
	// Experimental.
	SetMinSize(val *float64)
	// Experimental.
	MinSizeInput() *float64
	// Experimental.
	PoolState() *string
	// Experimental.
	SetPoolState(val *string)
	// Experimental.
	PoolStateInput() *string
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
	PutInstanceReusePolicy(value *TfGroup_InstanceReusePolicyProperty)
	// Experimental.
	ResetInstanceReusePolicy()
	// Experimental.
	ResetMaxGroupPreparedCapacity()
	// Experimental.
	ResetMinSize()
	// Experimental.
	ResetPoolState()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfGroup_WarmPoolPropertyOutputReference
type jsiiProxy_TfGroup_WarmPoolPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) InstanceReusePolicy() TfGroup_InstanceReusePolicyPropertyOutputReference {
	var returns TfGroup_InstanceReusePolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"instanceReusePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) InstanceReusePolicyInput() *TfGroup_InstanceReusePolicyProperty {
	var returns *TfGroup_InstanceReusePolicyProperty
	_jsii_.Get(
		j,
		"instanceReusePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) InternalValue() *TfGroup_WarmPoolProperty {
	var returns *TfGroup_WarmPoolProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) MaxGroupPreparedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxGroupPreparedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) MaxGroupPreparedCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxGroupPreparedCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) PoolState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) PoolStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGroup_WarmPoolPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfGroup_WarmPoolPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGroup_WarmPoolPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGroup_WarmPoolPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.TfGroup.WarmPoolPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGroup_WarmPoolPropertyOutputReference_Override(t TfGroup_WarmPoolPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.TfGroup.WarmPoolPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference)SetInternalValue(val *TfGroup_WarmPoolProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference)SetMaxGroupPreparedCapacity(val *float64) {
	if err := j.validateSetMaxGroupPreparedCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxGroupPreparedCapacity",
		val,
	)
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference)SetPoolState(val *string) {
	if err := j.validateSetPoolStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"poolState",
		val,
	)
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) PutInstanceReusePolicy(value *TfGroup_InstanceReusePolicyProperty) {
	if err := t.validatePutInstanceReusePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInstanceReusePolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) ResetInstanceReusePolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceReusePolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) ResetMaxGroupPreparedCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxGroupPreparedCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) ResetMinSize() {
	_jsii_.InvokeVoid(
		t,
		"resetMinSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) ResetPoolState() {
	_jsii_.InvokeVoid(
		t,
		"resetPoolState",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGroup_WarmPoolPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

