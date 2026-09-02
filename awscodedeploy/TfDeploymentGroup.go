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
type TfDeploymentGroup interface {
	cdktn.TerraformResource
	// Experimental.
	AlarmConfiguration() TfDeploymentGroup_AlarmConfigurationPropertyOutputReference
	// Experimental.
	AlarmConfigurationInput() *TfDeploymentGroup_AlarmConfigurationProperty
	// Experimental.
	AppName() *string
	// Experimental.
	SetAppName(val *string)
	// Experimental.
	AppNameInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	AutoRollbackConfiguration() TfDeploymentGroup_AutoRollbackConfigurationPropertyOutputReference
	// Experimental.
	AutoRollbackConfigurationInput() *TfDeploymentGroup_AutoRollbackConfigurationProperty
	// Experimental.
	AutoscalingGroups() *[]*string
	// Experimental.
	SetAutoscalingGroups(val *[]*string)
	// Experimental.
	AutoscalingGroupsInput() *[]*string
	// Experimental.
	BlueGreenDeploymentConfig() TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference
	// Experimental.
	BlueGreenDeploymentConfigInput() *TfDeploymentGroup_BlueGreenDeploymentConfigProperty
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
	DeploymentStyle() TfDeploymentGroup_DeploymentStylePropertyOutputReference
	// Experimental.
	DeploymentStyleInput() *TfDeploymentGroup_DeploymentStyleProperty
	// Experimental.
	Ec2TagFilter() TfDeploymentGroup_Ec2TagFilterPropertyList
	// Experimental.
	Ec2TagFilterInput() interface{}
	// Experimental.
	Ec2TagSet() TfDeploymentGroup_Ec2TagSetPropertyList
	// Experimental.
	Ec2TagSetInput() interface{}
	// Experimental.
	EcsService() TfDeploymentGroup_EcsServicePropertyOutputReference
	// Experimental.
	EcsServiceInput() *TfDeploymentGroup_EcsServiceProperty
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
	LoadBalancerInfo() TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference
	// Experimental.
	LoadBalancerInfoInput() *TfDeploymentGroup_LoadBalancerInfoProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OnPremisesInstanceTagFilter() TfDeploymentGroup_OnPremisesInstanceTagFilterPropertyList
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
	TriggerConfiguration() TfDeploymentGroup_TriggerConfigurationPropertyList
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
	PutAlarmConfiguration(value *TfDeploymentGroup_AlarmConfigurationProperty)
	// Experimental.
	PutAutoRollbackConfiguration(value *TfDeploymentGroup_AutoRollbackConfigurationProperty)
	// Experimental.
	PutBlueGreenDeploymentConfig(value *TfDeploymentGroup_BlueGreenDeploymentConfigProperty)
	// Experimental.
	PutDeploymentStyle(value *TfDeploymentGroup_DeploymentStyleProperty)
	// Experimental.
	PutEc2TagFilter(value interface{})
	// Experimental.
	PutEc2TagSet(value interface{})
	// Experimental.
	PutEcsService(value *TfDeploymentGroup_EcsServiceProperty)
	// Experimental.
	PutLoadBalancerInfo(value *TfDeploymentGroup_LoadBalancerInfoProperty)
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

// The jsii proxy struct for TfDeploymentGroup
type jsiiProxy_TfDeploymentGroup struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfDeploymentGroup) AlarmConfiguration() TfDeploymentGroup_AlarmConfigurationPropertyOutputReference {
	var returns TfDeploymentGroup_AlarmConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"alarmConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) AlarmConfigurationInput() *TfDeploymentGroup_AlarmConfigurationProperty {
	var returns *TfDeploymentGroup_AlarmConfigurationProperty
	_jsii_.Get(
		j,
		"alarmConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) AppName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) AppNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) AutoRollbackConfiguration() TfDeploymentGroup_AutoRollbackConfigurationPropertyOutputReference {
	var returns TfDeploymentGroup_AutoRollbackConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"autoRollbackConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) AutoRollbackConfigurationInput() *TfDeploymentGroup_AutoRollbackConfigurationProperty {
	var returns *TfDeploymentGroup_AutoRollbackConfigurationProperty
	_jsii_.Get(
		j,
		"autoRollbackConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) AutoscalingGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoscalingGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) AutoscalingGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoscalingGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) BlueGreenDeploymentConfig() TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference {
	var returns TfDeploymentGroup_BlueGreenDeploymentConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"blueGreenDeploymentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) BlueGreenDeploymentConfigInput() *TfDeploymentGroup_BlueGreenDeploymentConfigProperty {
	var returns *TfDeploymentGroup_BlueGreenDeploymentConfigProperty
	_jsii_.Get(
		j,
		"blueGreenDeploymentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) ComputePlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) DeploymentConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) DeploymentConfigNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentConfigNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) DeploymentGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) DeploymentGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) DeploymentGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) DeploymentStyle() TfDeploymentGroup_DeploymentStylePropertyOutputReference {
	var returns TfDeploymentGroup_DeploymentStylePropertyOutputReference
	_jsii_.Get(
		j,
		"deploymentStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) DeploymentStyleInput() *TfDeploymentGroup_DeploymentStyleProperty {
	var returns *TfDeploymentGroup_DeploymentStyleProperty
	_jsii_.Get(
		j,
		"deploymentStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) Ec2TagFilter() TfDeploymentGroup_Ec2TagFilterPropertyList {
	var returns TfDeploymentGroup_Ec2TagFilterPropertyList
	_jsii_.Get(
		j,
		"ec2TagFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) Ec2TagFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2TagFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) Ec2TagSet() TfDeploymentGroup_Ec2TagSetPropertyList {
	var returns TfDeploymentGroup_Ec2TagSetPropertyList
	_jsii_.Get(
		j,
		"ec2TagSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) Ec2TagSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ec2TagSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) EcsService() TfDeploymentGroup_EcsServicePropertyOutputReference {
	var returns TfDeploymentGroup_EcsServicePropertyOutputReference
	_jsii_.Get(
		j,
		"ecsService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) EcsServiceInput() *TfDeploymentGroup_EcsServiceProperty {
	var returns *TfDeploymentGroup_EcsServiceProperty
	_jsii_.Get(
		j,
		"ecsServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) LoadBalancerInfo() TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference {
	var returns TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"loadBalancerInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) LoadBalancerInfoInput() *TfDeploymentGroup_LoadBalancerInfoProperty {
	var returns *TfDeploymentGroup_LoadBalancerInfoProperty
	_jsii_.Get(
		j,
		"loadBalancerInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) OnPremisesInstanceTagFilter() TfDeploymentGroup_OnPremisesInstanceTagFilterPropertyList {
	var returns TfDeploymentGroup_OnPremisesInstanceTagFilterPropertyList
	_jsii_.Get(
		j,
		"onPremisesInstanceTagFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) OnPremisesInstanceTagFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onPremisesInstanceTagFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) OutdatedInstancesStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdatedInstancesStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) OutdatedInstancesStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outdatedInstancesStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) ServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) TerminationHookEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationHookEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) TerminationHookEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminationHookEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) TriggerConfiguration() TfDeploymentGroup_TriggerConfigurationPropertyList {
	var returns TfDeploymentGroup_TriggerConfigurationPropertyList
	_jsii_.Get(
		j,
		"triggerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup) TriggerConfigurationInput() interface{} {
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
func NewTfDeploymentGroup(scope constructs.Construct, id *string, config *TfDeploymentGroupConfig) TfDeploymentGroup {
	_init_.Initialize()

	if err := validateNewTfDeploymentGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDeploymentGroup{}

	_jsii_.Create(
		"@cdktn/aws-codedeploy.TfDeploymentGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group aws_codedeploy_deployment_group} Resource.
// Experimental.
func NewTfDeploymentGroup_Override(t TfDeploymentGroup, scope constructs.Construct, id *string, config *TfDeploymentGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codedeploy.TfDeploymentGroup",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetAppName(val *string) {
	if err := j.validateSetAppNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appName",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetAutoscalingGroups(val *[]*string) {
	if err := j.validateSetAutoscalingGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingGroups",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetDeploymentConfigName(val *string) {
	if err := j.validateSetDeploymentConfigNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentConfigName",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetDeploymentGroupName(val *string) {
	if err := j.validateSetDeploymentGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentGroupName",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetOutdatedInstancesStrategy(val *string) {
	if err := j.validateSetOutdatedInstancesStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outdatedInstancesStrategy",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetServiceRoleArn(val *string) {
	if err := j.validateSetServiceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup)SetTerminationHookEnabled(val interface{}) {
	if err := j.validateSetTerminationHookEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationHookEnabled",
		val,
	)
}

// Generates CDKTN code for importing a TfDeploymentGroup resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfDeploymentGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfDeploymentGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-codedeploy.TfDeploymentGroup",
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
func TfDeploymentGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDeploymentGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-codedeploy.TfDeploymentGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfDeploymentGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDeploymentGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-codedeploy.TfDeploymentGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfDeploymentGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDeploymentGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-codedeploy.TfDeploymentGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfDeploymentGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-codedeploy.TfDeploymentGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfDeploymentGroup) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDeploymentGroup) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeploymentGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDeploymentGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDeploymentGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDeploymentGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDeploymentGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDeploymentGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDeploymentGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDeploymentGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeploymentGroup) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeploymentGroup) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := t.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeploymentGroup) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) PutAlarmConfiguration(value *TfDeploymentGroup_AlarmConfigurationProperty) {
	if err := t.validatePutAlarmConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAlarmConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) PutAutoRollbackConfiguration(value *TfDeploymentGroup_AutoRollbackConfigurationProperty) {
	if err := t.validatePutAutoRollbackConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAutoRollbackConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) PutBlueGreenDeploymentConfig(value *TfDeploymentGroup_BlueGreenDeploymentConfigProperty) {
	if err := t.validatePutBlueGreenDeploymentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBlueGreenDeploymentConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) PutDeploymentStyle(value *TfDeploymentGroup_DeploymentStyleProperty) {
	if err := t.validatePutDeploymentStyleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeploymentStyle",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) PutEc2TagFilter(value interface{}) {
	if err := t.validatePutEc2TagFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEc2TagFilter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) PutEc2TagSet(value interface{}) {
	if err := t.validatePutEc2TagSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEc2TagSet",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) PutEcsService(value *TfDeploymentGroup_EcsServiceProperty) {
	if err := t.validatePutEcsServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEcsService",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) PutLoadBalancerInfo(value *TfDeploymentGroup_LoadBalancerInfoProperty) {
	if err := t.validatePutLoadBalancerInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLoadBalancerInfo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) PutOnPremisesInstanceTagFilter(value interface{}) {
	if err := t.validatePutOnPremisesInstanceTagFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOnPremisesInstanceTagFilter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) PutTriggerConfiguration(value interface{}) {
	if err := t.validatePutTriggerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTriggerConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetAlarmConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetAlarmConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetAutoRollbackConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoRollbackConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetAutoscalingGroups() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoscalingGroups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetBlueGreenDeploymentConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetBlueGreenDeploymentConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetDeploymentConfigName() {
	_jsii_.InvokeVoid(
		t,
		"resetDeploymentConfigName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetDeploymentStyle() {
	_jsii_.InvokeVoid(
		t,
		"resetDeploymentStyle",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetEc2TagFilter() {
	_jsii_.InvokeVoid(
		t,
		"resetEc2TagFilter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetEc2TagSet() {
	_jsii_.InvokeVoid(
		t,
		"resetEc2TagSet",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetEcsService() {
	_jsii_.InvokeVoid(
		t,
		"resetEcsService",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetLoadBalancerInfo() {
	_jsii_.InvokeVoid(
		t,
		"resetLoadBalancerInfo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetOnPremisesInstanceTagFilter() {
	_jsii_.InvokeVoid(
		t,
		"resetOnPremisesInstanceTagFilter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetOutdatedInstancesStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetOutdatedInstancesStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetTerminationHookEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetTerminationHookEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) ResetTriggerConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetTriggerConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeploymentGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeploymentGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeploymentGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeploymentGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeploymentGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeploymentGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		t,
		"with",
		args,
		&returns,
	)

	return returns
}

