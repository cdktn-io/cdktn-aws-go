package elb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elb/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/elb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group aws_lb_target_group}.
// Experimental.
type AwsTargetGroup interface {
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
	HealthCheck() AwsTargetGroup_HealthCheckPropertyOutputReference
	// Experimental.
	HealthCheckInput() *AwsTargetGroup_HealthCheckProperty
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
	Stickiness() AwsTargetGroup_StickinessPropertyOutputReference
	// Experimental.
	StickinessInput() *AwsTargetGroup_StickinessProperty
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
	TargetFailover() AwsTargetGroup_TargetFailoverPropertyList
	// Experimental.
	TargetFailoverInput() interface{}
	// Experimental.
	TargetGroupHealth() AwsTargetGroup_TargetGroupHealthPropertyOutputReference
	// Experimental.
	TargetGroupHealthInput() *AwsTargetGroup_TargetGroupHealthProperty
	// Experimental.
	TargetHealthState() AwsTargetGroup_TargetHealthStatePropertyList
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
	PutHealthCheck(value *AwsTargetGroup_HealthCheckProperty)
	// Experimental.
	PutStickiness(value *AwsTargetGroup_StickinessProperty)
	// Experimental.
	PutTargetFailover(value interface{})
	// Experimental.
	PutTargetGroupHealth(value *AwsTargetGroup_TargetGroupHealthProperty)
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

// The jsii proxy struct for AwsTargetGroup
type jsiiProxy_AwsTargetGroup struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsTargetGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) ArnSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) ConnectionTermination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionTermination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) ConnectionTerminationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionTerminationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) DeregistrationDelay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deregistrationDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) DeregistrationDelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deregistrationDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) HealthCheck() AwsTargetGroup_HealthCheckPropertyOutputReference {
	var returns AwsTargetGroup_HealthCheckPropertyOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) HealthCheckInput() *AwsTargetGroup_HealthCheckProperty {
	var returns *AwsTargetGroup_HealthCheckProperty
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) LambdaMultiValueHeadersEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaMultiValueHeadersEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) LambdaMultiValueHeadersEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaMultiValueHeadersEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) LoadBalancerArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) LoadBalancingAlgorithmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingAlgorithmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) LoadBalancingAlgorithmTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingAlgorithmTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) LoadBalancingAnomalyMitigation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingAnomalyMitigation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) LoadBalancingAnomalyMitigationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingAnomalyMitigationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) LoadBalancingCrossZoneEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingCrossZoneEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) LoadBalancingCrossZoneEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingCrossZoneEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) PreserveClientIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveClientIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) PreserveClientIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preserveClientIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) ProtocolVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) ProtocolVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) ProxyProtocolV2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyProtocolV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) ProxyProtocolV2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"proxyProtocolV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) SlowStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) SlowStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"slowStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) Stickiness() AwsTargetGroup_StickinessPropertyOutputReference {
	var returns AwsTargetGroup_StickinessPropertyOutputReference
	_jsii_.Get(
		j,
		"stickiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) StickinessInput() *AwsTargetGroup_StickinessProperty {
	var returns *AwsTargetGroup_StickinessProperty
	_jsii_.Get(
		j,
		"stickinessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) TargetControlPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetControlPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) TargetControlPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetControlPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) TargetFailover() AwsTargetGroup_TargetFailoverPropertyList {
	var returns AwsTargetGroup_TargetFailoverPropertyList
	_jsii_.Get(
		j,
		"targetFailover",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) TargetFailoverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetFailoverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) TargetGroupHealth() AwsTargetGroup_TargetGroupHealthPropertyOutputReference {
	var returns AwsTargetGroup_TargetGroupHealthPropertyOutputReference
	_jsii_.Get(
		j,
		"targetGroupHealth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) TargetGroupHealthInput() *AwsTargetGroup_TargetGroupHealthProperty {
	var returns *AwsTargetGroup_TargetGroupHealthProperty
	_jsii_.Get(
		j,
		"targetGroupHealthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) TargetHealthState() AwsTargetGroup_TargetHealthStatePropertyList {
	var returns AwsTargetGroup_TargetHealthStatePropertyList
	_jsii_.Get(
		j,
		"targetHealthState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) TargetHealthStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetHealthStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) TargetType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) TargetTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTargetGroup) VpcIdInput() *string {
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
func NewAwsTargetGroup(scope constructs.Construct, id *string, config *AwsTargetGroupConfig) AwsTargetGroup {
	_init_.Initialize()

	if err := validateNewAwsTargetGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTargetGroup{}

	_jsii_.Create(
		"@cdktn/aws-elb.AwsTargetGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group aws_lb_target_group} Resource.
// Experimental.
func NewAwsTargetGroup_Override(a AwsTargetGroup, scope constructs.Construct, id *string, config *AwsTargetGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.AwsTargetGroup",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetConnectionTermination(val interface{}) {
	if err := j.validateSetConnectionTerminationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionTermination",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetDeregistrationDelay(val *string) {
	if err := j.validateSetDeregistrationDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deregistrationDelay",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetLambdaMultiValueHeadersEnabled(val interface{}) {
	if err := j.validateSetLambdaMultiValueHeadersEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaMultiValueHeadersEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetLoadBalancingAlgorithmType(val *string) {
	if err := j.validateSetLoadBalancingAlgorithmTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingAlgorithmType",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetLoadBalancingAnomalyMitigation(val *string) {
	if err := j.validateSetLoadBalancingAnomalyMitigationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingAnomalyMitigation",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetLoadBalancingCrossZoneEnabled(val *string) {
	if err := j.validateSetLoadBalancingCrossZoneEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingCrossZoneEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetPreserveClientIp(val *string) {
	if err := j.validateSetPreserveClientIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveClientIp",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetProtocolVersion(val *string) {
	if err := j.validateSetProtocolVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolVersion",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetProxyProtocolV2(val interface{}) {
	if err := j.validateSetProxyProtocolV2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyProtocolV2",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetSlowStart(val *float64) {
	if err := j.validateSetSlowStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slowStart",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetTargetControlPort(val *float64) {
	if err := j.validateSetTargetControlPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetControlPort",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetTargetType(val *string) {
	if err := j.validateSetTargetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetType",
		val,
	)
}

func (j *jsiiProxy_AwsTargetGroup)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

// Generates CDKTN code for importing a AwsTargetGroup resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsTargetGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsTargetGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.AwsTargetGroup",
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
func AwsTargetGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsTargetGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.AwsTargetGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsTargetGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsTargetGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.AwsTargetGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsTargetGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsTargetGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.AwsTargetGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsTargetGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-elb.AwsTargetGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsTargetGroup) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsTargetGroup) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsTargetGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTargetGroup) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTargetGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTargetGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTargetGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTargetGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTargetGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTargetGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTargetGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTargetGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTargetGroup) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsTargetGroup) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTargetGroup) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsTargetGroup) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsTargetGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsTargetGroup) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsTargetGroup) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsTargetGroup) PutHealthCheck(value *AwsTargetGroup_HealthCheckProperty) {
	if err := a.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTargetGroup) PutStickiness(value *AwsTargetGroup_StickinessProperty) {
	if err := a.validatePutStickinessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStickiness",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTargetGroup) PutTargetFailover(value interface{}) {
	if err := a.validatePutTargetFailoverParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTargetFailover",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTargetGroup) PutTargetGroupHealth(value *AwsTargetGroup_TargetGroupHealthProperty) {
	if err := a.validatePutTargetGroupHealthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTargetGroupHealth",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTargetGroup) PutTargetHealthState(value interface{}) {
	if err := a.validatePutTargetHealthStateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTargetHealthState",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTargetGroup) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetConnectionTermination() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionTermination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetDeregistrationDelay() {
	_jsii_.InvokeVoid(
		a,
		"resetDeregistrationDelay",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		a,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetLambdaMultiValueHeadersEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaMultiValueHeadersEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetLoadBalancingAlgorithmType() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancingAlgorithmType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetLoadBalancingAnomalyMitigation() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancingAnomalyMitigation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetLoadBalancingCrossZoneEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancingCrossZoneEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetPreserveClientIp() {
	_jsii_.InvokeVoid(
		a,
		"resetPreserveClientIp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetProtocolVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocolVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetProxyProtocolV2() {
	_jsii_.InvokeVoid(
		a,
		"resetProxyProtocolV2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetSlowStart() {
	_jsii_.InvokeVoid(
		a,
		"resetSlowStart",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetStickiness() {
	_jsii_.InvokeVoid(
		a,
		"resetStickiness",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetTargetControlPort() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetControlPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetTargetFailover() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetFailover",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetTargetGroupHealth() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetGroupHealth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetTargetHealthState() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetHealthState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetTargetType() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) ResetVpcId() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTargetGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTargetGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTargetGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTargetGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTargetGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTargetGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTargetGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

