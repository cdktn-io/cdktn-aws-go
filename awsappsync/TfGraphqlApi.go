package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappsync/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsappsync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api aws_appsync_graphql_api}.
// Experimental.
type TfGraphqlApi interface {
	cdktn.TerraformResource
	// Experimental.
	AdditionalAuthenticationProvider() TfGraphqlApi_AdditionalAuthenticationProviderPropertyList
	// Experimental.
	AdditionalAuthenticationProviderInput() interface{}
	// Experimental.
	ApiType() *string
	// Experimental.
	SetApiType(val *string)
	// Experimental.
	ApiTypeInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	AuthenticationType() *string
	// Experimental.
	SetAuthenticationType(val *string)
	// Experimental.
	AuthenticationTypeInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	EnhancedMetricsConfig() TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference
	// Experimental.
	EnhancedMetricsConfigInput() *TfGraphqlApi_EnhancedMetricsConfigProperty
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
	IntrospectionConfig() *string
	// Experimental.
	SetIntrospectionConfig(val *string)
	// Experimental.
	IntrospectionConfigInput() *string
	// Experimental.
	LambdaAuthorizerConfig() TfGraphqlApi_LambdaAuthorizerConfigPropertyOutputReference
	// Experimental.
	LambdaAuthorizerConfigInput() *TfGraphqlApi_LambdaAuthorizerConfigProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LogConfig() TfGraphqlApi_LogConfigPropertyOutputReference
	// Experimental.
	LogConfigInput() *TfGraphqlApi_LogConfigProperty
	// Experimental.
	MergedApiExecutionRoleArn() *string
	// Experimental.
	SetMergedApiExecutionRoleArn(val *string)
	// Experimental.
	MergedApiExecutionRoleArnInput() *string
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OpenidConnectConfig() TfGraphqlApi_OpenidConnectConfigPropertyOutputReference
	// Experimental.
	OpenidConnectConfigInput() *TfGraphqlApi_OpenidConnectConfigProperty
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	QueryDepthLimit() *float64
	// Experimental.
	SetQueryDepthLimit(val *float64)
	// Experimental.
	QueryDepthLimitInput() *float64
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	ResolverCountLimit() *float64
	// Experimental.
	SetResolverCountLimit(val *float64)
	// Experimental.
	ResolverCountLimitInput() *float64
	// Experimental.
	Schema() *string
	// Experimental.
	SetSchema(val *string)
	// Experimental.
	SchemaInput() *string
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
	Uris() cdktn.StringMap
	// Experimental.
	UserPoolConfig() TfGraphqlApi_UserPoolConfigPropertyOutputReference
	// Experimental.
	UserPoolConfigInput() *TfGraphqlApi_UserPoolConfigProperty
	// Experimental.
	Visibility() *string
	// Experimental.
	SetVisibility(val *string)
	// Experimental.
	VisibilityInput() *string
	// Experimental.
	XrayEnabled() interface{}
	// Experimental.
	SetXrayEnabled(val interface{})
	// Experimental.
	XrayEnabledInput() interface{}
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
	PutAdditionalAuthenticationProvider(value interface{})
	// Experimental.
	PutEnhancedMetricsConfig(value *TfGraphqlApi_EnhancedMetricsConfigProperty)
	// Experimental.
	PutLambdaAuthorizerConfig(value *TfGraphqlApi_LambdaAuthorizerConfigProperty)
	// Experimental.
	PutLogConfig(value *TfGraphqlApi_LogConfigProperty)
	// Experimental.
	PutOpenidConnectConfig(value *TfGraphqlApi_OpenidConnectConfigProperty)
	// Experimental.
	PutUserPoolConfig(value *TfGraphqlApi_UserPoolConfigProperty)
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
	ResetAdditionalAuthenticationProvider()
	// Experimental.
	ResetApiType()
	// Experimental.
	ResetEnhancedMetricsConfig()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIntrospectionConfig()
	// Experimental.
	ResetLambdaAuthorizerConfig()
	// Experimental.
	ResetLogConfig()
	// Experimental.
	ResetMergedApiExecutionRoleArn()
	// Experimental.
	ResetOpenidConnectConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetQueryDepthLimit()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetResolverCountLimit()
	// Experimental.
	ResetSchema()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetUserPoolConfig()
	// Experimental.
	ResetVisibility()
	// Experimental.
	ResetXrayEnabled()
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

// The jsii proxy struct for TfGraphqlApi
type jsiiProxy_TfGraphqlApi struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfGraphqlApi) AdditionalAuthenticationProvider() TfGraphqlApi_AdditionalAuthenticationProviderPropertyList {
	var returns TfGraphqlApi_AdditionalAuthenticationProviderPropertyList
	_jsii_.Get(
		j,
		"additionalAuthenticationProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) AdditionalAuthenticationProviderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalAuthenticationProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) ApiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) ApiTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) EnhancedMetricsConfig() TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference {
	var returns TfGraphqlApi_EnhancedMetricsConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"enhancedMetricsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) EnhancedMetricsConfigInput() *TfGraphqlApi_EnhancedMetricsConfigProperty {
	var returns *TfGraphqlApi_EnhancedMetricsConfigProperty
	_jsii_.Get(
		j,
		"enhancedMetricsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) IntrospectionConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"introspectionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) IntrospectionConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"introspectionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) LambdaAuthorizerConfig() TfGraphqlApi_LambdaAuthorizerConfigPropertyOutputReference {
	var returns TfGraphqlApi_LambdaAuthorizerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) LambdaAuthorizerConfigInput() *TfGraphqlApi_LambdaAuthorizerConfigProperty {
	var returns *TfGraphqlApi_LambdaAuthorizerConfigProperty
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) LogConfig() TfGraphqlApi_LogConfigPropertyOutputReference {
	var returns TfGraphqlApi_LogConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) LogConfigInput() *TfGraphqlApi_LogConfigProperty {
	var returns *TfGraphqlApi_LogConfigProperty
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) MergedApiExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergedApiExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) MergedApiExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergedApiExecutionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) OpenidConnectConfig() TfGraphqlApi_OpenidConnectConfigPropertyOutputReference {
	var returns TfGraphqlApi_OpenidConnectConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"openidConnectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) OpenidConnectConfigInput() *TfGraphqlApi_OpenidConnectConfigProperty {
	var returns *TfGraphqlApi_OpenidConnectConfigProperty
	_jsii_.Get(
		j,
		"openidConnectConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) QueryDepthLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryDepthLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) QueryDepthLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryDepthLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) ResolverCountLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resolverCountLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) ResolverCountLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resolverCountLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) Uris() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"uris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) UserPoolConfig() TfGraphqlApi_UserPoolConfigPropertyOutputReference {
	var returns TfGraphqlApi_UserPoolConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"userPoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) UserPoolConfigInput() *TfGraphqlApi_UserPoolConfigProperty {
	var returns *TfGraphqlApi_UserPoolConfigProperty
	_jsii_.Get(
		j,
		"userPoolConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) VisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) XrayEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xrayEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGraphqlApi) XrayEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xrayEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api aws_appsync_graphql_api} Resource.
// Experimental.
func NewTfGraphqlApi(scope constructs.Construct, id *string, config *TfGraphqlApiConfig) TfGraphqlApi {
	_init_.Initialize()

	if err := validateNewTfGraphqlApiParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGraphqlApi{}

	_jsii_.Create(
		"@cdktn/aws-appsync.TfGraphqlApi",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api aws_appsync_graphql_api} Resource.
// Experimental.
func NewTfGraphqlApi_Override(t TfGraphqlApi, scope constructs.Construct, id *string, config *TfGraphqlApiConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appsync.TfGraphqlApi",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetApiType(val *string) {
	if err := j.validateSetApiTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiType",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetIntrospectionConfig(val *string) {
	if err := j.validateSetIntrospectionConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"introspectionConfig",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetMergedApiExecutionRoleArn(val *string) {
	if err := j.validateSetMergedApiExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergedApiExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetQueryDepthLimit(val *float64) {
	if err := j.validateSetQueryDepthLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryDepthLimit",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetResolverCountLimit(val *float64) {
	if err := j.validateSetResolverCountLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolverCountLimit",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetVisibility(val *string) {
	if err := j.validateSetVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibility",
		val,
	)
}

func (j *jsiiProxy_TfGraphqlApi)SetXrayEnabled(val interface{}) {
	if err := j.validateSetXrayEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xrayEnabled",
		val,
	)
}

// Generates CDKTN code for importing a TfGraphqlApi resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfGraphqlApi_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfGraphqlApi_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-appsync.TfGraphqlApi",
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
func TfGraphqlApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfGraphqlApi_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-appsync.TfGraphqlApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfGraphqlApi_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfGraphqlApi_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-appsync.TfGraphqlApi",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfGraphqlApi_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfGraphqlApi_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-appsync.TfGraphqlApi",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfGraphqlApi_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-appsync.TfGraphqlApi",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfGraphqlApi) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfGraphqlApi) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfGraphqlApi) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGraphqlApi) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGraphqlApi) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGraphqlApi) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGraphqlApi) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGraphqlApi) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGraphqlApi) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGraphqlApi) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGraphqlApi) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGraphqlApi) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGraphqlApi) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfGraphqlApi) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGraphqlApi) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfGraphqlApi) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfGraphqlApi) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfGraphqlApi) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfGraphqlApi) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfGraphqlApi) PutAdditionalAuthenticationProvider(value interface{}) {
	if err := t.validatePutAdditionalAuthenticationProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAdditionalAuthenticationProvider",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGraphqlApi) PutEnhancedMetricsConfig(value *TfGraphqlApi_EnhancedMetricsConfigProperty) {
	if err := t.validatePutEnhancedMetricsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEnhancedMetricsConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGraphqlApi) PutLambdaAuthorizerConfig(value *TfGraphqlApi_LambdaAuthorizerConfigProperty) {
	if err := t.validatePutLambdaAuthorizerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambdaAuthorizerConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGraphqlApi) PutLogConfig(value *TfGraphqlApi_LogConfigProperty) {
	if err := t.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGraphqlApi) PutOpenidConnectConfig(value *TfGraphqlApi_OpenidConnectConfigProperty) {
	if err := t.validatePutOpenidConnectConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOpenidConnectConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGraphqlApi) PutUserPoolConfig(value *TfGraphqlApi_UserPoolConfigProperty) {
	if err := t.validatePutUserPoolConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUserPoolConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGraphqlApi) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetAdditionalAuthenticationProvider() {
	_jsii_.InvokeVoid(
		t,
		"resetAdditionalAuthenticationProvider",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetApiType() {
	_jsii_.InvokeVoid(
		t,
		"resetApiType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetEnhancedMetricsConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetEnhancedMetricsConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetIntrospectionConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetIntrospectionConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetLambdaAuthorizerConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaAuthorizerConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetLogConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetMergedApiExecutionRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetMergedApiExecutionRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetOpenidConnectConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetOpenidConnectConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetQueryDepthLimit() {
	_jsii_.InvokeVoid(
		t,
		"resetQueryDepthLimit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetResolverCountLimit() {
	_jsii_.InvokeVoid(
		t,
		"resetResolverCountLimit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetSchema() {
	_jsii_.InvokeVoid(
		t,
		"resetSchema",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetUserPoolConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetUserPoolConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetVisibility() {
	_jsii_.InvokeVoid(
		t,
		"resetVisibility",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) ResetXrayEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetXrayEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGraphqlApi) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGraphqlApi) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGraphqlApi) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGraphqlApi) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGraphqlApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGraphqlApi) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGraphqlApi) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

