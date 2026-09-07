package autoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/autoscaling/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/autoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group aws_autoscaling_group}.
// Experimental.
type AwsGroup interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	AvailabilityZoneDistribution() AwsGroup_AvailabilityZoneDistributionPropertyOutputReference
	// Experimental.
	AvailabilityZoneDistributionInput() *AwsGroup_AvailabilityZoneDistributionProperty
	// Experimental.
	AvailabilityZones() *[]*string
	// Experimental.
	SetAvailabilityZones(val *[]*string)
	// Experimental.
	AvailabilityZonesInput() *[]*string
	// Experimental.
	CapacityRebalance() interface{}
	// Experimental.
	SetCapacityRebalance(val interface{})
	// Experimental.
	CapacityRebalanceInput() interface{}
	// Experimental.
	CapacityReservationSpecification() AwsGroup_CapacityReservationSpecificationPropertyOutputReference
	// Experimental.
	CapacityReservationSpecificationInput() *AwsGroup_CapacityReservationSpecificationProperty
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Context() *string
	// Experimental.
	SetContext(val *string)
	// Experimental.
	ContextInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DefaultCooldown() *float64
	// Experimental.
	SetDefaultCooldown(val *float64)
	// Experimental.
	DefaultCooldownInput() *float64
	// Experimental.
	DefaultInstanceWarmup() *float64
	// Experimental.
	SetDefaultInstanceWarmup(val *float64)
	// Experimental.
	DefaultInstanceWarmupInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DesiredCapacity() *float64
	// Experimental.
	SetDesiredCapacity(val *float64)
	// Experimental.
	DesiredCapacityInput() *float64
	// Experimental.
	DesiredCapacityType() *string
	// Experimental.
	SetDesiredCapacityType(val *string)
	// Experimental.
	DesiredCapacityTypeInput() *string
	// Experimental.
	EnabledMetrics() *[]*string
	// Experimental.
	SetEnabledMetrics(val *[]*string)
	// Experimental.
	EnabledMetricsInput() *[]*string
	// Experimental.
	ForceDelete() interface{}
	// Experimental.
	SetForceDelete(val interface{})
	// Experimental.
	ForceDeleteInput() interface{}
	// Experimental.
	ForceDeleteWarmPool() interface{}
	// Experimental.
	SetForceDeleteWarmPool(val interface{})
	// Experimental.
	ForceDeleteWarmPoolInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HealthCheckGracePeriod() *float64
	// Experimental.
	SetHealthCheckGracePeriod(val *float64)
	// Experimental.
	HealthCheckGracePeriodInput() *float64
	// Experimental.
	HealthCheckType() *string
	// Experimental.
	SetHealthCheckType(val *string)
	// Experimental.
	HealthCheckTypeInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	IgnoreFailedScalingActivities() interface{}
	// Experimental.
	SetIgnoreFailedScalingActivities(val interface{})
	// Experimental.
	IgnoreFailedScalingActivitiesInput() interface{}
	// Experimental.
	InitialLifecycleHook() AwsGroup_InitialLifecycleHookPropertyList
	// Experimental.
	InitialLifecycleHookInput() interface{}
	// Experimental.
	InstanceLifecyclePolicy() AwsGroup_InstanceLifecyclePolicyPropertyOutputReference
	// Experimental.
	InstanceLifecyclePolicyInput() *AwsGroup_InstanceLifecyclePolicyProperty
	// Experimental.
	InstanceMaintenancePolicy() AwsGroup_InstanceMaintenancePolicyPropertyOutputReference
	// Experimental.
	InstanceMaintenancePolicyInput() *AwsGroup_InstanceMaintenancePolicyProperty
	// Experimental.
	InstanceRefresh() AwsGroup_InstanceRefreshPropertyOutputReference
	// Experimental.
	InstanceRefreshInput() *AwsGroup_InstanceRefreshProperty
	// Experimental.
	LaunchConfiguration() *string
	// Experimental.
	SetLaunchConfiguration(val *string)
	// Experimental.
	LaunchConfigurationInput() *string
	// Experimental.
	LaunchTemplate() AwsGroup_LaunchTemplatePropertyOutputReference
	// Experimental.
	LaunchTemplateInput() *AwsGroup_LaunchTemplateProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LoadBalancers() *[]*string
	// Experimental.
	SetLoadBalancers(val *[]*string)
	// Experimental.
	LoadBalancersInput() *[]*string
	// Experimental.
	MaxInstanceLifetime() *float64
	// Experimental.
	SetMaxInstanceLifetime(val *float64)
	// Experimental.
	MaxInstanceLifetimeInput() *float64
	// Experimental.
	MaxSize() *float64
	// Experimental.
	SetMaxSize(val *float64)
	// Experimental.
	MaxSizeInput() *float64
	// Experimental.
	MetricsGranularity() *string
	// Experimental.
	SetMetricsGranularity(val *string)
	// Experimental.
	MetricsGranularityInput() *string
	// Experimental.
	MinElbCapacity() *float64
	// Experimental.
	SetMinElbCapacity(val *float64)
	// Experimental.
	MinElbCapacityInput() *float64
	// Experimental.
	MinSize() *float64
	// Experimental.
	SetMinSize(val *float64)
	// Experimental.
	MinSizeInput() *float64
	// Experimental.
	MixedInstancesPolicy() AwsGroup_MixedInstancesPolicyPropertyOutputReference
	// Experimental.
	MixedInstancesPolicyInput() *AwsGroup_MixedInstancesPolicyProperty
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NamePrefix() *string
	// Experimental.
	SetNamePrefix(val *string)
	// Experimental.
	NamePrefixInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PlacementGroup() *string
	// Experimental.
	SetPlacementGroup(val *string)
	// Experimental.
	PlacementGroupInput() *string
	// Experimental.
	PredictedCapacity() *float64
	// Experimental.
	ProtectFromScaleIn() interface{}
	// Experimental.
	SetProtectFromScaleIn(val interface{})
	// Experimental.
	ProtectFromScaleInInput() interface{}
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
	ServiceLinkedRoleArn() *string
	// Experimental.
	SetServiceLinkedRoleArn(val *string)
	// Experimental.
	ServiceLinkedRoleArnInput() *string
	// Experimental.
	SuspendedProcesses() *[]*string
	// Experimental.
	SetSuspendedProcesses(val *[]*string)
	// Experimental.
	SuspendedProcessesInput() *[]*string
	// Experimental.
	Tag() AwsGroup_TagPropertyList
	// Experimental.
	TagInput() interface{}
	// Experimental.
	TargetGroupArns() *[]*string
	// Experimental.
	SetTargetGroupArns(val *[]*string)
	// Experimental.
	TargetGroupArnsInput() *[]*string
	// Experimental.
	TerminationPolicies() *[]*string
	// Experimental.
	SetTerminationPolicies(val *[]*string)
	// Experimental.
	TerminationPoliciesInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsGroup_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TrafficSource() AwsGroup_TrafficSourcePropertyList
	// Experimental.
	TrafficSourceInput() interface{}
	// Experimental.
	VpcZoneIdentifier() *[]*string
	// Experimental.
	SetVpcZoneIdentifier(val *[]*string)
	// Experimental.
	VpcZoneIdentifierInput() *[]*string
	// Experimental.
	WaitForCapacityTimeout() *string
	// Experimental.
	SetWaitForCapacityTimeout(val *string)
	// Experimental.
	WaitForCapacityTimeoutInput() *string
	// Experimental.
	WaitForElbCapacity() *float64
	// Experimental.
	SetWaitForElbCapacity(val *float64)
	// Experimental.
	WaitForElbCapacityInput() *float64
	// Experimental.
	WarmPool() AwsGroup_WarmPoolPropertyOutputReference
	// Experimental.
	WarmPoolInput() *AwsGroup_WarmPoolProperty
	// Experimental.
	WarmPoolSize() *float64
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
	PutAvailabilityZoneDistribution(value *AwsGroup_AvailabilityZoneDistributionProperty)
	// Experimental.
	PutCapacityReservationSpecification(value *AwsGroup_CapacityReservationSpecificationProperty)
	// Experimental.
	PutInitialLifecycleHook(value interface{})
	// Experimental.
	PutInstanceLifecyclePolicy(value *AwsGroup_InstanceLifecyclePolicyProperty)
	// Experimental.
	PutInstanceMaintenancePolicy(value *AwsGroup_InstanceMaintenancePolicyProperty)
	// Experimental.
	PutInstanceRefresh(value *AwsGroup_InstanceRefreshProperty)
	// Experimental.
	PutLaunchTemplate(value *AwsGroup_LaunchTemplateProperty)
	// Experimental.
	PutMixedInstancesPolicy(value *AwsGroup_MixedInstancesPolicyProperty)
	// Experimental.
	PutTag(value interface{})
	// Experimental.
	PutTimeouts(value *AwsGroup_TimeoutsProperty)
	// Experimental.
	PutTrafficSource(value interface{})
	// Experimental.
	PutWarmPool(value *AwsGroup_WarmPoolProperty)
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
	ResetAvailabilityZoneDistribution()
	// Experimental.
	ResetAvailabilityZones()
	// Experimental.
	ResetCapacityRebalance()
	// Experimental.
	ResetCapacityReservationSpecification()
	// Experimental.
	ResetContext()
	// Experimental.
	ResetDefaultCooldown()
	// Experimental.
	ResetDefaultInstanceWarmup()
	// Experimental.
	ResetDesiredCapacity()
	// Experimental.
	ResetDesiredCapacityType()
	// Experimental.
	ResetEnabledMetrics()
	// Experimental.
	ResetForceDelete()
	// Experimental.
	ResetForceDeleteWarmPool()
	// Experimental.
	ResetHealthCheckGracePeriod()
	// Experimental.
	ResetHealthCheckType()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIgnoreFailedScalingActivities()
	// Experimental.
	ResetInitialLifecycleHook()
	// Experimental.
	ResetInstanceLifecyclePolicy()
	// Experimental.
	ResetInstanceMaintenancePolicy()
	// Experimental.
	ResetInstanceRefresh()
	// Experimental.
	ResetLaunchConfiguration()
	// Experimental.
	ResetLaunchTemplate()
	// Experimental.
	ResetLoadBalancers()
	// Experimental.
	ResetMaxInstanceLifetime()
	// Experimental.
	ResetMetricsGranularity()
	// Experimental.
	ResetMinElbCapacity()
	// Experimental.
	ResetMixedInstancesPolicy()
	// Experimental.
	ResetName()
	// Experimental.
	ResetNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPlacementGroup()
	// Experimental.
	ResetProtectFromScaleIn()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetServiceLinkedRoleArn()
	// Experimental.
	ResetSuspendedProcesses()
	// Experimental.
	ResetTag()
	// Experimental.
	ResetTargetGroupArns()
	// Experimental.
	ResetTerminationPolicies()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTrafficSource()
	// Experimental.
	ResetVpcZoneIdentifier()
	// Experimental.
	ResetWaitForCapacityTimeout()
	// Experimental.
	ResetWaitForElbCapacity()
	// Experimental.
	ResetWarmPool()
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

