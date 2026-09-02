package awsoracledatabaseaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster aws_odb_cloud_autonomous_vm_cluster}.
// Experimental.
type TfCloudAutonomousVmCluster interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	AutonomousDataStoragePercentage() *float64
	// Experimental.
	AutonomousDataStorageSizeInTbs() *float64
	// Experimental.
	SetAutonomousDataStorageSizeInTbs(val *float64)
	// Experimental.
	AutonomousDataStorageSizeInTbsInput() *float64
	// Experimental.
	AvailableAutonomousDataStorageSizeInTbs() *float64
	// Experimental.
	AvailableContainerDatabases() *float64
	// Experimental.
	AvailableCpus() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CloudExadataInfrastructureArn() *string
	// Experimental.
	SetCloudExadataInfrastructureArn(val *string)
	// Experimental.
	CloudExadataInfrastructureArnInput() *string
	// Experimental.
	CloudExadataInfrastructureId() *string
	// Experimental.
	SetCloudExadataInfrastructureId(val *string)
	// Experimental.
	CloudExadataInfrastructureIdInput() *string
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
	CpuCoreCount() *float64
	// Experimental.
	CpuCoreCountPerNode() *float64
	// Experimental.
	SetCpuCoreCountPerNode(val *float64)
	// Experimental.
	CpuCoreCountPerNodeInput() *float64
	// Experimental.
	CpuPercentage() *float64
	// Experimental.
	CreatedAt() *string
	// Experimental.
	DataStorageSizeInGbs() *float64
	// Experimental.
	DataStorageSizeInTbs() *float64
	// Experimental.
	DbServers() *[]*string
	// Experimental.
	SetDbServers(val *[]*string)
	// Experimental.
	DbServersInput() *[]*string
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
	DisplayName() *string
	// Experimental.
	SetDisplayName(val *string)
	// Experimental.
	DisplayNameInput() *string
	// Experimental.
	Domain() *string
	// Experimental.
	ExadataStorageInTbsLowestScaledValue() *float64
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Hostname() *string
	// Experimental.
	Id() *string
	// Experimental.
	IsMtlsEnabledVmCluster() interface{}
	// Experimental.
	SetIsMtlsEnabledVmCluster(val interface{})
	// Experimental.
	IsMtlsEnabledVmClusterInput() interface{}
	// Experimental.
	LicenseModel() *string
	// Experimental.
	SetLicenseModel(val *string)
	// Experimental.
	LicenseModelInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaintenanceWindow() TfCloudAutonomousVmCluster_MaintenanceWindowPropertyList
	// Experimental.
	MaintenanceWindowInput() interface{}
	// Experimental.
	MaxAcdsLowestScaledValue() *float64
	// Experimental.
	MemoryPerOracleComputeUnitInGbs() *float64
	// Experimental.
	SetMemoryPerOracleComputeUnitInGbs(val *float64)
	// Experimental.
	MemoryPerOracleComputeUnitInGbsInput() *float64
	// Experimental.
	MemorySizeInGbs() *float64
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	NodeCount() *float64
	// Experimental.
	NonProvisionableAutonomousContainerDatabases() *float64
	// Experimental.
	Ocid() *string
	// Experimental.
	OciResourceAnchorName() *string
	// Experimental.
	OciUrl() *string
	// Experimental.
	OdbNetworkArn() *string
	// Experimental.
	SetOdbNetworkArn(val *string)
	// Experimental.
	OdbNetworkArnInput() *string
	// Experimental.
	OdbNetworkId() *string
	// Experimental.
	SetOdbNetworkId(val *string)
	// Experimental.
	OdbNetworkIdInput() *string
	// Experimental.
	OdbNodeStorageSizeInGbs() *float64
	// Experimental.
	PercentProgress() *float64
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	ProvisionableAutonomousContainerDatabases() *float64
	// Experimental.
	ProvisionedAutonomousContainerDatabases() *float64
	// Experimental.
	ProvisionedCpus() *float64
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	ReclaimableCpus() *float64
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	ReservedCpus() *float64
	// Experimental.
	ScanListenerPortNonTls() *float64
	// Experimental.
	SetScanListenerPortNonTls(val *float64)
	// Experimental.
	ScanListenerPortNonTlsInput() *float64
	// Experimental.
	ScanListenerPortTls() *float64
	// Experimental.
	SetScanListenerPortTls(val *float64)
	// Experimental.
	ScanListenerPortTlsInput() *float64
	// Experimental.
	Shape() *string
	// Experimental.
	Status() *string
	// Experimental.
	StatusReason() *string
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
	TimeDatabaseSslCertificateExpires() *string
	// Experimental.
	TimeOrdsCertificateExpires() *string
	// Experimental.
	Timeouts() TfCloudAutonomousVmCluster_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TimeZone() *string
	// Experimental.
	SetTimeZone(val *string)
	// Experimental.
	TimeZoneInput() *string
	// Experimental.
	TotalContainerDatabases() *float64
	// Experimental.
	SetTotalContainerDatabases(val *float64)
	// Experimental.
	TotalContainerDatabasesInput() *float64
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
	PutMaintenanceWindow(value interface{})
	// Experimental.
	PutTimeouts(value *TfCloudAutonomousVmCluster_TimeoutsProperty)
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
	ResetCloudExadataInfrastructureArn()
	// Experimental.
	ResetCloudExadataInfrastructureId()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetIsMtlsEnabledVmCluster()
	// Experimental.
	ResetLicenseModel()
	// Experimental.
	ResetMaintenanceWindow()
	// Experimental.
	ResetOdbNetworkArn()
	// Experimental.
	ResetOdbNetworkId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTimeZone()
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

