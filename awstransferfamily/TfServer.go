package awstransferfamily

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstransferfamily/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awstransferfamily/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server aws_transfer_server}.
// Experimental.
type TfServer interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Certificate() *string
	// Experimental.
	SetCertificate(val *string)
	// Experimental.
	CertificateInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
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
	DirectoryId() *string
	// Experimental.
	SetDirectoryId(val *string)
	// Experimental.
	DirectoryIdInput() *string
	// Experimental.
	Domain() *string
	// Experimental.
	SetDomain(val *string)
	// Experimental.
	DomainInput() *string
	// Experimental.
	Endpoint() *string
	// Experimental.
	EndpointDetails() TfServer_EndpointDetailsPropertyOutputReference
	// Experimental.
	EndpointDetailsInput() *TfServer_EndpointDetailsProperty
	// Experimental.
	EndpointType() *string
	// Experimental.
	SetEndpointType(val *string)
	// Experimental.
	EndpointTypeInput() *string
	// Experimental.
	ForceDestroy() interface{}
	// Experimental.
	SetForceDestroy(val interface{})
	// Experimental.
	ForceDestroyInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Function() *string
	// Experimental.
	SetFunction(val *string)
	// Experimental.
	FunctionInput() *string
	// Experimental.
	HostKey() *string
	// Experimental.
	SetHostKey(val *string)
	// Experimental.
	HostKeyFingerprint() *string
	// Experimental.
	HostKeyInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdentityProviderType() *string
	// Experimental.
	SetIdentityProviderType(val *string)
	// Experimental.
	IdentityProviderTypeInput() *string
	// Experimental.
	IdInput() *string
	// Experimental.
	InvocationRole() *string
	// Experimental.
	SetInvocationRole(val *string)
	// Experimental.
	InvocationRoleInput() *string
	// Experimental.
	IpAddressType() *string
	// Experimental.
	SetIpAddressType(val *string)
	// Experimental.
	IpAddressTypeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LoggingRole() *string
	// Experimental.
	SetLoggingRole(val *string)
	// Experimental.
	LoggingRoleInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PostAuthenticationLoginBanner() *string
	// Experimental.
	SetPostAuthenticationLoginBanner(val *string)
	// Experimental.
	PostAuthenticationLoginBannerInput() *string
	// Experimental.
	PreAuthenticationLoginBanner() *string
	// Experimental.
	SetPreAuthenticationLoginBanner(val *string)
	// Experimental.
	PreAuthenticationLoginBannerInput() *string
	// Experimental.
	ProtocolDetails() TfServer_ProtocolDetailsPropertyOutputReference
	// Experimental.
	ProtocolDetailsInput() *TfServer_ProtocolDetailsProperty
	// Experimental.
	Protocols() *[]*string
	// Experimental.
	SetProtocols(val *[]*string)
	// Experimental.
	ProtocolsInput() *[]*string
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
	S3StorageOptions() TfServer_S3StorageOptionsPropertyOutputReference
	// Experimental.
	S3StorageOptionsInput() *TfServer_S3StorageOptionsProperty
	// Experimental.
	SecurityPolicyName() *string
	// Experimental.
	SetSecurityPolicyName(val *string)
	// Experimental.
	SecurityPolicyNameInput() *string
	// Experimental.
	SftpAuthenticationMethods() *string
	// Experimental.
	SetSftpAuthenticationMethods(val *string)
	// Experimental.
	SftpAuthenticationMethodsInput() *string
	// Experimental.
	StructuredLogDestinations() *[]*string
	// Experimental.
	SetStructuredLogDestinations(val *[]*string)
	// Experimental.
	StructuredLogDestinationsInput() *[]*string
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
	Url() *string
	// Experimental.
	SetUrl(val *string)
	// Experimental.
	UrlInput() *string
	// Experimental.
	WorkflowDetails() TfServer_WorkflowDetailsPropertyOutputReference
	// Experimental.
	WorkflowDetailsInput() *TfServer_WorkflowDetailsProperty
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
	PutEndpointDetails(value *TfServer_EndpointDetailsProperty)
	// Experimental.
	PutProtocolDetails(value *TfServer_ProtocolDetailsProperty)
	// Experimental.
	PutS3StorageOptions(value *TfServer_S3StorageOptionsProperty)
	// Experimental.
	PutWorkflowDetails(value *TfServer_WorkflowDetailsProperty)
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
	ResetCertificate()
	// Experimental.
	ResetDirectoryId()
	// Experimental.
	ResetDomain()
	// Experimental.
	ResetEndpointDetails()
	// Experimental.
	ResetEndpointType()
	// Experimental.
	ResetForceDestroy()
	// Experimental.
	ResetFunction()
	// Experimental.
	ResetHostKey()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIdentityProviderType()
	// Experimental.
	ResetInvocationRole()
	// Experimental.
	ResetIpAddressType()
	// Experimental.
	ResetLoggingRole()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPostAuthenticationLoginBanner()
	// Experimental.
	ResetPreAuthenticationLoginBanner()
	// Experimental.
	ResetProtocolDetails()
	// Experimental.
	ResetProtocols()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetS3StorageOptions()
	// Experimental.
	ResetSecurityPolicyName()
	// Experimental.
	ResetSftpAuthenticationMethods()
	// Experimental.
	ResetStructuredLogDestinations()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetUrl()
	// Experimental.
	ResetWorkflowDetails()
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

