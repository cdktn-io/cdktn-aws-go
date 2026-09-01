package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume aws_fsx_openzfs_volume}.
// Experimental.
type AwsFsxOpenzfsVolume interface {
	cdktn.TerraformResource
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
	CopyTagsToSnapshots() interface{}
	// Experimental.
	SetCopyTagsToSnapshots(val interface{})
	// Experimental.
	CopyTagsToSnapshotsInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DataCompressionType() *string
	// Experimental.
	SetDataCompressionType(val *string)
	// Experimental.
	DataCompressionTypeInput() *string
	// Experimental.
	DeleteVolumeOptions() *[]*string
	// Experimental.
	SetDeleteVolumeOptions(val *[]*string)
	// Experimental.
	DeleteVolumeOptionsInput() *[]*string
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
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NfsExports() AwsFsxOpenzfsVolume_NfsExportsPropertyOutputReference
	// Experimental.
	NfsExportsInput() *AwsFsxOpenzfsVolume_NfsExportsProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OriginSnapshot() AwsFsxOpenzfsVolume_OriginSnapshotPropertyOutputReference
	// Experimental.
	OriginSnapshotInput() *AwsFsxOpenzfsVolume_OriginSnapshotProperty
	// Experimental.
	ParentVolumeId() *string
	// Experimental.
	SetParentVolumeId(val *string)
	// Experimental.
	ParentVolumeIdInput() *string
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
	ReadOnly() interface{}
	// Experimental.
	SetReadOnly(val interface{})
	// Experimental.
	ReadOnlyInput() interface{}
	// Experimental.
	RecordSizeKib() *float64
	// Experimental.
	SetRecordSizeKib(val *float64)
	// Experimental.
	RecordSizeKibInput() *float64
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	StorageCapacityQuotaGib() *float64
	// Experimental.
	SetStorageCapacityQuotaGib(val *float64)
	// Experimental.
	StorageCapacityQuotaGibInput() *float64
	// Experimental.
	StorageCapacityReservationGib() *float64
	// Experimental.
	SetStorageCapacityReservationGib(val *float64)
	// Experimental.
	StorageCapacityReservationGibInput() *float64
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
	Timeouts() AwsFsxOpenzfsVolume_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	UserAndGroupQuotas() AwsFsxOpenzfsVolume_UserAndGroupQuotasPropertyList
	// Experimental.
	UserAndGroupQuotasInput() interface{}
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
	PutNfsExports(value *AwsFsxOpenzfsVolume_NfsExportsProperty)
	// Experimental.
	PutOriginSnapshot(value *AwsFsxOpenzfsVolume_OriginSnapshotProperty)
	// Experimental.
	PutTimeouts(value *AwsFsxOpenzfsVolume_TimeoutsProperty)
	// Experimental.
	PutUserAndGroupQuotas(value interface{})
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
	ResetCopyTagsToSnapshots()
	// Experimental.
	ResetDataCompressionType()
	// Experimental.
	ResetDeleteVolumeOptions()
	// Experimental.
	ResetId()
	// Experimental.
	ResetNfsExports()
	// Experimental.
	ResetOriginSnapshot()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetReadOnly()
	// Experimental.
	ResetRecordSizeKib()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetStorageCapacityQuotaGib()
	// Experimental.
	ResetStorageCapacityReservationGib()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetUserAndGroupQuotas()
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

// The jsii proxy struct for AwsFsxOpenzfsVolume
type jsiiProxy_AwsFsxOpenzfsVolume struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) CopyTagsToSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) CopyTagsToSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) DataCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) DataCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) DeleteVolumeOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deleteVolumeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) DeleteVolumeOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deleteVolumeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) NfsExports() AwsFsxOpenzfsVolume_NfsExportsPropertyOutputReference {
	var returns AwsFsxOpenzfsVolume_NfsExportsPropertyOutputReference
	_jsii_.Get(
		j,
		"nfsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) NfsExportsInput() *AwsFsxOpenzfsVolume_NfsExportsProperty {
	var returns *AwsFsxOpenzfsVolume_NfsExportsProperty
	_jsii_.Get(
		j,
		"nfsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) OriginSnapshot() AwsFsxOpenzfsVolume_OriginSnapshotPropertyOutputReference {
	var returns AwsFsxOpenzfsVolume_OriginSnapshotPropertyOutputReference
	_jsii_.Get(
		j,
		"originSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) OriginSnapshotInput() *AwsFsxOpenzfsVolume_OriginSnapshotProperty {
	var returns *AwsFsxOpenzfsVolume_OriginSnapshotProperty
	_jsii_.Get(
		j,
		"originSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) ParentVolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentVolumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) ParentVolumeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentVolumeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) RecordSizeKib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recordSizeKib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) RecordSizeKibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recordSizeKibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) StorageCapacityQuotaGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityQuotaGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) StorageCapacityQuotaGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityQuotaGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) StorageCapacityReservationGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityReservationGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) StorageCapacityReservationGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityReservationGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) Timeouts() AwsFsxOpenzfsVolume_TimeoutsPropertyOutputReference {
	var returns AwsFsxOpenzfsVolume_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) UserAndGroupQuotas() AwsFsxOpenzfsVolume_UserAndGroupQuotasPropertyList {
	var returns AwsFsxOpenzfsVolume_UserAndGroupQuotasPropertyList
	_jsii_.Get(
		j,
		"userAndGroupQuotas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) UserAndGroupQuotasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userAndGroupQuotasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume) VolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume aws_fsx_openzfs_volume} Resource.
// Experimental.
func NewAwsFsxOpenzfsVolume(scope constructs.Construct, id *string, config *AwsFsxOpenzfsVolumeConfig) AwsFsxOpenzfsVolume {
	_init_.Initialize()

	if err := validateNewAwsFsxOpenzfsVolumeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFsxOpenzfsVolume{}

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsFsxOpenzfsVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume aws_fsx_openzfs_volume} Resource.
// Experimental.
func NewAwsFsxOpenzfsVolume_Override(a AwsFsxOpenzfsVolume, scope constructs.Construct, id *string, config *AwsFsxOpenzfsVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsFsxOpenzfsVolume",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetCopyTagsToSnapshots(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshots",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetDataCompressionType(val *string) {
	if err := j.validateSetDataCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCompressionType",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetDeleteVolumeOptions(val *[]*string) {
	if err := j.validateSetDeleteVolumeOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteVolumeOptions",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetParentVolumeId(val *string) {
	if err := j.validateSetParentVolumeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentVolumeId",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetRecordSizeKib(val *float64) {
	if err := j.validateSetRecordSizeKibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordSizeKib",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetStorageCapacityQuotaGib(val *float64) {
	if err := j.validateSetStorageCapacityQuotaGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCapacityQuotaGib",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetStorageCapacityReservationGib(val *float64) {
	if err := j.validateSetStorageCapacityReservationGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCapacityReservationGib",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsVolume)SetVolumeType(val *string) {
	if err := j.validateSetVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeType",
		val,
	)
}

// Generates CDKTN code for importing a AwsFsxOpenzfsVolume resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsFsxOpenzfsVolume_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsFsxOpenzfsVolume_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.AwsFsxOpenzfsVolume",
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
func AwsFsxOpenzfsVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsFsxOpenzfsVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.AwsFsxOpenzfsVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsFsxOpenzfsVolume_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsFsxOpenzfsVolume_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.AwsFsxOpenzfsVolume",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsFsxOpenzfsVolume_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsFsxOpenzfsVolume_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.AwsFsxOpenzfsVolume",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsFsxOpenzfsVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-fsx.AwsFsxOpenzfsVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFsxOpenzfsVolume) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFsxOpenzfsVolume) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFsxOpenzfsVolume) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFsxOpenzfsVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFsxOpenzfsVolume) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFsxOpenzfsVolume) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFsxOpenzfsVolume) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFsxOpenzfsVolume) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFsxOpenzfsVolume) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFsxOpenzfsVolume) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsFsxOpenzfsVolume) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) PutNfsExports(value *AwsFsxOpenzfsVolume_NfsExportsProperty) {
	if err := a.validatePutNfsExportsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNfsExports",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) PutOriginSnapshot(value *AwsFsxOpenzfsVolume_OriginSnapshotProperty) {
	if err := a.validatePutOriginSnapshotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOriginSnapshot",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) PutTimeouts(value *AwsFsxOpenzfsVolume_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) PutUserAndGroupQuotas(value interface{}) {
	if err := a.validatePutUserAndGroupQuotasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserAndGroupQuotas",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetCopyTagsToSnapshots() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyTagsToSnapshots",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetDataCompressionType() {
	_jsii_.InvokeVoid(
		a,
		"resetDataCompressionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetDeleteVolumeOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteVolumeOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetNfsExports() {
	_jsii_.InvokeVoid(
		a,
		"resetNfsExports",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetOriginSnapshot() {
	_jsii_.InvokeVoid(
		a,
		"resetOriginSnapshot",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetReadOnly() {
	_jsii_.InvokeVoid(
		a,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetRecordSizeKib() {
	_jsii_.InvokeVoid(
		a,
		"resetRecordSizeKib",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetStorageCapacityQuotaGib() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageCapacityQuotaGib",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetStorageCapacityReservationGib() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageCapacityReservationGib",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetUserAndGroupQuotas() {
	_jsii_.InvokeVoid(
		a,
		"resetUserAndGroupQuotas",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ResetVolumeType() {
	_jsii_.InvokeVoid(
		a,
		"resetVolumeType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxOpenzfsVolume) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

