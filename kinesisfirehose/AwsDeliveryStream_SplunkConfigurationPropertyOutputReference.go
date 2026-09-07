package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/kinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDeliveryStream_SplunkConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
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
	CloudwatchLoggingOptions() AwsDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	// Experimental.
	CloudwatchLoggingOptionsInput() *AwsDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsProperty
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
	HecAcknowledgmentTimeout() *float64
	// Experimental.
	SetHecAcknowledgmentTimeout(val *float64)
	// Experimental.
	HecAcknowledgmentTimeoutInput() *float64
	// Experimental.
	HecEndpoint() *string
	// Experimental.
	SetHecEndpoint(val *string)
	// Experimental.
	HecEndpointInput() *string
	// Experimental.
	HecEndpointType() *string
	// Experimental.
	SetHecEndpointType(val *string)
	// Experimental.
	HecEndpointTypeInput() *string
	// Experimental.
	HecToken() *string
	// Experimental.
	SetHecToken(val *string)
	// Experimental.
	HecTokenInput() *string
	// Experimental.
	InternalValue() *AwsDeliveryStream_SplunkConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsDeliveryStream_SplunkConfigurationProperty)
	// Experimental.
	ProcessingConfiguration() AwsDeliveryStream_SplunkConfigurationProcessingConfigurationPropertyOutputReference
	// Experimental.
	ProcessingConfigurationInput() *AwsDeliveryStream_SplunkConfigurationProcessingConfigurationProperty
	// Experimental.
	RetryDuration() *float64
	// Experimental.
	SetRetryDuration(val *float64)
	// Experimental.
	RetryDurationInput() *float64
	// Experimental.
	S3BackupMode() *string
	// Experimental.
	SetS3BackupMode(val *string)
	// Experimental.
	S3BackupModeInput() *string
	// Experimental.
	S3Configuration() AwsDeliveryStream_SplunkConfigurationS3ConfigurationPropertyOutputReference
	// Experimental.
	S3ConfigurationInput() *AwsDeliveryStream_SplunkConfigurationS3ConfigurationProperty
	// Experimental.
	SecretsManagerConfiguration() AwsDeliveryStream_SplunkConfigurationSecretsManagerConfigurationPropertyOutputReference
	// Experimental.
	SecretsManagerConfigurationInput() *AwsDeliveryStream_SplunkConfigurationSecretsManagerConfigurationProperty
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
	PutCloudwatchLoggingOptions(value *AwsDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsProperty)
	// Experimental.
	PutProcessingConfiguration(value *AwsDeliveryStream_SplunkConfigurationProcessingConfigurationProperty)
	// Experimental.
	PutS3Configuration(value *AwsDeliveryStream_SplunkConfigurationS3ConfigurationProperty)
	// Experimental.
	PutSecretsManagerConfiguration(value *AwsDeliveryStream_SplunkConfigurationSecretsManagerConfigurationProperty)
	// Experimental.
	ResetBufferingInterval()
	// Experimental.
	ResetBufferingSize()
	// Experimental.
	ResetCloudwatchLoggingOptions()
	// Experimental.
	ResetHecAcknowledgmentTimeout()
	// Experimental.
	ResetHecEndpointType()
	// Experimental.
	ResetHecToken()
	// Experimental.
	ResetProcessingConfiguration()
	// Experimental.
	ResetRetryDuration()
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

