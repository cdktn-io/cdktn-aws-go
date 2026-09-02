package awselbclassic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselbclassic/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awselbclassic/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb aws_elb}.
// Experimental.
type TfElb interface {
	cdktn.TerraformResource
	// Experimental.
	AccessLogs() TfElb_AccessLogsPropertyOutputReference
	// Experimental.
	AccessLogsInput() *TfElb_AccessLogsProperty
	// Experimental.
	Arn() *string
	// Experimental.
	AvailabilityZones() *[]*string
	// Experimental.
	SetAvailabilityZones(val *[]*string)
	// Experimental.
	AvailabilityZonesInput() *[]*string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConnectionDraining() interface{}
	// Experimental.
	SetConnectionDraining(val interface{})
	// Experimental.
	ConnectionDrainingInput() interface{}
	// Experimental.
	ConnectionDrainingTimeout() *float64
	// Experimental.
	SetConnectionDrainingTimeout(val *float64)
	// Experimental.
	ConnectionDrainingTimeoutInput() *float64
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CrossZoneLoadBalancing() interface{}
	// Experimental.
	SetCrossZoneLoadBalancing(val interface{})
	// Experimental.
	CrossZoneLoadBalancingInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DesyncMitigationMode() *string
	// Experimental.
	SetDesyncMitigationMode(val *string)
	// Experimental.
	DesyncMitigationModeInput() *string
	// Experimental.
	DnsName() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HealthCheck() TfElb_HealthCheckPropertyOutputReference
	// Experimental.
	HealthCheckInput() *TfElb_HealthCheckProperty
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	IdleTimeout() *float64
	// Experimental.
	SetIdleTimeout(val *float64)
	// Experimental.
	IdleTimeoutInput() *float64
	// Experimental.
	Instances() *[]*string
	// Experimental.
	SetInstances(val *[]*string)
	// Experimental.
	InstancesInput() *[]*string
	// Experimental.
	Internal() interface{}
	// Experimental.
	SetInternal(val interface{})
	// Experimental.
	InternalInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	Listener() TfElb_ListenerPropertyList
	// Experimental.
	ListenerInput() interface{}
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
	SecurityGroups() *[]*string
	// Experimental.
	SetSecurityGroups(val *[]*string)
	// Experimental.
	SecurityGroupsInput() *[]*string
	// Experimental.
	SourceSecurityGroup() *string
	// Experimental.
	SetSourceSecurityGroup(val *string)
	// Experimental.
	SourceSecurityGroupId() *string
	// Experimental.
	SourceSecurityGroupInput() *string
	// Experimental.
	Subnets() *[]*string
	// Experimental.
	SetSubnets(val *[]*string)
	// Experimental.
	SubnetsInput() *[]*string
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
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() TfElb_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	ZoneId() *string
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
	PutAccessLogs(value *TfElb_AccessLogsProperty)
	// Experimental.
	PutHealthCheck(value *TfElb_HealthCheckProperty)
	// Experimental.
	PutListener(value interface{})
	// Experimental.
	PutTimeouts(value *TfElb_TimeoutsProperty)
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
	ResetAccessLogs()
	// Experimental.
	ResetAvailabilityZones()
	// Experimental.
	ResetConnectionDraining()
	// Experimental.
	ResetConnectionDrainingTimeout()
	// Experimental.
	ResetCrossZoneLoadBalancing()
	// Experimental.
	ResetDesyncMitigationMode()
	// Experimental.
	ResetHealthCheck()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIdleTimeout()
	// Experimental.
	ResetInstances()
	// Experimental.
	ResetInternal()
	// Experimental.
	ResetName()
	// Experimental.
	ResetNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSecurityGroups()
	// Experimental.
	ResetSourceSecurityGroup()
	// Experimental.
	ResetSubnets()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
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

// The jsii proxy struct for TfElb
type jsiiProxy_TfElb struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfElb) AccessLogs() TfElb_AccessLogsPropertyOutputReference {
	var returns TfElb_AccessLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"accessLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) AccessLogsInput() *TfElb_AccessLogsProperty {
	var returns *TfElb_AccessLogsProperty
	_jsii_.Get(
		j,
		"accessLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) ConnectionDraining() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionDraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) ConnectionDrainingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionDrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) ConnectionDrainingTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionDrainingTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) ConnectionDrainingTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionDrainingTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) CrossZoneLoadBalancing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossZoneLoadBalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) CrossZoneLoadBalancingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossZoneLoadBalancingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) DesyncMitigationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desyncMitigationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) DesyncMitigationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desyncMitigationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) HealthCheck() TfElb_HealthCheckPropertyOutputReference {
	var returns TfElb_HealthCheckPropertyOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) HealthCheckInput() *TfElb_HealthCheckProperty {
	var returns *TfElb_HealthCheckProperty
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) IdleTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) IdleTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Instances() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) InstancesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Internal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) InternalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Listener() TfElb_ListenerPropertyList {
	var returns TfElb_ListenerPropertyList
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) ListenerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"listenerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) SourceSecurityGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) SourceSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) SourceSecurityGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) SubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) Timeouts() TfElb_TimeoutsPropertyOutputReference {
	var returns TfElb_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfElb) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb aws_elb} Resource.
