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
type AwsGlueCrawler interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	CatalogTarget() AwsGlueCrawler_CatalogTargetPropertyList
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
	DeltaTarget() AwsGlueCrawler_DeltaTargetPropertyList
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
	DynamodbTarget() AwsGlueCrawler_DynamodbTargetPropertyList
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
	HudiTarget() AwsGlueCrawler_HudiTargetPropertyList
	// Experimental.
	HudiTargetInput() interface{}
	// Experimental.
	IcebergTarget() AwsGlueCrawler_IcebergTargetPropertyList
	// Experimental.
	IcebergTargetInput() interface{}
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	JdbcTarget() AwsGlueCrawler_JdbcTargetPropertyList
	// Experimental.
	JdbcTargetInput() interface{}
	// Experimental.
	LakeFormationConfiguration() AwsGlueCrawler_LakeFormationConfigurationPropertyOutputReference
	// Experimental.
	LakeFormationConfigurationInput() *AwsGlueCrawler_LakeFormationConfigurationProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LineageConfiguration() AwsGlueCrawler_LineageConfigurationPropertyOutputReference
	// Experimental.
	LineageConfigurationInput() *AwsGlueCrawler_LineageConfigurationProperty
	// Experimental.
	MongodbTarget() AwsGlueCrawler_MongodbTargetPropertyList
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
	RecrawlPolicy() AwsGlueCrawler_RecrawlPolicyPropertyOutputReference
	// Experimental.
	RecrawlPolicyInput() *AwsGlueCrawler_RecrawlPolicyProperty
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
	S3Target() AwsGlueCrawler_S3TargetPropertyList
	// Experimental.
	S3TargetInput() interface{}
	// Experimental.
	Schedule() *string
	// Experimental.
	SetSchedule(val *string)
	// Experimental.
	ScheduleInput() *string
	// Experimental.
	SchemaChangePolicy() AwsGlueCrawler_SchemaChangePolicyPropertyOutputReference
	// Experimental.
	SchemaChangePolicyInput() *AwsGlueCrawler_SchemaChangePolicyProperty
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
	PutLakeFormationConfiguration(value *AwsGlueCrawler_LakeFormationConfigurationProperty)
	// Experimental.
	PutLineageConfiguration(value *AwsGlueCrawler_LineageConfigurationProperty)
	// Experimental.
	PutMongodbTarget(value interface{})
	// Experimental.
	PutRecrawlPolicy(value *AwsGlueCrawler_RecrawlPolicyProperty)
	// Experimental.
	PutS3Target(value interface{})
	// Experimental.
	PutSchemaChangePolicy(value *AwsGlueCrawler_SchemaChangePolicyProperty)
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

