package awsoracledatabaseaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster aws_odb_cloud_vm_cluster}.
// Experimental.
type TfCloudVmCluster interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
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
	ClusterName() *string
	// Experimental.
	SetClusterName(val *string)
	// Experimental.
	ClusterNameInput() *string
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
	SetCpuCoreCount(val *float64)
	// Experimental.
	CpuCoreCountInput() *float64
	// Experimental.
	CreatedAt() *string
	// Experimental.
	DataCollectionOptions() TfCloudVmCluster_DataCollectionOptionsPropertyList
	// Experimental.
	DataCollectionOptionsInput() interface{}
	// Experimental.
	DataStorageSizeInTbs() *float64
	// Experimental.
	SetDataStorageSizeInTbs(val *float64)
	// Experimental.
	DataStorageSizeInTbsInput() *float64
	// Experimental.
	DbNodeStorageSizeInGbs() *float64
	// Experimental.
	SetDbNodeStorageSizeInGbs(val *float64)
	// Experimental.
	DbNodeStorageSizeInGbsInput() *float64
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
	DiskRedundancy() *string
	// Experimental.
	DisplayName() *string
	// Experimental.
	SetDisplayName(val *string)
	// Experimental.
	DisplayNameInput() *string
	// Experimental.
	Domain() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	GiVersion() *string
	// Experimental.
	SetGiVersion(val *string)
	// Experimental.
	GiVersionComputed() *string
	// Experimental.
	GiVersionInput() *string
	// Experimental.
	HostnamePrefix() *string
	// Experimental.
	SetHostnamePrefix(val *string)
	// Experimental.
	HostnamePrefixComputed() *string
	// Experimental.
	HostnamePrefixInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	IormConfigCache() TfCloudVmCluster_IormConfigCachePropertyList
	// Experimental.
	IsLocalBackupEnabled() interface{}
	// Experimental.
	SetIsLocalBackupEnabled(val interface{})
	// Experimental.
	IsLocalBackupEnabledInput() interface{}
	// Experimental.
	IsSparseDiskgroupEnabled() interface{}
	// Experimental.
	SetIsSparseDiskgroupEnabled(val interface{})
	// Experimental.
	IsSparseDiskgroupEnabledInput() interface{}
	// Experimental.
	LastUpdateHistoryEntryId() *string
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
	ListenerPort() *float64
	// Experimental.
	MemorySizeInGbs() *float64
	// Experimental.
	SetMemorySizeInGbs(val *float64)
	// Experimental.
	MemorySizeInGbsInput() *float64
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	NodeCount() *float64
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
	ScanDnsName() *string
	// Experimental.
	ScanDnsRecordId() *string
	// Experimental.
	ScanIpIds() *[]*string
	// Experimental.
	ScanListenerPortTcp() *float64
	// Experimental.
	SetScanListenerPortTcp(val *float64)
	// Experimental.
	ScanListenerPortTcpInput() *float64
	// Experimental.
	Shape() *string
	// Experimental.
	SshPublicKeys() *[]*string
	// Experimental.
	SetSshPublicKeys(val *[]*string)
	// Experimental.
	SshPublicKeysInput() *[]*string
	// Experimental.
	Status() *string
	// Experimental.
	StatusReason() *string
	// Experimental.
	StorageSizeInGbs() *float64
	// Experimental.
	SystemVersion() *string
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
	Timeouts() TfCloudVmCluster_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Timezone() *string
	// Experimental.
	SetTimezone(val *string)
	// Experimental.
	TimezoneInput() *string
	// Experimental.
	VipIds() *[]*string
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
	PutDataCollectionOptions(value interface{})
	// Experimental.
	PutTimeouts(value *TfCloudVmCluster_TimeoutsProperty)
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
	ResetClusterName()
	// Experimental.
	ResetDataCollectionOptions()
	// Experimental.
	ResetDbNodeStorageSizeInGbs()
	// Experimental.
	ResetIsLocalBackupEnabled()
	// Experimental.
	ResetIsSparseDiskgroupEnabled()
	// Experimental.
	ResetLicenseModel()
	// Experimental.
	ResetMemorySizeInGbs()
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
	ResetScanListenerPortTcp()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTimezone()
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

// The jsii proxy struct for TfCloudVmCluster
type jsiiProxy_TfCloudVmCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfCloudVmCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) CloudExadataInfrastructureArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) CloudExadataInfrastructureArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) CloudExadataInfrastructureId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) CloudExadataInfrastructureIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudExadataInfrastructureIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) ComputeModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) CpuCoreCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) DataCollectionOptions() TfCloudVmCluster_DataCollectionOptionsPropertyList {
	var returns TfCloudVmCluster_DataCollectionOptionsPropertyList
	_jsii_.Get(
		j,
		"dataCollectionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) DataCollectionOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataCollectionOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) DataStorageSizeInTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) DataStorageSizeInTbsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) DbNodeStorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) DbNodeStorageSizeInGbsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeInGbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) DbServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) DbServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) DiskRedundancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskRedundancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) GiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"giVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) GiVersionComputed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"giVersionComputed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) GiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"giVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) HostnamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) HostnamePrefixComputed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefixComputed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) HostnamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) IormConfigCache() TfCloudVmCluster_IormConfigCachePropertyList {
	var returns TfCloudVmCluster_IormConfigCachePropertyList
	_jsii_.Get(
		j,
		"iormConfigCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) IsLocalBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isLocalBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) IsLocalBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isLocalBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) IsSparseDiskgroupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSparseDiskgroupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) IsSparseDiskgroupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSparseDiskgroupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) LastUpdateHistoryEntryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdateHistoryEntryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) ListenerPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"listenerPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) MemorySizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) MemorySizeInGbsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInGbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) OciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) OdbNetworkArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) OdbNetworkArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) OdbNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) OdbNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"odbNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) PercentProgress() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentProgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) ScanDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) ScanDnsRecordId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanDnsRecordId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) ScanIpIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scanIpIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) ScanListenerPortTcp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) ScanListenerPortTcpInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) StorageSizeInGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSizeInGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) SystemVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) Timeouts() TfCloudVmCluster_TimeoutsPropertyOutputReference {
	var returns TfCloudVmCluster_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudVmCluster) VipIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vipIds",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster aws_odb_cloud_vm_cluster} Resource.
