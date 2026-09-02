package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream aws_kinesis_firehose_delivery_stream}.
// Experimental.
type TfDeliveryStream interface {
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
	ElasticsearchConfiguration() TfDeliveryStream_ElasticsearchConfigurationPropertyOutputReference
	// Experimental.
	ElasticsearchConfigurationInput() *TfDeliveryStream_ElasticsearchConfigurationProperty
	// Experimental.
	ExtendedS3Configuration() TfDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference
	// Experimental.
	ExtendedS3ConfigurationInput() *TfDeliveryStream_ExtendedS3ConfigurationProperty
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HttpEndpointConfiguration() TfDeliveryStream_HttpEndpointConfigurationPropertyOutputReference
	// Experimental.
	HttpEndpointConfigurationInput() *TfDeliveryStream_HttpEndpointConfigurationProperty
	// Experimental.
	IcebergConfiguration() TfDeliveryStream_IcebergConfigurationPropertyOutputReference
	// Experimental.
	IcebergConfigurationInput() *TfDeliveryStream_IcebergConfigurationProperty
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	KinesisSourceConfiguration() TfDeliveryStream_KinesisSourceConfigurationPropertyOutputReference
	// Experimental.
	KinesisSourceConfigurationInput() *TfDeliveryStream_KinesisSourceConfigurationProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MskSourceConfiguration() TfDeliveryStream_MskSourceConfigurationPropertyOutputReference
	// Experimental.
	MskSourceConfigurationInput() *TfDeliveryStream_MskSourceConfigurationProperty
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
	OpensearchConfiguration() TfDeliveryStream_OpensearchConfigurationPropertyOutputReference
	// Experimental.
	OpensearchConfigurationInput() *TfDeliveryStream_OpensearchConfigurationProperty
	// Experimental.
	OpensearchserverlessConfiguration() TfDeliveryStream_OpensearchserverlessConfigurationPropertyOutputReference
	// Experimental.
	OpensearchserverlessConfigurationInput() *TfDeliveryStream_OpensearchserverlessConfigurationProperty
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
	RedshiftConfiguration() TfDeliveryStream_RedshiftConfigurationPropertyOutputReference
	// Experimental.
	RedshiftConfigurationInput() *TfDeliveryStream_RedshiftConfigurationProperty
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	ServerSideEncryption() TfDeliveryStream_ServerSideEncryptionPropertyOutputReference
	// Experimental.
	ServerSideEncryptionInput() *TfDeliveryStream_ServerSideEncryptionProperty
	// Experimental.
	SnowflakeConfiguration() TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference
	// Experimental.
	SnowflakeConfigurationInput() *TfDeliveryStream_SnowflakeConfigurationProperty
	// Experimental.
	SplunkConfiguration() TfDeliveryStream_SplunkConfigurationPropertyOutputReference
	// Experimental.
	SplunkConfigurationInput() *TfDeliveryStream_SplunkConfigurationProperty
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
	Timeouts() TfDeliveryStream_TimeoutsPropertyOutputReference
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
	PutElasticsearchConfiguration(value *TfDeliveryStream_ElasticsearchConfigurationProperty)
	// Experimental.
	PutExtendedS3Configuration(value *TfDeliveryStream_ExtendedS3ConfigurationProperty)
	// Experimental.
	PutHttpEndpointConfiguration(value *TfDeliveryStream_HttpEndpointConfigurationProperty)
	// Experimental.
	PutIcebergConfiguration(value *TfDeliveryStream_IcebergConfigurationProperty)
	// Experimental.
	PutKinesisSourceConfiguration(value *TfDeliveryStream_KinesisSourceConfigurationProperty)
	// Experimental.
	PutMskSourceConfiguration(value *TfDeliveryStream_MskSourceConfigurationProperty)
	// Experimental.
	PutOpensearchConfiguration(value *TfDeliveryStream_OpensearchConfigurationProperty)
	// Experimental.
	PutOpensearchserverlessConfiguration(value *TfDeliveryStream_OpensearchserverlessConfigurationProperty)
	// Experimental.
	PutRedshiftConfiguration(value *TfDeliveryStream_RedshiftConfigurationProperty)
	// Experimental.
	PutServerSideEncryption(value *TfDeliveryStream_ServerSideEncryptionProperty)
	// Experimental.
	PutSnowflakeConfiguration(value *TfDeliveryStream_SnowflakeConfigurationProperty)
	// Experimental.
	PutSplunkConfiguration(value *TfDeliveryStream_SplunkConfigurationProperty)
	// Experimental.
	PutTimeouts(value *TfDeliveryStream_TimeoutsProperty)
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

// The jsii proxy struct for TfDeliveryStream
type jsiiProxy_TfDeliveryStream struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfDeliveryStream) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) DestinationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) DestinationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) DestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) ElasticsearchConfiguration() TfDeliveryStream_ElasticsearchConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_ElasticsearchConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"elasticsearchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) ElasticsearchConfigurationInput() *TfDeliveryStream_ElasticsearchConfigurationProperty {
	var returns *TfDeliveryStream_ElasticsearchConfigurationProperty
	_jsii_.Get(
		j,
		"elasticsearchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) ExtendedS3Configuration() TfDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"extendedS3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) ExtendedS3ConfigurationInput() *TfDeliveryStream_ExtendedS3ConfigurationProperty {
	var returns *TfDeliveryStream_ExtendedS3ConfigurationProperty
	_jsii_.Get(
		j,
		"extendedS3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) HttpEndpointConfiguration() TfDeliveryStream_HttpEndpointConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_HttpEndpointConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"httpEndpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) HttpEndpointConfigurationInput() *TfDeliveryStream_HttpEndpointConfigurationProperty {
	var returns *TfDeliveryStream_HttpEndpointConfigurationProperty
	_jsii_.Get(
		j,
		"httpEndpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) IcebergConfiguration() TfDeliveryStream_IcebergConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_IcebergConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"icebergConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) IcebergConfigurationInput() *TfDeliveryStream_IcebergConfigurationProperty {
	var returns *TfDeliveryStream_IcebergConfigurationProperty
	_jsii_.Get(
		j,
		"icebergConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) KinesisSourceConfiguration() TfDeliveryStream_KinesisSourceConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_KinesisSourceConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) KinesisSourceConfigurationInput() *TfDeliveryStream_KinesisSourceConfigurationProperty {
	var returns *TfDeliveryStream_KinesisSourceConfigurationProperty
	_jsii_.Get(
		j,
		"kinesisSourceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) MskSourceConfiguration() TfDeliveryStream_MskSourceConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_MskSourceConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"mskSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) MskSourceConfigurationInput() *TfDeliveryStream_MskSourceConfigurationProperty {
	var returns *TfDeliveryStream_MskSourceConfigurationProperty
	_jsii_.Get(
		j,
		"mskSourceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) OpensearchConfiguration() TfDeliveryStream_OpensearchConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_OpensearchConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"opensearchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) OpensearchConfigurationInput() *TfDeliveryStream_OpensearchConfigurationProperty {
	var returns *TfDeliveryStream_OpensearchConfigurationProperty
	_jsii_.Get(
		j,
		"opensearchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) OpensearchserverlessConfiguration() TfDeliveryStream_OpensearchserverlessConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_OpensearchserverlessConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"opensearchserverlessConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) OpensearchserverlessConfigurationInput() *TfDeliveryStream_OpensearchserverlessConfigurationProperty {
	var returns *TfDeliveryStream_OpensearchserverlessConfigurationProperty
	_jsii_.Get(
		j,
		"opensearchserverlessConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) RedshiftConfiguration() TfDeliveryStream_RedshiftConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_RedshiftConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"redshiftConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) RedshiftConfigurationInput() *TfDeliveryStream_RedshiftConfigurationProperty {
	var returns *TfDeliveryStream_RedshiftConfigurationProperty
	_jsii_.Get(
		j,
		"redshiftConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) ServerSideEncryption() TfDeliveryStream_ServerSideEncryptionPropertyOutputReference {
	var returns TfDeliveryStream_ServerSideEncryptionPropertyOutputReference
	_jsii_.Get(
		j,
		"serverSideEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) ServerSideEncryptionInput() *TfDeliveryStream_ServerSideEncryptionProperty {
	var returns *TfDeliveryStream_ServerSideEncryptionProperty
	_jsii_.Get(
		j,
		"serverSideEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) SnowflakeConfiguration() TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"snowflakeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) SnowflakeConfigurationInput() *TfDeliveryStream_SnowflakeConfigurationProperty {
	var returns *TfDeliveryStream_SnowflakeConfigurationProperty
	_jsii_.Get(
		j,
		"snowflakeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) SplunkConfiguration() TfDeliveryStream_SplunkConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_SplunkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"splunkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) SplunkConfigurationInput() *TfDeliveryStream_SplunkConfigurationProperty {
	var returns *TfDeliveryStream_SplunkConfigurationProperty
	_jsii_.Get(
		j,
		"splunkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) Timeouts() TfDeliveryStream_TimeoutsPropertyOutputReference {
	var returns TfDeliveryStream_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream) VersionIdInput() *string {
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
func NewTfDeliveryStream(scope constructs.Construct, id *string, config *TfDeliveryStreamConfig) TfDeliveryStream {
	_init_.Initialize()

	if err := validateNewTfDeliveryStreamParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDeliveryStream{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream aws_kinesis_firehose_delivery_stream} Resource.
// Experimental.
func NewTfDeliveryStream_Override(t TfDeliveryStream, scope constructs.Construct, id *string, config *TfDeliveryStreamConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfDeliveryStream)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream)SetDestination(val *string) {
	if err := j.validateSetDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream)SetDestinationId(val *string) {
	if err := j.validateSetDestinationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationId",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream)SetVersionId(val *string) {
	if err := j.validateSetVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionId",
		val,
	)
}

// Generates CDKTN code for importing a TfDeliveryStream resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfDeliveryStream_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfDeliveryStream_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream",
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
func TfDeliveryStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDeliveryStream_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfDeliveryStream_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDeliveryStream_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfDeliveryStream_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfDeliveryStream_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfDeliveryStream_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfDeliveryStream) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfDeliveryStream) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfDeliveryStream) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDeliveryStream) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDeliveryStream) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDeliveryStream) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDeliveryStream) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDeliveryStream) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDeliveryStream) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfDeliveryStream) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfDeliveryStream) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfDeliveryStream) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfDeliveryStream) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfDeliveryStream) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfDeliveryStream) PutElasticsearchConfiguration(value *TfDeliveryStream_ElasticsearchConfigurationProperty) {
	if err := t.validatePutElasticsearchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putElasticsearchConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream) PutExtendedS3Configuration(value *TfDeliveryStream_ExtendedS3ConfigurationProperty) {
	if err := t.validatePutExtendedS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExtendedS3Configuration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream) PutHttpEndpointConfiguration(value *TfDeliveryStream_HttpEndpointConfigurationProperty) {
	if err := t.validatePutHttpEndpointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHttpEndpointConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream) PutIcebergConfiguration(value *TfDeliveryStream_IcebergConfigurationProperty) {
	if err := t.validatePutIcebergConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIcebergConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream) PutKinesisSourceConfiguration(value *TfDeliveryStream_KinesisSourceConfigurationProperty) {
	if err := t.validatePutKinesisSourceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesisSourceConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream) PutMskSourceConfiguration(value *TfDeliveryStream_MskSourceConfigurationProperty) {
	if err := t.validatePutMskSourceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMskSourceConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream) PutOpensearchConfiguration(value *TfDeliveryStream_OpensearchConfigurationProperty) {
	if err := t.validatePutOpensearchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOpensearchConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream) PutOpensearchserverlessConfiguration(value *TfDeliveryStream_OpensearchserverlessConfigurationProperty) {
	if err := t.validatePutOpensearchserverlessConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOpensearchserverlessConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream) PutRedshiftConfiguration(value *TfDeliveryStream_RedshiftConfigurationProperty) {
	if err := t.validatePutRedshiftConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRedshiftConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream) PutServerSideEncryption(value *TfDeliveryStream_ServerSideEncryptionProperty) {
	if err := t.validatePutServerSideEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putServerSideEncryption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream) PutSnowflakeConfiguration(value *TfDeliveryStream_SnowflakeConfigurationProperty) {
	if err := t.validatePutSnowflakeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSnowflakeConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream) PutSplunkConfiguration(value *TfDeliveryStream_SplunkConfigurationProperty) {
	if err := t.validatePutSplunkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSplunkConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream) PutTimeouts(value *TfDeliveryStream_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetArn() {
	_jsii_.InvokeVoid(
		t,
		"resetArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetDestinationId() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetElasticsearchConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetElasticsearchConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetExtendedS3Configuration() {
	_jsii_.InvokeVoid(
		t,
		"resetExtendedS3Configuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetHttpEndpointConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpEndpointConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetIcebergConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetIcebergConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetKinesisSourceConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesisSourceConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetMskSourceConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetMskSourceConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetOpensearchConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetOpensearchConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetOpensearchserverlessConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetOpensearchserverlessConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetRedshiftConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetRedshiftConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetServerSideEncryption() {
	_jsii_.InvokeVoid(
		t,
		"resetServerSideEncryption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetSnowflakeConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSnowflakeConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetSplunkConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSplunkConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) ResetVersionId() {
	_jsii_.InvokeVoid(
		t,
		"resetVersionId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

