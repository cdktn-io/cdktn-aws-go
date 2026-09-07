package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kinesisfirehose/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/kinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream aws_kinesis_firehose_delivery_stream}.
// Experimental.
type AwsDeliveryStream interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	SetArn(val *string)
	// Experimental.
	ArnInput() *string
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
	Destination() *string
	// Experimental.
	SetDestination(val *string)
	// Experimental.
	DestinationId() *string
	// Experimental.
	SetDestinationId(val *string)
	// Experimental.
	DestinationIdInput() *string
	// Experimental.
	DestinationInput() *string
	// Experimental.
	ElasticsearchConfiguration() AwsDeliveryStream_ElasticsearchConfigurationPropertyOutputReference
	// Experimental.
	ElasticsearchConfigurationInput() *AwsDeliveryStream_ElasticsearchConfigurationProperty
	// Experimental.
	ExtendedS3Configuration() AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference
	// Experimental.
	ExtendedS3ConfigurationInput() *AwsDeliveryStream_ExtendedS3ConfigurationProperty
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HttpEndpointConfiguration() AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference
	// Experimental.
	HttpEndpointConfigurationInput() *AwsDeliveryStream_HttpEndpointConfigurationProperty
	// Experimental.
	IcebergConfiguration() AwsDeliveryStream_IcebergConfigurationPropertyOutputReference
	// Experimental.
	IcebergConfigurationInput() *AwsDeliveryStream_IcebergConfigurationProperty
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	KinesisSourceConfiguration() AwsDeliveryStream_KinesisSourceConfigurationPropertyOutputReference
	// Experimental.
	KinesisSourceConfigurationInput() *AwsDeliveryStream_KinesisSourceConfigurationProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MskSourceConfiguration() AwsDeliveryStream_MskSourceConfigurationPropertyOutputReference
	// Experimental.
	MskSourceConfigurationInput() *AwsDeliveryStream_MskSourceConfigurationProperty
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
	OpensearchConfiguration() AwsDeliveryStream_OpensearchConfigurationPropertyOutputReference
	// Experimental.
	OpensearchConfigurationInput() *AwsDeliveryStream_OpensearchConfigurationProperty
	// Experimental.
	OpensearchserverlessConfiguration() AwsDeliveryStream_OpensearchserverlessConfigurationPropertyOutputReference
	// Experimental.
	OpensearchserverlessConfigurationInput() *AwsDeliveryStream_OpensearchserverlessConfigurationProperty
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
	RedshiftConfiguration() AwsDeliveryStream_RedshiftConfigurationPropertyOutputReference
	// Experimental.
	RedshiftConfigurationInput() *AwsDeliveryStream_RedshiftConfigurationProperty
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	ServerSideEncryption() AwsDeliveryStream_ServerSideEncryptionPropertyOutputReference
	// Experimental.
	ServerSideEncryptionInput() *AwsDeliveryStream_ServerSideEncryptionProperty
	// Experimental.
	SnowflakeConfiguration() AwsDeliveryStream_SnowflakeConfigurationPropertyOutputReference
	// Experimental.
	SnowflakeConfigurationInput() *AwsDeliveryStream_SnowflakeConfigurationProperty
	// Experimental.
	SplunkConfiguration() AwsDeliveryStream_SplunkConfigurationPropertyOutputReference
	// Experimental.
	SplunkConfigurationInput() *AwsDeliveryStream_SplunkConfigurationProperty
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
	Timeouts() AwsDeliveryStream_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	VersionId() *string
	// Experimental.
	SetVersionId(val *string)
	// Experimental.
	VersionIdInput() *string
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
	PutElasticsearchConfiguration(value *AwsDeliveryStream_ElasticsearchConfigurationProperty)
	// Experimental.
	PutExtendedS3Configuration(value *AwsDeliveryStream_ExtendedS3ConfigurationProperty)
	// Experimental.
	PutHttpEndpointConfiguration(value *AwsDeliveryStream_HttpEndpointConfigurationProperty)
	// Experimental.
	PutIcebergConfiguration(value *AwsDeliveryStream_IcebergConfigurationProperty)
	// Experimental.
	PutKinesisSourceConfiguration(value *AwsDeliveryStream_KinesisSourceConfigurationProperty)
	// Experimental.
	PutMskSourceConfiguration(value *AwsDeliveryStream_MskSourceConfigurationProperty)
	// Experimental.
	PutOpensearchConfiguration(value *AwsDeliveryStream_OpensearchConfigurationProperty)
	// Experimental.
	PutOpensearchserverlessConfiguration(value *AwsDeliveryStream_OpensearchserverlessConfigurationProperty)
	// Experimental.
	PutRedshiftConfiguration(value *AwsDeliveryStream_RedshiftConfigurationProperty)
	// Experimental.
	PutServerSideEncryption(value *AwsDeliveryStream_ServerSideEncryptionProperty)
	// Experimental.
	PutSnowflakeConfiguration(value *AwsDeliveryStream_SnowflakeConfigurationProperty)
	// Experimental.
	PutSplunkConfiguration(value *AwsDeliveryStream_SplunkConfigurationProperty)
	// Experimental.
	PutTimeouts(value *AwsDeliveryStream_TimeoutsProperty)
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
	ResetArn()
	// Experimental.
	ResetDestinationId()
	// Experimental.
	ResetElasticsearchConfiguration()
	// Experimental.
	ResetExtendedS3Configuration()
	// Experimental.
	ResetHttpEndpointConfiguration()
	// Experimental.
	ResetIcebergConfiguration()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKinesisSourceConfiguration()
	// Experimental.
	ResetMskSourceConfiguration()
	// Experimental.
	ResetOpensearchConfiguration()
	// Experimental.
	ResetOpensearchserverlessConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRedshiftConfiguration()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetServerSideEncryption()
	// Experimental.
	ResetSnowflakeConfiguration()
	// Experimental.
	ResetSplunkConfiguration()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetVersionId()
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

