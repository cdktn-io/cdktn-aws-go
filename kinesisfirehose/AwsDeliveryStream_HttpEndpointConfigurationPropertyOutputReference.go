package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/kinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessKey() *string
	// Experimental.
	SetAccessKey(val *string)
	// Experimental.
	AccessKeyInput() *string
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
	CloudwatchLoggingOptions() AwsDeliveryStream_HttpEndpointConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	// Experimental.
	CloudwatchLoggingOptionsInput() *AwsDeliveryStream_HttpEndpointConfigurationCloudwatchLoggingOptionsProperty
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDeliveryStream_HttpEndpointConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsDeliveryStream_HttpEndpointConfigurationProperty)
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	ProcessingConfiguration() AwsDeliveryStream_HttpEndpointConfigurationProcessingConfigurationPropertyOutputReference
	// Experimental.
	ProcessingConfigurationInput() *AwsDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProperty
	// Experimental.
	RequestConfiguration() AwsDeliveryStream_RequestConfigurationPropertyOutputReference
	// Experimental.
	RequestConfigurationInput() *AwsDeliveryStream_RequestConfigurationProperty
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
	S3Configuration() AwsDeliveryStream_HttpEndpointConfigurationS3ConfigurationPropertyOutputReference
	// Experimental.
	S3ConfigurationInput() *AwsDeliveryStream_HttpEndpointConfigurationS3ConfigurationProperty
	// Experimental.
	SecretsManagerConfiguration() AwsDeliveryStream_HttpEndpointConfigurationSecretsManagerConfigurationPropertyOutputReference
	// Experimental.
	SecretsManagerConfigurationInput() *AwsDeliveryStream_HttpEndpointConfigurationSecretsManagerConfigurationProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Url() *string
	// Experimental.
	SetUrl(val *string)
	// Experimental.
	UrlInput() *string
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
	PutCloudwatchLoggingOptions(value *AwsDeliveryStream_HttpEndpointConfigurationCloudwatchLoggingOptionsProperty)
	// Experimental.
	PutProcessingConfiguration(value *AwsDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProperty)
	// Experimental.
	PutRequestConfiguration(value *AwsDeliveryStream_RequestConfigurationProperty)
	// Experimental.
	PutS3Configuration(value *AwsDeliveryStream_HttpEndpointConfigurationS3ConfigurationProperty)
	// Experimental.
	PutSecretsManagerConfiguration(value *AwsDeliveryStream_HttpEndpointConfigurationSecretsManagerConfigurationProperty)
	// Experimental.
	ResetAccessKey()
	// Experimental.
	ResetBufferingInterval()
	// Experimental.
	ResetBufferingSize()
	// Experimental.
	ResetCloudwatchLoggingOptions()
	// Experimental.
	ResetName()
	// Experimental.
	ResetProcessingConfiguration()
	// Experimental.
	ResetRequestConfiguration()
	// Experimental.
	ResetRetryDuration()
	// Experimental.
	ResetRoleArn()
	// Experimental.
	ResetS3BackupMode()
	// Experimental.
	ResetSecretsManagerConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference
type jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) AccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) BufferingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) BufferingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) BufferingSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) BufferingSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) CloudwatchLoggingOptions() AwsDeliveryStream_HttpEndpointConfigurationCloudwatchLoggingOptionsPropertyOutputReference {
	var returns AwsDeliveryStream_HttpEndpointConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) CloudwatchLoggingOptionsInput() *AwsDeliveryStream_HttpEndpointConfigurationCloudwatchLoggingOptionsProperty {
	var returns *AwsDeliveryStream_HttpEndpointConfigurationCloudwatchLoggingOptionsProperty
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) InternalValue() *AwsDeliveryStream_HttpEndpointConfigurationProperty {
	var returns *AwsDeliveryStream_HttpEndpointConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ProcessingConfiguration() AwsDeliveryStream_HttpEndpointConfigurationProcessingConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_HttpEndpointConfigurationProcessingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ProcessingConfigurationInput() *AwsDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProperty {
	var returns *AwsDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProperty
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) RequestConfiguration() AwsDeliveryStream_RequestConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_RequestConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"requestConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) RequestConfigurationInput() *AwsDeliveryStream_RequestConfigurationProperty {
	var returns *AwsDeliveryStream_RequestConfigurationProperty
	_jsii_.Get(
		j,
		"requestConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) RetryDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) RetryDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) S3BackupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) S3BackupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) S3Configuration() AwsDeliveryStream_HttpEndpointConfigurationS3ConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_HttpEndpointConfigurationS3ConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) S3ConfigurationInput() *AwsDeliveryStream_HttpEndpointConfigurationS3ConfigurationProperty {
	var returns *AwsDeliveryStream_HttpEndpointConfigurationS3ConfigurationProperty
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) SecretsManagerConfiguration() AwsDeliveryStream_HttpEndpointConfigurationSecretsManagerConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_HttpEndpointConfigurationSecretsManagerConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"secretsManagerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) SecretsManagerConfigurationInput() *AwsDeliveryStream_HttpEndpointConfigurationSecretsManagerConfigurationProperty {
	var returns *AwsDeliveryStream_HttpEndpointConfigurationSecretsManagerConfigurationProperty
	_jsii_.Get(
		j,
		"secretsManagerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream.HttpEndpointConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference_Override(a AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream.HttpEndpointConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference)SetAccessKey(val *string) {
	if err := j.validateSetAccessKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference)SetBufferingInterval(val *float64) {
	if err := j.validateSetBufferingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingInterval",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference)SetBufferingSize(val *float64) {
	if err := j.validateSetBufferingSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingSize",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference)SetInternalValue(val *AwsDeliveryStream_HttpEndpointConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference)SetRetryDuration(val *float64) {
	if err := j.validateSetRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryDuration",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference)SetS3BackupMode(val *string) {
	if err := j.validateSetS3BackupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BackupMode",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) PutCloudwatchLoggingOptions(value *AwsDeliveryStream_HttpEndpointConfigurationCloudwatchLoggingOptionsProperty) {
	if err := a.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) PutProcessingConfiguration(value *AwsDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProperty) {
	if err := a.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) PutRequestConfiguration(value *AwsDeliveryStream_RequestConfigurationProperty) {
	if err := a.validatePutRequestConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRequestConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) PutS3Configuration(value *AwsDeliveryStream_HttpEndpointConfigurationS3ConfigurationProperty) {
	if err := a.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) PutSecretsManagerConfiguration(value *AwsDeliveryStream_HttpEndpointConfigurationSecretsManagerConfigurationProperty) {
	if err := a.validatePutSecretsManagerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecretsManagerConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ResetAccessKey() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ResetBufferingInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetBufferingInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ResetBufferingSize() {
	_jsii_.InvokeVoid(
		a,
		"resetBufferingSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ResetRequestConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ResetRetryDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ResetS3BackupMode() {
	_jsii_.InvokeVoid(
		a,
		"resetS3BackupMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ResetSecretsManagerConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretsManagerConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDeliveryStream_HttpEndpointConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

