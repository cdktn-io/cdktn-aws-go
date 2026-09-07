package chimesdkmediapipelines

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/chimesdkmediapipelines/jsii"

	"github.com/cdktn-io/cdktn-aws-go/chimesdkmediapipelines/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AmazonTranscribeCallAnalyticsProcessorConfiguration() AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference
	// Experimental.
	AmazonTranscribeCallAnalyticsProcessorConfigurationInput() *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty
	// Experimental.
	AmazonTranscribeProcessorConfiguration() AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference
	// Experimental.
	AmazonTranscribeProcessorConfigurationInput() *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	KinesisDataStreamSinkConfiguration() AwsMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationPropertyOutputReference
	// Experimental.
	KinesisDataStreamSinkConfigurationInput() *AwsMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationProperty
	// Experimental.
	LambdaFunctionSinkConfiguration() AwsMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationPropertyOutputReference
	// Experimental.
	LambdaFunctionSinkConfigurationInput() *AwsMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationProperty
	// Experimental.
	S3RecordingSinkConfiguration() AwsMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationPropertyOutputReference
	// Experimental.
	S3RecordingSinkConfigurationInput() *AwsMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationProperty
	// Experimental.
	SnsTopicSinkConfiguration() AwsMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationPropertyOutputReference
	// Experimental.
	SnsTopicSinkConfigurationInput() *AwsMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationProperty
	// Experimental.
	SqsQueueSinkConfiguration() AwsMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationPropertyOutputReference
	// Experimental.
	SqsQueueSinkConfigurationInput() *AwsMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
	// Experimental.
	VoiceAnalyticsProcessorConfiguration() AwsMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference
	// Experimental.
	VoiceAnalyticsProcessorConfigurationInput() *AwsMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty
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
	PutAmazonTranscribeCallAnalyticsProcessorConfiguration(value *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty)
	// Experimental.
	PutAmazonTranscribeProcessorConfiguration(value *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty)
	// Experimental.
	PutKinesisDataStreamSinkConfiguration(value *AwsMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationProperty)
	// Experimental.
	PutLambdaFunctionSinkConfiguration(value *AwsMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationProperty)
	// Experimental.
	PutS3RecordingSinkConfiguration(value *AwsMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationProperty)
	// Experimental.
	PutSnsTopicSinkConfiguration(value *AwsMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationProperty)
	// Experimental.
	PutSqsQueueSinkConfiguration(value *AwsMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationProperty)
	// Experimental.
	PutVoiceAnalyticsProcessorConfiguration(value *AwsMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty)
	// Experimental.
	ResetAmazonTranscribeCallAnalyticsProcessorConfiguration()
	// Experimental.
	ResetAmazonTranscribeProcessorConfiguration()
	// Experimental.
	ResetKinesisDataStreamSinkConfiguration()
	// Experimental.
	ResetLambdaFunctionSinkConfiguration()
	// Experimental.
	ResetS3RecordingSinkConfiguration()
	// Experimental.
	ResetSnsTopicSinkConfiguration()
	// Experimental.
	ResetSqsQueueSinkConfiguration()
	// Experimental.
	ResetVoiceAnalyticsProcessorConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference
type jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) AmazonTranscribeCallAnalyticsProcessorConfiguration() AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference {
	var returns AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"amazonTranscribeCallAnalyticsProcessorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) AmazonTranscribeCallAnalyticsProcessorConfigurationInput() *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty {
	var returns *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty
	_jsii_.Get(
		j,
		"amazonTranscribeCallAnalyticsProcessorConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) AmazonTranscribeProcessorConfiguration() AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference {
	var returns AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"amazonTranscribeProcessorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) AmazonTranscribeProcessorConfigurationInput() *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty {
	var returns *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty
	_jsii_.Get(
		j,
		"amazonTranscribeProcessorConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) KinesisDataStreamSinkConfiguration() AwsMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationPropertyOutputReference {
	var returns AwsMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisDataStreamSinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) KinesisDataStreamSinkConfigurationInput() *AwsMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationProperty {
	var returns *AwsMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationProperty
	_jsii_.Get(
		j,
		"kinesisDataStreamSinkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) LambdaFunctionSinkConfiguration() AwsMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationPropertyOutputReference {
	var returns AwsMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"lambdaFunctionSinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) LambdaFunctionSinkConfigurationInput() *AwsMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationProperty {
	var returns *AwsMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationProperty
	_jsii_.Get(
		j,
		"lambdaFunctionSinkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) S3RecordingSinkConfiguration() AwsMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationPropertyOutputReference {
	var returns AwsMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3RecordingSinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) S3RecordingSinkConfigurationInput() *AwsMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationProperty {
	var returns *AwsMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationProperty
	_jsii_.Get(
		j,
		"s3RecordingSinkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) SnsTopicSinkConfiguration() AwsMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationPropertyOutputReference {
	var returns AwsMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"snsTopicSinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) SnsTopicSinkConfigurationInput() *AwsMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationProperty {
	var returns *AwsMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationProperty
	_jsii_.Get(
		j,
		"snsTopicSinkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) SqsQueueSinkConfiguration() AwsMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationPropertyOutputReference {
	var returns AwsMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"sqsQueueSinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) SqsQueueSinkConfigurationInput() *AwsMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationProperty {
	var returns *AwsMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationProperty
	_jsii_.Get(
		j,
		"sqsQueueSinkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) VoiceAnalyticsProcessorConfiguration() AwsMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference {
	var returns AwsMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"voiceAnalyticsProcessorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) VoiceAnalyticsProcessorConfigurationInput() *AwsMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty {
	var returns *AwsMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty
	_jsii_.Get(
		j,
		"voiceAnalyticsProcessorConfigurationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.ElementsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference_Override(a AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.AwsMediaInsightsPipelineConfiguration.ElementsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) PutAmazonTranscribeCallAnalyticsProcessorConfiguration(value *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty) {
	if err := a.validatePutAmazonTranscribeCallAnalyticsProcessorConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmazonTranscribeCallAnalyticsProcessorConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) PutAmazonTranscribeProcessorConfiguration(value *AwsMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty) {
	if err := a.validatePutAmazonTranscribeProcessorConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmazonTranscribeProcessorConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) PutKinesisDataStreamSinkConfiguration(value *AwsMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationProperty) {
	if err := a.validatePutKinesisDataStreamSinkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisDataStreamSinkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) PutLambdaFunctionSinkConfiguration(value *AwsMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationProperty) {
	if err := a.validatePutLambdaFunctionSinkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunctionSinkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) PutS3RecordingSinkConfiguration(value *AwsMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationProperty) {
	if err := a.validatePutS3RecordingSinkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3RecordingSinkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) PutSnsTopicSinkConfiguration(value *AwsMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationProperty) {
	if err := a.validatePutSnsTopicSinkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnsTopicSinkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) PutSqsQueueSinkConfiguration(value *AwsMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationProperty) {
	if err := a.validatePutSqsQueueSinkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqsQueueSinkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) PutVoiceAnalyticsProcessorConfiguration(value *AwsMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty) {
	if err := a.validatePutVoiceAnalyticsProcessorConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVoiceAnalyticsProcessorConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ResetAmazonTranscribeCallAnalyticsProcessorConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAmazonTranscribeCallAnalyticsProcessorConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ResetAmazonTranscribeProcessorConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAmazonTranscribeProcessorConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ResetKinesisDataStreamSinkConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisDataStreamSinkConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ResetLambdaFunctionSinkConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunctionSinkConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ResetS3RecordingSinkConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetS3RecordingSinkConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ResetSnsTopicSinkConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSnsTopicSinkConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ResetSqsQueueSinkConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSqsQueueSinkConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ResetVoiceAnalyticsProcessorConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetVoiceAnalyticsProcessorConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

