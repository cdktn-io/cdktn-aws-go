package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrock/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrock/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchConfig() TfModelInvocationLoggingConfiguration_CloudwatchConfigPropertyList
	// Experimental.
	CloudwatchConfigInput() interface{}
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
	EmbeddingDataDeliveryEnabled() interface{}
	// Experimental.
	SetEmbeddingDataDeliveryEnabled(val interface{})
	// Experimental.
	EmbeddingDataDeliveryEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	ImageDataDeliveryEnabled() interface{}
	// Experimental.
	SetImageDataDeliveryEnabled(val interface{})
	// Experimental.
	ImageDataDeliveryEnabledInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	S3Config() TfModelInvocationLoggingConfiguration_S3ConfigPropertyList
	// Experimental.
	S3ConfigInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TextDataDeliveryEnabled() interface{}
	// Experimental.
	SetTextDataDeliveryEnabled(val interface{})
	// Experimental.
	TextDataDeliveryEnabledInput() interface{}
	// Experimental.
	VideoDataDeliveryEnabled() interface{}
	// Experimental.
	SetVideoDataDeliveryEnabled(val interface{})
	// Experimental.
	VideoDataDeliveryEnabledInput() interface{}
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
	PutCloudwatchConfig(value interface{})
	// Experimental.
	PutS3Config(value interface{})
	// Experimental.
	ResetCloudwatchConfig()
	// Experimental.
	ResetEmbeddingDataDeliveryEnabled()
	// Experimental.
	ResetImageDataDeliveryEnabled()
	// Experimental.
	ResetS3Config()
	// Experimental.
	ResetTextDataDeliveryEnabled()
	// Experimental.
	ResetVideoDataDeliveryEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference
type jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) CloudwatchConfig() TfModelInvocationLoggingConfiguration_CloudwatchConfigPropertyList {
	var returns TfModelInvocationLoggingConfiguration_CloudwatchConfigPropertyList
	_jsii_.Get(
		j,
		"cloudwatchConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) CloudwatchConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) EmbeddingDataDeliveryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddingDataDeliveryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) EmbeddingDataDeliveryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"embeddingDataDeliveryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) ImageDataDeliveryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageDataDeliveryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) ImageDataDeliveryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageDataDeliveryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) S3Config() TfModelInvocationLoggingConfiguration_S3ConfigPropertyList {
	var returns TfModelInvocationLoggingConfiguration_S3ConfigPropertyList
	_jsii_.Get(
		j,
		"s3Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) S3ConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3ConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) TextDataDeliveryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"textDataDeliveryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) TextDataDeliveryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"textDataDeliveryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) VideoDataDeliveryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"videoDataDeliveryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) VideoDataDeliveryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"videoDataDeliveryEnabledInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock.TfModelInvocationLoggingConfiguration.LoggingConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference_Override(t TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock.TfModelInvocationLoggingConfiguration.LoggingConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference)SetEmbeddingDataDeliveryEnabled(val interface{}) {
	if err := j.validateSetEmbeddingDataDeliveryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"embeddingDataDeliveryEnabled",
		val,
	)
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference)SetImageDataDeliveryEnabled(val interface{}) {
	if err := j.validateSetImageDataDeliveryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageDataDeliveryEnabled",
		val,
	)
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference)SetTextDataDeliveryEnabled(val interface{}) {
	if err := j.validateSetTextDataDeliveryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textDataDeliveryEnabled",
		val,
	)
}

func (j *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference)SetVideoDataDeliveryEnabled(val interface{}) {
	if err := j.validateSetVideoDataDeliveryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"videoDataDeliveryEnabled",
		val,
	)
}

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) PutCloudwatchConfig(value interface{}) {
	if err := t.validatePutCloudwatchConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) PutS3Config(value interface{}) {
	if err := t.validatePutS3ConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Config",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) ResetCloudwatchConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) ResetEmbeddingDataDeliveryEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetEmbeddingDataDeliveryEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) ResetImageDataDeliveryEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetImageDataDeliveryEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) ResetS3Config() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Config",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) ResetTextDataDeliveryEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetTextDataDeliveryEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) ResetVideoDataDeliveryEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetVideoDataDeliveryEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfModelInvocationLoggingConfiguration_LoggingConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

