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
type AwsTransferServer interface {
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
	EndpointDetails() AwsTransferServer_EndpointDetailsPropertyOutputReference
	// Experimental.
	EndpointDetailsInput() *AwsTransferServer_EndpointDetailsProperty
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
	ProtocolDetails() AwsTransferServer_ProtocolDetailsPropertyOutputReference
	// Experimental.
	ProtocolDetailsInput() *AwsTransferServer_ProtocolDetailsProperty
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
	S3StorageOptions() AwsTransferServer_S3StorageOptionsPropertyOutputReference
	// Experimental.
	S3StorageOptionsInput() *AwsTransferServer_S3StorageOptionsProperty
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
	WorkflowDetails() AwsTransferServer_WorkflowDetailsPropertyOutputReference
	// Experimental.
	WorkflowDetailsInput() *AwsTransferServer_WorkflowDetailsProperty
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
	PutEndpointDetails(value *AwsTransferServer_EndpointDetailsProperty)
	// Experimental.
	PutProtocolDetails(value *AwsTransferServer_ProtocolDetailsProperty)
	// Experimental.
	PutS3StorageOptions(value *AwsTransferServer_S3StorageOptionsProperty)
	// Experimental.
	PutWorkflowDetails(value *AwsTransferServer_WorkflowDetailsProperty)
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

// The jsii proxy struct for AwsTransferServer
type jsiiProxy_AwsTransferServer struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsTransferServer) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) DirectoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) DirectoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) EndpointDetails() AwsTransferServer_EndpointDetailsPropertyOutputReference {
	var returns AwsTransferServer_EndpointDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"endpointDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) EndpointDetailsInput() *AwsTransferServer_EndpointDetailsProperty {
	var returns *AwsTransferServer_EndpointDetailsProperty
	_jsii_.Get(
		j,
		"endpointDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) Function() *string {
	var returns *string
	_jsii_.Get(
		j,
		"function",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) FunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) HostKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) HostKeyFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKeyFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) HostKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) IdentityProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) IdentityProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) InvocationRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invocationRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) InvocationRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invocationRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) LoggingRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) LoggingRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) PostAuthenticationLoginBanner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAuthenticationLoginBanner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) PostAuthenticationLoginBannerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postAuthenticationLoginBannerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) PreAuthenticationLoginBanner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preAuthenticationLoginBanner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) PreAuthenticationLoginBannerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preAuthenticationLoginBannerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) ProtocolDetails() AwsTransferServer_ProtocolDetailsPropertyOutputReference {
	var returns AwsTransferServer_ProtocolDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"protocolDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) ProtocolDetailsInput() *AwsTransferServer_ProtocolDetailsProperty {
	var returns *AwsTransferServer_ProtocolDetailsProperty
	_jsii_.Get(
		j,
		"protocolDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) S3StorageOptions() AwsTransferServer_S3StorageOptionsPropertyOutputReference {
	var returns AwsTransferServer_S3StorageOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"s3StorageOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) S3StorageOptionsInput() *AwsTransferServer_S3StorageOptionsProperty {
	var returns *AwsTransferServer_S3StorageOptionsProperty
	_jsii_.Get(
		j,
		"s3StorageOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) SecurityPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) SecurityPolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) SftpAuthenticationMethods() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sftpAuthenticationMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) SftpAuthenticationMethodsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sftpAuthenticationMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) StructuredLogDestinations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"structuredLogDestinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) StructuredLogDestinationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"structuredLogDestinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) WorkflowDetails() AwsTransferServer_WorkflowDetailsPropertyOutputReference {
	var returns AwsTransferServer_WorkflowDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"workflowDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer) WorkflowDetailsInput() *AwsTransferServer_WorkflowDetailsProperty {
	var returns *AwsTransferServer_WorkflowDetailsProperty
	_jsii_.Get(
		j,
		"workflowDetailsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server aws_transfer_server} Resource.
// Experimental.
func NewAwsTransferServer(scope constructs.Construct, id *string, config *AwsTransferServerConfig) AwsTransferServer {
	_init_.Initialize()

	if err := validateNewAwsTransferServerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTransferServer{}

	_jsii_.Create(
		"@cdktn/aws-transfer-family.AwsTransferServer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_server aws_transfer_server} Resource.
// Experimental.
func NewAwsTransferServer_Override(a AwsTransferServer, scope constructs.Construct, id *string, config *AwsTransferServerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transfer-family.AwsTransferServer",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetDirectoryId(val *string) {
	if err := j.validateSetDirectoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"directoryId",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetEndpointType(val *string) {
	if err := j.validateSetEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetFunction(val *string) {
	if err := j.validateSetFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"function",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetHostKey(val *string) {
	if err := j.validateSetHostKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostKey",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetIdentityProviderType(val *string) {
	if err := j.validateSetIdentityProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProviderType",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetInvocationRole(val *string) {
	if err := j.validateSetInvocationRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invocationRole",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetLoggingRole(val *string) {
	if err := j.validateSetLoggingRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingRole",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetPostAuthenticationLoginBanner(val *string) {
	if err := j.validateSetPostAuthenticationLoginBannerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postAuthenticationLoginBanner",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetPreAuthenticationLoginBanner(val *string) {
	if err := j.validateSetPreAuthenticationLoginBannerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preAuthenticationLoginBanner",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetSecurityPolicyName(val *string) {
	if err := j.validateSetSecurityPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityPolicyName",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetSftpAuthenticationMethods(val *string) {
	if err := j.validateSetSftpAuthenticationMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sftpAuthenticationMethods",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetStructuredLogDestinations(val *[]*string) {
	if err := j.validateSetStructuredLogDestinationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"structuredLogDestinations",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

// Generates CDKTN code for importing a AwsTransferServer resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsTransferServer_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsTransferServer_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-transfer-family.AwsTransferServer",
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
func AwsTransferServer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsTransferServer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-transfer-family.AwsTransferServer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsTransferServer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsTransferServer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-transfer-family.AwsTransferServer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsTransferServer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsTransferServer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-transfer-family.AwsTransferServer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsTransferServer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-transfer-family.AwsTransferServer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsTransferServer) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsTransferServer) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsTransferServer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTransferServer) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTransferServer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTransferServer) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTransferServer) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTransferServer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTransferServer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTransferServer) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTransferServer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTransferServer) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransferServer) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsTransferServer) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTransferServer) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsTransferServer) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsTransferServer) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsTransferServer) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsTransferServer) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsTransferServer) PutEndpointDetails(value *AwsTransferServer_EndpointDetailsProperty) {
	if err := a.validatePutEndpointDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEndpointDetails",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransferServer) PutProtocolDetails(value *AwsTransferServer_ProtocolDetailsProperty) {
	if err := a.validatePutProtocolDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProtocolDetails",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransferServer) PutS3StorageOptions(value *AwsTransferServer_S3StorageOptionsProperty) {
	if err := a.validatePutS3StorageOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3StorageOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransferServer) PutWorkflowDetails(value *AwsTransferServer_WorkflowDetailsProperty) {
	if err := a.validatePutWorkflowDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkflowDetails",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTransferServer) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetDirectoryId() {
	_jsii_.InvokeVoid(
		a,
		"resetDirectoryId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetEndpointDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetEndpointType() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		a,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetFunction() {
	_jsii_.InvokeVoid(
		a,
		"resetFunction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetHostKey() {
	_jsii_.InvokeVoid(
		a,
		"resetHostKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetIdentityProviderType() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityProviderType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetInvocationRole() {
	_jsii_.InvokeVoid(
		a,
		"resetInvocationRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		a,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetLoggingRole() {
	_jsii_.InvokeVoid(
		a,
		"resetLoggingRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetPostAuthenticationLoginBanner() {
	_jsii_.InvokeVoid(
		a,
		"resetPostAuthenticationLoginBanner",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetPreAuthenticationLoginBanner() {
	_jsii_.InvokeVoid(
		a,
		"resetPreAuthenticationLoginBanner",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetProtocolDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocolDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetProtocols() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocols",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetS3StorageOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetS3StorageOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetSecurityPolicyName() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityPolicyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetSftpAuthenticationMethods() {
	_jsii_.InvokeVoid(
		a,
		"resetSftpAuthenticationMethods",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetStructuredLogDestinations() {
	_jsii_.InvokeVoid(
		a,
		"resetStructuredLogDestinations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) ResetWorkflowDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkflowDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransferServer) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransferServer) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransferServer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransferServer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransferServer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransferServer) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

