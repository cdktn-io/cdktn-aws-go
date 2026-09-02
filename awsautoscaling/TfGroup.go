package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsautoscaling/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsautoscaling/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group aws_autoscaling_group}.
// Experimental.
type TfGroup interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	AvailabilityZoneDistribution() TfGroup_AvailabilityZoneDistributionPropertyOutputReference
	// Experimental.
	AvailabilityZoneDistributionInput() *TfGroup_AvailabilityZoneDistributionProperty
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
	CapacityReservationSpecification() TfGroup_CapacityReservationSpecificationPropertyOutputReference
	// Experimental.
	CapacityReservationSpecificationInput() *TfGroup_CapacityReservationSpecificationProperty
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
	InitialLifecycleHook() TfGroup_InitialLifecycleHookPropertyList
	// Experimental.
	InitialLifecycleHookInput() interface{}
	// Experimental.
	InstanceLifecyclePolicy() TfGroup_InstanceLifecyclePolicyPropertyOutputReference
	// Experimental.
	InstanceLifecyclePolicyInput() *TfGroup_InstanceLifecyclePolicyProperty
	// Experimental.
	InstanceMaintenancePolicy() TfGroup_InstanceMaintenancePolicyPropertyOutputReference
	// Experimental.
	InstanceMaintenancePolicyInput() *TfGroup_InstanceMaintenancePolicyProperty
	// Experimental.
	InstanceRefresh() TfGroup_InstanceRefreshPropertyOutputReference
	// Experimental.
	InstanceRefreshInput() *TfGroup_InstanceRefreshProperty
	// Experimental.
	LaunchConfiguration() *string
	// Experimental.
	SetLaunchConfiguration(val *string)
	// Experimental.
	LaunchConfigurationInput() *string
	// Experimental.
	LaunchTemplate() TfGroup_LaunchTemplatePropertyOutputReference
	// Experimental.
	LaunchTemplateInput() *TfGroup_LaunchTemplateProperty
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
	MixedInstancesPolicy() TfGroup_MixedInstancesPolicyPropertyOutputReference
	// Experimental.
	MixedInstancesPolicyInput() *TfGroup_MixedInstancesPolicyProperty
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
	Tag() TfGroup_TagPropertyList
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
	Timeouts() TfGroup_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TrafficSource() TfGroup_TrafficSourcePropertyList
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
	WarmPool() TfGroup_WarmPoolPropertyOutputReference
	// Experimental.
	WarmPoolInput() *TfGroup_WarmPoolProperty
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
	PutAvailabilityZoneDistribution(value *TfGroup_AvailabilityZoneDistributionProperty)
	// Experimental.
	PutCapacityReservationSpecification(value *TfGroup_CapacityReservationSpecificationProperty)
	// Experimental.
	PutInitialLifecycleHook(value interface{})
	// Experimental.
	PutInstanceLifecyclePolicy(value *TfGroup_InstanceLifecyclePolicyProperty)
	// Experimental.
	PutInstanceMaintenancePolicy(value *TfGroup_InstanceMaintenancePolicyProperty)
	// Experimental.
	PutInstanceRefresh(value *TfGroup_InstanceRefreshProperty)
	// Experimental.
	PutLaunchTemplate(value *TfGroup_LaunchTemplateProperty)
	// Experimental.
	PutMixedInstancesPolicy(value *TfGroup_MixedInstancesPolicyProperty)
	// Experimental.
	PutTag(value interface{})
	// Experimental.
	PutTimeouts(value *TfGroup_TimeoutsProperty)
	// Experimental.
	PutTrafficSource(value interface{})
	// Experimental.
	PutWarmPool(value *TfGroup_WarmPoolProperty)
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

