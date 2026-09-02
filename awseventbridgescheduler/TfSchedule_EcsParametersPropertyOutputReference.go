package awseventbridgescheduler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgescheduler/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgescheduler/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSchedule_EcsParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CapacityProviderStrategy() TfSchedule_CapacityProviderStrategyPropertyList
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
	InternalValue() *TfSchedule_EcsParametersProperty
	// Experimental.
	SetInternalValue(val *TfSchedule_EcsParametersProperty)
	// Experimental.
	LaunchType() *string
	// Experimental.
	SetLaunchType(val *string)
	// Experimental.
	LaunchTypeInput() *string
	// Experimental.
	NetworkConfiguration() TfSchedule_NetworkConfigurationPropertyOutputReference
	// Experimental.
	NetworkConfigurationInput() *TfSchedule_NetworkConfigurationProperty
	// Experimental.
	PlacementConstraints() TfSchedule_PlacementConstraintsPropertyList
	// Experimental.
	PlacementConstraintsInput() interface{}
	// Experimental.
	PlacementStrategy() TfSchedule_PlacementStrategyPropertyList
	// Experimental.
	PlacementStrategyInput() interface{}
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
	ReferenceId() *string
	// Experimental.
	SetReferenceId(val *string)
	// Experimental.
	ReferenceIdInput() *string
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
	PutNetworkConfiguration(value *TfSchedule_NetworkConfigurationProperty)
	// Experimental.
	PutPlacementConstraints(value interface{})
	// Experimental.
	PutPlacementStrategy(value interface{})
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
	ResetPlacementConstraints()
	// Experimental.
	ResetPlacementStrategy()
	// Experimental.
	ResetPlatformVersion()
	// Experimental.
	ResetPropagateTags()
	// Experimental.
	ResetReferenceId()
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

// The jsii proxy struct for TfSchedule_EcsParametersPropertyOutputReference
type jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) CapacityProviderStrategy() TfSchedule_CapacityProviderStrategyPropertyList {
	var returns TfSchedule_CapacityProviderStrategyPropertyList
	_jsii_.Get(
		j,
		"capacityProviderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) CapacityProviderStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityProviderStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) EnableEcsManagedTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) EnableEcsManagedTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) EnableExecuteCommand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) EnableExecuteCommandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) Group() *string {
	var returns *string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) GroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) InternalValue() *TfSchedule_EcsParametersProperty {
	var returns *TfSchedule_EcsParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) LaunchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) LaunchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) NetworkConfiguration() TfSchedule_NetworkConfigurationPropertyOutputReference {
	var returns TfSchedule_NetworkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) NetworkConfigurationInput() *TfSchedule_NetworkConfigurationProperty {
	var returns *TfSchedule_NetworkConfigurationProperty
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) PlacementConstraints() TfSchedule_PlacementConstraintsPropertyList {
	var returns TfSchedule_PlacementConstraintsPropertyList
	_jsii_.Get(
		j,
		"placementConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) PlacementConstraintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) PlacementStrategy() TfSchedule_PlacementStrategyPropertyList {
	var returns TfSchedule_PlacementStrategyPropertyList
	_jsii_.Get(
		j,
		"placementStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) PlacementStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) PlatformVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) PropagateTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ReferenceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ReferenceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) TaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) TaskCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) TaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) TaskDefinitionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfSchedule_EcsParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfSchedule_EcsParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfSchedule_EcsParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-scheduler.TfSchedule.EcsParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfSchedule_EcsParametersPropertyOutputReference_Override(t TfSchedule_EcsParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-scheduler.TfSchedule.EcsParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference)SetEnableEcsManagedTags(val interface{}) {
	if err := j.validateSetEnableEcsManagedTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEcsManagedTags",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference)SetEnableExecuteCommand(val interface{}) {
	if err := j.validateSetEnableExecuteCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableExecuteCommand",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference)SetGroup(val *string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference)SetInternalValue(val *TfSchedule_EcsParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference)SetLaunchType(val *string) {
	if err := j.validateSetLaunchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchType",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference)SetPlatformVersion(val *string) {
	if err := j.validateSetPlatformVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformVersion",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference)SetPropagateTags(val *string) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference)SetReferenceId(val *string) {
	if err := j.validateSetReferenceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceId",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference)SetTaskCount(val *float64) {
	if err := j.validateSetTaskCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskCount",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference)SetTaskDefinitionArn(val *string) {
	if err := j.validateSetTaskDefinitionArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDefinitionArn",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) PutCapacityProviderStrategy(value interface{}) {
	if err := t.validatePutCapacityProviderStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCapacityProviderStrategy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) PutNetworkConfiguration(value *TfSchedule_NetworkConfigurationProperty) {
	if err := t.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) PutPlacementConstraints(value interface{}) {
	if err := t.validatePutPlacementConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPlacementConstraints",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) PutPlacementStrategy(value interface{}) {
	if err := t.validatePutPlacementStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPlacementStrategy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ResetCapacityProviderStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetCapacityProviderStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ResetEnableEcsManagedTags() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableEcsManagedTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ResetEnableExecuteCommand() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableExecuteCommand",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		t,
		"resetGroup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ResetLaunchType() {
	_jsii_.InvokeVoid(
		t,
		"resetLaunchType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ResetPlacementConstraints() {
	_jsii_.InvokeVoid(
		t,
		"resetPlacementConstraints",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ResetPlacementStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetPlacementStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ResetPlatformVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetPlatformVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		t,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ResetReferenceId() {
	_jsii_.InvokeVoid(
		t,
		"resetReferenceId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ResetTaskCount() {
	_jsii_.InvokeVoid(
		t,
		"resetTaskCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfSchedule_EcsParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

