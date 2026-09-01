package awsamplify

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsamplify/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsamplify/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app aws_amplify_app}.
// Experimental.
type AwsAmplifyApp interface {
	cdktn.TerraformResource
	// Experimental.
	AccessToken() *string
	// Experimental.
	SetAccessToken(val *string)
	// Experimental.
	AccessTokenInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	AutoBranchCreationConfig() AwsAmplifyApp_AutoBranchCreationConfigPropertyOutputReference
	// Experimental.
	AutoBranchCreationConfigInput() *AwsAmplifyApp_AutoBranchCreationConfigProperty
	// Experimental.
	AutoBranchCreationPatterns() *[]*string
	// Experimental.
	SetAutoBranchCreationPatterns(val *[]*string)
	// Experimental.
	AutoBranchCreationPatternsInput() *[]*string
	// Experimental.
	BasicAuthCredentials() *string
	// Experimental.
	SetBasicAuthCredentials(val *string)
	// Experimental.
	BasicAuthCredentialsInput() *string
	// Experimental.
	BuildSpec() *string
	// Experimental.
	SetBuildSpec(val *string)
	// Experimental.
	BuildSpecInput() *string
	// Experimental.
	CacheConfig() AwsAmplifyApp_CacheConfigPropertyOutputReference
	// Experimental.
	CacheConfigInput() *AwsAmplifyApp_CacheConfigProperty
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ComputeRoleArn() *string
	// Experimental.
	SetComputeRoleArn(val *string)
	// Experimental.
	ComputeRoleArnInput() *string
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
	CustomHeaders() *string
	// Experimental.
	SetCustomHeaders(val *string)
	// Experimental.
	CustomHeadersInput() *string
	// Experimental.
	CustomRule() AwsAmplifyApp_CustomRulePropertyList
	// Experimental.
	CustomRuleInput() interface{}
	// Experimental.
	DefaultDomain() *string
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
	EnableAutoBranchCreation() interface{}
	// Experimental.
	SetEnableAutoBranchCreation(val interface{})
	// Experimental.
	EnableAutoBranchCreationInput() interface{}
	// Experimental.
	EnableBasicAuth() interface{}
	// Experimental.
	SetEnableBasicAuth(val interface{})
	// Experimental.
	EnableBasicAuthInput() interface{}
	// Experimental.
	EnableBranchAutoBuild() interface{}
	// Experimental.
	SetEnableBranchAutoBuild(val interface{})
	// Experimental.
	EnableBranchAutoBuildInput() interface{}
	// Experimental.
	EnableBranchAutoDeletion() interface{}
	// Experimental.
	SetEnableBranchAutoDeletion(val interface{})
	// Experimental.
	EnableBranchAutoDeletionInput() interface{}
	// Experimental.
	EnvironmentVariables() *map[string]*string
	// Experimental.
	SetEnvironmentVariables(val *map[string]*string)
	// Experimental.
	EnvironmentVariablesInput() *map[string]*string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	IamServiceRoleArn() *string
	// Experimental.
	SetIamServiceRoleArn(val *string)
	// Experimental.
	IamServiceRoleArnInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	JobConfig() AwsAmplifyApp_JobConfigPropertyOutputReference
	// Experimental.
	JobConfigInput() *AwsAmplifyApp_JobConfigProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
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
	OauthToken() *string
	// Experimental.
	SetOauthToken(val *string)
	// Experimental.
	OauthTokenInput() *string
	// Experimental.
	Platform() *string
	// Experimental.
	SetPlatform(val *string)
	// Experimental.
	PlatformInput() *string
	// Experimental.
	ProductionBranch() AwsAmplifyApp_ProductionBranchPropertyList
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
	Repository() *string
	// Experimental.
	SetRepository(val *string)
	// Experimental.
	RepositoryInput() *string
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
	PutAutoBranchCreationConfig(value *AwsAmplifyApp_AutoBranchCreationConfigProperty)
	// Experimental.
	PutCacheConfig(value *AwsAmplifyApp_CacheConfigProperty)
	// Experimental.
	PutCustomRule(value interface{})
	// Experimental.
	PutJobConfig(value *AwsAmplifyApp_JobConfigProperty)
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
	ResetAccessToken()
	// Experimental.
	ResetAutoBranchCreationConfig()
	// Experimental.
	ResetAutoBranchCreationPatterns()
	// Experimental.
	ResetBasicAuthCredentials()
	// Experimental.
	ResetBuildSpec()
	// Experimental.
	ResetCacheConfig()
	// Experimental.
	ResetComputeRoleArn()
	// Experimental.
	ResetCustomHeaders()
	// Experimental.
	ResetCustomRule()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetEnableAutoBranchCreation()
	// Experimental.
	ResetEnableBasicAuth()
	// Experimental.
	ResetEnableBranchAutoBuild()
	// Experimental.
	ResetEnableBranchAutoDeletion()
	// Experimental.
	ResetEnvironmentVariables()
	// Experimental.
	ResetIamServiceRoleArn()
	// Experimental.
	ResetId()
	// Experimental.
	ResetJobConfig()
	// Experimental.
	ResetOauthToken()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPlatform()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRepository()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
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

// The jsii proxy struct for AwsAmplifyApp
type jsiiProxy_AwsAmplifyApp struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsAmplifyApp) AccessToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) AccessTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) AutoBranchCreationConfig() AwsAmplifyApp_AutoBranchCreationConfigPropertyOutputReference {
	var returns AwsAmplifyApp_AutoBranchCreationConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"autoBranchCreationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) AutoBranchCreationConfigInput() *AwsAmplifyApp_AutoBranchCreationConfigProperty {
	var returns *AwsAmplifyApp_AutoBranchCreationConfigProperty
	_jsii_.Get(
		j,
		"autoBranchCreationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) AutoBranchCreationPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoBranchCreationPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) AutoBranchCreationPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"autoBranchCreationPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) BasicAuthCredentials() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicAuthCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) BasicAuthCredentialsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicAuthCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) BuildSpec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) BuildSpecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) CacheConfig() AwsAmplifyApp_CacheConfigPropertyOutputReference {
	var returns AwsAmplifyApp_CacheConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"cacheConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) CacheConfigInput() *AwsAmplifyApp_CacheConfigProperty {
	var returns *AwsAmplifyApp_CacheConfigProperty
	_jsii_.Get(
		j,
		"cacheConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) ComputeRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) ComputeRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) CustomHeaders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) CustomHeadersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) CustomRule() AwsAmplifyApp_CustomRulePropertyList {
	var returns AwsAmplifyApp_CustomRulePropertyList
	_jsii_.Get(
		j,
		"customRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) CustomRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) DefaultDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) EnableAutoBranchCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoBranchCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) EnableAutoBranchCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutoBranchCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) EnableBasicAuth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBasicAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) EnableBasicAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBasicAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) EnableBranchAutoBuild() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBranchAutoBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) EnableBranchAutoBuildInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBranchAutoBuildInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) EnableBranchAutoDeletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBranchAutoDeletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) EnableBranchAutoDeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBranchAutoDeletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) IamServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamServiceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) IamServiceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamServiceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) JobConfig() AwsAmplifyApp_JobConfigPropertyOutputReference {
	var returns AwsAmplifyApp_JobConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"jobConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) JobConfigInput() *AwsAmplifyApp_JobConfigProperty {
	var returns *AwsAmplifyApp_JobConfigProperty
	_jsii_.Get(
		j,
		"jobConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) OauthToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) OauthTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauthTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) PlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) ProductionBranch() AwsAmplifyApp_ProductionBranchPropertyList {
	var returns AwsAmplifyApp_ProductionBranchPropertyList
	_jsii_.Get(
		j,
		"productionBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) Repository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) RepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAmplifyApp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app aws_amplify_app} Resource.
// Experimental.
func NewAwsAmplifyApp(scope constructs.Construct, id *string, config *AwsAmplifyAppConfig) AwsAmplifyApp {
	_init_.Initialize()

	if err := validateNewAwsAmplifyAppParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAmplifyApp{}

	_jsii_.Create(
		"@cdktn/aws-amplify.AwsAmplifyApp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_app aws_amplify_app} Resource.
// Experimental.
func NewAwsAmplifyApp_Override(a AwsAmplifyApp, scope constructs.Construct, id *string, config *AwsAmplifyAppConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-amplify.AwsAmplifyApp",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetAccessToken(val *string) {
	if err := j.validateSetAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetAutoBranchCreationPatterns(val *[]*string) {
	if err := j.validateSetAutoBranchCreationPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoBranchCreationPatterns",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetBasicAuthCredentials(val *string) {
	if err := j.validateSetBasicAuthCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"basicAuthCredentials",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetBuildSpec(val *string) {
	if err := j.validateSetBuildSpecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildSpec",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetComputeRoleArn(val *string) {
	if err := j.validateSetComputeRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetCustomHeaders(val *string) {
	if err := j.validateSetCustomHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customHeaders",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetEnableAutoBranchCreation(val interface{}) {
	if err := j.validateSetEnableAutoBranchCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutoBranchCreation",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetEnableBasicAuth(val interface{}) {
	if err := j.validateSetEnableBasicAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBasicAuth",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetEnableBranchAutoBuild(val interface{}) {
	if err := j.validateSetEnableBranchAutoBuildParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBranchAutoBuild",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetEnableBranchAutoDeletion(val interface{}) {
	if err := j.validateSetEnableBranchAutoDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBranchAutoDeletion",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetIamServiceRoleArn(val *string) {
	if err := j.validateSetIamServiceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamServiceRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetOauthToken(val *string) {
	if err := j.validateSetOauthTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthToken",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetPlatform(val *string) {
	if err := j.validateSetPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platform",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetRepository(val *string) {
	if err := j.validateSetRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repository",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsAmplifyApp)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a AwsAmplifyApp resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsAmplifyApp_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsAmplifyApp_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-amplify.AwsAmplifyApp",
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
func AwsAmplifyApp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAmplifyApp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-amplify.AwsAmplifyApp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAmplifyApp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAmplifyApp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-amplify.AwsAmplifyApp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAmplifyApp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAmplifyApp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-amplify.AwsAmplifyApp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsAmplifyApp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-amplify.AwsAmplifyApp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsAmplifyApp) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsAmplifyApp) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsAmplifyApp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAmplifyApp) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAmplifyApp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAmplifyApp) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAmplifyApp) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAmplifyApp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAmplifyApp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAmplifyApp) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAmplifyApp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAmplifyApp) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAmplifyApp) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsAmplifyApp) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAmplifyApp) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsAmplifyApp) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAmplifyApp) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsAmplifyApp) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAmplifyApp) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsAmplifyApp) PutAutoBranchCreationConfig(value *AwsAmplifyApp_AutoBranchCreationConfigProperty) {
	if err := a.validatePutAutoBranchCreationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutoBranchCreationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAmplifyApp) PutCacheConfig(value *AwsAmplifyApp_CacheConfigProperty) {
	if err := a.validatePutCacheConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCacheConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAmplifyApp) PutCustomRule(value interface{}) {
	if err := a.validatePutCustomRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomRule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAmplifyApp) PutJobConfig(value *AwsAmplifyApp_JobConfigProperty) {
	if err := a.validatePutJobConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJobConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAmplifyApp) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetAccessToken() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetAutoBranchCreationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoBranchCreationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetAutoBranchCreationPatterns() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoBranchCreationPatterns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetBasicAuthCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetBasicAuthCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetBuildSpec() {
	_jsii_.InvokeVoid(
		a,
		"resetBuildSpec",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetCacheConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetComputeRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetComputeRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetCustomHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetCustomRule() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomRule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetEnableAutoBranchCreation() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableAutoBranchCreation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetEnableBasicAuth() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBasicAuth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetEnableBranchAutoBuild() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBranchAutoBuild",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetEnableBranchAutoDeletion() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableBranchAutoDeletion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetIamServiceRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetIamServiceRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetJobConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetJobConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetOauthToken() {
	_jsii_.InvokeVoid(
		a,
		"resetOauthToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetPlatform() {
	_jsii_.InvokeVoid(
		a,
		"resetPlatform",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetRepository() {
	_jsii_.InvokeVoid(
		a,
		"resetRepository",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAmplifyApp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAmplifyApp) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAmplifyApp) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAmplifyApp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAmplifyApp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAmplifyApp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAmplifyApp) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

