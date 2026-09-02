package awsoracledatabaseaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure aws_odb_cloud_exadata_infrastructure}.
// Experimental.
type TfCloudExadataInfrastructure interface {
	cdktn.TerraformResource
	// Experimental.
	ActivatedStorageCount() *float64
	// Experimental.
	AdditionalStorageCount() *float64
	// Experimental.
	Arn() *string
	// Experimental.
	AvailabilityZone() *string
	// Experimental.
	SetAvailabilityZone(val *string)
	// Experimental.
	AvailabilityZoneId() *string
	// Experimental.
	SetAvailabilityZoneId(val *string)
	// Experimental.
	AvailabilityZoneIdInput() *string
	// Experimental.
	AvailabilityZoneInput() *string
	// Experimental.
	AvailableStorageSizeInGbs() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ComputeCount() *float64
	// Experimental.
	SetComputeCount(val *float64)
	// Experimental.
	ComputeCountInput() *float64
	// Experimental.
	ComputeModel() *string
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
	CpuCount() *float64
	// Experimental.
	CreatedAt() *string
	// Experimental.
	CustomerContactsToSendToOci() TfCloudExadataInfrastructure_CustomerContactsToSendToOciPropertyList
	// Experimental.
	CustomerContactsToSendToOciInput() interface{}
	// Experimental.
	DatabaseServerType() *string
	// Experimental.
	SetDatabaseServerType(val *string)
	// Experimental.
	DatabaseServerTypeInput() *string
	// Experimental.
	DataStorageSizeInTbs() *float64
	// Experimental.
	DbNodeStorageSizeInGbs() *float64
	// Experimental.
	DbServerVersion() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DisplayName() *string
	// Experimental.
	SetDisplayName(val *string)
	// Experimental.
	DisplayNameInput() *string
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
	LastMaintenanceRunId() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaintenanceWindow() TfCloudExadataInfrastructure_MaintenanceWindowPropertyList
	// Experimental.
	MaintenanceWindowInput() interface{}
	// Experimental.
	MaxCpuCount() *float64
	// Experimental.
	MaxDataStorageInTbs() *float64
	// Experimental.
	MaxDbNodeStorageSizeInGbs() *float64
	// Experimental.
	MaxMemoryInGbs() *float64
	// Experimental.
	MemorySizeInGbs() *float64
	// Experimental.
	MonthlyDbServerVersion() *string
	// Experimental.
	MonthlyStorageServerVersion() *string
	// Experimental.
	NextMaintenanceRunId() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Ocid() *string
	// Experimental.
	OciResourceAnchorName() *string
	// Experimental.
	OciUrl() *string
	// Experimental.
	PercentProgress() *float64
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
	Shape() *string
	// Experimental.
	SetShape(val *string)
	// Experimental.
	ShapeInput() *string
	// Experimental.
	Status() *string
	// Experimental.
	StatusReason() *string
	// Experimental.
	StorageCount() *float64
	// Experimental.
	SetStorageCount(val *float64)
	// Experimental.
	StorageCountInput() *float64
	// Experimental.
	StorageServerType() *string
	// Experimental.
	SetStorageServerType(val *string)
	// Experimental.
	StorageServerTypeInput() *string
	// Experimental.
	StorageServerVersion() *string
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsAll() cdktn.StringMap
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() TfCloudExadataInfrastructure_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TotalStorageSizeInGbs() *float64
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
	PutCustomerContactsToSendToOci(value interface{})
	// Experimental.
	PutMaintenanceWindow(value interface{})
	// Experimental.
	PutTimeouts(value *TfCloudExadataInfrastructure_TimeoutsProperty)
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
	ResetAvailabilityZone()
	// Experimental.
	ResetComputeCount()
	// Experimental.
	ResetCustomerContactsToSendToOci()
	// Experimental.
	ResetDatabaseServerType()
	// Experimental.
	ResetMaintenanceWindow()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetStorageCount()
	// Experimental.
	ResetStorageServerType()
	// Experimental.
	ResetTags()
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

// The jsii proxy struct for TfCloudExadataInfrastructure
type jsiiProxy_TfCloudExadataInfrastructure struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) ActivatedStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"activatedStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) AdditionalStorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"additionalStorageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) AvailabilityZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) AvailableStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) ComputeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) CpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) CustomerContactsToSendToOci() TfCloudExadataInfrastructure_CustomerContactsToSendToOciPropertyList {
	var returns TfCloudExadataInfrastructure_CustomerContactsToSendToOciPropertyList
	_jsii_.Get(
		j,
		"customerContactsToSendToOci",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) CustomerContactsToSendToOciInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerContactsToSendToOciInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) DatabaseServerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseServerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) DatabaseServerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseServerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) DataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) DbNodeStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) DbServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) LastMaintenanceRunId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastMaintenanceRunId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) MaintenanceWindow() TfCloudExadataInfrastructure_MaintenanceWindowPropertyList {
	var returns TfCloudExadataInfrastructure_MaintenanceWindowPropertyList
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) MaintenanceWindowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) MaxCpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) MaxDataStorageInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDataStorageInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) MaxDbNodeStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDbNodeStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) MaxMemoryInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMemoryInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) MemorySizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) MonthlyDbServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlyDbServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) MonthlyStorageServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monthlyStorageServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) NextMaintenanceRunId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextMaintenanceRunId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) OciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) PercentProgress() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentProgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) ShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) StorageCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) StorageCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) StorageServerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageServerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) StorageServerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageServerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) StorageServerVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageServerVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) Timeouts() TfCloudExadataInfrastructure_TimeoutsPropertyOutputReference {
	var returns TfCloudExadataInfrastructure_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudExadataInfrastructure) TotalStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalStorageSizeInGbs",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure aws_odb_cloud_exadata_infrastructure} Resource.
