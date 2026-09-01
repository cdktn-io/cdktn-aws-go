package awseventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPipesPipe_TargetParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BatchJobParameters() AwsPipesPipe_BatchJobParametersPropertyOutputReference
	// Experimental.
	BatchJobParametersInput() *AwsPipesPipe_BatchJobParametersProperty
	// Experimental.
	CloudwatchLogsParameters() AwsPipesPipe_CloudwatchLogsParametersPropertyOutputReference
	// Experimental.
	CloudwatchLogsParametersInput() *AwsPipesPipe_CloudwatchLogsParametersProperty
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
	EcsTaskParameters() AwsPipesPipe_EcsTaskParametersPropertyOutputReference
	// Experimental.
	EcsTaskParametersInput() *AwsPipesPipe_EcsTaskParametersProperty
	// Experimental.
	EventbridgeEventBusParameters() AwsPipesPipe_EventbridgeEventBusParametersPropertyOutputReference
	// Experimental.
	EventbridgeEventBusParametersInput() *AwsPipesPipe_EventbridgeEventBusParametersProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	HttpParameters() AwsPipesPipe_TargetParametersHttpParametersPropertyOutputReference
	// Experimental.
	HttpParametersInput() *AwsPipesPipe_TargetParametersHttpParametersProperty
	// Experimental.
	InputTemplate() *string
	// Experimental.
	SetInputTemplate(val *string)
	// Experimental.
	InputTemplateInput() *string
	// Experimental.
	InternalValue() *AwsPipesPipe_TargetParametersProperty
	// Experimental.
	SetInternalValue(val *AwsPipesPipe_TargetParametersProperty)
	// Experimental.
	KinesisStreamParameters() AwsPipesPipe_TargetParametersKinesisStreamParametersPropertyOutputReference
	// Experimental.
	KinesisStreamParametersInput() *AwsPipesPipe_TargetParametersKinesisStreamParametersProperty
	// Experimental.
	LambdaFunctionParameters() AwsPipesPipe_LambdaFunctionParametersPropertyOutputReference
	// Experimental.
	LambdaFunctionParametersInput() *AwsPipesPipe_LambdaFunctionParametersProperty
	// Experimental.
	RedshiftDataParameters() AwsPipesPipe_RedshiftDataParametersPropertyOutputReference
	// Experimental.
	RedshiftDataParametersInput() *AwsPipesPipe_RedshiftDataParametersProperty
	// Experimental.
	SagemakerPipelineParameters() AwsPipesPipe_SagemakerPipelineParametersPropertyOutputReference
	// Experimental.
	SagemakerPipelineParametersInput() *AwsPipesPipe_SagemakerPipelineParametersProperty
	// Experimental.
	SqsQueueParameters() AwsPipesPipe_TargetParametersSqsQueueParametersPropertyOutputReference
	// Experimental.
	SqsQueueParametersInput() *AwsPipesPipe_TargetParametersSqsQueueParametersProperty
	// Experimental.
	StepFunctionStateMachineParameters() AwsPipesPipe_StepFunctionStateMachineParametersPropertyOutputReference
	// Experimental.
	StepFunctionStateMachineParametersInput() *AwsPipesPipe_StepFunctionStateMachineParametersProperty
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
	PutBatchJobParameters(value *AwsPipesPipe_BatchJobParametersProperty)
	// Experimental.
	PutCloudwatchLogsParameters(value *AwsPipesPipe_CloudwatchLogsParametersProperty)
	// Experimental.
	PutEcsTaskParameters(value *AwsPipesPipe_EcsTaskParametersProperty)
	// Experimental.
	PutEventbridgeEventBusParameters(value *AwsPipesPipe_EventbridgeEventBusParametersProperty)
	// Experimental.
	PutHttpParameters(value *AwsPipesPipe_TargetParametersHttpParametersProperty)
	// Experimental.
	PutKinesisStreamParameters(value *AwsPipesPipe_TargetParametersKinesisStreamParametersProperty)
	// Experimental.
	PutLambdaFunctionParameters(value *AwsPipesPipe_LambdaFunctionParametersProperty)
	// Experimental.
	PutRedshiftDataParameters(value *AwsPipesPipe_RedshiftDataParametersProperty)
	// Experimental.
	PutSagemakerPipelineParameters(value *AwsPipesPipe_SagemakerPipelineParametersProperty)
	// Experimental.
	PutSqsQueueParameters(value *AwsPipesPipe_TargetParametersSqsQueueParametersProperty)
	// Experimental.
	PutStepFunctionStateMachineParameters(value *AwsPipesPipe_StepFunctionStateMachineParametersProperty)
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

