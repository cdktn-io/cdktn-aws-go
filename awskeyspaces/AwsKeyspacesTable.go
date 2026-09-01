package awskeyspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskeyspaces/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awskeyspaces/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table aws_keyspaces_table}.
// Experimental.
type AwsKeyspacesTable interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	CapacitySpecification() AwsKeyspacesTable_CapacitySpecificationPropertyOutputReference
	// Experimental.
	CapacitySpecificationInput() *AwsKeyspacesTable_CapacitySpecificationProperty
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClientSideTimestamps() AwsKeyspacesTable_ClientSideTimestampsPropertyOutputReference
	// Experimental.
	ClientSideTimestampsInput() *AwsKeyspacesTable_ClientSideTimestampsProperty
	// Experimental.
	Comment() AwsKeyspacesTable_CommentPropertyOutputReference
	// Experimental.
	CommentInput() *AwsKeyspacesTable_CommentProperty
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
	DefaultTimeToLive() *float64
	// Experimental.
	SetDefaultTimeToLive(val *float64)
	// Experimental.
	DefaultTimeToLiveInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EncryptionSpecification() AwsKeyspacesTable_EncryptionSpecificationPropertyOutputReference
	// Experimental.
	EncryptionSpecificationInput() *AwsKeyspacesTable_EncryptionSpecificationProperty
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
	KeyspaceName() *string
	// Experimental.
	SetKeyspaceName(val *string)
	// Experimental.
	KeyspaceNameInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PointInTimeRecovery() AwsKeyspacesTable_PointInTimeRecoveryPropertyOutputReference
	// Experimental.
	PointInTimeRecoveryInput() *AwsKeyspacesTable_PointInTimeRecoveryProperty
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
	SchemaDefinition() AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference
	// Experimental.
	SchemaDefinitionInput() *AwsKeyspacesTable_SchemaDefinitionProperty
	// Experimental.
	TableName() *string
	// Experimental.
	SetTableName(val *string)
	// Experimental.
	TableNameInput() *string
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
	Timeouts() AwsKeyspacesTable_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Ttl() AwsKeyspacesTable_TtlPropertyOutputReference
	// Experimental.
	TtlInput() *AwsKeyspacesTable_TtlProperty
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
	PutCapacitySpecification(value *AwsKeyspacesTable_CapacitySpecificationProperty)
	// Experimental.
	PutClientSideTimestamps(value *AwsKeyspacesTable_ClientSideTimestampsProperty)
	// Experimental.
	PutComment(value *AwsKeyspacesTable_CommentProperty)
	// Experimental.
	PutEncryptionSpecification(value *AwsKeyspacesTable_EncryptionSpecificationProperty)
	// Experimental.
	PutPointInTimeRecovery(value *AwsKeyspacesTable_PointInTimeRecoveryProperty)
	// Experimental.
	PutSchemaDefinition(value *AwsKeyspacesTable_SchemaDefinitionProperty)
	// Experimental.
	PutTimeouts(value *AwsKeyspacesTable_TimeoutsProperty)
	// Experimental.
	PutTtl(value *AwsKeyspacesTable_TtlProperty)
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
	ResetCapacitySpecification()
	// Experimental.
	ResetClientSideTimestamps()
	// Experimental.
	ResetComment()
	// Experimental.
	ResetDefaultTimeToLive()
	// Experimental.
	ResetEncryptionSpecification()
	// Experimental.
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPointInTimeRecovery()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTtl()
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

