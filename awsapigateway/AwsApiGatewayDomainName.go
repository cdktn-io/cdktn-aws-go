package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapigateway/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsapigateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name aws_api_gateway_domain_name}.
// Experimental.
type AwsApiGatewayDomainName interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CertificateArn() *string
	// Experimental.
	SetCertificateArn(val *string)
	// Experimental.
	CertificateArnInput() *string
	// Experimental.
	CertificateBody() *string
	// Experimental.
	SetCertificateBody(val *string)
	// Experimental.
	CertificateBodyInput() *string
	// Experimental.
	CertificateChain() *string
	// Experimental.
	SetCertificateChain(val *string)
	// Experimental.
	CertificateChainInput() *string
	// Experimental.
	CertificateName() *string
	// Experimental.
	SetCertificateName(val *string)
	// Experimental.
	CertificateNameInput() *string
	// Experimental.
	CertificatePrivateKey() *string
	// Experimental.
	SetCertificatePrivateKey(val *string)
	// Experimental.
	CertificatePrivateKeyInput() *string
	// Experimental.
	CertificateUploadDate() *string
	// Experimental.
	CloudfrontDomainName() *string
	// Experimental.
	CloudfrontZoneId() *string
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
	DomainName() *string
	// Experimental.
	SetDomainName(val *string)
	// Experimental.
	DomainNameId() *string
	// Experimental.
	DomainNameInput() *string
	// Experimental.
	EndpointAccessMode() *string
	// Experimental.
	SetEndpointAccessMode(val *string)
	// Experimental.
	EndpointAccessModeInput() *string
	// Experimental.
	EndpointConfiguration() AwsApiGatewayDomainName_EndpointConfigurationPropertyOutputReference
	// Experimental.
	EndpointConfigurationInput() *AwsApiGatewayDomainName_EndpointConfigurationProperty
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
	MutualTlsAuthentication() AwsApiGatewayDomainName_MutualTlsAuthenticationPropertyOutputReference
	// Experimental.
	MutualTlsAuthenticationInput() *AwsApiGatewayDomainName_MutualTlsAuthenticationProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OwnershipVerificationCertificateArn() *string
	// Experimental.
	SetOwnershipVerificationCertificateArn(val *string)
	// Experimental.
	OwnershipVerificationCertificateArnInput() *string
	// Experimental.
	Policy() *string
	// Experimental.
	SetPolicy(val *string)
	// Experimental.
	PolicyInput() *string
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
	RegionalCertificateArn() *string
	// Experimental.
	SetRegionalCertificateArn(val *string)
	// Experimental.
	RegionalCertificateArnInput() *string
	// Experimental.
	RegionalCertificateName() *string
	// Experimental.
	SetRegionalCertificateName(val *string)
	// Experimental.
	RegionalCertificateNameInput() *string
	// Experimental.
	RegionalDomainName() *string
	// Experimental.
	RegionalZoneId() *string
	// Experimental.
	RegionInput() *string
	// Experimental.
	RoutingMode() *string
	// Experimental.
	SetRoutingMode(val *string)
	// Experimental.
	RoutingModeInput() *string
	// Experimental.
	SecurityPolicy() *string
	// Experimental.
	SetSecurityPolicy(val *string)
	// Experimental.
	SecurityPolicyInput() *string
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
	Timeouts() AwsApiGatewayDomainName_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
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
	PutEndpointConfiguration(value *AwsApiGatewayDomainName_EndpointConfigurationProperty)
	// Experimental.
	PutMutualTlsAuthentication(value *AwsApiGatewayDomainName_MutualTlsAuthenticationProperty)
	// Experimental.
	PutTimeouts(value *AwsApiGatewayDomainName_TimeoutsProperty)
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
	ResetCertificateArn()
	// Experimental.
	ResetCertificateBody()
	// Experimental.
	ResetCertificateChain()
	// Experimental.
	ResetCertificateName()
	// Experimental.
	ResetCertificatePrivateKey()
	// Experimental.
	ResetEndpointAccessMode()
	// Experimental.
	ResetEndpointConfiguration()
	// Experimental.
	ResetId()
	// Experimental.
	ResetMutualTlsAuthentication()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetOwnershipVerificationCertificateArn()
	// Experimental.
	ResetPolicy()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRegionalCertificateArn()
	// Experimental.
	ResetRegionalCertificateName()
	// Experimental.
	ResetRoutingMode()
	// Experimental.
	ResetSecurityPolicy()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
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

