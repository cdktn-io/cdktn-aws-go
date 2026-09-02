package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service aws_ecs_service}.
// Experimental.
type TfService interface {
	cdktn.TerraformResource
	// Experimental.
	Alarms() TfService_AlarmsPropertyOutputReference
	// Experimental.
	AlarmsInput() *TfService_AlarmsProperty
	// Experimental.
	Arn() *string
	// Experimental.
	AvailabilityZoneRebalancing() *string
	// Experimental.
	SetAvailabilityZoneRebalancing(val *string)
	// Experimental.
	AvailabilityZoneRebalancingInput() *string
	// Experimental.
	CapacityProviderStrategy() TfService_CapacityProviderStrategyPropertyList
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
	DeploymentCircuitBreaker() TfService_DeploymentCircuitBreakerPropertyOutputReference
	// Experimental.
	DeploymentCircuitBreakerInput() *TfService_DeploymentCircuitBreakerProperty
	// Experimental.
	DeploymentConfiguration() TfService_DeploymentConfigurationPropertyOutputReference
	// Experimental.
	DeploymentConfigurationInput() *TfService_DeploymentConfigurationProperty
	// Experimental.
	DeploymentController() TfService_DeploymentControllerPropertyOutputReference
	// Experimental.
	DeploymentControllerInput() *TfService_DeploymentControllerProperty
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
	LoadBalancer() TfService_LoadBalancerPropertyList
	// Experimental.
	LoadBalancerInput() interface{}
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NetworkConfiguration() TfService_NetworkConfigurationPropertyOutputReference
	// Experimental.
	NetworkConfigurationInput() *TfService_NetworkConfigurationProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OrderedPlacementStrategy() TfService_OrderedPlacementStrategyPropertyList
	// Experimental.
	OrderedPlacementStrategyInput() interface{}
	// Experimental.
	PlacementConstraints() TfService_PlacementConstraintsPropertyList
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
	ServiceConnectConfiguration() TfService_ServiceConnectConfigurationPropertyOutputReference
	// Experimental.
	ServiceConnectConfigurationInput() *TfService_ServiceConnectConfigurationProperty
	// Experimental.
	ServiceRegistries() TfService_ServiceRegistriesPropertyOutputReference
	// Experimental.
	ServiceRegistriesInput() *TfService_ServiceRegistriesProperty
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
	Timeouts() TfService_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Triggers() *map[string]*string
	// Experimental.
	SetTriggers(val *map[string]*string)
	// Experimental.
	TriggersInput() *map[string]*string
	// Experimental.
	VolumeConfiguration() TfService_VolumeConfigurationPropertyOutputReference
	// Experimental.
	VolumeConfigurationInput() *TfService_VolumeConfigurationProperty
	// Experimental.
	VpcLatticeConfigurations() TfService_VpcLatticeConfigurationsPropertyList
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
	PutAlarms(value *TfService_AlarmsProperty)
	// Experimental.
	PutCapacityProviderStrategy(value interface{})
	// Experimental.
	PutDeploymentCircuitBreaker(value *TfService_DeploymentCircuitBreakerProperty)
	// Experimental.
	PutDeploymentConfiguration(value *TfService_DeploymentConfigurationProperty)
	// Experimental.
	PutDeploymentController(value *TfService_DeploymentControllerProperty)
	// Experimental.
	PutLoadBalancer(value interface{})
	// Experimental.
	PutNetworkConfiguration(value *TfService_NetworkConfigurationProperty)
	// Experimental.
	PutOrderedPlacementStrategy(value interface{})
	// Experimental.
	PutPlacementConstraints(value interface{})
	// Experimental.
	PutServiceConnectConfiguration(value *TfService_ServiceConnectConfigurationProperty)
	// Experimental.
	PutServiceRegistries(value *TfService_ServiceRegistriesProperty)
	// Experimental.
	PutTimeouts(value *TfService_TimeoutsProperty)
	// Experimental.
	PutVolumeConfiguration(value *TfService_VolumeConfigurationProperty)
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

// The jsii proxy struct for TfService
type jsiiProxy_TfService struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfService) Alarms() TfService_AlarmsPropertyOutputReference {
	var returns TfService_AlarmsPropertyOutputReference
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) AlarmsInput() *TfService_AlarmsProperty {
	var returns *TfService_AlarmsProperty
	_jsii_.Get(
		j,
		"alarmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) AvailabilityZoneRebalancing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneRebalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) AvailabilityZoneRebalancingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneRebalancingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) CapacityProviderStrategy() TfService_CapacityProviderStrategyPropertyList {
	var returns TfService_CapacityProviderStrategyPropertyList
	_jsii_.Get(
		j,
		"capacityProviderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) CapacityProviderStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityProviderStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) ClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) DeploymentCircuitBreaker() TfService_DeploymentCircuitBreakerPropertyOutputReference {
	var returns TfService_DeploymentCircuitBreakerPropertyOutputReference
	_jsii_.Get(
		j,
		"deploymentCircuitBreaker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) DeploymentCircuitBreakerInput() *TfService_DeploymentCircuitBreakerProperty {
	var returns *TfService_DeploymentCircuitBreakerProperty
	_jsii_.Get(
		j,
		"deploymentCircuitBreakerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) DeploymentConfiguration() TfService_DeploymentConfigurationPropertyOutputReference {
	var returns TfService_DeploymentConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"deploymentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) DeploymentConfigurationInput() *TfService_DeploymentConfigurationProperty {
	var returns *TfService_DeploymentConfigurationProperty
	_jsii_.Get(
		j,
		"deploymentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) DeploymentController() TfService_DeploymentControllerPropertyOutputReference {
	var returns TfService_DeploymentControllerPropertyOutputReference
	_jsii_.Get(
		j,
		"deploymentController",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) DeploymentControllerInput() *TfService_DeploymentControllerProperty {
	var returns *TfService_DeploymentControllerProperty
	_jsii_.Get(
		j,
		"deploymentControllerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) DeploymentMaximumPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentMaximumPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) DeploymentMaximumPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentMaximumPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) DeploymentMinimumHealthyPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentMinimumHealthyPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) DeploymentMinimumHealthyPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deploymentMinimumHealthyPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) DesiredCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) DesiredCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) EnableEcsManagedTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) EnableEcsManagedTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) EnableExecuteCommand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) EnableExecuteCommandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) ForceDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) ForceDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) ForceNewDeployment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceNewDeployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) ForceNewDeploymentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceNewDeploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) HealthCheckGracePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) HealthCheckGracePeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) IamRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) IamRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) LaunchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) LaunchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) LoadBalancer() TfService_LoadBalancerPropertyList {
	var returns TfService_LoadBalancerPropertyList
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) LoadBalancerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) NetworkConfiguration() TfService_NetworkConfigurationPropertyOutputReference {
	var returns TfService_NetworkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) NetworkConfigurationInput() *TfService_NetworkConfigurationProperty {
	var returns *TfService_NetworkConfigurationProperty
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) OrderedPlacementStrategy() TfService_OrderedPlacementStrategyPropertyList {
	var returns TfService_OrderedPlacementStrategyPropertyList
	_jsii_.Get(
		j,
		"orderedPlacementStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) OrderedPlacementStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orderedPlacementStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) PlacementConstraints() TfService_PlacementConstraintsPropertyList {
	var returns TfService_PlacementConstraintsPropertyList
	_jsii_.Get(
		j,
		"placementConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) PlacementConstraintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) PlatformVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) PropagateTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) SchedulingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) SchedulingStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedulingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) ServiceConnectConfiguration() TfService_ServiceConnectConfigurationPropertyOutputReference {
	var returns TfService_ServiceConnectConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceConnectConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) ServiceConnectConfigurationInput() *TfService_ServiceConnectConfigurationProperty {
	var returns *TfService_ServiceConnectConfigurationProperty
	_jsii_.Get(
		j,
		"serviceConnectConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) ServiceRegistries() TfService_ServiceRegistriesPropertyOutputReference {
	var returns TfService_ServiceRegistriesPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceRegistries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) ServiceRegistriesInput() *TfService_ServiceRegistriesProperty {
	var returns *TfService_ServiceRegistriesProperty
	_jsii_.Get(
		j,
		"serviceRegistriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) SigintRollback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sigintRollback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) SigintRollbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sigintRollbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) TaskDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) TaskDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) Timeouts() TfService_TimeoutsPropertyOutputReference {
	var returns TfService_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) Triggers() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) TriggersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"triggersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) VolumeConfiguration() TfService_VolumeConfigurationPropertyOutputReference {
	var returns TfService_VolumeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"volumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) VolumeConfigurationInput() *TfService_VolumeConfigurationProperty {
	var returns *TfService_VolumeConfigurationProperty
	_jsii_.Get(
		j,
		"volumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) VpcLatticeConfigurations() TfService_VpcLatticeConfigurationsPropertyList {
	var returns TfService_VpcLatticeConfigurationsPropertyList
	_jsii_.Get(
		j,
		"vpcLatticeConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) VpcLatticeConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcLatticeConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) WaitForSteadyState() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForSteadyState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService) WaitForSteadyStateInput() interface{} {
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
func NewTfService(scope constructs.Construct, id *string, config *TfServiceConfig) TfService {
	_init_.Initialize()

	if err := validateNewTfServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfService{}

	_jsii_.Create(
		"@cdktn/aws-ecs.TfService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service aws_ecs_service} Resource.
// Experimental.
func NewTfService_Override(t TfService, scope constructs.Construct, id *string, config *TfServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.TfService",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfService)SetAvailabilityZoneRebalancing(val *string) {
	if err := j.validateSetAvailabilityZoneRebalancingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneRebalancing",
		val,
	)
}

func (j *jsiiProxy_TfService)SetCluster(val *string) {
	if err := j.validateSetClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_TfService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfService)SetDeploymentMaximumPercent(val *float64) {
	if err := j.validateSetDeploymentMaximumPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentMaximumPercent",
		val,
	)
}

func (j *jsiiProxy_TfService)SetDeploymentMinimumHealthyPercent(val *float64) {
	if err := j.validateSetDeploymentMinimumHealthyPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentMinimumHealthyPercent",
		val,
	)
}

func (j *jsiiProxy_TfService)SetDesiredCount(val *float64) {
	if err := j.validateSetDesiredCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCount",
		val,
	)
}

func (j *jsiiProxy_TfService)SetEnableEcsManagedTags(val interface{}) {
	if err := j.validateSetEnableEcsManagedTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEcsManagedTags",
		val,
	)
}

func (j *jsiiProxy_TfService)SetEnableExecuteCommand(val interface{}) {
	if err := j.validateSetEnableExecuteCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableExecuteCommand",
		val,
	)
}

func (j *jsiiProxy_TfService)SetForceDelete(val interface{}) {
	if err := j.validateSetForceDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDelete",
		val,
	)
}

func (j *jsiiProxy_TfService)SetForceNewDeployment(val interface{}) {
	if err := j.validateSetForceNewDeploymentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceNewDeployment",
		val,
	)
}

func (j *jsiiProxy_TfService)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfService)SetHealthCheckGracePeriodSeconds(val *float64) {
	if err := j.validateSetHealthCheckGracePeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckGracePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_TfService)SetIamRole(val *string) {
	if err := j.validateSetIamRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRole",
		val,
	)
}

func (j *jsiiProxy_TfService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfService)SetLaunchType(val *string) {
	if err := j.validateSetLaunchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchType",
		val,
	)
}

func (j *jsiiProxy_TfService)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfService)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfService)SetPlatformVersion(val *string) {
	if err := j.validateSetPlatformVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformVersion",
		val,
	)
}

