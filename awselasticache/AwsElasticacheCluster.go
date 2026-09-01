package awselasticache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselasticache/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awselasticache/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_cluster aws_elasticache_cluster}.
// Experimental.
type AwsElasticacheCluster interface {
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
	AutoMinorVersionUpgrade() *string
	// Experimental.
	SetAutoMinorVersionUpgrade(val *string)
	// Experimental.
	AutoMinorVersionUpgradeInput() *string
	// Experimental.
	AvailabilityZone() *string
	// Experimental.
	SetAvailabilityZone(val *string)
	// Experimental.
	AvailabilityZoneInput() *string
	// Experimental.
	AzMode() *string
	// Experimental.
	SetAzMode(val *string)
	// Experimental.
	AzModeInput() *string
	// Experimental.
	CacheNodes() AwsElasticacheCluster_CacheNodesPropertyList
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClusterAddress() *string
	// Experimental.
	ClusterId() *string
	// Experimental.
	SetClusterId(val *string)
	// Experimental.
	ClusterIdInput() *string
	// Experimental.
	ConfigurationEndpoint() *string
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
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LogDeliveryConfiguration() AwsElasticacheCluster_LogDeliveryConfigurationPropertyList
	// Experimental.
	LogDeliveryConfigurationInput() interface{}
	// Experimental.
	MaintenanceWindow() *string
	// Experimental.
	SetMaintenanceWindow(val *string)
	// Experimental.
	MaintenanceWindowInput() *string
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
	NumCacheNodes() *float64
	// Experimental.
	SetNumCacheNodes(val *float64)
	// Experimental.
	NumCacheNodesInput() *float64
	// Experimental.
	OutpostMode() *string
	// Experimental.
	SetOutpostMode(val *string)
	// Experimental.
	OutpostModeInput() *string
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
	PreferredAvailabilityZones() *[]*string
	// Experimental.
	SetPreferredAvailabilityZones(val *[]*string)
	// Experimental.
	PreferredAvailabilityZonesInput() *[]*string
	// Experimental.
	PreferredOutpostArn() *string
	// Experimental.
	SetPreferredOutpostArn(val *string)
	// Experimental.
	PreferredOutpostArnInput() *string
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
	Timeouts() AwsElasticacheCluster_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TransitEncryptionEnabled() interface{}
	// Experimental.
	SetTransitEncryptionEnabled(val interface{})
	// Experimental.
	TransitEncryptionEnabledInput() interface{}
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
	PutTimeouts(value *AwsElasticacheCluster_TimeoutsProperty)
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
	ResetAutoMinorVersionUpgrade()
	// Experimental.
	ResetAvailabilityZone()
	// Experimental.
	ResetAzMode()
	// Experimental.
	ResetEngine()
	// Experimental.
	ResetEngineVersion()
	// Experimental.
	ResetFinalSnapshotIdentifier()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIpDiscovery()
	// Experimental.
	ResetLogDeliveryConfiguration()
	// Experimental.
	ResetMaintenanceWindow()
	// Experimental.
	ResetNetworkType()
	// Experimental.
	ResetNodeType()
	// Experimental.
	ResetNotificationTopicArn()
	// Experimental.
	ResetNumCacheNodes()
	// Experimental.
	ResetOutpostMode()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetParameterGroupName()
	// Experimental.
	ResetPort()
	// Experimental.
	ResetPreferredAvailabilityZones()
	// Experimental.
	ResetPreferredOutpostArn()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetReplicationGroupId()
	// Experimental.
	ResetSecurityGroupIds()
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

// The jsii proxy struct for AwsElasticacheCluster
type jsiiProxy_AwsElasticacheCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsElasticacheCluster) ApplyImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) ApplyImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) AutoMinorVersionUpgrade() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) AutoMinorVersionUpgradeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) AzMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) AzModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) CacheNodes() AwsElasticacheCluster_CacheNodesPropertyList {
	var returns AwsElasticacheCluster_CacheNodesPropertyList
	_jsii_.Get(
		j,
		"cacheNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) ClusterAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) ConfigurationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) EngineVersionActual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionActual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) FinalSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) FinalSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) IpDiscovery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipDiscovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) IpDiscoveryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipDiscoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) LogDeliveryConfiguration() AwsElasticacheCluster_LogDeliveryConfigurationPropertyList {
	var returns AwsElasticacheCluster_LogDeliveryConfigurationPropertyList
	_jsii_.Get(
		j,
		"logDeliveryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) LogDeliveryConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDeliveryConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) MaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) MaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) NodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) NodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) NotificationTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) NotificationTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) NumCacheNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCacheNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) NumCacheNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCacheNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) OutpostMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) OutpostModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) ParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) ParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) PreferredAvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredAvailabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) PreferredAvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredAvailabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) PreferredOutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredOutpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) PreferredOutpostArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredOutpostArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) ReplicationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) ReplicationGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) SnapshotArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snapshotArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) SnapshotArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snapshotArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) SnapshotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) SnapshotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) SnapshotRetentionLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) SnapshotRetentionLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) SnapshotWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) SnapshotWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) SubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) SubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) Timeouts() AwsElasticacheCluster_TimeoutsPropertyOutputReference {
	var returns AwsElasticacheCluster_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) TransitEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsElasticacheCluster) TransitEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitEncryptionEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_cluster aws_elasticache_cluster} Resource.