// The jsii proxy struct for AwsKeyspacesTable
type jsiiProxy_AwsKeyspacesTable struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsKeyspacesTable) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) CapacitySpecification() AwsKeyspacesTable_CapacitySpecificationPropertyOutputReference {
	var returns AwsKeyspacesTable_CapacitySpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"capacitySpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) CapacitySpecificationInput() *AwsKeyspacesTable_CapacitySpecificationProperty {
	var returns *AwsKeyspacesTable_CapacitySpecificationProperty
	_jsii_.Get(
		j,
		"capacitySpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) ClientSideTimestamps() AwsKeyspacesTable_ClientSideTimestampsPropertyOutputReference {
	var returns AwsKeyspacesTable_ClientSideTimestampsPropertyOutputReference
	_jsii_.Get(
		j,
		"clientSideTimestamps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) ClientSideTimestampsInput() *AwsKeyspacesTable_ClientSideTimestampsProperty {
	var returns *AwsKeyspacesTable_ClientSideTimestampsProperty
	_jsii_.Get(
		j,
		"clientSideTimestampsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) Comment() AwsKeyspacesTable_CommentPropertyOutputReference {
	var returns AwsKeyspacesTable_CommentPropertyOutputReference
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) CommentInput() *AwsKeyspacesTable_CommentProperty {
	var returns *AwsKeyspacesTable_CommentProperty
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) DefaultTimeToLive() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTimeToLive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) DefaultTimeToLiveInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTimeToLiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) EncryptionSpecification() AwsKeyspacesTable_EncryptionSpecificationPropertyOutputReference {
	var returns AwsKeyspacesTable_EncryptionSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"encryptionSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) EncryptionSpecificationInput() *AwsKeyspacesTable_EncryptionSpecificationProperty {
	var returns *AwsKeyspacesTable_EncryptionSpecificationProperty
	_jsii_.Get(
		j,
		"encryptionSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) KeyspaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyspaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) KeyspaceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyspaceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) PointInTimeRecovery() AwsKeyspacesTable_PointInTimeRecoveryPropertyOutputReference {
	var returns AwsKeyspacesTable_PointInTimeRecoveryPropertyOutputReference
	_jsii_.Get(
		j,
		"pointInTimeRecovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) PointInTimeRecoveryInput() *AwsKeyspacesTable_PointInTimeRecoveryProperty {
	var returns *AwsKeyspacesTable_PointInTimeRecoveryProperty
	_jsii_.Get(
		j,
		"pointInTimeRecoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) SchemaDefinition() AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference {
	var returns AwsKeyspacesTable_SchemaDefinitionPropertyOutputReference
	_jsii_.Get(
		j,
		"schemaDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) SchemaDefinitionInput() *AwsKeyspacesTable_SchemaDefinitionProperty {
	var returns *AwsKeyspacesTable_SchemaDefinitionProperty
	_jsii_.Get(
		j,
		"schemaDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) Timeouts() AwsKeyspacesTable_TimeoutsPropertyOutputReference {
	var returns AwsKeyspacesTable_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) Ttl() AwsKeyspacesTable_TtlPropertyOutputReference {
	var returns AwsKeyspacesTable_TtlPropertyOutputReference
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKeyspacesTable) TtlInput() *AwsKeyspacesTable_TtlProperty {
	var returns *AwsKeyspacesTable_TtlProperty
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table aws_keyspaces_table} Resource.
// Experimental.
func NewAwsKeyspacesTable(scope constructs.Construct, id *string, config *AwsKeyspacesTableConfig) AwsKeyspacesTable {
	_init_.Initialize()

	if err := validateNewAwsKeyspacesTableParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKeyspacesTable{}

	_jsii_.Create(
		"@cdktn/aws-keyspaces.AwsKeyspacesTable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table aws_keyspaces_table} Resource.
// Experimental.
func NewAwsKeyspacesTable_Override(a AwsKeyspacesTable, scope constructs.Construct, id *string, config *AwsKeyspacesTableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-keyspaces.AwsKeyspacesTable",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable)SetDefaultTimeToLive(val *float64) {
	if err := j.validateSetDefaultTimeToLiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTimeToLive",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable)SetKeyspaceName(val *string) {
	if err := j.validateSetKeyspaceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyspaceName",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsKeyspacesTable)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a AwsKeyspacesTable resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsKeyspacesTable_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsKeyspacesTable_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-keyspaces.AwsKeyspacesTable",
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
func AwsKeyspacesTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsKeyspacesTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-keyspaces.AwsKeyspacesTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsKeyspacesTable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsKeyspacesTable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-keyspaces.AwsKeyspacesTable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsKeyspacesTable_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsKeyspacesTable_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-keyspaces.AwsKeyspacesTable",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsKeyspacesTable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-keyspaces.AwsKeyspacesTable",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsKeyspacesTable) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKeyspacesTable) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKeyspacesTable) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKeyspacesTable) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKeyspacesTable) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKeyspacesTable) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKeyspacesTable) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKeyspacesTable) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKeyspacesTable) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKeyspacesTable) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKeyspacesTable) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKeyspacesTable) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsKeyspacesTable) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) PutCapacitySpecification(value *AwsKeyspacesTable_CapacitySpecificationProperty) {
	if err := a.validatePutCapacitySpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCapacitySpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) PutClientSideTimestamps(value *AwsKeyspacesTable_ClientSideTimestampsProperty) {
	if err := a.validatePutClientSideTimestampsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putClientSideTimestamps",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) PutComment(value *AwsKeyspacesTable_CommentProperty) {
	if err := a.validatePutCommentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putComment",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) PutEncryptionSpecification(value *AwsKeyspacesTable_EncryptionSpecificationProperty) {
	if err := a.validatePutEncryptionSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEncryptionSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) PutPointInTimeRecovery(value *AwsKeyspacesTable_PointInTimeRecoveryProperty) {
	if err := a.validatePutPointInTimeRecoveryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPointInTimeRecovery",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) PutSchemaDefinition(value *AwsKeyspacesTable_SchemaDefinitionProperty) {
	if err := a.validatePutSchemaDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchemaDefinition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) PutTimeouts(value *AwsKeyspacesTable_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) PutTtl(value *AwsKeyspacesTable_TtlProperty) {
	if err := a.validatePutTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTtl",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) ResetCapacitySpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacitySpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) ResetClientSideTimestamps() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSideTimestamps",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) ResetComment() {
	_jsii_.InvokeVoid(
		a,
		"resetComment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) ResetDefaultTimeToLive() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultTimeToLive",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) ResetEncryptionSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) ResetPointInTimeRecovery() {
	_jsii_.InvokeVoid(
		a,
		"resetPointInTimeRecovery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) ResetTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKeyspacesTable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKeyspacesTable) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKeyspacesTable) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKeyspacesTable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKeyspacesTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKeyspacesTable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKeyspacesTable) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