// The jsii proxy struct for AwsApiGatewayDomainName
type jsiiProxy_AwsApiGatewayDomainName struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsApiGatewayDomainName) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) CertificateBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) CertificateBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) CertificateChain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) CertificateChainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateChainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) CertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) CertificateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) CertificatePrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) CertificatePrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePrivateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) CertificateUploadDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateUploadDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) CloudfrontDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudfrontDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) CloudfrontZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudfrontZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) DomainNameId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) EndpointAccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointAccessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) EndpointAccessModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointAccessModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) EndpointConfiguration() AwsApiGatewayDomainName_EndpointConfigurationPropertyOutputReference {
	var returns AwsApiGatewayDomainName_EndpointConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"endpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) EndpointConfigurationInput() *AwsApiGatewayDomainName_EndpointConfigurationProperty {
	var returns *AwsApiGatewayDomainName_EndpointConfigurationProperty
	_jsii_.Get(
		j,
		"endpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) MutualTlsAuthentication() AwsApiGatewayDomainName_MutualTlsAuthenticationPropertyOutputReference {
	var returns AwsApiGatewayDomainName_MutualTlsAuthenticationPropertyOutputReference
	_jsii_.Get(
		j,
		"mutualTlsAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) MutualTlsAuthenticationInput() *AwsApiGatewayDomainName_MutualTlsAuthenticationProperty {
	var returns *AwsApiGatewayDomainName_MutualTlsAuthenticationProperty
	_jsii_.Get(
		j,
		"mutualTlsAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) OwnershipVerificationCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownershipVerificationCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) OwnershipVerificationCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownershipVerificationCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) RegionalCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) RegionalCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) RegionalCertificateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalCertificateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) RegionalCertificateNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalCertificateNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) RegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) RegionalZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionalZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) RoutingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) RoutingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) SecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) SecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) Timeouts() AwsApiGatewayDomainName_TimeoutsPropertyOutputReference {
	var returns AwsApiGatewayDomainName_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayDomainName) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name aws_api_gateway_domain_name} Resource.
// Experimental.
func NewAwsApiGatewayDomainName(scope constructs.Construct, id *string, config *AwsApiGatewayDomainNameConfig) AwsApiGatewayDomainName {
	_init_.Initialize()

	if err := validateNewAwsApiGatewayDomainNameParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApiGatewayDomainName{}

	_jsii_.Create(
		"@cdktn/aws-api-gateway.AwsApiGatewayDomainName",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_domain_name aws_api_gateway_domain_name} Resource.
// Experimental.
func NewAwsApiGatewayDomainName_Override(a AwsApiGatewayDomainName, scope constructs.Construct, id *string, config *AwsApiGatewayDomainNameConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-api-gateway.AwsApiGatewayDomainName",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetCertificateArn(val *string) {
	if err := j.validateSetCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetCertificateBody(val *string) {
	if err := j.validateSetCertificateBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateBody",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetCertificateChain(val *string) {
	if err := j.validateSetCertificateChainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateChain",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetCertificateName(val *string) {
	if err := j.validateSetCertificateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateName",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetCertificatePrivateKey(val *string) {
	if err := j.validateSetCertificatePrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificatePrivateKey",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetEndpointAccessMode(val *string) {
	if err := j.validateSetEndpointAccessModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointAccessMode",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetOwnershipVerificationCertificateArn(val *string) {
	if err := j.validateSetOwnershipVerificationCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ownershipVerificationCertificateArn",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetPolicy(val *string) {
	if err := j.validateSetPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetRegionalCertificateArn(val *string) {
	if err := j.validateSetRegionalCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionalCertificateArn",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetRegionalCertificateName(val *string) {
	if err := j.validateSetRegionalCertificateNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionalCertificateName",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetRoutingMode(val *string) {
	if err := j.validateSetRoutingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingMode",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetSecurityPolicy(val *string) {
	if err := j.validateSetSecurityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayDomainName)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a AwsApiGatewayDomainName resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsApiGatewayDomainName_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsApiGatewayDomainName_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-api-gateway.AwsApiGatewayDomainName",
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
func AwsApiGatewayDomainName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsApiGatewayDomainName_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-api-gateway.AwsApiGatewayDomainName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsApiGatewayDomainName_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsApiGatewayDomainName_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-api-gateway.AwsApiGatewayDomainName",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsApiGatewayDomainName_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsApiGatewayDomainName_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-api-gateway.AwsApiGatewayDomainName",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsApiGatewayDomainName_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-api-gateway.AwsApiGatewayDomainName",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsApiGatewayDomainName) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsApiGatewayDomainName) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApiGatewayDomainName) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsApiGatewayDomainName) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsApiGatewayDomainName) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsApiGatewayDomainName) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsApiGatewayDomainName) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsApiGatewayDomainName) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsApiGatewayDomainName) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsApiGatewayDomainName) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApiGatewayDomainName) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsApiGatewayDomainName) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) PutEndpointConfiguration(value *AwsApiGatewayDomainName_EndpointConfigurationProperty) {
	if err := a.validatePutEndpointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEndpointConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) PutMutualTlsAuthentication(value *AwsApiGatewayDomainName_MutualTlsAuthenticationProperty) {
	if err := a.validatePutMutualTlsAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMutualTlsAuthentication",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) PutTimeouts(value *AwsApiGatewayDomainName_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetCertificateBody() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateBody",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetCertificateChain() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateChain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetCertificateName() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetCertificatePrivateKey() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificatePrivateKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetEndpointAccessMode() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointAccessMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetEndpointConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetMutualTlsAuthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetMutualTlsAuthentication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetOwnershipVerificationCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetOwnershipVerificationCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetRegionalCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRegionalCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetRegionalCertificateName() {
	_jsii_.InvokeVoid(
		a,
		"resetRegionalCertificateName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetRoutingMode() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetSecurityPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayDomainName) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiGatewayDomainName) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiGatewayDomainName) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiGatewayDomainName) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