// The jsii proxy struct for TfCloudAutonomousVmCluster
type jsiiProxy_TfCloudAutonomousVmCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) AutonomousDataStoragePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autonomousDataStoragePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) AutonomousDataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autonomousDataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) AutonomousDataStorageSizeInTbsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autonomousDataStorageSizeInTbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) AvailableAutonomousDataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableAutonomousDataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) AvailableContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) AvailableCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"availableCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) CloudExadataInfrastructureArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) CloudExadataInfrastructureArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) CloudExadataInfrastructureId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) CloudExadataInfrastructureIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) CpuCoreCountPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCountPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) CpuCoreCountPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCountPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) CpuPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) DataStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) DataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) DbServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) DbServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) ExadataStorageInTbsLowestScaledValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"exadataStorageInTbsLowestScaledValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) IsMtlsEnabledVmCluster() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMtlsEnabledVmCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) IsMtlsEnabledVmClusterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMtlsEnabledVmClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) MaintenanceWindow() TfCloudAutonomousVmCluster_MaintenanceWindowPropertyList {
	var returns TfCloudAutonomousVmCluster_MaintenanceWindowPropertyList
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) MaintenanceWindowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) MaxAcdsLowestScaledValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAcdsLowestScaledValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) MemoryPerOracleComputeUnitInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerOracleComputeUnitInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) MemoryPerOracleComputeUnitInGbsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerOracleComputeUnitInGbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) MemorySizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) NonProvisionableAutonomousContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nonProvisionableAutonomousContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) OciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) OdbNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) OdbNetworkArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) OdbNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) OdbNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) OdbNodeStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"odbNodeStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) PercentProgress() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentProgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) ProvisionableAutonomousContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionableAutonomousContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) ProvisionedAutonomousContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedAutonomousContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) ProvisionedCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) ReclaimableCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reclaimableCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) ReservedCpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) ScanListenerPortNonTls() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortNonTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) ScanListenerPortNonTlsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortNonTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) ScanListenerPortTls() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) ScanListenerPortTlsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) TimeDatabaseSslCertificateExpires() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeDatabaseSslCertificateExpires",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) TimeOrdsCertificateExpires() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOrdsCertificateExpires",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) Timeouts() TfCloudAutonomousVmCluster_TimeoutsPropertyOutputReference {
	var returns TfCloudAutonomousVmCluster_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) TotalContainerDatabases() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalContainerDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster) TotalContainerDatabasesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalContainerDatabasesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster aws_odb_cloud_autonomous_vm_cluster} Resource.
