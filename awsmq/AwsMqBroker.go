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
type AwsMqBroker interface {
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
	Configuration() AwsMqBroker_ConfigurationPropertyOutputReference
	// Experimental.
	ConfigurationInput() *AwsMqBroker_ConfigurationProperty
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
	EncryptionOptions() AwsMqBroker_EncryptionOptionsPropertyOutputReference
	// Experimental.
	EncryptionOptionsInput() *AwsMqBroker_EncryptionOptionsProperty
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
	Instances() AwsMqBroker_InstancesPropertyList
	// Experimental.
	LdapServerMetadata() AwsMqBroker_LdapServerMetadataPropertyOutputReference
	// Experimental.
	LdapServerMetadataInput() *AwsMqBroker_LdapServerMetadataProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	Logs() AwsMqBroker_LogsPropertyOutputReference
	// Experimental.
	LogsInput() *AwsMqBroker_LogsProperty
	// Experimental.
	MaintenanceWindowStartTime() AwsMqBroker_MaintenanceWindowStartTimePropertyOutputReference
	// Experimental.
	MaintenanceWindowStartTimeInput() *AwsMqBroker_MaintenanceWindowStartTimeProperty
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
	SharedResources() AwsMqBroker_SharedResourcesPropertyList
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
	Timeouts() AwsMqBroker_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	User() AwsMqBroker_UserPropertyList
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
	PutConfiguration(value *AwsMqBroker_ConfigurationProperty)
	// Experimental.
	PutEncryptionOptions(value *AwsMqBroker_EncryptionOptionsProperty)
	// Experimental.
	PutLdapServerMetadata(value *AwsMqBroker_LdapServerMetadataProperty)
	// Experimental.
	PutLogs(value *AwsMqBroker_LogsProperty)
	// Experimental.
	PutMaintenanceWindowStartTime(value *AwsMqBroker_MaintenanceWindowStartTimeProperty)
	// Experimental.
	PutTimeouts(value *AwsMqBroker_TimeoutsProperty)
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

// The jsii proxy struct for AwsMqBroker
type jsiiProxy_AwsMqBroker struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsMqBroker) ApplyImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) ApplyImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) AuthenticationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) AuthenticationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) AutoMinorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) BrokerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brokerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) BrokerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"brokerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) Configuration() AwsMqBroker_ConfigurationPropertyOutputReference {
	var returns AwsMqBroker_ConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) ConfigurationInput() *AwsMqBroker_ConfigurationProperty {
	var returns *AwsMqBroker_ConfigurationProperty
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) DataReplicationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataReplicationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) DataReplicationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataReplicationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) DataReplicationPrimaryBrokerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataReplicationPrimaryBrokerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) DataReplicationPrimaryBrokerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataReplicationPrimaryBrokerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) DeploymentMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) DeploymentModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) EncryptionOptions() AwsMqBroker_EncryptionOptionsPropertyOutputReference {
	var returns AwsMqBroker_EncryptionOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"encryptionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) EncryptionOptionsInput() *AwsMqBroker_EncryptionOptionsProperty {
	var returns *AwsMqBroker_EncryptionOptionsProperty
	_jsii_.Get(
		j,
		"encryptionOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) EngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) EngineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) HostInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) HostInstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) Instances() AwsMqBroker_InstancesPropertyList {
	var returns AwsMqBroker_InstancesPropertyList
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) LdapServerMetadata() AwsMqBroker_LdapServerMetadataPropertyOutputReference {
	var returns AwsMqBroker_LdapServerMetadataPropertyOutputReference
	_jsii_.Get(
		j,
		"ldapServerMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) LdapServerMetadataInput() *AwsMqBroker_LdapServerMetadataProperty {
	var returns *AwsMqBroker_LdapServerMetadataProperty
	_jsii_.Get(
		j,
		"ldapServerMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) Logs() AwsMqBroker_LogsPropertyOutputReference {
	var returns AwsMqBroker_LogsPropertyOutputReference
	_jsii_.Get(
		j,
		"logs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) LogsInput() *AwsMqBroker_LogsProperty {
	var returns *AwsMqBroker_LogsProperty
	_jsii_.Get(
		j,
		"logsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) MaintenanceWindowStartTime() AwsMqBroker_MaintenanceWindowStartTimePropertyOutputReference {
	var returns AwsMqBroker_MaintenanceWindowStartTimePropertyOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindowStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) MaintenanceWindowStartTimeInput() *AwsMqBroker_MaintenanceWindowStartTimeProperty {
	var returns *AwsMqBroker_MaintenanceWindowStartTimeProperty
	_jsii_.Get(
		j,
		"maintenanceWindowStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) PendingDataReplicationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pendingDataReplicationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) PubliclyAccessibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) ResourceShareArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceShareArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) ResourceShareArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceShareArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) SharedResources() AwsMqBroker_SharedResourcesPropertyList {
	var returns AwsMqBroker_SharedResourcesPropertyList
	_jsii_.Get(
		j,
		"sharedResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) Timeouts() AwsMqBroker_TimeoutsPropertyOutputReference {
	var returns AwsMqBroker_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) User() AwsMqBroker_UserPropertyList {
	var returns AwsMqBroker_UserPropertyList
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker) UserInput() interface{} {
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
func NewAwsMqBroker(scope constructs.Construct, id *string, config *AwsMqBrokerConfig) AwsMqBroker {
	_init_.Initialize()

	if err := validateNewAwsMqBrokerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMqBroker{}

	_jsii_.Create(
		"@cdktn/aws-mq.AwsMqBroker",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mq_broker aws_mq_broker} Resource.
// Experimental.
func NewAwsMqBroker_Override(a AwsMqBroker, scope constructs.Construct, id *string, config *AwsMqBrokerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-mq.AwsMqBroker",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetApplyImmediately(val interface{}) {
	if err := j.validateSetApplyImmediatelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyImmediately",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetAuthenticationStrategy(val *string) {
	if err := j.validateSetAuthenticationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationStrategy",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetAutoMinorVersionUpgrade(val interface{}) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetBrokerName(val *string) {
	if err := j.validateSetBrokerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"brokerName",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetDataReplicationMode(val *string) {
	if err := j.validateSetDataReplicationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataReplicationMode",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetDataReplicationPrimaryBrokerArn(val *string) {
	if err := j.validateSetDataReplicationPrimaryBrokerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataReplicationPrimaryBrokerArn",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetDeploymentMode(val *string) {
	if err := j.validateSetDeploymentModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentMode",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetEngineType(val *string) {
	if err := j.validateSetEngineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineType",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetHostInstanceType(val *string) {
	if err := j.validateSetHostInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostInstanceType",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetResourceShareArns(val *[]*string) {
	if err := j.validateSetResourceShareArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceShareArns",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a AwsMqBroker resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsMqBroker_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsMqBroker_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-mq.AwsMqBroker",
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
func AwsMqBroker_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsMqBroker_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-mq.AwsMqBroker",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsMqBroker_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsMqBroker_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-mq.AwsMqBroker",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsMqBroker_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsMqBroker_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-mq.AwsMqBroker",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsMqBroker_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-mq.AwsMqBroker",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsMqBroker) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsMqBroker) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsMqBroker) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMqBroker) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMqBroker) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMqBroker) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMqBroker) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMqBroker) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMqBroker) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMqBroker) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMqBroker) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMqBroker) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMqBroker) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsMqBroker) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMqBroker) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsMqBroker) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsMqBroker) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsMqBroker) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsMqBroker) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsMqBroker) PutConfiguration(value *AwsMqBroker_ConfigurationProperty) {
	if err := a.validatePutConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMqBroker) PutEncryptionOptions(value *AwsMqBroker_EncryptionOptionsProperty) {
	if err := a.validatePutEncryptionOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEncryptionOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMqBroker) PutLdapServerMetadata(value *AwsMqBroker_LdapServerMetadataProperty) {
	if err := a.validatePutLdapServerMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLdapServerMetadata",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMqBroker) PutLogs(value *AwsMqBroker_LogsProperty) {
	if err := a.validatePutLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMqBroker) PutMaintenanceWindowStartTime(value *AwsMqBroker_MaintenanceWindowStartTimeProperty) {
	if err := a.validatePutMaintenanceWindowStartTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMaintenanceWindowStartTime",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMqBroker) PutTimeouts(value *AwsMqBroker_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMqBroker) PutUser(value interface{}) {
	if err := a.validatePutUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUser",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMqBroker) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetApplyImmediately() {
	_jsii_.InvokeVoid(
		a,
		"resetApplyImmediately",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetAuthenticationStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticationStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetDataReplicationMode() {
	_jsii_.InvokeVoid(
		a,
		"resetDataReplicationMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetDataReplicationPrimaryBrokerArn() {
	_jsii_.InvokeVoid(
		a,
		"resetDataReplicationPrimaryBrokerArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetDeploymentMode() {
	_jsii_.InvokeVoid(
		a,
		"resetDeploymentMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetEncryptionOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetLdapServerMetadata() {
	_jsii_.InvokeVoid(
		a,
		"resetLdapServerMetadata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetMaintenanceWindowStartTime() {
	_jsii_.InvokeVoid(
		a,
		"resetMaintenanceWindowStartTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetPubliclyAccessible() {
	_jsii_.InvokeVoid(
		a,
		"resetPubliclyAccessible",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetResourceShareArns() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceShareArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetStorageType() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) ResetUser() {
	_jsii_.InvokeVoid(
		a,
		"resetUser",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMqBroker) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMqBroker) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMqBroker) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMqBroker) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMqBroker) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMqBroker) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

