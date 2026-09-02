package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdms/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsdms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task aws_dms_replication_task}.
// Experimental.
type TfReplicationTask interface {
	cdktn.TerraformResource
	// Experimental.
	CdcStartPosition() *string
	// Experimental.
	SetCdcStartPosition(val *string)
	// Experimental.
	CdcStartPositionInput() *string
	// Experimental.
	CdcStartTime() *string
	// Experimental.
	SetCdcStartTime(val *string)
	// Experimental.
	CdcStartTimeInput() *string
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
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MigrationType() *string
	// Experimental.
	SetMigrationType(val *string)
	// Experimental.
	MigrationTypeInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	ReplicationInstanceArn() *string
	// Experimental.
	SetReplicationInstanceArn(val *string)
	// Experimental.
	ReplicationInstanceArnInput() *string
	// Experimental.
	ReplicationTaskArn() *string
	// Experimental.
	ReplicationTaskId() *string
	// Experimental.
	SetReplicationTaskId(val *string)
	// Experimental.
	ReplicationTaskIdInput() *string
	// Experimental.
	ReplicationTaskSettings() *string
	// Experimental.
	SetReplicationTaskSettings(val *string)
	// Experimental.
	ReplicationTaskSettingsInput() *string
	// Experimental.
	ResourceIdentifier() *string
	// Experimental.
	SetResourceIdentifier(val *string)
	// Experimental.
	ResourceIdentifierInput() *string
	// Experimental.
	SourceEndpointArn() *string
	// Experimental.
	SetSourceEndpointArn(val *string)
	// Experimental.
	SourceEndpointArnInput() *string
	// Experimental.
	StartReplicationTask() interface{}
	// Experimental.
	SetStartReplicationTask(val interface{})
	// Experimental.
	StartReplicationTaskInput() interface{}
	// Experimental.
	Status() *string
	// Experimental.
	TableMappings() *string
	// Experimental.
	SetTableMappings(val *string)
	// Experimental.
	TableMappingsInput() *string
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
	TargetEndpointArn() *string
	// Experimental.
	SetTargetEndpointArn(val *string)
	// Experimental.
	TargetEndpointArnInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetCdcStartPosition()
	// Experimental.
	ResetCdcStartTime()
	// Experimental.
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetReplicationTaskSettings()
	// Experimental.
	ResetResourceIdentifier()
	// Experimental.
	ResetStartReplicationTask()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
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

// The jsii proxy struct for TfReplicationTask
type jsiiProxy_TfReplicationTask struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfReplicationTask) CdcStartPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcStartPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) CdcStartPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcStartPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) CdcStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) CdcStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) MigrationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) MigrationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) ReplicationInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) ReplicationInstanceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationInstanceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) ReplicationTaskArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationTaskArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) ReplicationTaskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationTaskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) ReplicationTaskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationTaskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) ReplicationTaskSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationTaskSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) ReplicationTaskSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationTaskSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) ResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) ResourceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) SourceEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) SourceEndpointArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceEndpointArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) StartReplicationTask() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startReplicationTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) StartReplicationTaskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startReplicationTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) TableMappings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) TableMappingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) TargetEndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetEndpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) TargetEndpointArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetEndpointArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationTask) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task aws_dms_replication_task} Resource.
// Experimental.
func NewTfReplicationTask(scope constructs.Construct, id *string, config *TfReplicationTaskConfig) TfReplicationTask {
	_init_.Initialize()

	if err := validateNewTfReplicationTaskParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfReplicationTask{}

	_jsii_.Create(
		"@cdktn/aws-dms.TfReplicationTask",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task aws_dms_replication_task} Resource.
// Experimental.
func NewTfReplicationTask_Override(t TfReplicationTask, scope constructs.Construct, id *string, config *TfReplicationTaskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.TfReplicationTask",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetCdcStartPosition(val *string) {
	if err := j.validateSetCdcStartPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcStartPosition",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetCdcStartTime(val *string) {
	if err := j.validateSetCdcStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcStartTime",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetMigrationType(val *string) {
	if err := j.validateSetMigrationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"migrationType",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetReplicationInstanceArn(val *string) {
	if err := j.validateSetReplicationInstanceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationInstanceArn",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetReplicationTaskId(val *string) {
	if err := j.validateSetReplicationTaskIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationTaskId",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetReplicationTaskSettings(val *string) {
	if err := j.validateSetReplicationTaskSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationTaskSettings",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetResourceIdentifier(val *string) {
	if err := j.validateSetResourceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetSourceEndpointArn(val *string) {
	if err := j.validateSetSourceEndpointArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceEndpointArn",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetStartReplicationTask(val interface{}) {
	if err := j.validateSetStartReplicationTaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startReplicationTask",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetTableMappings(val *string) {
	if err := j.validateSetTableMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableMappings",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfReplicationTask)SetTargetEndpointArn(val *string) {
	if err := j.validateSetTargetEndpointArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetEndpointArn",
		val,
	)
}

// Generates CDKTN code for importing a TfReplicationTask resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfReplicationTask_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfReplicationTask_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-dms.TfReplicationTask",
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
func TfReplicationTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfReplicationTask_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dms.TfReplicationTask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfReplicationTask_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfReplicationTask_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dms.TfReplicationTask",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfReplicationTask_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfReplicationTask_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dms.TfReplicationTask",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfReplicationTask_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-dms.TfReplicationTask",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfReplicationTask) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfReplicationTask) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfReplicationTask) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfReplicationTask) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReplicationTask) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfReplicationTask) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfReplicationTask) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfReplicationTask) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfReplicationTask) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfReplicationTask) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfReplicationTask) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfReplicationTask) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationTask) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfReplicationTask) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReplicationTask) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfReplicationTask) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfReplicationTask) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfReplicationTask) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfReplicationTask) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfReplicationTask) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfReplicationTask) ResetCdcStartPosition() {
	_jsii_.InvokeVoid(
		t,
		"resetCdcStartPosition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationTask) ResetCdcStartTime() {
	_jsii_.InvokeVoid(
		t,
		"resetCdcStartTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationTask) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationTask) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationTask) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationTask) ResetReplicationTaskSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetReplicationTaskSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationTask) ResetResourceIdentifier() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceIdentifier",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationTask) ResetStartReplicationTask() {
	_jsii_.InvokeVoid(
		t,
		"resetStartReplicationTask",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationTask) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationTask) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfReplicationTask) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationTask) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationTask) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationTask) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationTask) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationTask) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationTask) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