// The jsii proxy struct for AwsGroup
type jsiiProxy_AwsGroup struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) AvailabilityZoneDistribution() AwsGroup_AvailabilityZoneDistributionPropertyOutputReference {
	var returns AwsGroup_AvailabilityZoneDistributionPropertyOutputReference
	_jsii_.Get(
		j,
		"availabilityZoneDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) AvailabilityZoneDistributionInput() *AwsGroup_AvailabilityZoneDistributionProperty {
	var returns *AwsGroup_AvailabilityZoneDistributionProperty
	_jsii_.Get(
		j,
		"availabilityZoneDistributionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) CapacityRebalance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityRebalance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) CapacityRebalanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityRebalanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) CapacityReservationSpecification() AwsGroup_CapacityReservationSpecificationPropertyOutputReference {
	var returns AwsGroup_CapacityReservationSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"capacityReservationSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) CapacityReservationSpecificationInput() *AwsGroup_CapacityReservationSpecificationProperty {
	var returns *AwsGroup_CapacityReservationSpecificationProperty
	_jsii_.Get(
		j,
		"capacityReservationSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) DefaultCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) DefaultCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) DefaultInstanceWarmup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultInstanceWarmup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) DefaultInstanceWarmupInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultInstanceWarmupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) DesiredCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) DesiredCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) DesiredCapacityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) DesiredCapacityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) EnabledMetrics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) EnabledMetricsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) ForceDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) ForceDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) ForceDeleteWarmPool() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteWarmPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) ForceDeleteWarmPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteWarmPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) HealthCheckGracePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) HealthCheckGracePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) HealthCheckType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) HealthCheckTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) IgnoreFailedScalingActivities() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreFailedScalingActivities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) IgnoreFailedScalingActivitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreFailedScalingActivitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) InitialLifecycleHook() AwsGroup_InitialLifecycleHookPropertyList {
	var returns AwsGroup_InitialLifecycleHookPropertyList
	_jsii_.Get(
		j,
		"initialLifecycleHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) InitialLifecycleHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialLifecycleHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) InstanceLifecyclePolicy() AwsGroup_InstanceLifecyclePolicyPropertyOutputReference {
	var returns AwsGroup_InstanceLifecyclePolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"instanceLifecyclePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) InstanceLifecyclePolicyInput() *AwsGroup_InstanceLifecyclePolicyProperty {
	var returns *AwsGroup_InstanceLifecyclePolicyProperty
	_jsii_.Get(
		j,
		"instanceLifecyclePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) InstanceMaintenancePolicy() AwsGroup_InstanceMaintenancePolicyPropertyOutputReference {
	var returns AwsGroup_InstanceMaintenancePolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"instanceMaintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) InstanceMaintenancePolicyInput() *AwsGroup_InstanceMaintenancePolicyProperty {
	var returns *AwsGroup_InstanceMaintenancePolicyProperty
	_jsii_.Get(
		j,
		"instanceMaintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) InstanceRefresh() AwsGroup_InstanceRefreshPropertyOutputReference {
	var returns AwsGroup_InstanceRefreshPropertyOutputReference
	_jsii_.Get(
		j,
		"instanceRefresh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) InstanceRefreshInput() *AwsGroup_InstanceRefreshProperty {
	var returns *AwsGroup_InstanceRefreshProperty
	_jsii_.Get(
		j,
		"instanceRefreshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) LaunchConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) LaunchConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) LaunchTemplate() AwsGroup_LaunchTemplatePropertyOutputReference {
	var returns AwsGroup_LaunchTemplatePropertyOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) LaunchTemplateInput() *AwsGroup_LaunchTemplateProperty {
	var returns *AwsGroup_LaunchTemplateProperty
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) LoadBalancers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) LoadBalancersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) MaxInstanceLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) MaxInstanceLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) MetricsGranularity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsGranularity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) MetricsGranularityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsGranularityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) MinElbCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minElbCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) MinElbCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minElbCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) MixedInstancesPolicy() AwsGroup_MixedInstancesPolicyPropertyOutputReference {
	var returns AwsGroup_MixedInstancesPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"mixedInstancesPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) MixedInstancesPolicyInput() *AwsGroup_MixedInstancesPolicyProperty {
	var returns *AwsGroup_MixedInstancesPolicyProperty
	_jsii_.Get(
		j,
		"mixedInstancesPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) PlacementGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) PredictedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"predictedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) ProtectFromScaleIn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectFromScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) ProtectFromScaleInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectFromScaleInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) ServiceLinkedRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLinkedRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) ServiceLinkedRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLinkedRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) SuspendedProcesses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"suspendedProcesses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) SuspendedProcessesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"suspendedProcessesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) Tag() AwsGroup_TagPropertyList {
	var returns AwsGroup_TagPropertyList
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) TagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) TargetGroupArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) TargetGroupArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) TerminationPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"terminationPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) TerminationPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"terminationPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) Timeouts() AwsGroup_TimeoutsPropertyOutputReference {
	var returns AwsGroup_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) TrafficSource() AwsGroup_TrafficSourcePropertyList {
	var returns AwsGroup_TrafficSourcePropertyList
	_jsii_.Get(
		j,
		"trafficSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) TrafficSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) VpcZoneIdentifier() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcZoneIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) VpcZoneIdentifierInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcZoneIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) WaitForCapacityTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitForCapacityTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) WaitForCapacityTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitForCapacityTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) WaitForElbCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForElbCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) WaitForElbCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForElbCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) WarmPool() AwsGroup_WarmPoolPropertyOutputReference {
	var returns AwsGroup_WarmPoolPropertyOutputReference
	_jsii_.Get(
		j,
		"warmPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) WarmPoolInput() *AwsGroup_WarmPoolProperty {
	var returns *AwsGroup_WarmPoolProperty
	_jsii_.Get(
		j,
		"warmPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGroup) WarmPoolSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"warmPoolSize",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group aws_autoscaling_group} Resource.
// Experimental.
func NewAwsGroup(scope constructs.Construct, id *string, config *AwsGroupConfig) AwsGroup {
	_init_.Initialize()

	if err := validateNewAwsGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGroup{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.AwsGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group aws_autoscaling_group} Resource.
// Experimental.
func NewAwsGroup_Override(a AwsGroup, scope constructs.Construct, id *string, config *AwsGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.AwsGroup",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsGroup)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetCapacityRebalance(val interface{}) {
	if err := j.validateSetCapacityRebalanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityRebalance",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetContext(val *string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetDefaultCooldown(val *float64) {
	if err := j.validateSetDefaultCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultCooldown",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetDefaultInstanceWarmup(val *float64) {
	if err := j.validateSetDefaultInstanceWarmupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultInstanceWarmup",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetDesiredCapacity(val *float64) {
	if err := j.validateSetDesiredCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetDesiredCapacityType(val *string) {
	if err := j.validateSetDesiredCapacityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacityType",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetEnabledMetrics(val *[]*string) {
	if err := j.validateSetEnabledMetricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledMetrics",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetForceDelete(val interface{}) {
	if err := j.validateSetForceDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDelete",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetForceDeleteWarmPool(val interface{}) {
	if err := j.validateSetForceDeleteWarmPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDeleteWarmPool",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetHealthCheckGracePeriod(val *float64) {
	if err := j.validateSetHealthCheckGracePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckGracePeriod",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetHealthCheckType(val *string) {
	if err := j.validateSetHealthCheckTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckType",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetIgnoreFailedScalingActivities(val interface{}) {
	if err := j.validateSetIgnoreFailedScalingActivitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreFailedScalingActivities",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetLaunchConfiguration(val *string) {
	if err := j.validateSetLaunchConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchConfiguration",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetLoadBalancers(val *[]*string) {
	if err := j.validateSetLoadBalancersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancers",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetMaxInstanceLifetime(val *float64) {
	if err := j.validateSetMaxInstanceLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstanceLifetime",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetMaxSize(val *float64) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetMetricsGranularity(val *string) {
	if err := j.validateSetMetricsGranularityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsGranularity",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetMinElbCapacity(val *float64) {
	if err := j.validateSetMinElbCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minElbCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetPlacementGroup(val *string) {
	if err := j.validateSetPlacementGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroup",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetProtectFromScaleIn(val interface{}) {
	if err := j.validateSetProtectFromScaleInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectFromScaleIn",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetServiceLinkedRoleArn(val *string) {
	if err := j.validateSetServiceLinkedRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceLinkedRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetSuspendedProcesses(val *[]*string) {
	if err := j.validateSetSuspendedProcessesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendedProcesses",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetTargetGroupArns(val *[]*string) {
	if err := j.validateSetTargetGroupArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetGroupArns",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetTerminationPolicies(val *[]*string) {
	if err := j.validateSetTerminationPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationPolicies",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetVpcZoneIdentifier(val *[]*string) {
	if err := j.validateSetVpcZoneIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcZoneIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetWaitForCapacityTimeout(val *string) {
	if err := j.validateSetWaitForCapacityTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForCapacityTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsGroup)SetWaitForElbCapacity(val *float64) {
	if err := j.validateSetWaitForElbCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForElbCapacity",
		val,
	)
}

// Generates CDKTN code for importing a AwsGroup resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-auto-scaling.AwsGroup",
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
func AwsGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-auto-scaling.AwsGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-auto-scaling.AwsGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-auto-scaling.AwsGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-auto-scaling.AwsGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsGroup) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsGroup) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGroup) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGroup) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsGroup) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGroup) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsGroup) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsGroup) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsGroup) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsGroup) PutAvailabilityZoneDistribution(value *AwsGroup_AvailabilityZoneDistributionProperty) {
	if err := a.validatePutAvailabilityZoneDistributionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAvailabilityZoneDistribution",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGroup) PutCapacityReservationSpecification(value *AwsGroup_CapacityReservationSpecificationProperty) {
	if err := a.validatePutCapacityReservationSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCapacityReservationSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGroup) PutInitialLifecycleHook(value interface{}) {
	if err := a.validatePutInitialLifecycleHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInitialLifecycleHook",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGroup) PutInstanceLifecyclePolicy(value *AwsGroup_InstanceLifecyclePolicyProperty) {
	if err := a.validatePutInstanceLifecyclePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstanceLifecyclePolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGroup) PutInstanceMaintenancePolicy(value *AwsGroup_InstanceMaintenancePolicyProperty) {
	if err := a.validatePutInstanceMaintenancePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstanceMaintenancePolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGroup) PutInstanceRefresh(value *AwsGroup_InstanceRefreshProperty) {
	if err := a.validatePutInstanceRefreshParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstanceRefresh",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGroup) PutLaunchTemplate(value *AwsGroup_LaunchTemplateProperty) {
	if err := a.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGroup) PutMixedInstancesPolicy(value *AwsGroup_MixedInstancesPolicyProperty) {
	if err := a.validatePutMixedInstancesPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMixedInstancesPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGroup) PutTag(value interface{}) {
	if err := a.validatePutTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTag",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGroup) PutTimeouts(value *AwsGroup_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGroup) PutTrafficSource(value interface{}) {
	if err := a.validatePutTrafficSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrafficSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGroup) PutWarmPool(value *AwsGroup_WarmPoolProperty) {
	if err := a.validatePutWarmPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWarmPool",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGroup) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsGroup) ResetAvailabilityZoneDistribution() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZoneDistribution",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetAvailabilityZones() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZones",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetCapacityRebalance() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacityRebalance",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetCapacityReservationSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacityReservationSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetContext() {
	_jsii_.InvokeVoid(
		a,
		"resetContext",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetDefaultCooldown() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultCooldown",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetDefaultInstanceWarmup() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultInstanceWarmup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetDesiredCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetDesiredCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetDesiredCapacityType() {
	_jsii_.InvokeVoid(
		a,
		"resetDesiredCapacityType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetEnabledMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabledMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetForceDelete() {
	_jsii_.InvokeVoid(
		a,
		"resetForceDelete",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetForceDeleteWarmPool() {
	_jsii_.InvokeVoid(
		a,
		"resetForceDeleteWarmPool",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetHealthCheckGracePeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheckGracePeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetHealthCheckType() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheckType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetIgnoreFailedScalingActivities() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreFailedScalingActivities",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetInitialLifecycleHook() {
	_jsii_.InvokeVoid(
		a,
		"resetInitialLifecycleHook",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetInstanceLifecyclePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceLifecyclePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetInstanceMaintenancePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceMaintenancePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetInstanceRefresh() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceRefresh",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetLaunchConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetLoadBalancers() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetMaxInstanceLifetime() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxInstanceLifetime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetMetricsGranularity() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricsGranularity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetMinElbCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetMinElbCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetMixedInstancesPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetMixedInstancesPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetPlacementGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacementGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetProtectFromScaleIn() {
	_jsii_.InvokeVoid(
		a,
		"resetProtectFromScaleIn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetServiceLinkedRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceLinkedRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetSuspendedProcesses() {
	_jsii_.InvokeVoid(
		a,
		"resetSuspendedProcesses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetTag() {
	_jsii_.InvokeVoid(
		a,
		"resetTag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetTargetGroupArns() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetGroupArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetTerminationPolicies() {
	_jsii_.InvokeVoid(
		a,
		"resetTerminationPolicies",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetTrafficSource() {
	_jsii_.InvokeVoid(
		a,
		"resetTrafficSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetVpcZoneIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcZoneIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetWaitForCapacityTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetWaitForCapacityTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetWaitForElbCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetWaitForElbCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) ResetWarmPool() {
	_jsii_.InvokeVoid(
		a,
		"resetWarmPool",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

