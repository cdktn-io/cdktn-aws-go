package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/glue/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/glue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler aws_glue_crawler}.
// Experimental.
type AwsCrawler interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	CatalogTarget() AwsCrawler_CatalogTargetPropertyList
	// Experimental.
	CatalogTargetInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Classifiers() *[]*string
	// Experimental.
	SetClassifiers(val *[]*string)
	// Experimental.
	ClassifiersInput() *[]*string
	// Experimental.
	Configuration() *string
	// Experimental.
	SetConfiguration(val *string)
	// Experimental.
	ConfigurationInput() *string
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
	DatabaseName() *string
	// Experimental.
	SetDatabaseName(val *string)
	// Experimental.
	DatabaseNameInput() *string
	// Experimental.
	DeltaTarget() AwsCrawler_DeltaTargetPropertyList
	// Experimental.
	DeltaTargetInput() interface{}
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
	DynamodbTarget() AwsCrawler_DynamodbTargetPropertyList
	// Experimental.
	DynamodbTargetInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HudiTarget() AwsCrawler_HudiTargetPropertyList
	// Experimental.
	HudiTargetInput() interface{}
	// Experimental.
	IcebergTarget() AwsCrawler_IcebergTargetPropertyList
	// Experimental.
	IcebergTargetInput() interface{}
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	JdbcTarget() AwsCrawler_JdbcTargetPropertyList
	// Experimental.
	JdbcTargetInput() interface{}
	// Experimental.
	LakeFormationConfiguration() AwsCrawler_LakeFormationConfigurationPropertyOutputReference
	// Experimental.
	LakeFormationConfigurationInput() *AwsCrawler_LakeFormationConfigurationProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LineageConfiguration() AwsCrawler_LineageConfigurationPropertyOutputReference
	// Experimental.
	LineageConfigurationInput() *AwsCrawler_LineageConfigurationProperty
	// Experimental.
	MongodbTarget() AwsCrawler_MongodbTargetPropertyList
	// Experimental.
	MongodbTargetInput() interface{}
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
	RecrawlPolicy() AwsCrawler_RecrawlPolicyPropertyOutputReference
	// Experimental.
	RecrawlPolicyInput() *AwsCrawler_RecrawlPolicyProperty
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	Role() *string
	// Experimental.
	SetRole(val *string)
	// Experimental.
	RoleInput() *string
	// Experimental.
	S3Target() AwsCrawler_S3TargetPropertyList
	// Experimental.
	S3TargetInput() interface{}
	// Experimental.
	Schedule() *string
	// Experimental.
	SetSchedule(val *string)
	// Experimental.
	ScheduleInput() *string
	// Experimental.
	SchemaChangePolicy() AwsCrawler_SchemaChangePolicyPropertyOutputReference
	// Experimental.
	SchemaChangePolicyInput() *AwsCrawler_SchemaChangePolicyProperty
	// Experimental.
	SecurityConfiguration() *string
	// Experimental.
	SetSecurityConfiguration(val *string)
	// Experimental.
	SecurityConfigurationInput() *string
	// Experimental.
	TablePrefix() *string
	// Experimental.
	SetTablePrefix(val *string)
	// Experimental.
	TablePrefixInput() *string
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
	PutCatalogTarget(value interface{})
	// Experimental.
	PutDeltaTarget(value interface{})
	// Experimental.
	PutDynamodbTarget(value interface{})
	// Experimental.
	PutHudiTarget(value interface{})
	// Experimental.
	PutIcebergTarget(value interface{})
	// Experimental.
	PutJdbcTarget(value interface{})
	// Experimental.
	PutLakeFormationConfiguration(value *AwsCrawler_LakeFormationConfigurationProperty)
	// Experimental.
	PutLineageConfiguration(value *AwsCrawler_LineageConfigurationProperty)
	// Experimental.
	PutMongodbTarget(value interface{})
	// Experimental.
	PutRecrawlPolicy(value *AwsCrawler_RecrawlPolicyProperty)
	// Experimental.
	PutS3Target(value interface{})
	// Experimental.
	PutSchemaChangePolicy(value *AwsCrawler_SchemaChangePolicyProperty)
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
	ResetCatalogTarget()
	// Experimental.
	ResetClassifiers()
	// Experimental.
	ResetConfiguration()
	// Experimental.
	ResetDeltaTarget()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetDynamodbTarget()
	// Experimental.
	ResetHudiTarget()
	// Experimental.
	ResetIcebergTarget()
	// Experimental.
	ResetId()
	// Experimental.
	ResetJdbcTarget()
	// Experimental.
	ResetLakeFormationConfiguration()
	// Experimental.
	ResetLineageConfiguration()
	// Experimental.
	ResetMongodbTarget()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRecrawlPolicy()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetS3Target()
	// Experimental.
	ResetSchedule()
	// Experimental.
	ResetSchemaChangePolicy()
	// Experimental.
	ResetSecurityConfiguration()
	// Experimental.
	ResetTablePrefix()
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

