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
type AwsEc2ClientVpnEndpoint interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	AuthenticationOptions() AwsEc2ClientVpnEndpoint_AuthenticationOptionsPropertyList
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
	ClientConnectOptions() AwsEc2ClientVpnEndpoint_ClientConnectOptionsPropertyOutputReference
	// Experimental.
	ClientConnectOptionsInput() *AwsEc2ClientVpnEndpoint_ClientConnectOptionsProperty
	// Experimental.
	ClientLoginBannerOptions() AwsEc2ClientVpnEndpoint_ClientLoginBannerOptionsPropertyOutputReference
	// Experimental.
	ClientLoginBannerOptionsInput() *AwsEc2ClientVpnEndpoint_ClientLoginBannerOptionsProperty
	// Experimental.
	ClientRouteEnforcementOptions() AwsEc2ClientVpnEndpoint_ClientRouteEnforcementOptionsPropertyOutputReference
	// Experimental.
	ClientRouteEnforcementOptionsInput() *AwsEc2ClientVpnEndpoint_ClientRouteEnforcementOptionsProperty
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConnectionLogOptions() AwsEc2ClientVpnEndpoint_ConnectionLogOptionsPropertyOutputReference
	// Experimental.
	ConnectionLogOptionsInput() *AwsEc2ClientVpnEndpoint_ConnectionLogOptionsProperty
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
	TransitGatewayConfiguration() AwsEc2ClientVpnEndpoint_TransitGatewayConfigurationPropertyOutputReference
	// Experimental.
	TransitGatewayConfigurationInput() *AwsEc2ClientVpnEndpoint_TransitGatewayConfigurationProperty
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
	PutClientConnectOptions(value *AwsEc2ClientVpnEndpoint_ClientConnectOptionsProperty)
	// Experimental.
	PutClientLoginBannerOptions(value *AwsEc2ClientVpnEndpoint_ClientLoginBannerOptionsProperty)
	// Experimental.
	PutClientRouteEnforcementOptions(value *AwsEc2ClientVpnEndpoint_ClientRouteEnforcementOptionsProperty)
	// Experimental.
	PutConnectionLogOptions(value *AwsEc2ClientVpnEndpoint_ConnectionLogOptionsProperty)
	// Experimental.
	PutTransitGatewayConfiguration(value *AwsEc2ClientVpnEndpoint_TransitGatewayConfigurationProperty)
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

