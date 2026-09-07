package rds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/rds/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/rds/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance aws_db_instance}.
// Experimental.
type AwsDbInstance interface {
	cdktn.TerraformResource
	// Experimental.
	Address() *string
	// Experimental.
	AllocatedStorage() *float64
	// Experimental.
	SetAllocatedStorage(val *float64)
	// Experimental.
	AllocatedStorageInput() *float64
	// Experimental.
	AllowMajorVersionUpgrade() interface{}
	// Experimental.
	SetAllowMajorVersionUpgrade(val interface{})
	// Experimental.
	AllowMajorVersionUpgradeInput() interface{}
	// Experimental.
	ApplyImmediately() interface{}
	// Experimental.
	SetApplyImmediately(val interface{})
	// Experimental.
	ApplyImmediatelyInput() interface{}
	// Experimental.
	Arn() *string
	// Experimental.
	AutoMinorVersionUpgrade() interface{}
	// Experimental.
	SetAutoMinorVersionUpgrade(val interface{})
	// Experimental.
	AutoMinorVersionUpgradeInput() interface{}
	// Experimental.
	AvailabilityZone() *string
	// Experimental.
	SetAvailabilityZone(val *string)
	// Experimental.
	AvailabilityZoneInput() *string
	// Experimental.
	BackupRetentionPeriod() *float64
	// Experimental.
	SetBackupRetentionPeriod(val *float64)
	// Experimental.
	BackupRetentionPeriodInput() *float64
	// Experimental.
	BackupTarget() *string
	// Experimental.
	SetBackupTarget(val *string)
	// Experimental.
	BackupTargetInput() *string
	// Experimental.
	BackupWindow() *string
	// Experimental.
	SetBackupWindow(val *string)
	// Experimental.
	BackupWindowInput() *string
	// Experimental.
	BlueGreenUpdate() AwsDbInstance_BlueGreenUpdatePropertyOutputReference
	// Experimental.
	BlueGreenUpdateInput() *AwsDbInstance_BlueGreenUpdateProperty
	// Experimental.
	CaCertIdentifier() *string
	// Experimental.
	SetCaCertIdentifier(val *string)
	// Experimental.
	CaCertIdentifierInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CharacterSetName() *string
	// Experimental.
	SetCharacterSetName(val *string)
	// Experimental.
	CharacterSetNameInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	CopyTagsToSnapshot() interface{}
	// Experimental.
	SetCopyTagsToSnapshot(val interface{})
	// Experimental.
	CopyTagsToSnapshotInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CustomerOwnedIpEnabled() interface{}
	// Experimental.
	SetCustomerOwnedIpEnabled(val interface{})
	// Experimental.
	CustomerOwnedIpEnabledInput() interface{}
	// Experimental.
	CustomIamInstanceProfile() *string
	// Experimental.
	SetCustomIamInstanceProfile(val *string)
	// Experimental.
	CustomIamInstanceProfileInput() *string
	// Experimental.
	DatabaseInsightsMode() *string
	// Experimental.
	SetDatabaseInsightsMode(val *string)
	// Experimental.
	DatabaseInsightsModeInput() *string
	// Experimental.
	DbName() *string
	// Experimental.
	SetDbName(val *string)
	// Experimental.
	DbNameInput() *string
	// Experimental.
	DbSubnetGroupName() *string
	// Experimental.
	SetDbSubnetGroupName(val *string)
	// Experimental.
	DbSubnetGroupNameInput() *string
	// Experimental.
	DedicatedLogVolume() interface{}
	// Experimental.
	SetDedicatedLogVolume(val interface{})
	// Experimental.
	DedicatedLogVolumeInput() interface{}
	// Experimental.
	DeleteAutomatedBackups() interface{}
	// Experimental.
	SetDeleteAutomatedBackups(val interface{})
	// Experimental.
	DeleteAutomatedBackupsInput() interface{}
	// Experimental.
	DeletionProtection() interface{}
	// Experimental.
	SetDeletionProtection(val interface{})
	// Experimental.
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	Domain() *string
	// Experimental.
	SetDomain(val *string)
	// Experimental.
	DomainAuthSecretArn() *string
	// Experimental.
	SetDomainAuthSecretArn(val *string)
	// Experimental.
	DomainAuthSecretArnInput() *string
	// Experimental.
	DomainDnsIps() *[]*string
	// Experimental.
	SetDomainDnsIps(val *[]*string)
	// Experimental.
	DomainDnsIpsInput() *[]*string
	// Experimental.
	DomainFqdn() *string
	// Experimental.
	SetDomainFqdn(val *string)
	// Experimental.
	DomainFqdnInput() *string
	// Experimental.
	DomainIamRoleName() *string
	// Experimental.
	SetDomainIamRoleName(val *string)
	// Experimental.
	DomainIamRoleNameInput() *string
	// Experimental.
	DomainInput() *string
	// Experimental.
	DomainOu() *string
	// Experimental.
	SetDomainOu(val *string)
	// Experimental.
	DomainOuInput() *string
	// Experimental.
	EnabledCloudwatchLogsExports() *[]*string
	// Experimental.
	SetEnabledCloudwatchLogsExports(val *[]*string)
	// Experimental.
	EnabledCloudwatchLogsExportsInput() *[]*string
	// Experimental.
	Endpoint() *string
	// Experimental.
	Engine() *string
	// Experimental.
	SetEngine(val *string)
	// Experimental.
	EngineInput() *string
	// Experimental.
	EngineLifecycleSupport() *string
	// Experimental.
	SetEngineLifecycleSupport(val *string)
	// Experimental.
	EngineLifecycleSupportInput() *string
	// Experimental.
	EngineVersion() *string
	// Experimental.
	SetEngineVersion(val *string)
	// Experimental.
	EngineVersionActual() *string
	// Experimental.
	EngineVersionInput() *string
	// Experimental.
	FinalSnapshotIdentifier() *string
	// Experimental.
	SetFinalSnapshotIdentifier(val *string)
	// Experimental.
	FinalSnapshotIdentifierInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HostedZoneId() *string
	// Experimental.
	IamDatabaseAuthenticationEnabled() interface{}
	// Experimental.
	SetIamDatabaseAuthenticationEnabled(val interface{})
	// Experimental.
	IamDatabaseAuthenticationEnabledInput() interface{}
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	Identifier() *string
	// Experimental.
	SetIdentifier(val *string)
	// Experimental.
	IdentifierInput() *string
	// Experimental.
	IdentifierPrefix() *string
	// Experimental.
	SetIdentifierPrefix(val *string)
	// Experimental.
	IdentifierPrefixInput() *string
	// Experimental.
	IdInput() *string
	// Experimental.
	InstanceClass() *string
	// Experimental.
	SetInstanceClass(val *string)
	// Experimental.
	InstanceClassInput() *string
	// Experimental.
	Iops() *float64
	// Experimental.
	SetIops(val *float64)
	// Experimental.
	IopsInput() *float64
	// Experimental.
	KmsKeyId() *string
	// Experimental.
	SetKmsKeyId(val *string)
	// Experimental.
	KmsKeyIdInput() *string
	// Experimental.
	LatestRestorableTime() *string
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
	ListenerEndpoint() AwsDbInstance_ListenerEndpointPropertyList
	// Experimental.
	MaintenanceWindow() *string
	// Experimental.
	SetMaintenanceWindow(val *string)
	// Experimental.
	MaintenanceWindowInput() *string
	// Experimental.
	ManageMasterUserPassword() interface{}
	// Experimental.
	SetManageMasterUserPassword(val interface{})
	// Experimental.
	ManageMasterUserPasswordInput() interface{}
	// Experimental.
	MasterUserSecret() AwsDbInstance_MasterUserSecretPropertyList
	// Experimental.
	MasterUserSecretKmsKeyId() *string
	// Experimental.
	SetMasterUserSecretKmsKeyId(val *string)
	// Experimental.
	MasterUserSecretKmsKeyIdInput() *string
	// Experimental.
	MaxAllocatedStorage() *float64
	// Experimental.
	SetMaxAllocatedStorage(val *float64)
	// Experimental.
	MaxAllocatedStorageInput() *float64
	// Experimental.
	MonitoringInterval() *float64
	// Experimental.
	SetMonitoringInterval(val *float64)
	// Experimental.
	MonitoringIntervalInput() *float64
	// Experimental.
	MonitoringRoleArn() *string
	// Experimental.
	SetMonitoringRoleArn(val *string)
	// Experimental.
	MonitoringRoleArnInput() *string
	// Experimental.
	MultiAz() interface{}
	// Experimental.
	SetMultiAz(val interface{})
	// Experimental.
	MultiAzInput() interface{}
	// Experimental.
	NcharCharacterSetName() *string
	// Experimental.
	SetNcharCharacterSetName(val *string)
	// Experimental.
	NcharCharacterSetNameInput() *string
	// Experimental.
	NetworkType() *string
	// Experimental.
	SetNetworkType(val *string)
	// Experimental.
	NetworkTypeInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OptionGroupName() *string
	// Experimental.
	SetOptionGroupName(val *string)
	// Experimental.
	OptionGroupNameInput() *string
	// Experimental.
	ParameterGroupName() *string
	// Experimental.
	SetParameterGroupName(val *string)
	// Experimental.
	ParameterGroupNameInput() *string
	// Experimental.
	Password() *string
	// Experimental.
	SetPassword(val *string)
	// Experimental.
	PasswordInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	PasswordWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetPasswordWo(val *string)
	// Experimental.
	PasswordWoInput() *string
	// Experimental.
	PasswordWoVersion() *float64
	// Experimental.
	SetPasswordWoVersion(val *float64)
	// Experimental.
	PasswordWoVersionInput() *float64
	// Experimental.
	PerformanceInsightsEnabled() interface{}
	// Experimental.
	SetPerformanceInsightsEnabled(val interface{})
	// Experimental.
	PerformanceInsightsEnabledInput() interface{}
	// Experimental.
	PerformanceInsightsKmsKeyId() *string
	// Experimental.
	SetPerformanceInsightsKmsKeyId(val *string)
	// Experimental.
	PerformanceInsightsKmsKeyIdInput() *string
	// Experimental.
	PerformanceInsightsRetentionPeriod() *float64
	// Experimental.
	SetPerformanceInsightsRetentionPeriod(val *float64)
	// Experimental.
	PerformanceInsightsRetentionPeriodInput() *float64
	// Experimental.
	Port() *float64
	// Experimental.
	SetPort(val *float64)
	// Experimental.
	PortInput() *float64
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	PubliclyAccessible() interface{}
	// Experimental.
	SetPubliclyAccessible(val interface{})
	// Experimental.
	PubliclyAccessibleInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	ReplicaMode() *string
	// Experimental.
	SetReplicaMode(val *string)
	// Experimental.
	ReplicaModeInput() *string
	// Experimental.
	Replicas() *[]*string
	// Experimental.
	ReplicateSourceDb() *string
	// Experimental.
	SetReplicateSourceDb(val *string)
	// Experimental.
	ReplicateSourceDbInput() *string
	// Experimental.
	ResourceId() *string
	// Experimental.
	RestoreToPointInTime() AwsDbInstance_RestoreToPointInTimePropertyOutputReference
	// Experimental.
	RestoreToPointInTimeInput() *AwsDbInstance_RestoreToPointInTimeProperty
	// Experimental.
	S3Import() AwsDbInstance_S3ImportPropertyOutputReference
	// Experimental.
	S3ImportInput() *AwsDbInstance_S3ImportProperty
	// Experimental.
	SkipFinalSnapshot() interface{}
	// Experimental.
	SetSkipFinalSnapshot(val interface{})
	// Experimental.
	SkipFinalSnapshotInput() interface{}
	// Experimental.
	SnapshotIdentifier() *string
	// Experimental.
	SetSnapshotIdentifier(val *string)
	// Experimental.
	SnapshotIdentifierInput() *string
	// Experimental.
	Status() *string
	// Experimental.
	StorageEncrypted() interface{}
	// Experimental.
	SetStorageEncrypted(val interface{})
	// Experimental.
	StorageEncryptedInput() interface{}
	// Experimental.
	StorageThroughput() *float64
	// Experimental.
	SetStorageThroughput(val *float64)
	// Experimental.
	StorageThroughputInput() *float64
	// Experimental.
	StorageType() *string
	// Experimental.
	SetStorageType(val *string)
	// Experimental.
	StorageTypeInput() *string
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
	Timeouts() AwsDbInstance_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Timezone() *string
	// Experimental.
	SetTimezone(val *string)
	// Experimental.
	TimezoneInput() *string
	// Experimental.
	UpgradeRolloutOrder() *string
	// Experimental.
	UpgradeStorageConfig() interface{}
	// Experimental.
	SetUpgradeStorageConfig(val interface{})
	// Experimental.
	UpgradeStorageConfigInput() interface{}
	// Experimental.
	Username() *string
	// Experimental.
	SetUsername(val *string)
	// Experimental.
	UsernameInput() *string
	// Experimental.
	VpcSecurityGroupIds() *[]*string
	// Experimental.
	SetVpcSecurityGroupIds(val *[]*string)
	// Experimental.
	VpcSecurityGroupIdsInput() *[]*string
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
	PutBlueGreenUpdate(value *AwsDbInstance_BlueGreenUpdateProperty)
	// Experimental.
	PutRestoreToPointInTime(value *AwsDbInstance_RestoreToPointInTimeProperty)
	// Experimental.
	PutS3Import(value *AwsDbInstance_S3ImportProperty)
	// Experimental.
	PutTimeouts(value *AwsDbInstance_TimeoutsProperty)
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
	ResetAllocatedStorage()
	// Experimental.
	ResetAllowMajorVersionUpgrade()
	// Experimental.
	ResetApplyImmediately()
	// Experimental.
	ResetAutoMinorVersionUpgrade()
	// Experimental.
	ResetAvailabilityZone()
	// Experimental.
	ResetBackupRetentionPeriod()
	// Experimental.
	ResetBackupTarget()
	// Experimental.
	ResetBackupWindow()
	// Experimental.
	ResetBlueGreenUpdate()
	// Experimental.
	ResetCaCertIdentifier()
	// Experimental.
	ResetCharacterSetName()
	// Experimental.
	ResetCopyTagsToSnapshot()
	// Experimental.
	ResetCustomerOwnedIpEnabled()
	// Experimental.
	ResetCustomIamInstanceProfile()
	// Experimental.
	ResetDatabaseInsightsMode()
	// Experimental.
	ResetDbName()
	// Experimental.
	ResetDbSubnetGroupName()
	// Experimental.
	ResetDedicatedLogVolume()
	// Experimental.
	ResetDeleteAutomatedBackups()
	// Experimental.
	ResetDeletionProtection()
	// Experimental.
	ResetDomain()
	// Experimental.
	ResetDomainAuthSecretArn()
	// Experimental.
	ResetDomainDnsIps()
	// Experimental.
	ResetDomainFqdn()
	// Experimental.
	ResetDomainIamRoleName()
	// Experimental.
	ResetDomainOu()
	// Experimental.
	ResetEnabledCloudwatchLogsExports()
	// Experimental.
	ResetEngine()
	// Experimental.
	ResetEngineLifecycleSupport()
	// Experimental.
	ResetEngineVersion()
	// Experimental.
	ResetFinalSnapshotIdentifier()
	// Experimental.
	ResetIamDatabaseAuthenticationEnabled()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIdentifier()
	// Experimental.
	ResetIdentifierPrefix()
	// Experimental.
	ResetIops()
	// Experimental.
	ResetKmsKeyId()
	// Experimental.
	ResetLicenseModel()
	// Experimental.
	ResetMaintenanceWindow()
	// Experimental.
	ResetManageMasterUserPassword()
	// Experimental.
	ResetMasterUserSecretKmsKeyId()
	// Experimental.
	ResetMaxAllocatedStorage()
	// Experimental.
	ResetMonitoringInterval()
	// Experimental.
	ResetMonitoringRoleArn()
	// Experimental.
	ResetMultiAz()
	// Experimental.
	ResetNcharCharacterSetName()
	// Experimental.
	ResetNetworkType()
	// Experimental.
	ResetOptionGroupName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetParameterGroupName()
	// Experimental.
	ResetPassword()
	// Experimental.
	ResetPasswordWo()
	// Experimental.
	ResetPasswordWoVersion()
	// Experimental.
	ResetPerformanceInsightsEnabled()
	// Experimental.
	ResetPerformanceInsightsKmsKeyId()
	// Experimental.
	ResetPerformanceInsightsRetentionPeriod()
	// Experimental.
	ResetPort()
	// Experimental.
	ResetPubliclyAccessible()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetReplicaMode()
	// Experimental.
	ResetReplicateSourceDb()
	// Experimental.
	ResetRestoreToPointInTime()
	// Experimental.
	ResetS3Import()
	// Experimental.
	ResetSkipFinalSnapshot()
	// Experimental.
	ResetSnapshotIdentifier()
	// Experimental.
	ResetStorageEncrypted()
	// Experimental.
	ResetStorageThroughput()
	// Experimental.
	ResetStorageType()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTimezone()
	// Experimental.
	ResetUpgradeStorageConfig()
	// Experimental.
	ResetUsername()
	// Experimental.
	ResetVpcSecurityGroupIds()
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

// The jsii proxy struct for AwsDbInstance
type jsiiProxy_AwsDbInstance struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsDbInstance) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) AllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) AllocatedStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) AllowMajorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) AllowMajorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) ApplyImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) ApplyImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) AutoMinorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) BackupRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) BackupRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) BackupTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) BackupTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) BackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) BackupWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) BlueGreenUpdate() AwsDbInstance_BlueGreenUpdatePropertyOutputReference {
	var returns AwsDbInstance_BlueGreenUpdatePropertyOutputReference
	_jsii_.Get(
		j,
		"blueGreenUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) BlueGreenUpdateInput() *AwsDbInstance_BlueGreenUpdateProperty {
	var returns *AwsDbInstance_BlueGreenUpdateProperty
	_jsii_.Get(
		j,
		"blueGreenUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) CaCertIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) CaCertIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) CharacterSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) CharacterSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) CopyTagsToSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) CopyTagsToSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) CustomerOwnedIpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerOwnedIpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) CustomerOwnedIpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerOwnedIpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) CustomIamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customIamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) CustomIamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customIamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DatabaseInsightsMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInsightsMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DatabaseInsightsModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInsightsModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DbNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DbSubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DedicatedLogVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedLogVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DedicatedLogVolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedLogVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DeleteAutomatedBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DeleteAutomatedBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DomainAuthSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAuthSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DomainAuthSecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAuthSecretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DomainDnsIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainDnsIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DomainDnsIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainDnsIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DomainFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DomainFqdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainFqdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DomainIamRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DomainIamRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DomainOu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainOu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) DomainOuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainOuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) EnabledCloudwatchLogsExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledCloudwatchLogsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) EnabledCloudwatchLogsExportsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledCloudwatchLogsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) EngineLifecycleSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineLifecycleSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) EngineLifecycleSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineLifecycleSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) EngineVersionActual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionActual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) FinalSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) FinalSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) IamDatabaseAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamDatabaseAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) IamDatabaseAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamDatabaseAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) IdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) IdentifierPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) IdentifierPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) InstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) InstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) LatestRestorableTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestRestorableTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) ListenerEndpoint() AwsDbInstance_ListenerEndpointPropertyList {
	var returns AwsDbInstance_ListenerEndpointPropertyList
	_jsii_.Get(
		j,
		"listenerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) MaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) MaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) ManageMasterUserPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) ManageMasterUserPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) MasterUserSecret() AwsDbInstance_MasterUserSecretPropertyList {
	var returns AwsDbInstance_MasterUserSecretPropertyList
	_jsii_.Get(
		j,
		"masterUserSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) MasterUserSecretKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserSecretKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) MasterUserSecretKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserSecretKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) MaxAllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) MaxAllocatedStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) MonitoringInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) MonitoringIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) MonitoringRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) MonitoringRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) MultiAz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) MultiAzInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) NcharCharacterSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharCharacterSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) NcharCharacterSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharCharacterSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) OptionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) OptionGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) ParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) ParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) PasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) PasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) PasswordWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) PasswordWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) PerformanceInsightsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceInsightsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) PerformanceInsightsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceInsightsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) PerformanceInsightsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) PerformanceInsightsKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) PerformanceInsightsRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) PerformanceInsightsRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) PubliclyAccessibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) ReplicaMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) ReplicaModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Replicas() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) ReplicateSourceDb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicateSourceDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) ReplicateSourceDbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicateSourceDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) RestoreToPointInTime() AwsDbInstance_RestoreToPointInTimePropertyOutputReference {
	var returns AwsDbInstance_RestoreToPointInTimePropertyOutputReference
	_jsii_.Get(
		j,
		"restoreToPointInTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) RestoreToPointInTimeInput() *AwsDbInstance_RestoreToPointInTimeProperty {
	var returns *AwsDbInstance_RestoreToPointInTimeProperty
	_jsii_.Get(
		j,
		"restoreToPointInTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) S3Import() AwsDbInstance_S3ImportPropertyOutputReference {
	var returns AwsDbInstance_S3ImportPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Import",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) S3ImportInput() *AwsDbInstance_S3ImportProperty {
	var returns *AwsDbInstance_S3ImportProperty
	_jsii_.Get(
		j,
		"s3ImportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) SkipFinalSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) SkipFinalSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) SnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) SnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) StorageEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) StorageEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) StorageThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) StorageThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Timeouts() AwsDbInstance_TimeoutsPropertyOutputReference {
	var returns AwsDbInstance_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) UpgradeRolloutOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeRolloutOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) UpgradeStorageConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upgradeStorageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) UpgradeStorageConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upgradeStorageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbInstance) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance aws_db_instance} Resource.
// Experimental.
func NewAwsDbInstance(scope constructs.Construct, id *string, config *AwsDbInstanceConfig) AwsDbInstance {
	_init_.Initialize()

	if err := validateNewAwsDbInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDbInstance{}

	_jsii_.Create(
		"@cdktn/aws-rds.AwsDbInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance aws_db_instance} Resource.
// Experimental.
func NewAwsDbInstance_Override(a AwsDbInstance, scope constructs.Construct, id *string, config *AwsDbInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-rds.AwsDbInstance",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetAllocatedStorage(val *float64) {
	if err := j.validateSetAllocatedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocatedStorage",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetAllowMajorVersionUpgrade(val interface{}) {
	if err := j.validateSetAllowMajorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMajorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetApplyImmediately(val interface{}) {
	if err := j.validateSetApplyImmediatelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyImmediately",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetAutoMinorVersionUpgrade(val interface{}) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetBackupRetentionPeriod(val *float64) {
	if err := j.validateSetBackupRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetBackupTarget(val *string) {
	if err := j.validateSetBackupTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupTarget",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetBackupWindow(val *string) {
	if err := j.validateSetBackupWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupWindow",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetCaCertIdentifier(val *string) {
	if err := j.validateSetCaCertIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetCharacterSetName(val *string) {
	if err := j.validateSetCharacterSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"characterSetName",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetCopyTagsToSnapshot(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshot",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetCustomerOwnedIpEnabled(val interface{}) {
	if err := j.validateSetCustomerOwnedIpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerOwnedIpEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetCustomIamInstanceProfile(val *string) {
	if err := j.validateSetCustomIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customIamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetDatabaseInsightsMode(val *string) {
	if err := j.validateSetDatabaseInsightsModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseInsightsMode",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetDbName(val *string) {
	if err := j.validateSetDbNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetDbSubnetGroupName(val *string) {
	if err := j.validateSetDbSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetDedicatedLogVolume(val interface{}) {
	if err := j.validateSetDedicatedLogVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedLogVolume",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetDeleteAutomatedBackups(val interface{}) {
	if err := j.validateSetDeleteAutomatedBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAutomatedBackups",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetDomainAuthSecretArn(val *string) {
	if err := j.validateSetDomainAuthSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainAuthSecretArn",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetDomainDnsIps(val *[]*string) {
	if err := j.validateSetDomainDnsIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainDnsIps",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetDomainFqdn(val *string) {
	if err := j.validateSetDomainFqdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainFqdn",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetDomainIamRoleName(val *string) {
	if err := j.validateSetDomainIamRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainIamRoleName",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetDomainOu(val *string) {
	if err := j.validateSetDomainOuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainOu",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetEnabledCloudwatchLogsExports(val *[]*string) {
	if err := j.validateSetEnabledCloudwatchLogsExportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledCloudwatchLogsExports",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetEngineLifecycleSupport(val *string) {
	if err := j.validateSetEngineLifecycleSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineLifecycleSupport",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetFinalSnapshotIdentifier(val *string) {
	if err := j.validateSetFinalSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetIamDatabaseAuthenticationEnabled(val interface{}) {
	if err := j.validateSetIamDatabaseAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamDatabaseAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetIdentifier(val *string) {
	if err := j.validateSetIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identifier",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetIdentifierPrefix(val *string) {
	if err := j.validateSetIdentifierPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identifierPrefix",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetInstanceClass(val *string) {
	if err := j.validateSetInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceClass",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetIops(val *float64) {
	if err := j.validateSetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetMaintenanceWindow(val *string) {
	if err := j.validateSetMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetManageMasterUserPassword(val interface{}) {
	if err := j.validateSetManageMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageMasterUserPassword",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetMasterUserSecretKmsKeyId(val *string) {
	if err := j.validateSetMasterUserSecretKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserSecretKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetMaxAllocatedStorage(val *float64) {
	if err := j.validateSetMaxAllocatedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAllocatedStorage",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetMonitoringInterval(val *float64) {
	if err := j.validateSetMonitoringIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringInterval",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetMonitoringRoleArn(val *string) {
	if err := j.validateSetMonitoringRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetMultiAz(val interface{}) {
	if err := j.validateSetMultiAzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiAz",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetNcharCharacterSetName(val *string) {
	if err := j.validateSetNcharCharacterSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ncharCharacterSetName",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetOptionGroupName(val *string) {
	if err := j.validateSetOptionGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetParameterGroupName(val *string) {
	if err := j.validateSetParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetPasswordWo(val *string) {
	if err := j.validateSetPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWo",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetPasswordWoVersion(val *float64) {
	if err := j.validateSetPasswordWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWoVersion",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetPerformanceInsightsEnabled(val interface{}) {
	if err := j.validateSetPerformanceInsightsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetPerformanceInsightsKmsKeyId(val *string) {
	if err := j.validateSetPerformanceInsightsKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetPerformanceInsightsRetentionPeriod(val *float64) {
	if err := j.validateSetPerformanceInsightsRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetReplicaMode(val *string) {
	if err := j.validateSetReplicaModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaMode",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetReplicateSourceDb(val *string) {
	if err := j.validateSetReplicateSourceDbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicateSourceDb",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetSkipFinalSnapshot(val interface{}) {
	if err := j.validateSetSkipFinalSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipFinalSnapshot",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetSnapshotIdentifier(val *string) {
	if err := j.validateSetSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetStorageEncrypted(val interface{}) {
	if err := j.validateSetStorageEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetStorageThroughput(val *float64) {
	if err := j.validateSetStorageThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageThroughput",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetUpgradeStorageConfig(val interface{}) {
	if err := j.validateSetUpgradeStorageConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upgradeStorageConfig",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_AwsDbInstance)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Generates CDKTN code for importing a AwsDbInstance resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsDbInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsDbInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.AwsDbInstance",
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
func AwsDbInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDbInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.AwsDbInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsDbInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDbInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.AwsDbInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsDbInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDbInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.AwsDbInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsDbInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-rds.AwsDbInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsDbInstance) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsDbInstance) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsDbInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDbInstance) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDbInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDbInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDbInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDbInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDbInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDbInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDbInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDbInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDbInstance) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsDbInstance) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDbInstance) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsDbInstance) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsDbInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsDbInstance) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsDbInstance) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsDbInstance) PutBlueGreenUpdate(value *AwsDbInstance_BlueGreenUpdateProperty) {
	if err := a.validatePutBlueGreenUpdateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBlueGreenUpdate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDbInstance) PutRestoreToPointInTime(value *AwsDbInstance_RestoreToPointInTimeProperty) {
	if err := a.validatePutRestoreToPointInTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRestoreToPointInTime",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDbInstance) PutS3Import(value *AwsDbInstance_S3ImportProperty) {
	if err := a.validatePutS3ImportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Import",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDbInstance) PutTimeouts(value *AwsDbInstance_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDbInstance) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetAllocatedStorage() {
	_jsii_.InvokeVoid(
		a,
		"resetAllocatedStorage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetAllowMajorVersionUpgrade() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowMajorVersionUpgrade",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetApplyImmediately() {
	_jsii_.InvokeVoid(
		a,
		"resetApplyImmediately",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetBackupRetentionPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetBackupRetentionPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetBackupTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetBackupTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetBackupWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetBackupWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetBlueGreenUpdate() {
	_jsii_.InvokeVoid(
		a,
		"resetBlueGreenUpdate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetCaCertIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetCaCertIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetCharacterSetName() {
	_jsii_.InvokeVoid(
		a,
		"resetCharacterSetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetCopyTagsToSnapshot() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyTagsToSnapshot",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetCustomerOwnedIpEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerOwnedIpEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetCustomIamInstanceProfile() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomIamInstanceProfile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetDatabaseInsightsMode() {
	_jsii_.InvokeVoid(
		a,
		"resetDatabaseInsightsMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetDbName() {
	_jsii_.InvokeVoid(
		a,
		"resetDbName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetDbSubnetGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetDbSubnetGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetDedicatedLogVolume() {
	_jsii_.InvokeVoid(
		a,
		"resetDedicatedLogVolume",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetDeleteAutomatedBackups() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteAutomatedBackups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetDomainAuthSecretArn() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainAuthSecretArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetDomainDnsIps() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainDnsIps",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetDomainFqdn() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainFqdn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetDomainIamRoleName() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainIamRoleName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetDomainOu() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainOu",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetEnabledCloudwatchLogsExports() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabledCloudwatchLogsExports",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetEngine() {
	_jsii_.InvokeVoid(
		a,
		"resetEngine",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetEngineLifecycleSupport() {
	_jsii_.InvokeVoid(
		a,
		"resetEngineLifecycleSupport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetFinalSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetFinalSnapshotIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetIamDatabaseAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetIamDatabaseAuthenticationEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetIdentifierPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentifierPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetIops() {
	_jsii_.InvokeVoid(
		a,
		"resetIops",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetLicenseModel() {
	_jsii_.InvokeVoid(
		a,
		"resetLicenseModel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetManageMasterUserPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetManageMasterUserPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetMasterUserSecretKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterUserSecretKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetMaxAllocatedStorage() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxAllocatedStorage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetMonitoringInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitoringInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetMonitoringRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitoringRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetMultiAz() {
	_jsii_.InvokeVoid(
		a,
		"resetMultiAz",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetNcharCharacterSetName() {
	_jsii_.InvokeVoid(
		a,
		"resetNcharCharacterSetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetNetworkType() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetOptionGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetOptionGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetParameterGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetParameterGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetPasswordWo() {
	_jsii_.InvokeVoid(
		a,
		"resetPasswordWo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetPasswordWoVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetPasswordWoVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetPerformanceInsightsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetPerformanceInsightsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetPerformanceInsightsKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetPerformanceInsightsKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetPerformanceInsightsRetentionPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetPerformanceInsightsRetentionPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetPubliclyAccessible() {
	_jsii_.InvokeVoid(
		a,
		"resetPubliclyAccessible",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetReplicaMode() {
	_jsii_.InvokeVoid(
		a,
		"resetReplicaMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetReplicateSourceDb() {
	_jsii_.InvokeVoid(
		a,
		"resetReplicateSourceDb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetRestoreToPointInTime() {
	_jsii_.InvokeVoid(
		a,
		"resetRestoreToPointInTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetS3Import() {
	_jsii_.InvokeVoid(
		a,
		"resetS3Import",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetSkipFinalSnapshot() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipFinalSnapshot",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetStorageEncrypted() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageEncrypted",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetStorageThroughput() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageThroughput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetStorageType() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetTimezone() {
	_jsii_.InvokeVoid(
		a,
		"resetTimezone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetUpgradeStorageConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetUpgradeStorageConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDbInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDbInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDbInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDbInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDbInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDbInstance) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

