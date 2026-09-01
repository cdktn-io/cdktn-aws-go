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
type AwsFsxLustreFileSystem interface {
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
	DataReadCacheConfiguration() AwsFsxLustreFileSystem_DataReadCacheConfigurationPropertyOutputReference
	// Experimental.
	DataReadCacheConfigurationInput() *AwsFsxLustreFileSystem_DataReadCacheConfigurationProperty
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
	LogConfiguration() AwsFsxLustreFileSystem_LogConfigurationPropertyOutputReference
	// Experimental.
	LogConfigurationInput() *AwsFsxLustreFileSystem_LogConfigurationProperty
	// Experimental.
	MetadataConfiguration() AwsFsxLustreFileSystem_MetadataConfigurationPropertyOutputReference
	// Experimental.
	MetadataConfigurationInput() *AwsFsxLustreFileSystem_MetadataConfigurationProperty
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
	RootSquashConfiguration() AwsFsxLustreFileSystem_RootSquashConfigurationPropertyOutputReference
	// Experimental.
	RootSquashConfigurationInput() *AwsFsxLustreFileSystem_RootSquashConfigurationProperty
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
	Timeouts() AwsFsxLustreFileSystem_TimeoutsPropertyOutputReference
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
	PutDataReadCacheConfiguration(value *AwsFsxLustreFileSystem_DataReadCacheConfigurationProperty)
	// Experimental.
	PutLogConfiguration(value *AwsFsxLustreFileSystem_LogConfigurationProperty)
	// Experimental.
	PutMetadataConfiguration(value *AwsFsxLustreFileSystem_MetadataConfigurationProperty)
	// Experimental.
	PutRootSquashConfiguration(value *AwsFsxLustreFileSystem_RootSquashConfigurationProperty)
	// Experimental.
	PutTimeouts(value *AwsFsxLustreFileSystem_TimeoutsProperty)
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

// The jsii proxy struct for AwsFsxLustreFileSystem
type jsiiProxy_AwsFsxLustreFileSystem struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) AutoImportPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoImportPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) AutoImportPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoImportPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) AutomaticBackupRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) AutomaticBackupRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automaticBackupRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) BackupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) BackupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) CopyTagsToBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) CopyTagsToBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) DailyAutomaticBackupStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) DailyAutomaticBackupStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dailyAutomaticBackupStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) DataCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) DataCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) DataReadCacheConfiguration() AwsFsxLustreFileSystem_DataReadCacheConfigurationPropertyOutputReference {
	var returns AwsFsxLustreFileSystem_DataReadCacheConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"dataReadCacheConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) DataReadCacheConfigurationInput() *AwsFsxLustreFileSystem_DataReadCacheConfigurationProperty {
	var returns *AwsFsxLustreFileSystem_DataReadCacheConfigurationProperty
	_jsii_.Get(
		j,
		"dataReadCacheConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) DeploymentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) DeploymentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) DriveCacheType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driveCacheType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) DriveCacheTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driveCacheTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) EfaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efaEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) EfaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efaEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) ExportPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) ExportPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) FileSystemTypeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemTypeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) FileSystemTypeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemTypeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) FinalBackupTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"finalBackupTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) FinalBackupTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"finalBackupTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) ImportedFileChunkSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importedFileChunkSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) ImportedFileChunkSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importedFileChunkSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) ImportPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) ImportPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) LogConfiguration() AwsFsxLustreFileSystem_LogConfigurationPropertyOutputReference {
	var returns AwsFsxLustreFileSystem_LogConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) LogConfigurationInput() *AwsFsxLustreFileSystem_LogConfigurationProperty {
	var returns *AwsFsxLustreFileSystem_LogConfigurationProperty
	_jsii_.Get(
		j,
		"logConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) MetadataConfiguration() AwsFsxLustreFileSystem_MetadataConfigurationPropertyOutputReference {
	var returns AwsFsxLustreFileSystem_MetadataConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"metadataConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) MetadataConfigurationInput() *AwsFsxLustreFileSystem_MetadataConfigurationProperty {
	var returns *AwsFsxLustreFileSystem_MetadataConfigurationProperty
	_jsii_.Get(
		j,
		"metadataConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) MountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) NetworkInterfaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) OwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) PerUnitStorageThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perUnitStorageThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) PerUnitStorageThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perUnitStorageThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) RootSquashConfiguration() AwsFsxLustreFileSystem_RootSquashConfigurationPropertyOutputReference {
	var returns AwsFsxLustreFileSystem_RootSquashConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"rootSquashConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) RootSquashConfigurationInput() *AwsFsxLustreFileSystem_RootSquashConfigurationProperty {
	var returns *AwsFsxLustreFileSystem_RootSquashConfigurationProperty
	_jsii_.Get(
		j,
		"rootSquashConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) SkipFinalBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) SkipFinalBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) StorageCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) StorageCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) ThroughputCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) ThroughputCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"throughputCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) Timeouts() AwsFsxLustreFileSystem_TimeoutsPropertyOutputReference {
	var returns AwsFsxLustreFileSystem_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) WeeklyMaintenanceStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"weeklyMaintenanceStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxLustreFileSystem) WeeklyMaintenanceStartTimeInput() *string {
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
func NewAwsFsxLustreFileSystem(scope constructs.Construct, id *string, config *AwsFsxLustreFileSystemConfig) AwsFsxLustreFileSystem {
	_init_.Initialize()

	if err := validateNewAwsFsxLustreFileSystemParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFsxLustreFileSystem{}

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsFsxLustreFileSystem",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system aws_fsx_lustre_file_system} Resource.
// Experimental.
func NewAwsFsxLustreFileSystem_Override(a AwsFsxLustreFileSystem, scope constructs.Construct, id *string, config *AwsFsxLustreFileSystemConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsFsxLustreFileSystem",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetAutoImportPolicy(val *string) {
	if err := j.validateSetAutoImportPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoImportPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetAutomaticBackupRetentionDays(val *float64) {
	if err := j.validateSetAutomaticBackupRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticBackupRetentionDays",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetBackupId(val *string) {
	if err := j.validateSetBackupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupId",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetCopyTagsToBackups(val interface{}) {
	if err := j.validateSetCopyTagsToBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToBackups",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetDailyAutomaticBackupStartTime(val *string) {
	if err := j.validateSetDailyAutomaticBackupStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dailyAutomaticBackupStartTime",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetDataCompressionType(val *string) {
	if err := j.validateSetDataCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCompressionType",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetDeploymentType(val *string) {
	if err := j.validateSetDeploymentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentType",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetDriveCacheType(val *string) {
	if err := j.validateSetDriveCacheTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driveCacheType",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetEfaEnabled(val interface{}) {
	if err := j.validateSetEfaEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"efaEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetExportPath(val *string) {
	if err := j.validateSetExportPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportPath",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetFileSystemTypeVersion(val *string) {
	if err := j.validateSetFileSystemTypeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileSystemTypeVersion",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetFinalBackupTags(val *map[string]*string) {
	if err := j.validateSetFinalBackupTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalBackupTags",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetImportedFileChunkSize(val *float64) {
	if err := j.validateSetImportedFileChunkSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importedFileChunkSize",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetImportPath(val *string) {
	if err := j.validateSetImportPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importPath",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetPerUnitStorageThroughput(val *float64) {
	if err := j.validateSetPerUnitStorageThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perUnitStorageThroughput",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetSkipFinalBackup(val interface{}) {
	if err := j.validateSetSkipFinalBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipFinalBackup",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetStorageCapacity(val *float64) {
	if err := j.validateSetStorageCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetThroughputCapacity(val *float64) {
	if err := j.validateSetThroughputCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"throughputCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsFsxLustreFileSystem)SetWeeklyMaintenanceStartTime(val *string) {
	if err := j.validateSetWeeklyMaintenanceStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeklyMaintenanceStartTime",
		val,
	)
}

// Generates CDKTN code for importing a AwsFsxLustreFileSystem resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsFsxLustreFileSystem_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsFsxLustreFileSystem_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.AwsFsxLustreFileSystem",
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
func AwsFsxLustreFileSystem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsFsxLustreFileSystem_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.AwsFsxLustreFileSystem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsFsxLustreFileSystem_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsFsxLustreFileSystem_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.AwsFsxLustreFileSystem",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsFsxLustreFileSystem_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsFsxLustreFileSystem_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-fsx.AwsFsxLustreFileSystem",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsFsxLustreFileSystem_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-fsx.AwsFsxLustreFileSystem",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFsxLustreFileSystem) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFsxLustreFileSystem) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFsxLustreFileSystem) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFsxLustreFileSystem) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFsxLustreFileSystem) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFsxLustreFileSystem) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFsxLustreFileSystem) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFsxLustreFileSystem) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFsxLustreFileSystem) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFsxLustreFileSystem) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsFsxLustreFileSystem) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) PutDataReadCacheConfiguration(value *AwsFsxLustreFileSystem_DataReadCacheConfigurationProperty) {
	if err := a.validatePutDataReadCacheConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataReadCacheConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) PutLogConfiguration(value *AwsFsxLustreFileSystem_LogConfigurationProperty) {
	if err := a.validatePutLogConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) PutMetadataConfiguration(value *AwsFsxLustreFileSystem_MetadataConfigurationProperty) {
	if err := a.validatePutMetadataConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetadataConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) PutRootSquashConfiguration(value *AwsFsxLustreFileSystem_RootSquashConfigurationProperty) {
	if err := a.validatePutRootSquashConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRootSquashConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) PutTimeouts(value *AwsFsxLustreFileSystem_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetAutoImportPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoImportPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetAutomaticBackupRetentionDays() {
	_jsii_.InvokeVoid(
		a,
		"resetAutomaticBackupRetentionDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetBackupId() {
	_jsii_.InvokeVoid(
		a,
		"resetBackupId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetCopyTagsToBackups() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyTagsToBackups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetDailyAutomaticBackupStartTime() {
	_jsii_.InvokeVoid(
		a,
		"resetDailyAutomaticBackupStartTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetDataCompressionType() {
	_jsii_.InvokeVoid(
		a,
		"resetDataCompressionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetDataReadCacheConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetDataReadCacheConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetDeploymentType() {
	_jsii_.InvokeVoid(
		a,
		"resetDeploymentType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetDriveCacheType() {
	_jsii_.InvokeVoid(
		a,
		"resetDriveCacheType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetEfaEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEfaEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetExportPath() {
	_jsii_.InvokeVoid(
		a,
		"resetExportPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetFileSystemTypeVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetFileSystemTypeVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetFinalBackupTags() {
	_jsii_.InvokeVoid(
		a,
		"resetFinalBackupTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetImportedFileChunkSize() {
	_jsii_.InvokeVoid(
		a,
		"resetImportedFileChunkSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetImportPath() {
	_jsii_.InvokeVoid(
		a,
		"resetImportPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetLogConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLogConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetMetadataConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadataConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetPerUnitStorageThroughput() {
	_jsii_.InvokeVoid(
		a,
		"resetPerUnitStorageThroughput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetRootSquashConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRootSquashConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetSkipFinalBackup() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipFinalBackup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetStorageCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetStorageType() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetThroughputCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetThroughputCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ResetWeeklyMaintenanceStartTime() {
	_jsii_.InvokeVoid(
		a,
		"resetWeeklyMaintenanceStartTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxLustreFileSystem) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

