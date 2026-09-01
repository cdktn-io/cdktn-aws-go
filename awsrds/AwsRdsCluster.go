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
type AwsRdsCluster interface {
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
	MasterUserSecret() AwsRdsCluster_MasterUserSecretPropertyList
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
	RestoreToPointInTime() AwsRdsCluster_RestoreToPointInTimePropertyOutputReference
	// Experimental.
	RestoreToPointInTimeInput() *AwsRdsCluster_RestoreToPointInTimeProperty
	// Experimental.
	S3Import() AwsRdsCluster_S3ImportPropertyOutputReference
	// Experimental.
	S3ImportInput() *AwsRdsCluster_S3ImportProperty
	// Experimental.
	ScalingConfiguration() AwsRdsCluster_ScalingConfigurationPropertyOutputReference
	// Experimental.
	ScalingConfigurationInput() *AwsRdsCluster_ScalingConfigurationProperty
	// Experimental.
	Serverlessv2ScalingConfiguration() AwsRdsCluster_Serverlessv2ScalingConfigurationPropertyOutputReference
	// Experimental.
	Serverlessv2ScalingConfigurationInput() *AwsRdsCluster_Serverlessv2ScalingConfigurationProperty
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
	Timeouts() AwsRdsCluster_TimeoutsPropertyOutputReference
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
	PutRestoreToPointInTime(value *AwsRdsCluster_RestoreToPointInTimeProperty)
	// Experimental.
	PutS3Import(value *AwsRdsCluster_S3ImportProperty)
	// Experimental.
	PutScalingConfiguration(value *AwsRdsCluster_ScalingConfigurationProperty)
	// Experimental.
	PutServerlessv2ScalingConfiguration(value *AwsRdsCluster_Serverlessv2ScalingConfigurationProperty)
	// Experimental.
	PutTimeouts(value *AwsRdsCluster_TimeoutsProperty)
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

// The jsii proxy struct for AwsRdsCluster
type jsiiProxy_AwsRdsCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsRdsCluster) AllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) AllocatedStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) AllowMajorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) AllowMajorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ApplyImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ApplyImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) AutoMinorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) BacktrackWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backtrackWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) BacktrackWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backtrackWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) BackupRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) BackupRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) CaCertificateIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) CaCertificateIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) CaCertificateValidTill() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateValidTill",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ClusterIdentifierPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifierPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ClusterIdentifierPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifierPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ClusterMembers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterMembers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ClusterMembersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterMembersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ClusterResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ClusterScalabilityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterScalabilityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ClusterScalabilityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterScalabilityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) CopyTagsToSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) CopyTagsToSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DatabaseInsightsMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInsightsMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DatabaseInsightsModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInsightsModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DbClusterInstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterInstanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DbClusterInstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterInstanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DbClusterParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DbClusterParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DbInstanceParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DbInstanceParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DbSubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DbSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DbSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DeleteAutomatedBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DeleteAutomatedBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DomainIamRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DomainIamRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) EnabledCloudwatchLogsExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledCloudwatchLogsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) EnabledCloudwatchLogsExportsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledCloudwatchLogsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) EnableGlobalWriteForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGlobalWriteForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) EnableGlobalWriteForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGlobalWriteForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) EnableHttpEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) EnableHttpEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) EnableLocalWriteForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalWriteForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) EnableLocalWriteForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalWriteForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) EngineLifecycleSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineLifecycleSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) EngineLifecycleSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineLifecycleSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) EngineMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) EngineModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) EngineVersionActual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionActual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) FinalSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) FinalSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) GlobalClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) GlobalClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) IamDatabaseAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamDatabaseAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) IamDatabaseAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamDatabaseAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) IamRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) IamRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ManageMasterUserPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ManageMasterUserPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) MasterPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) MasterPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) MasterPasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) MasterPasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) MasterPasswordWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterPasswordWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) MasterPasswordWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterPasswordWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) MasterUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) MasterUserSecret() AwsRdsCluster_MasterUserSecretPropertyList {
	var returns AwsRdsCluster_MasterUserSecretPropertyList
	_jsii_.Get(
		j,
		"masterUserSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) MasterUserSecretKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserSecretKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) MasterUserSecretKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserSecretKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) MonitoringInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) MonitoringIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) MonitoringRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) MonitoringRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) PerformanceInsightsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceInsightsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) PerformanceInsightsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceInsightsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) PerformanceInsightsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) PerformanceInsightsKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) PerformanceInsightsRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) PerformanceInsightsRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) PreferredBackupWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) PreferredMaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ReaderEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ReplicationSourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ReplicationSourceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSourceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) RestoreToPointInTime() AwsRdsCluster_RestoreToPointInTimePropertyOutputReference {
	var returns AwsRdsCluster_RestoreToPointInTimePropertyOutputReference
	_jsii_.Get(
		j,
		"restoreToPointInTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) RestoreToPointInTimeInput() *AwsRdsCluster_RestoreToPointInTimeProperty {
	var returns *AwsRdsCluster_RestoreToPointInTimeProperty
	_jsii_.Get(
		j,
		"restoreToPointInTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) S3Import() AwsRdsCluster_S3ImportPropertyOutputReference {
	var returns AwsRdsCluster_S3ImportPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Import",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) S3ImportInput() *AwsRdsCluster_S3ImportProperty {
	var returns *AwsRdsCluster_S3ImportProperty
	_jsii_.Get(
		j,
		"s3ImportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ScalingConfiguration() AwsRdsCluster_ScalingConfigurationPropertyOutputReference {
	var returns AwsRdsCluster_ScalingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"scalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) ScalingConfigurationInput() *AwsRdsCluster_ScalingConfigurationProperty {
	var returns *AwsRdsCluster_ScalingConfigurationProperty
	_jsii_.Get(
		j,
		"scalingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Serverlessv2ScalingConfiguration() AwsRdsCluster_Serverlessv2ScalingConfigurationPropertyOutputReference {
	var returns AwsRdsCluster_Serverlessv2ScalingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"serverlessv2ScalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Serverlessv2ScalingConfigurationInput() *AwsRdsCluster_Serverlessv2ScalingConfigurationProperty {
	var returns *AwsRdsCluster_Serverlessv2ScalingConfigurationProperty
	_jsii_.Get(
		j,
		"serverlessv2ScalingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) SkipFinalSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) SkipFinalSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) SnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) SnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) SourceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) SourceRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) StorageEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) StorageEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) Timeouts() AwsRdsCluster_TimeoutsPropertyOutputReference {
	var returns AwsRdsCluster_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) UpgradeRolloutOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeRolloutOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster) VpcSecurityGroupIdsInput() *[]*string {
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
func NewAwsRdsCluster(scope constructs.Construct, id *string, config *AwsRdsClusterConfig) AwsRdsCluster {
	_init_.Initialize()

	if err := validateNewAwsRdsClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRdsCluster{}

	_jsii_.Create(
		"@cdktn/aws-rds.AwsRdsCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster aws_rds_cluster} Resource.
// Experimental.
func NewAwsRdsCluster_Override(a AwsRdsCluster, scope constructs.Construct, id *string, config *AwsRdsClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-rds.AwsRdsCluster",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetAllocatedStorage(val *float64) {
	if err := j.validateSetAllocatedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocatedStorage",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetAllowMajorVersionUpgrade(val interface{}) {
	if err := j.validateSetAllowMajorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMajorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetApplyImmediately(val interface{}) {
	if err := j.validateSetApplyImmediatelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyImmediately",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetAutoMinorVersionUpgrade(val interface{}) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetBacktrackWindow(val *float64) {
	if err := j.validateSetBacktrackWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backtrackWindow",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetBackupRetentionPeriod(val *float64) {
	if err := j.validateSetBackupRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetCaCertificateIdentifier(val *string) {
	if err := j.validateSetCaCertificateIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificateIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetClusterIdentifier(val *string) {
	if err := j.validateSetClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetClusterIdentifierPrefix(val *string) {
	if err := j.validateSetClusterIdentifierPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIdentifierPrefix",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetClusterMembers(val *[]*string) {
	if err := j.validateSetClusterMembersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterMembers",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetClusterScalabilityType(val *string) {
	if err := j.validateSetClusterScalabilityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterScalabilityType",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetCopyTagsToSnapshot(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshot",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetDatabaseInsightsMode(val *string) {
	if err := j.validateSetDatabaseInsightsModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseInsightsMode",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetDbClusterInstanceClass(val *string) {
	if err := j.validateSetDbClusterInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbClusterInstanceClass",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetDbClusterParameterGroupName(val *string) {
	if err := j.validateSetDbClusterParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbClusterParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetDbInstanceParameterGroupName(val *string) {
	if err := j.validateSetDbInstanceParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbInstanceParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetDbSubnetGroupName(val *string) {
	if err := j.validateSetDbSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetDbSystemId(val *string) {
	if err := j.validateSetDbSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSystemId",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetDeleteAutomatedBackups(val interface{}) {
	if err := j.validateSetDeleteAutomatedBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAutomatedBackups",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetDomainIamRoleName(val *string) {
	if err := j.validateSetDomainIamRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainIamRoleName",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetEnabledCloudwatchLogsExports(val *[]*string) {
	if err := j.validateSetEnabledCloudwatchLogsExportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledCloudwatchLogsExports",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetEnableGlobalWriteForwarding(val interface{}) {
	if err := j.validateSetEnableGlobalWriteForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableGlobalWriteForwarding",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetEnableHttpEndpoint(val interface{}) {
	if err := j.validateSetEnableHttpEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHttpEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetEnableLocalWriteForwarding(val interface{}) {
	if err := j.validateSetEnableLocalWriteForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLocalWriteForwarding",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetEngineLifecycleSupport(val *string) {
	if err := j.validateSetEngineLifecycleSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineLifecycleSupport",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetEngineMode(val *string) {
	if err := j.validateSetEngineModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineMode",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetFinalSnapshotIdentifier(val *string) {
	if err := j.validateSetFinalSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetGlobalClusterIdentifier(val *string) {
	if err := j.validateSetGlobalClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetIamDatabaseAuthenticationEnabled(val interface{}) {
	if err := j.validateSetIamDatabaseAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamDatabaseAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetIamRoles(val *[]*string) {
	if err := j.validateSetIamRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRoles",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetIops(val *float64) {
	if err := j.validateSetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetManageMasterUserPassword(val interface{}) {
	if err := j.validateSetManageMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageMasterUserPassword",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetMasterPassword(val *string) {
	if err := j.validateSetMasterPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterPassword",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetMasterPasswordWo(val *string) {
	if err := j.validateSetMasterPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterPasswordWo",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetMasterPasswordWoVersion(val *float64) {
	if err := j.validateSetMasterPasswordWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterPasswordWoVersion",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetMasterUsername(val *string) {
	if err := j.validateSetMasterUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetMasterUserSecretKmsKeyId(val *string) {
	if err := j.validateSetMasterUserSecretKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserSecretKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetMonitoringInterval(val *float64) {
	if err := j.validateSetMonitoringIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringInterval",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetMonitoringRoleArn(val *string) {
	if err := j.validateSetMonitoringRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetPerformanceInsightsEnabled(val interface{}) {
	if err := j.validateSetPerformanceInsightsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetPerformanceInsightsKmsKeyId(val *string) {
	if err := j.validateSetPerformanceInsightsKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetPerformanceInsightsRetentionPeriod(val *float64) {
	if err := j.validateSetPerformanceInsightsRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetPreferredBackupWindow(val *string) {
	if err := j.validateSetPreferredBackupWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetPreferredMaintenanceWindow(val *string) {
	if err := j.validateSetPreferredMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetReplicationSourceIdentifier(val *string) {
	if err := j.validateSetReplicationSourceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationSourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetSkipFinalSnapshot(val interface{}) {
	if err := j.validateSetSkipFinalSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipFinalSnapshot",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetSnapshotIdentifier(val *string) {
	if err := j.validateSetSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetSourceRegion(val *string) {
	if err := j.validateSetSourceRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceRegion",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetStorageEncrypted(val interface{}) {
	if err := j.validateSetStorageEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Generates CDKTN code for importing a AwsRdsCluster resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsRdsCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsRdsCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.AwsRdsCluster",
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
func AwsRdsCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsRdsCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.AwsRdsCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsRdsCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsRdsCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.AwsRdsCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsRdsCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsRdsCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-rds.AwsRdsCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsRdsCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-rds.AwsRdsCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsRdsCluster) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsRdsCluster) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsRdsCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRdsCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRdsCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRdsCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRdsCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRdsCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRdsCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRdsCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRdsCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRdsCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRdsCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsRdsCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRdsCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsRdsCluster) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsRdsCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsRdsCluster) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsRdsCluster) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsRdsCluster) PutRestoreToPointInTime(value *AwsRdsCluster_RestoreToPointInTimeProperty) {
	if err := a.validatePutRestoreToPointInTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRestoreToPointInTime",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRdsCluster) PutS3Import(value *AwsRdsCluster_S3ImportProperty) {
	if err := a.validatePutS3ImportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Import",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRdsCluster) PutScalingConfiguration(value *AwsRdsCluster_ScalingConfigurationProperty) {
	if err := a.validatePutScalingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScalingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRdsCluster) PutServerlessv2ScalingConfiguration(value *AwsRdsCluster_Serverlessv2ScalingConfigurationProperty) {
	if err := a.validatePutServerlessv2ScalingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServerlessv2ScalingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRdsCluster) PutTimeouts(value *AwsRdsCluster_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRdsCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetAllocatedStorage() {
	_jsii_.InvokeVoid(
		a,
		"resetAllocatedStorage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetAllowMajorVersionUpgrade() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowMajorVersionUpgrade",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetApplyImmediately() {
	_jsii_.InvokeVoid(
		a,
		"resetApplyImmediately",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetAvailabilityZones() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZones",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetBacktrackWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetBacktrackWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetBackupRetentionPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetBackupRetentionPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetCaCertificateIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetCaCertificateIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetClusterIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetClusterIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetClusterIdentifierPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetClusterIdentifierPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetClusterMembers() {
	_jsii_.InvokeVoid(
		a,
		"resetClusterMembers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetClusterScalabilityType() {
	_jsii_.InvokeVoid(
		a,
		"resetClusterScalabilityType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetCopyTagsToSnapshot() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyTagsToSnapshot",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetDatabaseInsightsMode() {
	_jsii_.InvokeVoid(
		a,
		"resetDatabaseInsightsMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		a,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetDbClusterInstanceClass() {
	_jsii_.InvokeVoid(
		a,
		"resetDbClusterInstanceClass",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetDbClusterParameterGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetDbClusterParameterGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetDbInstanceParameterGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetDbInstanceParameterGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetDbSubnetGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetDbSubnetGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetDbSystemId() {
	_jsii_.InvokeVoid(
		a,
		"resetDbSystemId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetDeleteAutomatedBackups() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteAutomatedBackups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		a,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetDomainIamRoleName() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainIamRoleName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetEnabledCloudwatchLogsExports() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabledCloudwatchLogsExports",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetEnableGlobalWriteForwarding() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableGlobalWriteForwarding",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetEnableHttpEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableHttpEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetEnableLocalWriteForwarding() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableLocalWriteForwarding",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetEngineLifecycleSupport() {
	_jsii_.InvokeVoid(
		a,
		"resetEngineLifecycleSupport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetEngineMode() {
	_jsii_.InvokeVoid(
		a,
		"resetEngineMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetFinalSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetFinalSnapshotIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetGlobalClusterIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetGlobalClusterIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetIamDatabaseAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetIamDatabaseAuthenticationEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetIamRoles() {
	_jsii_.InvokeVoid(
		a,
		"resetIamRoles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetIops() {
	_jsii_.InvokeVoid(
		a,
		"resetIops",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetManageMasterUserPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetManageMasterUserPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetMasterPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetMasterPasswordWo() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterPasswordWo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetMasterPasswordWoVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterPasswordWoVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetMasterUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetMasterUserSecretKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterUserSecretKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetMonitoringInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitoringInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetMonitoringRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitoringRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetNetworkType() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetPerformanceInsightsEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetPerformanceInsightsEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetPerformanceInsightsKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetPerformanceInsightsKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetPerformanceInsightsRetentionPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetPerformanceInsightsRetentionPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetPreferredBackupWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetPreferredBackupWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetPreferredMaintenanceWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetPreferredMaintenanceWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetReplicationSourceIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetReplicationSourceIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetRestoreToPointInTime() {
	_jsii_.InvokeVoid(
		a,
		"resetRestoreToPointInTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetS3Import() {
	_jsii_.InvokeVoid(
		a,
		"resetS3Import",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetScalingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetScalingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetServerlessv2ScalingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetServerlessv2ScalingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetSkipFinalSnapshot() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipFinalSnapshot",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetSourceRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetStorageEncrypted() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageEncrypted",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetStorageType() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRdsCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRdsCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRdsCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRdsCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRdsCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRdsCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

