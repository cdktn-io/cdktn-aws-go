package vpnclient

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/vpnclient/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/vpnclient/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint aws_ec2_client_vpn_endpoint}.
// Experimental.
type AwsEndpoint interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	AuthenticationOptions() AwsEndpoint_AuthenticationOptionsPropertyList
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
	ClientConnectOptions() AwsEndpoint_ClientConnectOptionsPropertyOutputReference
	// Experimental.
	ClientConnectOptionsInput() *AwsEndpoint_ClientConnectOptionsProperty
	// Experimental.
	ClientLoginBannerOptions() AwsEndpoint_ClientLoginBannerOptionsPropertyOutputReference
	// Experimental.
	ClientLoginBannerOptionsInput() *AwsEndpoint_ClientLoginBannerOptionsProperty
	// Experimental.
	ClientRouteEnforcementOptions() AwsEndpoint_ClientRouteEnforcementOptionsPropertyOutputReference
	// Experimental.
	ClientRouteEnforcementOptionsInput() *AwsEndpoint_ClientRouteEnforcementOptionsProperty
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConnectionLogOptions() AwsEndpoint_ConnectionLogOptionsPropertyOutputReference
	// Experimental.
	ConnectionLogOptionsInput() *AwsEndpoint_ConnectionLogOptionsProperty
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
	TransitGatewayConfiguration() AwsEndpoint_TransitGatewayConfigurationPropertyOutputReference
	// Experimental.
	TransitGatewayConfigurationInput() *AwsEndpoint_TransitGatewayConfigurationProperty
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
	PutClientConnectOptions(value *AwsEndpoint_ClientConnectOptionsProperty)
	// Experimental.
	PutClientLoginBannerOptions(value *AwsEndpoint_ClientLoginBannerOptionsProperty)
	// Experimental.
	PutClientRouteEnforcementOptions(value *AwsEndpoint_ClientRouteEnforcementOptionsProperty)
	// Experimental.
	PutConnectionLogOptions(value *AwsEndpoint_ConnectionLogOptionsProperty)
	// Experimental.
	PutTransitGatewayConfiguration(value *AwsEndpoint_TransitGatewayConfigurationProperty)
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

