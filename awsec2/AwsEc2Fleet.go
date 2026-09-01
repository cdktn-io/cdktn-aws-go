package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet aws_ec2_fleet}.
// Experimental.
type AwsEc2Fleet interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
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
	FleetInstanceSet() AwsEc2Fleet_FleetInstanceSetPropertyList
	// Experimental.
	FleetInstanceSetInput() interface{}
	// Experimental.
	FleetState() *string
	// Experimental.
	SetFleetState(val *string)
	// Experimental.
	FleetStateInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	FulfilledCapacity() *float64
	// Experimental.
	SetFulfilledCapacity(val *float64)
	// Experimental.
	FulfilledCapacityInput() *float64
	// Experimental.
	FulfilledOnDemandCapacity() *float64
	// Experimental.
	SetFulfilledOnDemandCapacity(val *float64)
	// Experimental.
	FulfilledOnDemandCapacityInput() *float64
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	LaunchTemplateConfig() AwsEc2Fleet_LaunchTemplateConfigPropertyList
	// Experimental.
	LaunchTemplateConfigInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OnDemandOptions() AwsEc2Fleet_OnDemandOptionsPropertyOutputReference
	// Experimental.
	OnDemandOptionsInput() *AwsEc2Fleet_OnDemandOptionsProperty
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
	SpotOptions() AwsEc2Fleet_SpotOptionsPropertyOutputReference
	// Experimental.
	SpotOptionsInput() *AwsEc2Fleet_SpotOptionsProperty
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
	TargetCapacitySpecification() AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference
	// Experimental.
	TargetCapacitySpecificationInput() *AwsEc2Fleet_TargetCapacitySpecificationProperty
	// Experimental.
	TerminateInstances() interface{}
	// Experimental.
	SetTerminateInstances(val interface{})
	// Experimental.
	TerminateInstancesInput() interface{}
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
	Timeouts() AwsEc2Fleet_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	PutFleetInstanceSet(value interface{})
	// Experimental.
	PutLaunchTemplateConfig(value interface{})
	// Experimental.
	PutOnDemandOptions(value *AwsEc2Fleet_OnDemandOptionsProperty)
	// Experimental.
	PutSpotOptions(value *AwsEc2Fleet_SpotOptionsProperty)
	// Experimental.
	PutTargetCapacitySpecification(value *AwsEc2Fleet_TargetCapacitySpecificationProperty)
	// Experimental.
	PutTimeouts(value *AwsEc2Fleet_TimeoutsProperty)
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
	ResetContext()
	// Experimental.
	ResetExcessCapacityTerminationPolicy()
	// Experimental.
	ResetFleetInstanceSet()
	// Experimental.
	ResetFleetState()
	// Experimental.
	ResetFulfilledCapacity()
	// Experimental.
	ResetFulfilledOnDemandCapacity()
	// Experimental.
	ResetId()
	// Experimental.
	ResetOnDemandOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetReplaceUnhealthyInstances()
	// Experimental.
	ResetSpotOptions()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTerminateInstances()
	// Experimental.
	ResetTerminateInstancesWithExpiration()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetType()
	// Experimental.
	ResetValidFrom()
	// Experimental.
	ResetValidUntil()
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

