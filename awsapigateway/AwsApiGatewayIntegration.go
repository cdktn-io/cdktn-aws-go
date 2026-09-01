package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapigateway/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsapigateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_integration aws_api_gateway_integration}.
// Experimental.
type AwsApiGatewayIntegration interface {
	cdktn.TerraformResource
	// Experimental.
	CacheKeyParameters() *[]*string
	// Experimental.
	SetCacheKeyParameters(val *[]*string)
	// Experimental.
	CacheKeyParametersInput() *[]*string
	// Experimental.
	CacheNamespace() *string
	// Experimental.
	SetCacheNamespace(val *string)
	// Experimental.
	CacheNamespaceInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConnectionId() *string
	// Experimental.
	SetConnectionId(val *string)
	// Experimental.
	ConnectionIdInput() *string
	// Experimental.
	ConnectionType() *string
	// Experimental.
	SetConnectionType(val *string)
	// Experimental.
	ConnectionTypeInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	ContentHandling() *string
	// Experimental.
	SetContentHandling(val *string)
	// Experimental.
	ContentHandlingInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	Credentials() *string
	// Experimental.
	SetCredentials(val *string)
	// Experimental.
	CredentialsInput() *string
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
	HttpMethod() *string
	// Experimental.
	SetHttpMethod(val *string)
	// Experimental.
	HttpMethodInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	IntegrationHttpMethod() *string
	// Experimental.
	SetIntegrationHttpMethod(val *string)
	// Experimental.
	IntegrationHttpMethodInput() *string
	// Experimental.
	IntegrationTarget() *string
	// Experimental.
	SetIntegrationTarget(val *string)
	// Experimental.
	IntegrationTargetInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PassthroughBehavior() *string
	// Experimental.
	SetPassthroughBehavior(val *string)
	// Experimental.
	PassthroughBehaviorInput() *string
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
	RequestParameters() *map[string]*string
	// Experimental.
	SetRequestParameters(val *map[string]*string)
	// Experimental.
	RequestParametersInput() *map[string]*string
	// Experimental.
	RequestTemplates() *map[string]*string
	// Experimental.
	SetRequestTemplates(val *map[string]*string)
	// Experimental.
	RequestTemplatesInput() *map[string]*string
	// Experimental.
	ResourceId() *string
	// Experimental.
	SetResourceId(val *string)
	// Experimental.
	ResourceIdInput() *string
	// Experimental.
	ResponseTransferMode() *string
	// Experimental.
	SetResponseTransferMode(val *string)
	// Experimental.
	ResponseTransferModeInput() *string
	// Experimental.
	RestApiId() *string
	// Experimental.
	SetRestApiId(val *string)
	// Experimental.
	RestApiIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	TimeoutMilliseconds() *float64
	// Experimental.
	SetTimeoutMilliseconds(val *float64)
	// Experimental.
	TimeoutMillisecondsInput() *float64
	// Experimental.
	TlsConfig() AwsApiGatewayIntegration_TlsConfigPropertyOutputReference
	// Experimental.
	TlsConfigInput() *AwsApiGatewayIntegration_TlsConfigProperty
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
	// Experimental.
	Uri() *string
	// Experimental.
	SetUri(val *string)
	// Experimental.
	UriInput() *string
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
	PutTlsConfig(value *AwsApiGatewayIntegration_TlsConfigProperty)
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
	ResetCacheKeyParameters()
	// Experimental.
	ResetCacheNamespace()
	// Experimental.
	ResetConnectionId()
	// Experimental.
	ResetConnectionType()
	// Experimental.
	ResetContentHandling()
	// Experimental.
	ResetCredentials()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIntegrationHttpMethod()
	// Experimental.
	ResetIntegrationTarget()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPassthroughBehavior()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRequestParameters()
	// Experimental.
	ResetRequestTemplates()
	// Experimental.
	ResetResponseTransferMode()
	// Experimental.
	ResetTimeoutMilliseconds()
	// Experimental.
	ResetTlsConfig()
	// Experimental.
	ResetUri()
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

// The jsii proxy struct for AwsApiGatewayIntegration
type jsiiProxy_AwsApiGatewayIntegration struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsApiGatewayIntegration) CacheKeyParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheKeyParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) CacheKeyParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheKeyParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) CacheNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) CacheNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) ConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) ConnectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) ConnectionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) ContentHandling() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) ContentHandlingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentHandlingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) Credentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) CredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) IntegrationHttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationHttpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) IntegrationHttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationHttpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) IntegrationTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) IntegrationTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) PassthroughBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passthroughBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) PassthroughBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passthroughBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) RequestParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) RequestParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) RequestTemplates() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) RequestTemplatesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestTemplatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) ResponseTransferMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTransferMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) ResponseTransferModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTransferModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) RestApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) RestApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restApiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) TimeoutMilliseconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMilliseconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) TimeoutMillisecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMillisecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) TlsConfig() AwsApiGatewayIntegration_TlsConfigPropertyOutputReference {
	var returns AwsApiGatewayIntegration_TlsConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"tlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) TlsConfigInput() *AwsApiGatewayIntegration_TlsConfigProperty {
	var returns *AwsApiGatewayIntegration_TlsConfigProperty
	_jsii_.Get(
		j,
		"tlsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiGatewayIntegration) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_integration aws_api_gateway_integration} Resource.
// Experimental.
func NewAwsApiGatewayIntegration(scope constructs.Construct, id *string, config *AwsApiGatewayIntegrationConfig) AwsApiGatewayIntegration {
	_init_.Initialize()

	if err := validateNewAwsApiGatewayIntegrationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApiGatewayIntegration{}

	_jsii_.Create(
		"@cdktn/aws-api-gateway.AwsApiGatewayIntegration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/api_gateway_integration aws_api_gateway_integration} Resource.
// Experimental.
func NewAwsApiGatewayIntegration_Override(a AwsApiGatewayIntegration, scope constructs.Construct, id *string, config *AwsApiGatewayIntegrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-api-gateway.AwsApiGatewayIntegration",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetCacheKeyParameters(val *[]*string) {
	if err := j.validateSetCacheKeyParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheKeyParameters",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetCacheNamespace(val *string) {
	if err := j.validateSetCacheNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheNamespace",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetConnectionId(val *string) {
	if err := j.validateSetConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionId",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetConnectionType(val *string) {
	if err := j.validateSetConnectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetContentHandling(val *string) {
	if err := j.validateSetContentHandlingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentHandling",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetCredentials(val *string) {
	if err := j.validateSetCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentials",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetIntegrationHttpMethod(val *string) {
	if err := j.validateSetIntegrationHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationHttpMethod",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetIntegrationTarget(val *string) {
	if err := j.validateSetIntegrationTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationTarget",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetPassthroughBehavior(val *string) {
	if err := j.validateSetPassthroughBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passthroughBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetRequestParameters(val *map[string]*string) {
	if err := j.validateSetRequestParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestParameters",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetRequestTemplates(val *map[string]*string) {
	if err := j.validateSetRequestTemplatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTemplates",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetResponseTransferMode(val *string) {
	if err := j.validateSetResponseTransferModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseTransferMode",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetRestApiId(val *string) {
	if err := j.validateSetRestApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restApiId",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetTimeoutMilliseconds(val *float64) {
	if err := j.validateSetTimeoutMillisecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutMilliseconds",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AwsApiGatewayIntegration)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

// Generates CDKTN code for importing a AwsApiGatewayIntegration resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsApiGatewayIntegration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsApiGatewayIntegration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-api-gateway.AwsApiGatewayIntegration",
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
func AwsApiGatewayIntegration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsApiGatewayIntegration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-api-gateway.AwsApiGatewayIntegration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsApiGatewayIntegration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsApiGatewayIntegration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-api-gateway.AwsApiGatewayIntegration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsApiGatewayIntegration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsApiGatewayIntegration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-api-gateway.AwsApiGatewayIntegration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsApiGatewayIntegration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-api-gateway.AwsApiGatewayIntegration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsApiGatewayIntegration) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsApiGatewayIntegration) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApiGatewayIntegration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsApiGatewayIntegration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsApiGatewayIntegration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsApiGatewayIntegration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsApiGatewayIntegration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsApiGatewayIntegration) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsApiGatewayIntegration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsApiGatewayIntegration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsApiGatewayIntegration) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsApiGatewayIntegration) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) PutTlsConfig(value *AwsApiGatewayIntegration_TlsConfigProperty) {
	if err := a.validatePutTlsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTlsConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetCacheKeyParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheKeyParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetCacheNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetConnectionId() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetConnectionType() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetContentHandling() {
	_jsii_.InvokeVoid(
		a,
		"resetContentHandling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetIntegrationHttpMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegrationHttpMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetIntegrationTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegrationTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetPassthroughBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetPassthroughBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetRequestParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetRequestTemplates() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestTemplates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetResponseTransferMode() {
	_jsii_.InvokeVoid(
		a,
		"resetResponseTransferMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetTimeoutMilliseconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutMilliseconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetTlsConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ResetUri() {
	_jsii_.InvokeVoid(
		a,
		"resetUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiGatewayIntegration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiGatewayIntegration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiGatewayIntegration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiGatewayIntegration) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

