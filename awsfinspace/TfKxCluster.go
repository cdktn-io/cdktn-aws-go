package awsfinspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfinspace/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsfinspace/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster aws_finspace_kx_cluster}.
// Experimental.
type TfKxCluster interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	AutoScalingConfiguration() TfKxCluster_AutoScalingConfigurationPropertyOutputReference
	// Experimental.
	AutoScalingConfigurationInput() *TfKxCluster_AutoScalingConfigurationProperty
	// Experimental.
	AvailabilityZoneId() *string
	// Experimental.
	SetAvailabilityZoneId(val *string)
	// Experimental.
	AvailabilityZoneIdInput() *string
	// Experimental.
	AzMode() *string
	// Experimental.
	SetAzMode(val *string)
	// Experimental.
	AzModeInput() *string
	// Experimental.
	CacheStorageConfigurations() TfKxCluster_CacheStorageConfigurationsPropertyList
	// Experimental.
	CacheStorageConfigurationsInput() interface{}
	// Experimental.
	CapacityConfiguration() TfKxCluster_CapacityConfigurationPropertyOutputReference
	// Experimental.
	CapacityConfigurationInput() *TfKxCluster_CapacityConfigurationProperty
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Code() TfKxCluster_CodePropertyOutputReference
	// Experimental.
	CodeInput() *TfKxCluster_CodeProperty
	// Experimental.
	CommandLineArguments() *map[string]*string
	// Experimental.
	SetCommandLineArguments(val *map[string]*string)
	// Experimental.
	CommandLineArgumentsInput() *map[string]*string
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
	CreatedTimestamp() *string
	// Experimental.
	Database() TfKxCluster_DatabasePropertyList
	// Experimental.
	DatabaseInput() interface{}
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
	EnvironmentId() *string
	// Experimental.
	SetEnvironmentId(val *string)
	// Experimental.
	EnvironmentIdInput() *string
	// Experimental.
	ExecutionRole() *string
	// Experimental.
	SetExecutionRole(val *string)
	// Experimental.
	ExecutionRoleInput() *string
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
	InitializationScript() *string
	// Experimental.
	SetInitializationScript(val *string)
	// Experimental.
	InitializationScriptInput() *string
	// Experimental.
	LastModifiedTimestamp() *string
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
	ReleaseLabel() *string
	// Experimental.
	SetReleaseLabel(val *string)
	// Experimental.
	ReleaseLabelInput() *string
	// Experimental.
	SavedownStorageConfiguration() TfKxCluster_SavedownStorageConfigurationPropertyOutputReference
	// Experimental.
	SavedownStorageConfigurationInput() *TfKxCluster_SavedownStorageConfigurationProperty
	// Experimental.
	ScalingGroupConfiguration() TfKxCluster_ScalingGroupConfigurationPropertyOutputReference
	// Experimental.
	ScalingGroupConfigurationInput() *TfKxCluster_ScalingGroupConfigurationProperty
	// Experimental.
	Status() *string
	// Experimental.
	StatusReason() *string
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
	TickerplantLogConfiguration() TfKxCluster_TickerplantLogConfigurationPropertyList
	// Experimental.
	TickerplantLogConfigurationInput() interface{}
	// Experimental.
	Timeouts() TfKxCluster_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
	// Experimental.
	VpcConfiguration() TfKxCluster_VpcConfigurationPropertyOutputReference
	// Experimental.
	VpcConfigurationInput() *TfKxCluster_VpcConfigurationProperty
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
	PutAutoScalingConfiguration(value *TfKxCluster_AutoScalingConfigurationProperty)
	// Experimental.
	PutCacheStorageConfigurations(value interface{})
	// Experimental.
	PutCapacityConfiguration(value *TfKxCluster_CapacityConfigurationProperty)
	// Experimental.
	PutCode(value *TfKxCluster_CodeProperty)
	// Experimental.
	PutDatabase(value interface{})
	// Experimental.
	PutSavedownStorageConfiguration(value *TfKxCluster_SavedownStorageConfigurationProperty)
	// Experimental.
	PutScalingGroupConfiguration(value *TfKxCluster_ScalingGroupConfigurationProperty)
	// Experimental.
	PutTickerplantLogConfiguration(value interface{})
	// Experimental.
	PutTimeouts(value *TfKxCluster_TimeoutsProperty)
	// Experimental.
	PutVpcConfiguration(value *TfKxCluster_VpcConfigurationProperty)
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
	ResetAutoScalingConfiguration()
	// Experimental.
	ResetAvailabilityZoneId()
	// Experimental.
	ResetCacheStorageConfigurations()
	// Experimental.
	ResetCapacityConfiguration()
	// Experimental.
	ResetCode()
	// Experimental.
	ResetCommandLineArguments()
	// Experimental.
	ResetDatabase()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetExecutionRole()
	// Experimental.
	ResetId()
	// Experimental.
	ResetInitializationScript()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSavedownStorageConfiguration()
	// Experimental.
	ResetScalingGroupConfiguration()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTickerplantLogConfiguration()
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

