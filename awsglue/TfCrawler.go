package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler aws_glue_crawler}.
// Experimental.
type TfCrawler interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	CatalogTarget() TfCrawler_CatalogTargetPropertyList
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
	DeltaTarget() TfCrawler_DeltaTargetPropertyList
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
	DynamodbTarget() TfCrawler_DynamodbTargetPropertyList
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
	HudiTarget() TfCrawler_HudiTargetPropertyList
	// Experimental.
	HudiTargetInput() interface{}
	// Experimental.
	IcebergTarget() TfCrawler_IcebergTargetPropertyList
	// Experimental.
	IcebergTargetInput() interface{}
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	JdbcTarget() TfCrawler_JdbcTargetPropertyList
	// Experimental.
	JdbcTargetInput() interface{}
	// Experimental.
	LakeFormationConfiguration() TfCrawler_LakeFormationConfigurationPropertyOutputReference
	// Experimental.
	LakeFormationConfigurationInput() *TfCrawler_LakeFormationConfigurationProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LineageConfiguration() TfCrawler_LineageConfigurationPropertyOutputReference
	// Experimental.
	LineageConfigurationInput() *TfCrawler_LineageConfigurationProperty
	// Experimental.
	MongodbTarget() TfCrawler_MongodbTargetPropertyList
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
	RecrawlPolicy() TfCrawler_RecrawlPolicyPropertyOutputReference
	// Experimental.
	RecrawlPolicyInput() *TfCrawler_RecrawlPolicyProperty
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
	S3Target() TfCrawler_S3TargetPropertyList
	// Experimental.
	S3TargetInput() interface{}
	// Experimental.
	Schedule() *string
	// Experimental.
	SetSchedule(val *string)
	// Experimental.
	ScheduleInput() *string
	// Experimental.
	SchemaChangePolicy() TfCrawler_SchemaChangePolicyPropertyOutputReference
	// Experimental.
	SchemaChangePolicyInput() *TfCrawler_SchemaChangePolicyProperty
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
	PutLakeFormationConfiguration(value *TfCrawler_LakeFormationConfigurationProperty)
	// Experimental.
	PutLineageConfiguration(value *TfCrawler_LineageConfigurationProperty)
	// Experimental.
	PutMongodbTarget(value interface{})
	// Experimental.
	PutRecrawlPolicy(value *TfCrawler_RecrawlPolicyProperty)
	// Experimental.
	PutS3Target(value interface{})
	// Experimental.
	PutSchemaChangePolicy(value *TfCrawler_SchemaChangePolicyProperty)
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

