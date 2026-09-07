package vpnsitetosite

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/vpnsitetosite/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/vpnsitetosite/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection aws_vpn_connection}.
// Experimental.
type AwsConnection interface {
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
	CoreNetworkArn() *string
	// Experimental.
	CoreNetworkAttachmentArn() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CustomerGatewayConfiguration() *string
	// Experimental.
	CustomerGatewayId() *string
	// Experimental.
	SetCustomerGatewayId(val *string)
	// Experimental.
	CustomerGatewayIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EnableAcceleration() interface{}
	// Experimental.
	SetEnableAcceleration(val interface{})
	// Experimental.
	EnableAccelerationInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LocalIpv4NetworkCidr() *string
	// Experimental.
	SetLocalIpv4NetworkCidr(val *string)
	// Experimental.
	LocalIpv4NetworkCidrInput() *string
	// Experimental.
	LocalIpv6NetworkCidr() *string
	// Experimental.
	SetLocalIpv6NetworkCidr(val *string)
	// Experimental.
	LocalIpv6NetworkCidrInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OutsideIpAddressType() *string
	// Experimental.
	SetOutsideIpAddressType(val *string)
	// Experimental.
	OutsideIpAddressTypeInput() *string
	// Experimental.
	PresharedKeyArn() *string
	// Experimental.
	PresharedKeyStorage() *string
	// Experimental.
	SetPresharedKeyStorage(val *string)
	// Experimental.
	PresharedKeyStorageInput() *string
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
	RemoteIpv4NetworkCidr() *string
	// Experimental.
	SetRemoteIpv4NetworkCidr(val *string)
	// Experimental.
	RemoteIpv4NetworkCidrInput() *string
	// Experimental.
	RemoteIpv6NetworkCidr() *string
	// Experimental.
	SetRemoteIpv6NetworkCidr(val *string)
	// Experimental.
	RemoteIpv6NetworkCidrInput() *string
	// Experimental.
	Routes() AwsConnection_RoutesPropertyList
	// Experimental.
	StaticRoutesOnly() interface{}
	// Experimental.
	SetStaticRoutesOnly(val interface{})
	// Experimental.
	StaticRoutesOnlyInput() interface{}
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
	TransitGatewayAttachmentId() *string
	// Experimental.
	TransitGatewayId() *string
	// Experimental.
	SetTransitGatewayId(val *string)
	// Experimental.
	TransitGatewayIdInput() *string
	// Experimental.
	TransportTransitGatewayAttachmentId() *string
	// Experimental.
	SetTransportTransitGatewayAttachmentId(val *string)
	// Experimental.
	TransportTransitGatewayAttachmentIdInput() *string
	// Experimental.
	Tunnel1Address() *string
	// Experimental.
	Tunnel1BgpAsn() *string
	// Experimental.
	Tunnel1BgpHoldtime() *float64
	// Experimental.
	Tunnel1CgwInsideAddress() *string
	// Experimental.
	Tunnel1DpdTimeoutAction() *string
	// Experimental.
	SetTunnel1DpdTimeoutAction(val *string)
	// Experimental.
	Tunnel1DpdTimeoutActionInput() *string
	// Experimental.
	Tunnel1DpdTimeoutSeconds() *float64
	// Experimental.
	SetTunnel1DpdTimeoutSeconds(val *float64)
	// Experimental.
	Tunnel1DpdTimeoutSecondsInput() *float64
	// Experimental.
	Tunnel1EnableTunnelLifecycleControl() interface{}
	// Experimental.
	SetTunnel1EnableTunnelLifecycleControl(val interface{})
	// Experimental.
	Tunnel1EnableTunnelLifecycleControlInput() interface{}
	// Experimental.
	Tunnel1IkeVersions() *[]*string
	// Experimental.
	SetTunnel1IkeVersions(val *[]*string)
	// Experimental.
	Tunnel1IkeVersionsInput() *[]*string
	// Experimental.
	Tunnel1InsideCidr() *string
	// Experimental.
	SetTunnel1InsideCidr(val *string)
	// Experimental.
	Tunnel1InsideCidrInput() *string
	// Experimental.
	Tunnel1InsideIpv6Cidr() *string
	// Experimental.
	SetTunnel1InsideIpv6Cidr(val *string)
	// Experimental.
	Tunnel1InsideIpv6CidrInput() *string
	// Experimental.
	Tunnel1LogOptions() AwsConnection_Tunnel1LogOptionsPropertyOutputReference
	// Experimental.
	Tunnel1LogOptionsInput() *AwsConnection_Tunnel1LogOptionsProperty
	// Experimental.
	Tunnel1Phase1DhGroupNumbers() *[]*float64
	// Experimental.
	SetTunnel1Phase1DhGroupNumbers(val *[]*float64)
	// Experimental.
	Tunnel1Phase1DhGroupNumbersInput() *[]*float64
	// Experimental.
	Tunnel1Phase1EncryptionAlgorithms() *[]*string
	// Experimental.
	SetTunnel1Phase1EncryptionAlgorithms(val *[]*string)
	// Experimental.
	Tunnel1Phase1EncryptionAlgorithmsInput() *[]*string
	// Experimental.
	Tunnel1Phase1IntegrityAlgorithms() *[]*string
	// Experimental.
	SetTunnel1Phase1IntegrityAlgorithms(val *[]*string)
	// Experimental.
	Tunnel1Phase1IntegrityAlgorithmsInput() *[]*string
	// Experimental.
	Tunnel1Phase1LifetimeSeconds() *float64
	// Experimental.
	SetTunnel1Phase1LifetimeSeconds(val *float64)
	// Experimental.
	Tunnel1Phase1LifetimeSecondsInput() *float64
	// Experimental.
	Tunnel1Phase2DhGroupNumbers() *[]*float64
	// Experimental.
	SetTunnel1Phase2DhGroupNumbers(val *[]*float64)
	// Experimental.
	Tunnel1Phase2DhGroupNumbersInput() *[]*float64
	// Experimental.
	Tunnel1Phase2EncryptionAlgorithms() *[]*string
	// Experimental.
	SetTunnel1Phase2EncryptionAlgorithms(val *[]*string)
	// Experimental.
	Tunnel1Phase2EncryptionAlgorithmsInput() *[]*string
	// Experimental.
	Tunnel1Phase2IntegrityAlgorithms() *[]*string
	// Experimental.
	SetTunnel1Phase2IntegrityAlgorithms(val *[]*string)
	// Experimental.
	Tunnel1Phase2IntegrityAlgorithmsInput() *[]*string
	// Experimental.
	Tunnel1Phase2LifetimeSeconds() *float64
	// Experimental.
	SetTunnel1Phase2LifetimeSeconds(val *float64)
	// Experimental.
	Tunnel1Phase2LifetimeSecondsInput() *float64
	// Experimental.
	Tunnel1PresharedKey() *string
	// Experimental.
	SetTunnel1PresharedKey(val *string)
	// Experimental.
	Tunnel1PresharedKeyInput() *string
	// Experimental.
	Tunnel1RekeyFuzzPercentage() *float64
	// Experimental.
	SetTunnel1RekeyFuzzPercentage(val *float64)
	// Experimental.
	Tunnel1RekeyFuzzPercentageInput() *float64
	// Experimental.
	Tunnel1RekeyMarginTimeSeconds() *float64
	// Experimental.
	SetTunnel1RekeyMarginTimeSeconds(val *float64)
	// Experimental.
	Tunnel1RekeyMarginTimeSecondsInput() *float64
	// Experimental.
	Tunnel1ReplayWindowSize() *float64
	// Experimental.
	SetTunnel1ReplayWindowSize(val *float64)
	// Experimental.
	Tunnel1ReplayWindowSizeInput() *float64
	// Experimental.
	Tunnel1StartupAction() *string
	// Experimental.
	SetTunnel1StartupAction(val *string)
	// Experimental.
	Tunnel1StartupActionInput() *string
	// Experimental.
	Tunnel1VgwInsideAddress() *string
	// Experimental.
	Tunnel2Address() *string
	// Experimental.
	Tunnel2BgpAsn() *string
	// Experimental.
	Tunnel2BgpHoldtime() *float64
	// Experimental.
	Tunnel2CgwInsideAddress() *string
	// Experimental.
	Tunnel2DpdTimeoutAction() *string
	// Experimental.
	SetTunnel2DpdTimeoutAction(val *string)
	// Experimental.
	Tunnel2DpdTimeoutActionInput() *string
	// Experimental.
	Tunnel2DpdTimeoutSeconds() *float64
	// Experimental.
	SetTunnel2DpdTimeoutSeconds(val *float64)
	// Experimental.
	Tunnel2DpdTimeoutSecondsInput() *float64
	// Experimental.
	Tunnel2EnableTunnelLifecycleControl() interface{}
	// Experimental.
	SetTunnel2EnableTunnelLifecycleControl(val interface{})
	// Experimental.
	Tunnel2EnableTunnelLifecycleControlInput() interface{}
	// Experimental.
	Tunnel2IkeVersions() *[]*string
	// Experimental.
	SetTunnel2IkeVersions(val *[]*string)
	// Experimental.
	Tunnel2IkeVersionsInput() *[]*string
	// Experimental.
	Tunnel2InsideCidr() *string
	// Experimental.
	SetTunnel2InsideCidr(val *string)
	// Experimental.
	Tunnel2InsideCidrInput() *string
	// Experimental.
	Tunnel2InsideIpv6Cidr() *string
	// Experimental.
	SetTunnel2InsideIpv6Cidr(val *string)
	// Experimental.
	Tunnel2InsideIpv6CidrInput() *string
	// Experimental.
	Tunnel2LogOptions() AwsConnection_Tunnel2LogOptionsPropertyOutputReference
	// Experimental.
	Tunnel2LogOptionsInput() *AwsConnection_Tunnel2LogOptionsProperty
	// Experimental.
	Tunnel2Phase1DhGroupNumbers() *[]*float64
	// Experimental.
	SetTunnel2Phase1DhGroupNumbers(val *[]*float64)
	// Experimental.
	Tunnel2Phase1DhGroupNumbersInput() *[]*float64
	// Experimental.
	Tunnel2Phase1EncryptionAlgorithms() *[]*string
	// Experimental.
	SetTunnel2Phase1EncryptionAlgorithms(val *[]*string)
	// Experimental.
	Tunnel2Phase1EncryptionAlgorithmsInput() *[]*string
	// Experimental.
	Tunnel2Phase1IntegrityAlgorithms() *[]*string
	// Experimental.
	SetTunnel2Phase1IntegrityAlgorithms(val *[]*string)
	// Experimental.
	Tunnel2Phase1IntegrityAlgorithmsInput() *[]*string
	// Experimental.
	Tunnel2Phase1LifetimeSeconds() *float64
	// Experimental.
	SetTunnel2Phase1LifetimeSeconds(val *float64)
	// Experimental.
	Tunnel2Phase1LifetimeSecondsInput() *float64
	// Experimental.
	Tunnel2Phase2DhGroupNumbers() *[]*float64
	// Experimental.
	SetTunnel2Phase2DhGroupNumbers(val *[]*float64)
	// Experimental.
	Tunnel2Phase2DhGroupNumbersInput() *[]*float64
	// Experimental.
	Tunnel2Phase2EncryptionAlgorithms() *[]*string
	// Experimental.
	SetTunnel2Phase2EncryptionAlgorithms(val *[]*string)
	// Experimental.
	Tunnel2Phase2EncryptionAlgorithmsInput() *[]*string
	// Experimental.
	Tunnel2Phase2IntegrityAlgorithms() *[]*string
	// Experimental.
	SetTunnel2Phase2IntegrityAlgorithms(val *[]*string)
	// Experimental.
	Tunnel2Phase2IntegrityAlgorithmsInput() *[]*string
	// Experimental.
	Tunnel2Phase2LifetimeSeconds() *float64
	// Experimental.
	SetTunnel2Phase2LifetimeSeconds(val *float64)
	// Experimental.
	Tunnel2Phase2LifetimeSecondsInput() *float64
	// Experimental.
	Tunnel2PresharedKey() *string
	// Experimental.
	SetTunnel2PresharedKey(val *string)
	// Experimental.
	Tunnel2PresharedKeyInput() *string
	// Experimental.
	Tunnel2RekeyFuzzPercentage() *float64
	// Experimental.
	SetTunnel2RekeyFuzzPercentage(val *float64)
	// Experimental.
	Tunnel2RekeyFuzzPercentageInput() *float64
	// Experimental.
	Tunnel2RekeyMarginTimeSeconds() *float64
	// Experimental.
	SetTunnel2RekeyMarginTimeSeconds(val *float64)
	// Experimental.
	Tunnel2RekeyMarginTimeSecondsInput() *float64
	// Experimental.
	Tunnel2ReplayWindowSize() *float64
	// Experimental.
	SetTunnel2ReplayWindowSize(val *float64)
	// Experimental.
	Tunnel2ReplayWindowSizeInput() *float64
	// Experimental.
	Tunnel2StartupAction() *string
	// Experimental.
	SetTunnel2StartupAction(val *string)
	// Experimental.
	Tunnel2StartupActionInput() *string
	// Experimental.
	Tunnel2VgwInsideAddress() *string
	// Experimental.
	TunnelBandwidth() *string
	// Experimental.
	SetTunnelBandwidth(val *string)
	// Experimental.
	TunnelBandwidthInput() *string
	// Experimental.
	TunnelInsideIpVersion() *string
	// Experimental.
	SetTunnelInsideIpVersion(val *string)
	// Experimental.
	TunnelInsideIpVersionInput() *string
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
	// Experimental.
	VgwTelemetry() AwsConnection_VgwTelemetryPropertyList
	// Experimental.
	VpnConcentratorId() *string
	// Experimental.
	SetVpnConcentratorId(val *string)
	// Experimental.
	VpnConcentratorIdInput() *string
	// Experimental.
	VpnGatewayId() *string
	// Experimental.
	SetVpnGatewayId(val *string)
	// Experimental.
	VpnGatewayIdInput() *string
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
	PutTunnel1LogOptions(value *AwsConnection_Tunnel1LogOptionsProperty)
	// Experimental.
	PutTunnel2LogOptions(value *AwsConnection_Tunnel2LogOptionsProperty)
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
	ResetEnableAcceleration()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLocalIpv4NetworkCidr()
	// Experimental.
	ResetLocalIpv6NetworkCidr()
	// Experimental.
	ResetOutsideIpAddressType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPresharedKeyStorage()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRemoteIpv4NetworkCidr()
	// Experimental.
	ResetRemoteIpv6NetworkCidr()
	// Experimental.
	ResetStaticRoutesOnly()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTransitGatewayId()
	// Experimental.
	ResetTransportTransitGatewayAttachmentId()
	// Experimental.
	ResetTunnel1DpdTimeoutAction()
	// Experimental.
	ResetTunnel1DpdTimeoutSeconds()
	// Experimental.
	ResetTunnel1EnableTunnelLifecycleControl()
	// Experimental.
	ResetTunnel1IkeVersions()
	// Experimental.
	ResetTunnel1InsideCidr()
	// Experimental.
	ResetTunnel1InsideIpv6Cidr()
	// Experimental.
	ResetTunnel1LogOptions()
	// Experimental.
	ResetTunnel1Phase1DhGroupNumbers()
	// Experimental.
	ResetTunnel1Phase1EncryptionAlgorithms()
	// Experimental.
	ResetTunnel1Phase1IntegrityAlgorithms()
	// Experimental.
	ResetTunnel1Phase1LifetimeSeconds()
	// Experimental.
	ResetTunnel1Phase2DhGroupNumbers()
	// Experimental.
	ResetTunnel1Phase2EncryptionAlgorithms()
	// Experimental.
	ResetTunnel1Phase2IntegrityAlgorithms()
	// Experimental.
	ResetTunnel1Phase2LifetimeSeconds()
	// Experimental.
	ResetTunnel1PresharedKey()
	// Experimental.
	ResetTunnel1RekeyFuzzPercentage()
	// Experimental.
	ResetTunnel1RekeyMarginTimeSeconds()
	// Experimental.
	ResetTunnel1ReplayWindowSize()
	// Experimental.
	ResetTunnel1StartupAction()
	// Experimental.
	ResetTunnel2DpdTimeoutAction()
	// Experimental.
	ResetTunnel2DpdTimeoutSeconds()
	// Experimental.
	ResetTunnel2EnableTunnelLifecycleControl()
	// Experimental.
	ResetTunnel2IkeVersions()
	// Experimental.
	ResetTunnel2InsideCidr()
	// Experimental.
	ResetTunnel2InsideIpv6Cidr()
	// Experimental.
	ResetTunnel2LogOptions()
	// Experimental.
	ResetTunnel2Phase1DhGroupNumbers()
	// Experimental.
	ResetTunnel2Phase1EncryptionAlgorithms()
	// Experimental.
	ResetTunnel2Phase1IntegrityAlgorithms()
	// Experimental.
	ResetTunnel2Phase1LifetimeSeconds()
	// Experimental.
	ResetTunnel2Phase2DhGroupNumbers()
	// Experimental.
	ResetTunnel2Phase2EncryptionAlgorithms()
	// Experimental.
	ResetTunnel2Phase2IntegrityAlgorithms()
	// Experimental.
	ResetTunnel2Phase2LifetimeSeconds()
	// Experimental.
	ResetTunnel2PresharedKey()
	// Experimental.
	ResetTunnel2RekeyFuzzPercentage()
	// Experimental.
	ResetTunnel2RekeyMarginTimeSeconds()
	// Experimental.
	ResetTunnel2ReplayWindowSize()
	// Experimental.
	ResetTunnel2StartupAction()
	// Experimental.
	ResetTunnelBandwidth()
	// Experimental.
	ResetTunnelInsideIpVersion()
	// Experimental.
	ResetVpnConcentratorId()
	// Experimental.
	ResetVpnGatewayId()
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

// The jsii proxy struct for AwsConnection
type jsiiProxy_AwsConnection struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsConnection) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) CoreNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) CoreNetworkAttachmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkAttachmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) CustomerGatewayConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) CustomerGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) CustomerGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) EnableAcceleration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceleration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) EnableAccelerationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAccelerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) LocalIpv4NetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpv4NetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) LocalIpv4NetworkCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpv4NetworkCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) LocalIpv6NetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpv6NetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) LocalIpv6NetworkCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpv6NetworkCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) OutsideIpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outsideIpAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) OutsideIpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outsideIpAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) PresharedKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"presharedKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) PresharedKeyStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"presharedKeyStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) PresharedKeyStorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"presharedKeyStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) RemoteIpv4NetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpv4NetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) RemoteIpv4NetworkCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpv4NetworkCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) RemoteIpv6NetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpv6NetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) RemoteIpv6NetworkCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpv6NetworkCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Routes() AwsConnection_RoutesPropertyList {
	var returns AwsConnection_RoutesPropertyList
	_jsii_.Get(
		j,
		"routes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) StaticRoutesOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticRoutesOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) StaticRoutesOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticRoutesOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) TransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) TransitGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) TransportTransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportTransitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) TransportTransitGatewayAttachmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportTransitGatewayAttachmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1BgpAsn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1BgpAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1BgpHoldtime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1BgpHoldtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1CgwInsideAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1CgwInsideAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1DpdTimeoutAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1DpdTimeoutAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1DpdTimeoutActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1DpdTimeoutActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1DpdTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1DpdTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1DpdTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1DpdTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1EnableTunnelLifecycleControl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tunnel1EnableTunnelLifecycleControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1EnableTunnelLifecycleControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tunnel1EnableTunnelLifecycleControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1IkeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1IkeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1IkeVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1IkeVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1InsideCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1InsideCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1InsideCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1InsideCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1InsideIpv6Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1InsideIpv6Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1InsideIpv6CidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1InsideIpv6CidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1LogOptions() AwsConnection_Tunnel1LogOptionsPropertyOutputReference {
	var returns AwsConnection_Tunnel1LogOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"tunnel1LogOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1LogOptionsInput() *AwsConnection_Tunnel1LogOptionsProperty {
	var returns *AwsConnection_Tunnel1LogOptionsProperty
	_jsii_.Get(
		j,
		"tunnel1LogOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Phase1DhGroupNumbers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel1Phase1DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Phase1DhGroupNumbersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel1Phase1DhGroupNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Phase1EncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase1EncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Phase1EncryptionAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase1EncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Phase1IntegrityAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase1IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Phase1IntegrityAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase1IntegrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Phase1LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1Phase1LifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Phase1LifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1Phase1LifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Phase2DhGroupNumbers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel1Phase2DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Phase2DhGroupNumbersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel1Phase2DhGroupNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Phase2EncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase2EncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Phase2EncryptionAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase2EncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Phase2IntegrityAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase2IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Phase2IntegrityAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase2IntegrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Phase2LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1Phase2LifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1Phase2LifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1Phase2LifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1PresharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1PresharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1PresharedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1PresharedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1RekeyFuzzPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1RekeyFuzzPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1RekeyFuzzPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1RekeyFuzzPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1RekeyMarginTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1RekeyMarginTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1RekeyMarginTimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1RekeyMarginTimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1ReplayWindowSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1ReplayWindowSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1ReplayWindowSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1ReplayWindowSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1StartupAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1StartupAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1StartupActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1StartupActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel1VgwInsideAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1VgwInsideAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2BgpAsn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2BgpAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2BgpHoldtime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2BgpHoldtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2CgwInsideAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2CgwInsideAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2DpdTimeoutAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2DpdTimeoutAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2DpdTimeoutActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2DpdTimeoutActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2DpdTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2DpdTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2DpdTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2DpdTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2EnableTunnelLifecycleControl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tunnel2EnableTunnelLifecycleControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2EnableTunnelLifecycleControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tunnel2EnableTunnelLifecycleControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2IkeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2IkeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2IkeVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2IkeVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2InsideCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2InsideCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2InsideCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2InsideCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2InsideIpv6Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2InsideIpv6Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2InsideIpv6CidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2InsideIpv6CidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2LogOptions() AwsConnection_Tunnel2LogOptionsPropertyOutputReference {
	var returns AwsConnection_Tunnel2LogOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"tunnel2LogOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2LogOptionsInput() *AwsConnection_Tunnel2LogOptionsProperty {
	var returns *AwsConnection_Tunnel2LogOptionsProperty
	_jsii_.Get(
		j,
		"tunnel2LogOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Phase1DhGroupNumbers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel2Phase1DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Phase1DhGroupNumbersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel2Phase1DhGroupNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Phase1EncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase1EncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Phase1EncryptionAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase1EncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Phase1IntegrityAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase1IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Phase1IntegrityAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase1IntegrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Phase1LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2Phase1LifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Phase1LifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2Phase1LifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Phase2DhGroupNumbers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel2Phase2DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Phase2DhGroupNumbersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel2Phase2DhGroupNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Phase2EncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase2EncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Phase2EncryptionAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase2EncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Phase2IntegrityAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase2IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Phase2IntegrityAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase2IntegrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Phase2LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2Phase2LifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2Phase2LifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2Phase2LifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2PresharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2PresharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2PresharedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2PresharedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2RekeyFuzzPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2RekeyFuzzPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2RekeyFuzzPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2RekeyFuzzPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2RekeyMarginTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2RekeyMarginTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2RekeyMarginTimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2RekeyMarginTimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2ReplayWindowSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2ReplayWindowSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2ReplayWindowSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2ReplayWindowSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2StartupAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2StartupAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2StartupActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2StartupActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Tunnel2VgwInsideAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2VgwInsideAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) TunnelBandwidth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) TunnelBandwidthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelBandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) TunnelInsideIpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelInsideIpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) TunnelInsideIpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelInsideIpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) VgwTelemetry() AwsConnection_VgwTelemetryPropertyList {
	var returns AwsConnection_VgwTelemetryPropertyList
	_jsii_.Get(
		j,
		"vgwTelemetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) VpnConcentratorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnConcentratorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) VpnConcentratorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnConcentratorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) VpnGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection) VpnGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection aws_vpn_connection} Resource.
// Experimental.
func NewAwsConnection(scope constructs.Construct, id *string, config *AwsConnectionConfig) AwsConnection {
	_init_.Initialize()

	if err := validateNewAwsConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConnection{}

	_jsii_.Create(
		"@cdktn/aws-vpn-site-to-site.AwsConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection aws_vpn_connection} Resource.
// Experimental.
func NewAwsConnection_Override(a AwsConnection, scope constructs.Construct, id *string, config *AwsConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpn-site-to-site.AwsConnection",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetCustomerGatewayId(val *string) {
	if err := j.validateSetCustomerGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerGatewayId",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetEnableAcceleration(val interface{}) {
	if err := j.validateSetEnableAccelerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAcceleration",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetLocalIpv4NetworkCidr(val *string) {
	if err := j.validateSetLocalIpv4NetworkCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localIpv4NetworkCidr",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetLocalIpv6NetworkCidr(val *string) {
	if err := j.validateSetLocalIpv6NetworkCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localIpv6NetworkCidr",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetOutsideIpAddressType(val *string) {
	if err := j.validateSetOutsideIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outsideIpAddressType",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetPresharedKeyStorage(val *string) {
	if err := j.validateSetPresharedKeyStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"presharedKeyStorage",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetRemoteIpv4NetworkCidr(val *string) {
	if err := j.validateSetRemoteIpv4NetworkCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteIpv4NetworkCidr",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetRemoteIpv6NetworkCidr(val *string) {
	if err := j.validateSetRemoteIpv6NetworkCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteIpv6NetworkCidr",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetStaticRoutesOnly(val interface{}) {
	if err := j.validateSetStaticRoutesOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticRoutesOnly",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTransitGatewayId(val *string) {
	if err := j.validateSetTransitGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayId",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTransportTransitGatewayAttachmentId(val *string) {
	if err := j.validateSetTransportTransitGatewayAttachmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportTransitGatewayAttachmentId",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1DpdTimeoutAction(val *string) {
	if err := j.validateSetTunnel1DpdTimeoutActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1DpdTimeoutAction",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1DpdTimeoutSeconds(val *float64) {
	if err := j.validateSetTunnel1DpdTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1DpdTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1EnableTunnelLifecycleControl(val interface{}) {
	if err := j.validateSetTunnel1EnableTunnelLifecycleControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1EnableTunnelLifecycleControl",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1IkeVersions(val *[]*string) {
	if err := j.validateSetTunnel1IkeVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1IkeVersions",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1InsideCidr(val *string) {
	if err := j.validateSetTunnel1InsideCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1InsideCidr",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1InsideIpv6Cidr(val *string) {
	if err := j.validateSetTunnel1InsideIpv6CidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1InsideIpv6Cidr",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1Phase1DhGroupNumbers(val *[]*float64) {
	if err := j.validateSetTunnel1Phase1DhGroupNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase1DhGroupNumbers",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1Phase1EncryptionAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel1Phase1EncryptionAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase1EncryptionAlgorithms",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1Phase1IntegrityAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel1Phase1IntegrityAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase1IntegrityAlgorithms",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1Phase1LifetimeSeconds(val *float64) {
	if err := j.validateSetTunnel1Phase1LifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase1LifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1Phase2DhGroupNumbers(val *[]*float64) {
	if err := j.validateSetTunnel1Phase2DhGroupNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase2DhGroupNumbers",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1Phase2EncryptionAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel1Phase2EncryptionAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase2EncryptionAlgorithms",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1Phase2IntegrityAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel1Phase2IntegrityAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase2IntegrityAlgorithms",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1Phase2LifetimeSeconds(val *float64) {
	if err := j.validateSetTunnel1Phase2LifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase2LifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1PresharedKey(val *string) {
	if err := j.validateSetTunnel1PresharedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1PresharedKey",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1RekeyFuzzPercentage(val *float64) {
	if err := j.validateSetTunnel1RekeyFuzzPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1RekeyFuzzPercentage",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1RekeyMarginTimeSeconds(val *float64) {
	if err := j.validateSetTunnel1RekeyMarginTimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1RekeyMarginTimeSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1ReplayWindowSize(val *float64) {
	if err := j.validateSetTunnel1ReplayWindowSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1ReplayWindowSize",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel1StartupAction(val *string) {
	if err := j.validateSetTunnel1StartupActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1StartupAction",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2DpdTimeoutAction(val *string) {
	if err := j.validateSetTunnel2DpdTimeoutActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2DpdTimeoutAction",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2DpdTimeoutSeconds(val *float64) {
	if err := j.validateSetTunnel2DpdTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2DpdTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2EnableTunnelLifecycleControl(val interface{}) {
	if err := j.validateSetTunnel2EnableTunnelLifecycleControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2EnableTunnelLifecycleControl",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2IkeVersions(val *[]*string) {
	if err := j.validateSetTunnel2IkeVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2IkeVersions",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2InsideCidr(val *string) {
	if err := j.validateSetTunnel2InsideCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2InsideCidr",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2InsideIpv6Cidr(val *string) {
	if err := j.validateSetTunnel2InsideIpv6CidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2InsideIpv6Cidr",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2Phase1DhGroupNumbers(val *[]*float64) {
	if err := j.validateSetTunnel2Phase1DhGroupNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase1DhGroupNumbers",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2Phase1EncryptionAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel2Phase1EncryptionAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase1EncryptionAlgorithms",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2Phase1IntegrityAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel2Phase1IntegrityAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase1IntegrityAlgorithms",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2Phase1LifetimeSeconds(val *float64) {
	if err := j.validateSetTunnel2Phase1LifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase1LifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2Phase2DhGroupNumbers(val *[]*float64) {
	if err := j.validateSetTunnel2Phase2DhGroupNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase2DhGroupNumbers",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2Phase2EncryptionAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel2Phase2EncryptionAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase2EncryptionAlgorithms",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2Phase2IntegrityAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel2Phase2IntegrityAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase2IntegrityAlgorithms",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2Phase2LifetimeSeconds(val *float64) {
	if err := j.validateSetTunnel2Phase2LifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase2LifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2PresharedKey(val *string) {
	if err := j.validateSetTunnel2PresharedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2PresharedKey",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2RekeyFuzzPercentage(val *float64) {
	if err := j.validateSetTunnel2RekeyFuzzPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2RekeyFuzzPercentage",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2RekeyMarginTimeSeconds(val *float64) {
	if err := j.validateSetTunnel2RekeyMarginTimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2RekeyMarginTimeSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2ReplayWindowSize(val *float64) {
	if err := j.validateSetTunnel2ReplayWindowSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2ReplayWindowSize",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnel2StartupAction(val *string) {
	if err := j.validateSetTunnel2StartupActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2StartupAction",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnelBandwidth(val *string) {
	if err := j.validateSetTunnelBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelBandwidth",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetTunnelInsideIpVersion(val *string) {
	if err := j.validateSetTunnelInsideIpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelInsideIpVersion",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetVpnConcentratorId(val *string) {
	if err := j.validateSetVpnConcentratorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnConcentratorId",
		val,
	)
}

func (j *jsiiProxy_AwsConnection)SetVpnGatewayId(val *string) {
	if err := j.validateSetVpnGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnGatewayId",
		val,
	)
}

// Generates CDKTN code for importing a AwsConnection resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsConnection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsConnection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-site-to-site.AwsConnection",
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
func AwsConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-site-to-site.AwsConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-site-to-site.AwsConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-site-to-site.AwsConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-vpn-site-to-site.AwsConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsConnection) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsConnection) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConnection) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConnection) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConnection) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConnection) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnection) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsConnection) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnection) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsConnection) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsConnection) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsConnection) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsConnection) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsConnection) PutTunnel1LogOptions(value *AwsConnection_Tunnel1LogOptionsProperty) {
	if err := a.validatePutTunnel1LogOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTunnel1LogOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnection) PutTunnel2LogOptions(value *AwsConnection_Tunnel2LogOptionsProperty) {
	if err := a.validatePutTunnel2LogOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTunnel2LogOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnection) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsConnection) ResetEnableAcceleration() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableAcceleration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetLocalIpv4NetworkCidr() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalIpv4NetworkCidr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetLocalIpv6NetworkCidr() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalIpv6NetworkCidr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetOutsideIpAddressType() {
	_jsii_.InvokeVoid(
		a,
		"resetOutsideIpAddressType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetPresharedKeyStorage() {
	_jsii_.InvokeVoid(
		a,
		"resetPresharedKeyStorage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetRemoteIpv4NetworkCidr() {
	_jsii_.InvokeVoid(
		a,
		"resetRemoteIpv4NetworkCidr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetRemoteIpv6NetworkCidr() {
	_jsii_.InvokeVoid(
		a,
		"resetRemoteIpv6NetworkCidr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetStaticRoutesOnly() {
	_jsii_.InvokeVoid(
		a,
		"resetStaticRoutesOnly",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTransitGatewayId() {
	_jsii_.InvokeVoid(
		a,
		"resetTransitGatewayId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTransportTransitGatewayAttachmentId() {
	_jsii_.InvokeVoid(
		a,
		"resetTransportTransitGatewayAttachmentId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1DpdTimeoutAction() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1DpdTimeoutAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1DpdTimeoutSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1DpdTimeoutSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1EnableTunnelLifecycleControl() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1EnableTunnelLifecycleControl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1IkeVersions() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1IkeVersions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1InsideCidr() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1InsideCidr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1InsideIpv6Cidr() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1InsideIpv6Cidr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1LogOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1LogOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1Phase1DhGroupNumbers() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1Phase1DhGroupNumbers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1Phase1EncryptionAlgorithms() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1Phase1EncryptionAlgorithms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1Phase1IntegrityAlgorithms() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1Phase1IntegrityAlgorithms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1Phase1LifetimeSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1Phase1LifetimeSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1Phase2DhGroupNumbers() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1Phase2DhGroupNumbers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1Phase2EncryptionAlgorithms() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1Phase2EncryptionAlgorithms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1Phase2IntegrityAlgorithms() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1Phase2IntegrityAlgorithms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1Phase2LifetimeSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1Phase2LifetimeSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1PresharedKey() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1PresharedKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1RekeyFuzzPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1RekeyFuzzPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1RekeyMarginTimeSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1RekeyMarginTimeSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1ReplayWindowSize() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1ReplayWindowSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel1StartupAction() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel1StartupAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2DpdTimeoutAction() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2DpdTimeoutAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2DpdTimeoutSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2DpdTimeoutSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2EnableTunnelLifecycleControl() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2EnableTunnelLifecycleControl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2IkeVersions() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2IkeVersions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2InsideCidr() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2InsideCidr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2InsideIpv6Cidr() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2InsideIpv6Cidr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2LogOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2LogOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2Phase1DhGroupNumbers() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2Phase1DhGroupNumbers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2Phase1EncryptionAlgorithms() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2Phase1EncryptionAlgorithms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2Phase1IntegrityAlgorithms() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2Phase1IntegrityAlgorithms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2Phase1LifetimeSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2Phase1LifetimeSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2Phase2DhGroupNumbers() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2Phase2DhGroupNumbers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2Phase2EncryptionAlgorithms() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2Phase2EncryptionAlgorithms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2Phase2IntegrityAlgorithms() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2Phase2IntegrityAlgorithms",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2Phase2LifetimeSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2Phase2LifetimeSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2PresharedKey() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2PresharedKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2RekeyFuzzPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2RekeyFuzzPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2RekeyMarginTimeSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2RekeyMarginTimeSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2ReplayWindowSize() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2ReplayWindowSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnel2StartupAction() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnel2StartupAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnelBandwidth() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnelBandwidth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetTunnelInsideIpVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetTunnelInsideIpVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetVpnConcentratorId() {
	_jsii_.InvokeVoid(
		a,
		"resetVpnConcentratorId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) ResetVpnGatewayId() {
	_jsii_.InvokeVoid(
		a,
		"resetVpnGatewayId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnection) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnection) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnection) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

