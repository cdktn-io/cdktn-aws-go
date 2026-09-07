package fsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/fsx/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/fsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system aws_fsx_lustre_file_system}.
// Experimental.
type AwsLustreFileSystem interface {
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
	DataReadCacheConfiguration() AwsLustreFileSystem_DataReadCacheConfigurationPropertyOutputReference
	// Experimental.
	DataReadCacheConfigurationInput() *AwsLustreFileSystem_DataReadCacheConfigurationProperty
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
	LogConfiguration() AwsLustreFileSystem_LogConfigurationPropertyOutputReference
	// Experimental.
	LogConfigurationInput() *AwsLustreFileSystem_LogConfigurationProperty
	// Experimental.
	MetadataConfiguration() AwsLustreFileSystem_MetadataConfigurationPropertyOutputReference
	// Experimental.
	MetadataConfigurationInput() *AwsLustreFileSystem_MetadataConfigurationProperty
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
	RootSquashConfiguration() AwsLustreFileSystem_RootSquashConfigurationPropertyOutputReference
	// Experimental.
	RootSquashConfigurationInput() *AwsLustreFileSystem_RootSquashConfigurationProperty
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
	Timeouts() AwsLustreFileSystem_TimeoutsPropertyOutputReference
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
	PutDataReadCacheConfiguration(value *AwsLustreFileSystem_DataReadCacheConfigurationProperty)
	// Experimental.
	PutLogConfiguration(value *AwsLustreFileSystem_LogConfigurationProperty)
	// Experimental.
	PutMetadataConfiguration(value *AwsLustreFileSystem_MetadataConfigurationProperty)
	// Experimental.
	PutRootSquashConfiguration(value *AwsLustreFileSystem_RootSquashConfigurationProperty)
	// Experimental.
	PutTimeouts(value *AwsLustreFileSystem_TimeoutsProperty)
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

// The jsii proxy struct for AwsLustreFileSystem
type jsiiProxy_AwsLustreFileSystem struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsLustreFileSystem) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) AutoImportPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoImportPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) AutoImportPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoImportPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) AutomaticBackupRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) AutomaticBackupRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) BackupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) BackupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) CopyTagsToBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) CopyTagsToBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) DailyAutomaticBackupStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) DailyAutomaticBackupStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) DataCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) DataCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) DataReadCacheConfiguration() AwsLustreFileSystem_DataReadCacheConfigurationPropertyOutputReference {
	var returns AwsLustreFileSystem_DataReadCacheConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"dataReadCacheConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) DataReadCacheConfigurationInput() *AwsLustreFileSystem_DataReadCacheConfigurationProperty {
	var returns *AwsLustreFileSystem_DataReadCacheConfigurationProperty
	_jsii_.Get(
		j,
		"dataReadCacheConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) DeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) DeploymentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) DriveCacheType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driveCacheType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) DriveCacheTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driveCacheTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) EfaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efaEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) EfaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efaEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) ExportPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) ExportPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) FileSystemTypeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemTypeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) FileSystemTypeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemTypeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) FinalBackupTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"finalBackupTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) FinalBackupTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"finalBackupTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) ImportedFileChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importedFileChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) ImportedFileChunkSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importedFileChunkSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) ImportPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) ImportPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) LogConfiguration() AwsLustreFileSystem_LogConfigurationPropertyOutputReference {
	var returns AwsLustreFileSystem_LogConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) LogConfigurationInput() *AwsLustreFileSystem_LogConfigurationProperty {
	var returns *AwsLustreFileSystem_LogConfigurationProperty
	_jsii_.Get(
		j,
		"logConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) MetadataConfiguration() AwsLustreFileSystem_MetadataConfigurationPropertyOutputReference {
	var returns AwsLustreFileSystem_MetadataConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"metadataConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) MetadataConfigurationInput() *AwsLustreFileSystem_MetadataConfigurationProperty {
	var returns *AwsLustreFileSystem_MetadataConfigurationProperty
	_jsii_.Get(
		j,
		"metadataConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) MountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) NetworkInterfaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) PerUnitStorageThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perUnitStorageThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) PerUnitStorageThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perUnitStorageThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) RootSquashConfiguration() AwsLustreFileSystem_RootSquashConfigurationPropertyOutputReference {
	var returns AwsLustreFileSystem_RootSquashConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"rootSquashConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) RootSquashConfigurationInput() *AwsLustreFileSystem_RootSquashConfigurationProperty {
	var returns *AwsLustreFileSystem_RootSquashConfigurationProperty
	_jsii_.Get(
		j,
		"rootSquashConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) SkipFinalBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) SkipFinalBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) StorageCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) StorageCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) ThroughputCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) ThroughputCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) Timeouts() AwsLustreFileSystem_TimeoutsPropertyOutputReference {
	var returns AwsLustreFileSystem_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) WeeklyMaintenanceStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLustreFileSystem) WeeklyMaintenanceStartTimeInput() *string {
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
func NewAwsLustreFileSystem(scope constructs.Construct, id *string, config *AwsLustreFileSystemConfig) AwsLustreFileSystem {
	_init_.Initialize()

	if err := validateNewAwsLustreFileSystemParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLustreFileSystem{}

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsLustreFileSystem",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system aws_fsx_lustre_file_system} Resource.
// Experimental.
func NewAwsLustreFileSystem_Override(a AwsLustreFileSystem, scope constructs.Construct, id *string, config *AwsLustreFileSystemConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsLustreFileSystem",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetAutoImportPolicy(val *string) {
	if err := j.validateSetAutoImportPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoImportPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetAutomaticBackupRetentionDays(val *float64) {
	if err := j.validateSetAutomaticBackupRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticBackupRetentionDays",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetBackupId(val *string) {
	if err := j.validateSetBackupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupId",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetCopyTagsToBackups(val interface{}) {
	if err := j.validateSetCopyTagsToBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToBackups",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetDailyAutomaticBackupStartTime(val *string) {
	if err := j.validateSetDailyAutomaticBackupStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailyAutomaticBackupStartTime",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetDataCompressionType(val *string) {
	if err := j.validateSetDataCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCompressionType",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetDeploymentType(val *string) {
	if err := j.validateSetDeploymentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentType",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetDriveCacheType(val *string) {
	if err := j.validateSetDriveCacheTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driveCacheType",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetEfaEnabled(val interface{}) {
	if err := j.validateSetEfaEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"efaEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetExportPath(val *string) {
	if err := j.validateSetExportPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportPath",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetFileSystemTypeVersion(val *string) {
	if err := j.validateSetFileSystemTypeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemTypeVersion",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetFinalBackupTags(val *map[string]*string) {
	if err := j.validateSetFinalBackupTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalBackupTags",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetImportedFileChunkSize(val *float64) {
	if err := j.validateSetImportedFileChunkSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importedFileChunkSize",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetImportPath(val *string) {
	if err := j.validateSetImportPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importPath",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetPerUnitStorageThroughput(val *float64) {
	if err := j.validateSetPerUnitStorageThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perUnitStorageThroughput",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetSkipFinalBackup(val interface{}) {
	if err := j.validateSetSkipFinalBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipFinalBackup",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetStorageCapacity(val *float64) {
	if err := j.validateSetStorageCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetThroughputCapacity(val *float64) {
	if err := j.validateSetThroughputCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughputCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsLustreFileSystem)SetWeeklyMaintenanceStartTime(val *string) {
	if err := j.validateSetWeeklyMaintenanceStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeklyMaintenanceStartTime",
		val,
	)
}

// Generates CDKTN code for importing a AwsLustreFileSystem resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsLustreFileSystem_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsLustreFileSystem_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.AwsLustreFileSystem",
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
func AwsLustreFileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsLustreFileSystem_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.AwsLustreFileSystem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsLustreFileSystem_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsLustreFileSystem_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.AwsLustreFileSystem",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsLustreFileSystem_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsLustreFileSystem_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.AwsLustreFileSystem",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsLustreFileSystem_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-fsx.AwsLustreFileSystem",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsLustreFileSystem) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLustreFileSystem) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLustreFileSystem) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLustreFileSystem) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLustreFileSystem) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLustreFileSystem) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLustreFileSystem) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLustreFileSystem) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLustreFileSystem) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLustreFileSystem) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLustreFileSystem) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLustreFileSystem) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsLustreFileSystem) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) PutDataReadCacheConfiguration(value *AwsLustreFileSystem_DataReadCacheConfigurationProperty) {
	if err := a.validatePutDataReadCacheConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataReadCacheConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) PutLogConfiguration(value *AwsLustreFileSystem_LogConfigurationProperty) {
	if err := a.validatePutLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) PutMetadataConfiguration(value *AwsLustreFileSystem_MetadataConfigurationProperty) {
	if err := a.validatePutMetadataConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetadataConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) PutRootSquashConfiguration(value *AwsLustreFileSystem_RootSquashConfigurationProperty) {
	if err := a.validatePutRootSquashConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRootSquashConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) PutTimeouts(value *AwsLustreFileSystem_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetAutoImportPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoImportPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetAutomaticBackupRetentionDays() {
	_jsii_.InvokeVoid(
		a,
		"resetAutomaticBackupRetentionDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetBackupId() {
	_jsii_.InvokeVoid(
		a,
		"resetBackupId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetCopyTagsToBackups() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyTagsToBackups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetDailyAutomaticBackupStartTime() {
	_jsii_.InvokeVoid(
		a,
		"resetDailyAutomaticBackupStartTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetDataCompressionType() {
	_jsii_.InvokeVoid(
		a,
		"resetDataCompressionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetDataReadCacheConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetDataReadCacheConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetDeploymentType() {
	_jsii_.InvokeVoid(
		a,
		"resetDeploymentType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetDriveCacheType() {
	_jsii_.InvokeVoid(
		a,
		"resetDriveCacheType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetEfaEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEfaEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetExportPath() {
	_jsii_.InvokeVoid(
		a,
		"resetExportPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetFileSystemTypeVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetFileSystemTypeVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetFinalBackupTags() {
	_jsii_.InvokeVoid(
		a,
		"resetFinalBackupTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetImportedFileChunkSize() {
	_jsii_.InvokeVoid(
		a,
		"resetImportedFileChunkSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetImportPath() {
	_jsii_.InvokeVoid(
		a,
		"resetImportPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetLogConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLogConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetMetadataConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadataConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetPerUnitStorageThroughput() {
	_jsii_.InvokeVoid(
		a,
		"resetPerUnitStorageThroughput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetRootSquashConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRootSquashConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetSkipFinalBackup() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipFinalBackup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetStorageCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetStorageType() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetThroughputCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetThroughputCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) ResetWeeklyMaintenanceStartTime() {
	_jsii_.InvokeVoid(
		a,
		"resetWeeklyMaintenanceStartTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLustreFileSystem) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLustreFileSystem) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLustreFileSystem) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLustreFileSystem) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLustreFileSystem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLustreFileSystem) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLustreFileSystem) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

