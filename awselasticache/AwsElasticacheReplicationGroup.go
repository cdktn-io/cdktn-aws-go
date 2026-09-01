package awselasticache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselasticache/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awselasticache/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group aws_elasticache_replication_group}.
// Experimental.
type AwsElasticacheReplicationGroup interface {
	cdktn.TerraformResource
	// Experimental.
	ApplyImmediately() interface{}
	// Experimental.
	SetApplyImmediately(val interface{})
	// Experimental.
	ApplyImmediatelyInput() interface{}
	// Experimental.
	Arn() *string
	// Experimental.
	AtRestEncryptionEnabled() *string
	// Experimental.
	SetAtRestEncryptionEnabled(val *string)
	// Experimental.
	AtRestEncryptionEnabledInput() *string
	// Experimental.
	AuthToken() *string
	// Experimental.
	SetAuthToken(val *string)
	// Experimental.
	AuthTokenInput() *string
	// Experimental.
	AuthTokenUpdateStrategy() *string
	// Experimental.
	SetAuthTokenUpdateStrategy(val *string)
	// Experimental.
	AuthTokenUpdateStrategyInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	AuthTokenWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetAuthTokenWo(val *string)
	// Experimental.
	AuthTokenWoInput() *string
	// Experimental.
	AuthTokenWoVersion() *float64
	// Experimental.
	SetAuthTokenWoVersion(val *float64)
	// Experimental.
	AuthTokenWoVersionInput() *float64
	// Experimental.
	AutomaticFailoverEnabled() interface{}
	// Experimental.
	SetAutomaticFailoverEnabled(val interface{})
	// Experimental.
	AutomaticFailoverEnabledInput() interface{}
	// Experimental.
	AutoMinorVersionUpgrade() *string
	// Experimental.
	SetAutoMinorVersionUpgrade(val *string)
	// Experimental.
	AutoMinorVersionUpgradeInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClusterEnabled() cdktn.IResolvable
	// Experimental.
	ClusterMode() *string
	// Experimental.
	SetClusterMode(val *string)
	// Experimental.
	ClusterModeInput() *string
	// Experimental.
	ConfigurationEndpointAddress() *string
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
	DataTieringEnabled() interface{}
	// Experimental.
	SetDataTieringEnabled(val interface{})
	// Experimental.
	DataTieringEnabledInput() interface{}
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
	Durability() *string
	// Experimental.
	SetDurability(val *string)
	// Experimental.
	DurabilityInput() *string
	// Experimental.
	Engine() *string
	// Experimental.
	SetEngine(val *string)
	// Experimental.
	EngineInput() *string
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
	GlobalReplicationGroupId() *string
	// Experimental.
	SetGlobalReplicationGroupId(val *string)
	// Experimental.
	GlobalReplicationGroupIdInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	IpDiscovery() *string
	// Experimental.
	SetIpDiscovery(val *string)
	// Experimental.
	IpDiscoveryInput() *string
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
	LogDeliveryConfiguration() AwsElasticacheReplicationGroup_LogDeliveryConfigurationPropertyList
	// Experimental.
	LogDeliveryConfigurationInput() interface{}
	// Experimental.
	MaintenanceWindow() *string
	// Experimental.
	SetMaintenanceWindow(val *string)
	// Experimental.
	MaintenanceWindowInput() *string
	// Experimental.
	MemberClusters() *[]*string
	// Experimental.
	MultiAzEnabled() interface{}
	// Experimental.
	SetMultiAzEnabled(val interface{})
	// Experimental.
	MultiAzEnabledInput() interface{}
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
	NodeGroupConfiguration() AwsElasticacheReplicationGroup_NodeGroupConfigurationPropertyList
	// Experimental.
	NodeGroupConfigurationInput() interface{}
	// Experimental.
	NodeType() *string
	// Experimental.
	SetNodeType(val *string)
	// Experimental.
	NodeTypeInput() *string
	// Experimental.
	NotificationTopicArn() *string
	// Experimental.
	SetNotificationTopicArn(val *string)
	// Experimental.
	NotificationTopicArnInput() *string
	// Experimental.
	NumCacheClusters() *float64
	// Experimental.
	SetNumCacheClusters(val *float64)
	// Experimental.
	NumCacheClustersInput() *float64
	// Experimental.
	NumNodeGroups() *float64
	// Experimental.
	SetNumNodeGroups(val *float64)
	// Experimental.
	NumNodeGroupsInput() *float64
	// Experimental.
	ParameterGroupName() *string
	// Experimental.
	SetParameterGroupName(val *string)
	// Experimental.
	ParameterGroupNameInput() *string
	// Experimental.
	Port() *float64
	// Experimental.
	SetPort(val *float64)
	// Experimental.
	PortInput() *float64
	// Experimental.
	PreferredCacheClusterAzs() *[]*string
	// Experimental.
	SetPreferredCacheClusterAzs(val *[]*string)
	// Experimental.
	PreferredCacheClusterAzsInput() *[]*string
	// Experimental.
	PrimaryEndpointAddress() *string
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
	ReaderEndpointAddress() *string
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	ReplicasPerNodeGroup() *float64
	// Experimental.
	SetReplicasPerNodeGroup(val *float64)
	// Experimental.
	ReplicasPerNodeGroupInput() *float64
	// Experimental.
	ReplicationGroupId() *string
	// Experimental.
	SetReplicationGroupId(val *string)
	// Experimental.
	ReplicationGroupIdInput() *string
	// Experimental.
	SecurityGroupIds() *[]*string
	// Experimental.
	SetSecurityGroupIds(val *[]*string)
	// Experimental.
	SecurityGroupIdsInput() *[]*string
	// Experimental.
	SecurityGroupNames() *[]*string
	// Experimental.
	SetSecurityGroupNames(val *[]*string)
	// Experimental.
	SecurityGroupNamesInput() *[]*string
	// Experimental.
	SnapshotArns() *[]*string
	// Experimental.
	SetSnapshotArns(val *[]*string)
	// Experimental.
	SnapshotArnsInput() *[]*string
	// Experimental.
	SnapshotName() *string
	// Experimental.
	SetSnapshotName(val *string)
	// Experimental.
	SnapshotNameInput() *string
	// Experimental.
	SnapshotRetentionLimit() *float64
	// Experimental.
	SetSnapshotRetentionLimit(val *float64)
	// Experimental.
	SnapshotRetentionLimitInput() *float64
	// Experimental.
	SnapshotWindow() *string
	// Experimental.
	SetSnapshotWindow(val *string)
	// Experimental.
	SnapshotWindowInput() *string
	// Experimental.
	SubnetGroupName() *string
	// Experimental.
	SetSubnetGroupName(val *string)
	// Experimental.
	SubnetGroupNameInput() *string
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
	Timeouts() AwsElasticacheReplicationGroup_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TransitEncryptionEnabled() interface{}
	// Experimental.
	SetTransitEncryptionEnabled(val interface{})
	// Experimental.
	TransitEncryptionEnabledInput() interface{}
	// Experimental.
	TransitEncryptionMode() *string
	// Experimental.
	SetTransitEncryptionMode(val *string)
	// Experimental.
	TransitEncryptionModeInput() *string
	// Experimental.
	UserGroupIds() *[]*string
	// Experimental.
	SetUserGroupIds(val *[]*string)
	// Experimental.
	UserGroupIdsInput() *[]*string
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
	PutLogDeliveryConfiguration(value interface{})
	// Experimental.
	PutNodeGroupConfiguration(value interface{})
	// Experimental.
	PutTimeouts(value *AwsElasticacheReplicationGroup_TimeoutsProperty)
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
	ResetApplyImmediately()
	// Experimental.
	ResetAtRestEncryptionEnabled()
	// Experimental.
	ResetAuthToken()
	// Experimental.
	ResetAuthTokenUpdateStrategy()
	// Experimental.
	ResetAuthTokenWo()
	// Experimental.
	ResetAuthTokenWoVersion()
	// Experimental.
	ResetAutomaticFailoverEnabled()
	// Experimental.
	ResetAutoMinorVersionUpgrade()
	// Experimental.
	ResetClusterMode()
	// Experimental.
	ResetDataTieringEnabled()
	// Experimental.
	ResetDurability()
	// Experimental.
	ResetEngine()
	// Experimental.
	ResetEngineVersion()
	// Experimental.
	ResetFinalSnapshotIdentifier()
	// Experimental.
	ResetGlobalReplicationGroupId()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIpDiscovery()
	// Experimental.
	ResetKmsKeyId()
	// Experimental.
	ResetLogDeliveryConfiguration()
	// Experimental.
	ResetMaintenanceWindow()
	// Experimental.
	ResetMultiAzEnabled()
	// Experimental.
	ResetNetworkType()
	// Experimental.
	ResetNodeGroupConfiguration()
	// Experimental.
	ResetNodeType()
	// Experimental.
	ResetNotificationTopicArn()
	// Experimental.
	ResetNumCacheClusters()
	// Experimental.
	ResetNumNodeGroups()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetParameterGroupName()
	// Experimental.
	ResetPort()
	// Experimental.
	ResetPreferredCacheClusterAzs()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetReplicasPerNodeGroup()
	// Experimental.
	ResetSecurityGroupIds()
	// Experimental.
	ResetSecurityGroupNames()
	// Experimental.
	ResetSnapshotArns()
	// Experimental.
	ResetSnapshotName()
	// Experimental.
	ResetSnapshotRetentionLimit()
	// Experimental.
	ResetSnapshotWindow()
	// Experimental.
	ResetSubnetGroupName()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTransitEncryptionEnabled()
	// Experimental.
	ResetTransitEncryptionMode()
	// Experimental.
	ResetUserGroupIds()
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

// The jsii proxy struct for AwsElasticacheReplicationGroup
type jsiiProxy_AwsElasticacheReplicationGroup struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) ApplyImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) ApplyImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) AtRestEncryptionEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"atRestEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) AtRestEncryptionEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"atRestEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) AuthToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) AuthTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) AuthTokenUpdateStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenUpdateStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) AuthTokenUpdateStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenUpdateStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) AuthTokenWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) AuthTokenWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTokenWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) AuthTokenWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authTokenWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) AuthTokenWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"authTokenWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) AutomaticFailoverEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticFailoverEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) AutomaticFailoverEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticFailoverEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) AutoMinorVersionUpgrade() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) AutoMinorVersionUpgradeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) ClusterEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"clusterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) ClusterMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) ClusterModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) ConfigurationEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) DataTieringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTieringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) DataTieringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTieringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) Durability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) DurabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) EngineVersionActual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionActual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) FinalSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) FinalSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) GlobalReplicationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalReplicationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) GlobalReplicationGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalReplicationGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) IpDiscovery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipDiscovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) IpDiscoveryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipDiscoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) LogDeliveryConfiguration() AwsElasticacheReplicationGroup_LogDeliveryConfigurationPropertyList {
	var returns AwsElasticacheReplicationGroup_LogDeliveryConfigurationPropertyList
	_jsii_.Get(
		j,
		"logDeliveryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) LogDeliveryConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDeliveryConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) MaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) MaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) MemberClusters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"memberClusters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) MultiAzEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAzEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) MultiAzEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAzEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) NodeGroupConfiguration() AwsElasticacheReplicationGroup_NodeGroupConfigurationPropertyList {
	var returns AwsElasticacheReplicationGroup_NodeGroupConfigurationPropertyList
	_jsii_.Get(
		j,
		"nodeGroupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) NodeGroupConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodeGroupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) NodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) NodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) NotificationTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) NotificationTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) NumCacheClusters() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCacheClusters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) NumCacheClustersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCacheClustersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) NumNodeGroups() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numNodeGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) NumNodeGroupsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numNodeGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) ParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) ParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) PreferredCacheClusterAzs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredCacheClusterAzs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) PreferredCacheClusterAzsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredCacheClusterAzsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) PrimaryEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) ReaderEndpointAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readerEndpointAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) ReplicasPerNodeGroup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicasPerNodeGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) ReplicasPerNodeGroupInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicasPerNodeGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) ReplicationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) ReplicationGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) SecurityGroupNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) SecurityGroupNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) SnapshotArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snapshotArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) SnapshotArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snapshotArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) SnapshotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) SnapshotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) SnapshotRetentionLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) SnapshotRetentionLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) SnapshotWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) SnapshotWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) SubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) SubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) Timeouts() AwsElasticacheReplicationGroup_TimeoutsPropertyOutputReference {
	var returns AwsElasticacheReplicationGroup_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) TransitEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) TransitEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) TransitEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) TransitEncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) UserGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup) UserGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userGroupIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group aws_elasticache_replication_group} Resource.
