package awsmwaa

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmwaa/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsmwaa/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment aws_mwaa_environment}.
// Experimental.
type TfEnvironment interface {
	cdktn.TerraformResource
	// Experimental.
	AirflowConfigurationOptions() *map[string]*string
	// Experimental.
	SetAirflowConfigurationOptions(val *map[string]*string)
	// Experimental.
	AirflowConfigurationOptionsInput() *map[string]*string
	// Experimental.
	AirflowVersion() *string
	// Experimental.
	SetAirflowVersion(val *string)
	// Experimental.
	AirflowVersionInput() *string
	// Experimental.
	Arn() *string
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
	CreatedAt() *string
	// Experimental.
	DagS3Path() *string
	// Experimental.
	SetDagS3Path(val *string)
	// Experimental.
	DagS3PathInput() *string
	// Experimental.
	DatabaseVpcEndpointService() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EndpointManagement() *string
	// Experimental.
	SetEndpointManagement(val *string)
	// Experimental.
	EndpointManagementInput() *string
	// Experimental.
	EnvironmentClass() *string
	// Experimental.
	SetEnvironmentClass(val *string)
	// Experimental.
	EnvironmentClassInput() *string
	// Experimental.
	ExecutionRoleArn() *string
	// Experimental.
	SetExecutionRoleArn(val *string)
	// Experimental.
	ExecutionRoleArnInput() *string
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
	KmsKey() *string
	// Experimental.
	SetKmsKey(val *string)
	// Experimental.
	KmsKeyInput() *string
	// Experimental.
	LastUpdated() TfEnvironment_LastUpdatedPropertyList
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LoggingConfiguration() TfEnvironment_LoggingConfigurationPropertyOutputReference
	// Experimental.
	LoggingConfigurationInput() *TfEnvironment_LoggingConfigurationProperty
	// Experimental.
	MaxWebservers() *float64
	// Experimental.
	SetMaxWebservers(val *float64)
	// Experimental.
	MaxWebserversInput() *float64
	// Experimental.
	MaxWorkers() *float64
	// Experimental.
	SetMaxWorkers(val *float64)
	// Experimental.
	MaxWorkersInput() *float64
	// Experimental.
	MinWebservers() *float64
	// Experimental.
	SetMinWebservers(val *float64)
	// Experimental.
	MinWebserversInput() *float64
	// Experimental.
	MinWorkers() *float64
	// Experimental.
	SetMinWorkers(val *float64)
	// Experimental.
	MinWorkersInput() *float64
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NetworkConfiguration() TfEnvironment_NetworkConfigurationPropertyOutputReference
	// Experimental.
	NetworkConfigurationInput() *TfEnvironment_NetworkConfigurationProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PluginsS3ObjectVersion() *string
	// Experimental.
	SetPluginsS3ObjectVersion(val *string)
	// Experimental.
	PluginsS3ObjectVersionInput() *string
	// Experimental.
	PluginsS3Path() *string
	// Experimental.
	SetPluginsS3Path(val *string)
	// Experimental.
	PluginsS3PathInput() *string
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
	RequirementsS3ObjectVersion() *string
	// Experimental.
	SetRequirementsS3ObjectVersion(val *string)
	// Experimental.
	RequirementsS3ObjectVersionInput() *string
	// Experimental.
	RequirementsS3Path() *string
	// Experimental.
	SetRequirementsS3Path(val *string)
	// Experimental.
	RequirementsS3PathInput() *string
	// Experimental.
	Schedulers() *float64
	// Experimental.
	SetSchedulers(val *float64)
	// Experimental.
	SchedulersInput() *float64
	// Experimental.
	ServiceRoleArn() *string
	// Experimental.
	SourceBucketArn() *string
	// Experimental.
	SetSourceBucketArn(val *string)
	// Experimental.
	SourceBucketArnInput() *string
	// Experimental.
	StartupScriptS3ObjectVersion() *string
	// Experimental.
	SetStartupScriptS3ObjectVersion(val *string)
	// Experimental.
	StartupScriptS3ObjectVersionInput() *string
	// Experimental.
	StartupScriptS3Path() *string
	// Experimental.
	SetStartupScriptS3Path(val *string)
	// Experimental.
	StartupScriptS3PathInput() *string
	// Experimental.
	Status() *string
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
	Timeouts() TfEnvironment_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	WebserverAccessMode() *string
	// Experimental.
	SetWebserverAccessMode(val *string)
	// Experimental.
	WebserverAccessModeInput() *string
	// Experimental.
	WebserverUrl() *string
	// Experimental.
	WebserverVpcEndpointService() *string
	// Experimental.
	WeeklyMaintenanceWindowStart() *string
	// Experimental.
	SetWeeklyMaintenanceWindowStart(val *string)
	// Experimental.
	WeeklyMaintenanceWindowStartInput() *string
	// Experimental.
	WorkerReplacementStrategy() *string
	// Experimental.
	SetWorkerReplacementStrategy(val *string)
	// Experimental.
	WorkerReplacementStrategyInput() *string
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
	PutLoggingConfiguration(value *TfEnvironment_LoggingConfigurationProperty)
	// Experimental.
	PutNetworkConfiguration(value *TfEnvironment_NetworkConfigurationProperty)
	// Experimental.
	PutTimeouts(value *TfEnvironment_TimeoutsProperty)
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
	ResetAirflowConfigurationOptions()
	// Experimental.
	ResetAirflowVersion()
	// Experimental.
	ResetEndpointManagement()
	// Experimental.
	ResetEnvironmentClass()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKmsKey()
	// Experimental.
	ResetLoggingConfiguration()
	// Experimental.
	ResetMaxWebservers()
	// Experimental.
	ResetMaxWorkers()
	// Experimental.
	ResetMinWebservers()
	// Experimental.
	ResetMinWorkers()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPluginsS3ObjectVersion()
	// Experimental.
	ResetPluginsS3Path()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRequirementsS3ObjectVersion()
	// Experimental.
	ResetRequirementsS3Path()
	// Experimental.
	ResetSchedulers()
	// Experimental.
	ResetStartupScriptS3ObjectVersion()
	// Experimental.
	ResetStartupScriptS3Path()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetWebserverAccessMode()
	// Experimental.
	ResetWeeklyMaintenanceWindowStart()
	// Experimental.
	ResetWorkerReplacementStrategy()
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

// The jsii proxy struct for TfEnvironment
type jsiiProxy_TfEnvironment struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfEnvironment) AirflowConfigurationOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"airflowConfigurationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) AirflowConfigurationOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"airflowConfigurationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) AirflowVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"airflowVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) AirflowVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"airflowVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) DagS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dagS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) DagS3PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dagS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) DatabaseVpcEndpointService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVpcEndpointService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) EndpointManagement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) EndpointManagementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) EnvironmentClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) EnvironmentClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) LastUpdated() TfEnvironment_LastUpdatedPropertyList {
	var returns TfEnvironment_LastUpdatedPropertyList
	_jsii_.Get(
		j,
		"lastUpdated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) LoggingConfiguration() TfEnvironment_LoggingConfigurationPropertyOutputReference {
	var returns TfEnvironment_LoggingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"loggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) LoggingConfigurationInput() *TfEnvironment_LoggingConfigurationProperty {
	var returns *TfEnvironment_LoggingConfigurationProperty
	_jsii_.Get(
		j,
		"loggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) MaxWebservers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWebservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) MaxWebserversInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWebserversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) MaxWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) MaxWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) MinWebservers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWebservers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) MinWebserversInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWebserversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) MinWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) MinWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) NetworkConfiguration() TfEnvironment_NetworkConfigurationPropertyOutputReference {
	var returns TfEnvironment_NetworkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) NetworkConfigurationInput() *TfEnvironment_NetworkConfigurationProperty {
	var returns *TfEnvironment_NetworkConfigurationProperty
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) PluginsS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) PluginsS3ObjectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3ObjectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) PluginsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) PluginsS3PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginsS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) RequirementsS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) RequirementsS3ObjectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3ObjectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) RequirementsS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) RequirementsS3PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requirementsS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) Schedulers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) SchedulersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"schedulersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) ServiceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) SourceBucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) SourceBucketArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceBucketArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) StartupScriptS3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptS3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) StartupScriptS3ObjectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptS3ObjectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) StartupScriptS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) StartupScriptS3PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) Timeouts() TfEnvironment_TimeoutsPropertyOutputReference {
	var returns TfEnvironment_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) WebserverAccessMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverAccessMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) WebserverAccessModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverAccessModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) WebserverUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) WebserverVpcEndpointService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webserverVpcEndpointService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) WeeklyMaintenanceWindowStart() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceWindowStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) WeeklyMaintenanceWindowStartInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceWindowStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) WorkerReplacementStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerReplacementStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEnvironment) WorkerReplacementStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerReplacementStrategyInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment aws_mwaa_environment} Resource.