// The jsii proxy struct for AwsDeliveryStream_SplunkConfigurationPropertyOutputReference
type jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) BufferingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) BufferingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) BufferingSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) BufferingSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) CloudwatchLoggingOptions() AwsDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsPropertyOutputReference {
	var returns AwsDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) CloudwatchLoggingOptionsInput() *AwsDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsProperty {
	var returns *AwsDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsProperty
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) HecAcknowledgmentTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hecAcknowledgmentTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) HecAcknowledgmentTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hecAcknowledgmentTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) HecEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) HecEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) HecEndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecEndpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) HecEndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecEndpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) HecToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) HecTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) InternalValue() *AwsDeliveryStream_SplunkConfigurationProperty {
	var returns *AwsDeliveryStream_SplunkConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) ProcessingConfiguration() AwsDeliveryStream_SplunkConfigurationProcessingConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_SplunkConfigurationProcessingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) ProcessingConfigurationInput() *AwsDeliveryStream_SplunkConfigurationProcessingConfigurationProperty {
	var returns *AwsDeliveryStream_SplunkConfigurationProcessingConfigurationProperty
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) RetryDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) RetryDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) S3BackupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) S3BackupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) S3Configuration() AwsDeliveryStream_SplunkConfigurationS3ConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_SplunkConfigurationS3ConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) S3ConfigurationInput() *AwsDeliveryStream_SplunkConfigurationS3ConfigurationProperty {
	var returns *AwsDeliveryStream_SplunkConfigurationS3ConfigurationProperty
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) SecretsManagerConfiguration() AwsDeliveryStream_SplunkConfigurationSecretsManagerConfigurationPropertyOutputReference {
	var returns AwsDeliveryStream_SplunkConfigurationSecretsManagerConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"secretsManagerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) SecretsManagerConfigurationInput() *AwsDeliveryStream_SplunkConfigurationSecretsManagerConfigurationProperty {
	var returns *AwsDeliveryStream_SplunkConfigurationSecretsManagerConfigurationProperty
	_jsii_.Get(
		j,
		"secretsManagerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDeliveryStream_SplunkConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDeliveryStream_SplunkConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDeliveryStream_SplunkConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream.SplunkConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDeliveryStream_SplunkConfigurationPropertyOutputReference_Override(a AwsDeliveryStream_SplunkConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.AwsDeliveryStream.SplunkConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference)SetBufferingInterval(val *float64) {
	if err := j.validateSetBufferingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingInterval",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference)SetBufferingSize(val *float64) {
	if err := j.validateSetBufferingSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingSize",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference)SetHecAcknowledgmentTimeout(val *float64) {
	if err := j.validateSetHecAcknowledgmentTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hecAcknowledgmentTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference)SetHecEndpoint(val *string) {
	if err := j.validateSetHecEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hecEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference)SetHecEndpointType(val *string) {
	if err := j.validateSetHecEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hecEndpointType",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference)SetHecToken(val *string) {
	if err := j.validateSetHecTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hecToken",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference)SetInternalValue(val *AwsDeliveryStream_SplunkConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference)SetRetryDuration(val *float64) {
	if err := j.validateSetRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryDuration",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference)SetS3BackupMode(val *string) {
	if err := j.validateSetS3BackupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BackupMode",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) PutCloudwatchLoggingOptions(value *AwsDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsProperty) {
	if err := a.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) PutProcessingConfiguration(value *AwsDeliveryStream_SplunkConfigurationProcessingConfigurationProperty) {
	if err := a.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) PutS3Configuration(value *AwsDeliveryStream_SplunkConfigurationS3ConfigurationProperty) {
	if err := a.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) PutSecretsManagerConfiguration(value *AwsDeliveryStream_SplunkConfigurationSecretsManagerConfigurationProperty) {
	if err := a.validatePutSecretsManagerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecretsManagerConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetBufferingInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetBufferingInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetBufferingSize() {
	_jsii_.InvokeVoid(
		a,
		"resetBufferingSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetHecAcknowledgmentTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetHecAcknowledgmentTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetHecEndpointType() {
	_jsii_.InvokeVoid(
		a,
		"resetHecEndpointType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetHecToken() {
	_jsii_.InvokeVoid(
		a,
		"resetHecToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetRetryDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetS3BackupMode() {
	_jsii_.InvokeVoid(
		a,
		"resetS3BackupMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetSecretsManagerConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretsManagerConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDeliveryStream_SplunkConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

