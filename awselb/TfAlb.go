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
type TfAlb interface {
	cdktn.TerraformResource
	// Experimental.
	AccessLogs() TfAlb_AccessLogsPropertyOutputReference
	// Experimental.
	AccessLogsInput() *TfAlb_AccessLogsProperty
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
	ConnectionLogs() TfAlb_ConnectionLogsPropertyOutputReference
	// Experimental.
	ConnectionLogsInput() *TfAlb_ConnectionLogsProperty
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
	HealthCheckLogs() TfAlb_HealthCheckLogsPropertyOutputReference
	// Experimental.
	HealthCheckLogsInput() *TfAlb_HealthCheckLogsProperty
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
	IpamPools() TfAlb_IpamPoolsPropertyOutputReference
	// Experimental.
	IpamPoolsInput() *TfAlb_IpamPoolsProperty
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
	MinimumLoadBalancerCapacity() TfAlb_MinimumLoadBalancerCapacityPropertyOutputReference
	// Experimental.
	MinimumLoadBalancerCapacityInput() *TfAlb_MinimumLoadBalancerCapacityProperty
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
	SubnetMapping() TfAlb_SubnetMappingPropertyList
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
	Timeouts() TfAlb_TimeoutsPropertyOutputReference
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
	PutAccessLogs(value *TfAlb_AccessLogsProperty)
	// Experimental.
	PutConnectionLogs(value *TfAlb_ConnectionLogsProperty)
	// Experimental.
	PutHealthCheckLogs(value *TfAlb_HealthCheckLogsProperty)
	// Experimental.
	PutIpamPools(value *TfAlb_IpamPoolsProperty)
	// Experimental.
	PutMinimumLoadBalancerCapacity(value *TfAlb_MinimumLoadBalancerCapacityProperty)
	// Experimental.
	PutSubnetMapping(value interface{})
	// Experimental.
	PutTimeouts(value *TfAlb_TimeoutsProperty)
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

// The jsii proxy struct for TfAlb
type jsiiProxy_TfAlb struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfAlb) AccessLogs() TfAlb_AccessLogsPropertyOutputReference {
	var returns TfAlb_AccessLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"accessLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) AccessLogsInput() *TfAlb_AccessLogsProperty {
	var returns *TfAlb_AccessLogsProperty
	_jsii_.Get(
		j,
		"accessLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) ArnSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) ClientKeepAlive() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientKeepAlive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) ClientKeepAliveInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientKeepAliveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) ConnectionLogs() TfAlb_ConnectionLogsPropertyOutputReference {
	var returns TfAlb_ConnectionLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"connectionLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) ConnectionLogsInput() *TfAlb_ConnectionLogsProperty {
	var returns *TfAlb_ConnectionLogsProperty
	_jsii_.Get(
		j,
		"connectionLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) CustomerOwnedIpv4Pool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerOwnedIpv4Pool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) CustomerOwnedIpv4PoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerOwnedIpv4PoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) DesyncMitigationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desyncMitigationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) DesyncMitigationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desyncMitigationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) DnsRecordClientRoutingPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsRecordClientRoutingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) DnsRecordClientRoutingPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsRecordClientRoutingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) DropInvalidHeaderFields() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropInvalidHeaderFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) DropInvalidHeaderFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dropInvalidHeaderFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnableCrossZoneLoadBalancing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCrossZoneLoadBalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnableCrossZoneLoadBalancingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCrossZoneLoadBalancingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnableDeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDeletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnableDeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDeletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnableHttp2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttp2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnableHttp2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttp2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnablePrefixForIpv6SourceNat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enablePrefixForIpv6SourceNat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnablePrefixForIpv6SourceNatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enablePrefixForIpv6SourceNatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnableTlsVersionAndCipherSuiteHeaders() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTlsVersionAndCipherSuiteHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnableTlsVersionAndCipherSuiteHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTlsVersionAndCipherSuiteHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnableWafFailOpen() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWafFailOpen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnableWafFailOpenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableWafFailOpenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnableXffClientPort() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableXffClientPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnableXffClientPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableXffClientPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnableZonalShift() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableZonalShift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnableZonalShiftInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableZonalShiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) EnforceSecurityGroupInboundRulesOnPrivateLinkTrafficInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceSecurityGroupInboundRulesOnPrivateLinkTrafficInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) HealthCheckLogs() TfAlb_HealthCheckLogsPropertyOutputReference {
	var returns TfAlb_HealthCheckLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"healthCheckLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) HealthCheckLogsInput() *TfAlb_HealthCheckLogsProperty {
	var returns *TfAlb_HealthCheckLogsProperty
	_jsii_.Get(
		j,
		"healthCheckLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) IdleTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) IdleTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) Internal() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) InternalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) IpamPools() TfAlb_IpamPoolsPropertyOutputReference {
	var returns TfAlb_IpamPoolsPropertyOutputReference
	_jsii_.Get(
		j,
		"ipamPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) IpamPoolsInput() *TfAlb_IpamPoolsProperty {
	var returns *TfAlb_IpamPoolsProperty
	_jsii_.Get(
		j,
		"ipamPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) LoadBalancerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) LoadBalancerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) MinimumLoadBalancerCapacity() TfAlb_MinimumLoadBalancerCapacityPropertyOutputReference {
	var returns TfAlb_MinimumLoadBalancerCapacityPropertyOutputReference
	_jsii_.Get(
		j,
		"minimumLoadBalancerCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) MinimumLoadBalancerCapacityInput() *TfAlb_MinimumLoadBalancerCapacityProperty {
	var returns *TfAlb_MinimumLoadBalancerCapacityProperty
	_jsii_.Get(
		j,
		"minimumLoadBalancerCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) PreserveHostHeader() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveHostHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) PreserveHostHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveHostHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) SecondaryIpsAutoAssignedPerSubnet() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondaryIpsAutoAssignedPerSubnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) SecondaryIpsAutoAssignedPerSubnetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondaryIpsAutoAssignedPerSubnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) SubnetMapping() TfAlb_SubnetMappingPropertyList {
	var returns TfAlb_SubnetMappingPropertyList
	_jsii_.Get(
		j,
		"subnetMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) SubnetMappingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) SubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) Timeouts() TfAlb_TimeoutsPropertyOutputReference {
	var returns TfAlb_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) XffHeaderProcessingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xffHeaderProcessingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) XffHeaderProcessingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"xffHeaderProcessingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlb) ZoneId() *string {
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
func NewTfAlb(scope constructs.Construct, id *string, config *TfAlbConfig) TfAlb {
	_init_.Initialize()

	if err := validateNewTfAlbParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAlb{}

	_jsii_.Create(
		"@cdktn/aws-elb.TfAlb",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb aws_alb} Resource.
// Experimental.
func NewTfAlb_Override(t TfAlb, scope constructs.Construct, id *string, config *TfAlbConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.TfAlb",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfAlb)SetClientKeepAlive(val *float64) {
	if err := j.validateSetClientKeepAliveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientKeepAlive",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetCustomerOwnedIpv4Pool(val *string) {
	if err := j.validateSetCustomerOwnedIpv4PoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerOwnedIpv4Pool",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetDesyncMitigationMode(val *string) {
	if err := j.validateSetDesyncMitigationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desyncMitigationMode",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetDnsRecordClientRoutingPolicy(val *string) {
	if err := j.validateSetDnsRecordClientRoutingPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsRecordClientRoutingPolicy",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetDropInvalidHeaderFields(val interface{}) {
	if err := j.validateSetDropInvalidHeaderFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dropInvalidHeaderFields",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetEnableCrossZoneLoadBalancing(val interface{}) {
	if err := j.validateSetEnableCrossZoneLoadBalancingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCrossZoneLoadBalancing",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetEnableDeletionProtection(val interface{}) {
	if err := j.validateSetEnableDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDeletionProtection",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetEnableHttp2(val interface{}) {
	if err := j.validateSetEnableHttp2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHttp2",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetEnablePrefixForIpv6SourceNat(val *string) {
	if err := j.validateSetEnablePrefixForIpv6SourceNatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrefixForIpv6SourceNat",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetEnableTlsVersionAndCipherSuiteHeaders(val interface{}) {
	if err := j.validateSetEnableTlsVersionAndCipherSuiteHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTlsVersionAndCipherSuiteHeaders",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetEnableWafFailOpen(val interface{}) {
	if err := j.validateSetEnableWafFailOpenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableWafFailOpen",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetEnableXffClientPort(val interface{}) {
	if err := j.validateSetEnableXffClientPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableXffClientPort",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetEnableZonalShift(val interface{}) {
	if err := j.validateSetEnableZonalShiftParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableZonalShift",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic(val *string) {
	if err := j.validateSetEnforceSecurityGroupInboundRulesOnPrivateLinkTrafficParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetIdleTimeout(val *float64) {
	if err := j.validateSetIdleTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTimeout",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetInternal(val interface{}) {
	if err := j.validateSetInternalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internal",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetLoadBalancerType(val *string) {
	if err := j.validateSetLoadBalancerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerType",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetPreserveHostHeader(val interface{}) {
	if err := j.validateSetPreserveHostHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveHostHeader",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetSecondaryIpsAutoAssignedPerSubnet(val *float64) {
	if err := j.validateSetSecondaryIpsAutoAssignedPerSubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryIpsAutoAssignedPerSubnet",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetSubnets(val *[]*string) {
	if err := j.validateSetSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfAlb)SetXffHeaderProcessingMode(val *string) {
	if err := j.validateSetXffHeaderProcessingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xffHeaderProcessingMode",
		val,
	)
}

// Generates CDKTN code for importing a TfAlb resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfAlb_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfAlb_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.TfAlb",
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
func TfAlb_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfAlb_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.TfAlb",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfAlb_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfAlb_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.TfAlb",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfAlb_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfAlb_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.TfAlb",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfAlb_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-elb.TfAlb",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfAlb) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfAlb) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfAlb) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAlb) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAlb) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAlb) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAlb) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAlb) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAlb) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAlb) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAlb) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAlb) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlb) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfAlb) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAlb) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfAlb) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfAlb) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfAlb) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfAlb) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfAlb) PutAccessLogs(value *TfAlb_AccessLogsProperty) {
	if err := t.validatePutAccessLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAccessLogs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlb) PutConnectionLogs(value *TfAlb_ConnectionLogsProperty) {
	if err := t.validatePutConnectionLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConnectionLogs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlb) PutHealthCheckLogs(value *TfAlb_HealthCheckLogsProperty) {
	if err := t.validatePutHealthCheckLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHealthCheckLogs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlb) PutIpamPools(value *TfAlb_IpamPoolsProperty) {
	if err := t.validatePutIpamPoolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIpamPools",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlb) PutMinimumLoadBalancerCapacity(value *TfAlb_MinimumLoadBalancerCapacityProperty) {
	if err := t.validatePutMinimumLoadBalancerCapacityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMinimumLoadBalancerCapacity",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlb) PutSubnetMapping(value interface{}) {
	if err := t.validatePutSubnetMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSubnetMapping",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlb) PutTimeouts(value *TfAlb_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlb) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfAlb) ResetAccessLogs() {
	_jsii_.InvokeVoid(
		t,
		"resetAccessLogs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetClientKeepAlive() {
	_jsii_.InvokeVoid(
		t,
		"resetClientKeepAlive",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetConnectionLogs() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectionLogs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetCustomerOwnedIpv4Pool() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomerOwnedIpv4Pool",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetDesyncMitigationMode() {
	_jsii_.InvokeVoid(
		t,
		"resetDesyncMitigationMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetDnsRecordClientRoutingPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetDnsRecordClientRoutingPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetDropInvalidHeaderFields() {
	_jsii_.InvokeVoid(
		t,
		"resetDropInvalidHeaderFields",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetEnableCrossZoneLoadBalancing() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableCrossZoneLoadBalancing",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetEnableDeletionProtection() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableDeletionProtection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetEnableHttp2() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableHttp2",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetEnablePrefixForIpv6SourceNat() {
	_jsii_.InvokeVoid(
		t,
		"resetEnablePrefixForIpv6SourceNat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetEnableTlsVersionAndCipherSuiteHeaders() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableTlsVersionAndCipherSuiteHeaders",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetEnableWafFailOpen() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableWafFailOpen",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetEnableXffClientPort() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableXffClientPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetEnableZonalShift() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableZonalShift",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic() {
	_jsii_.InvokeVoid(
		t,
		"resetEnforceSecurityGroupInboundRulesOnPrivateLinkTraffic",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetHealthCheckLogs() {
	_jsii_.InvokeVoid(
		t,
		"resetHealthCheckLogs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetIdleTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetIdleTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetInternal() {
	_jsii_.InvokeVoid(
		t,
		"resetInternal",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		t,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetIpamPools() {
	_jsii_.InvokeVoid(
		t,
		"resetIpamPools",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetLoadBalancerType() {
	_jsii_.InvokeVoid(
		t,
		"resetLoadBalancerType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetMinimumLoadBalancerCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetMinimumLoadBalancerCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetPreserveHostHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetPreserveHostHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetSecondaryIpsAutoAssignedPerSubnet() {
	_jsii_.InvokeVoid(
		t,
		"resetSecondaryIpsAutoAssignedPerSubnet",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetSubnetMapping() {
	_jsii_.InvokeVoid(
		t,
		"resetSubnetMapping",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetSubnets() {
	_jsii_.InvokeVoid(
		t,
		"resetSubnets",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) ResetXffHeaderProcessingMode() {
	_jsii_.InvokeVoid(
		t,
		"resetXffHeaderProcessingMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlb) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlb) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlb) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlb) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlb) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlb) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlb) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

