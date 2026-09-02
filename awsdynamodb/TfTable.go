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
type TfTable interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	Attribute() TfTable_AttributePropertyList
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
	GlobalSecondaryIndex() TfTable_GlobalSecondaryIndexPropertyList
	// Experimental.
	GlobalSecondaryIndexInput() interface{}
	// Experimental.
	GlobalTableWitness() TfTable_GlobalTableWitnessPropertyOutputReference
	// Experimental.
	GlobalTableWitnessInput() *TfTable_GlobalTableWitnessProperty
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
	ImportTable() TfTable_ImportTablePropertyOutputReference
	// Experimental.
	ImportTableInput() *TfTable_ImportTableProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LocalSecondaryIndex() TfTable_LocalSecondaryIndexPropertyList
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
	OnDemandThroughput() TfTable_OnDemandThroughputPropertyOutputReference
	// Experimental.
	OnDemandThroughputInput() *TfTable_OnDemandThroughputProperty
	// Experimental.
	PointInTimeRecovery() TfTable_PointInTimeRecoveryPropertyOutputReference
	// Experimental.
	PointInTimeRecoveryInput() *TfTable_PointInTimeRecoveryProperty
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
	Replica() TfTable_ReplicaPropertyList
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
	ServerSideEncryption() TfTable_ServerSideEncryptionPropertyOutputReference
	// Experimental.
	ServerSideEncryptionInput() *TfTable_ServerSideEncryptionProperty
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
	Timeouts() TfTable_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Ttl() TfTable_TtlPropertyOutputReference
	// Experimental.
	TtlInput() *TfTable_TtlProperty
	// Experimental.
	WarmThroughput() TfTable_WarmThroughputPropertyOutputReference
	// Experimental.
	WarmThroughputInput() *TfTable_WarmThroughputProperty
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
	PutGlobalTableWitness(value *TfTable_GlobalTableWitnessProperty)
	// Experimental.
	PutImportTable(value *TfTable_ImportTableProperty)
	// Experimental.
	PutLocalSecondaryIndex(value interface{})
	// Experimental.
	PutOnDemandThroughput(value *TfTable_OnDemandThroughputProperty)
	// Experimental.
	PutPointInTimeRecovery(value *TfTable_PointInTimeRecoveryProperty)
	// Experimental.
	PutReplica(value interface{})
	// Experimental.
	PutServerSideEncryption(value *TfTable_ServerSideEncryptionProperty)
	// Experimental.
	PutTimeouts(value *TfTable_TimeoutsProperty)
	// Experimental.
	PutTtl(value *TfTable_TtlProperty)
	// Experimental.
	PutWarmThroughput(value *TfTable_WarmThroughputProperty)
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

