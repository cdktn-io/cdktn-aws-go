package awschimesdkmediapipelines

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awschimesdkmediapipelines/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awschimesdkmediapipelines/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AmazonTranscribeCallAnalyticsProcessorConfiguration() TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference
	// Experimental.
	AmazonTranscribeCallAnalyticsProcessorConfigurationInput() *TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty
	// Experimental.
	AmazonTranscribeProcessorConfiguration() TfMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference
	// Experimental.
	AmazonTranscribeProcessorConfigurationInput() *TfMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty
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
	KinesisDataStreamSinkConfiguration() TfMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationPropertyOutputReference
	// Experimental.
	KinesisDataStreamSinkConfigurationInput() *TfMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationProperty
	// Experimental.
	LambdaFunctionSinkConfiguration() TfMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationPropertyOutputReference
	// Experimental.
	LambdaFunctionSinkConfigurationInput() *TfMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationProperty
	// Experimental.
	S3RecordingSinkConfiguration() TfMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationPropertyOutputReference
	// Experimental.
	S3RecordingSinkConfigurationInput() *TfMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationProperty
	// Experimental.
	SnsTopicSinkConfiguration() TfMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationPropertyOutputReference
	// Experimental.
	SnsTopicSinkConfigurationInput() *TfMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationProperty
	// Experimental.
	SqsQueueSinkConfiguration() TfMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationPropertyOutputReference
	// Experimental.
	SqsQueueSinkConfigurationInput() *TfMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationProperty
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
	VoiceAnalyticsProcessorConfiguration() TfMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference
	// Experimental.
	VoiceAnalyticsProcessorConfigurationInput() *TfMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty
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
	PutAmazonTranscribeCallAnalyticsProcessorConfiguration(value *TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty)
	// Experimental.
	PutAmazonTranscribeProcessorConfiguration(value *TfMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty)
	// Experimental.
	PutKinesisDataStreamSinkConfiguration(value *TfMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationProperty)
	// Experimental.
	PutLambdaFunctionSinkConfiguration(value *TfMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationProperty)
	// Experimental.
	PutS3RecordingSinkConfiguration(value *TfMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationProperty)
	// Experimental.
	PutSnsTopicSinkConfiguration(value *TfMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationProperty)
	// Experimental.
	PutSqsQueueSinkConfiguration(value *TfMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationProperty)
	// Experimental.
	PutVoiceAnalyticsProcessorConfiguration(value *TfMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty)
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

