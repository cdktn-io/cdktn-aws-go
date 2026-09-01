package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodedeploy/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awscodedeploy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group aws_codedeploy_deployment_group}.
// Experimental.
type AwsCodedeployDeploymentGroup interface {
	cdktn.TerraformResource
	// Experimental.
	AlarmConfiguration() AwsCodedeployDeploymentGroup_AlarmConfigurationPropertyOutputReference
	// Experimental.
	AlarmConfigurationInput() *AwsCodedeployDeploymentGroup_AlarmConfigurationProperty
	// Experimental.
	AppName() *string
	// Experimental.
	SetAppName(val *string)
	// Experimental.
	AppNameInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	AutoRollbackConfiguration() AwsCodedeployDeploymentGroup_AutoRollbackConfigurationPropertyOutputReference
	// Experimental.
	AutoRollbackConfigurationInput() *AwsCodedeployDeploymentGroup_AutoRollbackConfigurationProperty
	// Experimental.
	AutoscalingGroups() *[]*string
	// Experimental.
	SetAutoscalingGroups(val *[]*string)
	// Experimental.
	AutoscalingGroupsInput() *[]*string
	// Experimental.
	BlueGreenDeploymentConfig() AwsCodedeployDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference
	// Experimental.
	BlueGreenDeploymentConfigInput() *AwsCodedeployDeploymentGroup_BlueGreenDeploymentConfigProperty
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ComputePlatform() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DeploymentConfigName() *string
	// Experimental.
	SetDeploymentConfigName(val *string)
	// Experimental.
	DeploymentConfigNameInput() *string
	// Experimental.
	DeploymentGroupId() *string
	// Experimental.
	DeploymentGroupName() *string
	// Experimental.
	SetDeploymentGroupName(val *string)
	// Experimental.
	DeploymentGroupNameInput() *string
	// Experimental.
	DeploymentStyle() AwsCodedeployDeploymentGroup_DeploymentStylePropertyOutputReference
	// Experimental.
	DeploymentStyleInput() *AwsCodedeployDeploymentGroup_DeploymentStyleProperty
	// Experimental.
	Ec2TagFilter() AwsCodedeployDeploymentGroup_Ec2TagFilterPropertyList
	// Experimental.
	Ec2TagFilterInput() interface{}
	// Experimental.
	Ec2TagSet() AwsCodedeployDeploymentGroup_Ec2TagSetPropertyList
	// Experimental.
	Ec2TagSetInput() interface{}
	// Experimental.
	EcsService() AwsCodedeployDeploymentGroup_EcsServicePropertyOutputReference
	// Experimental.
	EcsServiceInput() *AwsCodedeployDeploymentGroup_EcsServiceProperty
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LoadBalancerInfo() AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference
	// Experimental.
	LoadBalancerInfoInput() *AwsCodedeployDeploymentGroup_LoadBalancerInfoProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OnPremisesInstanceTagFilter() AwsCodedeployDeploymentGroup_OnPremisesInstanceTagFilterPropertyList
	// Experimental.
	OnPremisesInstanceTagFilterInput() interface{}
	// Experimental.
	OutdatedInstancesStrategy() *string
	// Experimental.
	SetOutdatedInstancesStrategy(val *string)
	// Experimental.
	OutdatedInstancesStrategyInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	ServiceRoleArn() *string
	// Experimental.
	SetServiceRoleArn(val *string)
	// Experimental.
	ServiceRoleArnInput() *string
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsAll() *map[string]*string
	// Experimental.
	SetTagsAll(val *map[string]*string)
	// Experimental.
	TagsAllInput() *map[string]*string
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TerminationHookEnabled() interface{}
	// Experimental.
	SetTerminationHookEnabled(val interface{})
	// Experimental.
	TerminationHookEnabledInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	TriggerConfiguration() AwsCodedeployDeploymentGroup_TriggerConfigurationPropertyList
	// Experimental.
	TriggerConfigurationInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Experimental.
	PutAlarmConfiguration(value *AwsCodedeployDeploymentGroup_AlarmConfigurationProperty)
	// Experimental.
	PutAutoRollbackConfiguration(value *AwsCodedeployDeploymentGroup_AutoRollbackConfigurationProperty)
	// Experimental.
	PutBlueGreenDeploymentConfig(value *AwsCodedeployDeploymentGroup_BlueGreenDeploymentConfigProperty)
	// Experimental.
	PutDeploymentStyle(value *AwsCodedeployDeploymentGroup_DeploymentStyleProperty)
	// Experimental.
	PutEc2TagFilter(value interface{})
	// Experimental.
	PutEc2TagSet(value interface{})
	// Experimental.
	PutEcsService(value *AwsCodedeployDeploymentGroup_EcsServiceProperty)
	// Experimental.
	PutLoadBalancerInfo(value *AwsCodedeployDeploymentGroup_LoadBalancerInfoProperty)
	// Experimental.
	PutOnPremisesInstanceTagFilter(value interface{})
	// Experimental.
	PutTriggerConfiguration(value interface{})
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	// Experimental.
	ResetAlarmConfiguration()
	// Experimental.
	ResetAutoRollbackConfiguration()
	// Experimental.
	ResetAutoscalingGroups()
	// Experimental.
	ResetBlueGreenDeploymentConfig()
	// Experimental.
	ResetDeploymentConfigName()
	// Experimental.
	ResetDeploymentStyle()
	// Experimental.
	ResetEc2TagFilter()
	// Experimental.
	ResetEc2TagSet()
	// Experimental.
	ResetEcsService()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLoadBalancerInfo()
	// Experimental.
	ResetOnPremisesInstanceTagFilter()
	// Experimental.
	ResetOutdatedInstancesStrategy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTerminationHookEnabled()
	// Experimental.
	ResetTriggerConfiguration()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for AwsCodedeployDeploymentGroup
type jsiiProxy_AwsCodedeployDeploymentGroup struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) AlarmConfiguration() AwsCodedeployDeploymentGroup_AlarmConfigurationPropertyOutputReference {
	var returns AwsCodedeployDeploymentGroup_AlarmConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"alarmConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) AlarmConfigurationInput() *AwsCodedeployDeploymentGroup_AlarmConfigurationProperty {
	var returns *AwsCodedeployDeploymentGroup_AlarmConfigurationProperty
	_jsii_.Get(
		j,
		"alarmConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) AppName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) AppNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) AutoRollbackConfiguration() AwsCodedeployDeploymentGroup_AutoRollbackConfigurationPropertyOutputReference {
	var returns AwsCodedeployDeploymentGroup_AutoRollbackConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"autoRollbackConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) AutoRollbackConfigurationInput() *AwsCodedeployDeploymentGroup_AutoRollbackConfigurationProperty {
	var returns *AwsCodedeployDeploymentGroup_AutoRollbackConfigurationProperty
	_jsii_.Get(
		j,
		"autoRollbackConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) AutoscalingGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoscalingGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) AutoscalingGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoscalingGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) BlueGreenDeploymentConfig() AwsCodedeployDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference {
	var returns AwsCodedeployDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"blueGreenDeploymentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) BlueGreenDeploymentConfigInput() *AwsCodedeployDeploymentGroup_BlueGreenDeploymentConfigProperty {
	var returns *AwsCodedeployDeploymentGroup_BlueGreenDeploymentConfigProperty
	_jsii_.Get(
		j,
		"blueGreenDeploymentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) ComputePlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) DeploymentConfigNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) DeploymentGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) DeploymentGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) DeploymentGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) DeploymentStyle() AwsCodedeployDeploymentGroup_DeploymentStylePropertyOutputReference {
	var returns AwsCodedeployDeploymentGroup_DeploymentStylePropertyOutputReference
	_jsii_.Get(
		j,
		"deploymentStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) DeploymentStyleInput() *AwsCodedeployDeploymentGroup_DeploymentStyleProperty {
	var returns *AwsCodedeployDeploymentGroup_DeploymentStyleProperty
	_jsii_.Get(
		j,
		"deploymentStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) Ec2TagFilter() AwsCodedeployDeploymentGroup_Ec2TagFilterPropertyList {
	var returns AwsCodedeployDeploymentGroup_Ec2TagFilterPropertyList
	_jsii_.Get(
		j,
		"ec2TagFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) Ec2TagFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2TagFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) Ec2TagSet() AwsCodedeployDeploymentGroup_Ec2TagSetPropertyList {
	var returns AwsCodedeployDeploymentGroup_Ec2TagSetPropertyList
	_jsii_.Get(
		j,
		"ec2TagSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) Ec2TagSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2TagSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) EcsService() AwsCodedeployDeploymentGroup_EcsServicePropertyOutputReference {
	var returns AwsCodedeployDeploymentGroup_EcsServicePropertyOutputReference
	_jsii_.Get(
		j,
		"ecsService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) EcsServiceInput() *AwsCodedeployDeploymentGroup_EcsServiceProperty {
	var returns *AwsCodedeployDeploymentGroup_EcsServiceProperty
	_jsii_.Get(
		j,
		"ecsServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) LoadBalancerInfo() AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference {
	var returns AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"loadBalancerInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) LoadBalancerInfoInput() *AwsCodedeployDeploymentGroup_LoadBalancerInfoProperty {
	var returns *AwsCodedeployDeploymentGroup_LoadBalancerInfoProperty
	_jsii_.Get(
		j,
		"loadBalancerInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) OnPremisesInstanceTagFilter() AwsCodedeployDeploymentGroup_OnPremisesInstanceTagFilterPropertyList {
	var returns AwsCodedeployDeploymentGroup_OnPremisesInstanceTagFilterPropertyList
	_jsii_.Get(
		j,
		"onPremisesInstanceTagFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) OnPremisesInstanceTagFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onPremisesInstanceTagFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) OutdatedInstancesStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdatedInstancesStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) OutdatedInstancesStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdatedInstancesStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) ServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) TerminationHookEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationHookEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) TerminationHookEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationHookEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) TriggerConfiguration() AwsCodedeployDeploymentGroup_TriggerConfigurationPropertyList {
	var returns AwsCodedeployDeploymentGroup_TriggerConfigurationPropertyList
	_jsii_.Get(
		j,
		"triggerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup) TriggerConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"triggerConfigurationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group aws_codedeploy_deployment_group} Resource.
// Experimental.
func NewAwsCodedeployDeploymentGroup(scope constructs.Construct, id *string, config *AwsCodedeployDeploymentGroupConfig) AwsCodedeployDeploymentGroup {
	_init_.Initialize()

	if err := validateNewAwsCodedeployDeploymentGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCodedeployDeploymentGroup{}

	_jsii_.Create(
		"@cdktn/aws-codedeploy.AwsCodedeployDeploymentGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group aws_codedeploy_deployment_group} Resource.
// Experimental.
func NewAwsCodedeployDeploymentGroup_Override(a AwsCodedeployDeploymentGroup, scope constructs.Construct, id *string, config *AwsCodedeployDeploymentGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codedeploy.AwsCodedeployDeploymentGroup",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetAppName(val *string) {
	if err := j.validateSetAppNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appName",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetAutoscalingGroups(val *[]*string) {
	if err := j.validateSetAutoscalingGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingGroups",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetDeploymentConfigName(val *string) {
	if err := j.validateSetDeploymentConfigNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentConfigName",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetDeploymentGroupName(val *string) {
	if err := j.validateSetDeploymentGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetOutdatedInstancesStrategy(val *string) {
	if err := j.validateSetOutdatedInstancesStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outdatedInstancesStrategy",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetServiceRoleArn(val *string) {
	if err := j.validateSetServiceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup)SetTerminationHookEnabled(val interface{}) {
	if err := j.validateSetTerminationHookEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationHookEnabled",
		val,
	)
}

// Generates CDKTN code for importing a AwsCodedeployDeploymentGroup resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsCodedeployDeploymentGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsCodedeployDeploymentGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-codedeploy.AwsCodedeployDeploymentGroup",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func AwsCodedeployDeploymentGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCodedeployDeploymentGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-codedeploy.AwsCodedeployDeploymentGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCodedeployDeploymentGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCodedeployDeploymentGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-codedeploy.AwsCodedeployDeploymentGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCodedeployDeploymentGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCodedeployDeploymentGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-codedeploy.AwsCodedeployDeploymentGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsCodedeployDeploymentGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-codedeploy.AwsCodedeployDeploymentGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := a.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) PutAlarmConfiguration(value *AwsCodedeployDeploymentGroup_AlarmConfigurationProperty) {
	if err := a.validatePutAlarmConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAlarmConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) PutAutoRollbackConfiguration(value *AwsCodedeployDeploymentGroup_AutoRollbackConfigurationProperty) {
	if err := a.validatePutAutoRollbackConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutoRollbackConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) PutBlueGreenDeploymentConfig(value *AwsCodedeployDeploymentGroup_BlueGreenDeploymentConfigProperty) {
	if err := a.validatePutBlueGreenDeploymentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBlueGreenDeploymentConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) PutDeploymentStyle(value *AwsCodedeployDeploymentGroup_DeploymentStyleProperty) {
	if err := a.validatePutDeploymentStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeploymentStyle",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) PutEc2TagFilter(value interface{}) {
	if err := a.validatePutEc2TagFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEc2TagFilter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) PutEc2TagSet(value interface{}) {
	if err := a.validatePutEc2TagSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEc2TagSet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) PutEcsService(value *AwsCodedeployDeploymentGroup_EcsServiceProperty) {
	if err := a.validatePutEcsServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcsService",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) PutLoadBalancerInfo(value *AwsCodedeployDeploymentGroup_LoadBalancerInfoProperty) {
	if err := a.validatePutLoadBalancerInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLoadBalancerInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) PutOnPremisesInstanceTagFilter(value interface{}) {
	if err := a.validatePutOnPremisesInstanceTagFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOnPremisesInstanceTagFilter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) PutTriggerConfiguration(value interface{}) {
	if err := a.validatePutTriggerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTriggerConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetAlarmConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAlarmConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetAutoRollbackConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoRollbackConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetAutoscalingGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoscalingGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetBlueGreenDeploymentConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetBlueGreenDeploymentConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetDeploymentConfigName() {
	_jsii_.InvokeVoid(
		a,
		"resetDeploymentConfigName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetDeploymentStyle() {
	_jsii_.InvokeVoid(
		a,
		"resetDeploymentStyle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetEc2TagFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2TagFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetEc2TagSet() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2TagSet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetEcsService() {
	_jsii_.InvokeVoid(
		a,
		"resetEcsService",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetLoadBalancerInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancerInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetOnPremisesInstanceTagFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetOnPremisesInstanceTagFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetOutdatedInstancesStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetOutdatedInstancesStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetTerminationHookEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTerminationHookEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ResetTriggerConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetTriggerConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

