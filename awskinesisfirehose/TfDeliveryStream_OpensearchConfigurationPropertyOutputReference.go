package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDeliveryStream_OpensearchConfigurationPropertyOutputReference interface {
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
	CloudwatchLoggingOptions() TfDeliveryStream_OpensearchConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	// Experimental.
	CloudwatchLoggingOptionsInput() *TfDeliveryStream_OpensearchConfigurationCloudwatchLoggingOptionsProperty
	// Experimental.
	ClusterEndpoint() *string
	// Experimental.
	SetClusterEndpoint(val *string)
	// Experimental.
	ClusterEndpointInput() *string
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
	DocumentIdOptions() TfDeliveryStream_DocumentIdOptionsPropertyOutputReference
	// Experimental.
	DocumentIdOptionsInput() *TfDeliveryStream_DocumentIdOptionsProperty
	// Experimental.
	DomainArn() *string
	// Experimental.
	SetDomainArn(val *string)
	// Experimental.
	DomainArnInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	IndexName() *string
	// Experimental.
	SetIndexName(val *string)
	// Experimental.
	IndexNameInput() *string
	// Experimental.
	IndexRotationPeriod() *string
	// Experimental.
	SetIndexRotationPeriod(val *string)
	// Experimental.
	IndexRotationPeriodInput() *string
	// Experimental.
	InternalValue() *TfDeliveryStream_OpensearchConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfDeliveryStream_OpensearchConfigurationProperty)
	// Experimental.
	ProcessingConfiguration() TfDeliveryStream_OpensearchConfigurationProcessingConfigurationPropertyOutputReference
	// Experimental.
	ProcessingConfigurationInput() *TfDeliveryStream_OpensearchConfigurationProcessingConfigurationProperty
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
	S3Configuration() TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference
	// Experimental.
	S3ConfigurationInput() *TfDeliveryStream_OpensearchConfigurationS3ConfigurationProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TypeName() *string
	// Experimental.
	SetTypeName(val *string)
	// Experimental.
	TypeNameInput() *string
	// Experimental.
	VpcConfig() TfDeliveryStream_OpensearchConfigurationVpcConfigPropertyOutputReference
	// Experimental.
	VpcConfigInput() *TfDeliveryStream_OpensearchConfigurationVpcConfigProperty
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
	PutCloudwatchLoggingOptions(value *TfDeliveryStream_OpensearchConfigurationCloudwatchLoggingOptionsProperty)
	// Experimental.
	PutDocumentIdOptions(value *TfDeliveryStream_DocumentIdOptionsProperty)
	// Experimental.
	PutProcessingConfiguration(value *TfDeliveryStream_OpensearchConfigurationProcessingConfigurationProperty)
	// Experimental.
	PutS3Configuration(value *TfDeliveryStream_OpensearchConfigurationS3ConfigurationProperty)
	// Experimental.
	PutVpcConfig(value *TfDeliveryStream_OpensearchConfigurationVpcConfigProperty)
	// Experimental.
	ResetBufferingInterval()
	// Experimental.
	ResetBufferingSize()
	// Experimental.
	ResetCloudwatchLoggingOptions()
	// Experimental.
	ResetClusterEndpoint()
	// Experimental.
	ResetDocumentIdOptions()
	// Experimental.
	ResetDomainArn()
	// Experimental.
	ResetIndexRotationPeriod()
	// Experimental.
	ResetProcessingConfiguration()
	// Experimental.
	ResetRetryDuration()
	// Experimental.
	ResetS3BackupMode()
	// Experimental.
	ResetTypeName()
	// Experimental.
	ResetVpcConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDeliveryStream_OpensearchConfigurationPropertyOutputReference
type jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) BufferingInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) BufferingIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) BufferingSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) BufferingSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferingSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) CloudwatchLoggingOptions() TfDeliveryStream_OpensearchConfigurationCloudwatchLoggingOptionsPropertyOutputReference {
	var returns TfDeliveryStream_OpensearchConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) CloudwatchLoggingOptionsInput() *TfDeliveryStream_OpensearchConfigurationCloudwatchLoggingOptionsProperty {
	var returns *TfDeliveryStream_OpensearchConfigurationCloudwatchLoggingOptionsProperty
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ClusterEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ClusterEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) DocumentIdOptions() TfDeliveryStream_DocumentIdOptionsPropertyOutputReference {
	var returns TfDeliveryStream_DocumentIdOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"documentIdOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) DocumentIdOptionsInput() *TfDeliveryStream_DocumentIdOptionsProperty {
	var returns *TfDeliveryStream_DocumentIdOptionsProperty
	_jsii_.Get(
		j,
		"documentIdOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) DomainArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) DomainArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) IndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) IndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) IndexRotationPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexRotationPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) IndexRotationPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexRotationPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) InternalValue() *TfDeliveryStream_OpensearchConfigurationProperty {
	var returns *TfDeliveryStream_OpensearchConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ProcessingConfiguration() TfDeliveryStream_OpensearchConfigurationProcessingConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_OpensearchConfigurationProcessingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ProcessingConfigurationInput() *TfDeliveryStream_OpensearchConfigurationProcessingConfigurationProperty {
	var returns *TfDeliveryStream_OpensearchConfigurationProcessingConfigurationProperty
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) RetryDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) RetryDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) S3BackupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) S3BackupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) S3Configuration() TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_OpensearchConfigurationS3ConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) S3ConfigurationInput() *TfDeliveryStream_OpensearchConfigurationS3ConfigurationProperty {
	var returns *TfDeliveryStream_OpensearchConfigurationS3ConfigurationProperty
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) TypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) TypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) VpcConfig() TfDeliveryStream_OpensearchConfigurationVpcConfigPropertyOutputReference {
	var returns TfDeliveryStream_OpensearchConfigurationVpcConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) VpcConfigInput() *TfDeliveryStream_OpensearchConfigurationVpcConfigProperty {
	var returns *TfDeliveryStream_OpensearchConfigurationVpcConfigProperty
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDeliveryStream_OpensearchConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDeliveryStream_OpensearchConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDeliveryStream_OpensearchConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream.OpensearchConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDeliveryStream_OpensearchConfigurationPropertyOutputReference_Override(t TfDeliveryStream_OpensearchConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream.OpensearchConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference)SetBufferingInterval(val *float64) {
	if err := j.validateSetBufferingIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingInterval",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference)SetBufferingSize(val *float64) {
	if err := j.validateSetBufferingSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferingSize",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference)SetClusterEndpoint(val *string) {
	if err := j.validateSetClusterEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterEndpoint",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference)SetDomainArn(val *string) {
	if err := j.validateSetDomainArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainArn",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference)SetIndexName(val *string) {
	if err := j.validateSetIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexName",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference)SetIndexRotationPeriod(val *string) {
	if err := j.validateSetIndexRotationPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexRotationPeriod",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference)SetInternalValue(val *TfDeliveryStream_OpensearchConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference)SetRetryDuration(val *float64) {
	if err := j.validateSetRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryDuration",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference)SetS3BackupMode(val *string) {
	if err := j.validateSetS3BackupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BackupMode",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference)SetTypeName(val *string) {
	if err := j.validateSetTypeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeName",
		val,
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) PutCloudwatchLoggingOptions(value *TfDeliveryStream_OpensearchConfigurationCloudwatchLoggingOptionsProperty) {
	if err := t.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) PutDocumentIdOptions(value *TfDeliveryStream_DocumentIdOptionsProperty) {
	if err := t.validatePutDocumentIdOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDocumentIdOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) PutProcessingConfiguration(value *TfDeliveryStream_OpensearchConfigurationProcessingConfigurationProperty) {
	if err := t.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) PutS3Configuration(value *TfDeliveryStream_OpensearchConfigurationS3ConfigurationProperty) {
	if err := t.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) PutVpcConfig(value *TfDeliveryStream_OpensearchConfigurationVpcConfigProperty) {
	if err := t.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ResetBufferingInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetBufferingInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ResetBufferingSize() {
	_jsii_.InvokeVoid(
		t,
		"resetBufferingSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ResetClusterEndpoint() {
	_jsii_.InvokeVoid(
		t,
		"resetClusterEndpoint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ResetDocumentIdOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetDocumentIdOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ResetDomainArn() {
	_jsii_.InvokeVoid(
		t,
		"resetDomainArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ResetIndexRotationPeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetIndexRotationPeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ResetRetryDuration() {
	_jsii_.InvokeVoid(
		t,
		"resetRetryDuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ResetS3BackupMode() {
	_jsii_.InvokeVoid(
		t,
		"resetS3BackupMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ResetTypeName() {
	_jsii_.InvokeVoid(
		t,
		"resetTypeName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDeliveryStream_OpensearchConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

