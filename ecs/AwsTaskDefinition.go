package ecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ecs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/ecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition aws_ecs_task_definition}.
// Experimental.
type AwsTaskDefinition interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	ArnWithoutRevision() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	ContainerDefinitions() *string
	// Experimental.
	SetContainerDefinitions(val *string)
	// Experimental.
	ContainerDefinitionsInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	Cpu() *string
	// Experimental.
	SetCpu(val *string)
	// Experimental.
	CpuInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EnableFaultInjection() interface{}
	// Experimental.
	SetEnableFaultInjection(val interface{})
	// Experimental.
	EnableFaultInjectionInput() interface{}
	// Experimental.
	EphemeralStorage() AwsTaskDefinition_EphemeralStoragePropertyOutputReference
	// Experimental.
	EphemeralStorageInput() *AwsTaskDefinition_EphemeralStorageProperty
	// Experimental.
	ExecutionRoleArn() *string
	// Experimental.
	SetExecutionRoleArn(val *string)
	// Experimental.
	ExecutionRoleArnInput() *string
	// Experimental.
	Family() *string
	// Experimental.
	SetFamily(val *string)
	// Experimental.
	FamilyInput() *string
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
	IpcMode() *string
	// Experimental.
	SetIpcMode(val *string)
	// Experimental.
	IpcModeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	Memory() *string
	// Experimental.
	SetMemory(val *string)
	// Experimental.
	MemoryInput() *string
	// Experimental.
	NetworkMode() *string
	// Experimental.
	SetNetworkMode(val *string)
	// Experimental.
	NetworkModeInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PidMode() *string
	// Experimental.
	SetPidMode(val *string)
	// Experimental.
	PidModeInput() *string
	// Experimental.
	PlacementConstraints() AwsTaskDefinition_PlacementConstraintsPropertyList
	// Experimental.
	PlacementConstraintsInput() interface{}
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	ProxyConfiguration() AwsTaskDefinition_ProxyConfigurationPropertyOutputReference
	// Experimental.
	ProxyConfigurationInput() *AwsTaskDefinition_ProxyConfigurationProperty
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	RequiresCompatibilities() *[]*string
	// Experimental.
	SetRequiresCompatibilities(val *[]*string)
	// Experimental.
	RequiresCompatibilitiesInput() *[]*string
	// Experimental.
	Revision() *float64
	// Experimental.
	RuntimePlatform() AwsTaskDefinition_RuntimePlatformPropertyOutputReference
	// Experimental.
	RuntimePlatformInput() *AwsTaskDefinition_RuntimePlatformProperty
	// Experimental.
	SkipDestroy() interface{}
	// Experimental.
	SetSkipDestroy(val interface{})
	// Experimental.
	SkipDestroyInput() interface{}
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
	TaskRoleArn() *string
	// Experimental.
	SetTaskRoleArn(val *string)
	// Experimental.
	TaskRoleArnInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	TrackLatest() interface{}
	// Experimental.
	SetTrackLatest(val interface{})
	// Experimental.
	TrackLatestInput() interface{}
	// Experimental.
	Volume() AwsTaskDefinition_VolumePropertyList
	// Experimental.
	VolumeInput() interface{}
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
	PutEphemeralStorage(value *AwsTaskDefinition_EphemeralStorageProperty)
	// Experimental.
	PutPlacementConstraints(value interface{})
	// Experimental.
	PutProxyConfiguration(value *AwsTaskDefinition_ProxyConfigurationProperty)
	// Experimental.
	PutRuntimePlatform(value *AwsTaskDefinition_RuntimePlatformProperty)
	// Experimental.
	PutVolume(value interface{})
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
	ResetCpu()
	// Experimental.
	ResetEnableFaultInjection()
	// Experimental.
	ResetEphemeralStorage()
	// Experimental.
	ResetExecutionRoleArn()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIpcMode()
	// Experimental.
	ResetMemory()
	// Experimental.
	ResetNetworkMode()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPidMode()
	// Experimental.
	ResetPlacementConstraints()
	// Experimental.
	ResetProxyConfiguration()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRequiresCompatibilities()
	// Experimental.
	ResetRuntimePlatform()
	// Experimental.
	ResetSkipDestroy()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTaskRoleArn()
	// Experimental.
	ResetTrackLatest()
	// Experimental.
	ResetVolume()
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

