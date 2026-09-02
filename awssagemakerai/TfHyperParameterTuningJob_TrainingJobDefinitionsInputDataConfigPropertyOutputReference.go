package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ChannelName() *string
	// Experimental.
	SetChannelName(val *string)
	// Experimental.
	ChannelNameInput() *string
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
	CompressionType() *string
	// Experimental.
	SetCompressionType(val *string)
	// Experimental.
	CompressionTypeInput() *string
	// Experimental.
	ContentType() *string
	// Experimental.
	SetContentType(val *string)
	// Experimental.
	ContentTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DataSource() TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigDataSourcePropertyList
	// Experimental.
	DataSourceInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InputMode() *string
	// Experimental.
	SetInputMode(val *string)
	// Experimental.
	InputModeInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	RecordWrapperType() *string
	// Experimental.
	SetRecordWrapperType(val *string)
	// Experimental.
	RecordWrapperTypeInput() *string
	// Experimental.
	ShuffleConfig() TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigShuffleConfigPropertyList
	// Experimental.
	ShuffleConfigInput() interface{}
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
	PutDataSource(value interface{})
	// Experimental.
	PutShuffleConfig(value interface{})
	// Experimental.
	ResetCompressionType()
	// Experimental.
	ResetContentType()
	// Experimental.
	ResetDataSource()
	// Experimental.
	ResetInputMode()
	// Experimental.
	ResetRecordWrapperType()
	// Experimental.
	ResetShuffleConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference
type jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) ChannelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) ChannelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) CompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) DataSource() TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigDataSourcePropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigDataSourcePropertyList
	_jsii_.Get(
		j,
		"dataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) DataSourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) InputMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) InputModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) RecordWrapperType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordWrapperType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) RecordWrapperTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordWrapperTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) ShuffleConfig() TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigShuffleConfigPropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigShuffleConfigPropertyList
	_jsii_.Get(
		j,
		"shuffleConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) ShuffleConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shuffleConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfHyperParameterTuningJob.TrainingJobDefinitionsInputDataConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference_Override(t TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfHyperParameterTuningJob.TrainingJobDefinitionsInputDataConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference)SetChannelName(val *string) {
	if err := j.validateSetChannelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelName",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference)SetCompressionType(val *string) {
	if err := j.validateSetCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionType",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference)SetInputMode(val *string) {
	if err := j.validateSetInputModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputMode",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference)SetRecordWrapperType(val *string) {
	if err := j.validateSetRecordWrapperTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordWrapperType",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) PutDataSource(value interface{}) {
	if err := t.validatePutDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDataSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) PutShuffleConfig(value interface{}) {
	if err := t.validatePutShuffleConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putShuffleConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) ResetCompressionType() {
	_jsii_.InvokeVoid(
		t,
		"resetCompressionType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) ResetContentType() {
	_jsii_.InvokeVoid(
		t,
		"resetContentType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) ResetDataSource() {
	_jsii_.InvokeVoid(
		t,
		"resetDataSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) ResetInputMode() {
	_jsii_.InvokeVoid(
		t,
		"resetInputMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) ResetRecordWrapperType() {
	_jsii_.InvokeVoid(
		t,
		"resetRecordWrapperType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) ResetShuffleConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetShuffleConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

