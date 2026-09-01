package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdms/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsdms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint aws_dms_endpoint}.
// Experimental.
type AwsDmsEndpoint interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CertificateArn() *string
	// Experimental.
	SetCertificateArn(val *string)
	// Experimental.
	CertificateArnInput() *string
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
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ElasticsearchSettings() AwsDmsEndpoint_ElasticsearchSettingsPropertyOutputReference
	// Experimental.
	ElasticsearchSettingsInput() *AwsDmsEndpoint_ElasticsearchSettingsProperty
	// Experimental.
	EndpointArn() *string
	// Experimental.
	EndpointId() *string
	// Experimental.
	SetEndpointId(val *string)
	// Experimental.
	EndpointIdInput() *string
	// Experimental.
	EndpointType() *string
	// Experimental.
	SetEndpointType(val *string)
	// Experimental.
	EndpointTypeInput() *string
	// Experimental.
	EngineName() *string
	// Experimental.
	SetEngineName(val *string)
	// Experimental.
	EngineNameInput() *string
	// Experimental.
	ExtraConnectionAttributes() *string
	// Experimental.
	SetExtraConnectionAttributes(val *string)
	// Experimental.
	ExtraConnectionAttributesInput() *string
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
	KafkaSettings() AwsDmsEndpoint_KafkaSettingsPropertyOutputReference
	// Experimental.
	KafkaSettingsInput() *AwsDmsEndpoint_KafkaSettingsProperty
	// Experimental.
	KinesisSettings() AwsDmsEndpoint_KinesisSettingsPropertyOutputReference
	// Experimental.
	KinesisSettingsInput() *AwsDmsEndpoint_KinesisSettingsProperty
	// Experimental.
	KmsKeyArn() *string
	// Experimental.
	SetKmsKeyArn(val *string)
	// Experimental.
	KmsKeyArnInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MongodbSettings() AwsDmsEndpoint_MongodbSettingsPropertyOutputReference
	// Experimental.
	MongodbSettingsInput() *AwsDmsEndpoint_MongodbSettingsProperty
	// Experimental.
	MysqlSettings() AwsDmsEndpoint_MysqlSettingsPropertyOutputReference
	// Experimental.
	MysqlSettingsInput() *AwsDmsEndpoint_MysqlSettingsProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OracleSettings() AwsDmsEndpoint_OracleSettingsPropertyOutputReference
	// Experimental.
	OracleSettingsInput() *AwsDmsEndpoint_OracleSettingsProperty
	// Experimental.
	Password() *string
	// Experimental.
	SetPassword(val *string)
	// Experimental.
	PasswordInput() *string
	// Experimental.
	PauseReplicationTasks() interface{}
	// Experimental.
	SetPauseReplicationTasks(val interface{})
	// Experimental.
	PauseReplicationTasksInput() interface{}
	// Experimental.
	Port() *float64
	// Experimental.
	SetPort(val *float64)
	// Experimental.
	PortInput() *float64
	// Experimental.
	PostgresSettings() AwsDmsEndpoint_PostgresSettingsPropertyOutputReference
	// Experimental.
	PostgresSettingsInput() *AwsDmsEndpoint_PostgresSettingsProperty
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
	RedisSettings() AwsDmsEndpoint_RedisSettingsPropertyOutputReference
	// Experimental.
	RedisSettingsInput() *AwsDmsEndpoint_RedisSettingsProperty
	// Experimental.
	RedshiftSettings() AwsDmsEndpoint_RedshiftSettingsPropertyOutputReference
	// Experimental.
	RedshiftSettingsInput() *AwsDmsEndpoint_RedshiftSettingsProperty
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	SecretsManagerAccessRoleArn() *string
	// Experimental.
	SetSecretsManagerAccessRoleArn(val *string)
	// Experimental.
	SecretsManagerAccessRoleArnInput() *string
	// Experimental.
	SecretsManagerArn() *string
	// Experimental.
	SetSecretsManagerArn(val *string)
	// Experimental.
	SecretsManagerArnInput() *string
	// Experimental.
	ServerName() *string
	// Experimental.
	SetServerName(val *string)
	// Experimental.
	ServerNameInput() *string
	// Experimental.
	ServiceAccessRole() *string
	// Experimental.
	SetServiceAccessRole(val *string)
	// Experimental.
	ServiceAccessRoleInput() *string
	// Experimental.
	SslMode() *string
	// Experimental.
	SetSslMode(val *string)
	// Experimental.
	SslModeInput() *string
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
	Timeouts() AwsDmsEndpoint_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Username() *string
	// Experimental.
	SetUsername(val *string)
	// Experimental.
	UsernameInput() *string
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
	PutElasticsearchSettings(value *AwsDmsEndpoint_ElasticsearchSettingsProperty)
	// Experimental.
	PutKafkaSettings(value *AwsDmsEndpoint_KafkaSettingsProperty)
	// Experimental.
	PutKinesisSettings(value *AwsDmsEndpoint_KinesisSettingsProperty)
	// Experimental.
	PutMongodbSettings(value *AwsDmsEndpoint_MongodbSettingsProperty)
	// Experimental.
	PutMysqlSettings(value *AwsDmsEndpoint_MysqlSettingsProperty)
	// Experimental.
	PutOracleSettings(value *AwsDmsEndpoint_OracleSettingsProperty)
	// Experimental.
	PutPostgresSettings(value *AwsDmsEndpoint_PostgresSettingsProperty)
	// Experimental.
	PutRedisSettings(value *AwsDmsEndpoint_RedisSettingsProperty)
	// Experimental.
	PutRedshiftSettings(value *AwsDmsEndpoint_RedshiftSettingsProperty)
	// Experimental.
	PutTimeouts(value *AwsDmsEndpoint_TimeoutsProperty)
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
	ResetCertificateArn()
	// Experimental.
	ResetDatabaseName()
	// Experimental.
	ResetElasticsearchSettings()
	// Experimental.
	ResetExtraConnectionAttributes()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKafkaSettings()
	// Experimental.
	ResetKinesisSettings()
	// Experimental.
	ResetKmsKeyArn()
	// Experimental.
	ResetMongodbSettings()
	// Experimental.
	ResetMysqlSettings()
	// Experimental.
	ResetOracleSettings()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPassword()
	// Experimental.
	ResetPauseReplicationTasks()
	// Experimental.
	ResetPort()
	// Experimental.
	ResetPostgresSettings()
	// Experimental.
	ResetRedisSettings()
	// Experimental.
	ResetRedshiftSettings()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSecretsManagerAccessRoleArn()
	// Experimental.
	ResetSecretsManagerArn()
	// Experimental.
	ResetServerName()
	// Experimental.
	ResetServiceAccessRole()
	// Experimental.
	ResetSslMode()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetUsername()
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

