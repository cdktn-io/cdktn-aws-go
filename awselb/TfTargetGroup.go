package awselb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselb/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awselb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group aws_lb_target_group}.
// Experimental.
type TfTargetGroup interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	ArnSuffix() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConnectionTermination() interface{}
	// Experimental.
	SetConnectionTermination(val interface{})
	// Experimental.
	ConnectionTerminationInput() interface{}
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
	DeregistrationDelay() *string
	// Experimental.
	SetDeregistrationDelay(val *string)
	// Experimental.
	DeregistrationDelayInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HealthCheck() TfTargetGroup_HealthCheckPropertyOutputReference
	// Experimental.
	HealthCheckInput() *TfTargetGroup_HealthCheckProperty
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	IpAddressType() *string
	// Experimental.
	SetIpAddressType(val *string)
	// Experimental.
	IpAddressTypeInput() *string
	// Experimental.
	LambdaMultiValueHeadersEnabled() interface{}
	// Experimental.
	SetLambdaMultiValueHeadersEnabled(val interface{})
	// Experimental.
	LambdaMultiValueHeadersEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LoadBalancerArns() *[]*string
	// Experimental.
	LoadBalancingAlgorithmType() *string
	// Experimental.
	SetLoadBalancingAlgorithmType(val *string)
	// Experimental.
	LoadBalancingAlgorithmTypeInput() *string
	// Experimental.
	LoadBalancingAnomalyMitigation() *string
	// Experimental.
	SetLoadBalancingAnomalyMitigation(val *string)
	// Experimental.
	LoadBalancingAnomalyMitigationInput() *string
	// Experimental.
	LoadBalancingCrossZoneEnabled() *string
	// Experimental.
	SetLoadBalancingCrossZoneEnabled(val *string)
	// Experimental.
	LoadBalancingCrossZoneEnabledInput() *string
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
	Port() *float64
	// Experimental.
	SetPort(val *float64)
	// Experimental.
	PortInput() *float64
	// Experimental.
	PreserveClientIp() *string
	// Experimental.
	SetPreserveClientIp(val *string)
	// Experimental.
	PreserveClientIpInput() *string
	// Experimental.
	Protocol() *string
	// Experimental.
	SetProtocol(val *string)
	// Experimental.
	ProtocolInput() *string
	// Experimental.
	ProtocolVersion() *string
	// Experimental.
	SetProtocolVersion(val *string)
	// Experimental.
	ProtocolVersionInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	ProxyProtocolV2() interface{}
	// Experimental.
	SetProxyProtocolV2(val interface{})
	// Experimental.
	ProxyProtocolV2Input() interface{}
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	SlowStart() *float64
	// Experimental.
	SetSlowStart(val *float64)
	// Experimental.
	SlowStartInput() *float64
	// Experimental.
	Stickiness() TfTargetGroup_StickinessPropertyOutputReference
	// Experimental.
	StickinessInput() *TfTargetGroup_StickinessProperty
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
	TargetControlPort() *float64
	// Experimental.
	SetTargetControlPort(val *float64)
	// Experimental.
	TargetControlPortInput() *float64
	// Experimental.
	TargetFailover() TfTargetGroup_TargetFailoverPropertyList
	// Experimental.
	TargetFailoverInput() interface{}
	// Experimental.
	TargetGroupHealth() TfTargetGroup_TargetGroupHealthPropertyOutputReference
	// Experimental.
	TargetGroupHealthInput() *TfTargetGroup_TargetGroupHealthProperty
	// Experimental.
	TargetHealthState() TfTargetGroup_TargetHealthStatePropertyList
	// Experimental.
	TargetHealthStateInput() interface{}
	// Experimental.
	TargetType() *string
	// Experimental.
	SetTargetType(val *string)
	// Experimental.
	TargetTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	VpcId() *string
	// Experimental.
	SetVpcId(val *string)
	// Experimental.
	VpcIdInput() *string
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
	PutHealthCheck(value *TfTargetGroup_HealthCheckProperty)
	// Experimental.
	PutStickiness(value *TfTargetGroup_StickinessProperty)
	// Experimental.
	PutTargetFailover(value interface{})
	// Experimental.
	PutTargetGroupHealth(value *TfTargetGroup_TargetGroupHealthProperty)
	// Experimental.
	PutTargetHealthState(value interface{})
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
	ResetConnectionTermination()
	// Experimental.
	ResetDeregistrationDelay()
	// Experimental.
	ResetHealthCheck()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIpAddressType()
	// Experimental.
	ResetLambdaMultiValueHeadersEnabled()
	// Experimental.
	ResetLoadBalancingAlgorithmType()
	// Experimental.
	ResetLoadBalancingAnomalyMitigation()
	// Experimental.
	ResetLoadBalancingCrossZoneEnabled()
	// Experimental.
	ResetName()
	// Experimental.
	ResetNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPort()
	// Experimental.
	ResetPreserveClientIp()
	// Experimental.
	ResetProtocol()
	// Experimental.
	ResetProtocolVersion()
	// Experimental.
	ResetProxyProtocolV2()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSlowStart()
	// Experimental.
	ResetStickiness()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTargetControlPort()
	// Experimental.
	ResetTargetFailover()
	// Experimental.
	ResetTargetGroupHealth()
	// Experimental.
	ResetTargetHealthState()
	// Experimental.
	ResetTargetType()
	// Experimental.
	ResetVpcId()
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

