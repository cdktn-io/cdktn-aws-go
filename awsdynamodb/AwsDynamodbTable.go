package awsdynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdynamodb/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsdynamodb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table aws_dynamodb_table}.
// Experimental.
type AwsDynamodbTable interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	Attribute() AwsDynamodbTable_AttributePropertyList
	// Experimental.
	AttributeInput() interface{}
	// Experimental.
	BillingMode() *string
	// Experimental.
	SetBillingMode(val *string)
	// Experimental.
	BillingModeInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	DeletionProtectionEnabled() interface{}
	// Experimental.
	SetDeletionProtectionEnabled(val interface{})
	// Experimental.
	DeletionProtectionEnabledInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	GlobalSecondaryIndex() AwsDynamodbTable_GlobalSecondaryIndexPropertyList
	// Experimental.
	GlobalSecondaryIndexInput() interface{}
	// Experimental.
	GlobalTableWitness() AwsDynamodbTable_GlobalTableWitnessPropertyOutputReference
	// Experimental.
	GlobalTableWitnessInput() *AwsDynamodbTable_GlobalTableWitnessProperty
	// Experimental.
	HashKey() *string
	// Experimental.
	SetHashKey(val *string)
	// Experimental.
	HashKeyInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	ImportTable() AwsDynamodbTable_ImportTablePropertyOutputReference
	// Experimental.
	ImportTableInput() *AwsDynamodbTable_ImportTableProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LocalSecondaryIndex() AwsDynamodbTable_LocalSecondaryIndexPropertyList
	// Experimental.
	LocalSecondaryIndexInput() interface{}
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OnDemandThroughput() AwsDynamodbTable_OnDemandThroughputPropertyOutputReference
	// Experimental.
	OnDemandThroughputInput() *AwsDynamodbTable_OnDemandThroughputProperty
	// Experimental.
	PointInTimeRecovery() AwsDynamodbTable_PointInTimeRecoveryPropertyOutputReference
	// Experimental.
	PointInTimeRecoveryInput() *AwsDynamodbTable_PointInTimeRecoveryProperty
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RangeKey() *string
	// Experimental.
	SetRangeKey(val *string)
	// Experimental.
	RangeKeyInput() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	ReadCapacity() *float64
	// Experimental.
	SetReadCapacity(val *float64)
	// Experimental.
	ReadCapacityInput() *float64
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	Replica() AwsDynamodbTable_ReplicaPropertyList
	// Experimental.
	ReplicaInput() interface{}
	// Experimental.
	RestoreBackupArn() *string
	// Experimental.
	SetRestoreBackupArn(val *string)
	// Experimental.
	RestoreBackupArnInput() *string
	// Experimental.
	RestoreDateTime() *string
	// Experimental.
	SetRestoreDateTime(val *string)
	// Experimental.
	RestoreDateTimeInput() *string
	// Experimental.
	RestoreSourceName() *string
	// Experimental.
	SetRestoreSourceName(val *string)
	// Experimental.
	RestoreSourceNameInput() *string
	// Experimental.
	RestoreSourceTableArn() *string
	// Experimental.
	SetRestoreSourceTableArn(val *string)
	// Experimental.
	RestoreSourceTableArnInput() *string
	// Experimental.
	RestoreToLatestTime() interface{}
	// Experimental.
	SetRestoreToLatestTime(val interface{})
	// Experimental.
	RestoreToLatestTimeInput() interface{}
	// Experimental.
	ServerSideEncryption() AwsDynamodbTable_ServerSideEncryptionPropertyOutputReference
	// Experimental.
	ServerSideEncryptionInput() *AwsDynamodbTable_ServerSideEncryptionProperty
	// Experimental.
	StreamArn() *string
	// Experimental.
	StreamEnabled() interface{}
	// Experimental.
	SetStreamEnabled(val interface{})
	// Experimental.
	StreamEnabledInput() interface{}
	// Experimental.
	StreamLabel() *string
	// Experimental.
	StreamViewType() *string
	// Experimental.
	SetStreamViewType(val *string)
	// Experimental.
	StreamViewTypeInput() *string
	// Experimental.
	TableClass() *string
	// Experimental.
	SetTableClass(val *string)
	// Experimental.
	TableClassInput() *string
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
	Timeouts() AwsDynamodbTable_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Ttl() AwsDynamodbTable_TtlPropertyOutputReference
	// Experimental.
	TtlInput() *AwsDynamodbTable_TtlProperty
	// Experimental.
	WarmThroughput() AwsDynamodbTable_WarmThroughputPropertyOutputReference
	// Experimental.
	WarmThroughputInput() *AwsDynamodbTable_WarmThroughputProperty
	// Experimental.
	WriteCapacity() *float64
	// Experimental.
	SetWriteCapacity(val *float64)
	// Experimental.
	WriteCapacityInput() *float64
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
	PutAttribute(value interface{})
	// Experimental.
	PutGlobalSecondaryIndex(value interface{})
	// Experimental.
	PutGlobalTableWitness(value *AwsDynamodbTable_GlobalTableWitnessProperty)
	// Experimental.
	PutImportTable(value *AwsDynamodbTable_ImportTableProperty)
	// Experimental.
	PutLocalSecondaryIndex(value interface{})
	// Experimental.
	PutOnDemandThroughput(value *AwsDynamodbTable_OnDemandThroughputProperty)
	// Experimental.
	PutPointInTimeRecovery(value *AwsDynamodbTable_PointInTimeRecoveryProperty)
	// Experimental.
	PutReplica(value interface{})
	// Experimental.
	PutServerSideEncryption(value *AwsDynamodbTable_ServerSideEncryptionProperty)
	// Experimental.
	PutTimeouts(value *AwsDynamodbTable_TimeoutsProperty)
	// Experimental.
	PutTtl(value *AwsDynamodbTable_TtlProperty)
	// Experimental.
	PutWarmThroughput(value *AwsDynamodbTable_WarmThroughputProperty)
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
	ResetAttribute()
	// Experimental.
	ResetBillingMode()
	// Experimental.
	ResetDeletionProtectionEnabled()
	// Experimental.
	ResetGlobalSecondaryIndex()
	// Experimental.
	ResetGlobalTableWitness()
	// Experimental.
	ResetHashKey()
	// Experimental.
	ResetId()
	// Experimental.
	ResetImportTable()
	// Experimental.
	ResetLocalSecondaryIndex()
	// Experimental.
	ResetOnDemandThroughput()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPointInTimeRecovery()
	// Experimental.
	ResetRangeKey()
	// Experimental.
	ResetReadCapacity()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetReplica()
	// Experimental.
	ResetRestoreBackupArn()
	// Experimental.
	ResetRestoreDateTime()
	// Experimental.
	ResetRestoreSourceName()
	// Experimental.
	ResetRestoreSourceTableArn()
	// Experimental.
	ResetRestoreToLatestTime()
	// Experimental.
	ResetServerSideEncryption()
	// Experimental.
	ResetStreamEnabled()
	// Experimental.
	ResetStreamViewType()
	// Experimental.
	ResetTableClass()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTtl()
	// Experimental.
	ResetWarmThroughput()
	// Experimental.
	ResetWriteCapacity()
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

