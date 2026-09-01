package awsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmsk/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsmsk/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster aws_msk_cluster}.
// Experimental.
type AwsMskCluster interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	BootstrapBrokers() *string
	// Experimental.
	BootstrapBrokersIpv6() *string
	// Experimental.
	BootstrapBrokersPublicSaslIam() *string
	// Experimental.
	BootstrapBrokersPublicSaslScram() *string
	// Experimental.
	BootstrapBrokersPublicTls() *string
	// Experimental.
	BootstrapBrokersSaslIam() *string
	// Experimental.
	BootstrapBrokersSaslIamIpv6() *string
	// Experimental.
	BootstrapBrokersSaslScram() *string
	// Experimental.
	BootstrapBrokersSaslScramIpv6() *string
	// Experimental.
	BootstrapBrokersTls() *string
	// Experimental.
	BootstrapBrokersTlsIpv6() *string
	// Experimental.
	BootstrapBrokersVpcConnectivitySaslIam() *string
	// Experimental.
	BootstrapBrokersVpcConnectivitySaslScram() *string
	// Experimental.
	BootstrapBrokersVpcConnectivityTls() *string
	// Experimental.
	BrokerNodeGroupInfo() AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference
	// Experimental.
	BrokerNodeGroupInfoInput() *AwsMskCluster_BrokerNodeGroupInfoProperty
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClientAuthentication() AwsMskCluster_ClientAuthenticationPropertyOutputReference
	// Experimental.
	ClientAuthenticationInput() *AwsMskCluster_ClientAuthenticationProperty
	// Experimental.
	ClusterName() *string
	// Experimental.
	SetClusterName(val *string)
	// Experimental.
	ClusterNameInput() *string
	// Experimental.
	ClusterUuid() *string
	// Experimental.
	ConfigurationInfo() AwsMskCluster_ConfigurationInfoPropertyOutputReference
	// Experimental.
	ConfigurationInfoInput() *AwsMskCluster_ConfigurationInfoProperty
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
	CurrentVersion() *string
	// Experimental.
	CustomerActionStatus() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EncryptionInfo() AwsMskCluster_EncryptionInfoPropertyOutputReference
	// Experimental.
	EncryptionInfoInput() *AwsMskCluster_EncryptionInfoProperty
	// Experimental.
	EnhancedMonitoring() *string
	// Experimental.
	SetEnhancedMonitoring(val *string)
	// Experimental.
	EnhancedMonitoringInput() *string
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
	KafkaVersion() *string
	// Experimental.
	SetKafkaVersion(val *string)
	// Experimental.
	KafkaVersionInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LoggingInfo() AwsMskCluster_LoggingInfoPropertyOutputReference
	// Experimental.
	LoggingInfoInput() *AwsMskCluster_LoggingInfoProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	NumberOfBrokerNodes() *float64
	// Experimental.
	SetNumberOfBrokerNodes(val *float64)
	// Experimental.
	NumberOfBrokerNodesInput() *float64
	// Experimental.
	OpenMonitoring() AwsMskCluster_OpenMonitoringPropertyOutputReference
	// Experimental.
	OpenMonitoringInput() *AwsMskCluster_OpenMonitoringProperty
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
	Rebalancing() AwsMskCluster_RebalancingPropertyOutputReference
	// Experimental.
	RebalancingInput() *AwsMskCluster_RebalancingProperty
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	StorageMode() *string
	// Experimental.
	SetStorageMode(val *string)
	// Experimental.
	StorageModeInput() *string
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
	Timeouts() AwsMskCluster_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	ZookeeperConnectString() *string
	// Experimental.
	ZookeeperConnectStringTls() *string
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
	PutBrokerNodeGroupInfo(value *AwsMskCluster_BrokerNodeGroupInfoProperty)
	// Experimental.
	PutClientAuthentication(value *AwsMskCluster_ClientAuthenticationProperty)
	// Experimental.
	PutConfigurationInfo(value *AwsMskCluster_ConfigurationInfoProperty)
	// Experimental.
	PutEncryptionInfo(value *AwsMskCluster_EncryptionInfoProperty)
	// Experimental.
	PutLoggingInfo(value *AwsMskCluster_LoggingInfoProperty)
	// Experimental.
	PutOpenMonitoring(value *AwsMskCluster_OpenMonitoringProperty)
	// Experimental.
	PutRebalancing(value *AwsMskCluster_RebalancingProperty)
	// Experimental.
	PutTimeouts(value *AwsMskCluster_TimeoutsProperty)
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
	ResetClientAuthentication()
	// Experimental.
	ResetConfigurationInfo()
	// Experimental.
	ResetEncryptionInfo()
	// Experimental.
	ResetEnhancedMonitoring()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLoggingInfo()
	// Experimental.
	ResetOpenMonitoring()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRebalancing()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetStorageMode()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
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