// The jsii proxy struct for AwsCrawler
type jsiiProxy_AwsCrawler struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsCrawler) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) CatalogTarget() AwsCrawler_CatalogTargetPropertyList {
	var returns AwsCrawler_CatalogTargetPropertyList
	_jsii_.Get(
		j,
		"catalogTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) CatalogTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) Classifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) ClassifiersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) Configuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) ConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) DeltaTarget() AwsCrawler_DeltaTargetPropertyList {
	var returns AwsCrawler_DeltaTargetPropertyList
	_jsii_.Get(
		j,
		"deltaTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) DeltaTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deltaTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) DynamodbTarget() AwsCrawler_DynamodbTargetPropertyList {
	var returns AwsCrawler_DynamodbTargetPropertyList
	_jsii_.Get(
		j,
		"dynamodbTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) DynamodbTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamodbTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) HudiTarget() AwsCrawler_HudiTargetPropertyList {
	var returns AwsCrawler_HudiTargetPropertyList
	_jsii_.Get(
		j,
		"hudiTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) HudiTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hudiTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) IcebergTarget() AwsCrawler_IcebergTargetPropertyList {
	var returns AwsCrawler_IcebergTargetPropertyList
	_jsii_.Get(
		j,
		"icebergTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) IcebergTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"icebergTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) JdbcTarget() AwsCrawler_JdbcTargetPropertyList {
	var returns AwsCrawler_JdbcTargetPropertyList
	_jsii_.Get(
		j,
		"jdbcTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) JdbcTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) LakeFormationConfiguration() AwsCrawler_LakeFormationConfigurationPropertyOutputReference {
	var returns AwsCrawler_LakeFormationConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"lakeFormationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) LakeFormationConfigurationInput() *AwsCrawler_LakeFormationConfigurationProperty {
	var returns *AwsCrawler_LakeFormationConfigurationProperty
	_jsii_.Get(
		j,
		"lakeFormationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) LineageConfiguration() AwsCrawler_LineageConfigurationPropertyOutputReference {
	var returns AwsCrawler_LineageConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"lineageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) LineageConfigurationInput() *AwsCrawler_LineageConfigurationProperty {
	var returns *AwsCrawler_LineageConfigurationProperty
	_jsii_.Get(
		j,
		"lineageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) MongodbTarget() AwsCrawler_MongodbTargetPropertyList {
	var returns AwsCrawler_MongodbTargetPropertyList
	_jsii_.Get(
		j,
		"mongodbTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) MongodbTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongodbTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) RecrawlPolicy() AwsCrawler_RecrawlPolicyPropertyOutputReference {
	var returns AwsCrawler_RecrawlPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"recrawlPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) RecrawlPolicyInput() *AwsCrawler_RecrawlPolicyProperty {
	var returns *AwsCrawler_RecrawlPolicyProperty
	_jsii_.Get(
		j,
		"recrawlPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) S3Target() AwsCrawler_S3TargetPropertyList {
	var returns AwsCrawler_S3TargetPropertyList
	_jsii_.Get(
		j,
		"s3Target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) S3TargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3TargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) SchemaChangePolicy() AwsCrawler_SchemaChangePolicyPropertyOutputReference {
	var returns AwsCrawler_SchemaChangePolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"schemaChangePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) SchemaChangePolicyInput() *AwsCrawler_SchemaChangePolicyProperty {
	var returns *AwsCrawler_SchemaChangePolicyProperty
	_jsii_.Get(
		j,
		"schemaChangePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) SecurityConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) TablePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) TablePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCrawler) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler aws_glue_crawler} Resource.
// Experimental.
func NewAwsCrawler(scope constructs.Construct, id *string, config *AwsCrawlerConfig) AwsCrawler {
	_init_.Initialize()

	if err := validateNewAwsCrawlerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCrawler{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsCrawler",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler aws_glue_crawler} Resource.
// Experimental.
func NewAwsCrawler_Override(a AwsCrawler, scope constructs.Construct, id *string, config *AwsCrawlerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsCrawler",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsCrawler)SetClassifiers(val *[]*string) {
	if err := j.validateSetClassifiersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classifiers",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetConfiguration(val *string) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetSchedule(val *string) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetSecurityConfiguration(val *string) {
	if err := j.validateSetSecurityConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetTablePrefix(val *string) {
	if err := j.validateSetTablePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tablePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsCrawler)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a AwsCrawler resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsCrawler_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsCrawler_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-glue.AwsCrawler",
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
func AwsCrawler_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCrawler_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-glue.AwsCrawler",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCrawler_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCrawler_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-glue.AwsCrawler",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsCrawler_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsCrawler_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-glue.AwsCrawler",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsCrawler_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-glue.AwsCrawler",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsCrawler) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsCrawler) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsCrawler) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCrawler) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCrawler) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCrawler) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCrawler) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCrawler) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCrawler) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCrawler) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCrawler) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCrawler) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCrawler) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsCrawler) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCrawler) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsCrawler) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCrawler) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsCrawler) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsCrawler) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsCrawler) PutCatalogTarget(value interface{}) {
	if err := a.validatePutCatalogTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCatalogTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCrawler) PutDeltaTarget(value interface{}) {
	if err := a.validatePutDeltaTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeltaTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCrawler) PutDynamodbTarget(value interface{}) {
	if err := a.validatePutDynamodbTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynamodbTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCrawler) PutHudiTarget(value interface{}) {
	if err := a.validatePutHudiTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHudiTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCrawler) PutIcebergTarget(value interface{}) {
	if err := a.validatePutIcebergTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIcebergTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCrawler) PutJdbcTarget(value interface{}) {
	if err := a.validatePutJdbcTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJdbcTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCrawler) PutLakeFormationConfiguration(value *AwsCrawler_LakeFormationConfigurationProperty) {
	if err := a.validatePutLakeFormationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLakeFormationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCrawler) PutLineageConfiguration(value *AwsCrawler_LineageConfigurationProperty) {
	if err := a.validatePutLineageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLineageConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCrawler) PutMongodbTarget(value interface{}) {
	if err := a.validatePutMongodbTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMongodbTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCrawler) PutRecrawlPolicy(value *AwsCrawler_RecrawlPolicyProperty) {
	if err := a.validatePutRecrawlPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRecrawlPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCrawler) PutS3Target(value interface{}) {
	if err := a.validatePutS3TargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Target",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCrawler) PutSchemaChangePolicy(value *AwsCrawler_SchemaChangePolicyProperty) {
	if err := a.validatePutSchemaChangePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchemaChangePolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCrawler) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsCrawler) ResetCatalogTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetCatalogTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetClassifiers() {
	_jsii_.InvokeVoid(
		a,
		"resetClassifiers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetDeltaTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetDeltaTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetDynamodbTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamodbTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetHudiTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetHudiTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetIcebergTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetIcebergTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetJdbcTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetJdbcTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetLakeFormationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLakeFormationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetLineageConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLineageConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetMongodbTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetMongodbTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetRecrawlPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetRecrawlPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetS3Target() {
	_jsii_.InvokeVoid(
		a,
		"resetS3Target",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetSchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetSchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetSchemaChangePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetSchemaChangePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetSecurityConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetTablePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetTablePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCrawler) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCrawler) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCrawler) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCrawler) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCrawler) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCrawler) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCrawler) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

