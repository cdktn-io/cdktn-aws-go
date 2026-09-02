package awseventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPipe_TargetParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BatchJobParameters() TfPipe_BatchJobParametersPropertyOutputReference
	// Experimental.
	BatchJobParametersInput() *TfPipe_BatchJobParametersProperty
	// Experimental.
	CloudwatchLogsParameters() TfPipe_CloudwatchLogsParametersPropertyOutputReference
	// Experimental.
	CloudwatchLogsParametersInput() *TfPipe_CloudwatchLogsParametersProperty
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
	EcsTaskParameters() TfPipe_EcsTaskParametersPropertyOutputReference
	// Experimental.
	EcsTaskParametersInput() *TfPipe_EcsTaskParametersProperty
	// Experimental.
	EventbridgeEventBusParameters() TfPipe_EventbridgeEventBusParametersPropertyOutputReference
	// Experimental.
	EventbridgeEventBusParametersInput() *TfPipe_EventbridgeEventBusParametersProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	HttpParameters() TfPipe_TargetParametersHttpParametersPropertyOutputReference
	// Experimental.
	HttpParametersInput() *TfPipe_TargetParametersHttpParametersProperty
	// Experimental.
	InputTemplate() *string
	// Experimental.
	SetInputTemplate(val *string)
	// Experimental.
	InputTemplateInput() *string
	// Experimental.
	InternalValue() *TfPipe_TargetParametersProperty
	// Experimental.
	SetInternalValue(val *TfPipe_TargetParametersProperty)
	// Experimental.
	KinesisStreamParameters() TfPipe_TargetParametersKinesisStreamParametersPropertyOutputReference
	// Experimental.
	KinesisStreamParametersInput() *TfPipe_TargetParametersKinesisStreamParametersProperty
	// Experimental.
	LambdaFunctionParameters() TfPipe_LambdaFunctionParametersPropertyOutputReference
	// Experimental.
	LambdaFunctionParametersInput() *TfPipe_LambdaFunctionParametersProperty
	// Experimental.
	RedshiftDataParameters() TfPipe_RedshiftDataParametersPropertyOutputReference
	// Experimental.
	RedshiftDataParametersInput() *TfPipe_RedshiftDataParametersProperty
	// Experimental.
	SagemakerPipelineParameters() TfPipe_SagemakerPipelineParametersPropertyOutputReference
	// Experimental.
	SagemakerPipelineParametersInput() *TfPipe_SagemakerPipelineParametersProperty
	// Experimental.
	SqsQueueParameters() TfPipe_TargetParametersSqsQueueParametersPropertyOutputReference
	// Experimental.
	SqsQueueParametersInput() *TfPipe_TargetParametersSqsQueueParametersProperty
	// Experimental.
	StepFunctionStateMachineParameters() TfPipe_StepFunctionStateMachineParametersPropertyOutputReference
	// Experimental.
	StepFunctionStateMachineParametersInput() *TfPipe_StepFunctionStateMachineParametersProperty
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
	PutBatchJobParameters(value *TfPipe_BatchJobParametersProperty)
	// Experimental.
	PutCloudwatchLogsParameters(value *TfPipe_CloudwatchLogsParametersProperty)
	// Experimental.
	PutEcsTaskParameters(value *TfPipe_EcsTaskParametersProperty)
	// Experimental.
	PutEventbridgeEventBusParameters(value *TfPipe_EventbridgeEventBusParametersProperty)
	// Experimental.
	PutHttpParameters(value *TfPipe_TargetParametersHttpParametersProperty)
	// Experimental.
	PutKinesisStreamParameters(value *TfPipe_TargetParametersKinesisStreamParametersProperty)
	// Experimental.
	PutLambdaFunctionParameters(value *TfPipe_LambdaFunctionParametersProperty)
	// Experimental.
	PutRedshiftDataParameters(value *TfPipe_RedshiftDataParametersProperty)
	// Experimental.
	PutSagemakerPipelineParameters(value *TfPipe_SagemakerPipelineParametersProperty)
	// Experimental.
	PutSqsQueueParameters(value *TfPipe_TargetParametersSqsQueueParametersProperty)
	// Experimental.
	PutStepFunctionStateMachineParameters(value *TfPipe_StepFunctionStateMachineParametersProperty)
	// Experimental.
	ResetBatchJobParameters()
	// Experimental.
	ResetCloudwatchLogsParameters()
	// Experimental.
	ResetEcsTaskParameters()
	// Experimental.
	ResetEventbridgeEventBusParameters()
	// Experimental.
	ResetHttpParameters()
	// Experimental.
	ResetInputTemplate()
	// Experimental.
	ResetKinesisStreamParameters()
	// Experimental.
	ResetLambdaFunctionParameters()
	// Experimental.
	ResetRedshiftDataParameters()
	// Experimental.
	ResetSagemakerPipelineParameters()
	// Experimental.
	ResetSqsQueueParameters()
	// Experimental.
	ResetStepFunctionStateMachineParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPipe_TargetParametersPropertyOutputReference