// The jsii proxy struct for AwsDeliveryStream
type jsiiProxy_AwsDeliveryStream struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsDeliveryStream) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) DestinationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) DestinationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) DestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) ElasticsearchConfiguration() AwsDeliveryStream_ElasticsearchConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_ElasticsearchConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"elasticsearchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) ElasticsearchConfigurationInput() *AwsDeliveryStream_ElasticsearchConfigurationProperty {
	var returns *AwsDeliveryStream_ElasticsearchConfigurationProperty
	_jsii_.Get(
		j,
		"elasticsearchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) ExtendedS3Configuration() AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"extendedS3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) ExtendedS3ConfigurationInput() *AwsDeliveryStream_ExtendedS3ConfigurationProperty {
	var returns *AwsDeliveryStream_ExtendedS3ConfigurationProperty
	_jsii_.Get(
		j,
		"extendedS3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) HttpEndpointConfiguration() AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"httpEndpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) HttpEndpointConfigurationInput() *AwsDeliveryStream_HttpEndpointConfigurationProperty {
	var returns *AwsDeliveryStream_HttpEndpointConfigurationProperty
	_jsii_.Get(
		j,
		"httpEndpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) IcebergConfiguration() AwsDeliveryStream_IcebergConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_IcebergConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"icebergConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) IcebergConfigurationInput() *AwsDeliveryStream_IcebergConfigurationProperty {
	var returns *AwsDeliveryStream_IcebergConfigurationProperty
	_jsii_.Get(
		j,
		"icebergConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) KinesisSourceConfiguration() AwsDeliveryStream_KinesisSourceConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_KinesisSourceConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) KinesisSourceConfigurationInput() *AwsDeliveryStream_KinesisSourceConfigurationProperty {
	var returns *AwsDeliveryStream_KinesisSourceConfigurationProperty
	_jsii_.Get(
		j,
		"kinesisSourceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) MskSourceConfiguration() AwsDeliveryStream_MskSourceConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_MskSourceConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"mskSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) MskSourceConfigurationInput() *AwsDeliveryStream_MskSourceConfigurationProperty {
	var returns *AwsDeliveryStream_MskSourceConfigurationProperty
	_jsii_.Get(
		j,
		"mskSourceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) OpensearchConfiguration() AwsDeliveryStream_OpensearchConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_OpensearchConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"opensearchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) OpensearchConfigurationInput() *AwsDeliveryStream_OpensearchConfigurationProperty {
	var returns *AwsDeliveryStream_OpensearchConfigurationProperty
	_jsii_.Get(
		j,
		"opensearchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) OpensearchserverlessConfiguration() AwsDeliveryStream_OpensearchserverlessConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_OpensearchserverlessConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"opensearchserverlessConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) OpensearchserverlessConfigurationInput() *AwsDeliveryStream_OpensearchserverlessConfigurationProperty {
	var returns *AwsDeliveryStream_OpensearchserverlessConfigurationProperty
	_jsii_.Get(
		j,
		"opensearchserverlessConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) RedshiftConfiguration() AwsDeliveryStream_RedshiftConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_RedshiftConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"redshiftConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) RedshiftConfigurationInput() *AwsDeliveryStream_RedshiftConfigurationProperty {
	var returns *AwsDeliveryStream_RedshiftConfigurationProperty
	_jsii_.Get(
		j,
		"redshiftConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) ServerSideEncryption() AwsDeliveryStream_ServerSideEncryptionPropertyOutputReference {
	var returns AwsDeliveryStream_ServerSideEncryptionPropertyOutputReference
	_jsii_.Get(
		j,
		"serverSideEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) ServerSideEncryptionInput() *AwsDeliveryStream_ServerSideEncryptionProperty {
	var returns *AwsDeliveryStream_ServerSideEncryptionProperty
	_jsii_.Get(
		j,
		"serverSideEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) SnowflakeConfiguration() AwsDeliveryStream_SnowflakeConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_SnowflakeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"snowflakeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) SnowflakeConfigurationInput() *AwsDeliveryStream_SnowflakeConfigurationProperty {
	var returns *AwsDeliveryStream_SnowflakeConfigurationProperty
	_jsii_.Get(
		j,
		"snowflakeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) SplunkConfiguration() AwsDeliveryStream_SplunkConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_SplunkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"splunkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) SplunkConfigurationInput() *AwsDeliveryStream_SplunkConfigurationProperty {
	var returns *AwsDeliveryStream_SplunkConfigurationProperty
	_jsii_.Get(
		j,
		"splunkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) Timeouts() AwsDeliveryStream_TimeoutsPropertyOutputReference {
	var returns AwsDeliveryStream_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream) VersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream aws_kinesis_firehose_delivery_stream} Resource.
// Experimental.
func NewAwsDeliveryStream(scope constructs.Construct, id *string, config *AwsDeliveryStreamConfig) AwsDeliveryStream {
	_init_.Initialize()

	if err := validateNewAwsDeliveryStreamParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDeliveryStream{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream aws_kinesis_firehose_delivery_stream} Resource.
// Experimental.
func NewAwsDeliveryStream_Override(a AwsDeliveryStream, scope constructs.Construct, id *string, config *AwsDeliveryStreamConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsDeliveryStream)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream)SetDestination(val *string) {
	if err := j.validateSetDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream)SetDestinationId(val *string) {
	if err := j.validateSetDestinationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationId",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream)SetVersionId(val *string) {
	if err := j.validateSetVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionId",
		val,
	)
}

// Generates CDKTN code for importing a AwsDeliveryStream resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsDeliveryStream_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsDeliveryStream_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream",
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
func AwsDeliveryStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDeliveryStream_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsDeliveryStream_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDeliveryStream_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsDeliveryStream_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsDeliveryStream_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsDeliveryStream_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsDeliveryStream) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDeliveryStream) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeliveryStream) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDeliveryStream) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDeliveryStream) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDeliveryStream) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDeliveryStream) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDeliveryStream) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDeliveryStream) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDeliveryStream) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeliveryStream) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsDeliveryStream) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) PutElasticsearchConfiguration(value *AwsDeliveryStream_ElasticsearchConfigurationProperty) {
	if err := a.validatePutElasticsearchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putElasticsearchConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) PutExtendedS3Configuration(value *AwsDeliveryStream_ExtendedS3ConfigurationProperty) {
	if err := a.validatePutExtendedS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExtendedS3Configuration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) PutHttpEndpointConfiguration(value *AwsDeliveryStream_HttpEndpointConfigurationProperty) {
	if err := a.validatePutHttpEndpointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpEndpointConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) PutIcebergConfiguration(value *AwsDeliveryStream_IcebergConfigurationProperty) {
	if err := a.validatePutIcebergConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIcebergConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) PutKinesisSourceConfiguration(value *AwsDeliveryStream_KinesisSourceConfigurationProperty) {
	if err := a.validatePutKinesisSourceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisSourceConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) PutMskSourceConfiguration(value *AwsDeliveryStream_MskSourceConfigurationProperty) {
	if err := a.validatePutMskSourceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMskSourceConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) PutOpensearchConfiguration(value *AwsDeliveryStream_OpensearchConfigurationProperty) {
	if err := a.validatePutOpensearchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpensearchConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) PutOpensearchserverlessConfiguration(value *AwsDeliveryStream_OpensearchserverlessConfigurationProperty) {
	if err := a.validatePutOpensearchserverlessConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpensearchserverlessConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) PutRedshiftConfiguration(value *AwsDeliveryStream_RedshiftConfigurationProperty) {
	if err := a.validatePutRedshiftConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshiftConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) PutServerSideEncryption(value *AwsDeliveryStream_ServerSideEncryptionProperty) {
	if err := a.validatePutServerSideEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServerSideEncryption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) PutSnowflakeConfiguration(value *AwsDeliveryStream_SnowflakeConfigurationProperty) {
	if err := a.validatePutSnowflakeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnowflakeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) PutSplunkConfiguration(value *AwsDeliveryStream_SplunkConfigurationProperty) {
	if err := a.validatePutSplunkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSplunkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) PutTimeouts(value *AwsDeliveryStream_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetArn() {
	_jsii_.InvokeVoid(
		a,
		"resetArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetDestinationId() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetElasticsearchConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetExtendedS3Configuration() {
	_jsii_.InvokeVoid(
		a,
		"resetExtendedS3Configuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetHttpEndpointConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpEndpointConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetIcebergConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetIcebergConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetKinesisSourceConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisSourceConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetMskSourceConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetMskSourceConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetOpensearchConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetOpensearchConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetOpensearchserverlessConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetOpensearchserverlessConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetRedshiftConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshiftConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetServerSideEncryption() {
	_jsii_.InvokeVoid(
		a,
		"resetServerSideEncryption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetSnowflakeConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowflakeConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetSplunkConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSplunkConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) ResetVersionId() {
	_jsii_.InvokeVoid(
		a,
		"resetVersionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

