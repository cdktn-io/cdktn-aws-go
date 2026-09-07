package dms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/dms/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/dms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint aws_dms_s3_endpoint}.
// Experimental.
type AwsS3Endpoint interface {
	cdktn.TerraformResource
	// Experimental.
	AddColumnName() interface{}
	// Experimental.
	SetAddColumnName(val interface{})
	// Experimental.
	AddColumnNameInput() interface{}
	// Experimental.
	AddTrailingPaddingCharacter() interface{}
	// Experimental.
	SetAddTrailingPaddingCharacter(val interface{})
	// Experimental.
	AddTrailingPaddingCharacterInput() interface{}
	// Experimental.
	BucketFolder() *string
	// Experimental.
	SetBucketFolder(val *string)
	// Experimental.
	BucketFolderInput() *string
	// Experimental.
	BucketName() *string
	// Experimental.
	SetBucketName(val *string)
	// Experimental.
	BucketNameInput() *string
	// Experimental.
	CannedAclForObjects() *string
	// Experimental.
	SetCannedAclForObjects(val *string)
	// Experimental.
	CannedAclForObjectsInput() *string
	// Experimental.
	CdcInsertsAndUpdates() interface{}
	// Experimental.
	SetCdcInsertsAndUpdates(val interface{})
	// Experimental.
	CdcInsertsAndUpdatesInput() interface{}
	// Experimental.
	CdcInsertsOnly() interface{}
	// Experimental.
	SetCdcInsertsOnly(val interface{})
	// Experimental.
	CdcInsertsOnlyInput() interface{}
	// Experimental.
	CdcMaxBatchInterval() *float64
	// Experimental.
	SetCdcMaxBatchInterval(val *float64)
	// Experimental.
	CdcMaxBatchIntervalInput() *float64
	// Experimental.
	CdcMinFileSize() *float64
	// Experimental.
	SetCdcMinFileSize(val *float64)
	// Experimental.
	CdcMinFileSizeInput() *float64
	// Experimental.
	CdcPath() *string
	// Experimental.
	SetCdcPath(val *string)
	// Experimental.
	CdcPathInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CertificateArn() *string
	// Experimental.
	SetCertificateArn(val *string)
	// Experimental.
	CertificateArnInput() *string
	// Experimental.
	CompressionType() *string
	// Experimental.
	SetCompressionType(val *string)
	// Experimental.
	CompressionTypeInput() *string
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
	CsvDelimiter() *string
	// Experimental.
	SetCsvDelimiter(val *string)
	// Experimental.
	CsvDelimiterInput() *string
	// Experimental.
	CsvNoSupValue() *string
	// Experimental.
	SetCsvNoSupValue(val *string)
	// Experimental.
	CsvNoSupValueInput() *string
	// Experimental.
	CsvNullValue() *string
	// Experimental.
	SetCsvNullValue(val *string)
	// Experimental.
	CsvNullValueInput() *string
	// Experimental.
	CsvRowDelimiter() *string
	// Experimental.
	SetCsvRowDelimiter(val *string)
	// Experimental.
	CsvRowDelimiterInput() *string
	// Experimental.
	DataFormat() *string
	// Experimental.
	SetDataFormat(val *string)
	// Experimental.
	DataFormatInput() *string
	// Experimental.
	DataPageSize() *float64
	// Experimental.
	SetDataPageSize(val *float64)
	// Experimental.
	DataPageSizeInput() *float64
	// Experimental.
	DatePartitionDelimiter() *string
	// Experimental.
	SetDatePartitionDelimiter(val *string)
	// Experimental.
	DatePartitionDelimiterInput() *string
	// Experimental.
	DatePartitionEnabled() interface{}
	// Experimental.
	SetDatePartitionEnabled(val interface{})
	// Experimental.
	DatePartitionEnabledInput() interface{}
	// Experimental.
	DatePartitionSequence() *string
	// Experimental.
	SetDatePartitionSequence(val *string)
	// Experimental.
	DatePartitionSequenceInput() *string
	// Experimental.
	DatePartitionTimezone() *string
	// Experimental.
	SetDatePartitionTimezone(val *string)
	// Experimental.
	DatePartitionTimezoneInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DetachTargetOnLobLookupFailureParquet() interface{}
	// Experimental.
	SetDetachTargetOnLobLookupFailureParquet(val interface{})
	// Experimental.
	DetachTargetOnLobLookupFailureParquetInput() interface{}
	// Experimental.
	DictPageSizeLimit() *float64
	// Experimental.
	SetDictPageSizeLimit(val *float64)
	// Experimental.
	DictPageSizeLimitInput() *float64
	// Experimental.
	EnableStatistics() interface{}
	// Experimental.
	SetEnableStatistics(val interface{})
	// Experimental.
	EnableStatisticsInput() interface{}
	// Experimental.
	EncodingType() *string
	// Experimental.
	SetEncodingType(val *string)
	// Experimental.
	EncodingTypeInput() *string
	// Experimental.
	EncryptionMode() *string
	// Experimental.
	SetEncryptionMode(val *string)
	// Experimental.
	EncryptionModeInput() *string
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
	EngineDisplayName() *string
	// Experimental.
	ExpectedBucketOwner() *string
	// Experimental.
	SetExpectedBucketOwner(val *string)
	// Experimental.
	ExpectedBucketOwnerInput() *string
	// Experimental.
	ExternalId() *string
	// Experimental.
	ExternalTableDefinition() *string
	// Experimental.
	SetExternalTableDefinition(val *string)
	// Experimental.
	ExternalTableDefinitionInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	GlueCatalogGeneration() interface{}
	// Experimental.
	SetGlueCatalogGeneration(val interface{})
	// Experimental.
	GlueCatalogGenerationInput() interface{}
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	IgnoreHeaderRows() *float64
	// Experimental.
	SetIgnoreHeaderRows(val *float64)
	// Experimental.
	IgnoreHeaderRowsInput() *float64
	// Experimental.
	IncludeOpForFullLoad() interface{}
	// Experimental.
	SetIncludeOpForFullLoad(val interface{})
	// Experimental.
	IncludeOpForFullLoadInput() interface{}
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
	MaxFileSize() *float64
	// Experimental.
	SetMaxFileSize(val *float64)
	// Experimental.
	MaxFileSizeInput() *float64
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	ParquetTimestampInMillisecond() interface{}
	// Experimental.
	SetParquetTimestampInMillisecond(val interface{})
	// Experimental.
	ParquetTimestampInMillisecondInput() interface{}
	// Experimental.
	ParquetVersion() *string
	// Experimental.
	SetParquetVersion(val *string)
	// Experimental.
	ParquetVersionInput() *string
	// Experimental.
	PreserveTransactions() interface{}
	// Experimental.
	SetPreserveTransactions(val interface{})
	// Experimental.
	PreserveTransactionsInput() interface{}
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
	Rfc4180() interface{}
	// Experimental.
	SetRfc4180(val interface{})
	// Experimental.
	Rfc4180Input() interface{}
	// Experimental.
	RowGroupLength() *float64
	// Experimental.
	SetRowGroupLength(val *float64)
	// Experimental.
	RowGroupLengthInput() *float64
	// Experimental.
	ServerSideEncryptionKmsKeyId() *string
	// Experimental.
	SetServerSideEncryptionKmsKeyId(val *string)
	// Experimental.
	ServerSideEncryptionKmsKeyIdInput() *string
	// Experimental.
	ServiceAccessRoleArn() *string
	// Experimental.
	SetServiceAccessRoleArn(val *string)
	// Experimental.
	ServiceAccessRoleArnInput() *string
	// Experimental.
	SslMode() *string
	// Experimental.
	SetSslMode(val *string)
	// Experimental.
	SslModeInput() *string
	// Experimental.
	Status() *string
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
	Timeouts() AwsS3Endpoint_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TimestampColumnName() *string
	// Experimental.
	SetTimestampColumnName(val *string)
	// Experimental.
	TimestampColumnNameInput() *string
	// Experimental.
	UseCsvNoSupValue() interface{}
	// Experimental.
	SetUseCsvNoSupValue(val interface{})
	// Experimental.
	UseCsvNoSupValueInput() interface{}
	// Experimental.
	UseTaskStartTimeForFullLoadTimestamp() interface{}
	// Experimental.
	SetUseTaskStartTimeForFullLoadTimestamp(val interface{})
	// Experimental.
	UseTaskStartTimeForFullLoadTimestampInput() interface{}
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
	PutTimeouts(value *AwsS3Endpoint_TimeoutsProperty)
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
	ResetAddColumnName()
	// Experimental.
	ResetAddTrailingPaddingCharacter()
	// Experimental.
	ResetBucketFolder()
	// Experimental.
	ResetCannedAclForObjects()
	// Experimental.
	ResetCdcInsertsAndUpdates()
	// Experimental.
	ResetCdcInsertsOnly()
	// Experimental.
	ResetCdcMaxBatchInterval()
	// Experimental.
	ResetCdcMinFileSize()
	// Experimental.
	ResetCdcPath()
	// Experimental.
	ResetCertificateArn()
	// Experimental.
	ResetCompressionType()
	// Experimental.
	ResetCsvDelimiter()
	// Experimental.
	ResetCsvNoSupValue()
	// Experimental.
	ResetCsvNullValue()
	// Experimental.
	ResetCsvRowDelimiter()
	// Experimental.
	ResetDataFormat()
	// Experimental.
	ResetDataPageSize()
	// Experimental.
	ResetDatePartitionDelimiter()
	// Experimental.
	ResetDatePartitionEnabled()
	// Experimental.
	ResetDatePartitionSequence()
	// Experimental.
	ResetDatePartitionTimezone()
	// Experimental.
	ResetDetachTargetOnLobLookupFailureParquet()
	// Experimental.
	ResetDictPageSizeLimit()
	// Experimental.
	ResetEnableStatistics()
	// Experimental.
	ResetEncodingType()
	// Experimental.
	ResetEncryptionMode()
	// Experimental.
	ResetExpectedBucketOwner()
	// Experimental.
	ResetExternalTableDefinition()
	// Experimental.
	ResetGlueCatalogGeneration()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIgnoreHeaderRows()
	// Experimental.
	ResetIncludeOpForFullLoad()
	// Experimental.
	ResetKmsKeyArn()
	// Experimental.
	ResetMaxFileSize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetParquetTimestampInMillisecond()
	// Experimental.
	ResetParquetVersion()
	// Experimental.
	ResetPreserveTransactions()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRfc4180()
	// Experimental.
	ResetRowGroupLength()
	// Experimental.
	ResetServerSideEncryptionKmsKeyId()
	// Experimental.
	ResetSslMode()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTimestampColumnName()
	// Experimental.
	ResetUseCsvNoSupValue()
	// Experimental.
	ResetUseTaskStartTimeForFullLoadTimestamp()
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

// The jsii proxy struct for AwsS3Endpoint
type jsiiProxy_AwsS3Endpoint struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsS3Endpoint) AddColumnName() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) AddColumnNameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) AddTrailingPaddingCharacter() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addTrailingPaddingCharacter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) AddTrailingPaddingCharacterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addTrailingPaddingCharacterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) BucketFolder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketFolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) BucketFolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketFolderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CannedAclForObjects() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cannedAclForObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CannedAclForObjectsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cannedAclForObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CdcInsertsAndUpdates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdcInsertsAndUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CdcInsertsAndUpdatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdcInsertsAndUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CdcInsertsOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdcInsertsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CdcInsertsOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cdcInsertsOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CdcMaxBatchInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcMaxBatchInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CdcMaxBatchIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcMaxBatchIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CdcMinFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcMinFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CdcMinFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcMinFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CdcPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CdcPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CsvDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CsvDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CsvNoSupValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvNoSupValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CsvNoSupValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvNoSupValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CsvNullValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvNullValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CsvNullValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvNullValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CsvRowDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvRowDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) CsvRowDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvRowDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DataFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DataFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DataPageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataPageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DataPageSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataPageSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DatePartitionDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DatePartitionDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DatePartitionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datePartitionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DatePartitionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datePartitionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DatePartitionSequence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionSequence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DatePartitionSequenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionSequenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DatePartitionTimezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionTimezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DatePartitionTimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionTimezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DetachTargetOnLobLookupFailureParquet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detachTargetOnLobLookupFailureParquet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DetachTargetOnLobLookupFailureParquetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detachTargetOnLobLookupFailureParquetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DictPageSizeLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dictPageSizeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) DictPageSizeLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dictPageSizeLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) EnableStatistics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStatistics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) EnableStatisticsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableStatisticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) EncodingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) EncodingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) EncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) EncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) EndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) EndpointId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) EndpointIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) EngineDisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineDisplayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) ExpectedBucketOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBucketOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) ExpectedBucketOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBucketOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) ExternalTableDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTableDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) ExternalTableDefinitionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTableDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) GlueCatalogGeneration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"glueCatalogGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) GlueCatalogGenerationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"glueCatalogGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) IgnoreHeaderRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ignoreHeaderRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) IgnoreHeaderRowsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ignoreHeaderRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) IncludeOpForFullLoad() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeOpForFullLoad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) IncludeOpForFullLoadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeOpForFullLoadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) MaxFileSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) ParquetTimestampInMillisecond() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parquetTimestampInMillisecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) ParquetTimestampInMillisecondInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parquetTimestampInMillisecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) ParquetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parquetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) ParquetVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parquetVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) PreserveTransactions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveTransactions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) PreserveTransactionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveTransactionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) Rfc4180() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rfc4180",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) Rfc4180Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rfc4180Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) RowGroupLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowGroupLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) RowGroupLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowGroupLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) ServerSideEncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) ServerSideEncryptionKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) SslMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) SslModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) Timeouts() AwsS3Endpoint_TimeoutsPropertyOutputReference {
	var returns AwsS3Endpoint_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) TimestampColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) TimestampColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) UseCsvNoSupValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCsvNoSupValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) UseCsvNoSupValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCsvNoSupValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) UseTaskStartTimeForFullLoadTimestamp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTaskStartTimeForFullLoadTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3Endpoint) UseTaskStartTimeForFullLoadTimestampInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTaskStartTimeForFullLoadTimestampInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint aws_dms_s3_endpoint} Resource.