type jsiiProxy_TfPipe_TargetParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) BatchJobParameters() TfPipe_BatchJobParametersPropertyOutputReference {
	var returns TfPipe_BatchJobParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"batchJobParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) BatchJobParametersInput() *TfPipe_BatchJobParametersProperty {
	var returns *TfPipe_BatchJobParametersProperty
	_jsii_.Get(
		j,
		"batchJobParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) CloudwatchLogsParameters() TfPipe_CloudwatchLogsParametersPropertyOutputReference {
	var returns TfPipe_CloudwatchLogsParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) CloudwatchLogsParametersInput() *TfPipe_CloudwatchLogsParametersProperty {
	var returns *TfPipe_CloudwatchLogsParametersProperty
	_jsii_.Get(
		j,
		"cloudwatchLogsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) EcsTaskParameters() TfPipe_EcsTaskParametersPropertyOutputReference {
	var returns TfPipe_EcsTaskParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"ecsTaskParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) EcsTaskParametersInput() *TfPipe_EcsTaskParametersProperty {
	var returns *TfPipe_EcsTaskParametersProperty
	_jsii_.Get(
		j,
		"ecsTaskParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) EventbridgeEventBusParameters() TfPipe_EventbridgeEventBusParametersPropertyOutputReference {
	var returns TfPipe_EventbridgeEventBusParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"eventbridgeEventBusParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) EventbridgeEventBusParametersInput() *TfPipe_EventbridgeEventBusParametersProperty {
	var returns *TfPipe_EventbridgeEventBusParametersProperty
	_jsii_.Get(
		j,
		"eventbridgeEventBusParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) HttpParameters() TfPipe_TargetParametersHttpParametersPropertyOutputReference {
	var returns TfPipe_TargetParametersHttpParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"httpParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) HttpParametersInput() *TfPipe_TargetParametersHttpParametersProperty {
	var returns *TfPipe_TargetParametersHttpParametersProperty
	_jsii_.Get(
		j,
		"httpParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) InputTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) InputTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) InternalValue() *TfPipe_TargetParametersProperty {
	var returns *TfPipe_TargetParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) KinesisStreamParameters() TfPipe_TargetParametersKinesisStreamParametersPropertyOutputReference {
	var returns TfPipe_TargetParametersKinesisStreamParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) KinesisStreamParametersInput() *TfPipe_TargetParametersKinesisStreamParametersProperty {
	var returns *TfPipe_TargetParametersKinesisStreamParametersProperty
	_jsii_.Get(
		j,
		"kinesisStreamParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) LambdaFunctionParameters() TfPipe_LambdaFunctionParametersPropertyOutputReference {
	var returns TfPipe_LambdaFunctionParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"lambdaFunctionParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) LambdaFunctionParametersInput() *TfPipe_LambdaFunctionParametersProperty {
	var returns *TfPipe_LambdaFunctionParametersProperty
	_jsii_.Get(
		j,
		"lambdaFunctionParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) RedshiftDataParameters() TfPipe_RedshiftDataParametersPropertyOutputReference {
	var returns TfPipe_RedshiftDataParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"redshiftDataParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) RedshiftDataParametersInput() *TfPipe_RedshiftDataParametersProperty {
	var returns *TfPipe_RedshiftDataParametersProperty
	_jsii_.Get(
		j,
		"redshiftDataParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) SagemakerPipelineParameters() TfPipe_SagemakerPipelineParametersPropertyOutputReference {
	var returns TfPipe_SagemakerPipelineParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"sagemakerPipelineParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) SagemakerPipelineParametersInput() *TfPipe_SagemakerPipelineParametersProperty {
	var returns *TfPipe_SagemakerPipelineParametersProperty
	_jsii_.Get(
		j,
		"sagemakerPipelineParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) SqsQueueParameters() TfPipe_TargetParametersSqsQueueParametersPropertyOutputReference {
	var returns TfPipe_TargetParametersSqsQueueParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"sqsQueueParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) SqsQueueParametersInput() *TfPipe_TargetParametersSqsQueueParametersProperty {
	var returns *TfPipe_TargetParametersSqsQueueParametersProperty
	_jsii_.Get(
		j,
		"sqsQueueParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) StepFunctionStateMachineParameters() TfPipe_StepFunctionStateMachineParametersPropertyOutputReference {
	var returns TfPipe_StepFunctionStateMachineParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"stepFunctionStateMachineParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) StepFunctionStateMachineParametersInput() *TfPipe_StepFunctionStateMachineParametersProperty {
	var returns *TfPipe_StepFunctionStateMachineParametersProperty
	_jsii_.Get(
		j,
		"stepFunctionStateMachineParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPipe_TargetParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPipe_TargetParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPipe_TargetParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPipe_TargetParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.TargetParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPipe_TargetParametersPropertyOutputReference_Override(t TfPipe_TargetParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.TargetParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference)SetInputTemplate(val *string) {
	if err := j.validateSetInputTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputTemplate",
		val,
	)
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference)SetInternalValue(val *TfPipe_TargetParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) PutBatchJobParameters(value *TfPipe_BatchJobParametersProperty) {
	if err := t.validatePutBatchJobParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBatchJobParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) PutCloudwatchLogsParameters(value *TfPipe_CloudwatchLogsParametersProperty) {
	if err := t.validatePutCloudwatchLogsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchLogsParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) PutEcsTaskParameters(value *TfPipe_EcsTaskParametersProperty) {
	if err := t.validatePutEcsTaskParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEcsTaskParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) PutEventbridgeEventBusParameters(value *TfPipe_EventbridgeEventBusParametersProperty) {
	if err := t.validatePutEventbridgeEventBusParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEventbridgeEventBusParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) PutHttpParameters(value *TfPipe_TargetParametersHttpParametersProperty) {
	if err := t.validatePutHttpParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHttpParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) PutKinesisStreamParameters(value *TfPipe_TargetParametersKinesisStreamParametersProperty) {
	if err := t.validatePutKinesisStreamParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesisStreamParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) PutLambdaFunctionParameters(value *TfPipe_LambdaFunctionParametersProperty) {
	if err := t.validatePutLambdaFunctionParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambdaFunctionParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) PutRedshiftDataParameters(value *TfPipe_RedshiftDataParametersProperty) {
	if err := t.validatePutRedshiftDataParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRedshiftDataParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) PutSagemakerPipelineParameters(value *TfPipe_SagemakerPipelineParametersProperty) {
	if err := t.validatePutSagemakerPipelineParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSagemakerPipelineParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) PutSqsQueueParameters(value *TfPipe_TargetParametersSqsQueueParametersProperty) {
	if err := t.validatePutSqsQueueParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSqsQueueParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) PutStepFunctionStateMachineParameters(value *TfPipe_StepFunctionStateMachineParametersProperty) {
	if err := t.validatePutStepFunctionStateMachineParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStepFunctionStateMachineParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) ResetBatchJobParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetBatchJobParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) ResetCloudwatchLogsParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchLogsParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) ResetEcsTaskParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetEcsTaskParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) ResetEventbridgeEventBusParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetEventbridgeEventBusParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) ResetHttpParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) ResetInputTemplate() {
	_jsii_.InvokeVoid(
		t,
		"resetInputTemplate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) ResetKinesisStreamParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesisStreamParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) ResetLambdaFunctionParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaFunctionParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) ResetRedshiftDataParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetRedshiftDataParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) ResetSagemakerPipelineParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetSagemakerPipelineParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) ResetSqsQueueParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetSqsQueueParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) ResetStepFunctionStateMachineParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetStepFunctionStateMachineParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPipe_TargetParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