// The jsii proxy struct for TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference
type jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) AmazonTranscribeCallAnalyticsProcessorConfiguration() TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference {
	var returns TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"amazonTranscribeCallAnalyticsProcessorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) AmazonTranscribeCallAnalyticsProcessorConfigurationInput() *TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty {
	var returns *TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty
	_jsii_.Get(
		j,
		"amazonTranscribeCallAnalyticsProcessorConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) AmazonTranscribeProcessorConfiguration() TfMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference {
	var returns TfMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"amazonTranscribeProcessorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) AmazonTranscribeProcessorConfigurationInput() *TfMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty {
	var returns *TfMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty
	_jsii_.Get(
		j,
		"amazonTranscribeProcessorConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) KinesisDataStreamSinkConfiguration() TfMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationPropertyOutputReference {
	var returns TfMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisDataStreamSinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) KinesisDataStreamSinkConfigurationInput() *TfMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationProperty {
	var returns *TfMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationProperty
	_jsii_.Get(
		j,
		"kinesisDataStreamSinkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) LambdaFunctionSinkConfiguration() TfMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationPropertyOutputReference {
	var returns TfMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"lambdaFunctionSinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) LambdaFunctionSinkConfigurationInput() *TfMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationProperty {
	var returns *TfMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationProperty
	_jsii_.Get(
		j,
		"lambdaFunctionSinkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) S3RecordingSinkConfiguration() TfMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationPropertyOutputReference {
	var returns TfMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3RecordingSinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) S3RecordingSinkConfigurationInput() *TfMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationProperty {
	var returns *TfMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationProperty
	_jsii_.Get(
		j,
		"s3RecordingSinkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) SnsTopicSinkConfiguration() TfMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationPropertyOutputReference {
	var returns TfMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"snsTopicSinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) SnsTopicSinkConfigurationInput() *TfMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationProperty {
	var returns *TfMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationProperty
	_jsii_.Get(
		j,
		"snsTopicSinkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) SqsQueueSinkConfiguration() TfMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationPropertyOutputReference {
	var returns TfMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"sqsQueueSinkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) SqsQueueSinkConfigurationInput() *TfMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationProperty {
	var returns *TfMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationProperty
	_jsii_.Get(
		j,
		"sqsQueueSinkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) VoiceAnalyticsProcessorConfiguration() TfMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference {
	var returns TfMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"voiceAnalyticsProcessorConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) VoiceAnalyticsProcessorConfigurationInput() *TfMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty {
	var returns *TfMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty
	_jsii_.Get(
		j,
		"voiceAnalyticsProcessorConfigurationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.TfMediaInsightsPipelineConfiguration.ElementsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference_Override(t TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-chime-sdk-media-pipelines.TfMediaInsightsPipelineConfiguration.ElementsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) PutAmazonTranscribeCallAnalyticsProcessorConfiguration(value *TfMediaInsightsPipelineConfiguration_AmazonTranscribeCallAnalyticsProcessorConfigurationProperty) {
	if err := t.validatePutAmazonTranscribeCallAnalyticsProcessorConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAmazonTranscribeCallAnalyticsProcessorConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) PutAmazonTranscribeProcessorConfiguration(value *TfMediaInsightsPipelineConfiguration_AmazonTranscribeProcessorConfigurationProperty) {
	if err := t.validatePutAmazonTranscribeProcessorConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAmazonTranscribeProcessorConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) PutKinesisDataStreamSinkConfiguration(value *TfMediaInsightsPipelineConfiguration_KinesisDataStreamSinkConfigurationProperty) {
	if err := t.validatePutKinesisDataStreamSinkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesisDataStreamSinkConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) PutLambdaFunctionSinkConfiguration(value *TfMediaInsightsPipelineConfiguration_LambdaFunctionSinkConfigurationProperty) {
	if err := t.validatePutLambdaFunctionSinkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambdaFunctionSinkConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) PutS3RecordingSinkConfiguration(value *TfMediaInsightsPipelineConfiguration_S3RecordingSinkConfigurationProperty) {
	if err := t.validatePutS3RecordingSinkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3RecordingSinkConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) PutSnsTopicSinkConfiguration(value *TfMediaInsightsPipelineConfiguration_SnsTopicSinkConfigurationProperty) {
	if err := t.validatePutSnsTopicSinkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSnsTopicSinkConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) PutSqsQueueSinkConfiguration(value *TfMediaInsightsPipelineConfiguration_SqsQueueSinkConfigurationProperty) {
	if err := t.validatePutSqsQueueSinkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSqsQueueSinkConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) PutVoiceAnalyticsProcessorConfiguration(value *TfMediaInsightsPipelineConfiguration_VoiceAnalyticsProcessorConfigurationProperty) {
	if err := t.validatePutVoiceAnalyticsProcessorConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVoiceAnalyticsProcessorConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ResetAmazonTranscribeCallAnalyticsProcessorConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetAmazonTranscribeCallAnalyticsProcessorConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ResetAmazonTranscribeProcessorConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetAmazonTranscribeProcessorConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ResetKinesisDataStreamSinkConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesisDataStreamSinkConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ResetLambdaFunctionSinkConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaFunctionSinkConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ResetS3RecordingSinkConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetS3RecordingSinkConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ResetSnsTopicSinkConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSnsTopicSinkConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ResetSqsQueueSinkConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSqsQueueSinkConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ResetVoiceAnalyticsProcessorConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetVoiceAnalyticsProcessorConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfMediaInsightsPipelineConfiguration_ElementsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

