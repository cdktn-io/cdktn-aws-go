package awsemrserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemrserverless/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsemrserverless/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application aws_emrserverless_application}.
// Experimental.
type AwsEmrserverlessApplication interface {
	cdktn.TerraformResource
	// Experimental.
	Architecture() *string
	// Experimental.
	SetArchitecture(val *string)
	// Experimental.
	ArchitectureInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	AutoStartConfiguration() AwsEmrserverlessApplication_AutoStartConfigurationPropertyOutputReference
	// Experimental.
	AutoStartConfigurationInput() *AwsEmrserverlessApplication_AutoStartConfigurationProperty
	// Experimental.
	AutoStopConfiguration() AwsEmrserverlessApplication_AutoStopConfigurationPropertyOutputReference
	// Experimental.
	AutoStopConfigurationInput() *AwsEmrserverlessApplication_AutoStopConfigurationProperty
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
	ImageConfiguration() AwsEmrserverlessApplication_ImageConfigurationPropertyOutputReference
	// Experimental.
	ImageConfigurationInput() *AwsEmrserverlessApplication_ImageConfigurationProperty
	// Experimental.
	InitialCapacity() AwsEmrserverlessApplication_InitialCapacityPropertyList
	// Experimental.
	InitialCapacityInput() interface{}
	// Experimental.
	InteractiveConfiguration() AwsEmrserverlessApplication_InteractiveConfigurationPropertyOutputReference
	// Experimental.
	InteractiveConfigurationInput() *AwsEmrserverlessApplication_InteractiveConfigurationProperty
	// Experimental.
	JobLevelCostAllocationConfiguration() AwsEmrserverlessApplication_JobLevelCostAllocationConfigurationPropertyOutputReference
	// Experimental.
	JobLevelCostAllocationConfigurationInput() *AwsEmrserverlessApplication_JobLevelCostAllocationConfigurationProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaximumCapacity() AwsEmrserverlessApplication_MaximumCapacityPropertyOutputReference
	// Experimental.
	MaximumCapacityInput() *AwsEmrserverlessApplication_MaximumCapacityProperty
	// Experimental.
	MonitoringConfiguration() AwsEmrserverlessApplication_MonitoringConfigurationPropertyOutputReference
	// Experimental.
	MonitoringConfigurationInput() *AwsEmrserverlessApplication_MonitoringConfigurationProperty
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NetworkConfiguration() AwsEmrserverlessApplication_NetworkConfigurationPropertyOutputReference
	// Experimental.
	NetworkConfigurationInput() *AwsEmrserverlessApplication_NetworkConfigurationProperty
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
	RuntimeConfiguration() AwsEmrserverlessApplication_RuntimeConfigurationPropertyList
	// Experimental.
	RuntimeConfigurationInput() interface{}
	// Experimental.
	SchedulerConfiguration() AwsEmrserverlessApplication_SchedulerConfigurationPropertyOutputReference
	// Experimental.
	SchedulerConfigurationInput() *AwsEmrserverlessApplication_SchedulerConfigurationProperty
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
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	PutAutoStartConfiguration(value *AwsEmrserverlessApplication_AutoStartConfigurationProperty)
	// Experimental.
	PutAutoStopConfiguration(value *AwsEmrserverlessApplication_AutoStopConfigurationProperty)
	// Experimental.
	PutImageConfiguration(value *AwsEmrserverlessApplication_ImageConfigurationProperty)
	// Experimental.
	PutInitialCapacity(value interface{})
	// Experimental.
	PutInteractiveConfiguration(value *AwsEmrserverlessApplication_InteractiveConfigurationProperty)
	// Experimental.
	PutJobLevelCostAllocationConfiguration(value *AwsEmrserverlessApplication_JobLevelCostAllocationConfigurationProperty)
	// Experimental.
	PutMaximumCapacity(value *AwsEmrserverlessApplication_MaximumCapacityProperty)
	// Experimental.
	PutMonitoringConfiguration(value *AwsEmrserverlessApplication_MonitoringConfigurationProperty)
	// Experimental.
	PutNetworkConfiguration(value *AwsEmrserverlessApplication_NetworkConfigurationProperty)
	// Experimental.
	PutRuntimeConfiguration(value interface{})
	// Experimental.
	PutSchedulerConfiguration(value *AwsEmrserverlessApplication_SchedulerConfigurationProperty)
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
	ResetArchitecture()
	// Experimental.
	ResetAutoStartConfiguration()
	// Experimental.
	ResetAutoStopConfiguration()
	// Experimental.
	ResetId()
	// Experimental.
	ResetImageConfiguration()
	// Experimental.
	ResetInitialCapacity()
	// Experimental.
	ResetInteractiveConfiguration()
	// Experimental.
	ResetJobLevelCostAllocationConfiguration()
	// Experimental.
	ResetMaximumCapacity()
	// Experimental.
	ResetMonitoringConfiguration()
	// Experimental.
	ResetNetworkConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRuntimeConfiguration()
	// Experimental.
	ResetSchedulerConfiguration()
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

// The jsii proxy struct for AwsEmrserverlessApplication
type jsiiProxy_AwsEmrserverlessApplication struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsEmrserverlessApplication) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) ArchitectureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) AutoStartConfiguration() AwsEmrserverlessApplication_AutoStartConfigurationPropertyOutputReference {
	var returns AwsEmrserverlessApplication_AutoStartConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"autoStartConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) AutoStartConfigurationInput() *AwsEmrserverlessApplication_AutoStartConfigurationProperty {
	var returns *AwsEmrserverlessApplication_AutoStartConfigurationProperty
	_jsii_.Get(
		j,
		"autoStartConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) AutoStopConfiguration() AwsEmrserverlessApplication_AutoStopConfigurationPropertyOutputReference {
	var returns AwsEmrserverlessApplication_AutoStopConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"autoStopConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) AutoStopConfigurationInput() *AwsEmrserverlessApplication_AutoStopConfigurationProperty {
	var returns *AwsEmrserverlessApplication_AutoStopConfigurationProperty
	_jsii_.Get(
		j,
		"autoStopConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) ImageConfiguration() AwsEmrserverlessApplication_ImageConfigurationPropertyOutputReference {
	var returns AwsEmrserverlessApplication_ImageConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"imageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) ImageConfigurationInput() *AwsEmrserverlessApplication_ImageConfigurationProperty {
	var returns *AwsEmrserverlessApplication_ImageConfigurationProperty
	_jsii_.Get(
		j,
		"imageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) InitialCapacity() AwsEmrserverlessApplication_InitialCapacityPropertyList {
	var returns AwsEmrserverlessApplication_InitialCapacityPropertyList
	_jsii_.Get(
		j,
		"initialCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) InitialCapacityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) InteractiveConfiguration() AwsEmrserverlessApplication_InteractiveConfigurationPropertyOutputReference {
	var returns AwsEmrserverlessApplication_InteractiveConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"interactiveConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) InteractiveConfigurationInput() *AwsEmrserverlessApplication_InteractiveConfigurationProperty {
	var returns *AwsEmrserverlessApplication_InteractiveConfigurationProperty
	_jsii_.Get(
		j,
		"interactiveConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) JobLevelCostAllocationConfiguration() AwsEmrserverlessApplication_JobLevelCostAllocationConfigurationPropertyOutputReference {
	var returns AwsEmrserverlessApplication_JobLevelCostAllocationConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"jobLevelCostAllocationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) JobLevelCostAllocationConfigurationInput() *AwsEmrserverlessApplication_JobLevelCostAllocationConfigurationProperty {
	var returns *AwsEmrserverlessApplication_JobLevelCostAllocationConfigurationProperty
	_jsii_.Get(
		j,
		"jobLevelCostAllocationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) MaximumCapacity() AwsEmrserverlessApplication_MaximumCapacityPropertyOutputReference {
	var returns AwsEmrserverlessApplication_MaximumCapacityPropertyOutputReference
	_jsii_.Get(
		j,
		"maximumCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) MaximumCapacityInput() *AwsEmrserverlessApplication_MaximumCapacityProperty {
	var returns *AwsEmrserverlessApplication_MaximumCapacityProperty
	_jsii_.Get(
		j,
		"maximumCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) MonitoringConfiguration() AwsEmrserverlessApplication_MonitoringConfigurationPropertyOutputReference {
	var returns AwsEmrserverlessApplication_MonitoringConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) MonitoringConfigurationInput() *AwsEmrserverlessApplication_MonitoringConfigurationProperty {
	var returns *AwsEmrserverlessApplication_MonitoringConfigurationProperty
	_jsii_.Get(
		j,
		"monitoringConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) NetworkConfiguration() AwsEmrserverlessApplication_NetworkConfigurationPropertyOutputReference {
	var returns AwsEmrserverlessApplication_NetworkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) NetworkConfigurationInput() *AwsEmrserverlessApplication_NetworkConfigurationProperty {
	var returns *AwsEmrserverlessApplication_NetworkConfigurationProperty
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) ReleaseLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) ReleaseLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) RuntimeConfiguration() AwsEmrserverlessApplication_RuntimeConfigurationPropertyList {
	var returns AwsEmrserverlessApplication_RuntimeConfigurationPropertyList
	_jsii_.Get(
		j,
		"runtimeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) RuntimeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) SchedulerConfiguration() AwsEmrserverlessApplication_SchedulerConfigurationPropertyOutputReference {
	var returns AwsEmrserverlessApplication_SchedulerConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"schedulerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) SchedulerConfigurationInput() *AwsEmrserverlessApplication_SchedulerConfigurationProperty {
	var returns *AwsEmrserverlessApplication_SchedulerConfigurationProperty
	_jsii_.Get(
		j,
		"schedulerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrserverlessApplication) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application aws_emrserverless_application} Resource.
// Experimental.
func NewAwsEmrserverlessApplication(scope constructs.Construct, id *string, config *AwsEmrserverlessApplicationConfig) AwsEmrserverlessApplication {
	_init_.Initialize()

	if err := validateNewAwsEmrserverlessApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEmrserverlessApplication{}

	_jsii_.Create(
		"@cdktn/aws-emr-serverless.AwsEmrserverlessApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application aws_emrserverless_application} Resource.
// Experimental.
func NewAwsEmrserverlessApplication_Override(a AwsEmrserverlessApplication, scope constructs.Construct, id *string, config *AwsEmrserverlessApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr-serverless.AwsEmrserverlessApplication",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsEmrserverlessApplication)SetArchitecture(val *string) {
	if err := j.validateSetArchitectureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_AwsEmrserverlessApplication)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsEmrserverlessApplication)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsEmrserverlessApplication)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsEmrserverlessApplication)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsEmrserverlessApplication)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsEmrserverlessApplication)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsEmrserverlessApplication)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsEmrserverlessApplication)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsEmrserverlessApplication)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsEmrserverlessApplication)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsEmrserverlessApplication)SetReleaseLabel(val *string) {
	if err := j.validateSetReleaseLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseLabel",
		val,
	)
}

func (j *jsiiProxy_AwsEmrserverlessApplication)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsEmrserverlessApplication)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsEmrserverlessApplication)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTN code for importing a AwsEmrserverlessApplication resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsEmrserverlessApplication_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsEmrserverlessApplication_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-emr-serverless.AwsEmrserverlessApplication",
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
func AwsEmrserverlessApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEmrserverlessApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-emr-serverless.AwsEmrserverlessApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsEmrserverlessApplication_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEmrserverlessApplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-emr-serverless.AwsEmrserverlessApplication",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsEmrserverlessApplication_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsEmrserverlessApplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-emr-serverless.AwsEmrserverlessApplication",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsEmrserverlessApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-emr-serverless.AwsEmrserverlessApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsEmrserverlessApplication) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEmrserverlessApplication) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmrserverlessApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEmrserverlessApplication) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEmrserverlessApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEmrserverlessApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEmrserverlessApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEmrserverlessApplication) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEmrserverlessApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEmrserverlessApplication) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmrserverlessApplication) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsEmrserverlessApplication) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) PutAutoStartConfiguration(value *AwsEmrserverlessApplication_AutoStartConfigurationProperty) {
	if err := a.validatePutAutoStartConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutoStartConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) PutAutoStopConfiguration(value *AwsEmrserverlessApplication_AutoStopConfigurationProperty) {
	if err := a.validatePutAutoStopConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutoStopConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) PutImageConfiguration(value *AwsEmrserverlessApplication_ImageConfigurationProperty) {
	if err := a.validatePutImageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putImageConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) PutInitialCapacity(value interface{}) {
	if err := a.validatePutInitialCapacityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInitialCapacity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) PutInteractiveConfiguration(value *AwsEmrserverlessApplication_InteractiveConfigurationProperty) {
	if err := a.validatePutInteractiveConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInteractiveConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) PutJobLevelCostAllocationConfiguration(value *AwsEmrserverlessApplication_JobLevelCostAllocationConfigurationProperty) {
	if err := a.validatePutJobLevelCostAllocationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJobLevelCostAllocationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) PutMaximumCapacity(value *AwsEmrserverlessApplication_MaximumCapacityProperty) {
	if err := a.validatePutMaximumCapacityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMaximumCapacity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) PutMonitoringConfiguration(value *AwsEmrserverlessApplication_MonitoringConfigurationProperty) {
	if err := a.validatePutMonitoringConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonitoringConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) PutNetworkConfiguration(value *AwsEmrserverlessApplication_NetworkConfigurationProperty) {
	if err := a.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) PutRuntimeConfiguration(value interface{}) {
	if err := a.validatePutRuntimeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRuntimeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) PutSchedulerConfiguration(value *AwsEmrserverlessApplication_SchedulerConfigurationProperty) {
	if err := a.validatePutSchedulerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchedulerConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetArchitecture() {
	_jsii_.InvokeVoid(
		a,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetAutoStartConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoStartConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetAutoStopConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoStopConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetImageConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetImageConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetInitialCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetInitialCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetInteractiveConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetInteractiveConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetJobLevelCostAllocationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetJobLevelCostAllocationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetMaximumCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetMonitoringConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitoringConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetRuntimeConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntimeConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetSchedulerConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSchedulerConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrserverlessApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrserverlessApplication) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrserverlessApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrserverlessApplication) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