// The jsii proxy struct for AwsMskCluster
type jsiiProxy_AwsMskCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsMskCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) BootstrapBrokers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) BootstrapBrokersIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) BootstrapBrokersPublicSaslIam() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersPublicSaslIam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) BootstrapBrokersPublicSaslScram() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersPublicSaslScram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) BootstrapBrokersPublicTls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersPublicTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) BootstrapBrokersSaslIam() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersSaslIam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) BootstrapBrokersSaslIamIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersSaslIamIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) BootstrapBrokersSaslScram() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersSaslScram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) BootstrapBrokersSaslScramIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersSaslScramIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) BootstrapBrokersTls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) BootstrapBrokersTlsIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersTlsIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) BootstrapBrokersVpcConnectivitySaslIam() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersVpcConnectivitySaslIam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) BootstrapBrokersVpcConnectivitySaslScram() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersVpcConnectivitySaslScram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) BootstrapBrokersVpcConnectivityTls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersVpcConnectivityTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) BrokerNodeGroupInfo() AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference {
	var returns AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"brokerNodeGroupInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) BrokerNodeGroupInfoInput() *AwsMskCluster_BrokerNodeGroupInfoProperty {
	var returns *AwsMskCluster_BrokerNodeGroupInfoProperty
	_jsii_.Get(
		j,
		"brokerNodeGroupInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) ClientAuthentication() AwsMskCluster_ClientAuthenticationPropertyOutputReference {
	var returns AwsMskCluster_ClientAuthenticationPropertyOutputReference
	_jsii_.Get(
		j,
		"clientAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) ClientAuthenticationInput() *AwsMskCluster_ClientAuthenticationProperty {
	var returns *AwsMskCluster_ClientAuthenticationProperty
	_jsii_.Get(
		j,
		"clientAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) ClusterUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) ConfigurationInfo() AwsMskCluster_ConfigurationInfoPropertyOutputReference {
	var returns AwsMskCluster_ConfigurationInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"configurationInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) ConfigurationInfoInput() *AwsMskCluster_ConfigurationInfoProperty {
	var returns *AwsMskCluster_ConfigurationInfoProperty
	_jsii_.Get(
		j,
		"configurationInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) CurrentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) CustomerActionStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerActionStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) EncryptionInfo() AwsMskCluster_EncryptionInfoPropertyOutputReference {
	var returns AwsMskCluster_EncryptionInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"encryptionInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) EncryptionInfoInput() *AwsMskCluster_EncryptionInfoProperty {
	var returns *AwsMskCluster_EncryptionInfoProperty
	_jsii_.Get(
		j,
		"encryptionInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) EnhancedMonitoring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enhancedMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) EnhancedMonitoringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enhancedMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) KafkaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) KafkaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) LoggingInfo() AwsMskCluster_LoggingInfoPropertyOutputReference {
	var returns AwsMskCluster_LoggingInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"loggingInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) LoggingInfoInput() *AwsMskCluster_LoggingInfoProperty {
	var returns *AwsMskCluster_LoggingInfoProperty
	_jsii_.Get(
		j,
		"loggingInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) NumberOfBrokerNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBrokerNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) NumberOfBrokerNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBrokerNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) OpenMonitoring() AwsMskCluster_OpenMonitoringPropertyOutputReference {
	var returns AwsMskCluster_OpenMonitoringPropertyOutputReference
	_jsii_.Get(
		j,
		"openMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) OpenMonitoringInput() *AwsMskCluster_OpenMonitoringProperty {
	var returns *AwsMskCluster_OpenMonitoringProperty
	_jsii_.Get(
		j,
		"openMonitoringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) Rebalancing() AwsMskCluster_RebalancingPropertyOutputReference {
	var returns AwsMskCluster_RebalancingPropertyOutputReference
	_jsii_.Get(
		j,
		"rebalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) RebalancingInput() *AwsMskCluster_RebalancingProperty {
	var returns *AwsMskCluster_RebalancingProperty
	_jsii_.Get(
		j,
		"rebalancingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) StorageMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) StorageModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) Timeouts() AwsMskCluster_TimeoutsPropertyOutputReference {
	var returns AwsMskCluster_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) ZookeeperConnectString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zookeeperConnectString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster) ZookeeperConnectStringTls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zookeeperConnectStringTls",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster aws_msk_cluster} Resource.
// Experimental.
func NewAwsMskCluster(scope constructs.Construct, id *string, config *AwsMskClusterConfig) AwsMskCluster {
	_init_.Initialize()

	if err := validateNewAwsMskClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMskCluster{}

	_jsii_.Create(
		"@cdktn/aws-msk.AwsMskCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster aws_msk_cluster} Resource.
// Experimental.
func NewAwsMskCluster_Override(a AwsMskCluster, scope constructs.Construct, id *string, config *AwsMskClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-msk.AwsMskCluster",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsMskCluster)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster)SetEnhancedMonitoring(val *string) {
	if err := j.validateSetEnhancedMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enhancedMonitoring",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster)SetKafkaVersion(val *string) {
	if err := j.validateSetKafkaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kafkaVersion",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster)SetNumberOfBrokerNodes(val *float64) {
	if err := j.validateSetNumberOfBrokerNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfBrokerNodes",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster)SetStorageMode(val *string) {
	if err := j.validateSetStorageModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageMode",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a AwsMskCluster resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsMskCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsMskCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-msk.AwsMskCluster",
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
func AwsMskCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsMskCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-msk.AwsMskCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsMskCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsMskCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-msk.AwsMskCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsMskCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsMskCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-msk.AwsMskCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsMskCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-msk.AwsMskCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsMskCluster) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsMskCluster) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsMskCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMskCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMskCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMskCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMskCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMskCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMskCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMskCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMskCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMskCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMskCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsMskCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMskCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsMskCluster) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsMskCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsMskCluster) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsMskCluster) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsMskCluster) PutBrokerNodeGroupInfo(value *AwsMskCluster_BrokerNodeGroupInfoProperty) {
	if err := a.validatePutBrokerNodeGroupInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBrokerNodeGroupInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMskCluster) PutClientAuthentication(value *AwsMskCluster_ClientAuthenticationProperty) {
	if err := a.validatePutClientAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putClientAuthentication",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMskCluster) PutConfigurationInfo(value *AwsMskCluster_ConfigurationInfoProperty) {
	if err := a.validatePutConfigurationInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConfigurationInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMskCluster) PutEncryptionInfo(value *AwsMskCluster_EncryptionInfoProperty) {
	if err := a.validatePutEncryptionInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEncryptionInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMskCluster) PutLoggingInfo(value *AwsMskCluster_LoggingInfoProperty) {
	if err := a.validatePutLoggingInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLoggingInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMskCluster) PutOpenMonitoring(value *AwsMskCluster_OpenMonitoringProperty) {
	if err := a.validatePutOpenMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpenMonitoring",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMskCluster) PutRebalancing(value *AwsMskCluster_RebalancingProperty) {
	if err := a.validatePutRebalancingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRebalancing",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMskCluster) PutTimeouts(value *AwsMskCluster_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMskCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsMskCluster) ResetClientAuthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetClientAuthentication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster) ResetConfigurationInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetConfigurationInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster) ResetEncryptionInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster) ResetEnhancedMonitoring() {
	_jsii_.InvokeVoid(
		a,
		"resetEnhancedMonitoring",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster) ResetLoggingInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetLoggingInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster) ResetOpenMonitoring() {
	_jsii_.InvokeVoid(
		a,
		"resetOpenMonitoring",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster) ResetRebalancing() {
	_jsii_.InvokeVoid(
		a,
		"resetRebalancing",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster) ResetStorageMode() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMskCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMskCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMskCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMskCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMskCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMskCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

