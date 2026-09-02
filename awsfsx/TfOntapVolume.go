package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume aws_fsx_ontap_volume}.
// Experimental.
type TfOntapVolume interface {
	cdktn.TerraformResource
	// Experimental.
	AggregateConfiguration() TfOntapVolume_AggregateConfigurationPropertyOutputReference
	// Experimental.
	AggregateConfigurationInput() *TfOntapVolume_AggregateConfigurationProperty
	// Experimental.
	Arn() *string
	// Experimental.
	BypassSnaplockEnterpriseRetention() interface{}
	// Experimental.
	SetBypassSnaplockEnterpriseRetention(val interface{})
	// Experimental.
	BypassSnaplockEnterpriseRetentionInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	CopyTagsToBackups() interface{}
	// Experimental.
	SetCopyTagsToBackups(val interface{})
	// Experimental.
	CopyTagsToBackupsInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	FileSystemId() *string
	// Experimental.
	FinalBackupTags() *map[string]*string
	// Experimental.
	SetFinalBackupTags(val *map[string]*string)
	// Experimental.
	FinalBackupTagsInput() *map[string]*string
	// Experimental.
	FlexcacheEndpointType() *string
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
	JunctionPath() *string
	// Experimental.
	SetJunctionPath(val *string)
	// Experimental.
	JunctionPathInput() *string
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
	OntapVolumeType() *string
	// Experimental.
	SetOntapVolumeType(val *string)
	// Experimental.
	OntapVolumeTypeInput() *string
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
	SecurityStyle() *string
	// Experimental.
	SetSecurityStyle(val *string)
	// Experimental.
	SecurityStyleInput() *string
	// Experimental.
	SizeInBytes() *string
	// Experimental.
	SetSizeInBytes(val *string)
	// Experimental.
	SizeInBytesInput() *string
	// Experimental.
	SizeInMegabytes() *float64
	// Experimental.
	SetSizeInMegabytes(val *float64)
	// Experimental.
	SizeInMegabytesInput() *float64
	// Experimental.
	SkipFinalBackup() interface{}
	// Experimental.
	SetSkipFinalBackup(val interface{})
	// Experimental.
	SkipFinalBackupInput() interface{}
	// Experimental.
	SnaplockConfiguration() TfOntapVolume_SnaplockConfigurationPropertyOutputReference
	// Experimental.
	SnaplockConfigurationInput() *TfOntapVolume_SnaplockConfigurationProperty
	// Experimental.
	SnapshotPolicy() *string
	// Experimental.
	SetSnapshotPolicy(val *string)
	// Experimental.
	SnapshotPolicyInput() *string
	// Experimental.
	StorageEfficiencyEnabled() interface{}
	// Experimental.
	SetStorageEfficiencyEnabled(val interface{})
	// Experimental.
	StorageEfficiencyEnabledInput() interface{}
	// Experimental.
	StorageVirtualMachineId() *string
	// Experimental.
	SetStorageVirtualMachineId(val *string)
	// Experimental.
	StorageVirtualMachineIdInput() *string
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
	TieringPolicy() TfOntapVolume_TieringPolicyPropertyOutputReference
	// Experimental.
	TieringPolicyInput() *TfOntapVolume_TieringPolicyProperty
	// Experimental.
	Timeouts() TfOntapVolume_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Uuid() *string
	// Experimental.
	VolumeStyle() *string
	// Experimental.
	SetVolumeStyle(val *string)
	// Experimental.
	VolumeStyleInput() *string
	// Experimental.
	VolumeType() *string
	// Experimental.
	SetVolumeType(val *string)
	// Experimental.
	VolumeTypeInput() *string
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
	PutAggregateConfiguration(value *TfOntapVolume_AggregateConfigurationProperty)
	// Experimental.
	PutSnaplockConfiguration(value *TfOntapVolume_SnaplockConfigurationProperty)
	// Experimental.
	PutTieringPolicy(value *TfOntapVolume_TieringPolicyProperty)
	// Experimental.
	PutTimeouts(value *TfOntapVolume_TimeoutsProperty)
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
	ResetAggregateConfiguration()
	// Experimental.
	ResetBypassSnaplockEnterpriseRetention()
	// Experimental.
	ResetCopyTagsToBackups()
	// Experimental.
	ResetFinalBackupTags()
	// Experimental.
	ResetId()
	// Experimental.
	ResetJunctionPath()
	// Experimental.
	ResetOntapVolumeType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSecurityStyle()
	// Experimental.
	ResetSizeInBytes()
	// Experimental.
	ResetSizeInMegabytes()
	// Experimental.
	ResetSkipFinalBackup()
	// Experimental.
	ResetSnaplockConfiguration()
	// Experimental.
	ResetSnapshotPolicy()
	// Experimental.
	ResetStorageEfficiencyEnabled()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTieringPolicy()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetVolumeStyle()
	// Experimental.
	ResetVolumeType()
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

// The jsii proxy struct for TfOntapVolume
type jsiiProxy_TfOntapVolume struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfOntapVolume) AggregateConfiguration() TfOntapVolume_AggregateConfigurationPropertyOutputReference {
	var returns TfOntapVolume_AggregateConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"aggregateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) AggregateConfigurationInput() *TfOntapVolume_AggregateConfigurationProperty {
	var returns *TfOntapVolume_AggregateConfigurationProperty
	_jsii_.Get(
		j,
		"aggregateConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) BypassSnaplockEnterpriseRetention() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassSnaplockEnterpriseRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) BypassSnaplockEnterpriseRetentionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassSnaplockEnterpriseRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) CopyTagsToBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) CopyTagsToBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) FinalBackupTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"finalBackupTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) FinalBackupTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"finalBackupTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) FlexcacheEndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flexcacheEndpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) JunctionPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"junctionPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) JunctionPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"junctionPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) OntapVolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ontapVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) OntapVolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ontapVolumeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) SecurityStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) SecurityStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) SizeInBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) SizeInBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) SizeInMegabytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInMegabytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) SizeInMegabytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInMegabytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) SkipFinalBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) SkipFinalBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) SnaplockConfiguration() TfOntapVolume_SnaplockConfigurationPropertyOutputReference {
	var returns TfOntapVolume_SnaplockConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"snaplockConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) SnaplockConfigurationInput() *TfOntapVolume_SnaplockConfigurationProperty {
	var returns *TfOntapVolume_SnaplockConfigurationProperty
	_jsii_.Get(
		j,
		"snaplockConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) SnapshotPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) SnapshotPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) StorageEfficiencyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEfficiencyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) StorageEfficiencyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEfficiencyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) StorageVirtualMachineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageVirtualMachineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) StorageVirtualMachineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageVirtualMachineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) TieringPolicy() TfOntapVolume_TieringPolicyPropertyOutputReference {
	var returns TfOntapVolume_TieringPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"tieringPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) TieringPolicyInput() *TfOntapVolume_TieringPolicyProperty {
	var returns *TfOntapVolume_TieringPolicyProperty
	_jsii_.Get(
		j,
		"tieringPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) Timeouts() TfOntapVolume_TimeoutsPropertyOutputReference {
	var returns TfOntapVolume_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) VolumeStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) VolumeStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume) VolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume aws_fsx_ontap_volume} Resource.
