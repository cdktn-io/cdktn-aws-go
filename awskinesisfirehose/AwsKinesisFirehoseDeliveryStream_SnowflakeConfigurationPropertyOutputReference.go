package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference interface {
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
	CloudwatchLoggingOptions() AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	// Experimental.
	CloudwatchLoggingOptionsInput() *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationCloudwatchLoggingOptionsProperty
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
	InternalValue() *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProperty)
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
	ProcessingConfiguration() AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProcessingConfigurationPropertyOutputReference
	// Experimental.
	ProcessingConfigurationInput() *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProcessingConfigurationProperty
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
	S3Configuration() AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationS3ConfigurationPropertyOutputReference
	// Experimental.
	S3ConfigurationInput() *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationS3ConfigurationProperty
	// Experimental.
	Schema() *string
	// Experimental.
	SetSchema(val *string)
	// Experimental.
	SchemaInput() *string
	// Experimental.
	SecretsManagerConfiguration() AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationSecretsManagerConfigurationPropertyOutputReference
	// Experimental.
	SecretsManagerConfigurationInput() *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationSecretsManagerConfigurationProperty
	// Experimental.
	SnowflakeRoleConfiguration() AwsKinesisFirehoseDeliveryStream_SnowflakeRoleConfigurationPropertyOutputReference
	// Experimental.
	SnowflakeRoleConfigurationInput() *AwsKinesisFirehoseDeliveryStream_SnowflakeRoleConfigurationProperty
	// Experimental.
	SnowflakeVpcConfiguration() AwsKinesisFirehoseDeliveryStream_SnowflakeVpcConfigurationPropertyOutputReference
	// Experimental.
	SnowflakeVpcConfigurationInput() *AwsKinesisFirehoseDeliveryStream_SnowflakeVpcConfigurationProperty
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
	PutCloudwatchLoggingOptions(value *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationCloudwatchLoggingOptionsProperty)
	// Experimental.
	PutProcessingConfiguration(value *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProcessingConfigurationProperty)
	// Experimental.
	PutS3Configuration(value *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationS3ConfigurationProperty)
	// Experimental.
	PutSecretsManagerConfiguration(value *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationSecretsManagerConfigurationProperty)
	// Experimental.
	PutSnowflakeRoleConfiguration(value *AwsKinesisFirehoseDeliveryStream_SnowflakeRoleConfigurationProperty)
	// Experimental.
	PutSnowflakeVpcConfiguration(value *AwsKinesisFirehoseDeliveryStream_SnowflakeVpcConfigurationProperty)
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