// Experimental.
func NewTfCloudExadataInfrastructure(scope constructs.Construct, id *string, config *TfCloudExadataInfrastructureConfig) TfCloudExadataInfrastructure {
	_init_.Initialize()

	if err := validateNewTfCloudExadataInfrastructureParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCloudExadataInfrastructure{}

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.TfCloudExadataInfrastructure",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure aws_odb_cloud_exadata_infrastructure} Resource.
// Experimental.
func NewTfCloudExadataInfrastructure_Override(t TfCloudExadataInfrastructure, scope constructs.Construct, id *string, config *TfCloudExadataInfrastructureConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.TfCloudExadataInfrastructure",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetAvailabilityZoneId(val *string) {
	if err := j.validateSetAvailabilityZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneId",
		val,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetComputeCount(val *float64) {
	if err := j.validateSetComputeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeCount",
		val,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetDatabaseServerType(val *string) {
	if err := j.validateSetDatabaseServerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseServerType",
		val,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetShape(val *string) {
	if err := j.validateSetShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shape",
		val,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetStorageCount(val *float64) {
	if err := j.validateSetStorageCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCount",
		val,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetStorageServerType(val *string) {
	if err := j.validateSetStorageServerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageServerType",
		val,
	)
}

func (j *jsiiProxy_TfCloudExadataInfrastructure)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTN code for importing a TfCloudExadataInfrastructure resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfCloudExadataInfrastructure_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfCloudExadataInfrastructure_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.TfCloudExadataInfrastructure",
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
func TfCloudExadataInfrastructure_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCloudExadataInfrastructure_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.TfCloudExadataInfrastructure",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCloudExadataInfrastructure_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCloudExadataInfrastructure_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.TfCloudExadataInfrastructure",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCloudExadataInfrastructure_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCloudExadataInfrastructure_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.TfCloudExadataInfrastructure",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfCloudExadataInfrastructure_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-oracle-database-aws.TfCloudExadataInfrastructure",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCloudExadataInfrastructure) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCloudExadataInfrastructure) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCloudExadataInfrastructure) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCloudExadataInfrastructure) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCloudExadataInfrastructure) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCloudExadataInfrastructure) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCloudExadataInfrastructure) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCloudExadataInfrastructure) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCloudExadataInfrastructure) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCloudExadataInfrastructure) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfCloudExadataInfrastructure) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) PutCustomerContactsToSendToOci(value interface{}) {
	if err := t.validatePutCustomerContactsToSendToOciParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomerContactsToSendToOci",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) PutMaintenanceWindow(value interface{}) {
	if err := t.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) PutTimeouts(value *TfCloudExadataInfrastructure_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		t,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) ResetComputeCount() {
	_jsii_.InvokeVoid(
		t,
		"resetComputeCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) ResetCustomerContactsToSendToOci() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomerContactsToSendToOci",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) ResetDatabaseServerType() {
	_jsii_.InvokeVoid(
		t,
		"resetDatabaseServerType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		t,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) ResetStorageCount() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) ResetStorageServerType() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageServerType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudExadataInfrastructure) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

