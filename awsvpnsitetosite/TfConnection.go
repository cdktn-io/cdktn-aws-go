package awsvpnsitetosite

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpnsitetosite/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsvpnsitetosite/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection aws_vpn_connection}.
// Experimental.
type TfConnection interface {
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
	Routes() TfConnection_RoutesPropertyList
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
	Tunnel1LogOptions() TfConnection_Tunnel1LogOptionsPropertyOutputReference
	// Experimental.
	Tunnel1LogOptionsInput() *TfConnection_Tunnel1LogOptionsProperty
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
	Tunnel2LogOptions() TfConnection_Tunnel2LogOptionsPropertyOutputReference
	// Experimental.
	Tunnel2LogOptionsInput() *TfConnection_Tunnel2LogOptionsProperty
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
	VgwTelemetry() TfConnection_VgwTelemetryPropertyList
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
	PutTunnel1LogOptions(value *TfConnection_Tunnel1LogOptionsProperty)
	// Experimental.
	PutTunnel2LogOptions(value *TfConnection_Tunnel2LogOptionsProperty)
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

// The jsii proxy struct for TfConnection
type jsiiProxy_TfConnection struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfConnection) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) CoreNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) CoreNetworkAttachmentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coreNetworkAttachmentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) CustomerGatewayConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) CustomerGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) CustomerGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) EnableAcceleration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAcceleration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) EnableAccelerationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAccelerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) LocalIpv4NetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpv4NetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) LocalIpv4NetworkCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpv4NetworkCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) LocalIpv6NetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpv6NetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) LocalIpv6NetworkCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localIpv6NetworkCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) OutsideIpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outsideIpAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) OutsideIpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outsideIpAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) PresharedKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"presharedKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) PresharedKeyStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"presharedKeyStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) PresharedKeyStorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"presharedKeyStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) RemoteIpv4NetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpv4NetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) RemoteIpv4NetworkCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpv4NetworkCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) RemoteIpv6NetworkCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpv6NetworkCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) RemoteIpv6NetworkCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteIpv6NetworkCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Routes() TfConnection_RoutesPropertyList {
	var returns TfConnection_RoutesPropertyList
	_jsii_.Get(
		j,
		"routes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) StaticRoutesOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticRoutesOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) StaticRoutesOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticRoutesOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) TransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) TransitGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) TransitGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) TransportTransitGatewayAttachmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportTransitGatewayAttachmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) TransportTransitGatewayAttachmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportTransitGatewayAttachmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1BgpAsn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1BgpAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1BgpHoldtime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1BgpHoldtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1CgwInsideAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1CgwInsideAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1DpdTimeoutAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1DpdTimeoutAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1DpdTimeoutActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1DpdTimeoutActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1DpdTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1DpdTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1DpdTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1DpdTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1EnableTunnelLifecycleControl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tunnel1EnableTunnelLifecycleControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1EnableTunnelLifecycleControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tunnel1EnableTunnelLifecycleControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1IkeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1IkeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1IkeVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1IkeVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1InsideCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1InsideCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1InsideCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1InsideCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1InsideIpv6Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1InsideIpv6Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1InsideIpv6CidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1InsideIpv6CidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1LogOptions() TfConnection_Tunnel1LogOptionsPropertyOutputReference {
	var returns TfConnection_Tunnel1LogOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"tunnel1LogOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1LogOptionsInput() *TfConnection_Tunnel1LogOptionsProperty {
	var returns *TfConnection_Tunnel1LogOptionsProperty
	_jsii_.Get(
		j,
		"tunnel1LogOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Phase1DhGroupNumbers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel1Phase1DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Phase1DhGroupNumbersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel1Phase1DhGroupNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Phase1EncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase1EncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Phase1EncryptionAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase1EncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Phase1IntegrityAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase1IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Phase1IntegrityAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase1IntegrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Phase1LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1Phase1LifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Phase1LifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1Phase1LifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Phase2DhGroupNumbers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel1Phase2DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Phase2DhGroupNumbersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel1Phase2DhGroupNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Phase2EncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase2EncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Phase2EncryptionAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase2EncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Phase2IntegrityAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase2IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Phase2IntegrityAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel1Phase2IntegrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Phase2LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1Phase2LifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1Phase2LifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1Phase2LifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1PresharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1PresharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1PresharedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1PresharedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1RekeyFuzzPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1RekeyFuzzPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1RekeyFuzzPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1RekeyFuzzPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1RekeyMarginTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1RekeyMarginTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1RekeyMarginTimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1RekeyMarginTimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1ReplayWindowSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1ReplayWindowSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1ReplayWindowSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel1ReplayWindowSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1StartupAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1StartupAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1StartupActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1StartupActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel1VgwInsideAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel1VgwInsideAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2BgpAsn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2BgpAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2BgpHoldtime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2BgpHoldtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2CgwInsideAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2CgwInsideAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2DpdTimeoutAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2DpdTimeoutAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2DpdTimeoutActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2DpdTimeoutActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2DpdTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2DpdTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2DpdTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2DpdTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2EnableTunnelLifecycleControl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tunnel2EnableTunnelLifecycleControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2EnableTunnelLifecycleControlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tunnel2EnableTunnelLifecycleControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2IkeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2IkeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2IkeVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2IkeVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2InsideCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2InsideCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2InsideCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2InsideCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2InsideIpv6Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2InsideIpv6Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2InsideIpv6CidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2InsideIpv6CidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2LogOptions() TfConnection_Tunnel2LogOptionsPropertyOutputReference {
	var returns TfConnection_Tunnel2LogOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"tunnel2LogOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2LogOptionsInput() *TfConnection_Tunnel2LogOptionsProperty {
	var returns *TfConnection_Tunnel2LogOptionsProperty
	_jsii_.Get(
		j,
		"tunnel2LogOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Phase1DhGroupNumbers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel2Phase1DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Phase1DhGroupNumbersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel2Phase1DhGroupNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Phase1EncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase1EncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Phase1EncryptionAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase1EncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Phase1IntegrityAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase1IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Phase1IntegrityAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase1IntegrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Phase1LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2Phase1LifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Phase1LifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2Phase1LifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Phase2DhGroupNumbers() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel2Phase2DhGroupNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Phase2DhGroupNumbersInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"tunnel2Phase2DhGroupNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Phase2EncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase2EncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Phase2EncryptionAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase2EncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Phase2IntegrityAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase2IntegrityAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Phase2IntegrityAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tunnel2Phase2IntegrityAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Phase2LifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2Phase2LifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2Phase2LifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2Phase2LifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2PresharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2PresharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2PresharedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2PresharedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2RekeyFuzzPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2RekeyFuzzPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2RekeyFuzzPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2RekeyFuzzPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2RekeyMarginTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2RekeyMarginTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2RekeyMarginTimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2RekeyMarginTimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2ReplayWindowSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2ReplayWindowSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2ReplayWindowSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tunnel2ReplayWindowSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2StartupAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2StartupAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2StartupActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2StartupActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Tunnel2VgwInsideAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnel2VgwInsideAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) TunnelBandwidth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelBandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) TunnelBandwidthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelBandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) TunnelInsideIpVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelInsideIpVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) TunnelInsideIpVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelInsideIpVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) VgwTelemetry() TfConnection_VgwTelemetryPropertyList {
	var returns TfConnection_VgwTelemetryPropertyList
	_jsii_.Get(
		j,
		"vgwTelemetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) VpnConcentratorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnConcentratorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) VpnConcentratorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnConcentratorIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) VpnGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection) VpnGatewayIdInput() *string {
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
func NewTfConnection(scope constructs.Construct, id *string, config *TfConnectionConfig) TfConnection {
	_init_.Initialize()

	if err := validateNewTfConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConnection{}

	_jsii_.Create(
		"@cdktn/aws-vpn-site-to-site.TfConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection aws_vpn_connection} Resource.
// Experimental.
func NewTfConnection_Override(t TfConnection, scope constructs.Construct, id *string, config *TfConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpn-site-to-site.TfConnection",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetCustomerGatewayId(val *string) {
	if err := j.validateSetCustomerGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerGatewayId",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetEnableAcceleration(val interface{}) {
	if err := j.validateSetEnableAccelerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAcceleration",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetLocalIpv4NetworkCidr(val *string) {
	if err := j.validateSetLocalIpv4NetworkCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localIpv4NetworkCidr",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetLocalIpv6NetworkCidr(val *string) {
	if err := j.validateSetLocalIpv6NetworkCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localIpv6NetworkCidr",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetOutsideIpAddressType(val *string) {
	if err := j.validateSetOutsideIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outsideIpAddressType",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetPresharedKeyStorage(val *string) {
	if err := j.validateSetPresharedKeyStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"presharedKeyStorage",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetRemoteIpv4NetworkCidr(val *string) {
	if err := j.validateSetRemoteIpv4NetworkCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteIpv4NetworkCidr",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetRemoteIpv6NetworkCidr(val *string) {
	if err := j.validateSetRemoteIpv6NetworkCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteIpv6NetworkCidr",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetStaticRoutesOnly(val interface{}) {
	if err := j.validateSetStaticRoutesOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticRoutesOnly",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTransitGatewayId(val *string) {
	if err := j.validateSetTransitGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitGatewayId",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTransportTransitGatewayAttachmentId(val *string) {
	if err := j.validateSetTransportTransitGatewayAttachmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportTransitGatewayAttachmentId",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1DpdTimeoutAction(val *string) {
	if err := j.validateSetTunnel1DpdTimeoutActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1DpdTimeoutAction",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1DpdTimeoutSeconds(val *float64) {
	if err := j.validateSetTunnel1DpdTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1DpdTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1EnableTunnelLifecycleControl(val interface{}) {
	if err := j.validateSetTunnel1EnableTunnelLifecycleControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1EnableTunnelLifecycleControl",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1IkeVersions(val *[]*string) {
	if err := j.validateSetTunnel1IkeVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1IkeVersions",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1InsideCidr(val *string) {
	if err := j.validateSetTunnel1InsideCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1InsideCidr",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1InsideIpv6Cidr(val *string) {
	if err := j.validateSetTunnel1InsideIpv6CidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1InsideIpv6Cidr",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1Phase1DhGroupNumbers(val *[]*float64) {
	if err := j.validateSetTunnel1Phase1DhGroupNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase1DhGroupNumbers",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1Phase1EncryptionAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel1Phase1EncryptionAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase1EncryptionAlgorithms",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1Phase1IntegrityAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel1Phase1IntegrityAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase1IntegrityAlgorithms",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1Phase1LifetimeSeconds(val *float64) {
	if err := j.validateSetTunnel1Phase1LifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase1LifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1Phase2DhGroupNumbers(val *[]*float64) {
	if err := j.validateSetTunnel1Phase2DhGroupNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase2DhGroupNumbers",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1Phase2EncryptionAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel1Phase2EncryptionAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase2EncryptionAlgorithms",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1Phase2IntegrityAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel1Phase2IntegrityAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase2IntegrityAlgorithms",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1Phase2LifetimeSeconds(val *float64) {
	if err := j.validateSetTunnel1Phase2LifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1Phase2LifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1PresharedKey(val *string) {
	if err := j.validateSetTunnel1PresharedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1PresharedKey",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1RekeyFuzzPercentage(val *float64) {
	if err := j.validateSetTunnel1RekeyFuzzPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1RekeyFuzzPercentage",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1RekeyMarginTimeSeconds(val *float64) {
	if err := j.validateSetTunnel1RekeyMarginTimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1RekeyMarginTimeSeconds",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1ReplayWindowSize(val *float64) {
	if err := j.validateSetTunnel1ReplayWindowSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1ReplayWindowSize",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel1StartupAction(val *string) {
	if err := j.validateSetTunnel1StartupActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel1StartupAction",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2DpdTimeoutAction(val *string) {
	if err := j.validateSetTunnel2DpdTimeoutActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2DpdTimeoutAction",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2DpdTimeoutSeconds(val *float64) {
	if err := j.validateSetTunnel2DpdTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2DpdTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2EnableTunnelLifecycleControl(val interface{}) {
	if err := j.validateSetTunnel2EnableTunnelLifecycleControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2EnableTunnelLifecycleControl",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2IkeVersions(val *[]*string) {
	if err := j.validateSetTunnel2IkeVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2IkeVersions",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2InsideCidr(val *string) {
	if err := j.validateSetTunnel2InsideCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2InsideCidr",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2InsideIpv6Cidr(val *string) {
	if err := j.validateSetTunnel2InsideIpv6CidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2InsideIpv6Cidr",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2Phase1DhGroupNumbers(val *[]*float64) {
	if err := j.validateSetTunnel2Phase1DhGroupNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase1DhGroupNumbers",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2Phase1EncryptionAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel2Phase1EncryptionAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase1EncryptionAlgorithms",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2Phase1IntegrityAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel2Phase1IntegrityAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase1IntegrityAlgorithms",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2Phase1LifetimeSeconds(val *float64) {
	if err := j.validateSetTunnel2Phase1LifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase1LifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2Phase2DhGroupNumbers(val *[]*float64) {
	if err := j.validateSetTunnel2Phase2DhGroupNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase2DhGroupNumbers",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2Phase2EncryptionAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel2Phase2EncryptionAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase2EncryptionAlgorithms",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2Phase2IntegrityAlgorithms(val *[]*string) {
	if err := j.validateSetTunnel2Phase2IntegrityAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase2IntegrityAlgorithms",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2Phase2LifetimeSeconds(val *float64) {
	if err := j.validateSetTunnel2Phase2LifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2Phase2LifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2PresharedKey(val *string) {
	if err := j.validateSetTunnel2PresharedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2PresharedKey",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2RekeyFuzzPercentage(val *float64) {
	if err := j.validateSetTunnel2RekeyFuzzPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2RekeyFuzzPercentage",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2RekeyMarginTimeSeconds(val *float64) {
	if err := j.validateSetTunnel2RekeyMarginTimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2RekeyMarginTimeSeconds",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2ReplayWindowSize(val *float64) {
	if err := j.validateSetTunnel2ReplayWindowSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2ReplayWindowSize",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnel2StartupAction(val *string) {
	if err := j.validateSetTunnel2StartupActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnel2StartupAction",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnelBandwidth(val *string) {
	if err := j.validateSetTunnelBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelBandwidth",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetTunnelInsideIpVersion(val *string) {
	if err := j.validateSetTunnelInsideIpVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelInsideIpVersion",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetVpnConcentratorId(val *string) {
	if err := j.validateSetVpnConcentratorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnConcentratorId",
		val,
	)
}

func (j *jsiiProxy_TfConnection)SetVpnGatewayId(val *string) {
	if err := j.validateSetVpnGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnGatewayId",
		val,
	)
}

// Generates CDKTN code for importing a TfConnection resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfConnection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfConnection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-site-to-site.TfConnection",
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
func TfConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-site-to-site.TfConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-site-to-site.TfConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-site-to-site.TfConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-vpn-site-to-site.TfConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfConnection) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfConnection) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConnection) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConnection) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConnection) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConnection) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfConnection) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnection) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfConnection) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfConnection) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfConnection) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfConnection) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfConnection) PutTunnel1LogOptions(value *TfConnection_Tunnel1LogOptionsProperty) {
	if err := t.validatePutTunnel1LogOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTunnel1LogOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnection) PutTunnel2LogOptions(value *TfConnection_Tunnel2LogOptionsProperty) {
	if err := t.validatePutTunnel2LogOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTunnel2LogOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnection) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfConnection) ResetEnableAcceleration() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableAcceleration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetLocalIpv4NetworkCidr() {
	_jsii_.InvokeVoid(
		t,
		"resetLocalIpv4NetworkCidr",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetLocalIpv6NetworkCidr() {
	_jsii_.InvokeVoid(
		t,
		"resetLocalIpv6NetworkCidr",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetOutsideIpAddressType() {
	_jsii_.InvokeVoid(
		t,
		"resetOutsideIpAddressType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetPresharedKeyStorage() {
	_jsii_.InvokeVoid(
		t,
		"resetPresharedKeyStorage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetRemoteIpv4NetworkCidr() {
	_jsii_.InvokeVoid(
		t,
		"resetRemoteIpv4NetworkCidr",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetRemoteIpv6NetworkCidr() {
	_jsii_.InvokeVoid(
		t,
		"resetRemoteIpv6NetworkCidr",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetStaticRoutesOnly() {
	_jsii_.InvokeVoid(
		t,
		"resetStaticRoutesOnly",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTransitGatewayId() {
	_jsii_.InvokeVoid(
		t,
		"resetTransitGatewayId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTransportTransitGatewayAttachmentId() {
	_jsii_.InvokeVoid(
		t,
		"resetTransportTransitGatewayAttachmentId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1DpdTimeoutAction() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1DpdTimeoutAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1DpdTimeoutSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1DpdTimeoutSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1EnableTunnelLifecycleControl() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1EnableTunnelLifecycleControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1IkeVersions() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1IkeVersions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1InsideCidr() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1InsideCidr",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1InsideIpv6Cidr() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1InsideIpv6Cidr",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1LogOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1LogOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1Phase1DhGroupNumbers() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1Phase1DhGroupNumbers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1Phase1EncryptionAlgorithms() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1Phase1EncryptionAlgorithms",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1Phase1IntegrityAlgorithms() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1Phase1IntegrityAlgorithms",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1Phase1LifetimeSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1Phase1LifetimeSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1Phase2DhGroupNumbers() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1Phase2DhGroupNumbers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1Phase2EncryptionAlgorithms() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1Phase2EncryptionAlgorithms",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1Phase2IntegrityAlgorithms() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1Phase2IntegrityAlgorithms",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1Phase2LifetimeSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1Phase2LifetimeSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1PresharedKey() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1PresharedKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1RekeyFuzzPercentage() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1RekeyFuzzPercentage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1RekeyMarginTimeSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1RekeyMarginTimeSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1ReplayWindowSize() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1ReplayWindowSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel1StartupAction() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel1StartupAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2DpdTimeoutAction() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2DpdTimeoutAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2DpdTimeoutSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2DpdTimeoutSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2EnableTunnelLifecycleControl() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2EnableTunnelLifecycleControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2IkeVersions() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2IkeVersions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2InsideCidr() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2InsideCidr",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2InsideIpv6Cidr() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2InsideIpv6Cidr",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2LogOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2LogOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2Phase1DhGroupNumbers() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2Phase1DhGroupNumbers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2Phase1EncryptionAlgorithms() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2Phase1EncryptionAlgorithms",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2Phase1IntegrityAlgorithms() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2Phase1IntegrityAlgorithms",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2Phase1LifetimeSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2Phase1LifetimeSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2Phase2DhGroupNumbers() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2Phase2DhGroupNumbers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2Phase2EncryptionAlgorithms() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2Phase2EncryptionAlgorithms",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2Phase2IntegrityAlgorithms() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2Phase2IntegrityAlgorithms",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2Phase2LifetimeSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2Phase2LifetimeSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2PresharedKey() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2PresharedKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2RekeyFuzzPercentage() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2RekeyFuzzPercentage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2RekeyMarginTimeSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2RekeyMarginTimeSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2ReplayWindowSize() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2ReplayWindowSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnel2StartupAction() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnel2StartupAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnelBandwidth() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnelBandwidth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetTunnelInsideIpVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetTunnelInsideIpVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetVpnConcentratorId() {
	_jsii_.InvokeVoid(
		t,
		"resetVpnConcentratorId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) ResetVpnGatewayId() {
	_jsii_.InvokeVoid(
		t,
		"resetVpnGatewayId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

