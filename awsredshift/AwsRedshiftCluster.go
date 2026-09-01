package awsredshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsredshift/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsredshift/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster aws_redshift_cluster}.
// Experimental.
type AwsRedshiftCluster interface {
	cdktn.TerraformResource
	// Experimental.
	AllowVersionUpgrade() interface{}
	// Experimental.
	SetAllowVersionUpgrade(val interface{})
	// Experimental.
	AllowVersionUpgradeInput() interface{}
	// Experimental.
	ApplyImmediately() interface{}
	// Experimental.
	SetApplyImmediately(val interface{})
	// Experimental.
	ApplyImmediatelyInput() interface{}
	// Experimental.
	AquaConfigurationStatus() *string
	// Experimental.
	SetAquaConfigurationStatus(val *string)
	// Experimental.
	AquaConfigurationStatusInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	AutomatedSnapshotRetentionPeriod() *float64
	// Experimental.
	SetAutomatedSnapshotRetentionPeriod(val *float64)
	// Experimental.
	AutomatedSnapshotRetentionPeriodInput() *float64
	// Experimental.
	AvailabilityZone() *string
	// Experimental.
	SetAvailabilityZone(val *string)
	// Experimental.
	AvailabilityZoneInput() *string
	// Experimental.
	AvailabilityZoneRelocationEnabled() interface{}
	// Experimental.
	SetAvailabilityZoneRelocationEnabled(val interface{})
	// Experimental.
	AvailabilityZoneRelocationEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClusterIdentifier() *string
	// Experimental.
	SetClusterIdentifier(val *string)
	// Experimental.
	ClusterIdentifierInput() *string
	// Experimental.
	ClusterNamespaceArn() *string
	// Experimental.
	ClusterNodes() AwsRedshiftCluster_ClusterNodesPropertyList
	// Experimental.
	ClusterParameterGroupName() *string
	// Experimental.
	SetClusterParameterGroupName(val *string)
	// Experimental.
	ClusterParameterGroupNameInput() *string
	// Experimental.
	ClusterPublicKey() *string
	// Experimental.
	ClusterRevisionNumber() *string
	// Experimental.
	ClusterSubnetGroupName() *string
	// Experimental.
	SetClusterSubnetGroupName(val *string)
	// Experimental.
	ClusterSubnetGroupNameInput() *string
	// Experimental.
	ClusterType() *string
	// Experimental.
	SetClusterType(val *string)
	// Experimental.
	ClusterTypeInput() *string
	// Experimental.
	ClusterVersion() *string
	// Experimental.
	SetClusterVersion(val *string)
	// Experimental.
	ClusterVersionInput() *string
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
	DatabaseName() *string
	// Experimental.
	SetDatabaseName(val *string)
	// Experimental.
	DatabaseNameInput() *string
	// Experimental.
	DefaultIamRoleArn() *string
	// Experimental.
	SetDefaultIamRoleArn(val *string)
	// Experimental.
	DefaultIamRoleArnInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DnsName() *string
	// Experimental.
	ElasticIp() *string
	// Experimental.
	SetElasticIp(val *string)
	// Experimental.
	ElasticIpInput() *string
	// Experimental.
	Encrypted() *string
	// Experimental.
	SetEncrypted(val *string)
	// Experimental.
	EncryptedInput() *string
	// Experimental.
	Endpoint() *string
	// Experimental.
	EnhancedVpcRouting() interface{}
	// Experimental.
	SetEnhancedVpcRouting(val interface{})
	// Experimental.
	EnhancedVpcRoutingInput() interface{}
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
	MaintenanceTrackName() *string
	// Experimental.
	SetMaintenanceTrackName(val *string)
	// Experimental.
	MaintenanceTrackNameInput() *string
	// Experimental.
	ManageMasterPassword() interface{}
	// Experimental.
	SetManageMasterPassword(val interface{})
	// Experimental.
	ManageMasterPasswordInput() interface{}
	// Experimental.
	ManualSnapshotRetentionPeriod() *float64
	// Experimental.
	SetManualSnapshotRetentionPeriod(val *float64)
	// Experimental.
	ManualSnapshotRetentionPeriodInput() *float64
	// Experimental.
	MasterPassword() *string
	// Experimental.
	SetMasterPassword(val *string)
	// Experimental.
	MasterPasswordInput() *string
	// Experimental.
	MasterPasswordSecretArn() *string
	// Experimental.
	MasterPasswordSecretKmsKeyId() *string
	// Experimental.
	SetMasterPasswordSecretKmsKeyId(val *string)
	// Experimental.
	MasterPasswordSecretKmsKeyIdInput() *string
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
	MultiAz() interface{}
	// Experimental.
	SetMultiAz(val interface{})
	// Experimental.
	MultiAzInput() interface{}
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	NodeType() *string
	// Experimental.
	SetNodeType(val *string)
	// Experimental.
	NodeTypeInput() *string
	// Experimental.
	NumberOfNodes() *float64
	// Experimental.
	SetNumberOfNodes(val *float64)
	// Experimental.
	NumberOfNodesInput() *float64
	// Experimental.
	OwnerAccount() *string
	// Experimental.
	SetOwnerAccount(val *string)
	// Experimental.
	OwnerAccountInput() *string
	// Experimental.
	Port() *float64
	// Experimental.
	SetPort(val *float64)
	// Experimental.
	PortInput() *float64
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
	SkipFinalSnapshot() interface{}
	// Experimental.
	SetSkipFinalSnapshot(val interface{})
	// Experimental.
	SkipFinalSnapshotInput() interface{}
	// Experimental.
	SnapshotArn() *string
	// Experimental.
	SetSnapshotArn(val *string)
	// Experimental.
	SnapshotArnInput() *string
	// Experimental.
	SnapshotClusterIdentifier() *string
	// Experimental.
	SetSnapshotClusterIdentifier(val *string)
	// Experimental.
	SnapshotClusterIdentifierInput() *string
	// Experimental.
	SnapshotIdentifier() *string
	// Experimental.
	SetSnapshotIdentifier(val *string)
	// Experimental.
	SnapshotIdentifierInput() *string
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
	Timeouts() AwsRedshiftCluster_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
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
	PutTimeouts(value *AwsRedshiftCluster_TimeoutsProperty)
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
	ResetAllowVersionUpgrade()
	// Experimental.
	ResetApplyImmediately()
	// Experimental.
	ResetAquaConfigurationStatus()
	// Experimental.
	ResetAutomatedSnapshotRetentionPeriod()
	// Experimental.
	ResetAvailabilityZone()
	// Experimental.
	ResetAvailabilityZoneRelocationEnabled()
	// Experimental.
	ResetClusterParameterGroupName()
	// Experimental.
	ResetClusterSubnetGroupName()
	// Experimental.
	ResetClusterType()
	// Experimental.
	ResetClusterVersion()
	// Experimental.
	ResetDatabaseName()
	// Experimental.
	ResetDefaultIamRoleArn()
	// Experimental.
	ResetElasticIp()
	// Experimental.
	ResetEncrypted()
	// Experimental.
	ResetEnhancedVpcRouting()
	// Experimental.
	ResetFinalSnapshotIdentifier()
	// Experimental.
	ResetIamRoles()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKmsKeyId()
	// Experimental.
	ResetMaintenanceTrackName()
	// Experimental.
	ResetManageMasterPassword()
	// Experimental.
	ResetManualSnapshotRetentionPeriod()
	// Experimental.
	ResetMasterPassword()
	// Experimental.
	ResetMasterPasswordSecretKmsKeyId()
	// Experimental.
	ResetMasterPasswordWo()
	// Experimental.
	ResetMasterPasswordWoVersion()
	// Experimental.
	ResetMasterUsername()
	// Experimental.
	ResetMultiAz()
	// Experimental.
	ResetNumberOfNodes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetOwnerAccount()
	// Experimental.
	ResetPort()
	// Experimental.
	ResetPreferredMaintenanceWindow()
	// Experimental.
	ResetPubliclyAccessible()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSkipFinalSnapshot()
	// Experimental.
	ResetSnapshotArn()
	// Experimental.
	ResetSnapshotClusterIdentifier()
	// Experimental.
	ResetSnapshotIdentifier()
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

// The jsii proxy struct for AwsRedshiftCluster
type jsiiProxy_AwsRedshiftCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsRedshiftCluster) AllowVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) AllowVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ApplyImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ApplyImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) AquaConfigurationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aquaConfigurationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) AquaConfigurationStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aquaConfigurationStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) AutomatedSnapshotRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automatedSnapshotRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) AutomatedSnapshotRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automatedSnapshotRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) AvailabilityZoneRelocationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availabilityZoneRelocationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) AvailabilityZoneRelocationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availabilityZoneRelocationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ClusterNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNamespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ClusterNodes() AwsRedshiftCluster_ClusterNodesPropertyList {
	var returns AwsRedshiftCluster_ClusterNodesPropertyList
	_jsii_.Get(
		j,
		"clusterNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ClusterParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ClusterParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ClusterPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ClusterRevisionNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterRevisionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ClusterSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ClusterSubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSubnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ClusterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ClusterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ClusterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ClusterVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) DefaultIamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultIamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) DefaultIamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultIamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ElasticIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ElasticIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) Encrypted() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) EncryptedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) EnhancedVpcRouting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedVpcRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) EnhancedVpcRoutingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedVpcRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) FinalSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) FinalSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) IamRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) IamRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) MaintenanceTrackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceTrackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) MaintenanceTrackNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceTrackNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ManageMasterPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ManageMasterPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ManualSnapshotRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"manualSnapshotRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) ManualSnapshotRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"manualSnapshotRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) MasterPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) MasterPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) MasterPasswordSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) MasterPasswordSecretKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordSecretKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) MasterPasswordSecretKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordSecretKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) MasterPasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) MasterPasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) MasterPasswordWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterPasswordWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) MasterPasswordWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"masterPasswordWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) MasterUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) MultiAz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) MultiAzInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) NodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) NodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) NumberOfNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) NumberOfNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) OwnerAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) OwnerAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) PreferredMaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) PubliclyAccessibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) SkipFinalSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) SkipFinalSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) SnapshotArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) SnapshotArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) SnapshotClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) SnapshotClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) SnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) SnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) Timeouts() AwsRedshiftCluster_TimeoutsPropertyOutputReference {
	var returns AwsRedshiftCluster_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRedshiftCluster) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster aws_redshift_cluster} Resource.