// The jsii proxy struct for AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference
type jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) AccountUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) AccountUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) BufferingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) BufferingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) BufferingSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) BufferingSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) CloudwatchLoggingOptions() AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationCloudwatchLoggingOptionsPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) CloudwatchLoggingOptionsInput() *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationCloudwatchLoggingOptionsProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationCloudwatchLoggingOptionsProperty
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ContentColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ContentColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) Database() *string {
	var returns *string
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) DatabaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) DataLoadingOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLoadingOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) DataLoadingOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLoadingOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) InternalValue() *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) KeyPassphrase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPassphrase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) KeyPassphraseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyPassphraseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) MetadataColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) MetadataColumnNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataColumnNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) PrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) PrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ProcessingConfiguration() AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProcessingConfigurationPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProcessingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ProcessingConfigurationInput() *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProcessingConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProcessingConfigurationProperty
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) RetryDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) RetryDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) S3BackupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) S3BackupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) S3Configuration() AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationS3ConfigurationPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationS3ConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) S3ConfigurationInput() *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationS3ConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationS3ConfigurationProperty
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) SchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) SecretsManagerConfiguration() AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationSecretsManagerConfigurationPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationSecretsManagerConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"secretsManagerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) SecretsManagerConfigurationInput() *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationSecretsManagerConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationSecretsManagerConfigurationProperty
	_jsii_.Get(
		j,
		"secretsManagerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) SnowflakeRoleConfiguration() AwsKinesisFirehoseDeliveryStream_SnowflakeRoleConfigurationPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_SnowflakeRoleConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"snowflakeRoleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) SnowflakeRoleConfigurationInput() *AwsKinesisFirehoseDeliveryStream_SnowflakeRoleConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_SnowflakeRoleConfigurationProperty
	_jsii_.Get(
		j,
		"snowflakeRoleConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) SnowflakeVpcConfiguration() AwsKinesisFirehoseDeliveryStream_SnowflakeVpcConfigurationPropertyOutputReference {
	var returns AwsKinesisFirehoseDeliveryStream_SnowflakeVpcConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"snowflakeVpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) SnowflakeVpcConfigurationInput() *AwsKinesisFirehoseDeliveryStream_SnowflakeVpcConfigurationProperty {
	var returns *AwsKinesisFirehoseDeliveryStream_SnowflakeVpcConfigurationProperty
	_jsii_.Get(
		j,
		"snowflakeVpcConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) Table() *string {
	var returns *string
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) TableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsKinesisFirehoseDeliveryStream.SnowflakeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference_Override(a AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsKinesisFirehoseDeliveryStream.SnowflakeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetAccountUrl(val *string) {
	if err := j.validateSetAccountUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountUrl",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetBufferingInterval(val *float64) {
	if err := j.validateSetBufferingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingInterval",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetBufferingSize(val *float64) {
	if err := j.validateSetBufferingSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingSize",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetContentColumnName(val *string) {
	if err := j.validateSetContentColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentColumnName",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetDatabase(val *string) {
	if err := j.validateSetDatabaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"database",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetDataLoadingOption(val *string) {
	if err := j.validateSetDataLoadingOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataLoadingOption",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetInternalValue(val *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetKeyPassphrase(val *string) {
	if err := j.validateSetKeyPassphraseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPassphrase",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetMetadataColumnName(val *string) {
	if err := j.validateSetMetadataColumnNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataColumnName",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetPrivateKey(val *string) {
	if err := j.validateSetPrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateKey",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetRetryDuration(val *float64) {
	if err := j.validateSetRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryDuration",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetS3BackupMode(val *string) {
	if err := j.validateSetS3BackupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BackupMode",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetSchema(val *string) {
	if err := j.validateSetSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schema",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetTable(val *string) {
	if err := j.validateSetTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"table",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) PutCloudwatchLoggingOptions(value *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationCloudwatchLoggingOptionsProperty) {
	if err := a.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) PutProcessingConfiguration(value *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationProcessingConfigurationProperty) {
	if err := a.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) PutS3Configuration(value *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationS3ConfigurationProperty) {
	if err := a.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) PutSecretsManagerConfiguration(value *AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationSecretsManagerConfigurationProperty) {
	if err := a.validatePutSecretsManagerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecretsManagerConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) PutSnowflakeRoleConfiguration(value *AwsKinesisFirehoseDeliveryStream_SnowflakeRoleConfigurationProperty) {
	if err := a.validatePutSnowflakeRoleConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnowflakeRoleConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) PutSnowflakeVpcConfiguration(value *AwsKinesisFirehoseDeliveryStream_SnowflakeVpcConfigurationProperty) {
	if err := a.validatePutSnowflakeVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnowflakeVpcConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetBufferingInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetBufferingInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetBufferingSize() {
	_jsii_.InvokeVoid(
		a,
		"resetBufferingSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetContentColumnName() {
	_jsii_.InvokeVoid(
		a,
		"resetContentColumnName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetDataLoadingOption() {
	_jsii_.InvokeVoid(
		a,
		"resetDataLoadingOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetKeyPassphrase() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyPassphrase",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetMetadataColumnName() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadataColumnName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetPrivateKey() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetRetryDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetS3BackupMode() {
	_jsii_.InvokeVoid(
		a,
		"resetS3BackupMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetSecretsManagerConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretsManagerConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetSnowflakeRoleConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowflakeRoleConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetSnowflakeVpcConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowflakeVpcConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		a,
		"resetUser",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_SnowflakeConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

