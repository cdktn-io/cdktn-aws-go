package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseks/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfNodeGroup_NodeRepairConfigPropertyOutputReference interface {
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
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfNodeGroup_NodeRepairConfigProperty
	// Experimental.
	SetInternalValue(val *TfNodeGroup_NodeRepairConfigProperty)
	// Experimental.
	MaxParallelNodesRepairedCount() *float64
	// Experimental.
	SetMaxParallelNodesRepairedCount(val *float64)
	// Experimental.
	MaxParallelNodesRepairedCountInput() *float64
	// Experimental.
	MaxParallelNodesRepairedPercentage() *float64
	// Experimental.
	SetMaxParallelNodesRepairedPercentage(val *float64)
	// Experimental.
	MaxParallelNodesRepairedPercentageInput() *float64
	// Experimental.
	MaxUnhealthyNodeThresholdCount() *float64
	// Experimental.
	SetMaxUnhealthyNodeThresholdCount(val *float64)
	// Experimental.
	MaxUnhealthyNodeThresholdCountInput() *float64
	// Experimental.
	MaxUnhealthyNodeThresholdPercentage() *float64
	// Experimental.
	SetMaxUnhealthyNodeThresholdPercentage(val *float64)
	// Experimental.
	MaxUnhealthyNodeThresholdPercentageInput() *float64
	// Experimental.
	NodeRepairConfigOverrides() TfNodeGroup_NodeRepairConfigOverridesPropertyList
	// Experimental.
	NodeRepairConfigOverridesInput() interface{}
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
	PutNodeRepairConfigOverrides(value interface{})
	// Experimental.
	ResetEnabled()
	// Experimental.
	ResetMaxParallelNodesRepairedCount()
	// Experimental.
	ResetMaxParallelNodesRepairedPercentage()
	// Experimental.
	ResetMaxUnhealthyNodeThresholdCount()
	// Experimental.
	ResetMaxUnhealthyNodeThresholdPercentage()
	// Experimental.
	ResetNodeRepairConfigOverrides()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfNodeGroup_NodeRepairConfigPropertyOutputReference
type jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) InternalValue() *TfNodeGroup_NodeRepairConfigProperty {
	var returns *TfNodeGroup_NodeRepairConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) MaxParallelNodesRepairedCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelNodesRepairedCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) MaxParallelNodesRepairedCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelNodesRepairedCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) MaxParallelNodesRepairedPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelNodesRepairedPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) MaxParallelNodesRepairedPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelNodesRepairedPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) MaxUnhealthyNodeThresholdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyNodeThresholdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) MaxUnhealthyNodeThresholdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyNodeThresholdCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) MaxUnhealthyNodeThresholdPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyNodeThresholdPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) MaxUnhealthyNodeThresholdPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyNodeThresholdPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) NodeRepairConfigOverrides() TfNodeGroup_NodeRepairConfigOverridesPropertyList {
	var returns TfNodeGroup_NodeRepairConfigOverridesPropertyList
	_jsii_.Get(
		j,
		"nodeRepairConfigOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) NodeRepairConfigOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeRepairConfigOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfNodeGroup_NodeRepairConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfNodeGroup_NodeRepairConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfNodeGroup_NodeRepairConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eks.TfNodeGroup.NodeRepairConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfNodeGroup_NodeRepairConfigPropertyOutputReference_Override(t TfNodeGroup_NodeRepairConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.TfNodeGroup.NodeRepairConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference)SetInternalValue(val *TfNodeGroup_NodeRepairConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference)SetMaxParallelNodesRepairedCount(val *float64) {
	if err := j.validateSetMaxParallelNodesRepairedCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallelNodesRepairedCount",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference)SetMaxParallelNodesRepairedPercentage(val *float64) {
	if err := j.validateSetMaxParallelNodesRepairedPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallelNodesRepairedPercentage",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference)SetMaxUnhealthyNodeThresholdCount(val *float64) {
	if err := j.validateSetMaxUnhealthyNodeThresholdCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnhealthyNodeThresholdCount",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference)SetMaxUnhealthyNodeThresholdPercentage(val *float64) {
	if err := j.validateSetMaxUnhealthyNodeThresholdPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnhealthyNodeThresholdPercentage",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) PutNodeRepairConfigOverrides(value interface{}) {
	if err := t.validatePutNodeRepairConfigOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNodeRepairConfigOverrides",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) ResetMaxParallelNodesRepairedCount() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxParallelNodesRepairedCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) ResetMaxParallelNodesRepairedPercentage() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxParallelNodesRepairedPercentage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) ResetMaxUnhealthyNodeThresholdCount() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxUnhealthyNodeThresholdCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) ResetMaxUnhealthyNodeThresholdPercentage() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxUnhealthyNodeThresholdPercentage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) ResetNodeRepairConfigOverrides() {
	_jsii_.InvokeVoid(
		t,
		"resetNodeRepairConfigOverrides",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfNodeGroup_NodeRepairConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

