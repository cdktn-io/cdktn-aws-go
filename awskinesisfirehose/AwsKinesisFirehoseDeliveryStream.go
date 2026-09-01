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
type AwsKinesisFirehoseDeliveryStream interface {
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
	ElasticsearchConfiguration() AwsKinesisFirehoseDeliveryStream_ElasticsearchConfigurationPropertyOutputReference
	// Experimental.
	ElasticsearchConfigurationInput() *AwsKinesisFirehoseDeliveryStream_ElasticsearchConfigurationProperty
	// Experimental.
	ExtendedS3Configuration() AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference
	// Experimental.
	ExtendedS3ConfigurationInput() *AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProperty
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HttpEndpointConfiguration() AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationPropertyOutputReference
	// Experimental.
	HttpEndpointConfigurationInput() *AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProperty
	// Experimental.
	IcebergConfiguration() AwsKinesisFirehoseDeliveryStream_IcebergConfigurationPropertyOutputReference
	// Experimental.
	IcebergConfigurationInput() *AwsKinesisFirehoseDeliveryStream_IcebergConfigurationProperty
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	KinesisSourceConfiguration() AwsKinesisFirehoseDeliveryStream_KinesisSourceConfigurationPropertyOutputReference
	// Experimental.
	KinesisSourceConfigurationInput() *AwsKinesisFirehoseDeliveryStream_KinesisSourceConfigurationProperty
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MskSourceConfiguration() AwsKinesisFirehoseDeliveryStream_MskSourceConfigurationPropertyOutputReference
	// Experimental.
	MskSourceConfigurationInput() *AwsKinesisFirehoseDeliveryStream_MskSourceConfigurationProperty
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
	OpensearchConfiguration() AwsKinesisFirehoseDeliveryStream_OpensearchConfigurationPropertyOutputReference
	// Experimental.
	OpensearchConfigurationInput() *AwsKinesisFirehoseDeliveryStream_OpensearchConfigurationProperty
	// Experimental.
	OpensearchserverlessConfiguration() AwsKinesisFirehoseDeliveryStream_OpensearchserverlessConfigurationPropertyOutputReference
	// Experimental.
	OpensearchserverlessConfigurationInput() *AwsKinesisFirehoseDeliveryStream_OpensearchserverlessConfigurationProperty
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
	RedshiftConfiguration() AwsKinesisFirehoseDeliveryStream_RedshiftConfigurationPropertyOutputReference
	// Experimental.
	RedshiftConfigurationInput() *AwsKinesisFirehoseDeliveryStream_RedshiftConfigurationProperty
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	ServerSideEncryption() AwsKinesisFirehoseDeliveryStream_ServerSideEncryptionPropertyOutputReference
	// Experimental.
	ServerSideEncryptionInput() *AwsKinesisFirehoseDeliveryStream_ServerSideEncryptionProperty
	// Experimental.
	SnowflakeConfiguration() AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference
	// Experimental.
	SnowflakeConfigurationInput() *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProperty
	// Experimental.
	SplunkConfiguration() AwsKinesisFirehoseDeliveryStream_SplunkConfigurationPropertyOutputReference
	// Experimental.
	SplunkConfigurationInput() *AwsKinesisFirehoseDeliveryStream_SplunkConfigurationProperty
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
	Timeouts() AwsKinesisFirehoseDeliveryStream_TimeoutsPropertyOutputReference
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
	PutElasticsearchConfiguration(value *AwsKinesisFirehoseDeliveryStream_ElasticsearchConfigurationProperty)
	// Experimental.
	PutExtendedS3Configuration(value *AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProperty)
	// Experimental.
	PutHttpEndpointConfiguration(value *AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProperty)
	// Experimental.
	PutIcebergConfiguration(value *AwsKinesisFirehoseDeliveryStream_IcebergConfigurationProperty)
	// Experimental.
	PutKinesisSourceConfiguration(value *AwsKinesisFirehoseDeliveryStream_KinesisSourceConfigurationProperty)
	// Experimental.
	PutMskSourceConfiguration(value *AwsKinesisFirehoseDeliveryStream_MskSourceConfigurationProperty)
	// Experimental.
	PutOpensearchConfiguration(value *AwsKinesisFirehoseDeliveryStream_OpensearchConfigurationProperty)
	// Experimental.
	PutOpensearchserverlessConfiguration(value *AwsKinesisFirehoseDeliveryStream_OpensearchserverlessConfigurationProperty)
	// Experimental.
	PutRedshiftConfiguration(value *AwsKinesisFirehoseDeliveryStream_RedshiftConfigurationProperty)
	// Experimental.
	PutServerSideEncryption(value *AwsKinesisFirehoseDeliveryStream_ServerSideEncryptionProperty)
	// Experimental.
	PutSnowflakeConfiguration(value *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProperty)
	// Experimental.
	PutSplunkConfiguration(value *AwsKinesisFirehoseDeliveryStream_SplunkConfigurationProperty)
	// Experimental.
	PutTimeouts(value *AwsKinesisFirehoseDeliveryStream_TimeoutsProperty)
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

// The jsii proxy struct for AwsKinesisFirehoseDeliveryStream
type jsiiProxy_AwsKinesisFirehoseDeliveryStream struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) Destination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) DestinationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) DestinationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) DestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ElasticsearchConfiguration() AwsKinesisFirehoseDeliveryStream_ElasticsearchConfigurationPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_ElasticsearchConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"elasticsearchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ElasticsearchConfigurationInput() *AwsKinesisFirehoseDeliveryStream_ElasticsearchConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_ElasticsearchConfigurationProperty
	_jsii_.Get(
		j,
		"elasticsearchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ExtendedS3Configuration() AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"extendedS3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ExtendedS3ConfigurationInput() *AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProperty
	_jsii_.Get(
		j,
		"extendedS3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) HttpEndpointConfiguration() AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"httpEndpointConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) HttpEndpointConfigurationInput() *AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProperty
	_jsii_.Get(
		j,
		"httpEndpointConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) IcebergConfiguration() AwsKinesisFirehoseDeliveryStream_IcebergConfigurationPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_IcebergConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"icebergConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) IcebergConfigurationInput() *AwsKinesisFirehoseDeliveryStream_IcebergConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_IcebergConfigurationProperty
	_jsii_.Get(
		j,
		"icebergConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) KinesisSourceConfiguration() AwsKinesisFirehoseDeliveryStream_KinesisSourceConfigurationPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_KinesisSourceConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) KinesisSourceConfigurationInput() *AwsKinesisFirehoseDeliveryStream_KinesisSourceConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_KinesisSourceConfigurationProperty
	_jsii_.Get(
		j,
		"kinesisSourceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) MskSourceConfiguration() AwsKinesisFirehoseDeliveryStream_MskSourceConfigurationPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_MskSourceConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"mskSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) MskSourceConfigurationInput() *AwsKinesisFirehoseDeliveryStream_MskSourceConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_MskSourceConfigurationProperty
	_jsii_.Get(
		j,
		"mskSourceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) OpensearchConfiguration() AwsKinesisFirehoseDeliveryStream_OpensearchConfigurationPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_OpensearchConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"opensearchConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) OpensearchConfigurationInput() *AwsKinesisFirehoseDeliveryStream_OpensearchConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_OpensearchConfigurationProperty
	_jsii_.Get(
		j,
		"opensearchConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) OpensearchserverlessConfiguration() AwsKinesisFirehoseDeliveryStream_OpensearchserverlessConfigurationPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_OpensearchserverlessConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"opensearchserverlessConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) OpensearchserverlessConfigurationInput() *AwsKinesisFirehoseDeliveryStream_OpensearchserverlessConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_OpensearchserverlessConfigurationProperty
	_jsii_.Get(
		j,
		"opensearchserverlessConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) RedshiftConfiguration() AwsKinesisFirehoseDeliveryStream_RedshiftConfigurationPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_RedshiftConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"redshiftConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) RedshiftConfigurationInput() *AwsKinesisFirehoseDeliveryStream_RedshiftConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_RedshiftConfigurationProperty
	_jsii_.Get(
		j,
		"redshiftConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ServerSideEncryption() AwsKinesisFirehoseDeliveryStream_ServerSideEncryptionPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_ServerSideEncryptionPropertyOutputReference
	_jsii_.Get(
		j,
		"serverSideEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ServerSideEncryptionInput() *AwsKinesisFirehoseDeliveryStream_ServerSideEncryptionProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_ServerSideEncryptionProperty
	_jsii_.Get(
		j,
		"serverSideEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) SnowflakeConfiguration() AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"snowflakeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) SnowflakeConfigurationInput() *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProperty
	_jsii_.Get(
		j,
		"snowflakeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) SplunkConfiguration() AwsKinesisFirehoseDeliveryStream_SplunkConfigurationPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_SplunkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"splunkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) SplunkConfigurationInput() *AwsKinesisFirehoseDeliveryStream_SplunkConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_SplunkConfigurationProperty
	_jsii_.Get(
		j,
		"splunkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) Timeouts() AwsKinesisFirehoseDeliveryStream_TimeoutsPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream) VersionIdInput() *string {
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
func NewAwsKinesisFirehoseDeliveryStream(scope constructs.Construct, id *string, config *AwsKinesisFirehoseDeliveryStreamConfig) AwsKinesisFirehoseDeliveryStream {
	_init_.Initialize()

	if err := validateNewAwsKinesisFirehoseDeliveryStreamParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKinesisFirehoseDeliveryStream{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsKinesisFirehoseDeliveryStream",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesis_firehose_delivery_stream aws_kinesis_firehose_delivery_stream} Resource.
// Experimental.
func NewAwsKinesisFirehoseDeliveryStream_Override(a AwsKinesisFirehoseDeliveryStream, scope constructs.Construct, id *string, config *AwsKinesisFirehoseDeliveryStreamConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsKinesisFirehoseDeliveryStream",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream)SetDestination(val *string) {
	if err := j.validateSetDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destination",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream)SetDestinationId(val *string) {
	if err := j.validateSetDestinationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationId",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream)SetVersionId(val *string) {
	if err := j.validateSetVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionId",
		val,
	)
}

// Generates CDKTN code for importing a AwsKinesisFirehoseDeliveryStream resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsKinesisFirehoseDeliveryStream_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsKinesisFirehoseDeliveryStream_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-kinesis-firehose.AwsKinesisFirehoseDeliveryStream",
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
func AwsKinesisFirehoseDeliveryStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsKinesisFirehoseDeliveryStream_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-kinesis-firehose.AwsKinesisFirehoseDeliveryStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsKinesisFirehoseDeliveryStream_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsKinesisFirehoseDeliveryStream_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-kinesis-firehose.AwsKinesisFirehoseDeliveryStream",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsKinesisFirehoseDeliveryStream_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsKinesisFirehoseDeliveryStream_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-kinesis-firehose.AwsKinesisFirehoseDeliveryStream",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsKinesisFirehoseDeliveryStream_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-kinesis-firehose.AwsKinesisFirehoseDeliveryStream",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) PutElasticsearchConfiguration(value *AwsKinesisFirehoseDeliveryStream_ElasticsearchConfigurationProperty) {
	if err := a.validatePutElasticsearchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putElasticsearchConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) PutExtendedS3Configuration(value *AwsKinesisFirehoseDeliveryStream_ExtendedS3ConfigurationProperty) {
	if err := a.validatePutExtendedS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExtendedS3Configuration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) PutHttpEndpointConfiguration(value *AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProperty) {
	if err := a.validatePutHttpEndpointConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpEndpointConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) PutIcebergConfiguration(value *AwsKinesisFirehoseDeliveryStream_IcebergConfigurationProperty) {
	if err := a.validatePutIcebergConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIcebergConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) PutKinesisSourceConfiguration(value *AwsKinesisFirehoseDeliveryStream_KinesisSourceConfigurationProperty) {
	if err := a.validatePutKinesisSourceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisSourceConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) PutMskSourceConfiguration(value *AwsKinesisFirehoseDeliveryStream_MskSourceConfigurationProperty) {
	if err := a.validatePutMskSourceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMskSourceConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) PutOpensearchConfiguration(value *AwsKinesisFirehoseDeliveryStream_OpensearchConfigurationProperty) {
	if err := a.validatePutOpensearchConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpensearchConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) PutOpensearchserverlessConfiguration(value *AwsKinesisFirehoseDeliveryStream_OpensearchserverlessConfigurationProperty) {
	if err := a.validatePutOpensearchserverlessConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOpensearchserverlessConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) PutRedshiftConfiguration(value *AwsKinesisFirehoseDeliveryStream_RedshiftConfigurationProperty) {
	if err := a.validatePutRedshiftConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshiftConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) PutServerSideEncryption(value *AwsKinesisFirehoseDeliveryStream_ServerSideEncryptionProperty) {
	if err := a.validatePutServerSideEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServerSideEncryption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) PutSnowflakeConfiguration(value *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProperty) {
	if err := a.validatePutSnowflakeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnowflakeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) PutSplunkConfiguration(value *AwsKinesisFirehoseDeliveryStream_SplunkConfigurationProperty) {
	if err := a.validatePutSplunkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSplunkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) PutTimeouts(value *AwsKinesisFirehoseDeliveryStream_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetArn() {
	_jsii_.InvokeVoid(
		a,
		"resetArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetDestinationId() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetElasticsearchConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearchConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetExtendedS3Configuration() {
	_jsii_.InvokeVoid(
		a,
		"resetExtendedS3Configuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetHttpEndpointConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpEndpointConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetIcebergConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetIcebergConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetKinesisSourceConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisSourceConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetMskSourceConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetMskSourceConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetOpensearchConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetOpensearchConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetOpensearchserverlessConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetOpensearchserverlessConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetRedshiftConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshiftConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetServerSideEncryption() {
	_jsii_.InvokeVoid(
		a,
		"resetServerSideEncryption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetSnowflakeConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowflakeConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetSplunkConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSplunkConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ResetVersionId() {
	_jsii_.InvokeVoid(
		a,
		"resetVersionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