// The jsii proxy struct for AwsEc2Fleet
type jsiiProxy_AwsEc2Fleet struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsEc2Fleet) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) ExcessCapacityTerminationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excessCapacityTerminationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) ExcessCapacityTerminationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excessCapacityTerminationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) FleetInstanceSet() AwsEc2Fleet_FleetInstanceSetPropertyList {
	var returns AwsEc2Fleet_FleetInstanceSetPropertyList
	_jsii_.Get(
		j,
		"fleetInstanceSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) FleetInstanceSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fleetInstanceSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) FleetState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) FleetStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) FulfilledCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fulfilledCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) FulfilledCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fulfilledCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) FulfilledOnDemandCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fulfilledOnDemandCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) FulfilledOnDemandCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fulfilledOnDemandCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) LaunchTemplateConfig() AwsEc2Fleet_LaunchTemplateConfigPropertyList {
	var returns AwsEc2Fleet_LaunchTemplateConfigPropertyList
	_jsii_.Get(
		j,
		"launchTemplateConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) LaunchTemplateConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) OnDemandOptions() AwsEc2Fleet_OnDemandOptionsPropertyOutputReference {
	var returns AwsEc2Fleet_OnDemandOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"onDemandOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) OnDemandOptionsInput() *AwsEc2Fleet_OnDemandOptionsProperty {
	var returns *AwsEc2Fleet_OnDemandOptionsProperty
	_jsii_.Get(
		j,
		"onDemandOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) ReplaceUnhealthyInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceUnhealthyInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) ReplaceUnhealthyInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceUnhealthyInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) SpotOptions() AwsEc2Fleet_SpotOptionsPropertyOutputReference {
	var returns AwsEc2Fleet_SpotOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"spotOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) SpotOptionsInput() *AwsEc2Fleet_SpotOptionsProperty {
	var returns *AwsEc2Fleet_SpotOptionsProperty
	_jsii_.Get(
		j,
		"spotOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) TargetCapacitySpecification() AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference {
	var returns AwsEc2Fleet_TargetCapacitySpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"targetCapacitySpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) TargetCapacitySpecificationInput() *AwsEc2Fleet_TargetCapacitySpecificationProperty {
	var returns *AwsEc2Fleet_TargetCapacitySpecificationProperty
	_jsii_.Get(
		j,
		"targetCapacitySpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) TerminateInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) TerminateInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) TerminateInstancesWithExpiration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstancesWithExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) TerminateInstancesWithExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstancesWithExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) Timeouts() AwsEc2Fleet_TimeoutsPropertyOutputReference {
	var returns AwsEc2Fleet_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) ValidFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) ValidFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) ValidUntil() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntil",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2Fleet) ValidUntilInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntilInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet aws_ec2_fleet} Resource.
// Experimental.
func NewAwsEc2Fleet(scope constructs.Construct, id *string, config *AwsEc2FleetConfig) AwsEc2Fleet {
	_init_.Initialize()

	if err := validateNewAwsEc2FleetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEc2Fleet{}

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsEc2Fleet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet aws_ec2_fleet} Resource.
// Experimental.
func NewAwsEc2Fleet_Override(a AwsEc2Fleet, scope constructs.Construct, id *string, config *AwsEc2FleetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.AwsEc2Fleet",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetContext(val *string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetExcessCapacityTerminationPolicy(val *string) {
	if err := j.validateSetExcessCapacityTerminationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excessCapacityTerminationPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetFleetState(val *string) {
	if err := j.validateSetFleetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fleetState",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetFulfilledCapacity(val *float64) {
	if err := j.validateSetFulfilledCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fulfilledCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetFulfilledOnDemandCapacity(val *float64) {
	if err := j.validateSetFulfilledOnDemandCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fulfilledOnDemandCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetReplaceUnhealthyInstances(val interface{}) {
	if err := j.validateSetReplaceUnhealthyInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceUnhealthyInstances",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetTerminateInstances(val interface{}) {
	if err := j.validateSetTerminateInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminateInstances",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetTerminateInstancesWithExpiration(val interface{}) {
	if err := j.validateSetTerminateInstancesWithExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminateInstancesWithExpiration",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetValidFrom(val *string) {
	if err := j.validateSetValidFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validFrom",
		val,
	)
}

func (j *jsiiProxy_AwsEc2Fleet)SetValidUntil(val *string) {
	if err := j.validateSetValidUntilParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validUntil",
		val,
	)
}

// Generates CDKTN code for importing a AwsEc2Fleet resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsEc2Fleet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsEc2Fleet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsEc2Fleet",
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
func AwsEc2Fleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEc2Fleet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsEc2Fleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsEc2Fleet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEc2Fleet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsEc2Fleet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsEc2Fleet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEc2Fleet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.AwsEc2Fleet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsEc2Fleet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ec2.AwsEc2Fleet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsEc2Fleet) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsEc2Fleet) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsEc2Fleet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEc2Fleet) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2Fleet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEc2Fleet) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEc2Fleet) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEc2Fleet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEc2Fleet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEc2Fleet) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEc2Fleet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEc2Fleet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2Fleet) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsEc2Fleet) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2Fleet) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsEc2Fleet) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsEc2Fleet) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsEc2Fleet) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsEc2Fleet) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsEc2Fleet) PutFleetInstanceSet(value interface{}) {
	if err := a.validatePutFleetInstanceSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFleetInstanceSet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2Fleet) PutLaunchTemplateConfig(value interface{}) {
	if err := a.validatePutLaunchTemplateConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLaunchTemplateConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2Fleet) PutOnDemandOptions(value *AwsEc2Fleet_OnDemandOptionsProperty) {
	if err := a.validatePutOnDemandOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOnDemandOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2Fleet) PutSpotOptions(value *AwsEc2Fleet_SpotOptionsProperty) {
	if err := a.validatePutSpotOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSpotOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2Fleet) PutTargetCapacitySpecification(value *AwsEc2Fleet_TargetCapacitySpecificationProperty) {
	if err := a.validatePutTargetCapacitySpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTargetCapacitySpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2Fleet) PutTimeouts(value *AwsEc2Fleet_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2Fleet) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetContext() {
	_jsii_.InvokeVoid(
		a,
		"resetContext",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetExcessCapacityTerminationPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetExcessCapacityTerminationPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetFleetInstanceSet() {
	_jsii_.InvokeVoid(
		a,
		"resetFleetInstanceSet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetFleetState() {
	_jsii_.InvokeVoid(
		a,
		"resetFleetState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetFulfilledCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetFulfilledCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetFulfilledOnDemandCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetFulfilledOnDemandCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetOnDemandOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetReplaceUnhealthyInstances() {
	_jsii_.InvokeVoid(
		a,
		"resetReplaceUnhealthyInstances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetSpotOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetSpotOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetTerminateInstances() {
	_jsii_.InvokeVoid(
		a,
		"resetTerminateInstances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetTerminateInstancesWithExpiration() {
	_jsii_.InvokeVoid(
		a,
		"resetTerminateInstancesWithExpiration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetValidFrom() {
	_jsii_.InvokeVoid(
		a,
		"resetValidFrom",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) ResetValidUntil() {
	_jsii_.InvokeVoid(
		a,
		"resetValidUntil",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2Fleet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2Fleet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2Fleet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2Fleet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2Fleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2Fleet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2Fleet) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