func (j *jsiiProxy_TfService)SetPropagateTags(val *string) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_TfService)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfService)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfService)SetSchedulingStrategy(val *string) {
	if err := j.validateSetSchedulingStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulingStrategy",
		val,
	)
}

func (j *jsiiProxy_TfService)SetSigintRollback(val interface{}) {
	if err := j.validateSetSigintRollbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sigintRollback",
		val,
	)
}

func (j *jsiiProxy_TfService)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfService)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfService)SetTaskDefinition(val *string) {
	if err := j.validateSetTaskDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDefinition",
		val,
	)
}

func (j *jsiiProxy_TfService)SetTriggers(val *map[string]*string) {
	if err := j.validateSetTriggersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggers",
		val,
	)
}

func (j *jsiiProxy_TfService)SetWaitForSteadyState(val interface{}) {
	if err := j.validateSetWaitForSteadyStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForSteadyState",
		val,
	)
}

// Generates CDKTN code for importing a TfService resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ecs.TfService",
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
func TfService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ecs.TfService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ecs.TfService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ecs.TfService",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ecs.TfService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfService) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfService) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfService) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfService) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfService) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfService) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfService) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfService) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfService) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfService) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfService) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfService) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfService) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfService) PutAlarms(value *TfService_AlarmsProperty) {
	if err := t.validatePutAlarmsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAlarms",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService) PutCapacityProviderStrategy(value interface{}) {
	if err := t.validatePutCapacityProviderStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCapacityProviderStrategy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService) PutDeploymentCircuitBreaker(value *TfService_DeploymentCircuitBreakerProperty) {
	if err := t.validatePutDeploymentCircuitBreakerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeploymentCircuitBreaker",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService) PutDeploymentConfiguration(value *TfService_DeploymentConfigurationProperty) {
	if err := t.validatePutDeploymentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeploymentConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService) PutDeploymentController(value *TfService_DeploymentControllerProperty) {
	if err := t.validatePutDeploymentControllerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeploymentController",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService) PutLoadBalancer(value interface{}) {
	if err := t.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService) PutNetworkConfiguration(value *TfService_NetworkConfigurationProperty) {
	if err := t.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService) PutOrderedPlacementStrategy(value interface{}) {
	if err := t.validatePutOrderedPlacementStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOrderedPlacementStrategy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService) PutPlacementConstraints(value interface{}) {
	if err := t.validatePutPlacementConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPlacementConstraints",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService) PutServiceConnectConfiguration(value *TfService_ServiceConnectConfigurationProperty) {
	if err := t.validatePutServiceConnectConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putServiceConnectConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService) PutServiceRegistries(value *TfService_ServiceRegistriesProperty) {
	if err := t.validatePutServiceRegistriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putServiceRegistries",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService) PutTimeouts(value *TfService_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService) PutVolumeConfiguration(value *TfService_VolumeConfigurationProperty) {
	if err := t.validatePutVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVolumeConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService) PutVpcLatticeConfigurations(value interface{}) {
	if err := t.validatePutVpcLatticeConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVpcLatticeConfigurations",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfService) ResetAlarms() {
	_jsii_.InvokeVoid(
		t,
		"resetAlarms",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetAvailabilityZoneRebalancing() {
	_jsii_.InvokeVoid(
		t,
		"resetAvailabilityZoneRebalancing",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetCapacityProviderStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetCapacityProviderStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetCluster() {
	_jsii_.InvokeVoid(
		t,
		"resetCluster",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetDeploymentCircuitBreaker() {
	_jsii_.InvokeVoid(
		t,
		"resetDeploymentCircuitBreaker",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetDeploymentConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetDeploymentConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetDeploymentController() {
	_jsii_.InvokeVoid(
		t,
		"resetDeploymentController",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetDeploymentMaximumPercent() {
	_jsii_.InvokeVoid(
		t,
		"resetDeploymentMaximumPercent",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetDeploymentMinimumHealthyPercent() {
	_jsii_.InvokeVoid(
		t,
		"resetDeploymentMinimumHealthyPercent",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetDesiredCount() {
	_jsii_.InvokeVoid(
		t,
		"resetDesiredCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetEnableEcsManagedTags() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableEcsManagedTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetEnableExecuteCommand() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableExecuteCommand",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetForceDelete() {
	_jsii_.InvokeVoid(
		t,
		"resetForceDelete",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetForceNewDeployment() {
	_jsii_.InvokeVoid(
		t,
		"resetForceNewDeployment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetHealthCheckGracePeriodSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetHealthCheckGracePeriodSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetIamRole() {
	_jsii_.InvokeVoid(
		t,
		"resetIamRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetLaunchType() {
	_jsii_.InvokeVoid(
		t,
		"resetLaunchType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetLoadBalancer() {
	_jsii_.InvokeVoid(
		t,
		"resetLoadBalancer",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetOrderedPlacementStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetOrderedPlacementStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetPlacementConstraints() {
	_jsii_.InvokeVoid(
		t,
		"resetPlacementConstraints",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetPlatformVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetPlatformVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		t,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetSchedulingStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetSchedulingStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetServiceConnectConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceConnectConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetServiceRegistries() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceRegistries",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetSigintRollback() {
	_jsii_.InvokeVoid(
		t,
		"resetSigintRollback",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetTaskDefinition() {
	_jsii_.InvokeVoid(
		t,
		"resetTaskDefinition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetTriggers() {
	_jsii_.InvokeVoid(
		t,
		"resetTriggers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetVolumeConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetVolumeConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetVpcLatticeConfigurations() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcLatticeConfigurations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) ResetWaitForSteadyState() {
	_jsii_.InvokeVoid(
		t,
		"resetWaitForSteadyState",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

