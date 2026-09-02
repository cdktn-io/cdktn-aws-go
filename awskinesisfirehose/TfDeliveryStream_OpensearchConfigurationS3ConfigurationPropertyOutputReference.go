package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference interface {
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
	CloudwatchLoggingOptions() TfDeliveryStream_OpensearchConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	// Experimental.
	CloudwatchLoggingOptionsInput() *TfDeliveryStream_OpensearchConfigurationS3ConfigurationCloudwatchLoggingOptionsProperty
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
	ErrorOutputPrefix() *string
	// Experimental.
	SetErrorOutputPrefix(val *string)
	// Experimental.
	ErrorOutputPrefixInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDeliveryStream_OpensearchConfigurationS3ConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfDeliveryStream_OpensearchConfigurationS3ConfigurationProperty)
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
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
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
	PutCloudwatchLoggingOptions(value *TfDeliveryStream_OpensearchConfigurationS3ConfigurationCloudwatchLoggingOptionsProperty)
	// Experimental.
	ResetBufferingInterval()
	// Experimental.
	ResetBufferingSize()
	// Experimental.
	ResetCloudwatchLoggingOptions()
	// Experimental.
	ResetCompressionFormat()
	// Experimental.
	ResetErrorOutputPrefix()
	// Experimental.
	ResetKmsKeyArn()
	// Experimental.
	ResetPrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference
type jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) BucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) BucketArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) BufferingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) BufferingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) BufferingSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) BufferingSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) CloudwatchLoggingOptions() TfDeliveryStream_OpensearchConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference {
	var returns TfDeliveryStream_OpensearchConfigurationS3ConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) CloudwatchLoggingOptionsInput() *TfDeliveryStream_OpensearchConfigurationS3ConfigurationCloudwatchLoggingOptionsProperty {
	var returns *TfDeliveryStream_OpensearchConfigurationS3ConfigurationCloudwatchLoggingOptionsProperty
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) CompressionFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) CompressionFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) ErrorOutputPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorOutputPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) ErrorOutputPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorOutputPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) InternalValue() *TfDeliveryStream_OpensearchConfigurationS3ConfigurationProperty {
	var returns *TfDeliveryStream_OpensearchConfigurationS3ConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream.OpensearchConfigurationS3ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference_Override(t TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream.OpensearchConfigurationS3ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference)SetBucketArn(val *string) {
	if err := j.validateSetBucketArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketArn",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference)SetBufferingInterval(val *float64) {
	if err := j.validateSetBufferingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingInterval",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference)SetBufferingSize(val *float64) {
	if err := j.validateSetBufferingSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingSize",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference)SetCompressionFormat(val *string) {
	if err := j.validateSetCompressionFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionFormat",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference)SetErrorOutputPrefix(val *string) {
	if err := j.validateSetErrorOutputPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorOutputPrefix",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference)SetInternalValue(val *TfDeliveryStream_OpensearchConfigurationS3ConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) PutCloudwatchLoggingOptions(value *TfDeliveryStream_OpensearchConfigurationS3ConfigurationCloudwatchLoggingOptionsProperty) {
	if err := t.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) ResetBufferingInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetBufferingInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) ResetBufferingSize() {
	_jsii_.InvokeVoid(
		t,
		"resetBufferingSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) ResetCompressionFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetCompressionFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) ResetErrorOutputPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetErrorOutputPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