// Experimental.
func NewAwsRedshiftCluster(scope constructs.Construct, id *string, config *AwsRedshiftClusterConfig) AwsRedshiftCluster {
	_init_.Initialize()

	if err := validateNewAwsRedshiftClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRedshiftCluster{}

	_jsii_.Create(
		"@cdktn/aws-redshift.AwsRedshiftCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster aws_redshift_cluster} Resource.
// Experimental.
func NewAwsRedshiftCluster_Override(a AwsRedshiftCluster, scope constructs.Construct, id *string, config *AwsRedshiftClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-redshift.AwsRedshiftCluster",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetAllowVersionUpgrade(val interface{}) {
	if err := j.validateSetAllowVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetApplyImmediately(val interface{}) {
	if err := j.validateSetApplyImmediatelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyImmediately",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetAquaConfigurationStatus(val *string) {
	if err := j.validateSetAquaConfigurationStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aquaConfigurationStatus",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetAutomatedSnapshotRetentionPeriod(val *float64) {
	if err := j.validateSetAutomatedSnapshotRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automatedSnapshotRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetAvailabilityZoneRelocationEnabled(val interface{}) {
	if err := j.validateSetAvailabilityZoneRelocationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneRelocationEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetClusterIdentifier(val *string) {
	if err := j.validateSetClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetClusterParameterGroupName(val *string) {
	if err := j.validateSetClusterParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetClusterSubnetGroupName(val *string) {
	if err := j.validateSetClusterSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetClusterType(val *string) {
	if err := j.validateSetClusterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterType",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetClusterVersion(val *string) {
	if err := j.validateSetClusterVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterVersion",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetDefaultIamRoleArn(val *string) {
	if err := j.validateSetDefaultIamRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultIamRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetElasticIp(val *string) {
	if err := j.validateSetElasticIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticIp",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetEncrypted(val *string) {
	if err := j.validateSetEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetEnhancedVpcRouting(val interface{}) {
	if err := j.validateSetEnhancedVpcRoutingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enhancedVpcRouting",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetFinalSnapshotIdentifier(val *string) {
	if err := j.validateSetFinalSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetIamRoles(val *[]*string) {
	if err := j.validateSetIamRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRoles",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetMaintenanceTrackName(val *string) {
	if err := j.validateSetMaintenanceTrackNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceTrackName",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetManageMasterPassword(val interface{}) {
	if err := j.validateSetManageMasterPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageMasterPassword",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetManualSnapshotRetentionPeriod(val *float64) {
	if err := j.validateSetManualSnapshotRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manualSnapshotRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetMasterPassword(val *string) {
	if err := j.validateSetMasterPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterPassword",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetMasterPasswordSecretKmsKeyId(val *string) {
	if err := j.validateSetMasterPasswordSecretKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterPasswordSecretKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetMasterPasswordWo(val *string) {
	if err := j.validateSetMasterPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterPasswordWo",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetMasterPasswordWoVersion(val *float64) {
	if err := j.validateSetMasterPasswordWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterPasswordWoVersion",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetMasterUsername(val *string) {
	if err := j.validateSetMasterUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetMultiAz(val interface{}) {
	if err := j.validateSetMultiAzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiAz",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetNodeType(val *string) {
	if err := j.validateSetNodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeType",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetNumberOfNodes(val *float64) {
	if err := j.validateSetNumberOfNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfNodes",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetOwnerAccount(val *string) {
	if err := j.validateSetOwnerAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ownerAccount",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetPreferredMaintenanceWindow(val *string) {
	if err := j.validateSetPreferredMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetSkipFinalSnapshot(val interface{}) {
	if err := j.validateSetSkipFinalSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipFinalSnapshot",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetSnapshotArn(val *string) {
	if err := j.validateSetSnapshotArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotArn",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetSnapshotClusterIdentifier(val *string) {
	if err := j.validateSetSnapshotClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetSnapshotIdentifier(val *string) {
	if err := j.validateSetSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsRedshiftCluster)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Generates CDKTN code for importing a AwsRedshiftCluster resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsRedshiftCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsRedshiftCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-redshift.AwsRedshiftCluster",
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
func AwsRedshiftCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsRedshiftCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-redshift.AwsRedshiftCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsRedshiftCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsRedshiftCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-redshift.AwsRedshiftCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsRedshiftCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsRedshiftCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-redshift.AwsRedshiftCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsRedshiftCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-redshift.AwsRedshiftCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsRedshiftCluster) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRedshiftCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRedshiftCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRedshiftCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRedshiftCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRedshiftCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRedshiftCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRedshiftCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRedshiftCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRedshiftCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRedshiftCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRedshiftCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsRedshiftCluster) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) PutTimeouts(value *AwsRedshiftCluster_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetAllowVersionUpgrade() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowVersionUpgrade",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetApplyImmediately() {
	_jsii_.InvokeVoid(
		a,
		"resetApplyImmediately",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetAquaConfigurationStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetAquaConfigurationStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetAutomatedSnapshotRetentionPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetAutomatedSnapshotRetentionPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetAvailabilityZoneRelocationEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZoneRelocationEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetClusterParameterGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetClusterParameterGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetClusterSubnetGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetClusterSubnetGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetClusterType() {
	_jsii_.InvokeVoid(
		a,
		"resetClusterType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetClusterVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetClusterVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		a,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetDefaultIamRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultIamRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetElasticIp() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticIp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetEncrypted() {
	_jsii_.InvokeVoid(
		a,
		"resetEncrypted",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetEnhancedVpcRouting() {
	_jsii_.InvokeVoid(
		a,
		"resetEnhancedVpcRouting",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetFinalSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetFinalSnapshotIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetIamRoles() {
	_jsii_.InvokeVoid(
		a,
		"resetIamRoles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetMaintenanceTrackName() {
	_jsii_.InvokeVoid(
		a,
		"resetMaintenanceTrackName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetManageMasterPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetManageMasterPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetManualSnapshotRetentionPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetManualSnapshotRetentionPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetMasterPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetMasterPasswordSecretKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterPasswordSecretKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetMasterPasswordWo() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterPasswordWo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetMasterPasswordWoVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterPasswordWoVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetMasterUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetMasterUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetMultiAz() {
	_jsii_.InvokeVoid(
		a,
		"resetMultiAz",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetNumberOfNodes() {
	_jsii_.InvokeVoid(
		a,
		"resetNumberOfNodes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetOwnerAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetOwnerAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetPreferredMaintenanceWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetPreferredMaintenanceWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetPubliclyAccessible() {
	_jsii_.InvokeVoid(
		a,
		"resetPubliclyAccessible",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetSkipFinalSnapshot() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipFinalSnapshot",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetSnapshotArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetSnapshotClusterIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotClusterIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRedshiftCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRedshiftCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRedshiftCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRedshiftCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRedshiftCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRedshiftCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRedshiftCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

