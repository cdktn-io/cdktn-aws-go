package eks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/eks/jsii"

	"github.com/cdktn-io/cdktn-aws-go/eks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsNodeGroup_NodeRepairConfigPropertyOutputReference interface {
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
	InternalValue() *AwsNodeGroup_NodeRepairConfigProperty
	// Experimental.
	SetInternalValue(val *AwsNodeGroup_NodeRepairConfigProperty)
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
	NodeRepairConfigOverrides() AwsNodeGroup_NodeRepairConfigOverridesPropertyList
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

// The jsii proxy struct for AwsNodeGroup_NodeRepairConfigPropertyOutputReference
type jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) InternalValue() *AwsNodeGroup_NodeRepairConfigProperty {
	var returns *AwsNodeGroup_NodeRepairConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) MaxParallelNodesRepairedCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelNodesRepairedCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) MaxParallelNodesRepairedCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelNodesRepairedCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) MaxParallelNodesRepairedPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelNodesRepairedPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) MaxParallelNodesRepairedPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelNodesRepairedPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) MaxUnhealthyNodeThresholdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyNodeThresholdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) MaxUnhealthyNodeThresholdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyNodeThresholdCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) MaxUnhealthyNodeThresholdPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyNodeThresholdPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) MaxUnhealthyNodeThresholdPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnhealthyNodeThresholdPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) NodeRepairConfigOverrides() AwsNodeGroup_NodeRepairConfigOverridesPropertyList {
	var returns AwsNodeGroup_NodeRepairConfigOverridesPropertyList
	_jsii_.Get(
		j,
		"nodeRepairConfigOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) NodeRepairConfigOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeRepairConfigOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsNodeGroup_NodeRepairConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsNodeGroup_NodeRepairConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsNodeGroup_NodeRepairConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eks.AwsNodeGroup.NodeRepairConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsNodeGroup_NodeRepairConfigPropertyOutputReference_Override(a AwsNodeGroup_NodeRepairConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.AwsNodeGroup.NodeRepairConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference)SetInternalValue(val *AwsNodeGroup_NodeRepairConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference)SetMaxParallelNodesRepairedCount(val *float64) {
	if err := j.validateSetMaxParallelNodesRepairedCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallelNodesRepairedCount",
		val,
	)
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference)SetMaxParallelNodesRepairedPercentage(val *float64) {
	if err := j.validateSetMaxParallelNodesRepairedPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallelNodesRepairedPercentage",
		val,
	)
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference)SetMaxUnhealthyNodeThresholdCount(val *float64) {
	if err := j.validateSetMaxUnhealthyNodeThresholdCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnhealthyNodeThresholdCount",
		val,
	)
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference)SetMaxUnhealthyNodeThresholdPercentage(val *float64) {
	if err := j.validateSetMaxUnhealthyNodeThresholdPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnhealthyNodeThresholdPercentage",
		val,
	)
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) PutNodeRepairConfigOverrides(value interface{}) {
	if err := a.validatePutNodeRepairConfigOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNodeRepairConfigOverrides",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) ResetMaxParallelNodesRepairedCount() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxParallelNodesRepairedCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) ResetMaxParallelNodesRepairedPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxParallelNodesRepairedPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) ResetMaxUnhealthyNodeThresholdCount() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxUnhealthyNodeThresholdCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) ResetMaxUnhealthyNodeThresholdPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxUnhealthyNodeThresholdPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) ResetNodeRepairConfigOverrides() {
	_jsii_.InvokeVoid(
		a,
		"resetNodeRepairConfigOverrides",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsNodeGroup_NodeRepairConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