// The jsii proxy struct for TfCrawler
type jsiiProxy_TfCrawler struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfCrawler) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) CatalogTarget() TfCrawler_CatalogTargetPropertyList {
	var returns TfCrawler_CatalogTargetPropertyList
	_jsii_.Get(
		j,
		"catalogTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) CatalogTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) Classifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) ClassifiersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) Configuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) ConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) DeltaTarget() TfCrawler_DeltaTargetPropertyList {
	var returns TfCrawler_DeltaTargetPropertyList
	_jsii_.Get(
		j,
		"deltaTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) DeltaTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deltaTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) DynamodbTarget() TfCrawler_DynamodbTargetPropertyList {
	var returns TfCrawler_DynamodbTargetPropertyList
	_jsii_.Get(
		j,
		"dynamodbTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) DynamodbTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamodbTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) HudiTarget() TfCrawler_HudiTargetPropertyList {
	var returns TfCrawler_HudiTargetPropertyList
	_jsii_.Get(
		j,
		"hudiTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) HudiTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hudiTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) IcebergTarget() TfCrawler_IcebergTargetPropertyList {
	var returns TfCrawler_IcebergTargetPropertyList
	_jsii_.Get(
		j,
		"icebergTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) IcebergTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"icebergTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) JdbcTarget() TfCrawler_JdbcTargetPropertyList {
	var returns TfCrawler_JdbcTargetPropertyList
	_jsii_.Get(
		j,
		"jdbcTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) JdbcTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) LakeFormationConfiguration() TfCrawler_LakeFormationConfigurationPropertyOutputReference {
	var returns TfCrawler_LakeFormationConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"lakeFormationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) LakeFormationConfigurationInput() *TfCrawler_LakeFormationConfigurationProperty {
	var returns *TfCrawler_LakeFormationConfigurationProperty
	_jsii_.Get(
		j,
		"lakeFormationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) LineageConfiguration() TfCrawler_LineageConfigurationPropertyOutputReference {
	var returns TfCrawler_LineageConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"lineageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) LineageConfigurationInput() *TfCrawler_LineageConfigurationProperty {
	var returns *TfCrawler_LineageConfigurationProperty
	_jsii_.Get(
		j,
		"lineageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) MongodbTarget() TfCrawler_MongodbTargetPropertyList {
	var returns TfCrawler_MongodbTargetPropertyList
	_jsii_.Get(
		j,
		"mongodbTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) MongodbTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongodbTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) RecrawlPolicy() TfCrawler_RecrawlPolicyPropertyOutputReference {
	var returns TfCrawler_RecrawlPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"recrawlPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) RecrawlPolicyInput() *TfCrawler_RecrawlPolicyProperty {
	var returns *TfCrawler_RecrawlPolicyProperty
	_jsii_.Get(
		j,
		"recrawlPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) S3Target() TfCrawler_S3TargetPropertyList {
	var returns TfCrawler_S3TargetPropertyList
	_jsii_.Get(
		j,
		"s3Target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) S3TargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3TargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) SchemaChangePolicy() TfCrawler_SchemaChangePolicyPropertyOutputReference {
	var returns TfCrawler_SchemaChangePolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"schemaChangePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) SchemaChangePolicyInput() *TfCrawler_SchemaChangePolicyProperty {
	var returns *TfCrawler_SchemaChangePolicyProperty
	_jsii_.Get(
		j,
		"schemaChangePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) SecurityConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) TablePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) TablePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCrawler) TerraformResourceType() *string {
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
func NewTfCrawler(scope constructs.Construct, id *string, config *TfCrawlerConfig) TfCrawler {
	_init_.Initialize()

	if err := validateNewTfCrawlerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCrawler{}

	_jsii_.Create(
		"@cdktn/aws-glue.TfCrawler",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler aws_glue_crawler} Resource.
// Experimental.
func NewTfCrawler_Override(t TfCrawler, scope constructs.Construct, id *string, config *TfCrawlerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.TfCrawler",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfCrawler)SetClassifiers(val *[]*string) {
	if err := j.validateSetClassifiersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classifiers",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetConfiguration(val *string) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetSchedule(val *string) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetSecurityConfiguration(val *string) {
	if err := j.validateSetSecurityConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetTablePrefix(val *string) {
	if err := j.validateSetTablePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tablePrefix",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfCrawler)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a TfCrawler resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfCrawler_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfCrawler_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-glue.TfCrawler",
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
func TfCrawler_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCrawler_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-glue.TfCrawler",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCrawler_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCrawler_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-glue.TfCrawler",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCrawler_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCrawler_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-glue.TfCrawler",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfCrawler_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-glue.TfCrawler",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfCrawler) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfCrawler) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfCrawler) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCrawler) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCrawler) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCrawler) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCrawler) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCrawler) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCrawler) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCrawler) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCrawler) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCrawler) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCrawler) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfCrawler) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCrawler) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfCrawler) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCrawler) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfCrawler) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCrawler) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfCrawler) PutCatalogTarget(value interface{}) {
	if err := t.validatePutCatalogTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCatalogTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCrawler) PutDeltaTarget(value interface{}) {
	if err := t.validatePutDeltaTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeltaTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCrawler) PutDynamodbTarget(value interface{}) {
	if err := t.validatePutDynamodbTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDynamodbTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCrawler) PutHudiTarget(value interface{}) {
	if err := t.validatePutHudiTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHudiTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCrawler) PutIcebergTarget(value interface{}) {
	if err := t.validatePutIcebergTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIcebergTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCrawler) PutJdbcTarget(value interface{}) {
	if err := t.validatePutJdbcTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJdbcTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCrawler) PutLakeFormationConfiguration(value *TfCrawler_LakeFormationConfigurationProperty) {
	if err := t.validatePutLakeFormationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLakeFormationConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCrawler) PutLineageConfiguration(value *TfCrawler_LineageConfigurationProperty) {
	if err := t.validatePutLineageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLineageConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCrawler) PutMongodbTarget(value interface{}) {
	if err := t.validatePutMongodbTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMongodbTarget",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCrawler) PutRecrawlPolicy(value *TfCrawler_RecrawlPolicyProperty) {
	if err := t.validatePutRecrawlPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRecrawlPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCrawler) PutS3Target(value interface{}) {
	if err := t.validatePutS3TargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Target",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCrawler) PutSchemaChangePolicy(value *TfCrawler_SchemaChangePolicyProperty) {
	if err := t.validatePutSchemaChangePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSchemaChangePolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCrawler) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfCrawler) ResetCatalogTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetCatalogTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetClassifiers() {
	_jsii_.InvokeVoid(
		t,
		"resetClassifiers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetDeltaTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetDeltaTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetDynamodbTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetDynamodbTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetHudiTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetHudiTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetIcebergTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetIcebergTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetJdbcTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetJdbcTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetLakeFormationConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetLakeFormationConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetLineageConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetLineageConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetMongodbTarget() {
	_jsii_.InvokeVoid(
		t,
		"resetMongodbTarget",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetRecrawlPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetRecrawlPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetS3Target() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Target",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetSchedule() {
	_jsii_.InvokeVoid(
		t,
		"resetSchedule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetSchemaChangePolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetSchemaChangePolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetSecurityConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetTablePrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetTablePrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCrawler) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCrawler) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCrawler) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCrawler) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCrawler) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCrawler) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCrawler) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

