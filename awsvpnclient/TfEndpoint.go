package awsvpnclient

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpnclient/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsvpnclient/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint aws_ec2_client_vpn_endpoint}.
// Experimental.
type TfEndpoint interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	AuthenticationOptions() TfEndpoint_AuthenticationOptionsPropertyList
	// Experimental.
	AuthenticationOptionsInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClientCidrBlock() *string
	// Experimental.
	SetClientCidrBlock(val *string)
	// Experimental.
	ClientCidrBlockInput() *string
	// Experimental.
	ClientConnectOptions() TfEndpoint_ClientConnectOptionsPropertyOutputReference
	// Experimental.
	ClientConnectOptionsInput() *TfEndpoint_ClientConnectOptionsProperty
	// Experimental.
	ClientLoginBannerOptions() TfEndpoint_ClientLoginBannerOptionsPropertyOutputReference
	// Experimental.
	ClientLoginBannerOptionsInput() *TfEndpoint_ClientLoginBannerOptionsProperty
	// Experimental.
	ClientRouteEnforcementOptions() TfEndpoint_ClientRouteEnforcementOptionsPropertyOutputReference
	// Experimental.
	ClientRouteEnforcementOptionsInput() *TfEndpoint_ClientRouteEnforcementOptionsProperty
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConnectionLogOptions() TfEndpoint_ConnectionLogOptionsPropertyOutputReference
	// Experimental.
	ConnectionLogOptionsInput() *TfEndpoint_ConnectionLogOptionsProperty
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
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	DisconnectOnSessionTimeout() interface{}
	// Experimental.
	SetDisconnectOnSessionTimeout(val interface{})
	// Experimental.
	DisconnectOnSessionTimeoutInput() interface{}
	// Experimental.
	DnsName() *string
	// Experimental.
	DnsServers() *[]*string
	// Experimental.
	SetDnsServers(val *[]*string)
	// Experimental.
	DnsServersInput() *[]*string
	// Experimental.
	EndpointIpAddressType() *string
	// Experimental.
	SetEndpointIpAddressType(val *string)
	// Experimental.
	EndpointIpAddressTypeInput() *string
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
	SecurityGroupIds() *[]*string
	// Experimental.
	SetSecurityGroupIds(val *[]*string)
	// Experimental.
	SecurityGroupIdsInput() *[]*string
	// Experimental.
	SelfServicePortal() *string
	// Experimental.
	SetSelfServicePortal(val *string)
	// Experimental.
	SelfServicePortalInput() *string
	// Experimental.
	SelfServicePortalUrl() *string
	// Experimental.
	ServerCertificateArn() *string
	// Experimental.
	SetServerCertificateArn(val *string)
	// Experimental.
	ServerCertificateArnInput() *string
	// Experimental.
	SessionTimeoutHours() *float64
	// Experimental.
	SetSessionTimeoutHours(val *float64)
	// Experimental.
	SessionTimeoutHoursInput() *float64
	// Experimental.
	SplitTunnel() interface{}
	// Experimental.
	SetSplitTunnel(val interface{})
	// Experimental.
	SplitTunnelInput() interface{}
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
	TrafficIpAddressType() *string
	// Experimental.
	SetTrafficIpAddressType(val *string)
	// Experimental.
	TrafficIpAddressTypeInput() *string
	// Experimental.
	TransitGatewayConfiguration() TfEndpoint_TransitGatewayConfigurationPropertyOutputReference
	// Experimental.
	TransitGatewayConfigurationInput() *TfEndpoint_TransitGatewayConfigurationProperty
	// Experimental.
	TransportProtocol() *string
	// Experimental.
	SetTransportProtocol(val *string)
	// Experimental.
	TransportProtocolInput() *string
	// Experimental.
	VpcId() *string
	// Experimental.
	SetVpcId(val *string)
	// Experimental.
	VpcIdInput() *string
	// Experimental.
	VpnPort() *float64
	// Experimental.
	SetVpnPort(val *float64)
	// Experimental.
	VpnPortInput() *float64
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
	PutAuthenticationOptions(value interface{})
	// Experimental.
	PutClientConnectOptions(value *TfEndpoint_ClientConnectOptionsProperty)
	// Experimental.
	PutClientLoginBannerOptions(value *TfEndpoint_ClientLoginBannerOptionsProperty)
	// Experimental.
	PutClientRouteEnforcementOptions(value *TfEndpoint_ClientRouteEnforcementOptionsProperty)
	// Experimental.
	PutConnectionLogOptions(value *TfEndpoint_ConnectionLogOptionsProperty)
	// Experimental.
	PutTransitGatewayConfiguration(value *TfEndpoint_TransitGatewayConfigurationProperty)
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
	ResetClientCidrBlock()
	// Experimental.
	ResetClientConnectOptions()
	// Experimental.
	ResetClientLoginBannerOptions()
	// Experimental.
	ResetClientRouteEnforcementOptions()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetDisconnectOnSessionTimeout()
	// Experimental.
	ResetDnsServers()
	// Experimental.
	ResetEndpointIpAddressType()
	// Experimental.
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSecurityGroupIds()
	// Experimental.
	ResetSelfServicePortal()
	// Experimental.
	ResetSessionTimeoutHours()
	// Experimental.
	ResetSplitTunnel()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTrafficIpAddressType()
	// Experimental.
	ResetTransitGatewayConfiguration()
	// Experimental.
	ResetTransportProtocol()
	// Experimental.
	ResetVpcId()
	// Experimental.
	ResetVpnPort()
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

