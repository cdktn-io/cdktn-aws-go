package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request aws_spot_fleet_request}.
// Experimental.
type AwsSpotFleetRequest interface {
	cdktn.TerraformResource
	// Experimental.
	AllocationStrategy() *string
	// Experimental.
	SetAllocationStrategy(val *string)
	// Experimental.
	AllocationStrategyInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClientToken() *string
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
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ExcessCapacityTerminationPolicy() *string
	// Experimental.
	SetExcessCapacityTerminationPolicy(val *string)
	// Experimental.
	ExcessCapacityTerminationPolicyInput() *string
	// Experimental.
	FleetType() *string
	// Experimental.
	SetFleetType(val *string)
	// Experimental.
	FleetTypeInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	IamFleetRole() *string
	// Experimental.
	SetIamFleetRole(val *string)
	// Experimental.
	IamFleetRoleInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	InstanceInterruptionBehaviour() *string
	// Experimental.
	SetInstanceInterruptionBehaviour(val *string)
	// Experimental.
	InstanceInterruptionBehaviourInput() *string
	// Experimental.
	InstancePoolsToUseCount() *float64
	// Experimental.
	SetInstancePoolsToUseCount(val *float64)
	// Experimental.
	InstancePoolsToUseCountInput() *float64
	// Experimental.
	LaunchSpecification() AwsSpotFleetRequest_LaunchSpecificationPropertyList
	// Experimental.
	LaunchSpecificationInput() interface{}
	// Experimental.
	LaunchTemplateConfig() AwsSpotFleetRequest_LaunchTemplateConfigPropertyList
	// Experimental.
	LaunchTemplateConfigInput() interface{}
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
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OnDemandAllocationStrategy() *string
	// Experimental.
	SetOnDemandAllocationStrategy(val *string)
	// Experimental.
	OnDemandAllocationStrategyInput() *string
	// Experimental.
	OnDemandMaxTotalPrice() *string
	// Experimental.
	SetOnDemandMaxTotalPrice(val *string)
	// Experimental.
	OnDemandMaxTotalPriceInput() *string
	// Experimental.
	OnDemandTargetCapacity() *float64
	// Experimental.
	SetOnDemandTargetCapacity(val *float64)
	// Experimental.
	OnDemandTargetCapacityInput() *float64
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
	ReplaceUnhealthyInstances() interface{}
	// Experimental.
	SetReplaceUnhealthyInstances(val interface{})
	// Experimental.
	ReplaceUnhealthyInstancesInput() interface{}
	// Experimental.
	SpotMaintenanceStrategies() AwsSpotFleetRequest_SpotMaintenanceStrategiesPropertyOutputReference
	// Experimental.
	SpotMaintenanceStrategiesInput() *AwsSpotFleetRequest_SpotMaintenanceStrategiesProperty
	// Experimental.
	SpotPrice() *string
	// Experimental.
	SetSpotPrice(val *string)
	// Experimental.
	SpotPriceInput() *string
	// Experimental.
	SpotRequestState() *string
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
	TargetCapacity() *float64
	// Experimental.
	SetTargetCapacity(val *float64)
	// Experimental.
	TargetCapacityInput() *float64
	// Experimental.
	TargetCapacityUnitType() *string
	// Experimental.
	SetTargetCapacityUnitType(val *string)
	// Experimental.
	TargetCapacityUnitTypeInput() *string
	// Experimental.
	TargetGroupArns() *[]*string
	// Experimental.
	SetTargetGroupArns(val *[]*string)
	// Experimental.
	TargetGroupArnsInput() *[]*string
	// Experimental.
	TerminateInstancesOnDelete() *string
	// Experimental.
	SetTerminateInstancesOnDelete(val *string)
	// Experimental.
	TerminateInstancesOnDeleteInput() *string
	// Experimental.
	TerminateInstancesWithExpiration() interface{}
	// Experimental.
	SetTerminateInstancesWithExpiration(val interface{})
	// Experimental.
	TerminateInstancesWithExpirationInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsSpotFleetRequest_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	ValidFrom() *string
	// Experimental.
	SetValidFrom(val *string)
	// Experimental.
	ValidFromInput() *string
	// Experimental.
	ValidUntil() *string
	// Experimental.
	SetValidUntil(val *string)
	// Experimental.
	ValidUntilInput() *string
	// Experimental.
	WaitForFulfillment() interface{}
	// Experimental.
	SetWaitForFulfillment(val interface{})
	// Experimental.
	WaitForFulfillmentInput() interface{}
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
	PutLaunchSpecification(value interface{})
	// Experimental.
	PutLaunchTemplateConfig(value interface{})
	// Experimental.
	PutSpotMaintenanceStrategies(value *AwsSpotFleetRequest_SpotMaintenanceStrategiesProperty)
	// Experimental.
	PutTimeouts(value *AwsSpotFleetRequest_TimeoutsProperty)
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
	ResetAllocationStrategy()
	// Experimental.
	ResetContext()
	// Experimental.
	ResetExcessCapacityTerminationPolicy()
	// Experimental.
	ResetFleetType()
	// Experimental.
	ResetId()
	// Experimental.
	ResetInstanceInterruptionBehaviour()
	// Experimental.
	ResetInstancePoolsToUseCount()
	// Experimental.
	ResetLaunchSpecification()
	// Experimental.
	ResetLaunchTemplateConfig()
	// Experimental.
	ResetLoadBalancers()
	// Experimental.
	ResetOnDemandAllocationStrategy()
	// Experimental.
	ResetOnDemandMaxTotalPrice()
	// Experimental.
	ResetOnDemandTargetCapacity()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetReplaceUnhealthyInstances()
	// Experimental.
	ResetSpotMaintenanceStrategies()
	// Experimental.
	ResetSpotPrice()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTargetCapacityUnitType()
	// Experimental.
	ResetTargetGroupArns()
	// Experimental.
	ResetTerminateInstancesOnDelete()
	// Experimental.
	ResetTerminateInstancesWithExpiration()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetValidFrom()
	// Experimental.
	ResetValidUntil()
	// Experimental.
	ResetWaitForFulfillment()
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

// The jsii proxy struct for AwsSpotFleetRequest
type jsiiProxy_AwsSpotFleetRequest struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsSpotFleetRequest) AllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) AllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) ExcessCapacityTerminationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excessCapacityTerminationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) ExcessCapacityTerminationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excessCapacityTerminationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) FleetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) FleetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) IamFleetRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamFleetRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) IamFleetRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamFleetRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) InstanceInterruptionBehaviour() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInterruptionBehaviour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) InstanceInterruptionBehaviourInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInterruptionBehaviourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) InstancePoolsToUseCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancePoolsToUseCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) InstancePoolsToUseCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancePoolsToUseCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) LaunchSpecification() AwsSpotFleetRequest_LaunchSpecificationPropertyList {
	var returns AwsSpotFleetRequest_LaunchSpecificationPropertyList
	_jsii_.Get(
		j,
		"launchSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) LaunchSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) LaunchTemplateConfig() AwsSpotFleetRequest_LaunchTemplateConfigPropertyList {
	var returns AwsSpotFleetRequest_LaunchTemplateConfigPropertyList
	_jsii_.Get(
		j,
		"launchTemplateConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) LaunchTemplateConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) LoadBalancers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) LoadBalancersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) OnDemandAllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandAllocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) OnDemandAllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandAllocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) OnDemandMaxTotalPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandMaxTotalPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) OnDemandMaxTotalPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onDemandMaxTotalPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) OnDemandTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) OnDemandTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"onDemandTargetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) ReplaceUnhealthyInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceUnhealthyInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) ReplaceUnhealthyInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceUnhealthyInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) SpotMaintenanceStrategies() AwsSpotFleetRequest_SpotMaintenanceStrategiesPropertyOutputReference {
	var returns AwsSpotFleetRequest_SpotMaintenanceStrategiesPropertyOutputReference
	_jsii_.Get(
		j,
		"spotMaintenanceStrategies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) SpotMaintenanceStrategiesInput() *AwsSpotFleetRequest_SpotMaintenanceStrategiesProperty {
	var returns *AwsSpotFleetRequest_SpotMaintenanceStrategiesProperty
	_jsii_.Get(
		j,
		"spotMaintenanceStrategiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) SpotPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) SpotPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) SpotRequestState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotRequestState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TargetCapacityUnitType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacityUnitType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TargetCapacityUnitTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCapacityUnitTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TargetGroupArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TargetGroupArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetGroupArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TerminateInstancesOnDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminateInstancesOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TerminateInstancesOnDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminateInstancesOnDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TerminateInstancesWithExpiration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstancesWithExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TerminateInstancesWithExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstancesWithExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) Timeouts() AwsSpotFleetRequest_TimeoutsPropertyOutputReference {
	var returns AwsSpotFleetRequest_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) ValidFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) ValidFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) ValidUntil() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntil",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) ValidUntilInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntilInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) WaitForFulfillment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForFulfillment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSpotFleetRequest) WaitForFulfillmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForFulfillmentInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request aws_spot_fleet_request} Resource.