// The jsii proxy struct for TfServer
type jsiiProxy_TfServer struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfServer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) EndpointDetails() TfServer_EndpointDetailsPropertyOutputReference {
	var returns TfServer_EndpointDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"endpointDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) EndpointDetailsInput() *TfServer_EndpointDetailsProperty {
	var returns *TfServer_EndpointDetailsProperty
	_jsii_.Get(
		j,
		"endpointDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) Function() *string {
	var returns *string
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) FunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) HostKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) HostKeyFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKeyFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) HostKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) IdentityProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) IdentityProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) InvocationRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invocationRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) InvocationRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invocationRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) LoggingRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) LoggingRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) PostAuthenticationLoginBanner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAuthenticationLoginBanner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) PostAuthenticationLoginBannerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAuthenticationLoginBannerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) PreAuthenticationLoginBanner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preAuthenticationLoginBanner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) PreAuthenticationLoginBannerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preAuthenticationLoginBannerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) ProtocolDetails() TfServer_ProtocolDetailsPropertyOutputReference {
	var returns TfServer_ProtocolDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"protocolDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) ProtocolDetailsInput() *TfServer_ProtocolDetailsProperty {
	var returns *TfServer_ProtocolDetailsProperty
	_jsii_.Get(
		j,
		"protocolDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) S3StorageOptions() TfServer_S3StorageOptionsPropertyOutputReference {
	var returns TfServer_S3StorageOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"s3StorageOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) S3StorageOptionsInput() *TfServer_S3StorageOptionsProperty {
	var returns *TfServer_S3StorageOptionsProperty
	_jsii_.Get(
		j,
		"s3StorageOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) SecurityPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) SecurityPolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) SftpAuthenticationMethods() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sftpAuthenticationMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) SftpAuthenticationMethodsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sftpAuthenticationMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) StructuredLogDestinations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"structuredLogDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) StructuredLogDestinationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"structuredLogDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) WorkflowDetails() TfServer_WorkflowDetailsPropertyOutputReference {
	var returns TfServer_WorkflowDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"workflowDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer) WorkflowDetailsInput() *TfServer_WorkflowDetailsProperty {
	var returns *TfServer_WorkflowDetailsProperty
	_jsii_.Get(
		j,
		"workflowDetailsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server aws_transfer_server} Resource.
// Experimental.
func NewTfServer(scope constructs.Construct, id *string, config *TfServerConfig) TfServer {
	_init_.Initialize()

	if err := validateNewTfServerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfServer{}

	_jsii_.Create(
		"@cdktn/aws-transfer-family.TfServer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server aws_transfer_server} Resource.
// Experimental.
func NewTfServer_Override(t TfServer, scope constructs.Construct, id *string, config *TfServerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transfer-family.TfServer",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfServer)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetDirectoryId(val *string) {
	if err := j.validateSetDirectoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetEndpointType(val *string) {
	if err := j.validateSetEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetFunction(val *string) {
	if err := j.validateSetFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"function",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetHostKey(val *string) {
	if err := j.validateSetHostKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostKey",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetIdentityProviderType(val *string) {
	if err := j.validateSetIdentityProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProviderType",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetInvocationRole(val *string) {
	if err := j.validateSetInvocationRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invocationRole",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetLoggingRole(val *string) {
	if err := j.validateSetLoggingRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingRole",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetPostAuthenticationLoginBanner(val *string) {
	if err := j.validateSetPostAuthenticationLoginBannerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postAuthenticationLoginBanner",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetPreAuthenticationLoginBanner(val *string) {
	if err := j.validateSetPreAuthenticationLoginBannerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preAuthenticationLoginBanner",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetSecurityPolicyName(val *string) {
	if err := j.validateSetSecurityPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityPolicyName",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetSftpAuthenticationMethods(val *string) {
	if err := j.validateSetSftpAuthenticationMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sftpAuthenticationMethods",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetStructuredLogDestinations(val *[]*string) {
	if err := j.validateSetStructuredLogDestinationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"structuredLogDestinations",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfServer)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

// Generates CDKTN code for importing a TfServer resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfServer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfServer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-transfer-family.TfServer",
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
func TfServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfServer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-transfer-family.TfServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfServer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfServer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-transfer-family.TfServer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfServer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfServer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-transfer-family.TfServer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfServer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-transfer-family.TfServer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfServer) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfServer) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfServer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfServer) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfServer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfServer) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfServer) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfServer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfServer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfServer) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfServer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfServer) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfServer) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfServer) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfServer) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfServer) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfServer) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfServer) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfServer) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfServer) PutEndpointDetails(value *TfServer_EndpointDetailsProperty) {
	if err := t.validatePutEndpointDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEndpointDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfServer) PutProtocolDetails(value *TfServer_ProtocolDetailsProperty) {
	if err := t.validatePutProtocolDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProtocolDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfServer) PutS3StorageOptions(value *TfServer_S3StorageOptionsProperty) {
	if err := t.validatePutS3StorageOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3StorageOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfServer) PutWorkflowDetails(value *TfServer_WorkflowDetailsProperty) {
	if err := t.validatePutWorkflowDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWorkflowDetails",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfServer) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfServer) ResetCertificate() {
	_jsii_.InvokeVoid(
		t,
		"resetCertificate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetDirectoryId() {
	_jsii_.InvokeVoid(
		t,
		"resetDirectoryId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetDomain() {
	_jsii_.InvokeVoid(
		t,
		"resetDomain",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetEndpointDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetEndpointDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetEndpointType() {
	_jsii_.InvokeVoid(
		t,
		"resetEndpointType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		t,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetFunction() {
	_jsii_.InvokeVoid(
		t,
		"resetFunction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetHostKey() {
	_jsii_.InvokeVoid(
		t,
		"resetHostKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetIdentityProviderType() {
	_jsii_.InvokeVoid(
		t,
		"resetIdentityProviderType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetInvocationRole() {
	_jsii_.InvokeVoid(
		t,
		"resetInvocationRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		t,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetLoggingRole() {
	_jsii_.InvokeVoid(
		t,
		"resetLoggingRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetPostAuthenticationLoginBanner() {
	_jsii_.InvokeVoid(
		t,
		"resetPostAuthenticationLoginBanner",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetPreAuthenticationLoginBanner() {
	_jsii_.InvokeVoid(
		t,
		"resetPreAuthenticationLoginBanner",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetProtocolDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetProtocolDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetProtocols() {
	_jsii_.InvokeVoid(
		t,
		"resetProtocols",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetS3StorageOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetS3StorageOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetSecurityPolicyName() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityPolicyName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetSftpAuthenticationMethods() {
	_jsii_.InvokeVoid(
		t,
		"resetSftpAuthenticationMethods",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetStructuredLogDestinations() {
	_jsii_.InvokeVoid(
		t,
		"resetStructuredLogDestinations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) ResetWorkflowDetails() {
	_jsii_.InvokeVoid(
		t,
		"resetWorkflowDetails",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfServer) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfServer) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfServer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfServer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfServer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfServer) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

