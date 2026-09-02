package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system aws_fsx_lustre_file_system}.
// Experimental.
type TfLustreFileSystem interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	AutoImportPolicy() *string
	// Experimental.
	SetAutoImportPolicy(val *string)
	// Experimental.
	AutoImportPolicyInput() *string
	// Experimental.
	AutomaticBackupRetentionDays() *float64
	// Experimental.
	SetAutomaticBackupRetentionDays(val *float64)
	// Experimental.
	AutomaticBackupRetentionDaysInput() *float64
	// Experimental.
	BackupId() *string
	// Experimental.
	SetBackupId(val *string)
	// Experimental.
	BackupIdInput() *string
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
	DailyAutomaticBackupStartTime() *string
	// Experimental.
	SetDailyAutomaticBackupStartTime(val *string)
	// Experimental.
	DailyAutomaticBackupStartTimeInput() *string
	// Experimental.
	DataCompressionType() *string
	// Experimental.
	SetDataCompressionType(val *string)
	// Experimental.
	DataCompressionTypeInput() *string
	// Experimental.
	DataReadCacheConfiguration() TfLustreFileSystem_DataReadCacheConfigurationPropertyOutputReference
	// Experimental.
	DataReadCacheConfigurationInput() *TfLustreFileSystem_DataReadCacheConfigurationProperty
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DeploymentType() *string
	// Experimental.
	SetDeploymentType(val *string)
	// Experimental.
	DeploymentTypeInput() *string
	// Experimental.
	DnsName() *string
	// Experimental.
	DriveCacheType() *string
	// Experimental.
	SetDriveCacheType(val *string)
	// Experimental.
	DriveCacheTypeInput() *string
	// Experimental.
	EfaEnabled() interface{}
	// Experimental.
	SetEfaEnabled(val interface{})
	// Experimental.
	EfaEnabledInput() interface{}
	// Experimental.
	ExportPath() *string
	// Experimental.
	SetExportPath(val *string)
	// Experimental.
	ExportPathInput() *string
	// Experimental.
	FileSystemTypeVersion() *string
	// Experimental.
	SetFileSystemTypeVersion(val *string)
	// Experimental.
	FileSystemTypeVersionInput() *string
	// Experimental.
	FinalBackupTags() *map[string]*string
	// Experimental.
	SetFinalBackupTags(val *map[string]*string)
	// Experimental.
	FinalBackupTagsInput() *map[string]*string
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
	ImportedFileChunkSize() *float64
	// Experimental.
	SetImportedFileChunkSize(val *float64)
	// Experimental.
	ImportedFileChunkSizeInput() *float64
	// Experimental.
	ImportPath() *string
	// Experimental.
	SetImportPath(val *string)
	// Experimental.
	ImportPathInput() *string
	// Experimental.
	KmsKeyId() *string
	// Experimental.
	SetKmsKeyId(val *string)
	// Experimental.
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LogConfiguration() TfLustreFileSystem_LogConfigurationPropertyOutputReference
	// Experimental.
	LogConfigurationInput() *TfLustreFileSystem_LogConfigurationProperty
	// Experimental.
	MetadataConfiguration() TfLustreFileSystem_MetadataConfigurationPropertyOutputReference
	// Experimental.
	MetadataConfigurationInput() *TfLustreFileSystem_MetadataConfigurationProperty
	// Experimental.
	MountName() *string
	// Experimental.
	NetworkInterfaceIds() *[]*string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OwnerId() *string
	// Experimental.
	PerUnitStorageThroughput() *float64
	// Experimental.
	SetPerUnitStorageThroughput(val *float64)
	// Experimental.
	PerUnitStorageThroughputInput() *float64
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
	RootSquashConfiguration() TfLustreFileSystem_RootSquashConfigurationPropertyOutputReference
	// Experimental.
	RootSquashConfigurationInput() *TfLustreFileSystem_RootSquashConfigurationProperty
	// Experimental.
	SecurityGroupIds() *[]*string
	// Experimental.
	SetSecurityGroupIds(val *[]*string)
	// Experimental.
	SecurityGroupIdsInput() *[]*string
	// Experimental.
	SkipFinalBackup() interface{}
	// Experimental.
	SetSkipFinalBackup(val interface{})
	// Experimental.
	SkipFinalBackupInput() interface{}
	// Experimental.
	StorageCapacity() *float64
	// Experimental.
	SetStorageCapacity(val *float64)
	// Experimental.
	StorageCapacityInput() *float64
	// Experimental.
	StorageType() *string
	// Experimental.
	SetStorageType(val *string)
	// Experimental.
	StorageTypeInput() *string
	// Experimental.
	SubnetIds() *[]*string
	// Experimental.
	SetSubnetIds(val *[]*string)
	// Experimental.
	SubnetIdsInput() *[]*string
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
	ThroughputCapacity() *float64
	// Experimental.
	SetThroughputCapacity(val *float64)
	// Experimental.
	ThroughputCapacityInput() *float64
	// Experimental.
	Timeouts() TfLustreFileSystem_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	VpcId() *string
	// Experimental.
	WeeklyMaintenanceStartTime() *string
	// Experimental.
	SetWeeklyMaintenanceStartTime(val *string)
	// Experimental.
	WeeklyMaintenanceStartTimeInput() *string
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
	PutDataReadCacheConfiguration(value *TfLustreFileSystem_DataReadCacheConfigurationProperty)
	// Experimental.
	PutLogConfiguration(value *TfLustreFileSystem_LogConfigurationProperty)
	// Experimental.
	PutMetadataConfiguration(value *TfLustreFileSystem_MetadataConfigurationProperty)
	// Experimental.
	PutRootSquashConfiguration(value *TfLustreFileSystem_RootSquashConfigurationProperty)
	// Experimental.
	PutTimeouts(value *TfLustreFileSystem_TimeoutsProperty)
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
	ResetAutoImportPolicy()
	// Experimental.
	ResetAutomaticBackupRetentionDays()
	// Experimental.
	ResetBackupId()
	// Experimental.
	ResetCopyTagsToBackups()
	// Experimental.
	ResetDailyAutomaticBackupStartTime()
	// Experimental.
	ResetDataCompressionType()
	// Experimental.
	ResetDataReadCacheConfiguration()
	// Experimental.
	ResetDeploymentType()
	// Experimental.
	ResetDriveCacheType()
	// Experimental.
	ResetEfaEnabled()
	// Experimental.
	ResetExportPath()
	// Experimental.
	ResetFileSystemTypeVersion()
	// Experimental.
	ResetFinalBackupTags()
	// Experimental.
	ResetId()
	// Experimental.
	ResetImportedFileChunkSize()
	// Experimental.
	ResetImportPath()
	// Experimental.
	ResetKmsKeyId()
	// Experimental.
	ResetLogConfiguration()
	// Experimental.
	ResetMetadataConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPerUnitStorageThroughput()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRootSquashConfiguration()
	// Experimental.
	ResetSecurityGroupIds()
	// Experimental.
	ResetSkipFinalBackup()
	// Experimental.
	ResetStorageCapacity()
	// Experimental.
	ResetStorageType()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetThroughputCapacity()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetWeeklyMaintenanceStartTime()
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