// The jsii proxy struct for AwsEc2ClientVpnEndpoint
type jsiiProxy_AwsEc2ClientVpnEndpoint struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) AuthenticationOptions() AwsEc2ClientVpnEndpoint_AuthenticationOptionsPropertyList {
	var returns AwsEc2ClientVpnEndpoint_AuthenticationOptionsPropertyList
	_jsii_.Get(
		j,
		"authenticationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) AuthenticationOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) ClientCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) ClientCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) ClientConnectOptions() AwsEc2ClientVpnEndpoint_ClientConnectOptionsPropertyOutputReference {
	var returns AwsEc2ClientVpnEndpoint_ClientConnectOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"clientConnectOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) ClientConnectOptionsInput() *AwsEc2ClientVpnEndpoint_ClientConnectOptionsProperty {
	var returns *AwsEc2ClientVpnEndpoint_ClientConnectOptionsProperty
	_jsii_.Get(
		j,
		"clientConnectOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) ClientLoginBannerOptions() AwsEc2ClientVpnEndpoint_ClientLoginBannerOptionsPropertyOutputReference {
	var returns AwsEc2ClientVpnEndpoint_ClientLoginBannerOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"clientLoginBannerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) ClientLoginBannerOptionsInput() *AwsEc2ClientVpnEndpoint_ClientLoginBannerOptionsProperty {
	var returns *AwsEc2ClientVpnEndpoint_ClientLoginBannerOptionsProperty
	_jsii_.Get(
		j,
		"clientLoginBannerOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) ClientRouteEnforcementOptions() AwsEc2ClientVpnEndpoint_ClientRouteEnforcementOptionsPropertyOutputReference {
	var returns AwsEc2ClientVpnEndpoint_ClientRouteEnforcementOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"clientRouteEnforcementOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) ClientRouteEnforcementOptionsInput() *AwsEc2ClientVpnEndpoint_ClientRouteEnforcementOptionsProperty {
	var returns *AwsEc2ClientVpnEndpoint_ClientRouteEnforcementOptionsProperty
	_jsii_.Get(
		j,
		"clientRouteEnforcementOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) ConnectionLogOptions() AwsEc2ClientVpnEndpoint_ConnectionLogOptionsPropertyOutputReference {
	var returns AwsEc2ClientVpnEndpoint_ConnectionLogOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"connectionLogOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) ConnectionLogOptionsInput() *AwsEc2ClientVpnEndpoint_ConnectionLogOptionsProperty {
	var returns *AwsEc2ClientVpnEndpoint_ConnectionLogOptionsProperty
	_jsii_.Get(
		j,
		"connectionLogOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) DisconnectOnSessionTimeout() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disconnectOnSessionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) DisconnectOnSessionTimeoutInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disconnectOnSessionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) DnsServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) EndpointIpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIpAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) EndpointIpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIpAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) SelfServicePortal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServicePortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) SelfServicePortalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServicePortalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) SelfServicePortalUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServicePortalUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) ServerCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) ServerCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) SessionTimeoutHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) SessionTimeoutHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) SplitTunnel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splitTunnel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) SplitTunnelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splitTunnelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) TrafficIpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficIpAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) TrafficIpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficIpAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) TransitGatewayConfiguration() AwsEc2ClientVpnEndpoint_TransitGatewayConfigurationPropertyOutputReference {
	var returns AwsEc2ClientVpnEndpoint_TransitGatewayConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"transitGatewayConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) TransitGatewayConfigurationInput() *AwsEc2ClientVpnEndpoint_TransitGatewayConfigurationProperty {
	var returns *AwsEc2ClientVpnEndpoint_TransitGatewayConfigurationProperty
	_jsii_.Get(
		j,
		"transitGatewayConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) TransportProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) TransportProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) VpnPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpnPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint) VpnPortInput() *float64 {
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
func NewAwsEc2ClientVpnEndpoint(scope constructs.Construct, id *string, config *AwsEc2ClientVpnEndpointConfig) AwsEc2ClientVpnEndpoint {
	_init_.Initialize()

	if err := validateNewAwsEc2ClientVpnEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEc2ClientVpnEndpoint{}

	_jsii_.Create(
		"@cdktn/aws-vpn-client.AwsEc2ClientVpnEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint aws_ec2_client_vpn_endpoint} Resource.
// Experimental.
func NewAwsEc2ClientVpnEndpoint_Override(a AwsEc2ClientVpnEndpoint, scope constructs.Construct, id *string, config *AwsEc2ClientVpnEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpn-client.AwsEc2ClientVpnEndpoint",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetClientCidrBlock(val *string) {
	if err := j.validateSetClientCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCidrBlock",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetDisconnectOnSessionTimeout(val interface{}) {
	if err := j.validateSetDisconnectOnSessionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disconnectOnSessionTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetDnsServers(val *[]*string) {
	if err := j.validateSetDnsServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServers",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetEndpointIpAddressType(val *string) {
	if err := j.validateSetEndpointIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointIpAddressType",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetSelfServicePortal(val *string) {
	if err := j.validateSetSelfServicePortalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfServicePortal",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetServerCertificateArn(val *string) {
	if err := j.validateSetServerCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCertificateArn",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetSessionTimeoutHours(val *float64) {
	if err := j.validateSetSessionTimeoutHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionTimeoutHours",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetSplitTunnel(val interface{}) {
	if err := j.validateSetSplitTunnelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"splitTunnel",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetTrafficIpAddressType(val *string) {
	if err := j.validateSetTrafficIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficIpAddressType",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetTransportProtocol(val *string) {
	if err := j.validateSetTransportProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportProtocol",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func (j *jsiiProxy_AwsEc2ClientVpnEndpoint)SetVpnPort(val *float64) {
	if err := j.validateSetVpnPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnPort",
		val,
	)
}

// Generates CDKTN code for importing a AwsEc2ClientVpnEndpoint resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsEc2ClientVpnEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsEc2ClientVpnEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-client.AwsEc2ClientVpnEndpoint",
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
func AwsEc2ClientVpnEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEc2ClientVpnEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-client.AwsEc2ClientVpnEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsEc2ClientVpnEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEc2ClientVpnEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-client.AwsEc2ClientVpnEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsEc2ClientVpnEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEc2ClientVpnEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-client.AwsEc2ClientVpnEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsEc2ClientVpnEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-vpn-client.AwsEc2ClientVpnEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) PutAuthenticationOptions(value interface{}) {
	if err := a.validatePutAuthenticationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthenticationOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) PutClientConnectOptions(value *AwsEc2ClientVpnEndpoint_ClientConnectOptionsProperty) {
	if err := a.validatePutClientConnectOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putClientConnectOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) PutClientLoginBannerOptions(value *AwsEc2ClientVpnEndpoint_ClientLoginBannerOptionsProperty) {
	if err := a.validatePutClientLoginBannerOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putClientLoginBannerOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) PutClientRouteEnforcementOptions(value *AwsEc2ClientVpnEndpoint_ClientRouteEnforcementOptionsProperty) {
	if err := a.validatePutClientRouteEnforcementOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putClientRouteEnforcementOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) PutConnectionLogOptions(value *AwsEc2ClientVpnEndpoint_ConnectionLogOptionsProperty) {
	if err := a.validatePutConnectionLogOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectionLogOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) PutTransitGatewayConfiguration(value *AwsEc2ClientVpnEndpoint_TransitGatewayConfigurationProperty) {
	if err := a.validatePutTransitGatewayConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTransitGatewayConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetClientCidrBlock() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCidrBlock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetClientConnectOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetClientConnectOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetClientLoginBannerOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetClientLoginBannerOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetClientRouteEnforcementOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetClientRouteEnforcementOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetDisconnectOnSessionTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetDisconnectOnSessionTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetDnsServers() {
	_jsii_.InvokeVoid(
		a,
		"resetDnsServers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetEndpointIpAddressType() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointIpAddressType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetSelfServicePortal() {
	_jsii_.InvokeVoid(
		a,
		"resetSelfServicePortal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetSessionTimeoutHours() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionTimeoutHours",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetSplitTunnel() {
	_jsii_.InvokeVoid(
		a,
		"resetSplitTunnel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetTrafficIpAddressType() {
	_jsii_.InvokeVoid(
		a,
		"resetTrafficIpAddressType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetTransitGatewayConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetTransitGatewayConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetTransportProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetTransportProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetVpcId() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ResetVpnPort() {
	_jsii_.InvokeVoid(
		a,
		"resetVpnPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEc2ClientVpnEndpoint) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