// Experimental.
func NewAwsElasticacheCluster(scope constructs.Construct, id *string, config *AwsElasticacheClusterConfig) AwsElasticacheCluster {
	_init_.Initialize()

	if err := validateNewAwsElasticacheClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsElasticacheCluster{}

	_jsii_.Create(
		"@cdktn/aws-elasticache.AwsElasticacheCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_cluster aws_elasticache_cluster} Resource.
// Experimental.
func NewAwsElasticacheCluster_Override(a AwsElasticacheCluster, scope constructs.Construct, id *string, config *AwsElasticacheClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elasticache.AwsElasticacheCluster",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetApplyImmediately(val interface{}) {
	if err := j.validateSetApplyImmediatelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyImmediately",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetAutoMinorVersionUpgrade(val *string) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetAzMode(val *string) {
	if err := j.validateSetAzModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azMode",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetFinalSnapshotIdentifier(val *string) {
	if err := j.validateSetFinalSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetIpDiscovery(val *string) {
	if err := j.validateSetIpDiscoveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipDiscovery",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetMaintenanceWindow(val *string) {
	if err := j.validateSetMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetNodeType(val *string) {
	if err := j.validateSetNodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeType",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetNotificationTopicArn(val *string) {
	if err := j.validateSetNotificationTopicArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationTopicArn",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetNumCacheNodes(val *float64) {
	if err := j.validateSetNumCacheNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numCacheNodes",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetOutpostMode(val *string) {
	if err := j.validateSetOutpostModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outpostMode",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetParameterGroupName(val *string) {
	if err := j.validateSetParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetPreferredAvailabilityZones(val *[]*string) {
	if err := j.validateSetPreferredAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredAvailabilityZones",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetPreferredOutpostArn(val *string) {
	if err := j.validateSetPreferredOutpostArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredOutpostArn",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetReplicationGroupId(val *string) {
	if err := j.validateSetReplicationGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationGroupId",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetSnapshotArns(val *[]*string) {
	if err := j.validateSetSnapshotArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotArns",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetSnapshotName(val *string) {
	if err := j.validateSetSnapshotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotName",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetSnapshotRetentionLimit(val *float64) {
	if err := j.validateSetSnapshotRetentionLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotRetentionLimit",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetSnapshotWindow(val *string) {
	if err := j.validateSetSnapshotWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotWindow",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetSubnetGroupName(val *string) {
	if err := j.validateSetSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsElasticacheCluster)SetTransitEncryptionEnabled(val interface{}) {
	if err := j.validateSetTransitEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitEncryptionEnabled",
		val,
	)
}

// Generates CDKTN code for importing a AwsElasticacheCluster resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsElasticacheCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsElasticacheCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-elasticache.AwsElasticacheCluster",
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
func AwsElasticacheCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsElasticacheCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elasticache.AwsElasticacheCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsElasticacheCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsElasticacheCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elasticache.AwsElasticacheCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsElasticacheCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsElasticacheCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-elasticache.AwsElasticacheCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsElasticacheCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-elasticache.AwsElasticacheCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsElasticacheCluster) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsElasticacheCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsElasticacheCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsElasticacheCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsElasticacheCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsElasticacheCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsElasticacheCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsElasticacheCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsElasticacheCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsElasticacheCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticacheCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsElasticacheCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsElasticacheCluster) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) PutLogDeliveryConfiguration(value interface{}) {
	if err := a.validatePutLogDeliveryConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogDeliveryConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) PutTimeouts(value *AwsElasticacheCluster_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetApplyImmediately() {
	_jsii_.InvokeVoid(
		a,
		"resetApplyImmediately",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetAzMode() {
	_jsii_.InvokeVoid(
		a,
		"resetAzMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetEngine() {
	_jsii_.InvokeVoid(
		a,
		"resetEngine",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetFinalSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		a,
		"resetFinalSnapshotIdentifier",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetIpDiscovery() {
	_jsii_.InvokeVoid(
		a,
		"resetIpDiscovery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetLogDeliveryConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLogDeliveryConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetNetworkType() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetNodeType() {
	_jsii_.InvokeVoid(
		a,
		"resetNodeType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetNotificationTopicArn() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationTopicArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetNumCacheNodes() {
	_jsii_.InvokeVoid(
		a,
		"resetNumCacheNodes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetOutpostMode() {
	_jsii_.InvokeVoid(
		a,
		"resetOutpostMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetParameterGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetParameterGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetPreferredAvailabilityZones() {
	_jsii_.InvokeVoid(
		a,
		"resetPreferredAvailabilityZones",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetPreferredOutpostArn() {
	_jsii_.InvokeVoid(
		a,
		"resetPreferredOutpostArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetReplicationGroupId() {
	_jsii_.InvokeVoid(
		a,
		"resetReplicationGroupId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetSnapshotArns() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetSnapshotName() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetSnapshotRetentionLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotRetentionLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetSnapshotWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetSubnetGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) ResetTransitEncryptionEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTransitEncryptionEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsElasticacheCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticacheCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticacheCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticacheCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticacheCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticacheCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsElasticacheCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