// The jsii proxy struct for TfLustreFileSystem
type jsiiProxy_TfLustreFileSystem struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfLustreFileSystem) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) AutoImportPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoImportPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) AutoImportPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoImportPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) AutomaticBackupRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) AutomaticBackupRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) BackupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) BackupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) CopyTagsToBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) CopyTagsToBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) DailyAutomaticBackupStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) DailyAutomaticBackupStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) DataCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) DataCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) DataReadCacheConfiguration() TfLustreFileSystem_DataReadCacheConfigurationPropertyOutputReference {
	var returns TfLustreFileSystem_DataReadCacheConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"dataReadCacheConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) DataReadCacheConfigurationInput() *TfLustreFileSystem_DataReadCacheConfigurationProperty {
	var returns *TfLustreFileSystem_DataReadCacheConfigurationProperty
	_jsii_.Get(
		j,
		"dataReadCacheConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) DeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) DeploymentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) DriveCacheType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driveCacheType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) DriveCacheTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driveCacheTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) EfaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efaEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) EfaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efaEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) ExportPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) ExportPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) FileSystemTypeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemTypeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) FileSystemTypeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemTypeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) FinalBackupTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"finalBackupTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) FinalBackupTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"finalBackupTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) ImportedFileChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importedFileChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) ImportedFileChunkSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importedFileChunkSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) ImportPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) ImportPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) LogConfiguration() TfLustreFileSystem_LogConfigurationPropertyOutputReference {
	var returns TfLustreFileSystem_LogConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) LogConfigurationInput() *TfLustreFileSystem_LogConfigurationProperty {
	var returns *TfLustreFileSystem_LogConfigurationProperty
	_jsii_.Get(
		j,
		"logConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) MetadataConfiguration() TfLustreFileSystem_MetadataConfigurationPropertyOutputReference {
	var returns TfLustreFileSystem_MetadataConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"metadataConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) MetadataConfigurationInput() *TfLustreFileSystem_MetadataConfigurationProperty {
	var returns *TfLustreFileSystem_MetadataConfigurationProperty
	_jsii_.Get(
		j,
		"metadataConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) MountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) NetworkInterfaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) PerUnitStorageThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perUnitStorageThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) PerUnitStorageThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perUnitStorageThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) RootSquashConfiguration() TfLustreFileSystem_RootSquashConfigurationPropertyOutputReference {
	var returns TfLustreFileSystem_RootSquashConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"rootSquashConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) RootSquashConfigurationInput() *TfLustreFileSystem_RootSquashConfigurationProperty {
	var returns *TfLustreFileSystem_RootSquashConfigurationProperty
	_jsii_.Get(
		j,
		"rootSquashConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) SkipFinalBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) SkipFinalBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) StorageCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) StorageCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) ThroughputCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) ThroughputCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) Timeouts() TfLustreFileSystem_TimeoutsPropertyOutputReference {
	var returns TfLustreFileSystem_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) WeeklyMaintenanceStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLustreFileSystem) WeeklyMaintenanceStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTimeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system aws_fsx_lustre_file_system} Resource.
// Experimental.
func NewTfLustreFileSystem(scope constructs.Construct, id *string, config *TfLustreFileSystemConfig) TfLustreFileSystem {
	_init_.Initialize()

	if err := validateNewTfLustreFileSystemParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLustreFileSystem{}

	_jsii_.Create(
		"@cdktn/aws-fsx.TfLustreFileSystem",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system aws_fsx_lustre_file_system} Resource.
// Experimental.
func NewTfLustreFileSystem_Override(t TfLustreFileSystem, scope constructs.Construct, id *string, config *TfLustreFileSystemConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.TfLustreFileSystem",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetAutoImportPolicy(val *string) {
	if err := j.validateSetAutoImportPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoImportPolicy",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetAutomaticBackupRetentionDays(val *float64) {
	if err := j.validateSetAutomaticBackupRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticBackupRetentionDays",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetBackupId(val *string) {
	if err := j.validateSetBackupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupId",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetCopyTagsToBackups(val interface{}) {
	if err := j.validateSetCopyTagsToBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToBackups",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetDailyAutomaticBackupStartTime(val *string) {
	if err := j.validateSetDailyAutomaticBackupStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailyAutomaticBackupStartTime",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetDataCompressionType(val *string) {
	if err := j.validateSetDataCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCompressionType",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetDeploymentType(val *string) {
	if err := j.validateSetDeploymentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentType",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetDriveCacheType(val *string) {
	if err := j.validateSetDriveCacheTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driveCacheType",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetEfaEnabled(val interface{}) {
	if err := j.validateSetEfaEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"efaEnabled",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetExportPath(val *string) {
	if err := j.validateSetExportPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportPath",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetFileSystemTypeVersion(val *string) {
	if err := j.validateSetFileSystemTypeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemTypeVersion",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetFinalBackupTags(val *map[string]*string) {
	if err := j.validateSetFinalBackupTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalBackupTags",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetImportedFileChunkSize(val *float64) {
	if err := j.validateSetImportedFileChunkSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importedFileChunkSize",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetImportPath(val *string) {
	if err := j.validateSetImportPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importPath",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetPerUnitStorageThroughput(val *float64) {
	if err := j.validateSetPerUnitStorageThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perUnitStorageThroughput",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetSkipFinalBackup(val interface{}) {
	if err := j.validateSetSkipFinalBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipFinalBackup",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetStorageCapacity(val *float64) {
	if err := j.validateSetStorageCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCapacity",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetThroughputCapacity(val *float64) {
	if err := j.validateSetThroughputCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughputCapacity",
		val,
	)
}

func (j *jsiiProxy_TfLustreFileSystem)SetWeeklyMaintenanceStartTime(val *string) {
	if err := j.validateSetWeeklyMaintenanceStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeklyMaintenanceStartTime",
		val,
	)
}

// Generates CDKTN code for importing a TfLustreFileSystem resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfLustreFileSystem_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfLustreFileSystem_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.TfLustreFileSystem",
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
func TfLustreFileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfLustreFileSystem_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.TfLustreFileSystem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfLustreFileSystem_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfLustreFileSystem_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.TfLustreFileSystem",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfLustreFileSystem_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfLustreFileSystem_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.TfLustreFileSystem",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfLustreFileSystem_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-fsx.TfLustreFileSystem",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfLustreFileSystem) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfLustreFileSystem) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfLustreFileSystem) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLustreFileSystem) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLustreFileSystem) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLustreFileSystem) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLustreFileSystem) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLustreFileSystem) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLustreFileSystem) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLustreFileSystem) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLustreFileSystem) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLustreFileSystem) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLustreFileSystem) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfLustreFileSystem) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLustreFileSystem) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfLustreFileSystem) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfLustreFileSystem) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfLustreFileSystem) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfLustreFileSystem) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfLustreFileSystem) PutDataReadCacheConfiguration(value *TfLustreFileSystem_DataReadCacheConfigurationProperty) {
	if err := t.validatePutDataReadCacheConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDataReadCacheConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLustreFileSystem) PutLogConfiguration(value *TfLustreFileSystem_LogConfigurationProperty) {
	if err := t.validatePutLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLogConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLustreFileSystem) PutMetadataConfiguration(value *TfLustreFileSystem_MetadataConfigurationProperty) {
	if err := t.validatePutMetadataConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMetadataConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLustreFileSystem) PutRootSquashConfiguration(value *TfLustreFileSystem_RootSquashConfigurationProperty) {
	if err := t.validatePutRootSquashConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRootSquashConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLustreFileSystem) PutTimeouts(value *TfLustreFileSystem_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLustreFileSystem) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetAutoImportPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoImportPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetAutomaticBackupRetentionDays() {
	_jsii_.InvokeVoid(
		t,
		"resetAutomaticBackupRetentionDays",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetBackupId() {
	_jsii_.InvokeVoid(
		t,
		"resetBackupId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetCopyTagsToBackups() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyTagsToBackups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetDailyAutomaticBackupStartTime() {
	_jsii_.InvokeVoid(
		t,
		"resetDailyAutomaticBackupStartTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetDataCompressionType() {
	_jsii_.InvokeVoid(
		t,
		"resetDataCompressionType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetDataReadCacheConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetDataReadCacheConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetDeploymentType() {
	_jsii_.InvokeVoid(
		t,
		"resetDeploymentType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetDriveCacheType() {
	_jsii_.InvokeVoid(
		t,
		"resetDriveCacheType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetEfaEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetEfaEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetExportPath() {
	_jsii_.InvokeVoid(
		t,
		"resetExportPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetFileSystemTypeVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetFileSystemTypeVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetFinalBackupTags() {
	_jsii_.InvokeVoid(
		t,
		"resetFinalBackupTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetImportedFileChunkSize() {
	_jsii_.InvokeVoid(
		t,
		"resetImportedFileChunkSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetImportPath() {
	_jsii_.InvokeVoid(
		t,
		"resetImportPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetLogConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetLogConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetMetadataConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetMetadataConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetPerUnitStorageThroughput() {
	_jsii_.InvokeVoid(
		t,
		"resetPerUnitStorageThroughput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetRootSquashConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetRootSquashConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetSkipFinalBackup() {
	_jsii_.InvokeVoid(
		t,
		"resetSkipFinalBackup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetStorageCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetStorageType() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetThroughputCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetThroughputCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) ResetWeeklyMaintenanceStartTime() {
	_jsii_.InvokeVoid(
		t,
		"resetWeeklyMaintenanceStartTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLustreFileSystem) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLustreFileSystem) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLustreFileSystem) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLustreFileSystem) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLustreFileSystem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLustreFileSystem) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLustreFileSystem) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