// Experimental.
func NewAwsElasticacheReplicationGroup(scope constructs.Construct, id *string, config *AwsElasticacheReplicationGroupConfig) AwsElasticacheReplicationGroup {
	_init_.Initialize()

	if err := validateNewAwsElasticacheReplicationGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsElasticacheReplicationGroup{}

	_jsii_.Create(
		"@cdktn/aws-elasticache.AwsElasticacheReplicationGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group aws_elasticache_replication_group} Resource.
// Experimental.
func NewAwsElasticacheReplicationGroup_Override(a AwsElasticacheReplicationGroup, scope constructs.Construct, id *string, config *AwsElasticacheReplicationGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elasticache.AwsElasticacheReplicationGroup",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetApplyImmediately(val interface{}) {
	if err := j.validateSetApplyImmediatelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyImmediately",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetAtRestEncryptionEnabled(val *string) {
	if err := j.validateSetAtRestEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"atRestEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetAuthToken(val *string) {
	if err := j.validateSetAuthTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authToken",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetAuthTokenUpdateStrategy(val *string) {
	if err := j.validateSetAuthTokenUpdateStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authTokenUpdateStrategy",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetAuthTokenWo(val *string) {
	if err := j.validateSetAuthTokenWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authTokenWo",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetAuthTokenWoVersion(val *float64) {
	if err := j.validateSetAuthTokenWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authTokenWoVersion",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetAutomaticFailoverEnabled(val interface{}) {
	if err := j.validateSetAutomaticFailoverEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticFailoverEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetAutoMinorVersionUpgrade(val *string) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetClusterMode(val *string) {
	if err := j.validateSetClusterModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterMode",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetDataTieringEnabled(val interface{}) {
	if err := j.validateSetDataTieringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTieringEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetDurability(val *string) {
	if err := j.validateSetDurabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durability",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetFinalSnapshotIdentifier(val *string) {
	if err := j.validateSetFinalSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetGlobalReplicationGroupId(val *string) {
	if err := j.validateSetGlobalReplicationGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalReplicationGroupId",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetIpDiscovery(val *string) {
	if err := j.validateSetIpDiscoveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipDiscovery",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetMaintenanceWindow(val *string) {
	if err := j.validateSetMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetMultiAzEnabled(val interface{}) {
	if err := j.validateSetMultiAzEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiAzEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetNodeType(val *string) {
	if err := j.validateSetNodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeType",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetNotificationTopicArn(val *string) {
	if err := j.validateSetNotificationTopicArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationTopicArn",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetNumCacheClusters(val *float64) {
	if err := j.validateSetNumCacheClustersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numCacheClusters",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetNumNodeGroups(val *float64) {
	if err := j.validateSetNumNodeGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numNodeGroups",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetParameterGroupName(val *string) {
	if err := j.validateSetParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetPreferredCacheClusterAzs(val *[]*string) {
	if err := j.validateSetPreferredCacheClusterAzsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredCacheClusterAzs",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetReplicasPerNodeGroup(val *float64) {
	if err := j.validateSetReplicasPerNodeGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicasPerNodeGroup",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetReplicationGroupId(val *string) {
	if err := j.validateSetReplicationGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationGroupId",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetSecurityGroupNames(val *[]*string) {
	if err := j.validateSetSecurityGroupNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupNames",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetSnapshotArns(val *[]*string) {
	if err := j.validateSetSnapshotArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotArns",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetSnapshotName(val *string) {
	if err := j.validateSetSnapshotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotName",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetSnapshotRetentionLimit(val *float64) {
	if err := j.validateSetSnapshotRetentionLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotRetentionLimit",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetSnapshotWindow(val *string) {
	if err := j.validateSetSnapshotWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotWindow",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetSubnetGroupName(val *string) {
	if err := j.validateSetSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetTransitEncryptionEnabled(val interface{}) {
	if err := j.validateSetTransitEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetTransitEncryptionMode(val *string) {
	if err := j.validateSetTransitEncryptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitEncryptionMode",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheReplicationGroup)SetUserGroupIds(val *[]*string) {
	if err := j.validateSetUserGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userGroupIds",
		val,
	)
}

// Generates CDKTN code for importing a AwsElasticacheReplicationGroup resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsElasticacheReplicationGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsElasticacheReplicationGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-elasticache.AwsElasticacheReplicationGroup",
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
func AwsElasticacheReplicationGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsElasticacheReplicationGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elasticache.AwsElasticacheReplicationGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsElasticacheReplicationGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsElasticacheReplicationGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elasticache.AwsElasticacheReplicationGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsElasticacheReplicationGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsElasticacheReplicationGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elasticache.AwsElasticacheReplicationGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsElasticacheReplicationGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-elasticache.AwsElasticacheReplicationGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsElasticacheReplicationGroup) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsElasticacheReplicationGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsElasticacheReplicationGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsElasticacheReplicationGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsElasticacheReplicationGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsElasticacheReplicationGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsElasticacheReplicationGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsElasticacheReplicationGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsElasticacheReplicationGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsElasticacheReplicationGroup) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsElasticacheReplicationGroup) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) PutLogDeliveryConfiguration(value interface{}) {
	if err := a.validatePutLogDeliveryConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogDeliveryConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) PutNodeGroupConfiguration(value interface{}) {
	if err := a.validatePutNodeGroupConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNodeGroupConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) PutTimeouts(value *AwsElasticacheReplicationGroup_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetApplyImmediately() {
	_jsii_.InvokeVoid(
		a,
		"resetApplyImmediately",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetAtRestEncryptionEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAtRestEncryptionEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetAuthToken() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetAuthTokenUpdateStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthTokenUpdateStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetAuthTokenWo() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthTokenWo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetAuthTokenWoVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthTokenWoVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetAutomaticFailoverEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAutomaticFailoverEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetClusterMode() {
	_jsii_.InvokeVoid(
		a,
		"resetClusterMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetDataTieringEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDataTieringEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetDurability() {
	_jsii_.InvokeVoid(
		a,
		"resetDurability",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetEngine() {
	_jsii_.InvokeVoid(
		a,
		"resetEngine",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetFinalSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetFinalSnapshotIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetGlobalReplicationGroupId() {
	_jsii_.InvokeVoid(
		a,
		"resetGlobalReplicationGroupId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetIpDiscovery() {
	_jsii_.InvokeVoid(
		a,
		"resetIpDiscovery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetLogDeliveryConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLogDeliveryConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetMultiAzEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetMultiAzEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetNetworkType() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetNodeGroupConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetNodeGroupConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetNodeType() {
	_jsii_.InvokeVoid(
		a,
		"resetNodeType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetNotificationTopicArn() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationTopicArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetNumCacheClusters() {
	_jsii_.InvokeVoid(
		a,
		"resetNumCacheClusters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetNumNodeGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetNumNodeGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetParameterGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetParameterGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetPreferredCacheClusterAzs() {
	_jsii_.InvokeVoid(
		a,
		"resetPreferredCacheClusterAzs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetReplicasPerNodeGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetReplicasPerNodeGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetSecurityGroupNames() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetSnapshotArns() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetSnapshotName() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetSnapshotRetentionLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotRetentionLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetSnapshotWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetSubnetGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetTransitEncryptionEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTransitEncryptionEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetTransitEncryptionMode() {
	_jsii_.InvokeVoid(
		a,
		"resetTransitEncryptionMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ResetUserGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetUserGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticacheReplicationGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

