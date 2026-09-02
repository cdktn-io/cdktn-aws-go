package awselb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselb/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awselb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener aws_lb_listener}.
// Experimental.
type TfListener interface {
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
	DefaultAction() TfListener_DefaultActionPropertyList
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
	MutualAuthentication() TfListener_MutualAuthenticationPropertyOutputReference
	// Experimental.
	MutualAuthenticationInput() *TfListener_MutualAuthenticationProperty
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
	Timeouts() TfListener_TimeoutsPropertyOutputReference
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
	PutMutualAuthentication(value *TfListener_MutualAuthenticationProperty)
	// Experimental.
	PutTimeouts(value *TfListener_TimeoutsProperty)
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

// The jsii proxy struct for TfListener
type jsiiProxy_TfListener struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfListener) AlpnPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alpnPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) AlpnPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alpnPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) DefaultAction() TfListener_DefaultActionPropertyList {
	var returns TfListener_DefaultActionPropertyList
	_jsii_.Get(
		j,
		"defaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) DefaultActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) LoadBalancerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) LoadBalancerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) MutualAuthentication() TfListener_MutualAuthenticationPropertyOutputReference {
	var returns TfListener_MutualAuthenticationPropertyOutputReference
	_jsii_.Get(
		j,
		"mutualAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) MutualAuthenticationInput() *TfListener_MutualAuthenticationProperty {
	var returns *TfListener_MutualAuthenticationProperty
	_jsii_.Get(
		j,
		"mutualAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpRequestXAmznMtlsClientcertHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpRequestXAmznMtlsClientcertHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertIssuerHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpRequestXAmznMtlsClientcertIssuerHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertIssuerHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpRequestXAmznMtlsClientcertLeafHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertLeafHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpRequestXAmznMtlsClientcertLeafHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertLeafHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertSerialNumberHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertSubjectHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpRequestXAmznMtlsClientcertSubjectHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertSubjectHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpRequestXAmznMtlsClientcertValidityHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertValidityHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpRequestXAmznMtlsClientcertValidityHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznMtlsClientcertValidityHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpRequestXAmznTlsCipherSuiteHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznTlsCipherSuiteHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpRequestXAmznTlsCipherSuiteHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznTlsCipherSuiteHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpRequestXAmznTlsVersionHeaderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznTlsVersionHeaderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpRequestXAmznTlsVersionHeaderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpRequestXAmznTlsVersionHeaderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseAccessControlAllowCredentialsHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowCredentialsHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseAccessControlAllowCredentialsHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowCredentialsHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseAccessControlAllowHeadersHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowHeadersHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseAccessControlAllowHeadersHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowHeadersHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseAccessControlAllowMethodsHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowMethodsHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseAccessControlAllowMethodsHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowMethodsHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseAccessControlAllowOriginHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowOriginHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseAccessControlAllowOriginHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlAllowOriginHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseAccessControlExposeHeadersHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlExposeHeadersHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseAccessControlExposeHeadersHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlExposeHeadersHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseAccessControlMaxAgeHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlMaxAgeHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseAccessControlMaxAgeHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseAccessControlMaxAgeHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseContentSecurityPolicyHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseContentSecurityPolicyHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseContentSecurityPolicyHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseContentSecurityPolicyHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseServerEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingHttpResponseServerEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseServerEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingHttpResponseServerEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseStrictTransportSecurityHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseStrictTransportSecurityHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseStrictTransportSecurityHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseStrictTransportSecurityHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseXContentTypeOptionsHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseXContentTypeOptionsHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseXContentTypeOptionsHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseXContentTypeOptionsHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseXFrameOptionsHeaderValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseXFrameOptionsHeaderValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) RoutingHttpResponseXFrameOptionsHeaderValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routingHttpResponseXFrameOptionsHeaderValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) SslPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) SslPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) TcpIdleTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpIdleTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) TcpIdleTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tcpIdleTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) Timeouts() TfListener_TimeoutsPropertyOutputReference {
	var returns TfListener_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener aws_lb_listener} Resource.
// Experimental.
func NewTfListener(scope constructs.Construct, id *string, config *TfListenerConfig) TfListener {
	_init_.Initialize()

	if err := validateNewTfListenerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfListener{}

	_jsii_.Create(
		"@cdktn/aws-elb.TfListener",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_listener aws_lb_listener} Resource.
// Experimental.
func NewTfListener_Override(t TfListener, scope constructs.Construct, id *string, config *TfListenerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.TfListener",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfListener)SetAlpnPolicy(val *string) {
	if err := j.validateSetAlpnPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alpnPolicy",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetCertificateArn(val *string) {
	if err := j.validateSetCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetLoadBalancerArn(val *string) {
	if err := j.validateSetLoadBalancerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerArn",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpRequestXAmznMtlsClientcertHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertHeaderName",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertIssuerHeaderName",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertLeafHeaderName",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertSubjectHeaderName",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznMtlsClientcertValidityHeaderName",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpRequestXAmznTlsCipherSuiteHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznTlsCipherSuiteHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznTlsCipherSuiteHeaderName",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpRequestXAmznTlsVersionHeaderName(val *string) {
	if err := j.validateSetRoutingHttpRequestXAmznTlsVersionHeaderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpRequestXAmznTlsVersionHeaderName",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpResponseAccessControlAllowCredentialsHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlAllowCredentialsHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlAllowCredentialsHeaderValue",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpResponseAccessControlAllowHeadersHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlAllowHeadersHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlAllowHeadersHeaderValue",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpResponseAccessControlAllowMethodsHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlAllowMethodsHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlAllowMethodsHeaderValue",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpResponseAccessControlAllowOriginHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlAllowOriginHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlAllowOriginHeaderValue",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpResponseAccessControlExposeHeadersHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlExposeHeadersHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlExposeHeadersHeaderValue",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpResponseAccessControlMaxAgeHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseAccessControlMaxAgeHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseAccessControlMaxAgeHeaderValue",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpResponseContentSecurityPolicyHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseContentSecurityPolicyHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseContentSecurityPolicyHeaderValue",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpResponseServerEnabled(val interface{}) {
	if err := j.validateSetRoutingHttpResponseServerEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseServerEnabled",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpResponseStrictTransportSecurityHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseStrictTransportSecurityHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseStrictTransportSecurityHeaderValue",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpResponseXContentTypeOptionsHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseXContentTypeOptionsHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseXContentTypeOptionsHeaderValue",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetRoutingHttpResponseXFrameOptionsHeaderValue(val *string) {
	if err := j.validateSetRoutingHttpResponseXFrameOptionsHeaderValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingHttpResponseXFrameOptionsHeaderValue",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetSslPolicy(val *string) {
	if err := j.validateSetSslPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslPolicy",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfListener)SetTcpIdleTimeoutSeconds(val *float64) {
	if err := j.validateSetTcpIdleTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpIdleTimeoutSeconds",
		val,
	)
}

// Generates CDKTN code for importing a TfListener resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfListener_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfListener_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.TfListener",
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
func TfListener_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfListener_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.TfListener",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfListener_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfListener_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.TfListener",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfListener_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfListener_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elb.TfListener",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfListener_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-elb.TfListener",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfListener) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfListener) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfListener) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfListener) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfListener) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfListener) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfListener) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfListener) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfListener) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfListener) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfListener) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfListener) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfListener) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfListener) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfListener) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfListener) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfListener) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfListener) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfListener) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfListener) PutDefaultAction(value interface{}) {
	if err := t.validatePutDefaultActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDefaultAction",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListener) PutMutualAuthentication(value *TfListener_MutualAuthenticationProperty) {
	if err := t.validatePutMutualAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMutualAuthentication",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListener) PutTimeouts(value *TfListener_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListener) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfListener) ResetAlpnPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetAlpnPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		t,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetMutualAuthentication() {
	_jsii_.InvokeVoid(
		t,
		"resetMutualAuthentication",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetPort() {
	_jsii_.InvokeVoid(
		t,
		"resetPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetProtocol() {
	_jsii_.InvokeVoid(
		t,
		"resetProtocol",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpRequestXAmznMtlsClientcertHeaderName() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpRequestXAmznMtlsClientcertHeaderName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderName() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpRequestXAmznMtlsClientcertLeafHeaderName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderName() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpRequestXAmznMtlsClientcertValidityHeaderName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpRequestXAmznTlsCipherSuiteHeaderName() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpRequestXAmznTlsCipherSuiteHeaderName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpRequestXAmznTlsVersionHeaderName() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpRequestXAmznTlsVersionHeaderName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpResponseAccessControlAllowCredentialsHeaderValue() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpResponseAccessControlAllowCredentialsHeaderValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpResponseAccessControlAllowHeadersHeaderValue() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpResponseAccessControlAllowHeadersHeaderValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpResponseAccessControlAllowMethodsHeaderValue() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpResponseAccessControlAllowMethodsHeaderValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpResponseAccessControlAllowOriginHeaderValue() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpResponseAccessControlAllowOriginHeaderValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpResponseAccessControlExposeHeadersHeaderValue() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpResponseAccessControlExposeHeadersHeaderValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpResponseAccessControlMaxAgeHeaderValue() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpResponseAccessControlMaxAgeHeaderValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpResponseContentSecurityPolicyHeaderValue() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpResponseContentSecurityPolicyHeaderValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpResponseServerEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpResponseServerEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpResponseStrictTransportSecurityHeaderValue() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpResponseStrictTransportSecurityHeaderValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpResponseXContentTypeOptionsHeaderValue() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpResponseXContentTypeOptionsHeaderValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetRoutingHttpResponseXFrameOptionsHeaderValue() {
	_jsii_.InvokeVoid(
		t,
		"resetRoutingHttpResponseXFrameOptionsHeaderValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetSslPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetSslPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetTcpIdleTimeoutSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTcpIdleTimeoutSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfListener) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfListener) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfListener) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfListener) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfListener) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfListener) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