// The jsii proxy struct for TfTargetGroup
type jsiiProxy_TfTargetGroup struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfTargetGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) ArnSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) ConnectionTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) ConnectionTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) DeregistrationDelay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deregistrationDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) DeregistrationDelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deregistrationDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) HealthCheck() TfTargetGroup_HealthCheckPropertyOutputReference {
	var returns TfTargetGroup_HealthCheckPropertyOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) HealthCheckInput() *TfTargetGroup_HealthCheckProperty {
	var returns *TfTargetGroup_HealthCheckProperty
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) LambdaMultiValueHeadersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaMultiValueHeadersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) LambdaMultiValueHeadersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaMultiValueHeadersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) LoadBalancerArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) LoadBalancingAlgorithmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingAlgorithmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) LoadBalancingAlgorithmTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingAlgorithmTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) LoadBalancingAnomalyMitigation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingAnomalyMitigation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) LoadBalancingAnomalyMitigationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingAnomalyMitigationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) LoadBalancingCrossZoneEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingCrossZoneEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) LoadBalancingCrossZoneEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingCrossZoneEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) PreserveClientIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveClientIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) PreserveClientIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveClientIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) ProtocolVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) ProtocolVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) ProxyProtocolV2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyProtocolV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) ProxyProtocolV2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyProtocolV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) SlowStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) SlowStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slowStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) Stickiness() TfTargetGroup_StickinessPropertyOutputReference {
	var returns TfTargetGroup_StickinessPropertyOutputReference
	_jsii_.Get(
		j,
		"stickiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) StickinessInput() *TfTargetGroup_StickinessProperty {
	var returns *TfTargetGroup_StickinessProperty
	_jsii_.Get(
		j,
		"stickinessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) TargetControlPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetControlPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) TargetControlPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetControlPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) TargetFailover() TfTargetGroup_TargetFailoverPropertyList {
	var returns TfTargetGroup_TargetFailoverPropertyList
	_jsii_.Get(
		j,
		"targetFailover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) TargetFailoverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetFailoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) TargetGroupHealth() TfTargetGroup_TargetGroupHealthPropertyOutputReference {
	var returns TfTargetGroup_TargetGroupHealthPropertyOutputReference
	_jsii_.Get(
		j,
		"targetGroupHealth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) TargetGroupHealthInput() *TfTargetGroup_TargetGroupHealthProperty {
	var returns *TfTargetGroup_TargetGroupHealthProperty
	_jsii_.Get(
		j,
		"targetGroupHealthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) TargetHealthState() TfTargetGroup_TargetHealthStatePropertyList {
	var returns TfTargetGroup_TargetHealthStatePropertyList
	_jsii_.Get(
		j,
		"targetHealthState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) TargetHealthStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetHealthStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) TargetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) TargetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTargetGroup) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group aws_lb_target_group} Resource.
// Experimental.
func NewTfTargetGroup(scope constructs.Construct, id *string, config *TfTargetGroupConfig) TfTargetGroup {
	_init_.Initialize()

	if err := validateNewTfTargetGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTargetGroup{}

	_jsii_.Create(
		"@cdktn/aws-elb.TfTargetGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group aws_lb_target_group} Resource.
// Experimental.
func NewTfTargetGroup_Override(t TfTargetGroup, scope constructs.Construct, id *string, config *TfTargetGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.TfTargetGroup",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetConnectionTermination(val interface{}) {
	if err := j.validateSetConnectionTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionTermination",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetDeregistrationDelay(val *string) {
	if err := j.validateSetDeregistrationDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deregistrationDelay",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetLambdaMultiValueHeadersEnabled(val interface{}) {
	if err := j.validateSetLambdaMultiValueHeadersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaMultiValueHeadersEnabled",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetLoadBalancingAlgorithmType(val *string) {
	if err := j.validateSetLoadBalancingAlgorithmTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingAlgorithmType",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetLoadBalancingAnomalyMitigation(val *string) {
	if err := j.validateSetLoadBalancingAnomalyMitigationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingAnomalyMitigation",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetLoadBalancingCrossZoneEnabled(val *string) {
	if err := j.validateSetLoadBalancingCrossZoneEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingCrossZoneEnabled",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetPreserveClientIp(val *string) {
	if err := j.validateSetPreserveClientIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveClientIp",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetProtocolVersion(val *string) {
	if err := j.validateSetProtocolVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolVersion",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetProxyProtocolV2(val interface{}) {
	if err := j.validateSetProxyProtocolV2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyProtocolV2",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetSlowStart(val *float64) {
	if err := j.validateSetSlowStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slowStart",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetTargetControlPort(val *float64) {
	if err := j.validateSetTargetControlPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetControlPort",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetTargetType(val *string) {
	if err := j.validateSetTargetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

func (j *jsiiProxy_TfTargetGroup)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTN code for importing a TfTargetGroup resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfTargetGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfTargetGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.TfTargetGroup",
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
func TfTargetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTargetGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.TfTargetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfTargetGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTargetGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.TfTargetGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfTargetGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTargetGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.TfTargetGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfTargetGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-elb.TfTargetGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfTargetGroup) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfTargetGroup) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfTargetGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTargetGroup) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTargetGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTargetGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTargetGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTargetGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTargetGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTargetGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTargetGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTargetGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTargetGroup) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfTargetGroup) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTargetGroup) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfTargetGroup) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfTargetGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfTargetGroup) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfTargetGroup) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfTargetGroup) PutHealthCheck(value *TfTargetGroup_HealthCheckProperty) {
	if err := t.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTargetGroup) PutStickiness(value *TfTargetGroup_StickinessProperty) {
	if err := t.validatePutStickinessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStickiness",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTargetGroup) PutTargetFailover(value interface{}) {
	if err := t.validatePutTargetFailoverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTargetFailover",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTargetGroup) PutTargetGroupHealth(value *TfTargetGroup_TargetGroupHealthProperty) {
	if err := t.validatePutTargetGroupHealthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTargetGroupHealth",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTargetGroup) PutTargetHealthState(value interface{}) {
	if err := t.validatePutTargetHealthStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTargetHealthState",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTargetGroup) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetConnectionTermination() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectionTermination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetDeregistrationDelay() {
	_jsii_.InvokeVoid(
		t,
		"resetDeregistrationDelay",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		t,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		t,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetLambdaMultiValueHeadersEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaMultiValueHeadersEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetLoadBalancingAlgorithmType() {
	_jsii_.InvokeVoid(
		t,
		"resetLoadBalancingAlgorithmType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetLoadBalancingAnomalyMitigation() {
	_jsii_.InvokeVoid(
		t,
		"resetLoadBalancingAnomalyMitigation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetLoadBalancingCrossZoneEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetLoadBalancingCrossZoneEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetPort() {
	_jsii_.InvokeVoid(
		t,
		"resetPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetPreserveClientIp() {
	_jsii_.InvokeVoid(
		t,
		"resetPreserveClientIp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetProtocol() {
	_jsii_.InvokeVoid(
		t,
		"resetProtocol",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetProtocolVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetProtocolVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetProxyProtocolV2() {
	_jsii_.InvokeVoid(
		t,
		"resetProxyProtocolV2",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetSlowStart() {
	_jsii_.InvokeVoid(
		t,
		"resetSlowStart",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetStickiness() {
	_jsii_.InvokeVoid(
		t,
		"resetStickiness",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetTargetControlPort() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetControlPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetTargetFailover() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetFailover",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetTargetGroupHealth() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetGroupHealth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetTargetHealthState() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetHealthState",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetTargetType() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) ResetVpcId() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTargetGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTargetGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTargetGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTargetGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTargetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTargetGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTargetGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