// The jsii proxy struct for AwsGlueCrawler
type jsiiProxy_AwsGlueCrawler struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsGlueCrawler) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) CatalogTarget() AwsGlueCrawler_CatalogTargetPropertyList {
	var returns AwsGlueCrawler_CatalogTargetPropertyList
	_jsii_.Get(
		j,
		"catalogTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) CatalogTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) Classifiers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classifiers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) ClassifiersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"classifiersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) Configuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) ConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) DeltaTarget() AwsGlueCrawler_DeltaTargetPropertyList {
	var returns AwsGlueCrawler_DeltaTargetPropertyList
	_jsii_.Get(
		j,
		"deltaTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) DeltaTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deltaTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) DynamodbTarget() AwsGlueCrawler_DynamodbTargetPropertyList {
	var returns AwsGlueCrawler_DynamodbTargetPropertyList
	_jsii_.Get(
		j,
		"dynamodbTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) DynamodbTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamodbTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) HudiTarget() AwsGlueCrawler_HudiTargetPropertyList {
	var returns AwsGlueCrawler_HudiTargetPropertyList
	_jsii_.Get(
		j,
		"hudiTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) HudiTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hudiTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) IcebergTarget() AwsGlueCrawler_IcebergTargetPropertyList {
	var returns AwsGlueCrawler_IcebergTargetPropertyList
	_jsii_.Get(
		j,
		"icebergTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) IcebergTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"icebergTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) JdbcTarget() AwsGlueCrawler_JdbcTargetPropertyList {
	var returns AwsGlueCrawler_JdbcTargetPropertyList
	_jsii_.Get(
		j,
		"jdbcTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) JdbcTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) LakeFormationConfiguration() AwsGlueCrawler_LakeFormationConfigurationPropertyOutputReference {
	var returns AwsGlueCrawler_LakeFormationConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"lakeFormationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) LakeFormationConfigurationInput() *AwsGlueCrawler_LakeFormationConfigurationProperty {
	var returns *AwsGlueCrawler_LakeFormationConfigurationProperty
	_jsii_.Get(
		j,
		"lakeFormationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) LineageConfiguration() AwsGlueCrawler_LineageConfigurationPropertyOutputReference {
	var returns AwsGlueCrawler_LineageConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"lineageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) LineageConfigurationInput() *AwsGlueCrawler_LineageConfigurationProperty {
	var returns *AwsGlueCrawler_LineageConfigurationProperty
	_jsii_.Get(
		j,
		"lineageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) MongodbTarget() AwsGlueCrawler_MongodbTargetPropertyList {
	var returns AwsGlueCrawler_MongodbTargetPropertyList
	_jsii_.Get(
		j,
		"mongodbTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) MongodbTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongodbTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) RecrawlPolicy() AwsGlueCrawler_RecrawlPolicyPropertyOutputReference {
	var returns AwsGlueCrawler_RecrawlPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"recrawlPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) RecrawlPolicyInput() *AwsGlueCrawler_RecrawlPolicyProperty {
	var returns *AwsGlueCrawler_RecrawlPolicyProperty
	_jsii_.Get(
		j,
		"recrawlPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) S3Target() AwsGlueCrawler_S3TargetPropertyList {
	var returns AwsGlueCrawler_S3TargetPropertyList
	_jsii_.Get(
		j,
		"s3Target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) S3TargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3TargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) SchemaChangePolicy() AwsGlueCrawler_SchemaChangePolicyPropertyOutputReference {
	var returns AwsGlueCrawler_SchemaChangePolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"schemaChangePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) SchemaChangePolicyInput() *AwsGlueCrawler_SchemaChangePolicyProperty {
	var returns *AwsGlueCrawler_SchemaChangePolicyProperty
	_jsii_.Get(
		j,
		"schemaChangePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) SecurityConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) TablePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) TablePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tablePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGlueCrawler) TerraformResourceType() *string {
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
func NewAwsGlueCrawler(scope constructs.Construct, id *string, config *AwsGlueCrawlerConfig) AwsGlueCrawler {
	_init_.Initialize()

	if err := validateNewAwsGlueCrawlerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGlueCrawler{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsGlueCrawler",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler aws_glue_crawler} Resource.
// Experimental.
func NewAwsGlueCrawler_Override(a AwsGlueCrawler, scope constructs.Construct, id *string, config *AwsGlueCrawlerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsGlueCrawler",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetClassifiers(val *[]*string) {
	if err := j.validateSetClassifiersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"classifiers",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetConfiguration(val *string) {
	if err := j.validateSetConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuration",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetSchedule(val *string) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetSecurityConfiguration(val *string) {
	if err := j.validateSetSecurityConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetTablePrefix(val *string) {
	if err := j.validateSetTablePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tablePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsGlueCrawler)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a AwsGlueCrawler resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsGlueCrawler_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsGlueCrawler_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-glue.AwsGlueCrawler",
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
func AwsGlueCrawler_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsGlueCrawler_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-glue.AwsGlueCrawler",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsGlueCrawler_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsGlueCrawler_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-glue.AwsGlueCrawler",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsGlueCrawler_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsGlueCrawler_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-glue.AwsGlueCrawler",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsGlueCrawler_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-glue.AwsGlueCrawler",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsGlueCrawler) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGlueCrawler) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGlueCrawler) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGlueCrawler) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGlueCrawler) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGlueCrawler) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGlueCrawler) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGlueCrawler) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGlueCrawler) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGlueCrawler) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueCrawler) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGlueCrawler) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsGlueCrawler) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) PutCatalogTarget(value interface{}) {
	if err := a.validatePutCatalogTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCatalogTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) PutDeltaTarget(value interface{}) {
	if err := a.validatePutDeltaTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeltaTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) PutDynamodbTarget(value interface{}) {
	if err := a.validatePutDynamodbTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynamodbTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) PutHudiTarget(value interface{}) {
	if err := a.validatePutHudiTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHudiTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) PutIcebergTarget(value interface{}) {
	if err := a.validatePutIcebergTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIcebergTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) PutJdbcTarget(value interface{}) {
	if err := a.validatePutJdbcTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJdbcTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) PutLakeFormationConfiguration(value *AwsGlueCrawler_LakeFormationConfigurationProperty) {
	if err := a.validatePutLakeFormationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLakeFormationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) PutLineageConfiguration(value *AwsGlueCrawler_LineageConfigurationProperty) {
	if err := a.validatePutLineageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLineageConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) PutMongodbTarget(value interface{}) {
	if err := a.validatePutMongodbTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMongodbTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) PutRecrawlPolicy(value *AwsGlueCrawler_RecrawlPolicyProperty) {
	if err := a.validatePutRecrawlPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRecrawlPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) PutS3Target(value interface{}) {
	if err := a.validatePutS3TargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Target",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) PutSchemaChangePolicy(value *AwsGlueCrawler_SchemaChangePolicyProperty) {
	if err := a.validatePutSchemaChangePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchemaChangePolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetCatalogTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetCatalogTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetClassifiers() {
	_jsii_.InvokeVoid(
		a,
		"resetClassifiers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetDeltaTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetDeltaTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetDynamodbTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamodbTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetHudiTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetHudiTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetIcebergTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetIcebergTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetJdbcTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetJdbcTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetLakeFormationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLakeFormationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetLineageConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLineageConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetMongodbTarget() {
	_jsii_.InvokeVoid(
		a,
		"resetMongodbTarget",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetRecrawlPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetRecrawlPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetS3Target() {
	_jsii_.InvokeVoid(
		a,
		"resetS3Target",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetSchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetSchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetSchemaChangePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetSchemaChangePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetSecurityConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetTablePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetTablePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGlueCrawler) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueCrawler) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueCrawler) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueCrawler) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueCrawler) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueCrawler) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGlueCrawler) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