// Experimental.
func NewTfElb(scope constructs.Construct, id *string, config *TfElbConfig) TfElb {
	_init_.Initialize()

	if err := validateNewTfElbParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfElb{}

	_jsii_.Create(
		"@cdktn/aws-elb-classic.TfElb",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb aws_elb} Resource.
// Experimental.
func NewTfElb_Override(t TfElb, scope constructs.Construct, id *string, config *TfElbConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb-classic.TfElb",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfElb)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetConnectionDraining(val interface{}) {
	if err := j.validateSetConnectionDrainingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionDraining",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetConnectionDrainingTimeout(val *float64) {
	if err := j.validateSetConnectionDrainingTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionDrainingTimeout",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetCrossZoneLoadBalancing(val interface{}) {
	if err := j.validateSetCrossZoneLoadBalancingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossZoneLoadBalancing",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetDesyncMitigationMode(val *string) {
	if err := j.validateSetDesyncMitigationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desyncMitigationMode",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetIdleTimeout(val *float64) {
	if err := j.validateSetIdleTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTimeout",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetInstances(val *[]*string) {
	if err := j.validateSetInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instances",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetInternal(val interface{}) {
	if err := j.validateSetInternalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internal",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetSourceSecurityGroup(val *string) {
	if err := j.validateSetSourceSecurityGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetSubnets(val *[]*string) {
	if err := j.validateSetSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfElb)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a TfElb resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfElb_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfElb_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb-classic.TfElb",
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
func TfElb_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfElb_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb-classic.TfElb",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfElb_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfElb_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb-classic.TfElb",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfElb_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfElb_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb-classic.TfElb",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfElb_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-elb-classic.TfElb",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfElb) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfElb) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfElb) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfElb) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfElb) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfElb) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfElb) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfElb) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfElb) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfElb) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfElb) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfElb) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfElb) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfElb) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfElb) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfElb) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfElb) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfElb) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfElb) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfElb) PutAccessLogs(value *TfElb_AccessLogsProperty) {
	if err := t.validatePutAccessLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAccessLogs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfElb) PutHealthCheck(value *TfElb_HealthCheckProperty) {
	if err := t.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfElb) PutListener(value interface{}) {
	if err := t.validatePutListenerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putListener",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfElb) PutTimeouts(value *TfElb_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfElb) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfElb) ResetAccessLogs() {
	_jsii_.InvokeVoid(
		t,
		"resetAccessLogs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetAvailabilityZones() {
	_jsii_.InvokeVoid(
		t,
		"resetAvailabilityZones",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetConnectionDraining() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectionDraining",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetConnectionDrainingTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectionDrainingTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetCrossZoneLoadBalancing() {
	_jsii_.InvokeVoid(
		t,
		"resetCrossZoneLoadBalancing",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetDesyncMitigationMode() {
	_jsii_.InvokeVoid(
		t,
		"resetDesyncMitigationMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		t,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetIdleTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetIdleTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetInstances() {
	_jsii_.InvokeVoid(
		t,
		"resetInstances",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetInternal() {
	_jsii_.InvokeVoid(
		t,
		"resetInternal",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetSourceSecurityGroup() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceSecurityGroup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetSubnets() {
	_jsii_.InvokeVoid(
		t,
		"resetSubnets",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfElb) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfElb) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfElb) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfElb) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfElb) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfElb) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfElb) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