// The jsii proxy struct for AwsEndpoint
type jsiiProxy_AwsEndpoint struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsEndpoint) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) AuthenticationOptions() AwsEndpoint_AuthenticationOptionsPropertyList {
	var returns AwsEndpoint_AuthenticationOptionsPropertyList
	_jsii_.Get(
		j,
		"authenticationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) AuthenticationOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) ClientCidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) ClientCidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) ClientConnectOptions() AwsEndpoint_ClientConnectOptionsPropertyOutputReference {
	var returns AwsEndpoint_ClientConnectOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"clientConnectOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) ClientConnectOptionsInput() *AwsEndpoint_ClientConnectOptionsProperty {
	var returns *AwsEndpoint_ClientConnectOptionsProperty
	_jsii_.Get(
		j,
		"clientConnectOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) ClientLoginBannerOptions() AwsEndpoint_ClientLoginBannerOptionsPropertyOutputReference {
	var returns AwsEndpoint_ClientLoginBannerOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"clientLoginBannerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) ClientLoginBannerOptionsInput() *AwsEndpoint_ClientLoginBannerOptionsProperty {
	var returns *AwsEndpoint_ClientLoginBannerOptionsProperty
	_jsii_.Get(
		j,
		"clientLoginBannerOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) ClientRouteEnforcementOptions() AwsEndpoint_ClientRouteEnforcementOptionsPropertyOutputReference {
	var returns AwsEndpoint_ClientRouteEnforcementOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"clientRouteEnforcementOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) ClientRouteEnforcementOptionsInput() *AwsEndpoint_ClientRouteEnforcementOptionsProperty {
	var returns *AwsEndpoint_ClientRouteEnforcementOptionsProperty
	_jsii_.Get(
		j,
		"clientRouteEnforcementOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) ConnectionLogOptions() AwsEndpoint_ConnectionLogOptionsPropertyOutputReference {
	var returns AwsEndpoint_ConnectionLogOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"connectionLogOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) ConnectionLogOptionsInput() *AwsEndpoint_ConnectionLogOptionsProperty {
	var returns *AwsEndpoint_ConnectionLogOptionsProperty
	_jsii_.Get(
		j,
		"connectionLogOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) DisconnectOnSessionTimeout() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disconnectOnSessionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) DisconnectOnSessionTimeoutInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disconnectOnSessionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) DnsServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) EndpointIpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIpAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) EndpointIpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIpAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) SelfServicePortal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServicePortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) SelfServicePortalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServicePortalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) SelfServicePortalUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfServicePortalUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) ServerCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) ServerCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) SessionTimeoutHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) SessionTimeoutHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionTimeoutHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) SplitTunnel() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splitTunnel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) SplitTunnelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splitTunnelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) TrafficIpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficIpAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) TrafficIpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trafficIpAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) TransitGatewayConfiguration() AwsEndpoint_TransitGatewayConfigurationPropertyOutputReference {
	var returns AwsEndpoint_TransitGatewayConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"transitGatewayConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) TransitGatewayConfigurationInput() *AwsEndpoint_TransitGatewayConfigurationProperty {
	var returns *AwsEndpoint_TransitGatewayConfigurationProperty
	_jsii_.Get(
		j,
		"transitGatewayConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) TransportProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) TransportProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transportProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) VpcIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) VpnPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vpnPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint) VpnPortInput() *float64 {
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
func NewAwsEndpoint(scope constructs.Construct, id *string, config *AwsEndpointConfig) AwsEndpoint {
	_init_.Initialize()

	if err := validateNewAwsEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEndpoint{}

	_jsii_.Create(
		"@cdktn/aws-vpn-client.AwsEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint aws_ec2_client_vpn_endpoint} Resource.
// Experimental.
func NewAwsEndpoint_Override(a AwsEndpoint, scope constructs.Construct, id *string, config *AwsEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpn-client.AwsEndpoint",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetClientCidrBlock(val *string) {
	if err := j.validateSetClientCidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCidrBlock",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetDisconnectOnSessionTimeout(val interface{}) {
	if err := j.validateSetDisconnectOnSessionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disconnectOnSessionTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetDnsServers(val *[]*string) {
	if err := j.validateSetDnsServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServers",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetEndpointIpAddressType(val *string) {
	if err := j.validateSetEndpointIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointIpAddressType",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetSelfServicePortal(val *string) {
	if err := j.validateSetSelfServicePortalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfServicePortal",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetServerCertificateArn(val *string) {
	if err := j.validateSetServerCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCertificateArn",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetSessionTimeoutHours(val *float64) {
	if err := j.validateSetSessionTimeoutHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionTimeoutHours",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetSplitTunnel(val interface{}) {
	if err := j.validateSetSplitTunnelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"splitTunnel",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetTrafficIpAddressType(val *string) {
	if err := j.validateSetTrafficIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trafficIpAddressType",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetTransportProtocol(val *string) {
	if err := j.validateSetTransportProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transportProtocol",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetVpcId(val *string) {
	if err := j.validateSetVpcIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcId",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint)SetVpnPort(val *float64) {
	if err := j.validateSetVpnPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnPort",
		val,
	)
}

// Generates CDKTN code for importing a AwsEndpoint resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-client.AwsEndpoint",
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
func AwsEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-client.AwsEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-client.AwsEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-vpn-client.AwsEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-vpn-client.AwsEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsEndpoint) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsEndpoint) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsEndpoint) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsEndpoint) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsEndpoint) PutAuthenticationOptions(value interface{}) {
	if err := a.validatePutAuthenticationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthenticationOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpoint) PutClientConnectOptions(value *AwsEndpoint_ClientConnectOptionsProperty) {
	if err := a.validatePutClientConnectOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putClientConnectOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpoint) PutClientLoginBannerOptions(value *AwsEndpoint_ClientLoginBannerOptionsProperty) {
	if err := a.validatePutClientLoginBannerOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putClientLoginBannerOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpoint) PutClientRouteEnforcementOptions(value *AwsEndpoint_ClientRouteEnforcementOptionsProperty) {
	if err := a.validatePutClientRouteEnforcementOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putClientRouteEnforcementOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpoint) PutConnectionLogOptions(value *AwsEndpoint_ConnectionLogOptionsProperty) {
	if err := a.validatePutConnectionLogOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectionLogOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpoint) PutTransitGatewayConfiguration(value *AwsEndpoint_TransitGatewayConfigurationProperty) {
	if err := a.validatePutTransitGatewayConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTransitGatewayConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpoint) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetClientCidrBlock() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCidrBlock",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetClientConnectOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetClientConnectOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetClientLoginBannerOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetClientLoginBannerOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetClientRouteEnforcementOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetClientRouteEnforcementOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetDisconnectOnSessionTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetDisconnectOnSessionTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetDnsServers() {
	_jsii_.InvokeVoid(
		a,
		"resetDnsServers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetEndpointIpAddressType() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointIpAddressType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetSelfServicePortal() {
	_jsii_.InvokeVoid(
		a,
		"resetSelfServicePortal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetSessionTimeoutHours() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionTimeoutHours",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetSplitTunnel() {
	_jsii_.InvokeVoid(
		a,
		"resetSplitTunnel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetTrafficIpAddressType() {
	_jsii_.InvokeVoid(
		a,
		"resetTrafficIpAddressType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetTransitGatewayConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetTransitGatewayConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetTransportProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetTransportProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetVpcId() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) ResetVpnPort() {
	_jsii_.InvokeVoid(
		a,
		"resetVpnPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