// Experimental.
func NewTfCloudAutonomousVmCluster(scope constructs.Construct, id *string, config *TfCloudAutonomousVmClusterConfig) TfCloudAutonomousVmCluster {
	_init_.Initialize()

	if err := validateNewTfCloudAutonomousVmClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCloudAutonomousVmCluster{}

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.TfCloudAutonomousVmCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster aws_odb_cloud_autonomous_vm_cluster} Resource.
// Experimental.
func NewTfCloudAutonomousVmCluster_Override(t TfCloudAutonomousVmCluster, scope constructs.Construct, id *string, config *TfCloudAutonomousVmClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.TfCloudAutonomousVmCluster",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetAutonomousDataStorageSizeInTbs(val *float64) {
	if err := j.validateSetAutonomousDataStorageSizeInTbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autonomousDataStorageSizeInTbs",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetCloudExadataInfrastructureArn(val *string) {
	if err := j.validateSetCloudExadataInfrastructureArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudExadataInfrastructureArn",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetCloudExadataInfrastructureId(val *string) {
	if err := j.validateSetCloudExadataInfrastructureIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudExadataInfrastructureId",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetCpuCoreCountPerNode(val *float64) {
	if err := j.validateSetCpuCoreCountPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCoreCountPerNode",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetDbServers(val *[]*string) {
	if err := j.validateSetDbServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbServers",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetIsMtlsEnabledVmCluster(val interface{}) {
	if err := j.validateSetIsMtlsEnabledVmClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMtlsEnabledVmCluster",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetMemoryPerOracleComputeUnitInGbs(val *float64) {
	if err := j.validateSetMemoryPerOracleComputeUnitInGbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryPerOracleComputeUnitInGbs",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetOdbNetworkArn(val *string) {
	if err := j.validateSetOdbNetworkArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odbNetworkArn",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetOdbNetworkId(val *string) {
	if err := j.validateSetOdbNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odbNetworkId",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetScanListenerPortNonTls(val *float64) {
	if err := j.validateSetScanListenerPortNonTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanListenerPortNonTls",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetScanListenerPortTls(val *float64) {
	if err := j.validateSetScanListenerPortTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanListenerPortTls",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster)SetTotalContainerDatabases(val *float64) {
	if err := j.validateSetTotalContainerDatabasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalContainerDatabases",
		val,
	)
}

// Generates CDKTN code for importing a TfCloudAutonomousVmCluster resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfCloudAutonomousVmCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfCloudAutonomousVmCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.TfCloudAutonomousVmCluster",
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
func TfCloudAutonomousVmCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCloudAutonomousVmCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.TfCloudAutonomousVmCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCloudAutonomousVmCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCloudAutonomousVmCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.TfCloudAutonomousVmCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCloudAutonomousVmCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCloudAutonomousVmCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.TfCloudAutonomousVmCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfCloudAutonomousVmCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-oracle-database-aws.TfCloudAutonomousVmCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) PutMaintenanceWindow(value interface{}) {
	if err := t.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) PutTimeouts(value *TfCloudAutonomousVmCluster_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ResetCloudExadataInfrastructureArn() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudExadataInfrastructureArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ResetCloudExadataInfrastructureId() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudExadataInfrastructureId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ResetIsMtlsEnabledVmCluster() {
	_jsii_.InvokeVoid(
		t,
		"resetIsMtlsEnabledVmCluster",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ResetLicenseModel() {
	_jsii_.InvokeVoid(
		t,
		"resetLicenseModel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		t,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ResetOdbNetworkArn() {
	_jsii_.InvokeVoid(
		t,
		"resetOdbNetworkArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ResetOdbNetworkId() {
	_jsii_.InvokeVoid(
		t,
		"resetOdbNetworkId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ResetTimeZone() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