// The jsii proxy struct for AwsDmsEndpoint
type jsiiProxy_AwsDmsEndpoint struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsDmsEndpoint) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) ElasticsearchSettings() AwsDmsEndpoint_ElasticsearchSettingsPropertyOutputReference {
	var returns AwsDmsEndpoint_ElasticsearchSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"elasticsearchSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) ElasticsearchSettingsInput() *AwsDmsEndpoint_ElasticsearchSettingsProperty {
	var returns *AwsDmsEndpoint_ElasticsearchSettingsProperty
	_jsii_.Get(
		j,
		"elasticsearchSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) EndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) EndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) EndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) EngineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) EngineNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) ExtraConnectionAttributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraConnectionAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) ExtraConnectionAttributesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraConnectionAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) KafkaSettings() AwsDmsEndpoint_KafkaSettingsPropertyOutputReference {
	var returns AwsDmsEndpoint_KafkaSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"kafkaSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) KafkaSettingsInput() *AwsDmsEndpoint_KafkaSettingsProperty {
	var returns *AwsDmsEndpoint_KafkaSettingsProperty
	_jsii_.Get(
		j,
		"kafkaSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) KinesisSettings() AwsDmsEndpoint_KinesisSettingsPropertyOutputReference {
	var returns AwsDmsEndpoint_KinesisSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) KinesisSettingsInput() *AwsDmsEndpoint_KinesisSettingsProperty {
	var returns *AwsDmsEndpoint_KinesisSettingsProperty
	_jsii_.Get(
		j,
		"kinesisSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) MongodbSettings() AwsDmsEndpoint_MongodbSettingsPropertyOutputReference {
	var returns AwsDmsEndpoint_MongodbSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"mongodbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) MongodbSettingsInput() *AwsDmsEndpoint_MongodbSettingsProperty {
	var returns *AwsDmsEndpoint_MongodbSettingsProperty
	_jsii_.Get(
		j,
		"mongodbSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) MysqlSettings() AwsDmsEndpoint_MysqlSettingsPropertyOutputReference {
	var returns AwsDmsEndpoint_MysqlSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"mysqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) MysqlSettingsInput() *AwsDmsEndpoint_MysqlSettingsProperty {
	var returns *AwsDmsEndpoint_MysqlSettingsProperty
	_jsii_.Get(
		j,
		"mysqlSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) OracleSettings() AwsDmsEndpoint_OracleSettingsPropertyOutputReference {
	var returns AwsDmsEndpoint_OracleSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"oracleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) OracleSettingsInput() *AwsDmsEndpoint_OracleSettingsProperty {
	var returns *AwsDmsEndpoint_OracleSettingsProperty
	_jsii_.Get(
		j,
		"oracleSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) PauseReplicationTasks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pauseReplicationTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) PauseReplicationTasksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pauseReplicationTasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) PostgresSettings() AwsDmsEndpoint_PostgresSettingsPropertyOutputReference {
	var returns AwsDmsEndpoint_PostgresSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"postgresSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) PostgresSettingsInput() *AwsDmsEndpoint_PostgresSettingsProperty {
	var returns *AwsDmsEndpoint_PostgresSettingsProperty
	_jsii_.Get(
		j,
		"postgresSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) RedisSettings() AwsDmsEndpoint_RedisSettingsPropertyOutputReference {
	var returns AwsDmsEndpoint_RedisSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"redisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) RedisSettingsInput() *AwsDmsEndpoint_RedisSettingsProperty {
	var returns *AwsDmsEndpoint_RedisSettingsProperty
	_jsii_.Get(
		j,
		"redisSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) RedshiftSettings() AwsDmsEndpoint_RedshiftSettingsPropertyOutputReference {
	var returns AwsDmsEndpoint_RedshiftSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"redshiftSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) RedshiftSettingsInput() *AwsDmsEndpoint_RedshiftSettingsProperty {
	var returns *AwsDmsEndpoint_RedshiftSettingsProperty
	_jsii_.Get(
		j,
		"redshiftSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) SecretsManagerAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) SecretsManagerAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) SecretsManagerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) SecretsManagerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) ServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) ServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) ServiceAccessRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) ServiceAccessRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) SslMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) SslModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) Timeouts() AwsDmsEndpoint_TimeoutsPropertyOutputReference {
	var returns AwsDmsEndpoint_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDmsEndpoint) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint aws_dms_endpoint} Resource.
// Experimental.
func NewAwsDmsEndpoint(scope constructs.Construct, id *string, config *AwsDmsEndpointConfig) AwsDmsEndpoint {
	_init_.Initialize()

	if err := validateNewAwsDmsEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDmsEndpoint{}

	_jsii_.Create(
		"@cdktn/aws-dms.AwsDmsEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint aws_dms_endpoint} Resource.
// Experimental.
func NewAwsDmsEndpoint_Override(a AwsDmsEndpoint, scope constructs.Construct, id *string, config *AwsDmsEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.AwsDmsEndpoint",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetCertificateArn(val *string) {
	if err := j.validateSetCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetEndpointId(val *string) {
	if err := j.validateSetEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointId",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetEndpointType(val *string) {
	if err := j.validateSetEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetEngineName(val *string) {
	if err := j.validateSetEngineNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineName",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetExtraConnectionAttributes(val *string) {
	if err := j.validateSetExtraConnectionAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraConnectionAttributes",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetPauseReplicationTasks(val interface{}) {
	if err := j.validateSetPauseReplicationTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pauseReplicationTasks",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetSecretsManagerAccessRoleArn(val *string) {
	if err := j.validateSetSecretsManagerAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetSecretsManagerArn(val *string) {
	if err := j.validateSetSecretsManagerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerArn",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetServerName(val *string) {
	if err := j.validateSetServerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverName",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetServiceAccessRole(val *string) {
	if err := j.validateSetServiceAccessRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRole",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetSslMode(val *string) {
	if err := j.validateSetSslModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslMode",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsDmsEndpoint)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Generates CDKTN code for importing a AwsDmsEndpoint resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsDmsEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsDmsEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-dms.AwsDmsEndpoint",
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
func AwsDmsEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDmsEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dms.AwsDmsEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsDmsEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDmsEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dms.AwsDmsEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsDmsEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDmsEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dms.AwsDmsEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsDmsEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-dms.AwsDmsEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsDmsEndpoint) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDmsEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDmsEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDmsEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDmsEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDmsEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDmsEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDmsEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDmsEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDmsEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDmsEndpoint) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDmsEndpoint) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsDmsEndpoint) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) PutElasticsearchSettings(value *AwsDmsEndpoint_ElasticsearchSettingsProperty) {
	if err := a.validatePutElasticsearchSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putElasticsearchSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) PutKafkaSettings(value *AwsDmsEndpoint_KafkaSettingsProperty) {
	if err := a.validatePutKafkaSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKafkaSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) PutKinesisSettings(value *AwsDmsEndpoint_KinesisSettingsProperty) {
	if err := a.validatePutKinesisSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) PutMongodbSettings(value *AwsDmsEndpoint_MongodbSettingsProperty) {
	if err := a.validatePutMongodbSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMongodbSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) PutMysqlSettings(value *AwsDmsEndpoint_MysqlSettingsProperty) {
	if err := a.validatePutMysqlSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMysqlSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) PutOracleSettings(value *AwsDmsEndpoint_OracleSettingsProperty) {
	if err := a.validatePutOracleSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOracleSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) PutPostgresSettings(value *AwsDmsEndpoint_PostgresSettingsProperty) {
	if err := a.validatePutPostgresSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPostgresSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) PutRedisSettings(value *AwsDmsEndpoint_RedisSettingsProperty) {
	if err := a.validatePutRedisSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedisSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) PutRedshiftSettings(value *AwsDmsEndpoint_RedshiftSettingsProperty) {
	if err := a.validatePutRedshiftSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshiftSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) PutTimeouts(value *AwsDmsEndpoint_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		a,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetElasticsearchSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetExtraConnectionAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetExtraConnectionAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetKafkaSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetKafkaSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetKinesisSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetMongodbSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetMongodbSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetMysqlSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetMysqlSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetOracleSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetOracleSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetPauseReplicationTasks() {
	_jsii_.InvokeVoid(
		a,
		"resetPauseReplicationTasks",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetPostgresSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetPostgresSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetRedisSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRedisSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetRedshiftSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshiftSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetSecretsManagerAccessRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretsManagerAccessRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetSecretsManagerArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretsManagerArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetServerName() {
	_jsii_.InvokeVoid(
		a,
		"resetServerName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetServiceAccessRole() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceAccessRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetSslMode() {
	_jsii_.InvokeVoid(
		a,
		"resetSslMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) ResetUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDmsEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDmsEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDmsEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDmsEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDmsEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDmsEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDmsEndpoint) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

