package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDeliveryStream_SplunkConfigurationPropertyOutputReference interface {
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
	CloudwatchLoggingOptions() TfDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	// Experimental.
	CloudwatchLoggingOptionsInput() *TfDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsProperty
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
	InternalValue() *TfDeliveryStream_SplunkConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfDeliveryStream_SplunkConfigurationProperty)
	// Experimental.
	ProcessingConfiguration() TfDeliveryStream_SplunkConfigurationProcessingConfigurationPropertyOutputReference
	// Experimental.
	ProcessingConfigurationInput() *TfDeliveryStream_SplunkConfigurationProcessingConfigurationProperty
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
	S3Configuration() TfDeliveryStream_SplunkConfigurationS3ConfigurationPropertyOutputReference
	// Experimental.
	S3ConfigurationInput() *TfDeliveryStream_SplunkConfigurationS3ConfigurationProperty
	// Experimental.
	SecretsManagerConfiguration() TfDeliveryStream_SplunkConfigurationSecretsManagerConfigurationPropertyOutputReference
	// Experimental.
	SecretsManagerConfigurationInput() *TfDeliveryStream_SplunkConfigurationSecretsManagerConfigurationProperty
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
	PutCloudwatchLoggingOptions(value *TfDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsProperty)
	// Experimental.
	PutProcessingConfiguration(value *TfDeliveryStream_SplunkConfigurationProcessingConfigurationProperty)
	// Experimental.
	PutS3Configuration(value *TfDeliveryStream_SplunkConfigurationS3ConfigurationProperty)
	// Experimental.
	PutSecretsManagerConfiguration(value *TfDeliveryStream_SplunkConfigurationSecretsManagerConfigurationProperty)
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

// The jsii proxy struct for TfDeliveryStream_SplunkConfigurationPropertyOutputReference
type jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) BufferingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) BufferingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) BufferingSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) BufferingSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) CloudwatchLoggingOptions() TfDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsPropertyOutputReference {
	var returns TfDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) CloudwatchLoggingOptionsInput() *TfDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsProperty {
	var returns *TfDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsProperty
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) HecAcknowledgmentTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hecAcknowledgmentTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) HecAcknowledgmentTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hecAcknowledgmentTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) HecEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) HecEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) HecEndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecEndpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) HecEndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecEndpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) HecToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) HecTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hecTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) InternalValue() *TfDeliveryStream_SplunkConfigurationProperty {
	var returns *TfDeliveryStream_SplunkConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) ProcessingConfiguration() TfDeliveryStream_SplunkConfigurationProcessingConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_SplunkConfigurationProcessingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) ProcessingConfigurationInput() *TfDeliveryStream_SplunkConfigurationProcessingConfigurationProperty {
	var returns *TfDeliveryStream_SplunkConfigurationProcessingConfigurationProperty
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) RetryDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) RetryDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) S3BackupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) S3BackupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) S3Configuration() TfDeliveryStream_SplunkConfigurationS3ConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_SplunkConfigurationS3ConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) S3ConfigurationInput() *TfDeliveryStream_SplunkConfigurationS3ConfigurationProperty {
	var returns *TfDeliveryStream_SplunkConfigurationS3ConfigurationProperty
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) SecretsManagerConfiguration() TfDeliveryStream_SplunkConfigurationSecretsManagerConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_SplunkConfigurationSecretsManagerConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"secretsManagerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) SecretsManagerConfigurationInput() *TfDeliveryStream_SplunkConfigurationSecretsManagerConfigurationProperty {
	var returns *TfDeliveryStream_SplunkConfigurationSecretsManagerConfigurationProperty
	_jsii_.Get(
		j,
		"secretsManagerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDeliveryStream_SplunkConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDeliveryStream_SplunkConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDeliveryStream_SplunkConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream.SplunkConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDeliveryStream_SplunkConfigurationPropertyOutputReference_Override(t TfDeliveryStream_SplunkConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream.SplunkConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference)SetBufferingInterval(val *float64) {
	if err := j.validateSetBufferingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingInterval",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference)SetBufferingSize(val *float64) {
	if err := j.validateSetBufferingSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingSize",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference)SetHecAcknowledgmentTimeout(val *float64) {
	if err := j.validateSetHecAcknowledgmentTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hecAcknowledgmentTimeout",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference)SetHecEndpoint(val *string) {
	if err := j.validateSetHecEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hecEndpoint",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference)SetHecEndpointType(val *string) {
	if err := j.validateSetHecEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hecEndpointType",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference)SetHecToken(val *string) {
	if err := j.validateSetHecTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hecToken",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference)SetInternalValue(val *TfDeliveryStream_SplunkConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference)SetRetryDuration(val *float64) {
	if err := j.validateSetRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryDuration",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference)SetS3BackupMode(val *string) {
	if err := j.validateSetS3BackupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BackupMode",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) PutCloudwatchLoggingOptions(value *TfDeliveryStream_SplunkConfigurationCloudwatchLoggingOptionsProperty) {
	if err := t.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) PutProcessingConfiguration(value *TfDeliveryStream_SplunkConfigurationProcessingConfigurationProperty) {
	if err := t.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) PutS3Configuration(value *TfDeliveryStream_SplunkConfigurationS3ConfigurationProperty) {
	if err := t.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) PutSecretsManagerConfiguration(value *TfDeliveryStream_SplunkConfigurationSecretsManagerConfigurationProperty) {
	if err := t.validatePutSecretsManagerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSecretsManagerConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetBufferingInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetBufferingInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetBufferingSize() {
	_jsii_.InvokeVoid(
		t,
		"resetBufferingSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetHecAcknowledgmentTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetHecAcknowledgmentTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetHecEndpointType() {
	_jsii_.InvokeVoid(
		t,
		"resetHecEndpointType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetHecToken() {
	_jsii_.InvokeVoid(
		t,
		"resetHecToken",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetRetryDuration() {
	_jsii_.InvokeVoid(
		t,
		"resetRetryDuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetS3BackupMode() {
	_jsii_.InvokeVoid(
		t,
		"resetS3BackupMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) ResetSecretsManagerConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSecretsManagerConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDeliveryStream_SplunkConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