// The jsii proxy struct for TfEndpoint
type jsiiProxy_TfEndpoint struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfEndpoint) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) AuthenticationOptions() TfEndpoint_AuthenticationOptionsPropertyList {
	var returns TfEndpoint_AuthenticationOptionsPropertyList
	_jsii_.Get(
		j,
		"authenticationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) AuthenticationOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ClientCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ClientCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ClientConnectOptions() TfEndpoint_ClientConnectOptionsPropertyOutputReference {
	var returns TfEndpoint_ClientConnectOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"clientConnectOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ClientConnectOptionsInput() *TfEndpoint_ClientConnectOptionsProperty {
	var returns *TfEndpoint_ClientConnectOptionsProperty
	_jsii_.Get(
		j,
		"clientConnectOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ClientLoginBannerOptions() TfEndpoint_ClientLoginBannerOptionsPropertyOutputReference {
	var returns TfEndpoint_ClientLoginBannerOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"clientLoginBannerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ClientLoginBannerOptionsInput() *TfEndpoint_ClientLoginBannerOptionsProperty {
	var returns *TfEndpoint_ClientLoginBannerOptionsProperty
	_jsii_.Get(
		j,
		"clientLoginBannerOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ClientRouteEnforcementOptions() TfEndpoint_ClientRouteEnforcementOptionsPropertyOutputReference {
	var returns TfEndpoint_ClientRouteEnforcementOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"clientRouteEnforcementOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ClientRouteEnforcementOptionsInput() *TfEndpoint_ClientRouteEnforcementOptionsProperty {
	var returns *TfEndpoint_ClientRouteEnforcementOptionsProperty
	_jsii_.Get(
		j,
		"clientRouteEnforcementOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ConnectionLogOptions() TfEndpoint_ConnectionLogOptionsPropertyOutputReference {
	var returns TfEndpoint_ConnectionLogOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"connectionLogOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ConnectionLogOptionsInput() *TfEndpoint_ConnectionLogOptionsProperty {
	var returns *TfEndpoint_ConnectionLogOptionsProperty
	_jsii_.Get(
		j,
		"connectionLogOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) DisconnectOnSessionTimeout() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disconnectOnSessionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) DisconnectOnSessionTimeoutInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disconnectOnSessionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) DnsServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) EndpointIpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIpAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) EndpointIpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIpAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) SelfServicePortal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServicePortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) SelfServicePortalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServicePortalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) SelfServicePortalUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServicePortalUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ServerCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ServerCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) SessionTimeoutHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) SessionTimeoutHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) SplitTunnel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splitTunnel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) SplitTunnelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splitTunnelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TrafficIpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficIpAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TrafficIpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficIpAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TransitGatewayConfiguration() TfEndpoint_TransitGatewayConfigurationPropertyOutputReference {
	var returns TfEndpoint_TransitGatewayConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"transitGatewayConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TransitGatewayConfigurationInput() *TfEndpoint_TransitGatewayConfigurationProperty {
	var returns *TfEndpoint_TransitGatewayConfigurationProperty
	_jsii_.Get(
		j,
		"transitGatewayConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TransportProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TransportProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) VpnPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpnPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) VpnPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpnPortInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint aws_ec2_client_vpn_endpoint} Resource.
// Experimental.
func NewTfEndpoint(scope constructs.Construct, id *string, config *TfEndpointConfig) TfEndpoint {
	_init_.Initialize()

	if err := validateNewTfEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpoint{}

	_jsii_.Create(
		"@cdktn/aws-vpn-client.TfEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint aws_ec2_client_vpn_endpoint} Resource.
// Experimental.
func NewTfEndpoint_Override(t TfEndpoint, scope constructs.Construct, id *string, config *TfEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpn-client.TfEndpoint",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfEndpoint)SetClientCidrBlock(val *string) {
	if err := j.validateSetClientCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCidrBlock",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetDisconnectOnSessionTimeout(val interface{}) {
	if err := j.validateSetDisconnectOnSessionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disconnectOnSessionTimeout",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetDnsServers(val *[]*string) {
	if err := j.validateSetDnsServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServers",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetEndpointIpAddressType(val *string) {
	if err := j.validateSetEndpointIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointIpAddressType",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetSelfServicePortal(val *string) {
	if err := j.validateSetSelfServicePortalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfServicePortal",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetServerCertificateArn(val *string) {
	if err := j.validateSetServerCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCertificateArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetSessionTimeoutHours(val *float64) {
	if err := j.validateSetSessionTimeoutHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionTimeoutHours",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetSplitTunnel(val interface{}) {
	if err := j.validateSetSplitTunnelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"splitTunnel",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetTrafficIpAddressType(val *string) {
	if err := j.validateSetTrafficIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficIpAddressType",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetTransportProtocol(val *string) {
	if err := j.validateSetTransportProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportProtocol",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetVpnPort(val *float64) {
	if err := j.validateSetVpnPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnPort",
		val,
	)
}

// Generates CDKTN code for importing a TfEndpoint resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-client.TfEndpoint",
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
func TfEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-client.TfEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-client.TfEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-client.TfEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-vpn-client.TfEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfEndpoint) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfEndpoint) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfEndpoint) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfEndpoint) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfEndpoint) PutAuthenticationOptions(value interface{}) {
	if err := t.validatePutAuthenticationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuthenticationOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint) PutClientConnectOptions(value *TfEndpoint_ClientConnectOptionsProperty) {
	if err := t.validatePutClientConnectOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putClientConnectOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint) PutClientLoginBannerOptions(value *TfEndpoint_ClientLoginBannerOptionsProperty) {
	if err := t.validatePutClientLoginBannerOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putClientLoginBannerOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint) PutClientRouteEnforcementOptions(value *TfEndpoint_ClientRouteEnforcementOptionsProperty) {
	if err := t.validatePutClientRouteEnforcementOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putClientRouteEnforcementOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint) PutConnectionLogOptions(value *TfEndpoint_ConnectionLogOptionsProperty) {
	if err := t.validatePutConnectionLogOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConnectionLogOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint) PutTransitGatewayConfiguration(value *TfEndpoint_TransitGatewayConfigurationProperty) {
	if err := t.validatePutTransitGatewayConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTransitGatewayConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfEndpoint) ResetClientCidrBlock() {
	_jsii_.InvokeVoid(
		t,
		"resetClientCidrBlock",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetClientConnectOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetClientConnectOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetClientLoginBannerOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetClientLoginBannerOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetClientRouteEnforcementOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetClientRouteEnforcementOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetDisconnectOnSessionTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetDisconnectOnSessionTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetDnsServers() {
	_jsii_.InvokeVoid(
		t,
		"resetDnsServers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetEndpointIpAddressType() {
	_jsii_.InvokeVoid(
		t,
		"resetEndpointIpAddressType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetSelfServicePortal() {
	_jsii_.InvokeVoid(
		t,
		"resetSelfServicePortal",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetSessionTimeoutHours() {
	_jsii_.InvokeVoid(
		t,
		"resetSessionTimeoutHours",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetSplitTunnel() {
	_jsii_.InvokeVoid(
		t,
		"resetSplitTunnel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetTrafficIpAddressType() {
	_jsii_.InvokeVoid(
		t,
		"resetTrafficIpAddressType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetTransitGatewayConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetTransitGatewayConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetTransportProtocol() {
	_jsii_.InvokeVoid(
		t,
		"resetTransportProtocol",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetVpcId() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetVpnPort() {
	_jsii_.InvokeVoid(
		t,
		"resetVpnPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

