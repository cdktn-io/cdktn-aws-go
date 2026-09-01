package awselb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselb/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awselb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener aws_alb_listener}.
// Experimental.
type AwsAlbListener interface {
	cdktn.TerraformResource
	// Experimental.
	AlpnPolicy() *string
	// Experimental.
	SetAlpnPolicy(val *string)
	// Experimental.
	AlpnPolicyInput() *string
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
	DefaultAction() AwsAlbListener_DefaultActionPropertyList
	// Experimental.
	DefaultActionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	LoadBalancerArn() *string
	// Experimental.
	SetLoadBalancerArn(val *string)
	// Experimental.
	LoadBalancerArnInput() *string
	// Experimental.
	MutualAuthentication() AwsAlbListener_MutualAuthenticationPropertyOutputReference
	// Experimental.
	MutualAuthenticationInput() *AwsAlbListener_MutualAuthenticationProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Port() *float64
	// Experimental.
	SetPort(val *float64)
	// Experimental.
	PortInput() *float64
	// Experimental.
	Protocol() *string
	// Experimental.
	SetProtocol(val *string)
	// Experimental.
	ProtocolInput() *string
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
	RoutingHttpRequestXAmznMtlsClientcertHeaderName() *string
	// Experimental.
	SetRoutingHttpRequestXAmznMtlsClientcertHeaderName(val *string)
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertHeaderNameInput() *string
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName() *string
	// Experimental.
	SetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName(val *string)
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertIssuerHeaderNameInput() *string
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertLeafHeaderName() *string
	// Experimental.
	SetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderName(val *string)
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertLeafHeaderNameInput() *string
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName() *string
	// Experimental.
	SetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName(val *string)
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderNameInput() *string
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName() *string
	// Experimental.
	SetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName(val *string)
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertSubjectHeaderNameInput() *string
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertValidityHeaderName() *string
	// Experimental.
	SetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderName(val *string)
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertValidityHeaderNameInput() *string
	// Experimental.
	RoutingHttpRequestXAmznTlsCipherSuiteHeaderName() *string
	// Experimental.
	SetRoutingHttpRequestXAmznTlsCipherSuiteHeaderName(val *string)
	// Experimental.
	RoutingHttpRequestXAmznTlsCipherSuiteHeaderNameInput() *string
	// Experimental.
	RoutingHttpRequestXAmznTlsVersionHeaderName() *string
	// Experimental.
	SetRoutingHttpRequestXAmznTlsVersionHeaderName(val *string)
	// Experimental.
	RoutingHttpRequestXAmznTlsVersionHeaderNameInput() *string
	// Experimental.
	RoutingHttpResponseAccessControlAllowCredentialsHeaderValue() *string
	// Experimental.
	SetRoutingHttpResponseAccessControlAllowCredentialsHeaderValue(val *string)
	// Experimental.
	RoutingHttpResponseAccessControlAllowCredentialsHeaderValueInput() *string
	// Experimental.
	RoutingHttpResponseAccessControlAllowHeadersHeaderValue() *string
	// Experimental.
	SetRoutingHttpResponseAccessControlAllowHeadersHeaderValue(val *string)
	// Experimental.
	RoutingHttpResponseAccessControlAllowHeadersHeaderValueInput() *string
	// Experimental.
	RoutingHttpResponseAccessControlAllowMethodsHeaderValue() *string
	// Experimental.
	SetRoutingHttpResponseAccessControlAllowMethodsHeaderValue(val *string)
	// Experimental.
	RoutingHttpResponseAccessControlAllowMethodsHeaderValueInput() *string
	// Experimental.
	RoutingHttpResponseAccessControlAllowOriginHeaderValue() *string
	// Experimental.
	SetRoutingHttpResponseAccessControlAllowOriginHeaderValue(val *string)
	// Experimental.
	RoutingHttpResponseAccessControlAllowOriginHeaderValueInput() *string
	// Experimental.
	RoutingHttpResponseAccessControlExposeHeadersHeaderValue() *string
	// Experimental.
	SetRoutingHttpResponseAccessControlExposeHeadersHeaderValue(val *string)
	// Experimental.
	RoutingHttpResponseAccessControlExposeHeadersHeaderValueInput() *string
	// Experimental.
	RoutingHttpResponseAccessControlMaxAgeHeaderValue() *string
	// Experimental.
	SetRoutingHttpResponseAccessControlMaxAgeHeaderValue(val *string)
	// Experimental.
	RoutingHttpResponseAccessControlMaxAgeHeaderValueInput() *string
	// Experimental.
	RoutingHttpResponseContentSecurityPolicyHeaderValue() *string
	// Experimental.
	SetRoutingHttpResponseContentSecurityPolicyHeaderValue(val *string)
	// Experimental.
	RoutingHttpResponseContentSecurityPolicyHeaderValueInput() *string
	// Experimental.
	RoutingHttpResponseServerEnabled() interface{}
	// Experimental.
	SetRoutingHttpResponseServerEnabled(val interface{})
	// Experimental.
	RoutingHttpResponseServerEnabledInput() interface{}
	// Experimental.
	RoutingHttpResponseStrictTransportSecurityHeaderValue() *string
	// Experimental.
	SetRoutingHttpResponseStrictTransportSecurityHeaderValue(val *string)
	// Experimental.
	RoutingHttpResponseStrictTransportSecurityHeaderValueInput() *string
	// Experimental.
	RoutingHttpResponseXContentTypeOptionsHeaderValue() *string
	// Experimental.
	SetRoutingHttpResponseXContentTypeOptionsHeaderValue(val *string)
	// Experimental.
	RoutingHttpResponseXContentTypeOptionsHeaderValueInput() *string
	// Experimental.
	RoutingHttpResponseXFrameOptionsHeaderValue() *string
	// Experimental.
	SetRoutingHttpResponseXFrameOptionsHeaderValue(val *string)
	// Experimental.
	RoutingHttpResponseXFrameOptionsHeaderValueInput() *string
	// Experimental.
	SslPolicy() *string
	// Experimental.
	SetSslPolicy(val *string)
	// Experimental.
	SslPolicyInput() *string
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
	TcpIdleTimeoutSeconds() *float64
	// Experimental.
	SetTcpIdleTimeoutSeconds(val *float64)
	// Experimental.
	TcpIdleTimeoutSecondsInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsAlbListener_TimeoutsPropertyOutputReference
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
	PutDefaultAction(value interface{})
	// Experimental.
	PutMutualAuthentication(value *AwsAlbListener_MutualAuthenticationProperty)
	// Experimental.
	PutTimeouts(value *AwsAlbListener_TimeoutsProperty)
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
	ResetAlpnPolicy()
	// Experimental.
	ResetCertificateArn()
	// Experimental.
	ResetId()
	// Experimental.
	ResetMutualAuthentication()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPort()
	// Experimental.
	ResetProtocol()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRoutingHttpRequestXAmznMtlsClientcertHeaderName()
	// Experimental.
	ResetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName()
	// Experimental.
	ResetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderName()
	// Experimental.
	ResetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName()
	// Experimental.
	ResetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName()
	// Experimental.
	ResetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderName()
	// Experimental.
	ResetRoutingHttpRequestXAmznTlsCipherSuiteHeaderName()
	// Experimental.
	ResetRoutingHttpRequestXAmznTlsVersionHeaderName()
	// Experimental.
	ResetRoutingHttpResponseAccessControlAllowCredentialsHeaderValue()
	// Experimental.
	ResetRoutingHttpResponseAccessControlAllowHeadersHeaderValue()
	// Experimental.
	ResetRoutingHttpResponseAccessControlAllowMethodsHeaderValue()
	// Experimental.
	ResetRoutingHttpResponseAccessControlAllowOriginHeaderValue()
	// Experimental.
	ResetRoutingHttpResponseAccessControlExposeHeadersHeaderValue()
	// Experimental.
	ResetRoutingHttpResponseAccessControlMaxAgeHeaderValue()
	// Experimental.
	ResetRoutingHttpResponseContentSecurityPolicyHeaderValue()
	// Experimental.
	ResetRoutingHttpResponseServerEnabled()
	// Experimental.
	ResetRoutingHttpResponseStrictTransportSecurityHeaderValue()
	// Experimental.
	ResetRoutingHttpResponseXContentTypeOptionsHeaderValue()
	// Experimental.
	ResetRoutingHttpResponseXFrameOptionsHeaderValue()
	// Experimental.
	ResetSslPolicy()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTcpIdleTimeoutSeconds()
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

// The jsii proxy struct for AwsAlbListener
type jsiiProxy_AwsAlbListener struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsAlbListener) AlpnPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alpnPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) AlpnPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alpnPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) DefaultAction() AwsAlbListener_DefaultActionPropertyList {
	var returns AwsAlbListener_DefaultActionPropertyList
	_jsii_.Get(
		j,
		"defaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) DefaultActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) LoadBalancerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) MutualAuthentication() AwsAlbListener_MutualAuthenticationPropertyOutputReference {
	var returns AwsAlbListener_MutualAuthenticationPropertyOutputReference
	_jsii_.Get(
		j,
		"mutualAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) MutualAuthenticationInput() *AwsAlbListener_MutualAuthenticationProperty {
	var returns *AwsAlbListener_MutualAuthenticationProperty
	_jsii_.Get(
		j,
		"mutualAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpRequestXAmznMtlsClientcertHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpRequestXAmznMtlsClientcertHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertIssuerHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpRequestXAmznMtlsClientcertIssuerHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertIssuerHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpRequestXAmznMtlsClientcertLeafHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertLeafHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpRequestXAmznMtlsClientcertLeafHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertLeafHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertSerialNumberHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertSubjectHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpRequestXAmznMtlsClientcertSubjectHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertSubjectHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpRequestXAmznMtlsClientcertValidityHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertValidityHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpRequestXAmznMtlsClientcertValidityHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertValidityHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpRequestXAmznTlsCipherSuiteHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznTlsCipherSuiteHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpRequestXAmznTlsCipherSuiteHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznTlsCipherSuiteHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpRequestXAmznTlsVersionHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznTlsVersionHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpRequestXAmznTlsVersionHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznTlsVersionHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseAccessControlAllowCredentialsHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowCredentialsHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseAccessControlAllowCredentialsHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowCredentialsHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseAccessControlAllowHeadersHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowHeadersHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseAccessControlAllowHeadersHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowHeadersHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseAccessControlAllowMethodsHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowMethodsHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseAccessControlAllowMethodsHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowMethodsHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseAccessControlAllowOriginHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowOriginHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseAccessControlAllowOriginHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowOriginHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseAccessControlExposeHeadersHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlExposeHeadersHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseAccessControlExposeHeadersHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlExposeHeadersHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseAccessControlMaxAgeHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlMaxAgeHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseAccessControlMaxAgeHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlMaxAgeHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseContentSecurityPolicyHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseContentSecurityPolicyHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseContentSecurityPolicyHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseContentSecurityPolicyHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseServerEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingHttpResponseServerEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseServerEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingHttpResponseServerEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseStrictTransportSecurityHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseStrictTransportSecurityHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseStrictTransportSecurityHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseStrictTransportSecurityHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseXContentTypeOptionsHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseXContentTypeOptionsHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseXContentTypeOptionsHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseXContentTypeOptionsHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseXFrameOptionsHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseXFrameOptionsHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) RoutingHttpResponseXFrameOptionsHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseXFrameOptionsHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) SslPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) SslPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) TcpIdleTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpIdleTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) TcpIdleTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpIdleTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) Timeouts() AwsAlbListener_TimeoutsPropertyOutputReference {
	var returns AwsAlbListener_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAlbListener) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener aws_alb_listener} Resource.
// Experimental.
func NewAwsAlbListener(scope constructs.Construct, id *string, config *AwsAlbListenerConfig) AwsAlbListener {
	_init_.Initialize()

	if err := validateNewAwsAlbListenerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAlbListener{}

	_jsii_.Create(
		"@cdktn/aws-elb.AwsAlbListener",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener aws_alb_listener} Resource.
// Experimental.
func NewAwsAlbListener_Override(a AwsAlbListener, scope constructs.Construct, id *string, config *AwsAlbListenerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.AwsAlbListener",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetAlpnPolicy(val *string) {
	if err := j.validateSetAlpnPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alpnPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetCertificateArn(val *string) {
	if err := j.validateSetCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetLoadBalancerArn(val *string) {
	if err := j.validateSetLoadBalancerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerArn",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpRequestXAmznMtlsClientcertHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertHeaderName",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertIssuerHeaderName",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertLeafHeaderName",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertSubjectHeaderName",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertValidityHeaderName",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpRequestXAmznTlsCipherSuiteHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznTlsCipherSuiteHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznTlsCipherSuiteHeaderName",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpRequestXAmznTlsVersionHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznTlsVersionHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznTlsVersionHeaderName",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpResponseAccessControlAllowCredentialsHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlAllowCredentialsHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlAllowCredentialsHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpResponseAccessControlAllowHeadersHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlAllowHeadersHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlAllowHeadersHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpResponseAccessControlAllowMethodsHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlAllowMethodsHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlAllowMethodsHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpResponseAccessControlAllowOriginHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlAllowOriginHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlAllowOriginHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpResponseAccessControlExposeHeadersHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlExposeHeadersHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlExposeHeadersHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpResponseAccessControlMaxAgeHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlMaxAgeHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlMaxAgeHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpResponseContentSecurityPolicyHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseContentSecurityPolicyHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseContentSecurityPolicyHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpResponseServerEnabled(val interface{}) {
	if err := j.validateSetRoutingHttpResponseServerEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseServerEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpResponseStrictTransportSecurityHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseStrictTransportSecurityHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseStrictTransportSecurityHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpResponseXContentTypeOptionsHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseXContentTypeOptionsHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseXContentTypeOptionsHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetRoutingHttpResponseXFrameOptionsHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseXFrameOptionsHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseXFrameOptionsHeaderValue",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetSslPolicy(val *string) {
	if err := j.validateSetSslPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsAlbListener)SetTcpIdleTimeoutSeconds(val *float64) {
	if err := j.validateSetTcpIdleTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpIdleTimeoutSeconds",
		val,
	)
}

// Generates CDKTN code for importing a AwsAlbListener resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsAlbListener_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsAlbListener_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.AwsAlbListener",
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
func AwsAlbListener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAlbListener_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.AwsAlbListener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAlbListener_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAlbListener_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.AwsAlbListener",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAlbListener_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAlbListener_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.AwsAlbListener",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsAlbListener_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-elb.AwsAlbListener",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsAlbListener) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsAlbListener) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsAlbListener) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAlbListener) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAlbListener) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAlbListener) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAlbListener) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAlbListener) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAlbListener) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAlbListener) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAlbListener) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAlbListener) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlbListener) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsAlbListener) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAlbListener) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsAlbListener) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAlbListener) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsAlbListener) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAlbListener) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsAlbListener) PutDefaultAction(value interface{}) {
	if err := a.validatePutDefaultActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDefaultAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlbListener) PutMutualAuthentication(value *AwsAlbListener_MutualAuthenticationProperty) {
	if err := a.validatePutMutualAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMutualAuthentication",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlbListener) PutTimeouts(value *AwsAlbListener_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAlbListener) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetAlpnPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetAlpnPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetMutualAuthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetMutualAuthentication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetProtocol() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocol",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpRequestXAmznMtlsClientcertHeaderName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpRequestXAmznMtlsClientcertHeaderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpRequestXAmznTlsCipherSuiteHeaderName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpRequestXAmznTlsCipherSuiteHeaderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpRequestXAmznTlsVersionHeaderName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpRequestXAmznTlsVersionHeaderName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpResponseAccessControlAllowCredentialsHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseAccessControlAllowCredentialsHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpResponseAccessControlAllowHeadersHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseAccessControlAllowHeadersHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpResponseAccessControlAllowMethodsHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseAccessControlAllowMethodsHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpResponseAccessControlAllowOriginHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseAccessControlAllowOriginHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpResponseAccessControlExposeHeadersHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseAccessControlExposeHeadersHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpResponseAccessControlMaxAgeHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseAccessControlMaxAgeHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpResponseContentSecurityPolicyHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseContentSecurityPolicyHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpResponseServerEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseServerEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpResponseStrictTransportSecurityHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseStrictTransportSecurityHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpResponseXContentTypeOptionsHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseXContentTypeOptionsHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetRoutingHttpResponseXFrameOptionsHeaderValue() {
	_jsii_.InvokeVoid(
		a,
		"resetRoutingHttpResponseXFrameOptionsHeaderValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetSslPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetSslPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetTcpIdleTimeoutSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTcpIdleTimeoutSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAlbListener) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlbListener) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlbListener) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlbListener) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlbListener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlbListener) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAlbListener) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

