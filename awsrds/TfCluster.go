package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsrds/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsrds/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster aws_rds_cluster}.
// Experimental.
type TfCluster interface {
	cdktn.TerraformResource
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
	AvailabilityZones() *[]*string
	// Experimental.
	SetAvailabilityZones(val *[]*string)
	// Experimental.
	AvailabilityZonesInput() *[]*string
	// Experimental.
	BacktrackWindow() *float64
	// Experimental.
	SetBacktrackWindow(val *float64)
	// Experimental.
	BacktrackWindowInput() *float64
	// Experimental.
	BackupRetentionPeriod() *float64
	// Experimental.
	SetBackupRetentionPeriod(val *float64)
	// Experimental.
	BackupRetentionPeriodInput() *float64
	// Experimental.
	CaCertificateIdentifier() *string
	// Experimental.
	SetCaCertificateIdentifier(val *string)
	// Experimental.
	CaCertificateIdentifierInput() *string
	// Experimental.
	CaCertificateValidTill() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClusterIdentifier() *string
	// Experimental.
	SetClusterIdentifier(val *string)
	// Experimental.
	ClusterIdentifierInput() *string
	// Experimental.
	ClusterIdentifierPrefix() *string
	// Experimental.
	SetClusterIdentifierPrefix(val *string)
	// Experimental.
	ClusterIdentifierPrefixInput() *string
	// Experimental.
	ClusterMembers() *[]*string
	// Experimental.
	SetClusterMembers(val *[]*string)
	// Experimental.
	ClusterMembersInput() *[]*string
	// Experimental.
	ClusterResourceId() *string
	// Experimental.
	ClusterScalabilityType() *string
	// Experimental.
	SetClusterScalabilityType(val *string)
	// Experimental.
	ClusterScalabilityTypeInput() *string
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
	DatabaseInsightsMode() *string
	// Experimental.
	SetDatabaseInsightsMode(val *string)
	// Experimental.
	DatabaseInsightsModeInput() *string
	// Experimental.
	DatabaseName() *string
	// Experimental.
	SetDatabaseName(val *string)
	// Experimental.
	DatabaseNameInput() *string
	// Experimental.
	DbClusterInstanceClass() *string
	// Experimental.
	SetDbClusterInstanceClass(val *string)
	// Experimental.
	DbClusterInstanceClassInput() *string
	// Experimental.
	DbClusterParameterGroupName() *string
	// Experimental.
	SetDbClusterParameterGroupName(val *string)
	// Experimental.
	DbClusterParameterGroupNameInput() *string
	// Experimental.
	DbInstanceParameterGroupName() *string
	// Experimental.
	SetDbInstanceParameterGroupName(val *string)
	// Experimental.
	DbInstanceParameterGroupNameInput() *string
	// Experimental.
	DbSubnetGroupName() *string
	// Experimental.
	SetDbSubnetGroupName(val *string)
	// Experimental.
	DbSubnetGroupNameInput() *string
	// Experimental.
	DbSystemId() *string
	// Experimental.
	SetDbSystemId(val *string)
	// Experimental.
	DbSystemIdInput() *string
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
	DomainIamRoleName() *string
	// Experimental.
	SetDomainIamRoleName(val *string)
	// Experimental.
	DomainIamRoleNameInput() *string
	// Experimental.
	DomainInput() *string
	// Experimental.
	EnabledCloudwatchLogsExports() *[]*string
	// Experimental.
	SetEnabledCloudwatchLogsExports(val *[]*string)
	// Experimental.
	EnabledCloudwatchLogsExportsInput() *[]*string
	// Experimental.
	EnableGlobalWriteForwarding() interface{}
	// Experimental.
	SetEnableGlobalWriteForwarding(val interface{})
	// Experimental.
	EnableGlobalWriteForwardingInput() interface{}
	// Experimental.
	EnableHttpEndpoint() interface{}
	// Experimental.
	SetEnableHttpEndpoint(val interface{})
	// Experimental.
	EnableHttpEndpointInput() interface{}
	// Experimental.
	EnableLocalWriteForwarding() interface{}
	// Experimental.
	SetEnableLocalWriteForwarding(val interface{})
	// Experimental.
	EnableLocalWriteForwardingInput() interface{}
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
	EngineMode() *string
	// Experimental.
	SetEngineMode(val *string)
	// Experimental.
	EngineModeInput() *string
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
	GlobalClusterIdentifier() *string
	// Experimental.
	SetGlobalClusterIdentifier(val *string)
	// Experimental.
	GlobalClusterIdentifierInput() *string
	// Experimental.
	HostedZoneId() *string
	// Experimental.
	IamDatabaseAuthenticationEnabled() interface{}
	// Experimental.
	SetIamDatabaseAuthenticationEnabled(val interface{})
	// Experimental.
	IamDatabaseAuthenticationEnabledInput() interface{}
	// Experimental.
	IamRoles() *[]*string
	// Experimental.
	SetIamRoles(val *[]*string)
	// Experimental.
	IamRolesInput() *[]*string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
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
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	ManageMasterUserPassword() interface{}
	// Experimental.
	SetManageMasterUserPassword(val interface{})
	// Experimental.
	ManageMasterUserPasswordInput() interface{}
	// Experimental.
	MasterPassword() *string
	// Experimental.
	SetMasterPassword(val *string)
	// Experimental.
	MasterPasswordInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	MasterPasswordWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetMasterPasswordWo(val *string)
	// Experimental.
	MasterPasswordWoInput() *string
	// Experimental.
	MasterPasswordWoVersion() *float64
	// Experimental.
	SetMasterPasswordWoVersion(val *float64)
	// Experimental.
	MasterPasswordWoVersionInput() *float64
	// Experimental.
	MasterUsername() *string
	// Experimental.
	SetMasterUsername(val *string)
	// Experimental.
	MasterUsernameInput() *string
	// Experimental.
	MasterUserSecret() TfCluster_MasterUserSecretPropertyList
	// Experimental.
	MasterUserSecretKmsKeyId() *string
	// Experimental.
	SetMasterUserSecretKmsKeyId(val *string)
	// Experimental.
	MasterUserSecretKmsKeyIdInput() *string
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
	NetworkType() *string
	// Experimental.
	SetNetworkType(val *string)
	// Experimental.
	NetworkTypeInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	PreferredBackupWindow() *string
	// Experimental.
	SetPreferredBackupWindow(val *string)
	// Experimental.
	PreferredBackupWindowInput() *string
	// Experimental.
	PreferredMaintenanceWindow() *string
	// Experimental.
	SetPreferredMaintenanceWindow(val *string)
	// Experimental.
	PreferredMaintenanceWindowInput() *string
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
	ReaderEndpoint() *string
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	ReplicationSourceIdentifier() *string
	// Experimental.
	SetReplicationSourceIdentifier(val *string)
	// Experimental.
	ReplicationSourceIdentifierInput() *string
	// Experimental.
	RestoreToPointInTime() TfCluster_RestoreToPointInTimePropertyOutputReference
	// Experimental.
	RestoreToPointInTimeInput() *TfCluster_RestoreToPointInTimeProperty
	// Experimental.
	S3Import() TfCluster_S3ImportPropertyOutputReference
	// Experimental.
	S3ImportInput() *TfCluster_S3ImportProperty
	// Experimental.
	ScalingConfiguration() TfCluster_ScalingConfigurationPropertyOutputReference
	// Experimental.
	ScalingConfigurationInput() *TfCluster_ScalingConfigurationProperty
	// Experimental.
	Serverlessv2ScalingConfiguration() TfCluster_Serverlessv2ScalingConfigurationPropertyOutputReference
	// Experimental.
	Serverlessv2ScalingConfigurationInput() *TfCluster_Serverlessv2ScalingConfigurationProperty
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
	SourceRegion() *string
	// Experimental.
	SetSourceRegion(val *string)
	// Experimental.
	SourceRegionInput() *string
	// Experimental.
	StorageEncrypted() interface{}
	// Experimental.
	SetStorageEncrypted(val interface{})
	// Experimental.
	StorageEncryptedInput() interface{}
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
	Timeouts() TfCluster_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	UpgradeRolloutOrder() *string
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
	PutRestoreToPointInTime(value *TfCluster_RestoreToPointInTimeProperty)
	// Experimental.
	PutS3Import(value *TfCluster_S3ImportProperty)
	// Experimental.
	PutScalingConfiguration(value *TfCluster_ScalingConfigurationProperty)
	// Experimental.
	PutServerlessv2ScalingConfiguration(value *TfCluster_Serverlessv2ScalingConfigurationProperty)
	// Experimental.
	PutTimeouts(value *TfCluster_TimeoutsProperty)
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
	ResetAvailabilityZones()
	// Experimental.
	ResetBacktrackWindow()
	// Experimental.
	ResetBackupRetentionPeriod()
	// Experimental.
	ResetCaCertificateIdentifier()
	// Experimental.
	ResetClusterIdentifier()
	// Experimental.
	ResetClusterIdentifierPrefix()
	// Experimental.
	ResetClusterMembers()
	// Experimental.
	ResetClusterScalabilityType()
	// Experimental.
	ResetCopyTagsToSnapshot()
	// Experimental.
	ResetDatabaseInsightsMode()
	// Experimental.
	ResetDatabaseName()
	// Experimental.
	ResetDbClusterInstanceClass()
	// Experimental.
	ResetDbClusterParameterGroupName()
	// Experimental.
	ResetDbInstanceParameterGroupName()
	// Experimental.
	ResetDbSubnetGroupName()
	// Experimental.
	ResetDbSystemId()
	// Experimental.
	ResetDeleteAutomatedBackups()
	// Experimental.
	ResetDeletionProtection()
	// Experimental.
	ResetDomain()
	// Experimental.
	ResetDomainIamRoleName()
	// Experimental.
	ResetEnabledCloudwatchLogsExports()
	// Experimental.
	ResetEnableGlobalWriteForwarding()
	// Experimental.
	ResetEnableHttpEndpoint()
	// Experimental.
	ResetEnableLocalWriteForwarding()
	// Experimental.
	ResetEngineLifecycleSupport()
	// Experimental.
	ResetEngineMode()
	// Experimental.
	ResetEngineVersion()
	// Experimental.
	ResetFinalSnapshotIdentifier()
	// Experimental.
	ResetGlobalClusterIdentifier()
	// Experimental.
	ResetIamDatabaseAuthenticationEnabled()
	// Experimental.
	ResetIamRoles()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIops()
	// Experimental.
	ResetKmsKeyId()
	// Experimental.
	ResetManageMasterUserPassword()
	// Experimental.
	ResetMasterPassword()
	// Experimental.
	ResetMasterPasswordWo()
	// Experimental.
	ResetMasterPasswordWoVersion()
	// Experimental.
	ResetMasterUsername()
	// Experimental.
	ResetMasterUserSecretKmsKeyId()
	// Experimental.
	ResetMonitoringInterval()
	// Experimental.
	ResetMonitoringRoleArn()
	// Experimental.
	ResetNetworkType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPerformanceInsightsEnabled()
	// Experimental.
	ResetPerformanceInsightsKmsKeyId()
	// Experimental.
	ResetPerformanceInsightsRetentionPeriod()
	// Experimental.
	ResetPort()
	// Experimental.
	ResetPreferredBackupWindow()
	// Experimental.
	ResetPreferredMaintenanceWindow()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetReplicationSourceIdentifier()
	// Experimental.
	ResetRestoreToPointInTime()
	// Experimental.
	ResetS3Import()
	// Experimental.
	ResetScalingConfiguration()
	// Experimental.
	ResetServerlessv2ScalingConfiguration()
	// Experimental.
	ResetSkipFinalSnapshot()
	// Experimental.
	ResetSnapshotIdentifier()
	// Experimental.
	ResetSourceRegion()
	// Experimental.
	ResetStorageEncrypted()
	// Experimental.
	ResetStorageType()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
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

// The jsii proxy struct for TfCluster
type jsiiProxy_TfCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfCluster) AllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) AllocatedStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) AllowMajorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) AllowMajorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ApplyImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ApplyImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) AutoMinorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BacktrackWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backtrackWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BacktrackWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backtrackWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BackupRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BackupRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) CaCertificateIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) CaCertificateIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) CaCertificateValidTill() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateValidTill",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ClusterIdentifierPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifierPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ClusterIdentifierPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifierPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ClusterMembers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterMembers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ClusterMembersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterMembersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ClusterResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ClusterScalabilityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterScalabilityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ClusterScalabilityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterScalabilityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) CopyTagsToSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) CopyTagsToSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DatabaseInsightsMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInsightsMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DatabaseInsightsModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInsightsModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DbClusterInstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterInstanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DbClusterInstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterInstanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DbClusterParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DbClusterParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DbInstanceParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DbInstanceParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DbSubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DbSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DbSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DeleteAutomatedBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DeleteAutomatedBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DomainIamRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DomainIamRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EnabledCloudwatchLogsExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledCloudwatchLogsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EnabledCloudwatchLogsExportsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledCloudwatchLogsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EnableGlobalWriteForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGlobalWriteForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EnableGlobalWriteForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGlobalWriteForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EnableHttpEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EnableHttpEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EnableLocalWriteForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalWriteForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EnableLocalWriteForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalWriteForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EngineLifecycleSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineLifecycleSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EngineLifecycleSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineLifecycleSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EngineMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EngineModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EngineVersionActual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionActual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) FinalSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) FinalSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) GlobalClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) GlobalClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) IamDatabaseAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamDatabaseAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) IamDatabaseAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamDatabaseAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) IamRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) IamRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ManageMasterUserPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ManageMasterUserPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MasterPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MasterPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MasterPasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MasterPasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MasterPasswordWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterPasswordWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MasterPasswordWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterPasswordWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MasterUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MasterUserSecret() TfCluster_MasterUserSecretPropertyList {
	var returns TfCluster_MasterUserSecretPropertyList
	_jsii_.Get(
		j,
		"masterUserSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MasterUserSecretKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserSecretKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MasterUserSecretKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserSecretKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MonitoringInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MonitoringIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MonitoringRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) MonitoringRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) PerformanceInsightsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceInsightsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) PerformanceInsightsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceInsightsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) PerformanceInsightsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) PerformanceInsightsKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) PerformanceInsightsRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) PerformanceInsightsRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) PreferredBackupWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) PreferredMaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ReaderEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ReplicationSourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ReplicationSourceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSourceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) RestoreToPointInTime() TfCluster_RestoreToPointInTimePropertyOutputReference {
	var returns TfCluster_RestoreToPointInTimePropertyOutputReference
	_jsii_.Get(
		j,
		"restoreToPointInTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) RestoreToPointInTimeInput() *TfCluster_RestoreToPointInTimeProperty {
	var returns *TfCluster_RestoreToPointInTimeProperty
	_jsii_.Get(
		j,
		"restoreToPointInTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) S3Import() TfCluster_S3ImportPropertyOutputReference {
	var returns TfCluster_S3ImportPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Import",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) S3ImportInput() *TfCluster_S3ImportProperty {
	var returns *TfCluster_S3ImportProperty
	_jsii_.Get(
		j,
		"s3ImportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ScalingConfiguration() TfCluster_ScalingConfigurationPropertyOutputReference {
	var returns TfCluster_ScalingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"scalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ScalingConfigurationInput() *TfCluster_ScalingConfigurationProperty {
	var returns *TfCluster_ScalingConfigurationProperty
	_jsii_.Get(
		j,
		"scalingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Serverlessv2ScalingConfiguration() TfCluster_Serverlessv2ScalingConfigurationPropertyOutputReference {
	var returns TfCluster_Serverlessv2ScalingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"serverlessv2ScalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Serverlessv2ScalingConfigurationInput() *TfCluster_Serverlessv2ScalingConfigurationProperty {
	var returns *TfCluster_Serverlessv2ScalingConfigurationProperty
	_jsii_.Get(
		j,
		"serverlessv2ScalingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) SkipFinalSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) SkipFinalSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) SnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) SnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) SourceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) SourceRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) StorageEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) StorageEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) Timeouts() TfCluster_TimeoutsPropertyOutputReference {
	var returns TfCluster_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) UpgradeRolloutOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeRolloutOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster aws_rds_cluster} Resource.
// Experimental.
func NewTfCluster(scope constructs.Construct, id *string, config *TfClusterConfig) TfCluster {
	_init_.Initialize()

	if err := validateNewTfClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCluster{}

	_jsii_.Create(
		"@cdktn/aws-rds.TfCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster aws_rds_cluster} Resource.
// Experimental.
func NewTfCluster_Override(t TfCluster, scope constructs.Construct, id *string, config *TfClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-rds.TfCluster",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfCluster)SetAllocatedStorage(val *float64) {
	if err := j.validateSetAllocatedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocatedStorage",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetAllowMajorVersionUpgrade(val interface{}) {
	if err := j.validateSetAllowMajorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMajorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetApplyImmediately(val interface{}) {
	if err := j.validateSetApplyImmediatelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyImmediately",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetAutoMinorVersionUpgrade(val interface{}) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetBacktrackWindow(val *float64) {
	if err := j.validateSetBacktrackWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backtrackWindow",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetBackupRetentionPeriod(val *float64) {
	if err := j.validateSetBackupRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetCaCertificateIdentifier(val *string) {
	if err := j.validateSetCaCertificateIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificateIdentifier",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetClusterIdentifier(val *string) {
	if err := j.validateSetClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetClusterIdentifierPrefix(val *string) {
	if err := j.validateSetClusterIdentifierPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIdentifierPrefix",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetClusterMembers(val *[]*string) {
	if err := j.validateSetClusterMembersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterMembers",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetClusterScalabilityType(val *string) {
	if err := j.validateSetClusterScalabilityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterScalabilityType",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetCopyTagsToSnapshot(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshot",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetDatabaseInsightsMode(val *string) {
	if err := j.validateSetDatabaseInsightsModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseInsightsMode",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetDbClusterInstanceClass(val *string) {
	if err := j.validateSetDbClusterInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbClusterInstanceClass",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetDbClusterParameterGroupName(val *string) {
	if err := j.validateSetDbClusterParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbClusterParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetDbInstanceParameterGroupName(val *string) {
	if err := j.validateSetDbInstanceParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbInstanceParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetDbSubnetGroupName(val *string) {
	if err := j.validateSetDbSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetDbSystemId(val *string) {
	if err := j.validateSetDbSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSystemId",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetDeleteAutomatedBackups(val interface{}) {
	if err := j.validateSetDeleteAutomatedBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAutomatedBackups",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetDomainIamRoleName(val *string) {
	if err := j.validateSetDomainIamRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainIamRoleName",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetEnabledCloudwatchLogsExports(val *[]*string) {
	if err := j.validateSetEnabledCloudwatchLogsExportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledCloudwatchLogsExports",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetEnableGlobalWriteForwarding(val interface{}) {
	if err := j.validateSetEnableGlobalWriteForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableGlobalWriteForwarding",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetEnableHttpEndpoint(val interface{}) {
	if err := j.validateSetEnableHttpEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHttpEndpoint",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetEnableLocalWriteForwarding(val interface{}) {
	if err := j.validateSetEnableLocalWriteForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLocalWriteForwarding",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetEngineLifecycleSupport(val *string) {
	if err := j.validateSetEngineLifecycleSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineLifecycleSupport",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetEngineMode(val *string) {
	if err := j.validateSetEngineModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineMode",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetFinalSnapshotIdentifier(val *string) {
	if err := j.validateSetFinalSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetGlobalClusterIdentifier(val *string) {
	if err := j.validateSetGlobalClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetIamDatabaseAuthenticationEnabled(val interface{}) {
	if err := j.validateSetIamDatabaseAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamDatabaseAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetIamRoles(val *[]*string) {
	if err := j.validateSetIamRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRoles",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetIops(val *float64) {
	if err := j.validateSetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetManageMasterUserPassword(val interface{}) {
	if err := j.validateSetManageMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageMasterUserPassword",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetMasterPassword(val *string) {
	if err := j.validateSetMasterPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterPassword",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetMasterPasswordWo(val *string) {
	if err := j.validateSetMasterPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterPasswordWo",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetMasterPasswordWoVersion(val *float64) {
	if err := j.validateSetMasterPasswordWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterPasswordWoVersion",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetMasterUsername(val *string) {
	if err := j.validateSetMasterUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetMasterUserSecretKmsKeyId(val *string) {
	if err := j.validateSetMasterUserSecretKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserSecretKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetMonitoringInterval(val *float64) {
	if err := j.validateSetMonitoringIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringInterval",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetMonitoringRoleArn(val *string) {
	if err := j.validateSetMonitoringRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetPerformanceInsightsEnabled(val interface{}) {
	if err := j.validateSetPerformanceInsightsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsEnabled",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetPerformanceInsightsKmsKeyId(val *string) {
	if err := j.validateSetPerformanceInsightsKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetPerformanceInsightsRetentionPeriod(val *float64) {
	if err := j.validateSetPerformanceInsightsRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetPreferredBackupWindow(val *string) {
	if err := j.validateSetPreferredBackupWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetPreferredMaintenanceWindow(val *string) {
	if err := j.validateSetPreferredMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetReplicationSourceIdentifier(val *string) {
	if err := j.validateSetReplicationSourceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationSourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetSkipFinalSnapshot(val interface{}) {
	if err := j.validateSetSkipFinalSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipFinalSnapshot",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetSnapshotIdentifier(val *string) {
	if err := j.validateSetSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetSourceRegion(val *string) {
	if err := j.validateSetSourceRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceRegion",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetStorageEncrypted(val interface{}) {
	if err := j.validateSetStorageEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Generates CDKTN code for importing a TfCluster resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.TfCluster",
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
func TfCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.TfCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.TfCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.TfCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-rds.TfCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfCluster) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfCluster) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfCluster) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfCluster) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCluster) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfCluster) PutRestoreToPointInTime(value *TfCluster_RestoreToPointInTimeProperty) {
	if err := t.validatePutRestoreToPointInTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRestoreToPointInTime",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutS3Import(value *TfCluster_S3ImportProperty) {
	if err := t.validatePutS3ImportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Import",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutScalingConfiguration(value *TfCluster_ScalingConfigurationProperty) {
	if err := t.validatePutScalingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putScalingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutServerlessv2ScalingConfiguration(value *TfCluster_Serverlessv2ScalingConfigurationProperty) {
	if err := t.validatePutServerlessv2ScalingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putServerlessv2ScalingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutTimeouts(value *TfCluster_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfCluster) ResetAllocatedStorage() {
	_jsii_.InvokeVoid(
		t,
		"resetAllocatedStorage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetAllowMajorVersionUpgrade() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowMajorVersionUpgrade",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetApplyImmediately() {
	_jsii_.InvokeVoid(
		t,
		"resetApplyImmediately",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetAvailabilityZones() {
	_jsii_.InvokeVoid(
		t,
		"resetAvailabilityZones",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetBacktrackWindow() {
	_jsii_.InvokeVoid(
		t,
		"resetBacktrackWindow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetBackupRetentionPeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetBackupRetentionPeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetCaCertificateIdentifier() {
	_jsii_.InvokeVoid(
		t,
		"resetCaCertificateIdentifier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetClusterIdentifier() {
	_jsii_.InvokeVoid(
		t,
		"resetClusterIdentifier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetClusterIdentifierPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetClusterIdentifierPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetClusterMembers() {
	_jsii_.InvokeVoid(
		t,
		"resetClusterMembers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetClusterScalabilityType() {
	_jsii_.InvokeVoid(
		t,
		"resetClusterScalabilityType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetCopyTagsToSnapshot() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyTagsToSnapshot",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetDatabaseInsightsMode() {
	_jsii_.InvokeVoid(
		t,
		"resetDatabaseInsightsMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		t,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetDbClusterInstanceClass() {
	_jsii_.InvokeVoid(
		t,
		"resetDbClusterInstanceClass",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetDbClusterParameterGroupName() {
	_jsii_.InvokeVoid(
		t,
		"resetDbClusterParameterGroupName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetDbInstanceParameterGroupName() {
	_jsii_.InvokeVoid(
		t,
		"resetDbInstanceParameterGroupName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetDbSubnetGroupName() {
	_jsii_.InvokeVoid(
		t,
		"resetDbSubnetGroupName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetDbSystemId() {
	_jsii_.InvokeVoid(
		t,
		"resetDbSystemId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetDeleteAutomatedBackups() {
	_jsii_.InvokeVoid(
		t,
		"resetDeleteAutomatedBackups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		t,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetDomain() {
	_jsii_.InvokeVoid(
		t,
		"resetDomain",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetDomainIamRoleName() {
	_jsii_.InvokeVoid(
		t,
		"resetDomainIamRoleName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetEnabledCloudwatchLogsExports() {
	_jsii_.InvokeVoid(
		t,
		"resetEnabledCloudwatchLogsExports",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetEnableGlobalWriteForwarding() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableGlobalWriteForwarding",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetEnableHttpEndpoint() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableHttpEndpoint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetEnableLocalWriteForwarding() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableLocalWriteForwarding",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetEngineLifecycleSupport() {
	_jsii_.InvokeVoid(
		t,
		"resetEngineLifecycleSupport",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetEngineMode() {
	_jsii_.InvokeVoid(
		t,
		"resetEngineMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetFinalSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		t,
		"resetFinalSnapshotIdentifier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetGlobalClusterIdentifier() {
	_jsii_.InvokeVoid(
		t,
		"resetGlobalClusterIdentifier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetIamDatabaseAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetIamDatabaseAuthenticationEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetIamRoles() {
	_jsii_.InvokeVoid(
		t,
		"resetIamRoles",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetIops() {
	_jsii_.InvokeVoid(
		t,
		"resetIops",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetManageMasterUserPassword() {
	_jsii_.InvokeVoid(
		t,
		"resetManageMasterUserPassword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetMasterPassword() {
	_jsii_.InvokeVoid(
		t,
		"resetMasterPassword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetMasterPasswordWo() {
	_jsii_.InvokeVoid(
		t,
		"resetMasterPasswordWo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetMasterPasswordWoVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetMasterPasswordWoVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetMasterUsername() {
	_jsii_.InvokeVoid(
		t,
		"resetMasterUsername",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetMasterUserSecretKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetMasterUserSecretKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetMonitoringInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetMonitoringInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetMonitoringRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetMonitoringRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetNetworkType() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetPerformanceInsightsEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetPerformanceInsightsEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetPerformanceInsightsKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetPerformanceInsightsKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetPerformanceInsightsRetentionPeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetPerformanceInsightsRetentionPeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetPort() {
	_jsii_.InvokeVoid(
		t,
		"resetPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetPreferredBackupWindow() {
	_jsii_.InvokeVoid(
		t,
		"resetPreferredBackupWindow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetPreferredMaintenanceWindow() {
	_jsii_.InvokeVoid(
		t,
		"resetPreferredMaintenanceWindow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetReplicationSourceIdentifier() {
	_jsii_.InvokeVoid(
		t,
		"resetReplicationSourceIdentifier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetRestoreToPointInTime() {
	_jsii_.InvokeVoid(
		t,
		"resetRestoreToPointInTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetS3Import() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Import",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetScalingConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetScalingConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetServerlessv2ScalingConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetServerlessv2ScalingConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetSkipFinalSnapshot() {
	_jsii_.InvokeVoid(
		t,
		"resetSkipFinalSnapshot",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		t,
		"resetSnapshotIdentifier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetSourceRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetStorageEncrypted() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageEncrypted",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetStorageType() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