// Experimental.
func NewAwsSpotFleetRequest(scope constructs.Construct, id *string, config *AwsSpotFleetRequestConfig) AwsSpotFleetRequest {
	_init_.Initialize()

	if err := validateNewAwsSpotFleetRequestParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSpotFleetRequest{}

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsSpotFleetRequest",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request aws_spot_fleet_request} Resource.
// Experimental.
func NewAwsSpotFleetRequest_Override(a AwsSpotFleetRequest, scope constructs.Construct, id *string, config *AwsSpotFleetRequestConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsSpotFleetRequest",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetAllocationStrategy(val *string) {
	if err := j.validateSetAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationStrategy",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetContext(val *string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetExcessCapacityTerminationPolicy(val *string) {
	if err := j.validateSetExcessCapacityTerminationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excessCapacityTerminationPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetFleetType(val *string) {
	if err := j.validateSetFleetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fleetType",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetIamFleetRole(val *string) {
	if err := j.validateSetIamFleetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamFleetRole",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetInstanceInterruptionBehaviour(val *string) {
	if err := j.validateSetInstanceInterruptionBehaviourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInterruptionBehaviour",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetInstancePoolsToUseCount(val *float64) {
	if err := j.validateSetInstancePoolsToUseCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePoolsToUseCount",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetLoadBalancers(val *[]*string) {
	if err := j.validateSetLoadBalancersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancers",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetOnDemandAllocationStrategy(val *string) {
	if err := j.validateSetOnDemandAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandAllocationStrategy",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetOnDemandMaxTotalPrice(val *string) {
	if err := j.validateSetOnDemandMaxTotalPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandMaxTotalPrice",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetOnDemandTargetCapacity(val *float64) {
	if err := j.validateSetOnDemandTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onDemandTargetCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetReplaceUnhealthyInstances(val interface{}) {
	if err := j.validateSetReplaceUnhealthyInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceUnhealthyInstances",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetSpotPrice(val *string) {
	if err := j.validateSetSpotPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotPrice",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetTargetCapacity(val *float64) {
	if err := j.validateSetTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetTargetCapacityUnitType(val *string) {
	if err := j.validateSetTargetCapacityUnitTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCapacityUnitType",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetTargetGroupArns(val *[]*string) {
	if err := j.validateSetTargetGroupArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetGroupArns",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetTerminateInstancesOnDelete(val *string) {
	if err := j.validateSetTerminateInstancesOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminateInstancesOnDelete",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetTerminateInstancesWithExpiration(val interface{}) {
	if err := j.validateSetTerminateInstancesWithExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminateInstancesWithExpiration",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetValidFrom(val *string) {
	if err := j.validateSetValidFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validFrom",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetValidUntil(val *string) {
	if err := j.validateSetValidUntilParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validUntil",
		val,
	)
}

func (j *jsiiProxy_AwsSpotFleetRequest)SetWaitForFulfillment(val interface{}) {
	if err := j.validateSetWaitForFulfillmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForFulfillment",
		val,
	)
}

// Generates CDKTN code for importing a AwsSpotFleetRequest resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsSpotFleetRequest_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsSpotFleetRequest_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsSpotFleetRequest",
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
func AwsSpotFleetRequest_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSpotFleetRequest_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsSpotFleetRequest",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsSpotFleetRequest_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSpotFleetRequest_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsSpotFleetRequest",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsSpotFleetRequest_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSpotFleetRequest_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsSpotFleetRequest",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsSpotFleetRequest_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ec2.AwsSpotFleetRequest",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsSpotFleetRequest) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSpotFleetRequest) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSpotFleetRequest) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSpotFleetRequest) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSpotFleetRequest) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSpotFleetRequest) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSpotFleetRequest) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSpotFleetRequest) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSpotFleetRequest) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSpotFleetRequest) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotFleetRequest) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSpotFleetRequest) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsSpotFleetRequest) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) PutLaunchSpecification(value interface{}) {
	if err := a.validatePutLaunchSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLaunchSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) PutLaunchTemplateConfig(value interface{}) {
	if err := a.validatePutLaunchTemplateConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLaunchTemplateConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) PutSpotMaintenanceStrategies(value *AwsSpotFleetRequest_SpotMaintenanceStrategiesProperty) {
	if err := a.validatePutSpotMaintenanceStrategiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSpotMaintenanceStrategies",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) PutTimeouts(value *AwsSpotFleetRequest_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetAllocationStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetAllocationStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetContext() {
	_jsii_.InvokeVoid(
		a,
		"resetContext",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetExcessCapacityTerminationPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetExcessCapacityTerminationPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetFleetType() {
	_jsii_.InvokeVoid(
		a,
		"resetFleetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetInstanceInterruptionBehaviour() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceInterruptionBehaviour",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetInstancePoolsToUseCount() {
	_jsii_.InvokeVoid(
		a,
		"resetInstancePoolsToUseCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetLaunchSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetLaunchTemplateConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLaunchTemplateConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetLoadBalancers() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetOnDemandAllocationStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandAllocationStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetOnDemandMaxTotalPrice() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandMaxTotalPrice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetOnDemandTargetCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandTargetCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetReplaceUnhealthyInstances() {
	_jsii_.InvokeVoid(
		a,
		"resetReplaceUnhealthyInstances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetSpotMaintenanceStrategies() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotMaintenanceStrategies",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetSpotPrice() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotPrice",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetTargetCapacityUnitType() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetCapacityUnitType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetTargetGroupArns() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetGroupArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetTerminateInstancesOnDelete() {
	_jsii_.InvokeVoid(
		a,
		"resetTerminateInstancesOnDelete",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetTerminateInstancesWithExpiration() {
	_jsii_.InvokeVoid(
		a,
		"resetTerminateInstancesWithExpiration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetValidFrom() {
	_jsii_.InvokeVoid(
		a,
		"resetValidFrom",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetValidUntil() {
	_jsii_.InvokeVoid(
		a,
		"resetValidUntil",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) ResetWaitForFulfillment() {
	_jsii_.InvokeVoid(
		a,
		"resetWaitForFulfillment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSpotFleetRequest) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotFleetRequest) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotFleetRequest) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotFleetRequest) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotFleetRequest) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotFleetRequest) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSpotFleetRequest) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

