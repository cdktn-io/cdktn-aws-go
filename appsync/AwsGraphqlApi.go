package appsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appsync/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/appsync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api aws_appsync_graphql_api}.
// Experimental.
type AwsGraphqlApi interface {
	cdktn.TerraformResource
	// Experimental.
	AdditionalAuthenticationProvider() AwsGraphqlApi_AdditionalAuthenticationProviderPropertyList
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
	EnhancedMetricsConfig() AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference
	// Experimental.
	EnhancedMetricsConfigInput() *AwsGraphqlApi_EnhancedMetricsConfigProperty
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
	LambdaAuthorizerConfig() AwsGraphqlApi_LambdaAuthorizerConfigPropertyOutputReference
	// Experimental.
	LambdaAuthorizerConfigInput() *AwsGraphqlApi_LambdaAuthorizerConfigProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LogConfig() AwsGraphqlApi_LogConfigPropertyOutputReference
	// Experimental.
	LogConfigInput() *AwsGraphqlApi_LogConfigProperty
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
	OpenidConnectConfig() AwsGraphqlApi_OpenidConnectConfigPropertyOutputReference
	// Experimental.
	OpenidConnectConfigInput() *AwsGraphqlApi_OpenidConnectConfigProperty
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
	UserPoolConfig() AwsGraphqlApi_UserPoolConfigPropertyOutputReference
	// Experimental.
	UserPoolConfigInput() *AwsGraphqlApi_UserPoolConfigProperty
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
	PutEnhancedMetricsConfig(value *AwsGraphqlApi_EnhancedMetricsConfigProperty)
	// Experimental.
	PutLambdaAuthorizerConfig(value *AwsGraphqlApi_LambdaAuthorizerConfigProperty)
	// Experimental.
	PutLogConfig(value *AwsGraphqlApi_LogConfigProperty)
	// Experimental.
	PutOpenidConnectConfig(value *AwsGraphqlApi_OpenidConnectConfigProperty)
	// Experimental.
	PutUserPoolConfig(value *AwsGraphqlApi_UserPoolConfigProperty)
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

// The jsii proxy struct for AwsGraphqlApi
type jsiiProxy_AwsGraphqlApi struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsGraphqlApi) AdditionalAuthenticationProvider() AwsGraphqlApi_AdditionalAuthenticationProviderPropertyList {
	var returns AwsGraphqlApi_AdditionalAuthenticationProviderPropertyList
	_jsii_.Get(
		j,
		"additionalAuthenticationProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) AdditionalAuthenticationProviderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalAuthenticationProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) ApiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) ApiTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) EnhancedMetricsConfig() AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference {
	var returns AwsGraphqlApi_EnhancedMetricsConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"enhancedMetricsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) EnhancedMetricsConfigInput() *AwsGraphqlApi_EnhancedMetricsConfigProperty {
	var returns *AwsGraphqlApi_EnhancedMetricsConfigProperty
	_jsii_.Get(
		j,
		"enhancedMetricsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) IntrospectionConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"introspectionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) IntrospectionConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"introspectionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) LambdaAuthorizerConfig() AwsGraphqlApi_LambdaAuthorizerConfigPropertyOutputReference {
	var returns AwsGraphqlApi_LambdaAuthorizerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) LambdaAuthorizerConfigInput() *AwsGraphqlApi_LambdaAuthorizerConfigProperty {
	var returns *AwsGraphqlApi_LambdaAuthorizerConfigProperty
	_jsii_.Get(
		j,
		"lambdaAuthorizerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) LogConfig() AwsGraphqlApi_LogConfigPropertyOutputReference {
	var returns AwsGraphqlApi_LogConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) LogConfigInput() *AwsGraphqlApi_LogConfigProperty {
	var returns *AwsGraphqlApi_LogConfigProperty
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) MergedApiExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergedApiExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) MergedApiExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergedApiExecutionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) OpenidConnectConfig() AwsGraphqlApi_OpenidConnectConfigPropertyOutputReference {
	var returns AwsGraphqlApi_OpenidConnectConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"openidConnectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) OpenidConnectConfigInput() *AwsGraphqlApi_OpenidConnectConfigProperty {
	var returns *AwsGraphqlApi_OpenidConnectConfigProperty
	_jsii_.Get(
		j,
		"openidConnectConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) QueryDepthLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryDepthLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) QueryDepthLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryDepthLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) ResolverCountLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resolverCountLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) ResolverCountLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resolverCountLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) Uris() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"uris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) UserPoolConfig() AwsGraphqlApi_UserPoolConfigPropertyOutputReference {
	var returns AwsGraphqlApi_UserPoolConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"userPoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) UserPoolConfigInput() *AwsGraphqlApi_UserPoolConfigProperty {
	var returns *AwsGraphqlApi_UserPoolConfigProperty
	_jsii_.Get(
		j,
		"userPoolConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) VisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) XrayEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"xrayEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGraphqlApi) XrayEnabledInput() interface{} {
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
func NewAwsGraphqlApi(scope constructs.Construct, id *string, config *AwsGraphqlApiConfig) AwsGraphqlApi {
	_init_.Initialize()

	if err := validateNewAwsGraphqlApiParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGraphqlApi{}

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsGraphqlApi",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api aws_appsync_graphql_api} Resource.
// Experimental.
func NewAwsGraphqlApi_Override(a AwsGraphqlApi, scope constructs.Construct, id *string, config *AwsGraphqlApiConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsGraphqlApi",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetApiType(val *string) {
	if err := j.validateSetApiTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiType",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetIntrospectionConfig(val *string) {
	if err := j.validateSetIntrospectionConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"introspectionConfig",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetMergedApiExecutionRoleArn(val *string) {
	if err := j.validateSetMergedApiExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergedApiExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetQueryDepthLimit(val *float64) {
	if err := j.validateSetQueryDepthLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryDepthLimit",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetResolverCountLimit(val *float64) {
	if err := j.validateSetResolverCountLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resolverCountLimit",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetVisibility(val *string) {
	if err := j.validateSetVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibility",
		val,
	)
}

func (j *jsiiProxy_AwsGraphqlApi)SetXrayEnabled(val interface{}) {
	if err := j.validateSetXrayEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xrayEnabled",
		val,
	)
}

// Generates CDKTN code for importing a AwsGraphqlApi resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsGraphqlApi_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsGraphqlApi_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-appsync.AwsGraphqlApi",
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
func AwsGraphqlApi_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsGraphqlApi_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-appsync.AwsGraphqlApi",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsGraphqlApi_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsGraphqlApi_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-appsync.AwsGraphqlApi",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsGraphqlApi_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsGraphqlApi_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-appsync.AwsGraphqlApi",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsGraphqlApi_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-appsync.AwsGraphqlApi",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsGraphqlApi) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsGraphqlApi) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsGraphqlApi) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGraphqlApi) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGraphqlApi) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGraphqlApi) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGraphqlApi) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGraphqlApi) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGraphqlApi) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGraphqlApi) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGraphqlApi) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGraphqlApi) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsGraphqlApi) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGraphqlApi) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsGraphqlApi) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsGraphqlApi) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsGraphqlApi) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsGraphqlApi) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsGraphqlApi) PutAdditionalAuthenticationProvider(value interface{}) {
	if err := a.validatePutAdditionalAuthenticationProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdditionalAuthenticationProvider",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGraphqlApi) PutEnhancedMetricsConfig(value *AwsGraphqlApi_EnhancedMetricsConfigProperty) {
	if err := a.validatePutEnhancedMetricsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnhancedMetricsConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGraphqlApi) PutLambdaAuthorizerConfig(value *AwsGraphqlApi_LambdaAuthorizerConfigProperty) {
	if err := a.validatePutLambdaAuthorizerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaAuthorizerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGraphqlApi) PutLogConfig(value *AwsGraphqlApi_LogConfigProperty) {
	if err := a.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGraphqlApi) PutOpenidConnectConfig(value *AwsGraphqlApi_OpenidConnectConfigProperty) {
	if err := a.validatePutOpenidConnectConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpenidConnectConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGraphqlApi) PutUserPoolConfig(value *AwsGraphqlApi_UserPoolConfigProperty) {
	if err := a.validatePutUserPoolConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserPoolConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGraphqlApi) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetAdditionalAuthenticationProvider() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalAuthenticationProvider",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetApiType() {
	_jsii_.InvokeVoid(
		a,
		"resetApiType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetEnhancedMetricsConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEnhancedMetricsConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetIntrospectionConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetIntrospectionConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetLambdaAuthorizerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaAuthorizerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetLogConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetMergedApiExecutionRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetMergedApiExecutionRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetOpenidConnectConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetOpenidConnectConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetQueryDepthLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryDepthLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetResolverCountLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetResolverCountLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetSchema() {
	_jsii_.InvokeVoid(
		a,
		"resetSchema",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetUserPoolConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetUserPoolConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetVisibility() {
	_jsii_.InvokeVoid(
		a,
		"resetVisibility",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) ResetXrayEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetXrayEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGraphqlApi) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGraphqlApi) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

