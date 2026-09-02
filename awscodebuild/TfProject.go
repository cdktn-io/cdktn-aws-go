package awscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodebuild/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awscodebuild/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project aws_codebuild_project}.
// Experimental.
type TfProject interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	Artifacts() TfProject_ArtifactsPropertyOutputReference
	// Experimental.
	ArtifactsInput() *TfProject_ArtifactsProperty
	// Experimental.
	AutoRetryLimit() *float64
	// Experimental.
	SetAutoRetryLimit(val *float64)
	// Experimental.
	AutoRetryLimitInput() *float64
	// Experimental.
	BadgeEnabled() interface{}
	// Experimental.
	SetBadgeEnabled(val interface{})
	// Experimental.
	BadgeEnabledInput() interface{}
	// Experimental.
	BadgeUrl() *string
	// Experimental.
	BuildBatchConfig() TfProject_BuildBatchConfigPropertyOutputReference
	// Experimental.
	BuildBatchConfigInput() *TfProject_BuildBatchConfigProperty
	// Experimental.
	BuildTimeout() *float64
	// Experimental.
	SetBuildTimeout(val *float64)
	// Experimental.
	BuildTimeoutInput() *float64
	// Experimental.
	Cache() TfProject_CachePropertyOutputReference
	// Experimental.
	CacheInput() *TfProject_CacheProperty
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConcurrentBuildLimit() *float64
	// Experimental.
	SetConcurrentBuildLimit(val *float64)
	// Experimental.
	ConcurrentBuildLimitInput() *float64
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
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	EncryptionKey() *string
	// Experimental.
	SetEncryptionKey(val *string)
	// Experimental.
	EncryptionKeyInput() *string
	// Experimental.
	Environment() TfProject_EnvironmentPropertyOutputReference
	// Experimental.
	EnvironmentInput() *TfProject_EnvironmentProperty
	// Experimental.
	FileSystemLocations() TfProject_FileSystemLocationsPropertyList
	// Experimental.
	FileSystemLocationsInput() interface{}
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
	LogsConfig() TfProject_LogsConfigPropertyOutputReference
	// Experimental.
	LogsConfigInput() *TfProject_LogsConfigProperty
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
	ProjectVisibility() *string
	// Experimental.
	SetProjectVisibility(val *string)
	// Experimental.
	ProjectVisibilityInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	PublicProjectAlias() *string
	// Experimental.
	QueuedTimeout() *float64
	// Experimental.
	SetQueuedTimeout(val *float64)
	// Experimental.
	QueuedTimeoutInput() *float64
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	ResourceAccessRole() *string
	// Experimental.
	SetResourceAccessRole(val *string)
	// Experimental.
	ResourceAccessRoleInput() *string
	// Experimental.
	SecondaryArtifacts() TfProject_SecondaryArtifactsPropertyList
	// Experimental.
	SecondaryArtifactsInput() interface{}
	// Experimental.
	SecondarySources() TfProject_SecondarySourcesPropertyList
	// Experimental.
	SecondarySourcesInput() interface{}
	// Experimental.
	SecondarySourceVersion() TfProject_SecondarySourceVersionPropertyList
	// Experimental.
	SecondarySourceVersionInput() interface{}
	// Experimental.
	ServiceRole() *string
	// Experimental.
	SetServiceRole(val *string)
	// Experimental.
	ServiceRoleInput() *string
	// Experimental.
	Source() TfProject_SourcePropertyOutputReference
	// Experimental.
	SourceInput() *TfProject_SourceProperty
	// Experimental.
	SourceVersion() *string
	// Experimental.
	SetSourceVersion(val *string)
	// Experimental.
	SourceVersionInput() *string
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
	VpcConfig() TfProject_VpcConfigPropertyOutputReference
	// Experimental.
	VpcConfigInput() *TfProject_VpcConfigProperty
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
	PutArtifacts(value *TfProject_ArtifactsProperty)
	// Experimental.
	PutBuildBatchConfig(value *TfProject_BuildBatchConfigProperty)
	// Experimental.
	PutCache(value *TfProject_CacheProperty)
	// Experimental.
	PutEnvironment(value *TfProject_EnvironmentProperty)
	// Experimental.
	PutFileSystemLocations(value interface{})
	// Experimental.
	PutLogsConfig(value *TfProject_LogsConfigProperty)
	// Experimental.
	PutSecondaryArtifacts(value interface{})
	// Experimental.
	PutSecondarySources(value interface{})
	// Experimental.
	PutSecondarySourceVersion(value interface{})
	// Experimental.
	PutSource(value *TfProject_SourceProperty)
	// Experimental.
	PutVpcConfig(value *TfProject_VpcConfigProperty)
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
	ResetAutoRetryLimit()
	// Experimental.
	ResetBadgeEnabled()
	// Experimental.
	ResetBuildBatchConfig()
	// Experimental.
	ResetBuildTimeout()
	// Experimental.
	ResetCache()
	// Experimental.
	ResetConcurrentBuildLimit()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetEncryptionKey()
	// Experimental.
	ResetFileSystemLocations()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLogsConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetProjectVisibility()
	// Experimental.
	ResetQueuedTimeout()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetResourceAccessRole()
	// Experimental.
	ResetSecondaryArtifacts()
	// Experimental.
	ResetSecondarySources()
	// Experimental.
	ResetSecondarySourceVersion()
	// Experimental.
	ResetSourceVersion()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetVpcConfig()
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