// The jsii proxy struct for AwsPipesPipe_TargetParametersPropertyOutputReference
type jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) BatchJobParameters() AwsPipesPipe_BatchJobParametersPropertyOutputReference {
	var returns AwsPipesPipe_BatchJobParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"batchJobParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) BatchJobParametersInput() *AwsPipesPipe_BatchJobParametersProperty {
	var returns *AwsPipesPipe_BatchJobParametersProperty
	_jsii_.Get(
		j,
		"batchJobParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) CloudwatchLogsParameters() AwsPipesPipe_CloudwatchLogsParametersPropertyOutputReference {
	var returns AwsPipesPipe_CloudwatchLogsParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) CloudwatchLogsParametersInput() *AwsPipesPipe_CloudwatchLogsParametersProperty {
	var returns *AwsPipesPipe_CloudwatchLogsParametersProperty
	_jsii_.Get(
		j,
		"cloudwatchLogsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) EcsTaskParameters() AwsPipesPipe_EcsTaskParametersPropertyOutputReference {
	var returns AwsPipesPipe_EcsTaskParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"ecsTaskParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) EcsTaskParametersInput() *AwsPipesPipe_EcsTaskParametersProperty {
	var returns *AwsPipesPipe_EcsTaskParametersProperty
	_jsii_.Get(
		j,
		"ecsTaskParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) EventbridgeEventBusParameters() AwsPipesPipe_EventbridgeEventBusParametersPropertyOutputReference {
	var returns AwsPipesPipe_EventbridgeEventBusParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"eventbridgeEventBusParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) EventbridgeEventBusParametersInput() *AwsPipesPipe_EventbridgeEventBusParametersProperty {
	var returns *AwsPipesPipe_EventbridgeEventBusParametersProperty
	_jsii_.Get(
		j,
		"eventbridgeEventBusParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) HttpParameters() AwsPipesPipe_TargetParametersHttpParametersPropertyOutputReference {
	var returns AwsPipesPipe_TargetParametersHttpParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"httpParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) HttpParametersInput() *AwsPipesPipe_TargetParametersHttpParametersProperty {
	var returns *AwsPipesPipe_TargetParametersHttpParametersProperty
	_jsii_.Get(
		j,
		"httpParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) InputTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) InputTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) InternalValue() *AwsPipesPipe_TargetParametersProperty {
	var returns *AwsPipesPipe_TargetParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) KinesisStreamParameters() AwsPipesPipe_TargetParametersKinesisStreamParametersPropertyOutputReference {
	var returns AwsPipesPipe_TargetParametersKinesisStreamParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) KinesisStreamParametersInput() *AwsPipesPipe_TargetParametersKinesisStreamParametersProperty {
	var returns *AwsPipesPipe_TargetParametersKinesisStreamParametersProperty
	_jsii_.Get(
		j,
		"kinesisStreamParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) LambdaFunctionParameters() AwsPipesPipe_LambdaFunctionParametersPropertyOutputReference {
	var returns AwsPipesPipe_LambdaFunctionParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"lambdaFunctionParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) LambdaFunctionParametersInput() *AwsPipesPipe_LambdaFunctionParametersProperty {
	var returns *AwsPipesPipe_LambdaFunctionParametersProperty
	_jsii_.Get(
		j,
		"lambdaFunctionParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) RedshiftDataParameters() AwsPipesPipe_RedshiftDataParametersPropertyOutputReference {
	var returns AwsPipesPipe_RedshiftDataParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"redshiftDataParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) RedshiftDataParametersInput() *AwsPipesPipe_RedshiftDataParametersProperty {
	var returns *AwsPipesPipe_RedshiftDataParametersProperty
	_jsii_.Get(
		j,
		"redshiftDataParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) SagemakerPipelineParameters() AwsPipesPipe_SagemakerPipelineParametersPropertyOutputReference {
	var returns AwsPipesPipe_SagemakerPipelineParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"sagemakerPipelineParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) SagemakerPipelineParametersInput() *AwsPipesPipe_SagemakerPipelineParametersProperty {
	var returns *AwsPipesPipe_SagemakerPipelineParametersProperty
	_jsii_.Get(
		j,
		"sagemakerPipelineParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) SqsQueueParameters() AwsPipesPipe_TargetParametersSqsQueueParametersPropertyOutputReference {
	var returns AwsPipesPipe_TargetParametersSqsQueueParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"sqsQueueParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) SqsQueueParametersInput() *AwsPipesPipe_TargetParametersSqsQueueParametersProperty {
	var returns *AwsPipesPipe_TargetParametersSqsQueueParametersProperty
	_jsii_.Get(
		j,
		"sqsQueueParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) StepFunctionStateMachineParameters() AwsPipesPipe_StepFunctionStateMachineParametersPropertyOutputReference {
	var returns AwsPipesPipe_StepFunctionStateMachineParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"stepFunctionStateMachineParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) StepFunctionStateMachineParametersInput() *AwsPipesPipe_StepFunctionStateMachineParametersProperty {
	var returns *AwsPipesPipe_StepFunctionStateMachineParametersProperty
	_jsii_.Get(
		j,
		"stepFunctionStateMachineParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPipesPipe_TargetParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPipesPipe_TargetParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPipesPipe_TargetParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipesPipe.TargetParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPipesPipe_TargetParametersPropertyOutputReference_Override(a AwsPipesPipe_TargetParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipesPipe.TargetParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference)SetInputTemplate(val *string) {
	if err := j.validateSetInputTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputTemplate",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference)SetInternalValue(val *AwsPipesPipe_TargetParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) PutBatchJobParameters(value *AwsPipesPipe_BatchJobParametersProperty) {
	if err := a.validatePutBatchJobParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBatchJobParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) PutCloudwatchLogsParameters(value *AwsPipesPipe_CloudwatchLogsParametersProperty) {
	if err := a.validatePutCloudwatchLogsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchLogsParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) PutEcsTaskParameters(value *AwsPipesPipe_EcsTaskParametersProperty) {
	if err := a.validatePutEcsTaskParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcsTaskParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) PutEventbridgeEventBusParameters(value *AwsPipesPipe_EventbridgeEventBusParametersProperty) {
	if err := a.validatePutEventbridgeEventBusParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEventbridgeEventBusParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) PutHttpParameters(value *AwsPipesPipe_TargetParametersHttpParametersProperty) {
	if err := a.validatePutHttpParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) PutKinesisStreamParameters(value *AwsPipesPipe_TargetParametersKinesisStreamParametersProperty) {
	if err := a.validatePutKinesisStreamParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisStreamParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) PutLambdaFunctionParameters(value *AwsPipesPipe_LambdaFunctionParametersProperty) {
	if err := a.validatePutLambdaFunctionParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunctionParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) PutRedshiftDataParameters(value *AwsPipesPipe_RedshiftDataParametersProperty) {
	if err := a.validatePutRedshiftDataParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshiftDataParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) PutSagemakerPipelineParameters(value *AwsPipesPipe_SagemakerPipelineParametersProperty) {
	if err := a.validatePutSagemakerPipelineParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSagemakerPipelineParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) PutSqsQueueParameters(value *AwsPipesPipe_TargetParametersSqsQueueParametersProperty) {
	if err := a.validatePutSqsQueueParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqsQueueParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) PutStepFunctionStateMachineParameters(value *AwsPipesPipe_StepFunctionStateMachineParametersProperty) {
	if err := a.validatePutStepFunctionStateMachineParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStepFunctionStateMachineParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) ResetBatchJobParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetBatchJobParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) ResetCloudwatchLogsParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLogsParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) ResetEcsTaskParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetEcsTaskParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) ResetEventbridgeEventBusParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetEventbridgeEventBusParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) ResetHttpParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) ResetInputTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetInputTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) ResetKinesisStreamParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisStreamParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) ResetLambdaFunctionParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunctionParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) ResetRedshiftDataParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshiftDataParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) ResetSagemakerPipelineParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetSagemakerPipelineParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) ResetSqsQueueParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetSqsQueueParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) ResetStepFunctionStateMachineParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetStepFunctionStateMachineParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPipesPipe_TargetParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

