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
type TfCluster interface {
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
	BrokerNodeGroupInfo() TfCluster_BrokerNodeGroupInfoPropertyOutputReference
	// Experimental.
	BrokerNodeGroupInfoInput() *TfCluster_BrokerNodeGroupInfoProperty
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClientAuthentication() TfCluster_ClientAuthenticationPropertyOutputReference
	// Experimental.
	ClientAuthenticationInput() *TfCluster_ClientAuthenticationProperty
	// Experimental.
	ClusterName() *string
	// Experimental.
	SetClusterName(val *string)
	// Experimental.
	ClusterNameInput() *string
	// Experimental.
	ClusterUuid() *string
	// Experimental.
	ConfigurationInfo() TfCluster_ConfigurationInfoPropertyOutputReference
	// Experimental.
	ConfigurationInfoInput() *TfCluster_ConfigurationInfoProperty
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
	EncryptionInfo() TfCluster_EncryptionInfoPropertyOutputReference
	// Experimental.
	EncryptionInfoInput() *TfCluster_EncryptionInfoProperty
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
	LoggingInfo() TfCluster_LoggingInfoPropertyOutputReference
	// Experimental.
	LoggingInfoInput() *TfCluster_LoggingInfoProperty
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
	OpenMonitoring() TfCluster_OpenMonitoringPropertyOutputReference
	// Experimental.
	OpenMonitoringInput() *TfCluster_OpenMonitoringProperty
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
	Rebalancing() TfCluster_RebalancingPropertyOutputReference
	// Experimental.
	RebalancingInput() *TfCluster_RebalancingProperty
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
	Timeouts() TfCluster_TimeoutsPropertyOutputReference
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
	PutBrokerNodeGroupInfo(value *TfCluster_BrokerNodeGroupInfoProperty)
	// Experimental.
	PutClientAuthentication(value *TfCluster_ClientAuthenticationProperty)
	// Experimental.
	PutConfigurationInfo(value *TfCluster_ConfigurationInfoProperty)
	// Experimental.
	PutEncryptionInfo(value *TfCluster_EncryptionInfoProperty)
	// Experimental.
	PutLoggingInfo(value *TfCluster_LoggingInfoProperty)
	// Experimental.
	PutOpenMonitoring(value *TfCluster_OpenMonitoringProperty)
	// Experimental.
	PutRebalancing(value *TfCluster_RebalancingProperty)
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

// The jsii proxy struct for TfCluster
type jsiiProxy_TfCluster struct {
	internal.Type__cdktnTerraformResource
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

func (j *jsiiProxy_TfCluster) BootstrapBrokers() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapBrokersIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapBrokersPublicSaslIam() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersPublicSaslIam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapBrokersPublicSaslScram() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersPublicSaslScram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapBrokersPublicTls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersPublicTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapBrokersSaslIam() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersSaslIam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapBrokersSaslIamIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersSaslIamIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapBrokersSaslScram() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersSaslScram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapBrokersSaslScramIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersSaslScramIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapBrokersTls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapBrokersTlsIpv6() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersTlsIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapBrokersVpcConnectivitySaslIam() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersVpcConnectivitySaslIam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapBrokersVpcConnectivitySaslScram() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersVpcConnectivitySaslScram",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BootstrapBrokersVpcConnectivityTls() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapBrokersVpcConnectivityTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BrokerNodeGroupInfo() TfCluster_BrokerNodeGroupInfoPropertyOutputReference {
	var returns TfCluster_BrokerNodeGroupInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"brokerNodeGroupInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) BrokerNodeGroupInfoInput() *TfCluster_BrokerNodeGroupInfoProperty {
	var returns *TfCluster_BrokerNodeGroupInfoProperty
	_jsii_.Get(
		j,
		"brokerNodeGroupInfoInput",
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

func (j *jsiiProxy_TfCluster) ClientAuthentication() TfCluster_ClientAuthenticationPropertyOutputReference {
	var returns TfCluster_ClientAuthenticationPropertyOutputReference
	_jsii_.Get(
		j,
		"clientAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ClientAuthenticationInput() *TfCluster_ClientAuthenticationProperty {
	var returns *TfCluster_ClientAuthenticationProperty
	_jsii_.Get(
		j,
		"clientAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ClusterUuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterUuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ConfigurationInfo() TfCluster_ConfigurationInfoPropertyOutputReference {
	var returns TfCluster_ConfigurationInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"configurationInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ConfigurationInfoInput() *TfCluster_ConfigurationInfoProperty {
	var returns *TfCluster_ConfigurationInfoProperty
	_jsii_.Get(
		j,
		"configurationInfoInput",
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

func (j *jsiiProxy_TfCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) CurrentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) CustomerActionStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerActionStatus",
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

func (j *jsiiProxy_TfCluster) EncryptionInfo() TfCluster_EncryptionInfoPropertyOutputReference {
	var returns TfCluster_EncryptionInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"encryptionInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EncryptionInfoInput() *TfCluster_EncryptionInfoProperty {
	var returns *TfCluster_EncryptionInfoProperty
	_jsii_.Get(
		j,
		"encryptionInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EnhancedMonitoring() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enhancedMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) EnhancedMonitoringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enhancedMonitoringInput",
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

func (j *jsiiProxy_TfCluster) KafkaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) KafkaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kafkaVersionInput",
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

func (j *jsiiProxy_TfCluster) LoggingInfo() TfCluster_LoggingInfoPropertyOutputReference {
	var returns TfCluster_LoggingInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"loggingInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) LoggingInfoInput() *TfCluster_LoggingInfoProperty {
	var returns *TfCluster_LoggingInfoProperty
	_jsii_.Get(
		j,
		"loggingInfoInput",
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

func (j *jsiiProxy_TfCluster) NumberOfBrokerNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBrokerNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) NumberOfBrokerNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfBrokerNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) OpenMonitoring() TfCluster_OpenMonitoringPropertyOutputReference {
	var returns TfCluster_OpenMonitoringPropertyOutputReference
	_jsii_.Get(
		j,
		"openMonitoring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) OpenMonitoringInput() *TfCluster_OpenMonitoringProperty {
	var returns *TfCluster_OpenMonitoringProperty
	_jsii_.Get(
		j,
		"openMonitoringInput",
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

func (j *jsiiProxy_TfCluster) Rebalancing() TfCluster_RebalancingPropertyOutputReference {
	var returns TfCluster_RebalancingPropertyOutputReference
	_jsii_.Get(
		j,
		"rebalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) RebalancingInput() *TfCluster_RebalancingProperty {
	var returns *TfCluster_RebalancingProperty
	_jsii_.Get(
		j,
		"rebalancingInput",
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

func (j *jsiiProxy_TfCluster) StorageMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) StorageModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageModeInput",
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

func (j *jsiiProxy_TfCluster) ZookeeperConnectString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zookeeperConnectString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster) ZookeeperConnectStringTls() *string {
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
func NewTfCluster(scope constructs.Construct, id *string, config *TfClusterConfig) TfCluster {
	_init_.Initialize()

	if err := validateNewTfClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCluster{}

	_jsii_.Create(
		"@cdktn/aws-msk.TfCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster aws_msk_cluster} Resource.
// Experimental.
func NewTfCluster_Override(t TfCluster, scope constructs.Construct, id *string, config *TfClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-msk.TfCluster",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfCluster)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
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

func (j *jsiiProxy_TfCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfCluster)SetEnhancedMonitoring(val *string) {
	if err := j.validateSetEnhancedMonitoringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enhancedMonitoring",
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

func (j *jsiiProxy_TfCluster)SetKafkaVersion(val *string) {
	if err := j.validateSetKafkaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kafkaVersion",
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

func (j *jsiiProxy_TfCluster)SetNumberOfBrokerNodes(val *float64) {
	if err := j.validateSetNumberOfBrokerNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfBrokerNodes",
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

func (j *jsiiProxy_TfCluster)SetStorageMode(val *string) {
	if err := j.validateSetStorageModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageMode",
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

// Generates CDKTN code for importing a TfCluster resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-msk.TfCluster",
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
		"@cdktn/aws-msk.TfCluster",
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
		"@cdktn/aws-msk.TfCluster",
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
		"@cdktn/aws-msk.TfCluster",
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
		"@cdktn/aws-msk.TfCluster",
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

func (t *jsiiProxy_TfCluster) PutBrokerNodeGroupInfo(value *TfCluster_BrokerNodeGroupInfoProperty) {
	if err := t.validatePutBrokerNodeGroupInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBrokerNodeGroupInfo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutClientAuthentication(value *TfCluster_ClientAuthenticationProperty) {
	if err := t.validatePutClientAuthenticationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putClientAuthentication",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutConfigurationInfo(value *TfCluster_ConfigurationInfoProperty) {
	if err := t.validatePutConfigurationInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConfigurationInfo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutEncryptionInfo(value *TfCluster_EncryptionInfoProperty) {
	if err := t.validatePutEncryptionInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEncryptionInfo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutLoggingInfo(value *TfCluster_LoggingInfoProperty) {
	if err := t.validatePutLoggingInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLoggingInfo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutOpenMonitoring(value *TfCluster_OpenMonitoringProperty) {
	if err := t.validatePutOpenMonitoringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOpenMonitoring",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster) PutRebalancing(value *TfCluster_RebalancingProperty) {
	if err := t.validatePutRebalancingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRebalancing",
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

func (t *jsiiProxy_TfCluster) ResetClientAuthentication() {
	_jsii_.InvokeVoid(
		t,
		"resetClientAuthentication",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetConfigurationInfo() {
	_jsii_.InvokeVoid(
		t,
		"resetConfigurationInfo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetEncryptionInfo() {
	_jsii_.InvokeVoid(
		t,
		"resetEncryptionInfo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetEnhancedMonitoring() {
	_jsii_.InvokeVoid(
		t,
		"resetEnhancedMonitoring",
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

func (t *jsiiProxy_TfCluster) ResetLoggingInfo() {
	_jsii_.InvokeVoid(
		t,
		"resetLoggingInfo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster) ResetOpenMonitoring() {
	_jsii_.InvokeVoid(
		t,
		"resetOpenMonitoring",
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

func (t *jsiiProxy_TfCluster) ResetRebalancing() {
	_jsii_.InvokeVoid(
		t,
		"resetRebalancing",
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

func (t *jsiiProxy_TfCluster) ResetStorageMode() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageMode",
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