// The jsii proxy struct for TfKxCluster
type jsiiProxy_TfKxCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfKxCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) AutoScalingConfiguration() TfKxCluster_AutoScalingConfigurationPropertyOutputReference {
	var returns TfKxCluster_AutoScalingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"autoScalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) AutoScalingConfigurationInput() *TfKxCluster_AutoScalingConfigurationProperty {
	var returns *TfKxCluster_AutoScalingConfigurationProperty
	_jsii_.Get(
		j,
		"autoScalingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) AvailabilityZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) AzMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) AzModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) CacheStorageConfigurations() TfKxCluster_CacheStorageConfigurationsPropertyList {
	var returns TfKxCluster_CacheStorageConfigurationsPropertyList
	_jsii_.Get(
		j,
		"cacheStorageConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) CacheStorageConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheStorageConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) CapacityConfiguration() TfKxCluster_CapacityConfigurationPropertyOutputReference {
	var returns TfKxCluster_CapacityConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"capacityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) CapacityConfigurationInput() *TfKxCluster_CapacityConfigurationProperty {
	var returns *TfKxCluster_CapacityConfigurationProperty
	_jsii_.Get(
		j,
		"capacityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Code() TfKxCluster_CodePropertyOutputReference {
	var returns TfKxCluster_CodePropertyOutputReference
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) CodeInput() *TfKxCluster_CodeProperty {
	var returns *TfKxCluster_CodeProperty
	_jsii_.Get(
		j,
		"codeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) CommandLineArguments() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"commandLineArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) CommandLineArgumentsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"commandLineArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) CreatedTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Database() TfKxCluster_DatabasePropertyList {
	var returns TfKxCluster_DatabasePropertyList
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) EnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) EnvironmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) InitializationScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initializationScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) InitializationScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initializationScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) LastModifiedTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) ReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) ReleaseLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) SavedownStorageConfiguration() TfKxCluster_SavedownStorageConfigurationPropertyOutputReference {
	var returns TfKxCluster_SavedownStorageConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"savedownStorageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) SavedownStorageConfigurationInput() *TfKxCluster_SavedownStorageConfigurationProperty {
	var returns *TfKxCluster_SavedownStorageConfigurationProperty
	_jsii_.Get(
		j,
		"savedownStorageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) ScalingGroupConfiguration() TfKxCluster_ScalingGroupConfigurationPropertyOutputReference {
	var returns TfKxCluster_ScalingGroupConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"scalingGroupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) ScalingGroupConfigurationInput() *TfKxCluster_ScalingGroupConfigurationProperty {
	var returns *TfKxCluster_ScalingGroupConfigurationProperty
	_jsii_.Get(
		j,
		"scalingGroupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) TickerplantLogConfiguration() TfKxCluster_TickerplantLogConfigurationPropertyList {
	var returns TfKxCluster_TickerplantLogConfigurationPropertyList
	_jsii_.Get(
		j,
		"tickerplantLogConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) TickerplantLogConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tickerplantLogConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Timeouts() TfKxCluster_TimeoutsPropertyOutputReference {
	var returns TfKxCluster_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) VpcConfiguration() TfKxCluster_VpcConfigurationPropertyOutputReference {
	var returns TfKxCluster_VpcConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfKxCluster) VpcConfigurationInput() *TfKxCluster_VpcConfigurationProperty {
	var returns *TfKxCluster_VpcConfigurationProperty
	_jsii_.Get(
		j,
		"vpcConfigurationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster aws_finspace_kx_cluster} Resource.
// Experimental.
func NewTfKxCluster(scope constructs.Construct, id *string, config *TfKxClusterConfig) TfKxCluster {
	_init_.Initialize()

	if err := validateNewTfKxClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfKxCluster{}

	_jsii_.Create(
		"@cdktn/aws-finspace.TfKxCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster aws_finspace_kx_cluster} Resource.
// Experimental.
func NewTfKxCluster_Override(t TfKxCluster, scope constructs.Construct, id *string, config *TfKxClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-finspace.TfKxCluster",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfKxCluster)SetAvailabilityZoneId(val *string) {
	if err := j.validateSetAvailabilityZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneId",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetAzMode(val *string) {
	if err := j.validateSetAzModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azMode",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetCommandLineArguments(val *map[string]*string) {
	if err := j.validateSetCommandLineArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commandLineArguments",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetEnvironmentId(val *string) {
	if err := j.validateSetEnvironmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentId",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetInitializationScript(val *string) {
	if err := j.validateSetInitializationScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initializationScript",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetReleaseLabel(val *string) {
	if err := j.validateSetReleaseLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfKxCluster)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTN code for importing a TfKxCluster resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfKxCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfKxCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-finspace.TfKxCluster",
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
func TfKxCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfKxCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-finspace.TfKxCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfKxCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfKxCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-finspace.TfKxCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfKxCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfKxCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-finspace.TfKxCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfKxCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-finspace.TfKxCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfKxCluster) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfKxCluster) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfKxCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfKxCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfKxCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfKxCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfKxCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfKxCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfKxCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfKxCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfKxCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfKxCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfKxCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfKxCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfKxCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfKxCluster) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfKxCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfKxCluster) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfKxCluster) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfKxCluster) PutAutoScalingConfiguration(value *TfKxCluster_AutoScalingConfigurationProperty) {
	if err := t.validatePutAutoScalingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAutoScalingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKxCluster) PutCacheStorageConfigurations(value interface{}) {
	if err := t.validatePutCacheStorageConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCacheStorageConfigurations",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKxCluster) PutCapacityConfiguration(value *TfKxCluster_CapacityConfigurationProperty) {
	if err := t.validatePutCapacityConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCapacityConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKxCluster) PutCode(value *TfKxCluster_CodeProperty) {
	if err := t.validatePutCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCode",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKxCluster) PutDatabase(value interface{}) {
	if err := t.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDatabase",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKxCluster) PutSavedownStorageConfiguration(value *TfKxCluster_SavedownStorageConfigurationProperty) {
	if err := t.validatePutSavedownStorageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSavedownStorageConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKxCluster) PutScalingGroupConfiguration(value *TfKxCluster_ScalingGroupConfigurationProperty) {
	if err := t.validatePutScalingGroupConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putScalingGroupConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKxCluster) PutTickerplantLogConfiguration(value interface{}) {
	if err := t.validatePutTickerplantLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTickerplantLogConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKxCluster) PutTimeouts(value *TfKxCluster_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKxCluster) PutVpcConfiguration(value *TfKxCluster_VpcConfigurationProperty) {
	if err := t.validatePutVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVpcConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfKxCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfKxCluster) ResetAutoScalingConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoScalingConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetAvailabilityZoneId() {
	_jsii_.InvokeVoid(
		t,
		"resetAvailabilityZoneId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetCacheStorageConfigurations() {
	_jsii_.InvokeVoid(
		t,
		"resetCacheStorageConfigurations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetCapacityConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetCapacityConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetCode() {
	_jsii_.InvokeVoid(
		t,
		"resetCode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetCommandLineArguments() {
	_jsii_.InvokeVoid(
		t,
		"resetCommandLineArguments",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetDatabase() {
	_jsii_.InvokeVoid(
		t,
		"resetDatabase",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetExecutionRole() {
	_jsii_.InvokeVoid(
		t,
		"resetExecutionRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetInitializationScript() {
	_jsii_.InvokeVoid(
		t,
		"resetInitializationScript",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetSavedownStorageConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSavedownStorageConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetScalingGroupConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetScalingGroupConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetTickerplantLogConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetTickerplantLogConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfKxCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfKxCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfKxCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfKxCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfKxCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfKxCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfKxCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