// Experimental.
func NewTfCloudVmCluster(scope constructs.Construct, id *string, config *TfCloudVmClusterConfig) TfCloudVmCluster {
	_init_.Initialize()

	if err := validateNewTfCloudVmClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCloudVmCluster{}

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.TfCloudVmCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster aws_odb_cloud_vm_cluster} Resource.
// Experimental.
func NewTfCloudVmCluster_Override(t TfCloudVmCluster, scope constructs.Construct, id *string, config *TfCloudVmClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.TfCloudVmCluster",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetCloudExadataInfrastructureArn(val *string) {
	if err := j.validateSetCloudExadataInfrastructureArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudExadataInfrastructureArn",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetCloudExadataInfrastructureId(val *string) {
	if err := j.validateSetCloudExadataInfrastructureIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudExadataInfrastructureId",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetCpuCoreCount(val *float64) {
	if err := j.validateSetCpuCoreCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCoreCount",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetDataStorageSizeInTbs(val *float64) {
	if err := j.validateSetDataStorageSizeInTbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStorageSizeInTbs",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetDbNodeStorageSizeInGbs(val *float64) {
	if err := j.validateSetDbNodeStorageSizeInGbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbNodeStorageSizeInGbs",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetDbServers(val *[]*string) {
	if err := j.validateSetDbServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbServers",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetGiVersion(val *string) {
	if err := j.validateSetGiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"giVersion",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetHostnamePrefix(val *string) {
	if err := j.validateSetHostnamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnamePrefix",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetIsLocalBackupEnabled(val interface{}) {
	if err := j.validateSetIsLocalBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isLocalBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetIsSparseDiskgroupEnabled(val interface{}) {
	if err := j.validateSetIsSparseDiskgroupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSparseDiskgroupEnabled",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetMemorySizeInGbs(val *float64) {
	if err := j.validateSetMemorySizeInGbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySizeInGbs",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetOdbNetworkArn(val *string) {
	if err := j.validateSetOdbNetworkArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odbNetworkArn",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetOdbNetworkId(val *string) {
	if err := j.validateSetOdbNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"odbNetworkId",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetScanListenerPortTcp(val *float64) {
	if err := j.validateSetScanListenerPortTcpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanListenerPortTcp",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetSshPublicKeys(val *[]*string) {
	if err := j.validateSetSshPublicKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfCloudVmCluster)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

// Generates CDKTN code for importing a TfCloudVmCluster resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfCloudVmCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfCloudVmCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.TfCloudVmCluster",
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
func TfCloudVmCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCloudVmCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.TfCloudVmCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCloudVmCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCloudVmCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.TfCloudVmCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCloudVmCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCloudVmCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.TfCloudVmCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfCloudVmCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-oracle-database-aws.TfCloudVmCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfCloudVmCluster) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfCloudVmCluster) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfCloudVmCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCloudVmCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCloudVmCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCloudVmCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCloudVmCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCloudVmCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCloudVmCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCloudVmCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCloudVmCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCloudVmCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudVmCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfCloudVmCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCloudVmCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfCloudVmCluster) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCloudVmCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfCloudVmCluster) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCloudVmCluster) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfCloudVmCluster) PutDataCollectionOptions(value interface{}) {
	if err := t.validatePutDataCollectionOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDataCollectionOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCloudVmCluster) PutTimeouts(value *TfCloudVmCluster_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCloudVmCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetCloudExadataInfrastructureArn() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudExadataInfrastructureArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetCloudExadataInfrastructureId() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudExadataInfrastructureId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetClusterName() {
	_jsii_.InvokeVoid(
		t,
		"resetClusterName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetDataCollectionOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetDataCollectionOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetDbNodeStorageSizeInGbs() {
	_jsii_.InvokeVoid(
		t,
		"resetDbNodeStorageSizeInGbs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetIsLocalBackupEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetIsLocalBackupEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetIsSparseDiskgroupEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetIsSparseDiskgroupEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetLicenseModel() {
	_jsii_.InvokeVoid(
		t,
		"resetLicenseModel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetMemorySizeInGbs() {
	_jsii_.InvokeVoid(
		t,
		"resetMemorySizeInGbs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetOdbNetworkArn() {
	_jsii_.InvokeVoid(
		t,
		"resetOdbNetworkArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetOdbNetworkId() {
	_jsii_.InvokeVoid(
		t,
		"resetOdbNetworkId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetScanListenerPortTcp() {
	_jsii_.InvokeVoid(
		t,
		"resetScanListenerPortTcp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) ResetTimezone() {
	_jsii_.InvokeVoid(
		t,
		"resetTimezone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudVmCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudVmCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudVmCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudVmCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudVmCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudVmCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudVmCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

