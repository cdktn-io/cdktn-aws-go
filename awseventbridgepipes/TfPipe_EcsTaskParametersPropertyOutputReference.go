package awseventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPipe_EcsTaskParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CapacityProviderStrategy() TfPipe_CapacityProviderStrategyPropertyList
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
	InternalValue() *TfPipe_EcsTaskParametersProperty
	// Experimental.
	SetInternalValue(val *TfPipe_EcsTaskParametersProperty)
	// Experimental.
	LaunchType() *string
	// Experimental.
	SetLaunchType(val *string)
	// Experimental.
	LaunchTypeInput() *string
	// Experimental.
	NetworkConfiguration() TfPipe_NetworkConfigurationPropertyOutputReference
	// Experimental.
	NetworkConfigurationInput() *TfPipe_NetworkConfigurationProperty
	// Experimental.
	Overrides() TfPipe_OverridesPropertyOutputReference
	// Experimental.
	OverridesInput() *TfPipe_OverridesProperty
	// Experimental.
	PlacementConstraint() TfPipe_PlacementConstraintPropertyList
	// Experimental.
	PlacementConstraintInput() interface{}
	// Experimental.
	PlacementStrategy() TfPipe_PlacementStrategyPropertyList
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
	PutNetworkConfiguration(value *TfPipe_NetworkConfigurationProperty)
	// Experimental.
	PutOverrides(value *TfPipe_OverridesProperty)
	// Experimental.
	PutPlacementConstraint(value interface{})
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
	ResetOverrides()
	// Experimental.
	ResetPlacementConstraint()
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

// The jsii proxy struct for TfPipe_EcsTaskParametersPropertyOutputReference
type jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) CapacityProviderStrategy() TfPipe_CapacityProviderStrategyPropertyList {
	var returns TfPipe_CapacityProviderStrategyPropertyList
	_jsii_.Get(
		j,
		"capacityProviderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) CapacityProviderStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityProviderStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) EnableEcsManagedTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) EnableEcsManagedTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) EnableExecuteCommand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) EnableExecuteCommandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) Group() *string {
	var returns *string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) GroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) InternalValue() *TfPipe_EcsTaskParametersProperty {
	var returns *TfPipe_EcsTaskParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) LaunchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) LaunchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) NetworkConfiguration() TfPipe_NetworkConfigurationPropertyOutputReference {
	var returns TfPipe_NetworkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) NetworkConfigurationInput() *TfPipe_NetworkConfigurationProperty {
	var returns *TfPipe_NetworkConfigurationProperty
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) Overrides() TfPipe_OverridesPropertyOutputReference {
	var returns TfPipe_OverridesPropertyOutputReference
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) OverridesInput() *TfPipe_OverridesProperty {
	var returns *TfPipe_OverridesProperty
	_jsii_.Get(
		j,
		"overridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) PlacementConstraint() TfPipe_PlacementConstraintPropertyList {
	var returns TfPipe_PlacementConstraintPropertyList
	_jsii_.Get(
		j,
		"placementConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) PlacementConstraintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) PlacementStrategy() TfPipe_PlacementStrategyPropertyList {
	var returns TfPipe_PlacementStrategyPropertyList
	_jsii_.Get(
		j,
		"placementStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) PlacementStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) PlatformVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) PropagateTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ReferenceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ReferenceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) TaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) TaskCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) TaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) TaskDefinitionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPipe_EcsTaskParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPipe_EcsTaskParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPipe_EcsTaskParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.EcsTaskParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPipe_EcsTaskParametersPropertyOutputReference_Override(t TfPipe_EcsTaskParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.EcsTaskParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference)SetEnableEcsManagedTags(val interface{}) {
	if err := j.validateSetEnableEcsManagedTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEcsManagedTags",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference)SetEnableExecuteCommand(val interface{}) {
	if err := j.validateSetEnableExecuteCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableExecuteCommand",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference)SetGroup(val *string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference)SetInternalValue(val *TfPipe_EcsTaskParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference)SetLaunchType(val *string) {
	if err := j.validateSetLaunchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchType",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference)SetPlatformVersion(val *string) {
	if err := j.validateSetPlatformVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformVersion",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference)SetPropagateTags(val *string) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference)SetReferenceId(val *string) {
	if err := j.validateSetReferenceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceId",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference)SetTaskCount(val *float64) {
	if err := j.validateSetTaskCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskCount",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference)SetTaskDefinitionArn(val *string) {
	if err := j.validateSetTaskDefinitionArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDefinitionArn",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) PutCapacityProviderStrategy(value interface{}) {
	if err := t.validatePutCapacityProviderStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCapacityProviderStrategy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) PutNetworkConfiguration(value *TfPipe_NetworkConfigurationProperty) {
	if err := t.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) PutOverrides(value *TfPipe_OverridesProperty) {
	if err := t.validatePutOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOverrides",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) PutPlacementConstraint(value interface{}) {
	if err := t.validatePutPlacementConstraintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPlacementConstraint",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) PutPlacementStrategy(value interface{}) {
	if err := t.validatePutPlacementStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPlacementStrategy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ResetCapacityProviderStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetCapacityProviderStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ResetEnableEcsManagedTags() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableEcsManagedTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ResetEnableExecuteCommand() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableExecuteCommand",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		t,
		"resetGroup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ResetLaunchType() {
	_jsii_.InvokeVoid(
		t,
		"resetLaunchType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ResetOverrides() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrides",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ResetPlacementConstraint() {
	_jsii_.InvokeVoid(
		t,
		"resetPlacementConstraint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ResetPlacementStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetPlacementStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ResetPlatformVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetPlatformVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		t,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ResetReferenceId() {
	_jsii_.InvokeVoid(
		t,
		"resetReferenceId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ResetTaskCount() {
	_jsii_.InvokeVoid(
		t,
		"resetTaskCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPipe_EcsTaskParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

