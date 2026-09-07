package eventbridge

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/eventbridge/jsii"

	"github.com/cdktn-io/cdktn-aws-go/eventbridge/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTarget_EcsTargetPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CapacityProviderStrategy() AwsTarget_CapacityProviderStrategyPropertyList
	// Experimental.
	CapacityProviderStrategyInput() interface{}
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
	EnableEcsManagedTags() interface{}
	// Experimental.
	SetEnableEcsManagedTags(val interface{})
	// Experimental.
	EnableEcsManagedTagsInput() interface{}
	// Experimental.
	EnableExecuteCommand() interface{}
	// Experimental.
	SetEnableExecuteCommand(val interface{})
	// Experimental.
	EnableExecuteCommandInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	Group() *string
	// Experimental.
	SetGroup(val *string)
	// Experimental.
	GroupInput() *string
	// Experimental.
	InternalValue() *AwsTarget_EcsTargetProperty
	// Experimental.
	SetInternalValue(val *AwsTarget_EcsTargetProperty)
	// Experimental.
	LaunchType() *string
	// Experimental.
	SetLaunchType(val *string)
	// Experimental.
	LaunchTypeInput() *string
	// Experimental.
	NetworkConfiguration() AwsTarget_NetworkConfigurationPropertyOutputReference
	// Experimental.
	NetworkConfigurationInput() *AwsTarget_NetworkConfigurationProperty
	// Experimental.
	OrderedPlacementStrategy() AwsTarget_OrderedPlacementStrategyPropertyList
	// Experimental.
	OrderedPlacementStrategyInput() interface{}
	// Experimental.
	PlacementConstraint() AwsTarget_PlacementConstraintPropertyList
	// Experimental.
	PlacementConstraintInput() interface{}
	// Experimental.
	PlatformVersion() *string
	// Experimental.
	SetPlatformVersion(val *string)
	// Experimental.
	PlatformVersionInput() *string
	// Experimental.
	PropagateTags() *string
	// Experimental.
	SetPropagateTags(val *string)
	// Experimental.
	PropagateTagsInput() *string
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TaskCount() *float64
	// Experimental.
	SetTaskCount(val *float64)
	// Experimental.
	TaskCountInput() *float64
	// Experimental.
	TaskDefinitionArn() *string
	// Experimental.
	SetTaskDefinitionArn(val *string)
	// Experimental.
	TaskDefinitionArnInput() *string
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
	PutCapacityProviderStrategy(value interface{})
	// Experimental.
	PutNetworkConfiguration(value *AwsTarget_NetworkConfigurationProperty)
	// Experimental.
	PutOrderedPlacementStrategy(value interface{})
	// Experimental.
	PutPlacementConstraint(value interface{})
	// Experimental.
	ResetCapacityProviderStrategy()
	// Experimental.
	ResetEnableEcsManagedTags()
	// Experimental.
	ResetEnableExecuteCommand()
	// Experimental.
	ResetGroup()
	// Experimental.
	ResetLaunchType()
	// Experimental.
	ResetNetworkConfiguration()
	// Experimental.
	ResetOrderedPlacementStrategy()
	// Experimental.
	ResetPlacementConstraint()
	// Experimental.
	ResetPlatformVersion()
	// Experimental.
	ResetPropagateTags()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTaskCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsTarget_EcsTargetPropertyOutputReference
type jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) CapacityProviderStrategy() AwsTarget_CapacityProviderStrategyPropertyList {
	var returns AwsTarget_CapacityProviderStrategyPropertyList
	_jsii_.Get(
		j,
		"capacityProviderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) CapacityProviderStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityProviderStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) EnableEcsManagedTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) EnableEcsManagedTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) EnableExecuteCommand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) EnableExecuteCommandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) Group() *string {
	var returns *string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) GroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) InternalValue() *AwsTarget_EcsTargetProperty {
	var returns *AwsTarget_EcsTargetProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) LaunchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) LaunchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) NetworkConfiguration() AwsTarget_NetworkConfigurationPropertyOutputReference {
	var returns AwsTarget_NetworkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) NetworkConfigurationInput() *AwsTarget_NetworkConfigurationProperty {
	var returns *AwsTarget_NetworkConfigurationProperty
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) OrderedPlacementStrategy() AwsTarget_OrderedPlacementStrategyPropertyList {
	var returns AwsTarget_OrderedPlacementStrategyPropertyList
	_jsii_.Get(
		j,
		"orderedPlacementStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) OrderedPlacementStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orderedPlacementStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) PlacementConstraint() AwsTarget_PlacementConstraintPropertyList {
	var returns AwsTarget_PlacementConstraintPropertyList
	_jsii_.Get(
		j,
		"placementConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) PlacementConstraintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) PlatformVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) PropagateTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) TaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) TaskCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) TaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) TaskDefinitionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTarget_EcsTargetPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsTarget_EcsTargetPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTarget_EcsTargetPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge.AwsTarget.EcsTargetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTarget_EcsTargetPropertyOutputReference_Override(a AwsTarget_EcsTargetPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge.AwsTarget.EcsTargetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference)SetEnableEcsManagedTags(val interface{}) {
	if err := j.validateSetEnableEcsManagedTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEcsManagedTags",
		val,
	)
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference)SetEnableExecuteCommand(val interface{}) {
	if err := j.validateSetEnableExecuteCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableExecuteCommand",
		val,
	)
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference)SetGroup(val *string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference)SetInternalValue(val *AwsTarget_EcsTargetProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference)SetLaunchType(val *string) {
	if err := j.validateSetLaunchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchType",
		val,
	)
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference)SetPlatformVersion(val *string) {
	if err := j.validateSetPlatformVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformVersion",
		val,
	)
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference)SetPropagateTags(val *string) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference)SetTaskCount(val *float64) {
	if err := j.validateSetTaskCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskCount",
		val,
	)
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference)SetTaskDefinitionArn(val *string) {
	if err := j.validateSetTaskDefinitionArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDefinitionArn",
		val,
	)
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) PutCapacityProviderStrategy(value interface{}) {
	if err := a.validatePutCapacityProviderStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCapacityProviderStrategy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) PutNetworkConfiguration(value *AwsTarget_NetworkConfigurationProperty) {
	if err := a.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) PutOrderedPlacementStrategy(value interface{}) {
	if err := a.validatePutOrderedPlacementStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOrderedPlacementStrategy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) PutPlacementConstraint(value interface{}) {
	if err := a.validatePutPlacementConstraintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPlacementConstraint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) ResetCapacityProviderStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacityProviderStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) ResetEnableEcsManagedTags() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableEcsManagedTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) ResetEnableExecuteCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableExecuteCommand",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) ResetLaunchType() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) ResetOrderedPlacementStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetOrderedPlacementStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) ResetPlacementConstraint() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacementConstraint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) ResetPlatformVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetPlatformVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		a,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) ResetTaskCount() {
	_jsii_.InvokeVoid(
		a,
		"resetTaskCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTarget_EcsTargetPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