// Experimental.
func NewTfEnvironment(scope constructs.Construct, id *string, config *TfEnvironmentConfig) TfEnvironment {
	_init_.Initialize()

	if err := validateNewTfEnvironmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEnvironment{}

	_jsii_.Create(
		"@cdktn/aws-mwaa.TfEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment aws_mwaa_environment} Resource.
// Experimental.
func NewTfEnvironment_Override(t TfEnvironment, scope constructs.Construct, id *string, config *TfEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-mwaa.TfEnvironment",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfEnvironment)SetAirflowConfigurationOptions(val *map[string]*string) {
	if err := j.validateSetAirflowConfigurationOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"airflowConfigurationOptions",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetAirflowVersion(val *string) {
	if err := j.validateSetAirflowVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"airflowVersion",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetDagS3Path(val *string) {
	if err := j.validateSetDagS3PathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dagS3Path",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetEndpointManagement(val *string) {
	if err := j.validateSetEndpointManagementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointManagement",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetEnvironmentClass(val *string) {
	if err := j.validateSetEnvironmentClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentClass",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetMaxWebservers(val *float64) {
	if err := j.validateSetMaxWebserversParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWebservers",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetMaxWorkers(val *float64) {
	if err := j.validateSetMaxWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxWorkers",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetMinWebservers(val *float64) {
	if err := j.validateSetMinWebserversParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minWebservers",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetMinWorkers(val *float64) {
	if err := j.validateSetMinWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minWorkers",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetPluginsS3ObjectVersion(val *string) {
	if err := j.validateSetPluginsS3ObjectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginsS3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetPluginsS3Path(val *string) {
	if err := j.validateSetPluginsS3PathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginsS3Path",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetRequirementsS3ObjectVersion(val *string) {
	if err := j.validateSetRequirementsS3ObjectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirementsS3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetRequirementsS3Path(val *string) {
	if err := j.validateSetRequirementsS3PathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requirementsS3Path",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetSchedulers(val *float64) {
	if err := j.validateSetSchedulersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedulers",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetSourceBucketArn(val *string) {
	if err := j.validateSetSourceBucketArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceBucketArn",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetStartupScriptS3ObjectVersion(val *string) {
	if err := j.validateSetStartupScriptS3ObjectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startupScriptS3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetStartupScriptS3Path(val *string) {
	if err := j.validateSetStartupScriptS3PathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startupScriptS3Path",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetWebserverAccessMode(val *string) {
	if err := j.validateSetWebserverAccessModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webserverAccessMode",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetWeeklyMaintenanceWindowStart(val *string) {
	if err := j.validateSetWeeklyMaintenanceWindowStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeklyMaintenanceWindowStart",
		val,
	)
}

func (j *jsiiProxy_TfEnvironment)SetWorkerReplacementStrategy(val *string) {
	if err := j.validateSetWorkerReplacementStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerReplacementStrategy",
		val,
	)
}

// Generates CDKTN code for importing a TfEnvironment resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfEnvironment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfEnvironment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-mwaa.TfEnvironment",
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
func TfEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-mwaa.TfEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfEnvironment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfEnvironment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-mwaa.TfEnvironment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfEnvironment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfEnvironment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-mwaa.TfEnvironment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfEnvironment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-mwaa.TfEnvironment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfEnvironment) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfEnvironment) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfEnvironment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEnvironment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEnvironment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEnvironment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEnvironment) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEnvironment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEnvironment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEnvironment) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfEnvironment) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEnvironment) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfEnvironment) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfEnvironment) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfEnvironment) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfEnvironment) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfEnvironment) PutLoggingConfiguration(value *TfEnvironment_LoggingConfigurationProperty) {
	if err := t.validatePutLoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLoggingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEnvironment) PutNetworkConfiguration(value *TfEnvironment_NetworkConfigurationProperty) {
	if err := t.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEnvironment) PutTimeouts(value *TfEnvironment_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEnvironment) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfEnvironment) ResetAirflowConfigurationOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetAirflowConfigurationOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetAirflowVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetAirflowVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetEndpointManagement() {
	_jsii_.InvokeVoid(
		t,
		"resetEndpointManagement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetEnvironmentClass() {
	_jsii_.InvokeVoid(
		t,
		"resetEnvironmentClass",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetKmsKey() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetLoggingConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetLoggingConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetMaxWebservers() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxWebservers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetMaxWorkers() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxWorkers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetMinWebservers() {
	_jsii_.InvokeVoid(
		t,
		"resetMinWebservers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetMinWorkers() {
	_jsii_.InvokeVoid(
		t,
		"resetMinWorkers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetPluginsS3ObjectVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetPluginsS3ObjectVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetPluginsS3Path() {
	_jsii_.InvokeVoid(
		t,
		"resetPluginsS3Path",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetRequirementsS3ObjectVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetRequirementsS3ObjectVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetRequirementsS3Path() {
	_jsii_.InvokeVoid(
		t,
		"resetRequirementsS3Path",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetSchedulers() {
	_jsii_.InvokeVoid(
		t,
		"resetSchedulers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetStartupScriptS3ObjectVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetStartupScriptS3ObjectVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetStartupScriptS3Path() {
	_jsii_.InvokeVoid(
		t,
		"resetStartupScriptS3Path",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetWebserverAccessMode() {
	_jsii_.InvokeVoid(
		t,
		"resetWebserverAccessMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetWeeklyMaintenanceWindowStart() {
	_jsii_.InvokeVoid(
		t,
		"resetWeeklyMaintenanceWindowStart",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) ResetWorkerReplacementStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetWorkerReplacementStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEnvironment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEnvironment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEnvironment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEnvironment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEnvironment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEnvironment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEnvironment) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