// The jsii proxy struct for TfTable
type jsiiProxy_TfTable struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfTable) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) Attribute() TfTable_AttributePropertyList {
	var returns TfTable_AttributePropertyList
	_jsii_.Get(
		j,
		"attribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) AttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) BillingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) BillingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) DeletionProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) GlobalSecondaryIndex() TfTable_GlobalSecondaryIndexPropertyList {
	var returns TfTable_GlobalSecondaryIndexPropertyList
	_jsii_.Get(
		j,
		"globalSecondaryIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) GlobalSecondaryIndexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalSecondaryIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) GlobalTableWitness() TfTable_GlobalTableWitnessPropertyOutputReference {
	var returns TfTable_GlobalTableWitnessPropertyOutputReference
	_jsii_.Get(
		j,
		"globalTableWitness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) GlobalTableWitnessInput() *TfTable_GlobalTableWitnessProperty {
	var returns *TfTable_GlobalTableWitnessProperty
	_jsii_.Get(
		j,
		"globalTableWitnessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) HashKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) HashKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) ImportTable() TfTable_ImportTablePropertyOutputReference {
	var returns TfTable_ImportTablePropertyOutputReference
	_jsii_.Get(
		j,
		"importTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) ImportTableInput() *TfTable_ImportTableProperty {
	var returns *TfTable_ImportTableProperty
	_jsii_.Get(
		j,
		"importTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) LocalSecondaryIndex() TfTable_LocalSecondaryIndexPropertyList {
	var returns TfTable_LocalSecondaryIndexPropertyList
	_jsii_.Get(
		j,
		"localSecondaryIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) LocalSecondaryIndexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localSecondaryIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) OnDemandThroughput() TfTable_OnDemandThroughputPropertyOutputReference {
	var returns TfTable_OnDemandThroughputPropertyOutputReference
	_jsii_.Get(
		j,
		"onDemandThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) OnDemandThroughputInput() *TfTable_OnDemandThroughputProperty {
	var returns *TfTable_OnDemandThroughputProperty
	_jsii_.Get(
		j,
		"onDemandThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) PointInTimeRecovery() TfTable_PointInTimeRecoveryPropertyOutputReference {
	var returns TfTable_PointInTimeRecoveryPropertyOutputReference
	_jsii_.Get(
		j,
		"pointInTimeRecovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) PointInTimeRecoveryInput() *TfTable_PointInTimeRecoveryProperty {
	var returns *TfTable_PointInTimeRecoveryProperty
	_jsii_.Get(
		j,
		"pointInTimeRecoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) RangeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) RangeKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) ReadCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) ReadCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) Replica() TfTable_ReplicaPropertyList {
	var returns TfTable_ReplicaPropertyList
	_jsii_.Get(
		j,
		"replica",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) ReplicaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) RestoreBackupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreBackupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) RestoreBackupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreBackupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) RestoreDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) RestoreDateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreDateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) RestoreSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) RestoreSourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreSourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) RestoreSourceTableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreSourceTableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) RestoreSourceTableArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreSourceTableArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) RestoreToLatestTime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restoreToLatestTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) RestoreToLatestTimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restoreToLatestTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) ServerSideEncryption() TfTable_ServerSideEncryptionPropertyOutputReference {
	var returns TfTable_ServerSideEncryptionPropertyOutputReference
	_jsii_.Get(
		j,
		"serverSideEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) ServerSideEncryptionInput() *TfTable_ServerSideEncryptionProperty {
	var returns *TfTable_ServerSideEncryptionProperty
	_jsii_.Get(
		j,
		"serverSideEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) StreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) StreamEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) StreamEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) StreamLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) StreamViewType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamViewType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) StreamViewTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamViewTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) TableClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) TableClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) Timeouts() TfTable_TimeoutsPropertyOutputReference {
	var returns TfTable_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) Ttl() TfTable_TtlPropertyOutputReference {
	var returns TfTable_TtlPropertyOutputReference
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) TtlInput() *TfTable_TtlProperty {
	var returns *TfTable_TtlProperty
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) WarmThroughput() TfTable_WarmThroughputPropertyOutputReference {
	var returns TfTable_WarmThroughputPropertyOutputReference
	_jsii_.Get(
		j,
		"warmThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) WarmThroughputInput() *TfTable_WarmThroughputProperty {
	var returns *TfTable_WarmThroughputProperty
	_jsii_.Get(
		j,
		"warmThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) WriteCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTable) WriteCapacityInput() *float64 {
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
func NewTfTable(scope constructs.Construct, id *string, config *TfTableConfig) TfTable {
	_init_.Initialize()

	if err := validateNewTfTableParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTable{}

	_jsii_.Create(
		"@cdktn/aws-dynamodb.TfTable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table aws_dynamodb_table} Resource.
// Experimental.
func NewTfTable_Override(t TfTable, scope constructs.Construct, id *string, config *TfTableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dynamodb.TfTable",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfTable)SetBillingMode(val *string) {
	if err := j.validateSetBillingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingMode",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetHashKey(val *string) {
	if err := j.validateSetHashKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashKey",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetRangeKey(val *string) {
	if err := j.validateSetRangeKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rangeKey",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetReadCapacity(val *float64) {
	if err := j.validateSetReadCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readCapacity",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetRestoreBackupArn(val *string) {
	if err := j.validateSetRestoreBackupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreBackupArn",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetRestoreDateTime(val *string) {
	if err := j.validateSetRestoreDateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreDateTime",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetRestoreSourceName(val *string) {
	if err := j.validateSetRestoreSourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreSourceName",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetRestoreSourceTableArn(val *string) {
	if err := j.validateSetRestoreSourceTableArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreSourceTableArn",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetRestoreToLatestTime(val interface{}) {
	if err := j.validateSetRestoreToLatestTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreToLatestTime",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetStreamEnabled(val interface{}) {
	if err := j.validateSetStreamEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamEnabled",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetStreamViewType(val *string) {
	if err := j.validateSetStreamViewTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamViewType",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetTableClass(val *string) {
	if err := j.validateSetTableClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableClass",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfTable)SetWriteCapacity(val *float64) {
	if err := j.validateSetWriteCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeCapacity",
		val,
	)
}

// Generates CDKTN code for importing a TfTable resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfTable_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfTable_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-dynamodb.TfTable",
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
func TfTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dynamodb.TfTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfTable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dynamodb.TfTable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfTable_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTable_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dynamodb.TfTable",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfTable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-dynamodb.TfTable",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfTable) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfTable) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfTable) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTable) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTable) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTable) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTable) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTable) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTable) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTable) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTable) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTable) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfTable) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTable) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfTable) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfTable) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfTable) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfTable) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfTable) PutAttribute(value interface{}) {
	if err := t.validatePutAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAttribute",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable) PutGlobalSecondaryIndex(value interface{}) {
	if err := t.validatePutGlobalSecondaryIndexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGlobalSecondaryIndex",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable) PutGlobalTableWitness(value *TfTable_GlobalTableWitnessProperty) {
	if err := t.validatePutGlobalTableWitnessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGlobalTableWitness",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable) PutImportTable(value *TfTable_ImportTableProperty) {
	if err := t.validatePutImportTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putImportTable",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable) PutLocalSecondaryIndex(value interface{}) {
	if err := t.validatePutLocalSecondaryIndexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLocalSecondaryIndex",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable) PutOnDemandThroughput(value *TfTable_OnDemandThroughputProperty) {
	if err := t.validatePutOnDemandThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOnDemandThroughput",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable) PutPointInTimeRecovery(value *TfTable_PointInTimeRecoveryProperty) {
	if err := t.validatePutPointInTimeRecoveryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPointInTimeRecovery",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable) PutReplica(value interface{}) {
	if err := t.validatePutReplicaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putReplica",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable) PutServerSideEncryption(value *TfTable_ServerSideEncryptionProperty) {
	if err := t.validatePutServerSideEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putServerSideEncryption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable) PutTimeouts(value *TfTable_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable) PutTtl(value *TfTable_TtlProperty) {
	if err := t.validatePutTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTtl",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable) PutWarmThroughput(value *TfTable_WarmThroughputProperty) {
	if err := t.validatePutWarmThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWarmThroughput",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTable) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfTable) ResetAttribute() {
	_jsii_.InvokeVoid(
		t,
		"resetAttribute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetBillingMode() {
	_jsii_.InvokeVoid(
		t,
		"resetBillingMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetDeletionProtectionEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetDeletionProtectionEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetGlobalSecondaryIndex() {
	_jsii_.InvokeVoid(
		t,
		"resetGlobalSecondaryIndex",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetGlobalTableWitness() {
	_jsii_.InvokeVoid(
		t,
		"resetGlobalTableWitness",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetHashKey() {
	_jsii_.InvokeVoid(
		t,
		"resetHashKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetImportTable() {
	_jsii_.InvokeVoid(
		t,
		"resetImportTable",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetLocalSecondaryIndex() {
	_jsii_.InvokeVoid(
		t,
		"resetLocalSecondaryIndex",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetOnDemandThroughput() {
	_jsii_.InvokeVoid(
		t,
		"resetOnDemandThroughput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetPointInTimeRecovery() {
	_jsii_.InvokeVoid(
		t,
		"resetPointInTimeRecovery",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetRangeKey() {
	_jsii_.InvokeVoid(
		t,
		"resetRangeKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetReadCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetReadCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetReplica() {
	_jsii_.InvokeVoid(
		t,
		"resetReplica",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetRestoreBackupArn() {
	_jsii_.InvokeVoid(
		t,
		"resetRestoreBackupArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetRestoreDateTime() {
	_jsii_.InvokeVoid(
		t,
		"resetRestoreDateTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetRestoreSourceName() {
	_jsii_.InvokeVoid(
		t,
		"resetRestoreSourceName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetRestoreSourceTableArn() {
	_jsii_.InvokeVoid(
		t,
		"resetRestoreSourceTableArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetRestoreToLatestTime() {
	_jsii_.InvokeVoid(
		t,
		"resetRestoreToLatestTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetServerSideEncryption() {
	_jsii_.InvokeVoid(
		t,
		"resetServerSideEncryption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetStreamEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetStreamEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetStreamViewType() {
	_jsii_.InvokeVoid(
		t,
		"resetStreamViewType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetTableClass() {
	_jsii_.InvokeVoid(
		t,
		"resetTableClass",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetTtl() {
	_jsii_.InvokeVoid(
		t,
		"resetTtl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetWarmThroughput() {
	_jsii_.InvokeVoid(
		t,
		"resetWarmThroughput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) ResetWriteCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetWriteCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTable) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

