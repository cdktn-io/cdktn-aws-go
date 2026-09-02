package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseks/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfNodeGroup_WarmPoolConfigPropertyOutputReference interface {
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
	InternalValue() *TfNodeGroup_WarmPoolConfigProperty
	// Experimental.
	SetInternalValue(val *TfNodeGroup_WarmPoolConfigProperty)
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
	ReuseOnScaleIn() interface{}
	// Experimental.
	SetReuseOnScaleIn(val interface{})
	// Experimental.
	ReuseOnScaleInInput() interface{}
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
	ResetMaxGroupPreparedCapacity()
	// Experimental.
	ResetMinSize()
	// Experimental.
	ResetPoolState()
	// Experimental.
	ResetReuseOnScaleIn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfNodeGroup_WarmPoolConfigPropertyOutputReference
type jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) InternalValue() *TfNodeGroup_WarmPoolConfigProperty {
	var returns *TfNodeGroup_WarmPoolConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) MaxGroupPreparedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxGroupPreparedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) MaxGroupPreparedCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxGroupPreparedCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) PoolState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) PoolStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) ReuseOnScaleIn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reuseOnScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) ReuseOnScaleInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reuseOnScaleInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfNodeGroup_WarmPoolConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfNodeGroup_WarmPoolConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfNodeGroup_WarmPoolConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eks.TfNodeGroup.WarmPoolConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfNodeGroup_WarmPoolConfigPropertyOutputReference_Override(t TfNodeGroup_WarmPoolConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.TfNodeGroup.WarmPoolConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference)SetInternalValue(val *TfNodeGroup_WarmPoolConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference)SetMaxGroupPreparedCapacity(val *float64) {
	if err := j.validateSetMaxGroupPreparedCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxGroupPreparedCapacity",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference)SetPoolState(val *string) {
	if err := j.validateSetPoolStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"poolState",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference)SetReuseOnScaleIn(val interface{}) {
	if err := j.validateSetReuseOnScaleInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reuseOnScaleIn",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) ResetMaxGroupPreparedCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxGroupPreparedCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) ResetMinSize() {
	_jsii_.InvokeVoid(
		t,
		"resetMinSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) ResetPoolState() {
	_jsii_.InvokeVoid(
		t,
		"resetPoolState",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) ResetReuseOnScaleIn() {
	_jsii_.InvokeVoid(
		t,
		"resetReuseOnScaleIn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfNodeGroup_WarmPoolConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

