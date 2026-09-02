package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsrds/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsrds/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance aws_db_instance}.
// Experimental.
type TfDbInstance interface {
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
	BlueGreenUpdate() TfDbInstance_BlueGreenUpdatePropertyOutputReference
	// Experimental.
	BlueGreenUpdateInput() *TfDbInstance_BlueGreenUpdateProperty
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
	ListenerEndpoint() TfDbInstance_ListenerEndpointPropertyList
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
	MasterUserSecret() TfDbInstance_MasterUserSecretPropertyList
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
	RestoreToPointInTime() TfDbInstance_RestoreToPointInTimePropertyOutputReference
	// Experimental.
	RestoreToPointInTimeInput() *TfDbInstance_RestoreToPointInTimeProperty
	// Experimental.
	S3Import() TfDbInstance_S3ImportPropertyOutputReference
	// Experimental.
	S3ImportInput() *TfDbInstance_S3ImportProperty
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
	Timeouts() TfDbInstance_TimeoutsPropertyOutputReference
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
	PutBlueGreenUpdate(value *TfDbInstance_BlueGreenUpdateProperty)
	// Experimental.
	PutRestoreToPointInTime(value *TfDbInstance_RestoreToPointInTimeProperty)
	// Experimental.
	PutS3Import(value *TfDbInstance_S3ImportProperty)
	// Experimental.
	PutTimeouts(value *TfDbInstance_TimeoutsProperty)
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

// The jsii proxy struct for TfDbInstance
type jsiiProxy_TfDbInstance struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfDbInstance) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) AllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) AllocatedStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) AllowMajorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) AllowMajorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) ApplyImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) ApplyImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) AutoMinorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) BackupRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) BackupRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) BackupTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) BackupTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) BackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) BackupWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) BlueGreenUpdate() TfDbInstance_BlueGreenUpdatePropertyOutputReference {
	var returns TfDbInstance_BlueGreenUpdatePropertyOutputReference
	_jsii_.Get(
		j,
		"blueGreenUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) BlueGreenUpdateInput() *TfDbInstance_BlueGreenUpdateProperty {
	var returns *TfDbInstance_BlueGreenUpdateProperty
	_jsii_.Get(
		j,
		"blueGreenUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) CaCertIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) CaCertIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) CharacterSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) CharacterSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) CopyTagsToSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) CopyTagsToSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) CustomerOwnedIpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerOwnedIpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) CustomerOwnedIpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerOwnedIpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) CustomIamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customIamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) CustomIamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customIamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DatabaseInsightsMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInsightsMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DatabaseInsightsModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInsightsModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DbNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DbSubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DedicatedLogVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedLogVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DedicatedLogVolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedLogVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DeleteAutomatedBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DeleteAutomatedBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DomainAuthSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAuthSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DomainAuthSecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAuthSecretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DomainDnsIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainDnsIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DomainDnsIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainDnsIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DomainFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DomainFqdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainFqdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DomainIamRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DomainIamRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DomainOu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainOu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) DomainOuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainOuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) EnabledCloudwatchLogsExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledCloudwatchLogsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) EnabledCloudwatchLogsExportsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledCloudwatchLogsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) EngineLifecycleSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineLifecycleSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) EngineLifecycleSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineLifecycleSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) EngineVersionActual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionActual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) FinalSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) FinalSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) IamDatabaseAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamDatabaseAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) IamDatabaseAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamDatabaseAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) IdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) IdentifierPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) IdentifierPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) InstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) InstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) LatestRestorableTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestRestorableTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) ListenerEndpoint() TfDbInstance_ListenerEndpointPropertyList {
	var returns TfDbInstance_ListenerEndpointPropertyList
	_jsii_.Get(
		j,
		"listenerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) MaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) MaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) ManageMasterUserPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) ManageMasterUserPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) MasterUserSecret() TfDbInstance_MasterUserSecretPropertyList {
	var returns TfDbInstance_MasterUserSecretPropertyList
	_jsii_.Get(
		j,
		"masterUserSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) MasterUserSecretKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserSecretKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) MasterUserSecretKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserSecretKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) MaxAllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) MaxAllocatedStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) MonitoringInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) MonitoringIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) MonitoringRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) MonitoringRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) MultiAz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) MultiAzInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) NcharCharacterSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharCharacterSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) NcharCharacterSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharCharacterSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) OptionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) OptionGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) ParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) ParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) PasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) PasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) PasswordWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) PasswordWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) PerformanceInsightsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceInsightsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) PerformanceInsightsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceInsightsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) PerformanceInsightsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) PerformanceInsightsKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) PerformanceInsightsRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) PerformanceInsightsRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) PubliclyAccessibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) ReplicaMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) ReplicaModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Replicas() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) ReplicateSourceDb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicateSourceDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) ReplicateSourceDbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicateSourceDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) RestoreToPointInTime() TfDbInstance_RestoreToPointInTimePropertyOutputReference {
	var returns TfDbInstance_RestoreToPointInTimePropertyOutputReference
	_jsii_.Get(
		j,
		"restoreToPointInTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) RestoreToPointInTimeInput() *TfDbInstance_RestoreToPointInTimeProperty {
	var returns *TfDbInstance_RestoreToPointInTimeProperty
	_jsii_.Get(
		j,
		"restoreToPointInTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) S3Import() TfDbInstance_S3ImportPropertyOutputReference {
	var returns TfDbInstance_S3ImportPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Import",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) S3ImportInput() *TfDbInstance_S3ImportProperty {
	var returns *TfDbInstance_S3ImportProperty
	_jsii_.Get(
		j,
		"s3ImportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) SkipFinalSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) SkipFinalSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) SnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) SnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) StorageEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) StorageEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) StorageThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) StorageThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Timeouts() TfDbInstance_TimeoutsPropertyOutputReference {
	var returns TfDbInstance_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) UpgradeRolloutOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeRolloutOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) UpgradeStorageConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upgradeStorageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) UpgradeStorageConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upgradeStorageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDbInstance) VpcSecurityGroupIdsInput() *[]*string {
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
func NewTfDbInstance(scope constructs.Construct, id *string, config *TfDbInstanceConfig) TfDbInstance {
	_init_.Initialize()

	if err := validateNewTfDbInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDbInstance{}

	_jsii_.Create(
		"@cdktn/aws-rds.TfDbInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance aws_db_instance} Resource.
// Experimental.
func NewTfDbInstance_Override(t TfDbInstance, scope constructs.Construct, id *string, config *TfDbInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-rds.TfDbInstance",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfDbInstance)SetAllocatedStorage(val *float64) {
	if err := j.validateSetAllocatedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocatedStorage",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetAllowMajorVersionUpgrade(val interface{}) {
	if err := j.validateSetAllowMajorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMajorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetApplyImmediately(val interface{}) {
	if err := j.validateSetApplyImmediatelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyImmediately",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetAutoMinorVersionUpgrade(val interface{}) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetBackupRetentionPeriod(val *float64) {
	if err := j.validateSetBackupRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetBackupTarget(val *string) {
	if err := j.validateSetBackupTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupTarget",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetBackupWindow(val *string) {
	if err := j.validateSetBackupWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupWindow",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetCaCertIdentifier(val *string) {
	if err := j.validateSetCaCertIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertIdentifier",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetCharacterSetName(val *string) {
	if err := j.validateSetCharacterSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"characterSetName",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetCopyTagsToSnapshot(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshot",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetCustomerOwnedIpEnabled(val interface{}) {
	if err := j.validateSetCustomerOwnedIpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerOwnedIpEnabled",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetCustomIamInstanceProfile(val *string) {
	if err := j.validateSetCustomIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customIamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetDatabaseInsightsMode(val *string) {
	if err := j.validateSetDatabaseInsightsModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseInsightsMode",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetDbName(val *string) {
	if err := j.validateSetDbNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetDbSubnetGroupName(val *string) {
	if err := j.validateSetDbSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetDedicatedLogVolume(val interface{}) {
	if err := j.validateSetDedicatedLogVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedLogVolume",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetDeleteAutomatedBackups(val interface{}) {
	if err := j.validateSetDeleteAutomatedBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAutomatedBackups",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetDomainAuthSecretArn(val *string) {
	if err := j.validateSetDomainAuthSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainAuthSecretArn",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetDomainDnsIps(val *[]*string) {
	if err := j.validateSetDomainDnsIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainDnsIps",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetDomainFqdn(val *string) {
	if err := j.validateSetDomainFqdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainFqdn",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetDomainIamRoleName(val *string) {
	if err := j.validateSetDomainIamRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainIamRoleName",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetDomainOu(val *string) {
	if err := j.validateSetDomainOuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainOu",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetEnabledCloudwatchLogsExports(val *[]*string) {
	if err := j.validateSetEnabledCloudwatchLogsExportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledCloudwatchLogsExports",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetEngineLifecycleSupport(val *string) {
	if err := j.validateSetEngineLifecycleSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineLifecycleSupport",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetFinalSnapshotIdentifier(val *string) {
	if err := j.validateSetFinalSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetIamDatabaseAuthenticationEnabled(val interface{}) {
	if err := j.validateSetIamDatabaseAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamDatabaseAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetIdentifier(val *string) {
	if err := j.validateSetIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identifier",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetIdentifierPrefix(val *string) {
	if err := j.validateSetIdentifierPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identifierPrefix",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetInstanceClass(val *string) {
	if err := j.validateSetInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceClass",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetIops(val *float64) {
	if err := j.validateSetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetMaintenanceWindow(val *string) {
	if err := j.validateSetMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetManageMasterUserPassword(val interface{}) {
	if err := j.validateSetManageMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageMasterUserPassword",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetMasterUserSecretKmsKeyId(val *string) {
	if err := j.validateSetMasterUserSecretKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserSecretKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetMaxAllocatedStorage(val *float64) {
	if err := j.validateSetMaxAllocatedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAllocatedStorage",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetMonitoringInterval(val *float64) {
	if err := j.validateSetMonitoringIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringInterval",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetMonitoringRoleArn(val *string) {
	if err := j.validateSetMonitoringRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetMultiAz(val interface{}) {
	if err := j.validateSetMultiAzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiAz",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetNcharCharacterSetName(val *string) {
	if err := j.validateSetNcharCharacterSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ncharCharacterSetName",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetOptionGroupName(val *string) {
	if err := j.validateSetOptionGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionGroupName",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetParameterGroupName(val *string) {
	if err := j.validateSetParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterGroupName",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetPasswordWo(val *string) {
	if err := j.validateSetPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWo",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetPasswordWoVersion(val *float64) {
	if err := j.validateSetPasswordWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordWoVersion",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetPerformanceInsightsEnabled(val interface{}) {
	if err := j.validateSetPerformanceInsightsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsEnabled",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetPerformanceInsightsKmsKeyId(val *string) {
	if err := j.validateSetPerformanceInsightsKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetPerformanceInsightsRetentionPeriod(val *float64) {
	if err := j.validateSetPerformanceInsightsRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetReplicaMode(val *string) {
	if err := j.validateSetReplicaModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaMode",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetReplicateSourceDb(val *string) {
	if err := j.validateSetReplicateSourceDbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicateSourceDb",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetSkipFinalSnapshot(val interface{}) {
	if err := j.validateSetSkipFinalSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipFinalSnapshot",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetSnapshotIdentifier(val *string) {
	if err := j.validateSetSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetStorageEncrypted(val interface{}) {
	if err := j.validateSetStorageEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetStorageThroughput(val *float64) {
	if err := j.validateSetStorageThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageThroughput",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetUpgradeStorageConfig(val interface{}) {
	if err := j.validateSetUpgradeStorageConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upgradeStorageConfig",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_TfDbInstance)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Generates CDKTN code for importing a TfDbInstance resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfDbInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfDbInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.TfDbInstance",
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
func TfDbInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDbInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.TfDbInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfDbInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDbInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.TfDbInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfDbInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDbInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.TfDbInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfDbInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-rds.TfDbInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfDbInstance) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfDbInstance) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfDbInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDbInstance) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDbInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDbInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDbInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDbInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDbInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDbInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDbInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDbInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDbInstance) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfDbInstance) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDbInstance) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfDbInstance) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfDbInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfDbInstance) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfDbInstance) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfDbInstance) PutBlueGreenUpdate(value *TfDbInstance_BlueGreenUpdateProperty) {
	if err := t.validatePutBlueGreenUpdateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBlueGreenUpdate",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDbInstance) PutRestoreToPointInTime(value *TfDbInstance_RestoreToPointInTimeProperty) {
	if err := t.validatePutRestoreToPointInTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRestoreToPointInTime",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDbInstance) PutS3Import(value *TfDbInstance_S3ImportProperty) {
	if err := t.validatePutS3ImportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Import",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDbInstance) PutTimeouts(value *TfDbInstance_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDbInstance) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfDbInstance) ResetAllocatedStorage() {
	_jsii_.InvokeVoid(
		t,
		"resetAllocatedStorage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetAllowMajorVersionUpgrade() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowMajorVersionUpgrade",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetApplyImmediately() {
	_jsii_.InvokeVoid(
		t,
		"resetApplyImmediately",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		t,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetBackupRetentionPeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetBackupRetentionPeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetBackupTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetBackupTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetBackupWindow() {
	_jsii_.InvokeVoid(
		t,
		"resetBackupWindow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetBlueGreenUpdate() {
	_jsii_.InvokeVoid(
		t,
		"resetBlueGreenUpdate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetCaCertIdentifier() {
	_jsii_.InvokeVoid(
		t,
		"resetCaCertIdentifier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetCharacterSetName() {
	_jsii_.InvokeVoid(
		t,
		"resetCharacterSetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetCopyTagsToSnapshot() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyTagsToSnapshot",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetCustomerOwnedIpEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomerOwnedIpEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetCustomIamInstanceProfile() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomIamInstanceProfile",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetDatabaseInsightsMode() {
	_jsii_.InvokeVoid(
		t,
		"resetDatabaseInsightsMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetDbName() {
	_jsii_.InvokeVoid(
		t,
		"resetDbName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetDbSubnetGroupName() {
	_jsii_.InvokeVoid(
		t,
		"resetDbSubnetGroupName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetDedicatedLogVolume() {
	_jsii_.InvokeVoid(
		t,
		"resetDedicatedLogVolume",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetDeleteAutomatedBackups() {
	_jsii_.InvokeVoid(
		t,
		"resetDeleteAutomatedBackups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		t,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetDomain() {
	_jsii_.InvokeVoid(
		t,
		"resetDomain",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetDomainAuthSecretArn() {
	_jsii_.InvokeVoid(
		t,
		"resetDomainAuthSecretArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetDomainDnsIps() {
	_jsii_.InvokeVoid(
		t,
		"resetDomainDnsIps",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetDomainFqdn() {
	_jsii_.InvokeVoid(
		t,
		"resetDomainFqdn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetDomainIamRoleName() {
	_jsii_.InvokeVoid(
		t,
		"resetDomainIamRoleName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetDomainOu() {
	_jsii_.InvokeVoid(
		t,
		"resetDomainOu",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetEnabledCloudwatchLogsExports() {
	_jsii_.InvokeVoid(
		t,
		"resetEnabledCloudwatchLogsExports",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetEngine() {
	_jsii_.InvokeVoid(
		t,
		"resetEngine",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetEngineLifecycleSupport() {
	_jsii_.InvokeVoid(
		t,
		"resetEngineLifecycleSupport",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetFinalSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		t,
		"resetFinalSnapshotIdentifier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetIamDatabaseAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetIamDatabaseAuthenticationEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetIdentifier() {
	_jsii_.InvokeVoid(
		t,
		"resetIdentifier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetIdentifierPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetIdentifierPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetIops() {
	_jsii_.InvokeVoid(
		t,
		"resetIops",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetLicenseModel() {
	_jsii_.InvokeVoid(
		t,
		"resetLicenseModel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		t,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetManageMasterUserPassword() {
	_jsii_.InvokeVoid(
		t,
		"resetManageMasterUserPassword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetMasterUserSecretKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetMasterUserSecretKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetMaxAllocatedStorage() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxAllocatedStorage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetMonitoringInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetMonitoringInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetMonitoringRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetMonitoringRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetMultiAz() {
	_jsii_.InvokeVoid(
		t,
		"resetMultiAz",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetNcharCharacterSetName() {
	_jsii_.InvokeVoid(
		t,
		"resetNcharCharacterSetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetNetworkType() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetOptionGroupName() {
	_jsii_.InvokeVoid(
		t,
		"resetOptionGroupName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetParameterGroupName() {
	_jsii_.InvokeVoid(
		t,
		"resetParameterGroupName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetPassword() {
	_jsii_.InvokeVoid(
		t,
		"resetPassword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetPasswordWo() {
	_jsii_.InvokeVoid(
		t,
		"resetPasswordWo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetPasswordWoVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetPasswordWoVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetPerformanceInsightsEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetPerformanceInsightsEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetPerformanceInsightsKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetPerformanceInsightsKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetPerformanceInsightsRetentionPeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetPerformanceInsightsRetentionPeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetPort() {
	_jsii_.InvokeVoid(
		t,
		"resetPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetPubliclyAccessible() {
	_jsii_.InvokeVoid(
		t,
		"resetPubliclyAccessible",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetReplicaMode() {
	_jsii_.InvokeVoid(
		t,
		"resetReplicaMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetReplicateSourceDb() {
	_jsii_.InvokeVoid(
		t,
		"resetReplicateSourceDb",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetRestoreToPointInTime() {
	_jsii_.InvokeVoid(
		t,
		"resetRestoreToPointInTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetS3Import() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Import",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetSkipFinalSnapshot() {
	_jsii_.InvokeVoid(
		t,
		"resetSkipFinalSnapshot",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		t,
		"resetSnapshotIdentifier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetStorageEncrypted() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageEncrypted",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetStorageThroughput() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageThroughput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetStorageType() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetTimezone() {
	_jsii_.InvokeVoid(
		t,
		"resetTimezone",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetUpgradeStorageConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetUpgradeStorageConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetUsername() {
	_jsii_.InvokeVoid(
		t,
		"resetUsername",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDbInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDbInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDbInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDbInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDbInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDbInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDbInstance) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

