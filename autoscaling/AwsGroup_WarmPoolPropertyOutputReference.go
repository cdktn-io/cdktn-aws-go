package autoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/autoscaling/jsii"

	"github.com/cdktn-io/cdktn-aws-go/autoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGroup_WarmPoolPropertyOutputReference interface {
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
	InstanceReusePolicy() AwsGroup_InstanceReusePolicyPropertyOutputReference
	// Experimental.
	InstanceReusePolicyInput() *AwsGroup_InstanceReusePolicyProperty
	// Experimental.
	InternalValue() *AwsGroup_WarmPoolProperty
	// Experimental.
	SetInternalValue(val *AwsGroup_WarmPoolProperty)
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
	PutInstanceReusePolicy(value *AwsGroup_InstanceReusePolicyProperty)
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

// The jsii proxy struct for AwsGroup_WarmPoolPropertyOutputReference
type jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) InstanceReusePolicy() AwsGroup_InstanceReusePolicyPropertyOutputReference {
	var returns AwsGroup_InstanceReusePolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"instanceReusePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) InstanceReusePolicyInput() *AwsGroup_InstanceReusePolicyProperty {
	var returns *AwsGroup_InstanceReusePolicyProperty
	_jsii_.Get(
		j,
		"instanceReusePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) InternalValue() *AwsGroup_WarmPoolProperty {
	var returns *AwsGroup_WarmPoolProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) MaxGroupPreparedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxGroupPreparedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) MaxGroupPreparedCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxGroupPreparedCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) PoolState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) PoolStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGroup_WarmPoolPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsGroup_WarmPoolPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGroup_WarmPoolPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.AwsGroup.WarmPoolPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGroup_WarmPoolPropertyOutputReference_Override(a AwsGroup_WarmPoolPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.AwsGroup.WarmPoolPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference)SetInternalValue(val *AwsGroup_WarmPoolProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference)SetMaxGroupPreparedCapacity(val *float64) {
	if err := j.validateSetMaxGroupPreparedCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxGroupPreparedCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference)SetPoolState(val *string) {
	if err := j.validateSetPoolStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"poolState",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) PutInstanceReusePolicy(value *AwsGroup_InstanceReusePolicyProperty) {
	if err := a.validatePutInstanceReusePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstanceReusePolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) ResetInstanceReusePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceReusePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) ResetMaxGroupPreparedCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxGroupPreparedCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) ResetMinSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMinSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) ResetPoolState() {
	_jsii_.InvokeVoid(
		a,
		"resetPoolState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGroup_WarmPoolPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