// Experimental.
func NewAwsS3Endpoint(scope constructs.Construct, id *string, config *AwsS3EndpointConfig) AwsS3Endpoint {
	_init_.Initialize()

	if err := validateNewAwsS3EndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsS3Endpoint{}

	_jsii_.Create(
		"@cdktn/aws-dms.AwsS3Endpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_s3_endpoint aws_dms_s3_endpoint} Resource.
// Experimental.
func NewAwsS3Endpoint_Override(a AwsS3Endpoint, scope constructs.Construct, id *string, config *AwsS3EndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.AwsS3Endpoint",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetAddColumnName(val interface{}) {
	if err := j.validateSetAddColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addColumnName",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetAddTrailingPaddingCharacter(val interface{}) {
	if err := j.validateSetAddTrailingPaddingCharacterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addTrailingPaddingCharacter",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetBucketFolder(val *string) {
	if err := j.validateSetBucketFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketFolder",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetCannedAclForObjects(val *string) {
	if err := j.validateSetCannedAclForObjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cannedAclForObjects",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetCdcInsertsAndUpdates(val interface{}) {
	if err := j.validateSetCdcInsertsAndUpdatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcInsertsAndUpdates",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetCdcInsertsOnly(val interface{}) {
	if err := j.validateSetCdcInsertsOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcInsertsOnly",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetCdcMaxBatchInterval(val *float64) {
	if err := j.validateSetCdcMaxBatchIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcMaxBatchInterval",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetCdcMinFileSize(val *float64) {
	if err := j.validateSetCdcMinFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcMinFileSize",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetCdcPath(val *string) {
	if err := j.validateSetCdcPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cdcPath",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetCertificateArn(val *string) {
	if err := j.validateSetCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetCompressionType(val *string) {
	if err := j.validateSetCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetCsvDelimiter(val *string) {
	if err := j.validateSetCsvDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csvDelimiter",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetCsvNoSupValue(val *string) {
	if err := j.validateSetCsvNoSupValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csvNoSupValue",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetCsvNullValue(val *string) {
	if err := j.validateSetCsvNullValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csvNullValue",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetCsvRowDelimiter(val *string) {
	if err := j.validateSetCsvRowDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"csvRowDelimiter",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetDataFormat(val *string) {
	if err := j.validateSetDataFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFormat",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetDataPageSize(val *float64) {
	if err := j.validateSetDataPageSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataPageSize",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetDatePartitionDelimiter(val *string) {
	if err := j.validateSetDatePartitionDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datePartitionDelimiter",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetDatePartitionEnabled(val interface{}) {
	if err := j.validateSetDatePartitionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datePartitionEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetDatePartitionSequence(val *string) {
	if err := j.validateSetDatePartitionSequenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datePartitionSequence",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetDatePartitionTimezone(val *string) {
	if err := j.validateSetDatePartitionTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datePartitionTimezone",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetDetachTargetOnLobLookupFailureParquet(val interface{}) {
	if err := j.validateSetDetachTargetOnLobLookupFailureParquetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detachTargetOnLobLookupFailureParquet",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetDictPageSizeLimit(val *float64) {
	if err := j.validateSetDictPageSizeLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dictPageSizeLimit",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetEnableStatistics(val interface{}) {
	if err := j.validateSetEnableStatisticsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableStatistics",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetEncodingType(val *string) {
	if err := j.validateSetEncodingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encodingType",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetEncryptionMode(val *string) {
	if err := j.validateSetEncryptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionMode",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetEndpointId(val *string) {
	if err := j.validateSetEndpointIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointId",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetEndpointType(val *string) {
	if err := j.validateSetEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetExpectedBucketOwner(val *string) {
	if err := j.validateSetExpectedBucketOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedBucketOwner",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetExternalTableDefinition(val *string) {
	if err := j.validateSetExternalTableDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalTableDefinition",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetGlueCatalogGeneration(val interface{}) {
	if err := j.validateSetGlueCatalogGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"glueCatalogGeneration",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetIgnoreHeaderRows(val *float64) {
	if err := j.validateSetIgnoreHeaderRowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreHeaderRows",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetIncludeOpForFullLoad(val interface{}) {
	if err := j.validateSetIncludeOpForFullLoadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeOpForFullLoad",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetMaxFileSize(val *float64) {
	if err := j.validateSetMaxFileSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxFileSize",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetParquetTimestampInMillisecond(val interface{}) {
	if err := j.validateSetParquetTimestampInMillisecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parquetTimestampInMillisecond",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetParquetVersion(val *string) {
	if err := j.validateSetParquetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parquetVersion",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetPreserveTransactions(val interface{}) {
	if err := j.validateSetPreserveTransactionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveTransactions",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetRfc4180(val interface{}) {
	if err := j.validateSetRfc4180Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rfc4180",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetRowGroupLength(val *float64) {
	if err := j.validateSetRowGroupLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowGroupLength",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetServerSideEncryptionKmsKeyId(val *string) {
	if err := j.validateSetServerSideEncryptionKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverSideEncryptionKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetServiceAccessRoleArn(val *string) {
	if err := j.validateSetServiceAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetSslMode(val *string) {
	if err := j.validateSetSslModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslMode",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetTimestampColumnName(val *string) {
	if err := j.validateSetTimestampColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestampColumnName",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetUseCsvNoSupValue(val interface{}) {
	if err := j.validateSetUseCsvNoSupValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCsvNoSupValue",
		val,
	)
}

func (j *jsiiProxy_AwsS3Endpoint)SetUseTaskStartTimeForFullLoadTimestamp(val interface{}) {
	if err := j.validateSetUseTaskStartTimeForFullLoadTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTaskStartTimeForFullLoadTimestamp",
		val,
	)
}

// Generates CDKTN code for importing a AwsS3Endpoint resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsS3Endpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsS3Endpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-dms.AwsS3Endpoint",
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
func AwsS3Endpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsS3Endpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dms.AwsS3Endpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsS3Endpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsS3Endpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dms.AwsS3Endpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsS3Endpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsS3Endpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dms.AwsS3Endpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsS3Endpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-dms.AwsS3Endpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsS3Endpoint) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsS3Endpoint) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsS3Endpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsS3Endpoint) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3Endpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsS3Endpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsS3Endpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsS3Endpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsS3Endpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsS3Endpoint) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsS3Endpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsS3Endpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3Endpoint) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsS3Endpoint) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3Endpoint) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsS3Endpoint) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsS3Endpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsS3Endpoint) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsS3Endpoint) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsS3Endpoint) PutTimeouts(value *AwsS3Endpoint_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3Endpoint) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetAddColumnName() {
	_jsii_.InvokeVoid(
		a,
		"resetAddColumnName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetAddTrailingPaddingCharacter() {
	_jsii_.InvokeVoid(
		a,
		"resetAddTrailingPaddingCharacter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetBucketFolder() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketFolder",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetCannedAclForObjects() {
	_jsii_.InvokeVoid(
		a,
		"resetCannedAclForObjects",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetCdcInsertsAndUpdates() {
	_jsii_.InvokeVoid(
		a,
		"resetCdcInsertsAndUpdates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetCdcInsertsOnly() {
	_jsii_.InvokeVoid(
		a,
		"resetCdcInsertsOnly",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetCdcMaxBatchInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetCdcMaxBatchInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetCdcMinFileSize() {
	_jsii_.InvokeVoid(
		a,
		"resetCdcMinFileSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetCdcPath() {
	_jsii_.InvokeVoid(
		a,
		"resetCdcPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetCompressionType() {
	_jsii_.InvokeVoid(
		a,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetCsvDelimiter() {
	_jsii_.InvokeVoid(
		a,
		"resetCsvDelimiter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetCsvNoSupValue() {
	_jsii_.InvokeVoid(
		a,
		"resetCsvNoSupValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetCsvNullValue() {
	_jsii_.InvokeVoid(
		a,
		"resetCsvNullValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetCsvRowDelimiter() {
	_jsii_.InvokeVoid(
		a,
		"resetCsvRowDelimiter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetDataFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetDataFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetDataPageSize() {
	_jsii_.InvokeVoid(
		a,
		"resetDataPageSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetDatePartitionDelimiter() {
	_jsii_.InvokeVoid(
		a,
		"resetDatePartitionDelimiter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetDatePartitionEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetDatePartitionEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetDatePartitionSequence() {
	_jsii_.InvokeVoid(
		a,
		"resetDatePartitionSequence",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetDatePartitionTimezone() {
	_jsii_.InvokeVoid(
		a,
		"resetDatePartitionTimezone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetDetachTargetOnLobLookupFailureParquet() {
	_jsii_.InvokeVoid(
		a,
		"resetDetachTargetOnLobLookupFailureParquet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetDictPageSizeLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetDictPageSizeLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetEnableStatistics() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableStatistics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetEncodingType() {
	_jsii_.InvokeVoid(
		a,
		"resetEncodingType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetEncryptionMode() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetExpectedBucketOwner() {
	_jsii_.InvokeVoid(
		a,
		"resetExpectedBucketOwner",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetExternalTableDefinition() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalTableDefinition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetGlueCatalogGeneration() {
	_jsii_.InvokeVoid(
		a,
		"resetGlueCatalogGeneration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetIgnoreHeaderRows() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreHeaderRows",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetIncludeOpForFullLoad() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeOpForFullLoad",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetMaxFileSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxFileSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetParquetTimestampInMillisecond() {
	_jsii_.InvokeVoid(
		a,
		"resetParquetTimestampInMillisecond",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetParquetVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetParquetVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetPreserveTransactions() {
	_jsii_.InvokeVoid(
		a,
		"resetPreserveTransactions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetRfc4180() {
	_jsii_.InvokeVoid(
		a,
		"resetRfc4180",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetRowGroupLength() {
	_jsii_.InvokeVoid(
		a,
		"resetRowGroupLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetServerSideEncryptionKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetServerSideEncryptionKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetSslMode() {
	_jsii_.InvokeVoid(
		a,
		"resetSslMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetTimestampColumnName() {
	_jsii_.InvokeVoid(
		a,
		"resetTimestampColumnName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetUseCsvNoSupValue() {
	_jsii_.InvokeVoid(
		a,
		"resetUseCsvNoSupValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) ResetUseTaskStartTimeForFullLoadTimestamp() {
	_jsii_.InvokeVoid(
		a,
		"resetUseTaskStartTimeForFullLoadTimestamp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3Endpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3Endpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3Endpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3Endpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3Endpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3Endpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3Endpoint) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