// The jsii proxy struct for AwsTaskDefinition
type jsiiProxy_AwsTaskDefinition struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsTaskDefinition) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) ArnWithoutRevision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnWithoutRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) ContainerDefinitions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) ContainerDefinitionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) Cpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) CpuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) EnableFaultInjection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFaultInjection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) EnableFaultInjectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableFaultInjectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) EphemeralStorage() AwsTaskDefinition_EphemeralStoragePropertyOutputReference {
	var returns AwsTaskDefinition_EphemeralStoragePropertyOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) EphemeralStorageInput() *AwsTaskDefinition_EphemeralStorageProperty {
	var returns *AwsTaskDefinition_EphemeralStorageProperty
	_jsii_.Get(
		j,
		"ephemeralStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) FamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"familyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) IpcMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipcMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) IpcModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipcModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) Memory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) MemoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) NetworkMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) NetworkModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) PidMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pidMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) PidModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pidModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) PlacementConstraints() AwsTaskDefinition_PlacementConstraintsPropertyList {
	var returns AwsTaskDefinition_PlacementConstraintsPropertyList
	_jsii_.Get(
		j,
		"placementConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) PlacementConstraintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) ProxyConfiguration() AwsTaskDefinition_ProxyConfigurationPropertyOutputReference {
	var returns AwsTaskDefinition_ProxyConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"proxyConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) ProxyConfigurationInput() *AwsTaskDefinition_ProxyConfigurationProperty {
	var returns *AwsTaskDefinition_ProxyConfigurationProperty
	_jsii_.Get(
		j,
		"proxyConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) RequiresCompatibilities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiresCompatibilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) RequiresCompatibilitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requiresCompatibilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) Revision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"revision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) RuntimePlatform() AwsTaskDefinition_RuntimePlatformPropertyOutputReference {
	var returns AwsTaskDefinition_RuntimePlatformPropertyOutputReference
	_jsii_.Get(
		j,
		"runtimePlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) RuntimePlatformInput() *AwsTaskDefinition_RuntimePlatformProperty {
	var returns *AwsTaskDefinition_RuntimePlatformProperty
	_jsii_.Get(
		j,
		"runtimePlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) SkipDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) SkipDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) TaskRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) TaskRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) TrackLatest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trackLatest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) TrackLatestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trackLatestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) Volume() AwsTaskDefinition_VolumePropertyList {
	var returns AwsTaskDefinition_VolumePropertyList
	_jsii_.Get(
		j,
		"volume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTaskDefinition) VolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition aws_ecs_task_definition} Resource.
// Experimental.
func NewAwsTaskDefinition(scope constructs.Construct, id *string, config *AwsTaskDefinitionConfig) AwsTaskDefinition {
	_init_.Initialize()

	if err := validateNewAwsTaskDefinitionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTaskDefinition{}

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsTaskDefinition",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition aws_ecs_task_definition} Resource.
// Experimental.
func NewAwsTaskDefinition_Override(a AwsTaskDefinition, scope constructs.Construct, id *string, config *AwsTaskDefinitionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsTaskDefinition",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetContainerDefinitions(val *string) {
	if err := j.validateSetContainerDefinitionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerDefinitions",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetCpu(val *string) {
	if err := j.validateSetCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpu",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetEnableFaultInjection(val interface{}) {
	if err := j.validateSetEnableFaultInjectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableFaultInjection",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetFamily(val *string) {
	if err := j.validateSetFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"family",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetIpcMode(val *string) {
	if err := j.validateSetIpcModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipcMode",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetMemory(val *string) {
	if err := j.validateSetMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memory",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetNetworkMode(val *string) {
	if err := j.validateSetNetworkModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkMode",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetPidMode(val *string) {
	if err := j.validateSetPidModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pidMode",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetRequiresCompatibilities(val *[]*string) {
	if err := j.validateSetRequiresCompatibilitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiresCompatibilities",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetSkipDestroy(val interface{}) {
	if err := j.validateSetSkipDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipDestroy",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetTaskRoleArn(val *string) {
	if err := j.validateSetTaskRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsTaskDefinition)SetTrackLatest(val interface{}) {
	if err := j.validateSetTrackLatestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trackLatest",
		val,
	)
}

// Generates CDKTN code for importing a AwsTaskDefinition resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsTaskDefinition_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsTaskDefinition_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ecs.AwsTaskDefinition",
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
func AwsTaskDefinition_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsTaskDefinition_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ecs.AwsTaskDefinition",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsTaskDefinition_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsTaskDefinition_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ecs.AwsTaskDefinition",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsTaskDefinition_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsTaskDefinition_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ecs.AwsTaskDefinition",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsTaskDefinition_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ecs.AwsTaskDefinition",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsTaskDefinition) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsTaskDefinition) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsTaskDefinition) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTaskDefinition) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTaskDefinition) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTaskDefinition) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTaskDefinition) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTaskDefinition) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTaskDefinition) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTaskDefinition) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTaskDefinition) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTaskDefinition) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTaskDefinition) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsTaskDefinition) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTaskDefinition) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsTaskDefinition) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsTaskDefinition) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsTaskDefinition) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsTaskDefinition) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsTaskDefinition) PutEphemeralStorage(value *AwsTaskDefinition_EphemeralStorageProperty) {
	if err := a.validatePutEphemeralStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEphemeralStorage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTaskDefinition) PutPlacementConstraints(value interface{}) {
	if err := a.validatePutPlacementConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPlacementConstraints",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTaskDefinition) PutProxyConfiguration(value *AwsTaskDefinition_ProxyConfigurationProperty) {
	if err := a.validatePutProxyConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProxyConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTaskDefinition) PutRuntimePlatform(value *AwsTaskDefinition_RuntimePlatformProperty) {
	if err := a.validatePutRuntimePlatformParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRuntimePlatform",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTaskDefinition) PutVolume(value interface{}) {
	if err := a.validatePutVolumeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVolume",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTaskDefinition) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetCpu() {
	_jsii_.InvokeVoid(
		a,
		"resetCpu",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetEnableFaultInjection() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableFaultInjection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetEphemeralStorage() {
	_jsii_.InvokeVoid(
		a,
		"resetEphemeralStorage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetExecutionRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetIpcMode() {
	_jsii_.InvokeVoid(
		a,
		"resetIpcMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetMemory() {
	_jsii_.InvokeVoid(
		a,
		"resetMemory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetNetworkMode() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetPidMode() {
	_jsii_.InvokeVoid(
		a,
		"resetPidMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetPlacementConstraints() {
	_jsii_.InvokeVoid(
		a,
		"resetPlacementConstraints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetProxyConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetProxyConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetRequiresCompatibilities() {
	_jsii_.InvokeVoid(
		a,
		"resetRequiresCompatibilities",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetRuntimePlatform() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntimePlatform",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetSkipDestroy() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipDestroy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetTaskRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetTaskRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetTrackLatest() {
	_jsii_.InvokeVoid(
		a,
		"resetTrackLatest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) ResetVolume() {
	_jsii_.InvokeVoid(
		a,
		"resetVolume",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTaskDefinition) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTaskDefinition) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTaskDefinition) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTaskDefinition) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTaskDefinition) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTaskDefinition) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTaskDefinition) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

