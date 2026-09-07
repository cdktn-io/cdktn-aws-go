package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/kinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BucketArn() *string
	// Experimental.
	SetBucketArn(val *string)
	// Experimental.
	BucketArnInput() *string
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
	CloudwatchLoggingOptions() AwsDeliveryStream_ExtendedS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	// Experimental.
	CloudwatchLoggingOptionsInput() *AwsDeliveryStream_ExtendedS3ConfigurationCloudwatchLoggingOptionsProperty
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
	CompressionFormat() *string
	// Experimental.
	SetCompressionFormat(val *string)
	// Experimental.
	CompressionFormatInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CustomTimeZone() *string
	// Experimental.
	SetCustomTimeZone(val *string)
	// Experimental.
	CustomTimeZoneInput() *string
	// Experimental.
	DataFormatConversionConfiguration() AwsDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference
	// Experimental.
	DataFormatConversionConfigurationInput() *AwsDeliveryStream_DataFormatConversionConfigurationProperty
	// Experimental.
	DynamicPartitioningConfiguration() AwsDeliveryStream_DynamicPartitioningConfigurationPropertyOutputReference
	// Experimental.
	DynamicPartitioningConfigurationInput() *AwsDeliveryStream_DynamicPartitioningConfigurationProperty
	// Experimental.
	ErrorOutputPrefix() *string
	// Experimental.
	SetErrorOutputPrefix(val *string)
	// Experimental.
	ErrorOutputPrefixInput() *string
	// Experimental.
	FileExtension() *string
	// Experimental.
	SetFileExtension(val *string)
	// Experimental.
	FileExtensionInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDeliveryStream_ExtendedS3ConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsDeliveryStream_ExtendedS3ConfigurationProperty)
	// Experimental.
	KmsKeyArn() *string
	// Experimental.
	SetKmsKeyArn(val *string)
	// Experimental.
	KmsKeyArnInput() *string
	// Experimental.
	Prefix() *string
	// Experimental.
	SetPrefix(val *string)
	// Experimental.
	PrefixInput() *string
	// Experimental.
	ProcessingConfiguration() AwsDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationPropertyOutputReference
	// Experimental.
	ProcessingConfigurationInput() *AwsDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProperty
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	S3BackupConfiguration() AwsDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationPropertyOutputReference
	// Experimental.
	S3BackupConfigurationInput() *AwsDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationProperty
	// Experimental.
	S3BackupMode() *string
	// Experimental.
	SetS3BackupMode(val *string)
	// Experimental.
	S3BackupModeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutCloudwatchLoggingOptions(value *AwsDeliveryStream_ExtendedS3ConfigurationCloudwatchLoggingOptionsProperty)
	// Experimental.
	PutDataFormatConversionConfiguration(value *AwsDeliveryStream_DataFormatConversionConfigurationProperty)
	// Experimental.
	PutDynamicPartitioningConfiguration(value *AwsDeliveryStream_DynamicPartitioningConfigurationProperty)
	// Experimental.
	PutProcessingConfiguration(value *AwsDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProperty)
	// Experimental.
	PutS3BackupConfiguration(value *AwsDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationProperty)
	// Experimental.
	ResetBufferingInterval()
	// Experimental.
	ResetBufferingSize()
	// Experimental.
	ResetCloudwatchLoggingOptions()
	// Experimental.
	ResetCompressionFormat()
	// Experimental.
	ResetCustomTimeZone()
	// Experimental.
	ResetDataFormatConversionConfiguration()
	// Experimental.
	ResetDynamicPartitioningConfiguration()
	// Experimental.
	ResetErrorOutputPrefix()
	// Experimental.
	ResetFileExtension()
	// Experimental.
	ResetKmsKeyArn()
	// Experimental.
	ResetPrefix()
	// Experimental.
	ResetProcessingConfiguration()
	// Experimental.
	ResetS3BackupConfiguration()
	// Experimental.
	ResetS3BackupMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference
type jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) BucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) BucketArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) BufferingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) BufferingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) BufferingSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) BufferingSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) CloudwatchLoggingOptions() AwsDeliveryStream_ExtendedS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference {
	var returns AwsDeliveryStream_ExtendedS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) CloudwatchLoggingOptionsInput() *AwsDeliveryStream_ExtendedS3ConfigurationCloudwatchLoggingOptionsProperty {
	var returns *AwsDeliveryStream_ExtendedS3ConfigurationCloudwatchLoggingOptionsProperty
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) CompressionFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) CompressionFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) CustomTimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customTimeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) CustomTimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customTimeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) DataFormatConversionConfiguration() AwsDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_DataFormatConversionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"dataFormatConversionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) DataFormatConversionConfigurationInput() *AwsDeliveryStream_DataFormatConversionConfigurationProperty {
	var returns *AwsDeliveryStream_DataFormatConversionConfigurationProperty
	_jsii_.Get(
		j,
		"dataFormatConversionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) DynamicPartitioningConfiguration() AwsDeliveryStream_DynamicPartitioningConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_DynamicPartitioningConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"dynamicPartitioningConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) DynamicPartitioningConfigurationInput() *AwsDeliveryStream_DynamicPartitioningConfigurationProperty {
	var returns *AwsDeliveryStream_DynamicPartitioningConfigurationProperty
	_jsii_.Get(
		j,
		"dynamicPartitioningConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ErrorOutputPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorOutputPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ErrorOutputPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorOutputPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) FileExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) FileExtensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileExtensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) InternalValue() *AwsDeliveryStream_ExtendedS3ConfigurationProperty {
	var returns *AwsDeliveryStream_ExtendedS3ConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ProcessingConfiguration() AwsDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ProcessingConfigurationInput() *AwsDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProperty {
	var returns *AwsDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProperty
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) S3BackupConfiguration() AwsDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3BackupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) S3BackupConfigurationInput() *AwsDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationProperty {
	var returns *AwsDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationProperty
	_jsii_.Get(
		j,
		"s3BackupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) S3BackupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) S3BackupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream.ExtendedS3ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference_Override(a AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream.ExtendedS3ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference)SetBucketArn(val *string) {
	if err := j.validateSetBucketArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketArn",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference)SetBufferingInterval(val *float64) {
	if err := j.validateSetBufferingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingInterval",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference)SetBufferingSize(val *float64) {
	if err := j.validateSetBufferingSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingSize",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference)SetCompressionFormat(val *string) {
	if err := j.validateSetCompressionFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionFormat",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference)SetCustomTimeZone(val *string) {
	if err := j.validateSetCustomTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customTimeZone",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference)SetErrorOutputPrefix(val *string) {
	if err := j.validateSetErrorOutputPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOutputPrefix",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference)SetFileExtension(val *string) {
	if err := j.validateSetFileExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileExtension",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference)SetInternalValue(val *AwsDeliveryStream_ExtendedS3ConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference)SetS3BackupMode(val *string) {
	if err := j.validateSetS3BackupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BackupMode",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) PutCloudwatchLoggingOptions(value *AwsDeliveryStream_ExtendedS3ConfigurationCloudwatchLoggingOptionsProperty) {
	if err := a.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) PutDataFormatConversionConfiguration(value *AwsDeliveryStream_DataFormatConversionConfigurationProperty) {
	if err := a.validatePutDataFormatConversionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataFormatConversionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) PutDynamicPartitioningConfiguration(value *AwsDeliveryStream_DynamicPartitioningConfigurationProperty) {
	if err := a.validatePutDynamicPartitioningConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynamicPartitioningConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) PutProcessingConfiguration(value *AwsDeliveryStream_ExtendedS3ConfigurationProcessingConfigurationProperty) {
	if err := a.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) PutS3BackupConfiguration(value *AwsDeliveryStream_ExtendedS3ConfigurationS3BackupConfigurationProperty) {
	if err := a.validatePutS3BackupConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3BackupConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ResetBufferingInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetBufferingInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ResetBufferingSize() {
	_jsii_.InvokeVoid(
		a,
		"resetBufferingSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ResetCompressionFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetCompressionFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ResetCustomTimeZone() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomTimeZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ResetDataFormatConversionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetDataFormatConversionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ResetDynamicPartitioningConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamicPartitioningConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ResetErrorOutputPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetErrorOutputPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ResetFileExtension() {
	_jsii_.InvokeVoid(
		a,
		"resetFileExtension",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ResetS3BackupConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetS3BackupConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ResetS3BackupMode() {
	_jsii_.InvokeVoid(
		a,
		"resetS3BackupMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDeliveryStream_ExtendedS3ConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