// The jsii proxy struct for AwsDynamodbTable
type jsiiProxy_AwsDynamodbTable struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsDynamodbTable) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) Attribute() AwsDynamodbTable_AttributePropertyList {
	var returns AwsDynamodbTable_AttributePropertyList
	_jsii_.Get(
		j,
		"attribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) AttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) BillingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) BillingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) DeletionProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) GlobalSecondaryIndex() AwsDynamodbTable_GlobalSecondaryIndexPropertyList {
	var returns AwsDynamodbTable_GlobalSecondaryIndexPropertyList
	_jsii_.Get(
		j,
		"globalSecondaryIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) GlobalSecondaryIndexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalSecondaryIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) GlobalTableWitness() AwsDynamodbTable_GlobalTableWitnessPropertyOutputReference {
	var returns AwsDynamodbTable_GlobalTableWitnessPropertyOutputReference
	_jsii_.Get(
		j,
		"globalTableWitness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) GlobalTableWitnessInput() *AwsDynamodbTable_GlobalTableWitnessProperty {
	var returns *AwsDynamodbTable_GlobalTableWitnessProperty
	_jsii_.Get(
		j,
		"globalTableWitnessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) HashKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) HashKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) ImportTable() AwsDynamodbTable_ImportTablePropertyOutputReference {
	var returns AwsDynamodbTable_ImportTablePropertyOutputReference
	_jsii_.Get(
		j,
		"importTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) ImportTableInput() *AwsDynamodbTable_ImportTableProperty {
	var returns *AwsDynamodbTable_ImportTableProperty
	_jsii_.Get(
		j,
		"importTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) LocalSecondaryIndex() AwsDynamodbTable_LocalSecondaryIndexPropertyList {
	var returns AwsDynamodbTable_LocalSecondaryIndexPropertyList
	_jsii_.Get(
		j,
		"localSecondaryIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) LocalSecondaryIndexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localSecondaryIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) OnDemandThroughput() AwsDynamodbTable_OnDemandThroughputPropertyOutputReference {
	var returns AwsDynamodbTable_OnDemandThroughputPropertyOutputReference
	_jsii_.Get(
		j,
		"onDemandThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) OnDemandThroughputInput() *AwsDynamodbTable_OnDemandThroughputProperty {
	var returns *AwsDynamodbTable_OnDemandThroughputProperty
	_jsii_.Get(
		j,
		"onDemandThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) PointInTimeRecovery() AwsDynamodbTable_PointInTimeRecoveryPropertyOutputReference {
	var returns AwsDynamodbTable_PointInTimeRecoveryPropertyOutputReference
	_jsii_.Get(
		j,
		"pointInTimeRecovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) PointInTimeRecoveryInput() *AwsDynamodbTable_PointInTimeRecoveryProperty {
	var returns *AwsDynamodbTable_PointInTimeRecoveryProperty
	_jsii_.Get(
		j,
		"pointInTimeRecoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) RangeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) RangeKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) ReadCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) ReadCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) Replica() AwsDynamodbTable_ReplicaPropertyList {
	var returns AwsDynamodbTable_ReplicaPropertyList
	_jsii_.Get(
		j,
		"replica",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) ReplicaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) RestoreBackupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreBackupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) RestoreBackupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreBackupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) RestoreDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) RestoreDateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreDateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) RestoreSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) RestoreSourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreSourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) RestoreSourceTableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreSourceTableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) RestoreSourceTableArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreSourceTableArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) RestoreToLatestTime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restoreToLatestTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) RestoreToLatestTimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restoreToLatestTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) ServerSideEncryption() AwsDynamodbTable_ServerSideEncryptionPropertyOutputReference {
	var returns AwsDynamodbTable_ServerSideEncryptionPropertyOutputReference
	_jsii_.Get(
		j,
		"serverSideEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) ServerSideEncryptionInput() *AwsDynamodbTable_ServerSideEncryptionProperty {
	var returns *AwsDynamodbTable_ServerSideEncryptionProperty
	_jsii_.Get(
		j,
		"serverSideEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) StreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) StreamEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) StreamEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) StreamLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) StreamViewType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamViewType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) StreamViewTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamViewTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) TableClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) TableClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) Timeouts() AwsDynamodbTable_TimeoutsPropertyOutputReference {
	var returns AwsDynamodbTable_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) Ttl() AwsDynamodbTable_TtlPropertyOutputReference {
	var returns AwsDynamodbTable_TtlPropertyOutputReference
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) TtlInput() *AwsDynamodbTable_TtlProperty {
	var returns *AwsDynamodbTable_TtlProperty
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) WarmThroughput() AwsDynamodbTable_WarmThroughputPropertyOutputReference {
	var returns AwsDynamodbTable_WarmThroughputPropertyOutputReference
	_jsii_.Get(
		j,
		"warmThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) WarmThroughputInput() *AwsDynamodbTable_WarmThroughputProperty {
	var returns *AwsDynamodbTable_WarmThroughputProperty
	_jsii_.Get(
		j,
		"warmThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) WriteCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDynamodbTable) WriteCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeCapacityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table aws_dynamodb_table} Resource.
// Experimental.
func NewAwsDynamodbTable(scope constructs.Construct, id *string, config *AwsDynamodbTableConfig) AwsDynamodbTable {
	_init_.Initialize()

	if err := validateNewAwsDynamodbTableParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDynamodbTable{}

	_jsii_.Create(
		"@cdktn/aws-dynamodb.AwsDynamodbTable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table aws_dynamodb_table} Resource.
// Experimental.
func NewAwsDynamodbTable_Override(a AwsDynamodbTable, scope constructs.Construct, id *string, config *AwsDynamodbTableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dynamodb.AwsDynamodbTable",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetBillingMode(val *string) {
	if err := j.validateSetBillingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingMode",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetHashKey(val *string) {
	if err := j.validateSetHashKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashKey",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetRangeKey(val *string) {
	if err := j.validateSetRangeKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rangeKey",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetReadCapacity(val *float64) {
	if err := j.validateSetReadCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetRestoreBackupArn(val *string) {
	if err := j.validateSetRestoreBackupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreBackupArn",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetRestoreDateTime(val *string) {
	if err := j.validateSetRestoreDateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreDateTime",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetRestoreSourceName(val *string) {
	if err := j.validateSetRestoreSourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreSourceName",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetRestoreSourceTableArn(val *string) {
	if err := j.validateSetRestoreSourceTableArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreSourceTableArn",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetRestoreToLatestTime(val interface{}) {
	if err := j.validateSetRestoreToLatestTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreToLatestTime",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetStreamEnabled(val interface{}) {
	if err := j.validateSetStreamEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetStreamViewType(val *string) {
	if err := j.validateSetStreamViewTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamViewType",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetTableClass(val *string) {
	if err := j.validateSetTableClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableClass",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsDynamodbTable)SetWriteCapacity(val *float64) {
	if err := j.validateSetWriteCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeCapacity",
		val,
	)
}

// Generates CDKTN code for importing a AwsDynamodbTable resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsDynamodbTable_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsDynamodbTable_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-dynamodb.AwsDynamodbTable",
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
func AwsDynamodbTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDynamodbTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dynamodb.AwsDynamodbTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsDynamodbTable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDynamodbTable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dynamodb.AwsDynamodbTable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsDynamodbTable_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDynamodbTable_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dynamodb.AwsDynamodbTable",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsDynamodbTable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-dynamodb.AwsDynamodbTable",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsDynamodbTable) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDynamodbTable) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDynamodbTable) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDynamodbTable) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDynamodbTable) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDynamodbTable) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDynamodbTable) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDynamodbTable) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDynamodbTable) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDynamodbTable) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDynamodbTable) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDynamodbTable) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsDynamodbTable) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) PutAttribute(value interface{}) {
	if err := a.validatePutAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAttribute",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) PutGlobalSecondaryIndex(value interface{}) {
	if err := a.validatePutGlobalSecondaryIndexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGlobalSecondaryIndex",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) PutGlobalTableWitness(value *AwsDynamodbTable_GlobalTableWitnessProperty) {
	if err := a.validatePutGlobalTableWitnessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGlobalTableWitness",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) PutImportTable(value *AwsDynamodbTable_ImportTableProperty) {
	if err := a.validatePutImportTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putImportTable",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) PutLocalSecondaryIndex(value interface{}) {
	if err := a.validatePutLocalSecondaryIndexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLocalSecondaryIndex",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) PutOnDemandThroughput(value *AwsDynamodbTable_OnDemandThroughputProperty) {
	if err := a.validatePutOnDemandThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOnDemandThroughput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) PutPointInTimeRecovery(value *AwsDynamodbTable_PointInTimeRecoveryProperty) {
	if err := a.validatePutPointInTimeRecoveryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPointInTimeRecovery",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) PutReplica(value interface{}) {
	if err := a.validatePutReplicaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReplica",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) PutServerSideEncryption(value *AwsDynamodbTable_ServerSideEncryptionProperty) {
	if err := a.validatePutServerSideEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServerSideEncryption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) PutTimeouts(value *AwsDynamodbTable_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) PutTtl(value *AwsDynamodbTable_TtlProperty) {
	if err := a.validatePutTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTtl",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) PutWarmThroughput(value *AwsDynamodbTable_WarmThroughputProperty) {
	if err := a.validatePutWarmThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWarmThroughput",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetAttribute() {
	_jsii_.InvokeVoid(
		a,
		"resetAttribute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetBillingMode() {
	_jsii_.InvokeVoid(
		a,
		"resetBillingMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetDeletionProtectionEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDeletionProtectionEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetGlobalSecondaryIndex() {
	_jsii_.InvokeVoid(
		a,
		"resetGlobalSecondaryIndex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetGlobalTableWitness() {
	_jsii_.InvokeVoid(
		a,
		"resetGlobalTableWitness",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetHashKey() {
	_jsii_.InvokeVoid(
		a,
		"resetHashKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetImportTable() {
	_jsii_.InvokeVoid(
		a,
		"resetImportTable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetLocalSecondaryIndex() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalSecondaryIndex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetOnDemandThroughput() {
	_jsii_.InvokeVoid(
		a,
		"resetOnDemandThroughput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetPointInTimeRecovery() {
	_jsii_.InvokeVoid(
		a,
		"resetPointInTimeRecovery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetRangeKey() {
	_jsii_.InvokeVoid(
		a,
		"resetRangeKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetReadCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetReadCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetReplica() {
	_jsii_.InvokeVoid(
		a,
		"resetReplica",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetRestoreBackupArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRestoreBackupArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetRestoreDateTime() {
	_jsii_.InvokeVoid(
		a,
		"resetRestoreDateTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetRestoreSourceName() {
	_jsii_.InvokeVoid(
		a,
		"resetRestoreSourceName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetRestoreSourceTableArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRestoreSourceTableArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetRestoreToLatestTime() {
	_jsii_.InvokeVoid(
		a,
		"resetRestoreToLatestTime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetServerSideEncryption() {
	_jsii_.InvokeVoid(
		a,
		"resetServerSideEncryption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetStreamEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetStreamEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetStreamViewType() {
	_jsii_.InvokeVoid(
		a,
		"resetStreamViewType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetTableClass() {
	_jsii_.InvokeVoid(
		a,
		"resetTableClass",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetWarmThroughput() {
	_jsii_.InvokeVoid(
		a,
		"resetWarmThroughput",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) ResetWriteCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetWriteCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDynamodbTable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDynamodbTable) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDynamodbTable) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDynamodbTable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDynamodbTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDynamodbTable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDynamodbTable) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