// The jsii proxy struct for TfGroup
type jsiiProxy_TfGroup struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) AvailabilityZoneDistribution() TfGroup_AvailabilityZoneDistributionPropertyOutputReference {
	var returns TfGroup_AvailabilityZoneDistributionPropertyOutputReference
	_jsii_.Get(
		j,
		"availabilityZoneDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) AvailabilityZoneDistributionInput() *TfGroup_AvailabilityZoneDistributionProperty {
	var returns *TfGroup_AvailabilityZoneDistributionProperty
	_jsii_.Get(
		j,
		"availabilityZoneDistributionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) CapacityRebalance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityRebalance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) CapacityRebalanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityRebalanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) CapacityReservationSpecification() TfGroup_CapacityReservationSpecificationPropertyOutputReference {
	var returns TfGroup_CapacityReservationSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"capacityReservationSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) CapacityReservationSpecificationInput() *TfGroup_CapacityReservationSpecificationProperty {
	var returns *TfGroup_CapacityReservationSpecificationProperty
	_jsii_.Get(
		j,
		"capacityReservationSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) DefaultCooldown() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultCooldown",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) DefaultCooldownInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultCooldownInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) DefaultInstanceWarmup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultInstanceWarmup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) DefaultInstanceWarmupInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultInstanceWarmupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) DesiredCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) DesiredCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"desiredCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) DesiredCapacityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) DesiredCapacityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredCapacityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) EnabledMetrics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) EnabledMetricsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) ForceDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) ForceDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) ForceDeleteWarmPool() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteWarmPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) ForceDeleteWarmPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDeleteWarmPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) HealthCheckGracePeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) HealthCheckGracePeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckGracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) HealthCheckType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) HealthCheckTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) IgnoreFailedScalingActivities() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreFailedScalingActivities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) IgnoreFailedScalingActivitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreFailedScalingActivitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) InitialLifecycleHook() TfGroup_InitialLifecycleHookPropertyList {
	var returns TfGroup_InitialLifecycleHookPropertyList
	_jsii_.Get(
		j,
		"initialLifecycleHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) InitialLifecycleHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialLifecycleHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) InstanceLifecyclePolicy() TfGroup_InstanceLifecyclePolicyPropertyOutputReference {
	var returns TfGroup_InstanceLifecyclePolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"instanceLifecyclePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) InstanceLifecyclePolicyInput() *TfGroup_InstanceLifecyclePolicyProperty {
	var returns *TfGroup_InstanceLifecyclePolicyProperty
	_jsii_.Get(
		j,
		"instanceLifecyclePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) InstanceMaintenancePolicy() TfGroup_InstanceMaintenancePolicyPropertyOutputReference {
	var returns TfGroup_InstanceMaintenancePolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"instanceMaintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) InstanceMaintenancePolicyInput() *TfGroup_InstanceMaintenancePolicyProperty {
	var returns *TfGroup_InstanceMaintenancePolicyProperty
	_jsii_.Get(
		j,
		"instanceMaintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) InstanceRefresh() TfGroup_InstanceRefreshPropertyOutputReference {
	var returns TfGroup_InstanceRefreshPropertyOutputReference
	_jsii_.Get(
		j,
		"instanceRefresh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) InstanceRefreshInput() *TfGroup_InstanceRefreshProperty {
	var returns *TfGroup_InstanceRefreshProperty
	_jsii_.Get(
		j,
		"instanceRefreshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) LaunchConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) LaunchConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) LaunchTemplate() TfGroup_LaunchTemplatePropertyOutputReference {
	var returns TfGroup_LaunchTemplatePropertyOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) LaunchTemplateInput() *TfGroup_LaunchTemplateProperty {
	var returns *TfGroup_LaunchTemplateProperty
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) LoadBalancers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) LoadBalancersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) MaxInstanceLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) MaxInstanceLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) MaxSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) MaxSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) MetricsGranularity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsGranularity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) MetricsGranularityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsGranularityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) MinElbCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minElbCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) MinElbCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minElbCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) MinSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) MinSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) MixedInstancesPolicy() TfGroup_MixedInstancesPolicyPropertyOutputReference {
	var returns TfGroup_MixedInstancesPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"mixedInstancesPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) MixedInstancesPolicyInput() *TfGroup_MixedInstancesPolicyProperty {
	var returns *TfGroup_MixedInstancesPolicyProperty
	_jsii_.Get(
		j,
		"mixedInstancesPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) PlacementGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) PlacementGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"placementGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) PredictedCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"predictedCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) ProtectFromScaleIn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectFromScaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) ProtectFromScaleInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectFromScaleInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) ServiceLinkedRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLinkedRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) ServiceLinkedRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLinkedRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) SuspendedProcesses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"suspendedProcesses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) SuspendedProcessesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"suspendedProcessesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) Tag() TfGroup_TagPropertyList {
	var returns TfGroup_TagPropertyList
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) TagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) TargetGroupArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) TargetGroupArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) TerminationPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"terminationPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) TerminationPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"terminationPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) Timeouts() TfGroup_TimeoutsPropertyOutputReference {
	var returns TfGroup_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) TrafficSource() TfGroup_TrafficSourcePropertyList {
	var returns TfGroup_TrafficSourcePropertyList
	_jsii_.Get(
		j,
		"trafficSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) TrafficSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) VpcZoneIdentifier() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcZoneIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) VpcZoneIdentifierInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcZoneIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) WaitForCapacityTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitForCapacityTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) WaitForCapacityTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitForCapacityTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) WaitForElbCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForElbCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) WaitForElbCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitForElbCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) WarmPool() TfGroup_WarmPoolPropertyOutputReference {
	var returns TfGroup_WarmPoolPropertyOutputReference
	_jsii_.Get(
		j,
		"warmPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) WarmPoolInput() *TfGroup_WarmPoolProperty {
	var returns *TfGroup_WarmPoolProperty
	_jsii_.Get(
		j,
		"warmPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGroup) WarmPoolSize() *float64 {
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
func NewTfGroup(scope constructs.Construct, id *string, config *TfGroupConfig) TfGroup {
	_init_.Initialize()

	if err := validateNewTfGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGroup{}

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.TfGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group aws_autoscaling_group} Resource.
// Experimental.
func NewTfGroup_Override(t TfGroup, scope constructs.Construct, id *string, config *TfGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-auto-scaling.TfGroup",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfGroup)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetCapacityRebalance(val interface{}) {
	if err := j.validateSetCapacityRebalanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityRebalance",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetContext(val *string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetDefaultCooldown(val *float64) {
	if err := j.validateSetDefaultCooldownParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultCooldown",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetDefaultInstanceWarmup(val *float64) {
	if err := j.validateSetDefaultInstanceWarmupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultInstanceWarmup",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetDesiredCapacity(val *float64) {
	if err := j.validateSetDesiredCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacity",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetDesiredCapacityType(val *string) {
	if err := j.validateSetDesiredCapacityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredCapacityType",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetEnabledMetrics(val *[]*string) {
	if err := j.validateSetEnabledMetricsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledMetrics",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetForceDelete(val interface{}) {
	if err := j.validateSetForceDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDelete",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetForceDeleteWarmPool(val interface{}) {
	if err := j.validateSetForceDeleteWarmPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDeleteWarmPool",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetHealthCheckGracePeriod(val *float64) {
	if err := j.validateSetHealthCheckGracePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckGracePeriod",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetHealthCheckType(val *string) {
	if err := j.validateSetHealthCheckTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckType",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetIgnoreFailedScalingActivities(val interface{}) {
	if err := j.validateSetIgnoreFailedScalingActivitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreFailedScalingActivities",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetLaunchConfiguration(val *string) {
	if err := j.validateSetLaunchConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchConfiguration",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetLoadBalancers(val *[]*string) {
	if err := j.validateSetLoadBalancersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancers",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetMaxInstanceLifetime(val *float64) {
	if err := j.validateSetMaxInstanceLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstanceLifetime",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetMaxSize(val *float64) {
	if err := j.validateSetMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSize",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetMetricsGranularity(val *string) {
	if err := j.validateSetMetricsGranularityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsGranularity",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetMinElbCapacity(val *float64) {
	if err := j.validateSetMinElbCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minElbCapacity",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetMinSize(val *float64) {
	if err := j.validateSetMinSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSize",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetPlacementGroup(val *string) {
	if err := j.validateSetPlacementGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementGroup",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetProtectFromScaleIn(val interface{}) {
	if err := j.validateSetProtectFromScaleInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectFromScaleIn",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetServiceLinkedRoleArn(val *string) {
	if err := j.validateSetServiceLinkedRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceLinkedRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetSuspendedProcesses(val *[]*string) {
	if err := j.validateSetSuspendedProcessesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendedProcesses",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetTargetGroupArns(val *[]*string) {
	if err := j.validateSetTargetGroupArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetGroupArns",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetTerminationPolicies(val *[]*string) {
	if err := j.validateSetTerminationPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationPolicies",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetVpcZoneIdentifier(val *[]*string) {
	if err := j.validateSetVpcZoneIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcZoneIdentifier",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetWaitForCapacityTimeout(val *string) {
	if err := j.validateSetWaitForCapacityTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForCapacityTimeout",
		val,
	)
}

func (j *jsiiProxy_TfGroup)SetWaitForElbCapacity(val *float64) {
	if err := j.validateSetWaitForElbCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForElbCapacity",
		val,
	)
}

// Generates CDKTN code for importing a TfGroup resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-auto-scaling.TfGroup",
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
func TfGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-auto-scaling.TfGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-auto-scaling.TfGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-auto-scaling.TfGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-auto-scaling.TfGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfGroup) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfGroup) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGroup) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGroup) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfGroup) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGroup) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfGroup) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfGroup) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfGroup) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfGroup) PutAvailabilityZoneDistribution(value *TfGroup_AvailabilityZoneDistributionProperty) {
	if err := t.validatePutAvailabilityZoneDistributionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAvailabilityZoneDistribution",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGroup) PutCapacityReservationSpecification(value *TfGroup_CapacityReservationSpecificationProperty) {
	if err := t.validatePutCapacityReservationSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCapacityReservationSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGroup) PutInitialLifecycleHook(value interface{}) {
	if err := t.validatePutInitialLifecycleHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInitialLifecycleHook",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGroup) PutInstanceLifecyclePolicy(value *TfGroup_InstanceLifecyclePolicyProperty) {
	if err := t.validatePutInstanceLifecyclePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInstanceLifecyclePolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGroup) PutInstanceMaintenancePolicy(value *TfGroup_InstanceMaintenancePolicyProperty) {
	if err := t.validatePutInstanceMaintenancePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInstanceMaintenancePolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGroup) PutInstanceRefresh(value *TfGroup_InstanceRefreshProperty) {
	if err := t.validatePutInstanceRefreshParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInstanceRefresh",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGroup) PutLaunchTemplate(value *TfGroup_LaunchTemplateProperty) {
	if err := t.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGroup) PutMixedInstancesPolicy(value *TfGroup_MixedInstancesPolicyProperty) {
	if err := t.validatePutMixedInstancesPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMixedInstancesPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGroup) PutTag(value interface{}) {
	if err := t.validatePutTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTag",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGroup) PutTimeouts(value *TfGroup_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGroup) PutTrafficSource(value interface{}) {
	if err := t.validatePutTrafficSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTrafficSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGroup) PutWarmPool(value *TfGroup_WarmPoolProperty) {
	if err := t.validatePutWarmPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWarmPool",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGroup) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfGroup) ResetAvailabilityZoneDistribution() {
	_jsii_.InvokeVoid(
		t,
		"resetAvailabilityZoneDistribution",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetAvailabilityZones() {
	_jsii_.InvokeVoid(
		t,
		"resetAvailabilityZones",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetCapacityRebalance() {
	_jsii_.InvokeVoid(
		t,
		"resetCapacityRebalance",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetCapacityReservationSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetCapacityReservationSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetContext() {
	_jsii_.InvokeVoid(
		t,
		"resetContext",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetDefaultCooldown() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultCooldown",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetDefaultInstanceWarmup() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultInstanceWarmup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetDesiredCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetDesiredCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetDesiredCapacityType() {
	_jsii_.InvokeVoid(
		t,
		"resetDesiredCapacityType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetEnabledMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetEnabledMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetForceDelete() {
	_jsii_.InvokeVoid(
		t,
		"resetForceDelete",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetForceDeleteWarmPool() {
	_jsii_.InvokeVoid(
		t,
		"resetForceDeleteWarmPool",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetHealthCheckGracePeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetHealthCheckGracePeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetHealthCheckType() {
	_jsii_.InvokeVoid(
		t,
		"resetHealthCheckType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetIgnoreFailedScalingActivities() {
	_jsii_.InvokeVoid(
		t,
		"resetIgnoreFailedScalingActivities",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetInitialLifecycleHook() {
	_jsii_.InvokeVoid(
		t,
		"resetInitialLifecycleHook",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetInstanceLifecyclePolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceLifecyclePolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetInstanceMaintenancePolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceMaintenancePolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetInstanceRefresh() {
	_jsii_.InvokeVoid(
		t,
		"resetInstanceRefresh",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetLaunchConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetLaunchConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		t,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetLoadBalancers() {
	_jsii_.InvokeVoid(
		t,
		"resetLoadBalancers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetMaxInstanceLifetime() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxInstanceLifetime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetMetricsGranularity() {
	_jsii_.InvokeVoid(
		t,
		"resetMetricsGranularity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetMinElbCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetMinElbCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetMixedInstancesPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetMixedInstancesPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetPlacementGroup() {
	_jsii_.InvokeVoid(
		t,
		"resetPlacementGroup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetProtectFromScaleIn() {
	_jsii_.InvokeVoid(
		t,
		"resetProtectFromScaleIn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetServiceLinkedRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceLinkedRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetSuspendedProcesses() {
	_jsii_.InvokeVoid(
		t,
		"resetSuspendedProcesses",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetTag() {
	_jsii_.InvokeVoid(
		t,
		"resetTag",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetTargetGroupArns() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetGroupArns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetTerminationPolicies() {
	_jsii_.InvokeVoid(
		t,
		"resetTerminationPolicies",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetTrafficSource() {
	_jsii_.InvokeVoid(
		t,
		"resetTrafficSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetVpcZoneIdentifier() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcZoneIdentifier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetWaitForCapacityTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetWaitForCapacityTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetWaitForElbCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetWaitForElbCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) ResetWarmPool() {
	_jsii_.InvokeVoid(
		t,
		"resetWarmPool",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

