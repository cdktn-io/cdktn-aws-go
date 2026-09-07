package ecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ecs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/ecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service aws_ecs_service}.
// Experimental.
type AwsService interface {
	cdktn.TerraformResource
	// Experimental.
	Alarms() AwsService_AlarmsPropertyOutputReference
	// Experimental.
	AlarmsInput() *AwsService_AlarmsProperty
	// Experimental.
	Arn() *string
	// Experimental.
	AvailabilityZoneRebalancing() *string
	// Experimental.
	SetAvailabilityZoneRebalancing(val *string)
	// Experimental.
	AvailabilityZoneRebalancingInput() *string
	// Experimental.
	CapacityProviderStrategy() AwsService_CapacityProviderStrategyPropertyList
	// Experimental.
	CapacityProviderStrategyInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Cluster() *string
	// Experimental.
	SetCluster(val *string)
	// Experimental.
	ClusterInput() *string
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
	DeploymentCircuitBreaker() AwsService_DeploymentCircuitBreakerPropertyOutputReference
	// Experimental.
	DeploymentCircuitBreakerInput() *AwsService_DeploymentCircuitBreakerProperty
	// Experimental.
	DeploymentConfiguration() AwsService_DeploymentConfigurationPropertyOutputReference
	// Experimental.
	DeploymentConfigurationInput() *AwsService_DeploymentConfigurationProperty
	// Experimental.
	DeploymentController() AwsService_DeploymentControllerPropertyOutputReference
	// Experimental.
	DeploymentControllerInput() *AwsService_DeploymentControllerProperty
	// Experimental.
	DeploymentMaximumPercent() *float64
	// Experimental.
	SetDeploymentMaximumPercent(val *float64)
	// Experimental.
	DeploymentMaximumPercentInput() *float64
	// Experimental.
	DeploymentMinimumHealthyPercent() *float64
	// Experimental.
	SetDeploymentMinimumHealthyPercent(val *float64)
	// Experimental.
	DeploymentMinimumHealthyPercentInput() *float64
	// Experimental.
	DesiredCount() *float64
	// Experimental.
	SetDesiredCount(val *float64)
	// Experimental.
	DesiredCountInput() *float64
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
	ForceDelete() interface{}
	// Experimental.
	SetForceDelete(val interface{})
	// Experimental.
	ForceDeleteInput() interface{}
	// Experimental.
	ForceNewDeployment() interface{}
	// Experimental.
	SetForceNewDeployment(val interface{})
	// Experimental.
	ForceNewDeploymentInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HealthCheckGracePeriodSeconds() *float64
	// Experimental.
	SetHealthCheckGracePeriodSeconds(val *float64)
	// Experimental.
	HealthCheckGracePeriodSecondsInput() *float64
	// Experimental.
	IamRole() *string
	// Experimental.
	SetIamRole(val *string)
	// Experimental.
	IamRoleInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	LaunchType() *string
	// Experimental.
	SetLaunchType(val *string)
	// Experimental.
	LaunchTypeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LoadBalancer() AwsService_LoadBalancerPropertyList
	// Experimental.
	LoadBalancerInput() interface{}
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NetworkConfiguration() AwsService_NetworkConfigurationPropertyOutputReference
	// Experimental.
	NetworkConfigurationInput() *AwsService_NetworkConfigurationProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OrderedPlacementStrategy() AwsService_OrderedPlacementStrategyPropertyList
	// Experimental.
	OrderedPlacementStrategyInput() interface{}
	// Experimental.
	PlacementConstraints() AwsService_PlacementConstraintsPropertyList
	// Experimental.
	PlacementConstraintsInput() interface{}
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
	SchedulingStrategy() *string
	// Experimental.
	SetSchedulingStrategy(val *string)
	// Experimental.
	SchedulingStrategyInput() *string
	// Experimental.
	ServiceConnectConfiguration() AwsService_ServiceConnectConfigurationPropertyOutputReference
	// Experimental.
	ServiceConnectConfigurationInput() *AwsService_ServiceConnectConfigurationProperty
	// Experimental.
	ServiceRegistries() AwsService_ServiceRegistriesPropertyOutputReference
	// Experimental.
	ServiceRegistriesInput() *AwsService_ServiceRegistriesProperty
	// Experimental.
	SigintRollback() interface{}
	// Experimental.
	SetSigintRollback(val interface{})
	// Experimental.
	SigintRollbackInput() interface{}
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
	TaskDefinition() *string
	// Experimental.
	SetTaskDefinition(val *string)
	// Experimental.
	TaskDefinitionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsService_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Triggers() *map[string]*string
	// Experimental.
	SetTriggers(val *map[string]*string)
	// Experimental.
	TriggersInput() *map[string]*string
	// Experimental.
	VolumeConfiguration() AwsService_VolumeConfigurationPropertyOutputReference
	// Experimental.
	VolumeConfigurationInput() *AwsService_VolumeConfigurationProperty
	// Experimental.
	VpcLatticeConfigurations() AwsService_VpcLatticeConfigurationsPropertyList
	// Experimental.
	VpcLatticeConfigurationsInput() interface{}
	// Experimental.
	WaitForSteadyState() interface{}
	// Experimental.
	SetWaitForSteadyState(val interface{})
	// Experimental.
	WaitForSteadyStateInput() interface{}
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
	PutAlarms(value *AwsService_AlarmsProperty)
	// Experimental.
	PutCapacityProviderStrategy(value interface{})
	// Experimental.
	PutDeploymentCircuitBreaker(value *AwsService_DeploymentCircuitBreakerProperty)
	// Experimental.
	PutDeploymentConfiguration(value *AwsService_DeploymentConfigurationProperty)
	// Experimental.
	PutDeploymentController(value *AwsService_DeploymentControllerProperty)
	// Experimental.
	PutLoadBalancer(value interface{})
	// Experimental.
	PutNetworkConfiguration(value *AwsService_NetworkConfigurationProperty)
	// Experimental.
	PutOrderedPlacementStrategy(value interface{})
	// Experimental.
	PutPlacementConstraints(value interface{})
	// Experimental.
	PutServiceConnectConfiguration(value *AwsService_ServiceConnectConfigurationProperty)
	// Experimental.
	PutServiceRegistries(value *AwsService_ServiceRegistriesProperty)
	// Experimental.
	PutTimeouts(value *AwsService_TimeoutsProperty)
	// Experimental.
	PutVolumeConfiguration(value *AwsService_VolumeConfigurationProperty)
	// Experimental.
	PutVpcLatticeConfigurations(value interface{})
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
	ResetAlarms()
	// Experimental.
	ResetAvailabilityZoneRebalancing()
	// Experimental.
	ResetCapacityProviderStrategy()
	// Experimental.
	ResetCluster()
	// Experimental.
	ResetDeploymentCircuitBreaker()
	// Experimental.
	ResetDeploymentConfiguration()
	// Experimental.
	ResetDeploymentController()
	// Experimental.
	ResetDeploymentMaximumPercent()
	// Experimental.
	ResetDeploymentMinimumHealthyPercent()
	// Experimental.
	ResetDesiredCount()
	// Experimental.
	ResetEnableEcsManagedTags()
	// Experimental.
	ResetEnableExecuteCommand()
	// Experimental.
	ResetForceDelete()
	// Experimental.
	ResetForceNewDeployment()
	// Experimental.
	ResetHealthCheckGracePeriodSeconds()
	// Experimental.
	ResetIamRole()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLaunchType()
	// Experimental.
	ResetLoadBalancer()
	// Experimental.
	ResetNetworkConfiguration()
	// Experimental.
	ResetOrderedPlacementStrategy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPlacementConstraints()
	// Experimental.
	ResetPlatformVersion()
	// Experimental.
	ResetPropagateTags()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSchedulingStrategy()
	// Experimental.
	ResetServiceConnectConfiguration()
	// Experimental.
	ResetServiceRegistries()
	// Experimental.
	ResetSigintRollback()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTaskDefinition()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTriggers()
	// Experimental.
	ResetVolumeConfiguration()
	// Experimental.
	ResetVpcLatticeConfigurations()
	// Experimental.
	ResetWaitForSteadyState()
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

// The jsii proxy struct for AwsService
type jsiiProxy_AwsService struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsService) Alarms() AwsService_AlarmsPropertyOutputReference {
	var returns AwsService_AlarmsPropertyOutputReference
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) AlarmsInput() *AwsService_AlarmsProperty {
	var returns *AwsService_AlarmsProperty
	_jsii_.Get(
		j,
		"alarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) AvailabilityZoneRebalancing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneRebalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) AvailabilityZoneRebalancingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneRebalancingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) CapacityProviderStrategy() AwsService_CapacityProviderStrategyPropertyList {
	var returns AwsService_CapacityProviderStrategyPropertyList
	_jsii_.Get(
		j,
		"capacityProviderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) CapacityProviderStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityProviderStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) ClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) DeploymentCircuitBreaker() AwsService_DeploymentCircuitBreakerPropertyOutputReference {
	var returns AwsService_DeploymentCircuitBreakerPropertyOutputReference
	_jsii_.Get(
		j,
		"deploymentCircuitBreaker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) DeploymentCircuitBreakerInput() *AwsService_DeploymentCircuitBreakerProperty {
	var returns *AwsService_DeploymentCircuitBreakerProperty
	_jsii_.Get(
		j,
		"deploymentCircuitBreakerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) DeploymentConfiguration() AwsService_DeploymentConfigurationPropertyOutputReference {
	var returns AwsService_DeploymentConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"deploymentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) DeploymentConfigurationInput() *AwsService_DeploymentConfigurationProperty {
	var returns *AwsService_DeploymentConfigurationProperty
	_jsii_.Get(
		j,
		"deploymentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) DeploymentController() AwsService_DeploymentControllerPropertyOutputReference {
	var returns AwsService_DeploymentControllerPropertyOutputReference
	_jsii_.Get(
		j,
		"deploymentController",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) DeploymentControllerInput() *AwsService_DeploymentControllerProperty {
	var returns *AwsService_DeploymentControllerProperty
	_jsii_.Get(
		j,
		"deploymentControllerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) DeploymentMaximumPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentMaximumPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) DeploymentMaximumPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentMaximumPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) DeploymentMinimumHealthyPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentMinimumHealthyPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) DeploymentMinimumHealthyPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentMinimumHealthyPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) DesiredCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) EnableEcsManagedTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) EnableEcsManagedTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) EnableExecuteCommand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) EnableExecuteCommandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) ForceDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) ForceDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) ForceNewDeployment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceNewDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) ForceNewDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceNewDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) HealthCheckGracePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) HealthCheckGracePeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) IamRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) IamRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) LaunchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) LaunchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) LoadBalancer() AwsService_LoadBalancerPropertyList {
	var returns AwsService_LoadBalancerPropertyList
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) LoadBalancerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) NetworkConfiguration() AwsService_NetworkConfigurationPropertyOutputReference {
	var returns AwsService_NetworkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) NetworkConfigurationInput() *AwsService_NetworkConfigurationProperty {
	var returns *AwsService_NetworkConfigurationProperty
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) OrderedPlacementStrategy() AwsService_OrderedPlacementStrategyPropertyList {
	var returns AwsService_OrderedPlacementStrategyPropertyList
	_jsii_.Get(
		j,
		"orderedPlacementStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) OrderedPlacementStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orderedPlacementStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) PlacementConstraints() AwsService_PlacementConstraintsPropertyList {
	var returns AwsService_PlacementConstraintsPropertyList
	_jsii_.Get(
		j,
		"placementConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) PlacementConstraintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) PlatformVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) PropagateTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) SchedulingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) SchedulingStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) ServiceConnectConfiguration() AwsService_ServiceConnectConfigurationPropertyOutputReference {
	var returns AwsService_ServiceConnectConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceConnectConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) ServiceConnectConfigurationInput() *AwsService_ServiceConnectConfigurationProperty {
	var returns *AwsService_ServiceConnectConfigurationProperty
	_jsii_.Get(
		j,
		"serviceConnectConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) ServiceRegistries() AwsService_ServiceRegistriesPropertyOutputReference {
	var returns AwsService_ServiceRegistriesPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceRegistries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) ServiceRegistriesInput() *AwsService_ServiceRegistriesProperty {
	var returns *AwsService_ServiceRegistriesProperty
	_jsii_.Get(
		j,
		"serviceRegistriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) SigintRollback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sigintRollback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) SigintRollbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sigintRollbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) TaskDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) TaskDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) Timeouts() AwsService_TimeoutsPropertyOutputReference {
	var returns AwsService_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) Triggers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) TriggersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) VolumeConfiguration() AwsService_VolumeConfigurationPropertyOutputReference {
	var returns AwsService_VolumeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"volumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) VolumeConfigurationInput() *AwsService_VolumeConfigurationProperty {
	var returns *AwsService_VolumeConfigurationProperty
	_jsii_.Get(
		j,
		"volumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) VpcLatticeConfigurations() AwsService_VpcLatticeConfigurationsPropertyList {
	var returns AwsService_VpcLatticeConfigurationsPropertyList
	_jsii_.Get(
		j,
		"vpcLatticeConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) VpcLatticeConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcLatticeConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) WaitForSteadyState() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForSteadyState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService) WaitForSteadyStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForSteadyStateInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service aws_ecs_service} Resource.
// Experimental.
func NewAwsService(scope constructs.Construct, id *string, config *AwsServiceConfig) AwsService {
	_init_.Initialize()

	if err := validateNewAwsServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsService{}

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service aws_ecs_service} Resource.
// Experimental.
func NewAwsService_Override(a AwsService, scope constructs.Construct, id *string, config *AwsServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsService",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsService)SetAvailabilityZoneRebalancing(val *string) {
	if err := j.validateSetAvailabilityZoneRebalancingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneRebalancing",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetCluster(val *string) {
	if err := j.validateSetClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetDeploymentMaximumPercent(val *float64) {
	if err := j.validateSetDeploymentMaximumPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentMaximumPercent",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetDeploymentMinimumHealthyPercent(val *float64) {
	if err := j.validateSetDeploymentMinimumHealthyPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentMinimumHealthyPercent",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetDesiredCount(val *float64) {
	if err := j.validateSetDesiredCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCount",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetEnableEcsManagedTags(val interface{}) {
	if err := j.validateSetEnableEcsManagedTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEcsManagedTags",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetEnableExecuteCommand(val interface{}) {
	if err := j.validateSetEnableExecuteCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableExecuteCommand",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetForceDelete(val interface{}) {
	if err := j.validateSetForceDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDelete",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetForceNewDeployment(val interface{}) {
	if err := j.validateSetForceNewDeploymentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceNewDeployment",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetHealthCheckGracePeriodSeconds(val *float64) {
	if err := j.validateSetHealthCheckGracePeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckGracePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetIamRole(val *string) {
	if err := j.validateSetIamRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRole",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetLaunchType(val *string) {
	if err := j.validateSetLaunchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchType",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetPlatformVersion(val *string) {
	if err := j.validateSetPlatformVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformVersion",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetPropagateTags(val *string) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetSchedulingStrategy(val *string) {
	if err := j.validateSetSchedulingStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulingStrategy",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetSigintRollback(val interface{}) {
	if err := j.validateSetSigintRollbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sigintRollback",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetTaskDefinition(val *string) {
	if err := j.validateSetTaskDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDefinition",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetTriggers(val *map[string]*string) {
	if err := j.validateSetTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggers",
		val,
	)
}

func (j *jsiiProxy_AwsService)SetWaitForSteadyState(val interface{}) {
	if err := j.validateSetWaitForSteadyStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForSteadyState",
		val,
	)
}

// Generates CDKTN code for importing a AwsService resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ecs.AwsService",
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
func AwsService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ecs.AwsService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ecs.AwsService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ecs.AwsService",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ecs.AwsService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsService) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsService) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsService) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsService) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsService) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsService) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsService) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsService) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsService) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsService) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsService) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsService) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsService) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsService) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsService) PutAlarms(value *AwsService_AlarmsProperty) {
	if err := a.validatePutAlarmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAlarms",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsService) PutCapacityProviderStrategy(value interface{}) {
	if err := a.validatePutCapacityProviderStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCapacityProviderStrategy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsService) PutDeploymentCircuitBreaker(value *AwsService_DeploymentCircuitBreakerProperty) {
	if err := a.validatePutDeploymentCircuitBreakerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeploymentCircuitBreaker",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsService) PutDeploymentConfiguration(value *AwsService_DeploymentConfigurationProperty) {
	if err := a.validatePutDeploymentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeploymentConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsService) PutDeploymentController(value *AwsService_DeploymentControllerProperty) {
	if err := a.validatePutDeploymentControllerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeploymentController",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsService) PutLoadBalancer(value interface{}) {
	if err := a.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsService) PutNetworkConfiguration(value *AwsService_NetworkConfigurationProperty) {
	if err := a.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsService) PutOrderedPlacementStrategy(value interface{}) {
	if err := a.validatePutOrderedPlacementStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOrderedPlacementStrategy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsService) PutPlacementConstraints(value interface{}) {
	if err := a.validatePutPlacementConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPlacementConstraints",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsService) PutServiceConnectConfiguration(value *AwsService_ServiceConnectConfigurationProperty) {
	if err := a.validatePutServiceConnectConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceConnectConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsService) PutServiceRegistries(value *AwsService_ServiceRegistriesProperty) {
	if err := a.validatePutServiceRegistriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceRegistries",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsService) PutTimeouts(value *AwsService_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsService) PutVolumeConfiguration(value *AwsService_VolumeConfigurationProperty) {
	if err := a.validatePutVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVolumeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsService) PutVpcLatticeConfigurations(value interface{}) {
	if err := a.validatePutVpcLatticeConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcLatticeConfigurations",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsService) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsService) ResetAlarms() {
	_jsii_.InvokeVoid(
		a,
		"resetAlarms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetAvailabilityZoneRebalancing() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZoneRebalancing",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetCapacityProviderStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacityProviderStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetCluster() {
	_jsii_.InvokeVoid(
		a,
		"resetCluster",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetDeploymentCircuitBreaker() {
	_jsii_.InvokeVoid(
		a,
		"resetDeploymentCircuitBreaker",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetDeploymentConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetDeploymentConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetDeploymentController() {
	_jsii_.InvokeVoid(
		a,
		"resetDeploymentController",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetDeploymentMaximumPercent() {
	_jsii_.InvokeVoid(
		a,
		"resetDeploymentMaximumPercent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetDeploymentMinimumHealthyPercent() {
	_jsii_.InvokeVoid(
		a,
		"resetDeploymentMinimumHealthyPercent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetDesiredCount() {
	_jsii_.InvokeVoid(
		a,
		"resetDesiredCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetEnableEcsManagedTags() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableEcsManagedTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetEnableExecuteCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableExecuteCommand",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetForceDelete() {
	_jsii_.InvokeVoid(
		a,
		"resetForceDelete",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetForceNewDeployment() {
	_jsii_.InvokeVoid(
		a,
		"resetForceNewDeployment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetHealthCheckGracePeriodSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheckGracePeriodSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetIamRole() {
	_jsii_.InvokeVoid(
		a,
		"resetIamRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetLaunchType() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetLoadBalancer() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetOrderedPlacementStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetOrderedPlacementStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetPlacementConstraints() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacementConstraints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetPlatformVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetPlatformVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		a,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetSchedulingStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetSchedulingStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetServiceConnectConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceConnectConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetServiceRegistries() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceRegistries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetSigintRollback() {
	_jsii_.InvokeVoid(
		a,
		"resetSigintRollback",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetTaskDefinition() {
	_jsii_.InvokeVoid(
		a,
		"resetTaskDefinition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetTriggers() {
	_jsii_.InvokeVoid(
		a,
		"resetTriggers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetVolumeConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetVolumeConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetVpcLatticeConfigurations() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcLatticeConfigurations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) ResetWaitForSteadyState() {
	_jsii_.InvokeVoid(
		a,
		"resetWaitForSteadyState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsService) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsService) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsService) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