// The jsii proxy struct for TfProject
type jsiiProxy_TfProject struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfProject) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) Artifacts() TfProject_ArtifactsPropertyOutputReference {
	var returns TfProject_ArtifactsPropertyOutputReference
	_jsii_.Get(
		j,
		"artifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) ArtifactsInput() *TfProject_ArtifactsProperty {
	var returns *TfProject_ArtifactsProperty
	_jsii_.Get(
		j,
		"artifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) AutoRetryLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRetryLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) AutoRetryLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoRetryLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) BadgeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"badgeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) BadgeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"badgeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) BadgeUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"badgeUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) BuildBatchConfig() TfProject_BuildBatchConfigPropertyOutputReference {
	var returns TfProject_BuildBatchConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"buildBatchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) BuildBatchConfigInput() *TfProject_BuildBatchConfigProperty {
	var returns *TfProject_BuildBatchConfigProperty
	_jsii_.Get(
		j,
		"buildBatchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) BuildTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"buildTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) BuildTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"buildTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) Cache() TfProject_CachePropertyOutputReference {
	var returns TfProject_CachePropertyOutputReference
	_jsii_.Get(
		j,
		"cache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) CacheInput() *TfProject_CacheProperty {
	var returns *TfProject_CacheProperty
	_jsii_.Get(
		j,
		"cacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) ConcurrentBuildLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"concurrentBuildLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) ConcurrentBuildLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"concurrentBuildLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) EncryptionKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) EncryptionKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) Environment() TfProject_EnvironmentPropertyOutputReference {
	var returns TfProject_EnvironmentPropertyOutputReference
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) EnvironmentInput() *TfProject_EnvironmentProperty {
	var returns *TfProject_EnvironmentProperty
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) FileSystemLocations() TfProject_FileSystemLocationsPropertyList {
	var returns TfProject_FileSystemLocationsPropertyList
	_jsii_.Get(
		j,
		"fileSystemLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) FileSystemLocationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileSystemLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) LogsConfig() TfProject_LogsConfigPropertyOutputReference {
	var returns TfProject_LogsConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"logsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) LogsConfigInput() *TfProject_LogsConfigProperty {
	var returns *TfProject_LogsConfigProperty
	_jsii_.Get(
		j,
		"logsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) ProjectVisibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) ProjectVisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) PublicProjectAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicProjectAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) QueuedTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queuedTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) QueuedTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queuedTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) ResourceAccessRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceAccessRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) ResourceAccessRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceAccessRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) SecondaryArtifacts() TfProject_SecondaryArtifactsPropertyList {
	var returns TfProject_SecondaryArtifactsPropertyList
	_jsii_.Get(
		j,
		"secondaryArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) SecondaryArtifactsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) SecondarySources() TfProject_SecondarySourcesPropertyList {
	var returns TfProject_SecondarySourcesPropertyList
	_jsii_.Get(
		j,
		"secondarySources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) SecondarySourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondarySourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) SecondarySourceVersion() TfProject_SecondarySourceVersionPropertyList {
	var returns TfProject_SecondarySourceVersionPropertyList
	_jsii_.Get(
		j,
		"secondarySourceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) SecondarySourceVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondarySourceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) ServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) ServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) Source() TfProject_SourcePropertyOutputReference {
	var returns TfProject_SourcePropertyOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) SourceInput() *TfProject_SourceProperty {
	var returns *TfProject_SourceProperty
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) SourceVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) SourceVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) VpcConfig() TfProject_VpcConfigPropertyOutputReference {
	var returns TfProject_VpcConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProject) VpcConfigInput() *TfProject_VpcConfigProperty {
	var returns *TfProject_VpcConfigProperty
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project aws_codebuild_project} Resource.
// Experimental.
func NewTfProject(scope constructs.Construct, id *string, config *TfProjectConfig) TfProject {
	_init_.Initialize()

	if err := validateNewTfProjectParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfProject{}

	_jsii_.Create(
		"@cdktn/aws-codebuild.TfProject",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project aws_codebuild_project} Resource.
// Experimental.
func NewTfProject_Override(t TfProject, scope constructs.Construct, id *string, config *TfProjectConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codebuild.TfProject",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfProject)SetAutoRetryLimit(val *float64) {
	if err := j.validateSetAutoRetryLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRetryLimit",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetBadgeEnabled(val interface{}) {
	if err := j.validateSetBadgeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"badgeEnabled",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetBuildTimeout(val *float64) {
	if err := j.validateSetBuildTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildTimeout",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetConcurrentBuildLimit(val *float64) {
	if err := j.validateSetConcurrentBuildLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"concurrentBuildLimit",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetEncryptionKey(val *string) {
	if err := j.validateSetEncryptionKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionKey",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetProjectVisibility(val *string) {
	if err := j.validateSetProjectVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"projectVisibility",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetQueuedTimeout(val *float64) {
	if err := j.validateSetQueuedTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queuedTimeout",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetResourceAccessRole(val *string) {
	if err := j.validateSetResourceAccessRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceAccessRole",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetServiceRole(val *string) {
	if err := j.validateSetServiceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRole",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetSourceVersion(val *string) {
	if err := j.validateSetSourceVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceVersion",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfProject)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a TfProject resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfProject_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfProject_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-codebuild.TfProject",
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
func TfProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfProject_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-codebuild.TfProject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfProject_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfProject_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-codebuild.TfProject",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfProject_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfProject_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-codebuild.TfProject",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfProject_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-codebuild.TfProject",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfProject) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfProject) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfProject) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfProject) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProject) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfProject) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfProject) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfProject) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfProject) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfProject) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfProject) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfProject) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfProject) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProject) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfProject) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfProject) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfProject) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfProject) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfProject) PutArtifacts(value *TfProject_ArtifactsProperty) {
	if err := t.validatePutArtifactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putArtifacts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject) PutBuildBatchConfig(value *TfProject_BuildBatchConfigProperty) {
	if err := t.validatePutBuildBatchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBuildBatchConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject) PutCache(value *TfProject_CacheProperty) {
	if err := t.validatePutCacheParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCache",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject) PutEnvironment(value *TfProject_EnvironmentProperty) {
	if err := t.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject) PutFileSystemLocations(value interface{}) {
	if err := t.validatePutFileSystemLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFileSystemLocations",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject) PutLogsConfig(value *TfProject_LogsConfigProperty) {
	if err := t.validatePutLogsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLogsConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject) PutSecondaryArtifacts(value interface{}) {
	if err := t.validatePutSecondaryArtifactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSecondaryArtifacts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject) PutSecondarySources(value interface{}) {
	if err := t.validatePutSecondarySourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSecondarySources",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject) PutSecondarySourceVersion(value interface{}) {
	if err := t.validatePutSecondarySourceVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSecondarySourceVersion",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject) PutSource(value *TfProject_SourceProperty) {
	if err := t.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject) PutVpcConfig(value *TfProject_VpcConfigProperty) {
	if err := t.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfProject) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfProject) ResetAutoRetryLimit() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoRetryLimit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetBadgeEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetBadgeEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetBuildBatchConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetBuildBatchConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetBuildTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetBuildTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetCache() {
	_jsii_.InvokeVoid(
		t,
		"resetCache",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetConcurrentBuildLimit() {
	_jsii_.InvokeVoid(
		t,
		"resetConcurrentBuildLimit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetEncryptionKey() {
	_jsii_.InvokeVoid(
		t,
		"resetEncryptionKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetFileSystemLocations() {
	_jsii_.InvokeVoid(
		t,
		"resetFileSystemLocations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetLogsConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetLogsConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetProjectVisibility() {
	_jsii_.InvokeVoid(
		t,
		"resetProjectVisibility",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetQueuedTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetQueuedTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetResourceAccessRole() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceAccessRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetSecondaryArtifacts() {
	_jsii_.InvokeVoid(
		t,
		"resetSecondaryArtifacts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetSecondarySources() {
	_jsii_.InvokeVoid(
		t,
		"resetSecondarySources",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetSecondarySourceVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetSecondarySourceVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetSourceVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProject) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProject) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

