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
type TfFleet interface {
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
	FleetInstanceSet() TfFleet_FleetInstanceSetPropertyList
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
	LaunchTemplateConfig() TfFleet_LaunchTemplateConfigPropertyList
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
	OnDemandOptions() TfFleet_OnDemandOptionsPropertyOutputReference
	// Experimental.
	OnDemandOptionsInput() *TfFleet_OnDemandOptionsProperty
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
	SpotOptions() TfFleet_SpotOptionsPropertyOutputReference
	// Experimental.
	SpotOptionsInput() *TfFleet_SpotOptionsProperty
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
	TargetCapacitySpecification() TfFleet_TargetCapacitySpecificationPropertyOutputReference
	// Experimental.
	TargetCapacitySpecificationInput() *TfFleet_TargetCapacitySpecificationProperty
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
	Timeouts() TfFleet_TimeoutsPropertyOutputReference
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
	PutOnDemandOptions(value *TfFleet_OnDemandOptionsProperty)
	// Experimental.
	PutSpotOptions(value *TfFleet_SpotOptionsProperty)
	// Experimental.
	PutTargetCapacitySpecification(value *TfFleet_TargetCapacitySpecificationProperty)
	// Experimental.
	PutTimeouts(value *TfFleet_TimeoutsProperty)
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

// The jsii proxy struct for TfFleet
type jsiiProxy_TfFleet struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfFleet) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) ExcessCapacityTerminationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excessCapacityTerminationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) ExcessCapacityTerminationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excessCapacityTerminationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) FleetInstanceSet() TfFleet_FleetInstanceSetPropertyList {
	var returns TfFleet_FleetInstanceSetPropertyList
	_jsii_.Get(
		j,
		"fleetInstanceSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) FleetInstanceSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fleetInstanceSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) FleetState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) FleetStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) FulfilledCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fulfilledCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) FulfilledCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fulfilledCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) FulfilledOnDemandCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fulfilledOnDemandCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) FulfilledOnDemandCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fulfilledOnDemandCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) LaunchTemplateConfig() TfFleet_LaunchTemplateConfigPropertyList {
	var returns TfFleet_LaunchTemplateConfigPropertyList
	_jsii_.Get(
		j,
		"launchTemplateConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) LaunchTemplateConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) OnDemandOptions() TfFleet_OnDemandOptionsPropertyOutputReference {
	var returns TfFleet_OnDemandOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"onDemandOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) OnDemandOptionsInput() *TfFleet_OnDemandOptionsProperty {
	var returns *TfFleet_OnDemandOptionsProperty
	_jsii_.Get(
		j,
		"onDemandOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) ReplaceUnhealthyInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceUnhealthyInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) ReplaceUnhealthyInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceUnhealthyInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) SpotOptions() TfFleet_SpotOptionsPropertyOutputReference {
	var returns TfFleet_SpotOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"spotOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) SpotOptionsInput() *TfFleet_SpotOptionsProperty {
	var returns *TfFleet_SpotOptionsProperty
	_jsii_.Get(
		j,
		"spotOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) TargetCapacitySpecification() TfFleet_TargetCapacitySpecificationPropertyOutputReference {
	var returns TfFleet_TargetCapacitySpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"targetCapacitySpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) TargetCapacitySpecificationInput() *TfFleet_TargetCapacitySpecificationProperty {
	var returns *TfFleet_TargetCapacitySpecificationProperty
	_jsii_.Get(
		j,
		"targetCapacitySpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) TerminateInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) TerminateInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) TerminateInstancesWithExpiration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstancesWithExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) TerminateInstancesWithExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstancesWithExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) Timeouts() TfFleet_TimeoutsPropertyOutputReference {
	var returns TfFleet_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) ValidFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) ValidFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) ValidUntil() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntil",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet) ValidUntilInput() *string {
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
func NewTfFleet(scope constructs.Construct, id *string, config *TfFleetConfig) TfFleet {
	_init_.Initialize()

	if err := validateNewTfFleetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFleet{}

	_jsii_.Create(
		"@cdktn/aws-ec2.TfFleet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet aws_ec2_fleet} Resource.
// Experimental.
func NewTfFleet_Override(t TfFleet, scope constructs.Construct, id *string, config *TfFleetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.TfFleet",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfFleet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetContext(val *string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetExcessCapacityTerminationPolicy(val *string) {
	if err := j.validateSetExcessCapacityTerminationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excessCapacityTerminationPolicy",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetFleetState(val *string) {
	if err := j.validateSetFleetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fleetState",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetFulfilledCapacity(val *float64) {
	if err := j.validateSetFulfilledCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fulfilledCapacity",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetFulfilledOnDemandCapacity(val *float64) {
	if err := j.validateSetFulfilledOnDemandCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fulfilledOnDemandCapacity",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetReplaceUnhealthyInstances(val interface{}) {
	if err := j.validateSetReplaceUnhealthyInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceUnhealthyInstances",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetTerminateInstances(val interface{}) {
	if err := j.validateSetTerminateInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminateInstances",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetTerminateInstancesWithExpiration(val interface{}) {
	if err := j.validateSetTerminateInstancesWithExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminateInstancesWithExpiration",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetValidFrom(val *string) {
	if err := j.validateSetValidFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validFrom",
		val,
	)
}

func (j *jsiiProxy_TfFleet)SetValidUntil(val *string) {
	if err := j.validateSetValidUntilParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validUntil",
		val,
	)
}

// Generates CDKTN code for importing a TfFleet resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfFleet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfFleet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.TfFleet",
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
func TfFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfFleet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.TfFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfFleet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfFleet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.TfFleet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfFleet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfFleet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2.TfFleet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfFleet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ec2.TfFleet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfFleet) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfFleet) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfFleet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFleet) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFleet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFleet) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFleet) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFleet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFleet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFleet) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFleet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFleet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFleet) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfFleet) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFleet) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfFleet) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfFleet) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfFleet) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfFleet) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfFleet) PutFleetInstanceSet(value interface{}) {
	if err := t.validatePutFleetInstanceSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFleetInstanceSet",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFleet) PutLaunchTemplateConfig(value interface{}) {
	if err := t.validatePutLaunchTemplateConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLaunchTemplateConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFleet) PutOnDemandOptions(value *TfFleet_OnDemandOptionsProperty) {
	if err := t.validatePutOnDemandOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOnDemandOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFleet) PutSpotOptions(value *TfFleet_SpotOptionsProperty) {
	if err := t.validatePutSpotOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSpotOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFleet) PutTargetCapacitySpecification(value *TfFleet_TargetCapacitySpecificationProperty) {
	if err := t.validatePutTargetCapacitySpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTargetCapacitySpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFleet) PutTimeouts(value *TfFleet_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFleet) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfFleet) ResetContext() {
	_jsii_.InvokeVoid(
		t,
		"resetContext",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetExcessCapacityTerminationPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetExcessCapacityTerminationPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetFleetInstanceSet() {
	_jsii_.InvokeVoid(
		t,
		"resetFleetInstanceSet",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetFleetState() {
	_jsii_.InvokeVoid(
		t,
		"resetFleetState",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetFulfilledCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetFulfilledCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetFulfilledOnDemandCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetFulfilledOnDemandCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetOnDemandOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetOnDemandOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetReplaceUnhealthyInstances() {
	_jsii_.InvokeVoid(
		t,
		"resetReplaceUnhealthyInstances",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetSpotOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetSpotOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetTerminateInstances() {
	_jsii_.InvokeVoid(
		t,
		"resetTerminateInstances",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetTerminateInstancesWithExpiration() {
	_jsii_.InvokeVoid(
		t,
		"resetTerminateInstancesWithExpiration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetType() {
	_jsii_.InvokeVoid(
		t,
		"resetType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetValidFrom() {
	_jsii_.InvokeVoid(
		t,
		"resetValidFrom",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) ResetValidUntil() {
	_jsii_.InvokeVoid(
		t,
		"resetValidUntil",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFleet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFleet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFleet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFleet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFleet) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

