package awselb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselb/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awselb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb aws_alb}.
// Experimental.
type AwsAlb interface {
	cdktn.TerraformResource
	// Experimental.
	AccessLogs() AwsAlb_AccessLogsPropertyOutputReference
	// Experimental.
	AccessLogsInput() *AwsAlb_AccessLogsProperty
	// Experimental.
	Arn() *string
	// Experimental.
	ArnSuffix() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClientKeepAlive() *float64
	// Experimental.
	SetClientKeepAlive(val *float64)
	// Experimental.
	ClientKeepAliveInput() *float64
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConnectionLogs() AwsAlb_ConnectionLogsPropertyOutputReference
	// Experimental.
	ConnectionLogsInput() *AwsAlb_ConnectionLogsProperty
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CustomerOwnedIpv4Pool() *string
	// Experimental.
	SetCustomerOwnedIpv4Pool(val *string)
	// Experimental.
	CustomerOwnedIpv4PoolInput() *string
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
	DnsRecordClientRoutingPolicy() *string
	// Experimental.
	SetDnsRecordClientRoutingPolicy(val *string)
	// Experimental.
	DnsRecordClientRoutingPolicyInput() *string
	// Experimental.
	DropInvalidHeaderFields() interface{}
	// Experimental.
	SetDropInvalidHeaderFields(val interface{})
	// Experimental.
	DropInvalidHeaderFieldsInput() interface{}
	// Experimental.
	EnableCrossZoneLoadBalancing() interface{}
	// Experimental.
	SetEnableCrossZoneLoadBalancing(val interface{})
	// Experimental.
	EnableCrossZoneLoadBalancingInput() interface{}
	// Experimental.
	EnableDeletionProtection() interface{}
	// Experimental.
	SetEnableDeletionProtection(val interface{})
	// Experimental.
	EnableDeletionProtectionInput() interface{}
	// Experimental.
	EnableHttp2() interface{}
	// Experimental.
	SetEnableHttp2(val interface{})
	// Experimental.
	EnableHttp2Input() interface{}
	// Experimental.
	EnablePrefixForIpv6SourceNat() *string
	// Experimental.
	SetEnablePrefixForIpv6SourceNat(val *string)
	// Experimental.
	EnablePrefixForIpv6SourceNatInput() *string
	// Experimental.
	EnableTlsVersionAndCipherSuiteHeaders() interface{}
	// Experimental.
	SetEnableTlsVersionAndCipherSuiteHeaders(val interface{})
	// Experimental.
	EnableTlsVersionAndCipherSuiteHeadersInput() interface{}
	// Experimental.
	EnableWafFailOpen() interface{}
	// Experimental.
	SetEnableWafFailOpen(val interface{})
	// Experimental.
	EnableWafFailOpenInput() interface{}
	// Experimental.
	EnableXffClientPort() interface{}
	// Experimental.
	SetEnableXffClientPort(val interface{})
	// Experimental.
	EnableXffClientPortInput() interface{}
	// Experimental.
	EnableZonalShift() interface{}
	// Experimental.
	SetEnableZonalShift(val interface{})
	// Experimental.
	EnableZonalShiftInput() interface{}
	// Experimental.
	EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic() *string
	// Experimental.
	SetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic(val *string)
	// Experimental.
	EnforceSecurityGroupInboundRulesOnPrivateLinkTrafficInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HealthCheckLogs() AwsAlb_HealthCheckLogsPropertyOutputReference
	// Experimental.
	HealthCheckLogsInput() *AwsAlb_HealthCheckLogsProperty
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
	Internal() interface{}
	// Experimental.
	SetInternal(val interface{})
	// Experimental.
	InternalInput() interface{}
	// Experimental.
	IpAddressType() *string
	// Experimental.
	SetIpAddressType(val *string)
	// Experimental.
	IpAddressTypeInput() *string
	// Experimental.
	IpamPools() AwsAlb_IpamPoolsPropertyOutputReference
	// Experimental.
	IpamPoolsInput() *AwsAlb_IpamPoolsProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LoadBalancerType() *string
	// Experimental.
	SetLoadBalancerType(val *string)
	// Experimental.
	LoadBalancerTypeInput() *string
	// Experimental.
	MinimumLoadBalancerCapacity() AwsAlb_MinimumLoadBalancerCapacityPropertyOutputReference
	// Experimental.
	MinimumLoadBalancerCapacityInput() *AwsAlb_MinimumLoadBalancerCapacityProperty
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
	PreserveHostHeader() interface{}
	// Experimental.
	SetPreserveHostHeader(val interface{})
	// Experimental.
	PreserveHostHeaderInput() interface{}
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
	SecondaryIpsAutoAssignedPerSubnet() *float64
	// Experimental.
	SetSecondaryIpsAutoAssignedPerSubnet(val *float64)
	// Experimental.
	SecondaryIpsAutoAssignedPerSubnetInput() *float64
	// Experimental.
	SecurityGroups() *[]*string
	// Experimental.
	SetSecurityGroups(val *[]*string)
	// Experimental.
	SecurityGroupsInput() *[]*string
	// Experimental.
	SubnetMapping() AwsAlb_SubnetMappingPropertyList
	// Experimental.
	SubnetMappingInput() interface{}
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
	Timeouts() AwsAlb_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	VpcId() *string
	// Experimental.
	XffHeaderProcessingMode() *string
	// Experimental.
	SetXffHeaderProcessingMode(val *string)
	// Experimental.
	XffHeaderProcessingModeInput() *string
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
	PutAccessLogs(value *AwsAlb_AccessLogsProperty)
	// Experimental.
	PutConnectionLogs(value *AwsAlb_ConnectionLogsProperty)
	// Experimental.
	PutHealthCheckLogs(value *AwsAlb_HealthCheckLogsProperty)
	// Experimental.
	PutIpamPools(value *AwsAlb_IpamPoolsProperty)
	// Experimental.
	PutMinimumLoadBalancerCapacity(value *AwsAlb_MinimumLoadBalancerCapacityProperty)
	// Experimental.
	PutSubnetMapping(value interface{})
	// Experimental.
	PutTimeouts(value *AwsAlb_TimeoutsProperty)
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
	ResetClientKeepAlive()
	// Experimental.
	ResetConnectionLogs()
	// Experimental.
	ResetCustomerOwnedIpv4Pool()
	// Experimental.
	ResetDesyncMitigationMode()
	// Experimental.
	ResetDnsRecordClientRoutingPolicy()
	// Experimental.
	ResetDropInvalidHeaderFields()
	// Experimental.
	ResetEnableCrossZoneLoadBalancing()
	// Experimental.
	ResetEnableDeletionProtection()
	// Experimental.
	ResetEnableHttp2()
	// Experimental.
	ResetEnablePrefixForIpv6SourceNat()
	// Experimental.
	ResetEnableTlsVersionAndCipherSuiteHeaders()
	// Experimental.
	ResetEnableWafFailOpen()
	// Experimental.
	ResetEnableXffClientPort()
	// Experimental.
	ResetEnableZonalShift()
	// Experimental.
	ResetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic()
	// Experimental.
	ResetHealthCheckLogs()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIdleTimeout()
	// Experimental.
	ResetInternal()
	// Experimental.
	ResetIpAddressType()
	// Experimental.
	ResetIpamPools()
	// Experimental.
	ResetLoadBalancerType()
	// Experimental.
	ResetMinimumLoadBalancerCapacity()
	// Experimental.
	ResetName()
	// Experimental.
	ResetNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPreserveHostHeader()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSecondaryIpsAutoAssignedPerSubnet()
	// Experimental.
	ResetSecurityGroups()
	// Experimental.
	ResetSubnetMapping()
	// Experimental.
	ResetSubnets()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetXffHeaderProcessingMode()
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

// The jsii proxy struct for AwsAlb
type jsiiProxy_AwsAlb struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsAlb) AccessLogs() AwsAlb_AccessLogsPropertyOutputReference {
	var returns AwsAlb_AccessLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"accessLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) AccessLogsInput() *AwsAlb_AccessLogsProperty {
	var returns *AwsAlb_AccessLogsProperty
	_jsii_.Get(
		j,
		"accessLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) ArnSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) ClientKeepAlive() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) ClientKeepAliveInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) ConnectionLogs() AwsAlb_ConnectionLogsPropertyOutputReference {
	var returns AwsAlb_ConnectionLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"connectionLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) ConnectionLogsInput() *AwsAlb_ConnectionLogsProperty {
	var returns *AwsAlb_ConnectionLogsProperty
	_jsii_.Get(
		j,
		"connectionLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) CustomerOwnedIpv4Pool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerOwnedIpv4Pool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) CustomerOwnedIpv4PoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerOwnedIpv4PoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) DesyncMitigationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desyncMitigationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) DesyncMitigationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desyncMitigationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) DnsRecordClientRoutingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsRecordClientRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) DnsRecordClientRoutingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsRecordClientRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) DropInvalidHeaderFields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropInvalidHeaderFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) DropInvalidHeaderFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropInvalidHeaderFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnableCrossZoneLoadBalancing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCrossZoneLoadBalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnableCrossZoneLoadBalancingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCrossZoneLoadBalancingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnableDeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDeletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnableDeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDeletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnableHttp2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttp2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnableHttp2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttp2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnablePrefixForIpv6SourceNat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enablePrefixForIpv6SourceNat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnablePrefixForIpv6SourceNatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enablePrefixForIpv6SourceNatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnableTlsVersionAndCipherSuiteHeaders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTlsVersionAndCipherSuiteHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnableTlsVersionAndCipherSuiteHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTlsVersionAndCipherSuiteHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnableWafFailOpen() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWafFailOpen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnableWafFailOpenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWafFailOpenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnableXffClientPort() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableXffClientPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnableXffClientPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableXffClientPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnableZonalShift() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableZonalShift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnableZonalShiftInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableZonalShiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) EnforceSecurityGroupInboundRulesOnPrivateLinkTrafficInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceSecurityGroupInboundRulesOnPrivateLinkTrafficInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) HealthCheckLogs() AwsAlb_HealthCheckLogsPropertyOutputReference {
	var returns AwsAlb_HealthCheckLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"healthCheckLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) HealthCheckLogsInput() *AwsAlb_HealthCheckLogsProperty {
	var returns *AwsAlb_HealthCheckLogsProperty
	_jsii_.Get(
		j,
		"healthCheckLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) IdleTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) IdleTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) Internal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) InternalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) IpamPools() AwsAlb_IpamPoolsPropertyOutputReference {
	var returns AwsAlb_IpamPoolsPropertyOutputReference
	_jsii_.Get(
		j,
		"ipamPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) IpamPoolsInput() *AwsAlb_IpamPoolsProperty {
	var returns *AwsAlb_IpamPoolsProperty
	_jsii_.Get(
		j,
		"ipamPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) LoadBalancerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) LoadBalancerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) MinimumLoadBalancerCapacity() AwsAlb_MinimumLoadBalancerCapacityPropertyOutputReference {
	var returns AwsAlb_MinimumLoadBalancerCapacityPropertyOutputReference
	_jsii_.Get(
		j,
		"minimumLoadBalancerCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) MinimumLoadBalancerCapacityInput() *AwsAlb_MinimumLoadBalancerCapacityProperty {
	var returns *AwsAlb_MinimumLoadBalancerCapacityProperty
	_jsii_.Get(
		j,
		"minimumLoadBalancerCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) PreserveHostHeader() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveHostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) PreserveHostHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveHostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) SecondaryIpsAutoAssignedPerSubnet() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondaryIpsAutoAssignedPerSubnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) SecondaryIpsAutoAssignedPerSubnetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondaryIpsAutoAssignedPerSubnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) SubnetMapping() AwsAlb_SubnetMappingPropertyList {
	var returns AwsAlb_SubnetMappingPropertyList
	_jsii_.Get(
		j,
		"subnetMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) SubnetMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) SubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) Timeouts() AwsAlb_TimeoutsPropertyOutputReference {
	var returns AwsAlb_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) XffHeaderProcessingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xffHeaderProcessingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) XffHeaderProcessingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xffHeaderProcessingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlb) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb aws_alb} Resource.
// Experimental.
func NewAwsAlb(scope constructs.Construct, id *string, config *AwsAlbConfig) AwsAlb {
	_init_.Initialize()

	if err := validateNewAwsAlbParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAlb{}

	_jsii_.Create(
		"@cdktn/aws-elb.AwsAlb",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb aws_alb} Resource.
// Experimental.
func NewAwsAlb_Override(a AwsAlb, scope constructs.Construct, id *string, config *AwsAlbConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.AwsAlb",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsAlb)SetClientKeepAlive(val *float64) {
	if err := j.validateSetClientKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientKeepAlive",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetCustomerOwnedIpv4Pool(val *string) {
	if err := j.validateSetCustomerOwnedIpv4PoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerOwnedIpv4Pool",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetDesyncMitigationMode(val *string) {
	if err := j.validateSetDesyncMitigationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desyncMitigationMode",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetDnsRecordClientRoutingPolicy(val *string) {
	if err := j.validateSetDnsRecordClientRoutingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsRecordClientRoutingPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetDropInvalidHeaderFields(val interface{}) {
	if err := j.validateSetDropInvalidHeaderFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dropInvalidHeaderFields",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetEnableCrossZoneLoadBalancing(val interface{}) {
	if err := j.validateSetEnableCrossZoneLoadBalancingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCrossZoneLoadBalancing",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetEnableDeletionProtection(val interface{}) {
	if err := j.validateSetEnableDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDeletionProtection",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetEnableHttp2(val interface{}) {
	if err := j.validateSetEnableHttp2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHttp2",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetEnablePrefixForIpv6SourceNat(val *string) {
	if err := j.validateSetEnablePrefixForIpv6SourceNatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrefixForIpv6SourceNat",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetEnableTlsVersionAndCipherSuiteHeaders(val interface{}) {
	if err := j.validateSetEnableTlsVersionAndCipherSuiteHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTlsVersionAndCipherSuiteHeaders",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetEnableWafFailOpen(val interface{}) {
	if err := j.validateSetEnableWafFailOpenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableWafFailOpen",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetEnableXffClientPort(val interface{}) {
	if err := j.validateSetEnableXffClientPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableXffClientPort",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetEnableZonalShift(val interface{}) {
	if err := j.validateSetEnableZonalShiftParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableZonalShift",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic(val *string) {
	if err := j.validateSetEnforceSecurityGroupInboundRulesOnPrivateLinkTrafficParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetIdleTimeout(val *float64) {
	if err := j.validateSetIdleTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetInternal(val interface{}) {
	if err := j.validateSetInternalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internal",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetLoadBalancerType(val *string) {
	if err := j.validateSetLoadBalancerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerType",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetPreserveHostHeader(val interface{}) {
	if err := j.validateSetPreserveHostHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveHostHeader",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetSecondaryIpsAutoAssignedPerSubnet(val *float64) {
	if err := j.validateSetSecondaryIpsAutoAssignedPerSubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryIpsAutoAssignedPerSubnet",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetSubnets(val *[]*string) {
	if err := j.validateSetSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsAlb)SetXffHeaderProcessingMode(val *string) {
	if err := j.validateSetXffHeaderProcessingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xffHeaderProcessingMode",
		val,
	)
}

// Generates CDKTN code for importing a AwsAlb resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsAlb_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsAlb_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.AwsAlb",
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
func AwsAlb_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAlb_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.AwsAlb",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAlb_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAlb_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.AwsAlb",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAlb_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAlb_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.AwsAlb",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsAlb_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-elb.AwsAlb",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsAlb) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsAlb) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsAlb) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAlb) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAlb) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAlb) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAlb) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAlb) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAlb) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAlb) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAlb) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAlb) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlb) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsAlb) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAlb) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsAlb) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAlb) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsAlb) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAlb) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsAlb) PutAccessLogs(value *AwsAlb_AccessLogsProperty) {
	if err := a.validatePutAccessLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlb) PutConnectionLogs(value *AwsAlb_ConnectionLogsProperty) {
	if err := a.validatePutConnectionLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectionLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlb) PutHealthCheckLogs(value *AwsAlb_HealthCheckLogsProperty) {
	if err := a.validatePutHealthCheckLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHealthCheckLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlb) PutIpamPools(value *AwsAlb_IpamPoolsProperty) {
	if err := a.validatePutIpamPoolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIpamPools",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlb) PutMinimumLoadBalancerCapacity(value *AwsAlb_MinimumLoadBalancerCapacityProperty) {
	if err := a.validatePutMinimumLoadBalancerCapacityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMinimumLoadBalancerCapacity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlb) PutSubnetMapping(value interface{}) {
	if err := a.validatePutSubnetMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSubnetMapping",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlb) PutTimeouts(value *AwsAlb_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlb) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsAlb) ResetAccessLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetClientKeepAlive() {
	_jsii_.InvokeVoid(
		a,
		"resetClientKeepAlive",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetConnectionLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetCustomerOwnedIpv4Pool() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerOwnedIpv4Pool",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetDesyncMitigationMode() {
	_jsii_.InvokeVoid(
		a,
		"resetDesyncMitigationMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetDnsRecordClientRoutingPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetDnsRecordClientRoutingPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetDropInvalidHeaderFields() {
	_jsii_.InvokeVoid(
		a,
		"resetDropInvalidHeaderFields",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetEnableCrossZoneLoadBalancing() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableCrossZoneLoadBalancing",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetEnableDeletionProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableDeletionProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetEnableHttp2() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableHttp2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetEnablePrefixForIpv6SourceNat() {
	_jsii_.InvokeVoid(
		a,
		"resetEnablePrefixForIpv6SourceNat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetEnableTlsVersionAndCipherSuiteHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableTlsVersionAndCipherSuiteHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetEnableWafFailOpen() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableWafFailOpen",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetEnableXffClientPort() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableXffClientPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetEnableZonalShift() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableZonalShift",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic() {
	_jsii_.InvokeVoid(
		a,
		"resetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetHealthCheckLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheckLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetIdleTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetIdleTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetInternal() {
	_jsii_.InvokeVoid(
		a,
		"resetInternal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		a,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetIpamPools() {
	_jsii_.InvokeVoid(
		a,
		"resetIpamPools",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetLoadBalancerType() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancerType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetMinimumLoadBalancerCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetMinimumLoadBalancerCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetPreserveHostHeader() {
	_jsii_.InvokeVoid(
		a,
		"resetPreserveHostHeader",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetSecondaryIpsAutoAssignedPerSubnet() {
	_jsii_.InvokeVoid(
		a,
		"resetSecondaryIpsAutoAssignedPerSubnet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetSubnetMapping() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetMapping",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetSubnets() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) ResetXffHeaderProcessingMode() {
	_jsii_.InvokeVoid(
		a,
		"resetXffHeaderProcessingMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlb) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlb) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlb) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlb) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlb) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlb) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlb) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

