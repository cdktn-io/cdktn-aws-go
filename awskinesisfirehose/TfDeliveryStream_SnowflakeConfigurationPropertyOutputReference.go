package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccountUrl() *string
	// Experimental.
	SetAccountUrl(val *string)
	// Experimental.
	AccountUrlInput() *string
	// Experimental.
	BufferingInterval() *float64
	// Experimental.
	SetBufferingInterval(val *float64)
	// Experimental.
	BufferingIntervalInput() *float64
	// Experimental.
	BufferingSize() *float64
	// Experimental.
	SetBufferingSize(val *float64)
	// Experimental.
	BufferingSizeInput() *float64
	// Experimental.
	CloudwatchLoggingOptions() TfDeliveryStream_SnowflakeConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	// Experimental.
	CloudwatchLoggingOptionsInput() *TfDeliveryStream_SnowflakeConfigurationCloudwatchLoggingOptionsProperty
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// Experimental.
	ContentColumnName() *string
	// Experimental.
	SetContentColumnName(val *string)
	// Experimental.
	ContentColumnNameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Database() *string
	// Experimental.
	SetDatabase(val *string)
	// Experimental.
	DatabaseInput() *string
	// Experimental.
	DataLoadingOption() *string
	// Experimental.
	SetDataLoadingOption(val *string)
	// Experimental.
	DataLoadingOptionInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDeliveryStream_SnowflakeConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfDeliveryStream_SnowflakeConfigurationProperty)
	// Experimental.
	KeyPassphrase() *string
	// Experimental.
	SetKeyPassphrase(val *string)
	// Experimental.
	KeyPassphraseInput() *string
	// Experimental.
	MetadataColumnName() *string
	// Experimental.
	SetMetadataColumnName(val *string)
	// Experimental.
	MetadataColumnNameInput() *string
	// Experimental.
	PrivateKey() *string
	// Experimental.
	SetPrivateKey(val *string)
	// Experimental.
	PrivateKeyInput() *string
	// Experimental.
	ProcessingConfiguration() TfDeliveryStream_SnowflakeConfigurationProcessingConfigurationPropertyOutputReference
	// Experimental.
	ProcessingConfigurationInput() *TfDeliveryStream_SnowflakeConfigurationProcessingConfigurationProperty
	// Experimental.
	RetryDuration() *float64
	// Experimental.
	SetRetryDuration(val *float64)
	// Experimental.
	RetryDurationInput() *float64
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	S3BackupMode() *string
	// Experimental.
	SetS3BackupMode(val *string)
	// Experimental.
	S3BackupModeInput() *string
	// Experimental.
	S3Configuration() TfDeliveryStream_SnowflakeConfigurationS3ConfigurationPropertyOutputReference
	// Experimental.
	S3ConfigurationInput() *TfDeliveryStream_SnowflakeConfigurationS3ConfigurationProperty
	// Experimental.
	Schema() *string
	// Experimental.
	SetSchema(val *string)
	// Experimental.
	SchemaInput() *string
	// Experimental.
	SecretsManagerConfiguration() TfDeliveryStream_SnowflakeConfigurationSecretsManagerConfigurationPropertyOutputReference
	// Experimental.
	SecretsManagerConfigurationInput() *TfDeliveryStream_SnowflakeConfigurationSecretsManagerConfigurationProperty
	// Experimental.
	SnowflakeRoleConfiguration() TfDeliveryStream_SnowflakeRoleConfigurationPropertyOutputReference
	// Experimental.
	SnowflakeRoleConfigurationInput() *TfDeliveryStream_SnowflakeRoleConfigurationProperty
	// Experimental.
	SnowflakeVpcConfiguration() TfDeliveryStream_SnowflakeVpcConfigurationPropertyOutputReference
	// Experimental.
	SnowflakeVpcConfigurationInput() *TfDeliveryStream_SnowflakeVpcConfigurationProperty
	// Experimental.
	Table() *string
	// Experimental.
	SetTable(val *string)
	// Experimental.
	TableInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	User() *string
	// Experimental.
	SetUser(val *string)
	// Experimental.
	UserInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	PutCloudwatchLoggingOptions(value *TfDeliveryStream_SnowflakeConfigurationCloudwatchLoggingOptionsProperty)
	// Experimental.
	PutProcessingConfiguration(value *TfDeliveryStream_SnowflakeConfigurationProcessingConfigurationProperty)
	// Experimental.
	PutS3Configuration(value *TfDeliveryStream_SnowflakeConfigurationS3ConfigurationProperty)
	// Experimental.
	PutSecretsManagerConfiguration(value *TfDeliveryStream_SnowflakeConfigurationSecretsManagerConfigurationProperty)
	// Experimental.
	PutSnowflakeRoleConfiguration(value *TfDeliveryStream_SnowflakeRoleConfigurationProperty)
	// Experimental.
	PutSnowflakeVpcConfiguration(value *TfDeliveryStream_SnowflakeVpcConfigurationProperty)
	// Experimental.
	ResetBufferingInterval()
	// Experimental.
	ResetBufferingSize()
	// Experimental.
	ResetCloudwatchLoggingOptions()
	// Experimental.
	ResetContentColumnName()
	// Experimental.
	ResetDataLoadingOption()
	// Experimental.
	ResetKeyPassphrase()
	// Experimental.
	ResetMetadataColumnName()
	// Experimental.
	ResetPrivateKey()
	// Experimental.
	ResetProcessingConfiguration()
	// Experimental.
	ResetRetryDuration()
	// Experimental.
	ResetS3BackupMode()
	// Experimental.
	ResetSecretsManagerConfiguration()
	// Experimental.
	ResetSnowflakeRoleConfiguration()
	// Experimental.
	ResetSnowflakeVpcConfiguration()
	// Experimental.
	ResetUser()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference
type jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) AccountUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) AccountUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) BufferingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) BufferingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) BufferingSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) BufferingSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) CloudwatchLoggingOptions() TfDeliveryStream_SnowflakeConfigurationCloudwatchLoggingOptionsPropertyOutputReference {
	var returns TfDeliveryStream_SnowflakeConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) CloudwatchLoggingOptionsInput() *TfDeliveryStream_SnowflakeConfigurationCloudwatchLoggingOptionsProperty {
	var returns *TfDeliveryStream_SnowflakeConfigurationCloudwatchLoggingOptionsProperty
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ContentColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ContentColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) DataLoadingOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLoadingOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) DataLoadingOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLoadingOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) InternalValue() *TfDeliveryStream_SnowflakeConfigurationProperty {
	var returns *TfDeliveryStream_SnowflakeConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) KeyPassphrase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPassphrase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) KeyPassphraseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPassphraseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) MetadataColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) MetadataColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ProcessingConfiguration() TfDeliveryStream_SnowflakeConfigurationProcessingConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_SnowflakeConfigurationProcessingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ProcessingConfigurationInput() *TfDeliveryStream_SnowflakeConfigurationProcessingConfigurationProperty {
	var returns *TfDeliveryStream_SnowflakeConfigurationProcessingConfigurationProperty
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) RetryDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) RetryDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) S3BackupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) S3BackupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) S3Configuration() TfDeliveryStream_SnowflakeConfigurationS3ConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_SnowflakeConfigurationS3ConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) S3ConfigurationInput() *TfDeliveryStream_SnowflakeConfigurationS3ConfigurationProperty {
	var returns *TfDeliveryStream_SnowflakeConfigurationS3ConfigurationProperty
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) SecretsManagerConfiguration() TfDeliveryStream_SnowflakeConfigurationSecretsManagerConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_SnowflakeConfigurationSecretsManagerConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"secretsManagerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) SecretsManagerConfigurationInput() *TfDeliveryStream_SnowflakeConfigurationSecretsManagerConfigurationProperty {
	var returns *TfDeliveryStream_SnowflakeConfigurationSecretsManagerConfigurationProperty
	_jsii_.Get(
		j,
		"secretsManagerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) SnowflakeRoleConfiguration() TfDeliveryStream_SnowflakeRoleConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_SnowflakeRoleConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"snowflakeRoleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) SnowflakeRoleConfigurationInput() *TfDeliveryStream_SnowflakeRoleConfigurationProperty {
	var returns *TfDeliveryStream_SnowflakeRoleConfigurationProperty
	_jsii_.Get(
		j,
		"snowflakeRoleConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) SnowflakeVpcConfiguration() TfDeliveryStream_SnowflakeVpcConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_SnowflakeVpcConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"snowflakeVpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) SnowflakeVpcConfigurationInput() *TfDeliveryStream_SnowflakeVpcConfigurationProperty {
	var returns *TfDeliveryStream_SnowflakeVpcConfigurationProperty
	_jsii_.Get(
		j,
		"snowflakeVpcConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) Table() *string {
	var returns *string
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) TableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDeliveryStream_SnowflakeConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDeliveryStream_SnowflakeConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream.SnowflakeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDeliveryStream_SnowflakeConfigurationPropertyOutputReference_Override(t TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream.SnowflakeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetAccountUrl(val *string) {
	if err := j.validateSetAccountUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountUrl",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetBufferingInterval(val *float64) {
	if err := j.validateSetBufferingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingInterval",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetBufferingSize(val *float64) {
	if err := j.validateSetBufferingSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingSize",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetContentColumnName(val *string) {
	if err := j.validateSetContentColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentColumnName",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetDataLoadingOption(val *string) {
	if err := j.validateSetDataLoadingOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataLoadingOption",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetInternalValue(val *TfDeliveryStream_SnowflakeConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetKeyPassphrase(val *string) {
	if err := j.validateSetKeyPassphraseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPassphrase",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetMetadataColumnName(val *string) {
	if err := j.validateSetMetadataColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataColumnName",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetPrivateKey(val *string) {
	if err := j.validateSetPrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetRetryDuration(val *float64) {
	if err := j.validateSetRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryDuration",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetS3BackupMode(val *string) {
	if err := j.validateSetS3BackupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BackupMode",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetTable(val *string) {
	if err := j.validateSetTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"table",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) PutCloudwatchLoggingOptions(value *TfDeliveryStream_SnowflakeConfigurationCloudwatchLoggingOptionsProperty) {
	if err := t.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) PutProcessingConfiguration(value *TfDeliveryStream_SnowflakeConfigurationProcessingConfigurationProperty) {
	if err := t.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) PutS3Configuration(value *TfDeliveryStream_SnowflakeConfigurationS3ConfigurationProperty) {
	if err := t.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) PutSecretsManagerConfiguration(value *TfDeliveryStream_SnowflakeConfigurationSecretsManagerConfigurationProperty) {
	if err := t.validatePutSecretsManagerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSecretsManagerConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) PutSnowflakeRoleConfiguration(value *TfDeliveryStream_SnowflakeRoleConfigurationProperty) {
	if err := t.validatePutSnowflakeRoleConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSnowflakeRoleConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) PutSnowflakeVpcConfiguration(value *TfDeliveryStream_SnowflakeVpcConfigurationProperty) {
	if err := t.validatePutSnowflakeVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSnowflakeVpcConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetBufferingInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetBufferingInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetBufferingSize() {
	_jsii_.InvokeVoid(
		t,
		"resetBufferingSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetContentColumnName() {
	_jsii_.InvokeVoid(
		t,
		"resetContentColumnName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetDataLoadingOption() {
	_jsii_.InvokeVoid(
		t,
		"resetDataLoadingOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetKeyPassphrase() {
	_jsii_.InvokeVoid(
		t,
		"resetKeyPassphrase",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetMetadataColumnName() {
	_jsii_.InvokeVoid(
		t,
		"resetMetadataColumnName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		t,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetRetryDuration() {
	_jsii_.InvokeVoid(
		t,
		"resetRetryDuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetS3BackupMode() {
	_jsii_.InvokeVoid(
		t,
		"resetS3BackupMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetSecretsManagerConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSecretsManagerConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetSnowflakeRoleConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSnowflakeRoleConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetSnowflakeVpcConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSnowflakeVpcConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		t,
		"resetUser",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

