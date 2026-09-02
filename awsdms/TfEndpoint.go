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
type TfEndpoint interface {
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
	ElasticsearchSettings() TfEndpoint_ElasticsearchSettingsPropertyOutputReference
	// Experimental.
	ElasticsearchSettingsInput() *TfEndpoint_ElasticsearchSettingsProperty
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
	KafkaSettings() TfEndpoint_KafkaSettingsPropertyOutputReference
	// Experimental.
	KafkaSettingsInput() *TfEndpoint_KafkaSettingsProperty
	// Experimental.
	KinesisSettings() TfEndpoint_KinesisSettingsPropertyOutputReference
	// Experimental.
	KinesisSettingsInput() *TfEndpoint_KinesisSettingsProperty
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
	MongodbSettings() TfEndpoint_MongodbSettingsPropertyOutputReference
	// Experimental.
	MongodbSettingsInput() *TfEndpoint_MongodbSettingsProperty
	// Experimental.
	MysqlSettings() TfEndpoint_MysqlSettingsPropertyOutputReference
	// Experimental.
	MysqlSettingsInput() *TfEndpoint_MysqlSettingsProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OracleSettings() TfEndpoint_OracleSettingsPropertyOutputReference
	// Experimental.
	OracleSettingsInput() *TfEndpoint_OracleSettingsProperty
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
	PostgresSettings() TfEndpoint_PostgresSettingsPropertyOutputReference
	// Experimental.
	PostgresSettingsInput() *TfEndpoint_PostgresSettingsProperty
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
	RedisSettings() TfEndpoint_RedisSettingsPropertyOutputReference
	// Experimental.
	RedisSettingsInput() *TfEndpoint_RedisSettingsProperty
	// Experimental.
	RedshiftSettings() TfEndpoint_RedshiftSettingsPropertyOutputReference
	// Experimental.
	RedshiftSettingsInput() *TfEndpoint_RedshiftSettingsProperty
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
	Timeouts() TfEndpoint_TimeoutsPropertyOutputReference
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
	PutElasticsearchSettings(value *TfEndpoint_ElasticsearchSettingsProperty)
	// Experimental.
	PutKafkaSettings(value *TfEndpoint_KafkaSettingsProperty)
	// Experimental.
	PutKinesisSettings(value *TfEndpoint_KinesisSettingsProperty)
	// Experimental.
	PutMongodbSettings(value *TfEndpoint_MongodbSettingsProperty)
	// Experimental.
	PutMysqlSettings(value *TfEndpoint_MysqlSettingsProperty)
	// Experimental.
	PutOracleSettings(value *TfEndpoint_OracleSettingsProperty)
	// Experimental.
	PutPostgresSettings(value *TfEndpoint_PostgresSettingsProperty)
	// Experimental.
	PutRedisSettings(value *TfEndpoint_RedisSettingsProperty)
	// Experimental.
	PutRedshiftSettings(value *TfEndpoint_RedshiftSettingsProperty)
	// Experimental.
	PutTimeouts(value *TfEndpoint_TimeoutsProperty)
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

// The jsii proxy struct for TfEndpoint
type jsiiProxy_TfEndpoint struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfEndpoint) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ElasticsearchSettings() TfEndpoint_ElasticsearchSettingsPropertyOutputReference {
	var returns TfEndpoint_ElasticsearchSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"elasticsearchSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ElasticsearchSettingsInput() *TfEndpoint_ElasticsearchSettingsProperty {
	var returns *TfEndpoint_ElasticsearchSettingsProperty
	_jsii_.Get(
		j,
		"elasticsearchSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) EndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) EndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) EndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) EngineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) EngineNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ExtraConnectionAttributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraConnectionAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ExtraConnectionAttributesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraConnectionAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) KafkaSettings() TfEndpoint_KafkaSettingsPropertyOutputReference {
	var returns TfEndpoint_KafkaSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"kafkaSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) KafkaSettingsInput() *TfEndpoint_KafkaSettingsProperty {
	var returns *TfEndpoint_KafkaSettingsProperty
	_jsii_.Get(
		j,
		"kafkaSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) KinesisSettings() TfEndpoint_KinesisSettingsPropertyOutputReference {
	var returns TfEndpoint_KinesisSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) KinesisSettingsInput() *TfEndpoint_KinesisSettingsProperty {
	var returns *TfEndpoint_KinesisSettingsProperty
	_jsii_.Get(
		j,
		"kinesisSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) MongodbSettings() TfEndpoint_MongodbSettingsPropertyOutputReference {
	var returns TfEndpoint_MongodbSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"mongodbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) MongodbSettingsInput() *TfEndpoint_MongodbSettingsProperty {
	var returns *TfEndpoint_MongodbSettingsProperty
	_jsii_.Get(
		j,
		"mongodbSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) MysqlSettings() TfEndpoint_MysqlSettingsPropertyOutputReference {
	var returns TfEndpoint_MysqlSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"mysqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) MysqlSettingsInput() *TfEndpoint_MysqlSettingsProperty {
	var returns *TfEndpoint_MysqlSettingsProperty
	_jsii_.Get(
		j,
		"mysqlSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) OracleSettings() TfEndpoint_OracleSettingsPropertyOutputReference {
	var returns TfEndpoint_OracleSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"oracleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) OracleSettingsInput() *TfEndpoint_OracleSettingsProperty {
	var returns *TfEndpoint_OracleSettingsProperty
	_jsii_.Get(
		j,
		"oracleSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) PauseReplicationTasks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pauseReplicationTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) PauseReplicationTasksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pauseReplicationTasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) PostgresSettings() TfEndpoint_PostgresSettingsPropertyOutputReference {
	var returns TfEndpoint_PostgresSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"postgresSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) PostgresSettingsInput() *TfEndpoint_PostgresSettingsProperty {
	var returns *TfEndpoint_PostgresSettingsProperty
	_jsii_.Get(
		j,
		"postgresSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) RedisSettings() TfEndpoint_RedisSettingsPropertyOutputReference {
	var returns TfEndpoint_RedisSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"redisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) RedisSettingsInput() *TfEndpoint_RedisSettingsProperty {
	var returns *TfEndpoint_RedisSettingsProperty
	_jsii_.Get(
		j,
		"redisSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) RedshiftSettings() TfEndpoint_RedshiftSettingsPropertyOutputReference {
	var returns TfEndpoint_RedshiftSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"redshiftSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) RedshiftSettingsInput() *TfEndpoint_RedshiftSettingsProperty {
	var returns *TfEndpoint_RedshiftSettingsProperty
	_jsii_.Get(
		j,
		"redshiftSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) SecretsManagerAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) SecretsManagerAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) SecretsManagerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) SecretsManagerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ServiceAccessRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) ServiceAccessRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) SslMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) SslModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Timeouts() TfEndpoint_TimeoutsPropertyOutputReference {
	var returns TfEndpoint_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint) UsernameInput() *string {
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
func NewTfEndpoint(scope constructs.Construct, id *string, config *TfEndpointConfig) TfEndpoint {
	_init_.Initialize()

	if err := validateNewTfEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpoint{}

	_jsii_.Create(
		"@cdktn/aws-dms.TfEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_endpoint aws_dms_endpoint} Resource.
// Experimental.
func NewTfEndpoint_Override(t TfEndpoint, scope constructs.Construct, id *string, config *TfEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.TfEndpoint",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfEndpoint)SetCertificateArn(val *string) {
	if err := j.validateSetCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetEndpointId(val *string) {
	if err := j.validateSetEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointId",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetEndpointType(val *string) {
	if err := j.validateSetEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetEngineName(val *string) {
	if err := j.validateSetEngineNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineName",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetExtraConnectionAttributes(val *string) {
	if err := j.validateSetExtraConnectionAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraConnectionAttributes",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetPauseReplicationTasks(val interface{}) {
	if err := j.validateSetPauseReplicationTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pauseReplicationTasks",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetSecretsManagerAccessRoleArn(val *string) {
	if err := j.validateSetSecretsManagerAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetSecretsManagerArn(val *string) {
	if err := j.validateSetSecretsManagerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetServerName(val *string) {
	if err := j.validateSetServerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverName",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetServiceAccessRole(val *string) {
	if err := j.validateSetServiceAccessRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRole",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetSslMode(val *string) {
	if err := j.validateSetSslModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslMode",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Generates CDKTN code for importing a TfEndpoint resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-dms.TfEndpoint",
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
func TfEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dms.TfEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dms.TfEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dms.TfEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-dms.TfEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfEndpoint) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfEndpoint) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfEndpoint) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfEndpoint) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfEndpoint) PutElasticsearchSettings(value *TfEndpoint_ElasticsearchSettingsProperty) {
	if err := t.validatePutElasticsearchSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putElasticsearchSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint) PutKafkaSettings(value *TfEndpoint_KafkaSettingsProperty) {
	if err := t.validatePutKafkaSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKafkaSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint) PutKinesisSettings(value *TfEndpoint_KinesisSettingsProperty) {
	if err := t.validatePutKinesisSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesisSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint) PutMongodbSettings(value *TfEndpoint_MongodbSettingsProperty) {
	if err := t.validatePutMongodbSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMongodbSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint) PutMysqlSettings(value *TfEndpoint_MysqlSettingsProperty) {
	if err := t.validatePutMysqlSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMysqlSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint) PutOracleSettings(value *TfEndpoint_OracleSettingsProperty) {
	if err := t.validatePutOracleSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOracleSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint) PutPostgresSettings(value *TfEndpoint_PostgresSettingsProperty) {
	if err := t.validatePutPostgresSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPostgresSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint) PutRedisSettings(value *TfEndpoint_RedisSettingsProperty) {
	if err := t.validatePutRedisSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRedisSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint) PutRedshiftSettings(value *TfEndpoint_RedshiftSettingsProperty) {
	if err := t.validatePutRedshiftSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRedshiftSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint) PutTimeouts(value *TfEndpoint_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfEndpoint) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		t,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		t,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetElasticsearchSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetElasticsearchSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetExtraConnectionAttributes() {
	_jsii_.InvokeVoid(
		t,
		"resetExtraConnectionAttributes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetKafkaSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetKafkaSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetKinesisSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesisSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetMongodbSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetMongodbSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetMysqlSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetMysqlSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetOracleSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetOracleSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetPassword() {
	_jsii_.InvokeVoid(
		t,
		"resetPassword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetPauseReplicationTasks() {
	_jsii_.InvokeVoid(
		t,
		"resetPauseReplicationTasks",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetPort() {
	_jsii_.InvokeVoid(
		t,
		"resetPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetPostgresSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetPostgresSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetRedisSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetRedisSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetRedshiftSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetRedshiftSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetSecretsManagerAccessRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSecretsManagerAccessRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetSecretsManagerArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSecretsManagerArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetServerName() {
	_jsii_.InvokeVoid(
		t,
		"resetServerName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetServiceAccessRole() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceAccessRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetSslMode() {
	_jsii_.InvokeVoid(
		t,
		"resetSslMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) ResetUsername() {
	_jsii_.InvokeVoid(
		t,
		"resetUsername",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

