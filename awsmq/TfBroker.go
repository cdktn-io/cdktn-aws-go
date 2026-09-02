package awsmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmq/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsmq/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker aws_mq_broker}.
// Experimental.
type TfBroker interface {
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
	AuthenticationStrategy() *string
	// Experimental.
	SetAuthenticationStrategy(val *string)
	// Experimental.
	AuthenticationStrategyInput() *string
	// Experimental.
	AutoMinorVersionUpgrade() interface{}
	// Experimental.
	SetAutoMinorVersionUpgrade(val interface{})
	// Experimental.
	AutoMinorVersionUpgradeInput() interface{}
	// Experimental.
	BrokerName() *string
	// Experimental.
	SetBrokerName(val *string)
	// Experimental.
	BrokerNameInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Configuration() TfBroker_ConfigurationPropertyOutputReference
	// Experimental.
	ConfigurationInput() *TfBroker_ConfigurationProperty
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
	DataReplicationMode() *string
	// Experimental.
	SetDataReplicationMode(val *string)
	// Experimental.
	DataReplicationModeInput() *string
	// Experimental.
	DataReplicationPrimaryBrokerArn() *string
	// Experimental.
	SetDataReplicationPrimaryBrokerArn(val *string)
	// Experimental.
	DataReplicationPrimaryBrokerArnInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DeploymentMode() *string
	// Experimental.
	SetDeploymentMode(val *string)
	// Experimental.
	DeploymentModeInput() *string
	// Experimental.
	EncryptionOptions() TfBroker_EncryptionOptionsPropertyOutputReference
	// Experimental.
	EncryptionOptionsInput() *TfBroker_EncryptionOptionsProperty
	// Experimental.
	EngineType() *string
	// Experimental.
	SetEngineType(val *string)
	// Experimental.
	EngineTypeInput() *string
	// Experimental.
	EngineVersion() *string
	// Experimental.
	SetEngineVersion(val *string)
	// Experimental.
	EngineVersionInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HostInstanceType() *string
	// Experimental.
	SetHostInstanceType(val *string)
	// Experimental.
	HostInstanceTypeInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	Instances() TfBroker_InstancesPropertyList
	// Experimental.
	LdapServerMetadata() TfBroker_LdapServerMetadataPropertyOutputReference
	// Experimental.
	LdapServerMetadataInput() *TfBroker_LdapServerMetadataProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	Logs() TfBroker_LogsPropertyOutputReference
	// Experimental.
	LogsInput() *TfBroker_LogsProperty
	// Experimental.
	MaintenanceWindowStartTime() TfBroker_MaintenanceWindowStartTimePropertyOutputReference
	// Experimental.
	MaintenanceWindowStartTimeInput() *TfBroker_MaintenanceWindowStartTimeProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PendingDataReplicationMode() *string
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
	ResourceShareArns() *[]*string
	// Experimental.
	SetResourceShareArns(val *[]*string)
	// Experimental.
	ResourceShareArnsInput() *[]*string
	// Experimental.
	SecurityGroups() *[]*string
	// Experimental.
	SetSecurityGroups(val *[]*string)
	// Experimental.
	SecurityGroupsInput() *[]*string
	// Experimental.
	SharedResources() TfBroker_SharedResourcesPropertyList
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
	Timeouts() TfBroker_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	User() TfBroker_UserPropertyList
	// Experimental.
	UserInput() interface{}
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
	PutConfiguration(value *TfBroker_ConfigurationProperty)
	// Experimental.
	PutEncryptionOptions(value *TfBroker_EncryptionOptionsProperty)
	// Experimental.
	PutLdapServerMetadata(value *TfBroker_LdapServerMetadataProperty)
	// Experimental.
	PutLogs(value *TfBroker_LogsProperty)
	// Experimental.
	PutMaintenanceWindowStartTime(value *TfBroker_MaintenanceWindowStartTimeProperty)
	// Experimental.
	PutTimeouts(value *TfBroker_TimeoutsProperty)
	// Experimental.
	PutUser(value interface{})
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
	ResetAuthenticationStrategy()
	// Experimental.
	ResetAutoMinorVersionUpgrade()
	// Experimental.
	ResetConfiguration()
	// Experimental.
	ResetDataReplicationMode()
	// Experimental.
	ResetDataReplicationPrimaryBrokerArn()
	// Experimental.
	ResetDeploymentMode()
	// Experimental.
	ResetEncryptionOptions()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLdapServerMetadata()
	// Experimental.
	ResetLogs()
	// Experimental.
	ResetMaintenanceWindowStartTime()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPubliclyAccessible()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetResourceShareArns()
	// Experimental.
	ResetSecurityGroups()
	// Experimental.
	ResetStorageType()
	// Experimental.
	ResetSubnetIds()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetUser()
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

// The jsii proxy struct for TfBroker
type jsiiProxy_TfBroker struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfBroker) ApplyImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) ApplyImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) AuthenticationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) AuthenticationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) AutoMinorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) BrokerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brokerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) BrokerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brokerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) Configuration() TfBroker_ConfigurationPropertyOutputReference {
	var returns TfBroker_ConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) ConfigurationInput() *TfBroker_ConfigurationProperty {
	var returns *TfBroker_ConfigurationProperty
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) DataReplicationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataReplicationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) DataReplicationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataReplicationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) DataReplicationPrimaryBrokerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataReplicationPrimaryBrokerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) DataReplicationPrimaryBrokerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataReplicationPrimaryBrokerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) DeploymentMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) DeploymentModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) EncryptionOptions() TfBroker_EncryptionOptionsPropertyOutputReference {
	var returns TfBroker_EncryptionOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"encryptionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) EncryptionOptionsInput() *TfBroker_EncryptionOptionsProperty {
	var returns *TfBroker_EncryptionOptionsProperty
	_jsii_.Get(
		j,
		"encryptionOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) EngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) EngineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) HostInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) HostInstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) Instances() TfBroker_InstancesPropertyList {
	var returns TfBroker_InstancesPropertyList
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) LdapServerMetadata() TfBroker_LdapServerMetadataPropertyOutputReference {
	var returns TfBroker_LdapServerMetadataPropertyOutputReference
	_jsii_.Get(
		j,
		"ldapServerMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) LdapServerMetadataInput() *TfBroker_LdapServerMetadataProperty {
	var returns *TfBroker_LdapServerMetadataProperty
	_jsii_.Get(
		j,
		"ldapServerMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) Logs() TfBroker_LogsPropertyOutputReference {
	var returns TfBroker_LogsPropertyOutputReference
	_jsii_.Get(
		j,
		"logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) LogsInput() *TfBroker_LogsProperty {
	var returns *TfBroker_LogsProperty
	_jsii_.Get(
		j,
		"logsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) MaintenanceWindowStartTime() TfBroker_MaintenanceWindowStartTimePropertyOutputReference {
	var returns TfBroker_MaintenanceWindowStartTimePropertyOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindowStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) MaintenanceWindowStartTimeInput() *TfBroker_MaintenanceWindowStartTimeProperty {
	var returns *TfBroker_MaintenanceWindowStartTimeProperty
	_jsii_.Get(
		j,
		"maintenanceWindowStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) PendingDataReplicationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pendingDataReplicationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) PubliclyAccessibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) ResourceShareArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceShareArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) ResourceShareArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceShareArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) SharedResources() TfBroker_SharedResourcesPropertyList {
	var returns TfBroker_SharedResourcesPropertyList
	_jsii_.Get(
		j,
		"sharedResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) Timeouts() TfBroker_TimeoutsPropertyOutputReference {
	var returns TfBroker_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) User() TfBroker_UserPropertyList {
	var returns TfBroker_UserPropertyList
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBroker) UserInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker aws_mq_broker} Resource.
// Experimental.
func NewTfBroker(scope constructs.Construct, id *string, config *TfBrokerConfig) TfBroker {
	_init_.Initialize()

	if err := validateNewTfBrokerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfBroker{}

	_jsii_.Create(
		"@cdktn/aws-mq.TfBroker",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker aws_mq_broker} Resource.
// Experimental.
func NewTfBroker_Override(t TfBroker, scope constructs.Construct, id *string, config *TfBrokerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-mq.TfBroker",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfBroker)SetApplyImmediately(val interface{}) {
	if err := j.validateSetApplyImmediatelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyImmediately",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetAuthenticationStrategy(val *string) {
	if err := j.validateSetAuthenticationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationStrategy",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetAutoMinorVersionUpgrade(val interface{}) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetBrokerName(val *string) {
	if err := j.validateSetBrokerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"brokerName",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetDataReplicationMode(val *string) {
	if err := j.validateSetDataReplicationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataReplicationMode",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetDataReplicationPrimaryBrokerArn(val *string) {
	if err := j.validateSetDataReplicationPrimaryBrokerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataReplicationPrimaryBrokerArn",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetDeploymentMode(val *string) {
	if err := j.validateSetDeploymentModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentMode",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetEngineType(val *string) {
	if err := j.validateSetEngineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineType",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetHostInstanceType(val *string) {
	if err := j.validateSetHostInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostInstanceType",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetResourceShareArns(val *[]*string) {
	if err := j.validateSetResourceShareArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceShareArns",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfBroker)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a TfBroker resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfBroker_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfBroker_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-mq.TfBroker",
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
func TfBroker_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfBroker_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-mq.TfBroker",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfBroker_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfBroker_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-mq.TfBroker",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfBroker_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfBroker_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-mq.TfBroker",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfBroker_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-mq.TfBroker",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfBroker) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfBroker) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfBroker) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfBroker) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBroker) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfBroker) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfBroker) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfBroker) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfBroker) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfBroker) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfBroker) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfBroker) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBroker) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfBroker) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBroker) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfBroker) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfBroker) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfBroker) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfBroker) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfBroker) PutConfiguration(value *TfBroker_ConfigurationProperty) {
	if err := t.validatePutConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBroker) PutEncryptionOptions(value *TfBroker_EncryptionOptionsProperty) {
	if err := t.validatePutEncryptionOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEncryptionOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBroker) PutLdapServerMetadata(value *TfBroker_LdapServerMetadataProperty) {
	if err := t.validatePutLdapServerMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLdapServerMetadata",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBroker) PutLogs(value *TfBroker_LogsProperty) {
	if err := t.validatePutLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLogs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBroker) PutMaintenanceWindowStartTime(value *TfBroker_MaintenanceWindowStartTimeProperty) {
	if err := t.validatePutMaintenanceWindowStartTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMaintenanceWindowStartTime",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBroker) PutTimeouts(value *TfBroker_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBroker) PutUser(value interface{}) {
	if err := t.validatePutUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUser",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBroker) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfBroker) ResetApplyImmediately() {
	_jsii_.InvokeVoid(
		t,
		"resetApplyImmediately",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetAuthenticationStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthenticationStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetDataReplicationMode() {
	_jsii_.InvokeVoid(
		t,
		"resetDataReplicationMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetDataReplicationPrimaryBrokerArn() {
	_jsii_.InvokeVoid(
		t,
		"resetDataReplicationPrimaryBrokerArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetDeploymentMode() {
	_jsii_.InvokeVoid(
		t,
		"resetDeploymentMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetEncryptionOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetEncryptionOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetLdapServerMetadata() {
	_jsii_.InvokeVoid(
		t,
		"resetLdapServerMetadata",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetLogs() {
	_jsii_.InvokeVoid(
		t,
		"resetLogs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetMaintenanceWindowStartTime() {
	_jsii_.InvokeVoid(
		t,
		"resetMaintenanceWindowStartTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetPubliclyAccessible() {
	_jsii_.InvokeVoid(
		t,
		"resetPubliclyAccessible",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetResourceShareArns() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceShareArns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetStorageType() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		t,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) ResetUser() {
	_jsii_.InvokeVoid(
		t,
		"resetUser",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBroker) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBroker) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBroker) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBroker) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBroker) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBroker) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBroker) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