// Experimental.
func NewTfOntapVolume(scope constructs.Construct, id *string, config *TfOntapVolumeConfig) TfOntapVolume {
	_init_.Initialize()

	if err := validateNewTfOntapVolumeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfOntapVolume{}

	_jsii_.Create(
		"@cdktn/aws-fsx.TfOntapVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume aws_fsx_ontap_volume} Resource.
// Experimental.
func NewTfOntapVolume_Override(t TfOntapVolume, scope constructs.Construct, id *string, config *TfOntapVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.TfOntapVolume",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetBypassSnaplockEnterpriseRetention(val interface{}) {
	if err := j.validateSetBypassSnaplockEnterpriseRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bypassSnaplockEnterpriseRetention",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetCopyTagsToBackups(val interface{}) {
	if err := j.validateSetCopyTagsToBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToBackups",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetFinalBackupTags(val *map[string]*string) {
	if err := j.validateSetFinalBackupTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalBackupTags",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetJunctionPath(val *string) {
	if err := j.validateSetJunctionPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"junctionPath",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetOntapVolumeType(val *string) {
	if err := j.validateSetOntapVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ontapVolumeType",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetSecurityStyle(val *string) {
	if err := j.validateSetSecurityStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityStyle",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetSizeInBytes(val *string) {
	if err := j.validateSetSizeInBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeInBytes",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetSizeInMegabytes(val *float64) {
	if err := j.validateSetSizeInMegabytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeInMegabytes",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetSkipFinalBackup(val interface{}) {
	if err := j.validateSetSkipFinalBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipFinalBackup",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetSnapshotPolicy(val *string) {
	if err := j.validateSetSnapshotPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotPolicy",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetStorageEfficiencyEnabled(val interface{}) {
	if err := j.validateSetStorageEfficiencyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEfficiencyEnabled",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetStorageVirtualMachineId(val *string) {
	if err := j.validateSetStorageVirtualMachineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageVirtualMachineId",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetVolumeStyle(val *string) {
	if err := j.validateSetVolumeStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeStyle",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume)SetVolumeType(val *string) {
	if err := j.validateSetVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeType",
		val,
	)
}

// Generates CDKTN code for importing a TfOntapVolume resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfOntapVolume_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfOntapVolume_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.TfOntapVolume",
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
func TfOntapVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfOntapVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.TfOntapVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfOntapVolume_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfOntapVolume_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.TfOntapVolume",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfOntapVolume_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfOntapVolume_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.TfOntapVolume",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfOntapVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-fsx.TfOntapVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfOntapVolume) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfOntapVolume) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfOntapVolume) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfOntapVolume) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOntapVolume) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfOntapVolume) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfOntapVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfOntapVolume) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfOntapVolume) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfOntapVolume) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfOntapVolume) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfOntapVolume) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapVolume) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfOntapVolume) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOntapVolume) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfOntapVolume) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfOntapVolume) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfOntapVolume) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfOntapVolume) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfOntapVolume) PutAggregateConfiguration(value *TfOntapVolume_AggregateConfigurationProperty) {
	if err := t.validatePutAggregateConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAggregateConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOntapVolume) PutSnaplockConfiguration(value *TfOntapVolume_SnaplockConfigurationProperty) {
	if err := t.validatePutSnaplockConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSnaplockConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOntapVolume) PutTieringPolicy(value *TfOntapVolume_TieringPolicyProperty) {
	if err := t.validatePutTieringPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTieringPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOntapVolume) PutTimeouts(value *TfOntapVolume_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOntapVolume) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetAggregateConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetAggregateConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetBypassSnaplockEnterpriseRetention() {
	_jsii_.InvokeVoid(
		t,
		"resetBypassSnaplockEnterpriseRetention",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetCopyTagsToBackups() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyTagsToBackups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetFinalBackupTags() {
	_jsii_.InvokeVoid(
		t,
		"resetFinalBackupTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetJunctionPath() {
	_jsii_.InvokeVoid(
		t,
		"resetJunctionPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetOntapVolumeType() {
	_jsii_.InvokeVoid(
		t,
		"resetOntapVolumeType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetSecurityStyle() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityStyle",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetSizeInBytes() {
	_jsii_.InvokeVoid(
		t,
		"resetSizeInBytes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetSizeInMegabytes() {
	_jsii_.InvokeVoid(
		t,
		"resetSizeInMegabytes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetSkipFinalBackup() {
	_jsii_.InvokeVoid(
		t,
		"resetSkipFinalBackup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetSnaplockConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSnaplockConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetSnapshotPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetSnapshotPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetStorageEfficiencyEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageEfficiencyEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetTieringPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetTieringPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetVolumeStyle() {
	_jsii_.InvokeVoid(
		t,
		"resetVolumeStyle",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) ResetVolumeType() {
	_jsii_.InvokeVoid(
		t,
		"resetVolumeType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapVolume) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapVolume) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapVolume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapVolume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapVolume) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

