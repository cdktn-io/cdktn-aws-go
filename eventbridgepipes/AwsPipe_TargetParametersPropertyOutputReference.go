package eventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/eventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/eventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPipe_TargetParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BatchJobParameters() AwsPipe_BatchJobParametersPropertyOutputReference
	// Experimental.
	BatchJobParametersInput() *AwsPipe_BatchJobParametersProperty
	// Experimental.
	CloudwatchLogsParameters() AwsPipe_CloudwatchLogsParametersPropertyOutputReference
	// Experimental.
	CloudwatchLogsParametersInput() *AwsPipe_CloudwatchLogsParametersProperty
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
	EcsTaskParameters() AwsPipe_EcsTaskParametersPropertyOutputReference
	// Experimental.
	EcsTaskParametersInput() *AwsPipe_EcsTaskParametersProperty
	// Experimental.
	EventbridgeEventBusParameters() AwsPipe_EventbridgeEventBusParametersPropertyOutputReference
	// Experimental.
	EventbridgeEventBusParametersInput() *AwsPipe_EventbridgeEventBusParametersProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	HttpParameters() AwsPipe_TargetParametersHttpParametersPropertyOutputReference
	// Experimental.
	HttpParametersInput() *AwsPipe_TargetParametersHttpParametersProperty
	// Experimental.
	InputTemplate() *string
	// Experimental.
	SetInputTemplate(val *string)
	// Experimental.
	InputTemplateInput() *string
	// Experimental.
	InternalValue() *AwsPipe_TargetParametersProperty
	// Experimental.
	SetInternalValue(val *AwsPipe_TargetParametersProperty)
	// Experimental.
	KinesisStreamParameters() AwsPipe_TargetParametersKinesisStreamParametersPropertyOutputReference
	// Experimental.
	KinesisStreamParametersInput() *AwsPipe_TargetParametersKinesisStreamParametersProperty
	// Experimental.
	LambdaFunctionParameters() AwsPipe_LambdaFunctionParametersPropertyOutputReference
	// Experimental.
	LambdaFunctionParametersInput() *AwsPipe_LambdaFunctionParametersProperty
	// Experimental.
	RedshiftDataParameters() AwsPipe_RedshiftDataParametersPropertyOutputReference
	// Experimental.
	RedshiftDataParametersInput() *AwsPipe_RedshiftDataParametersProperty
	// Experimental.
	SagemakerPipelineParameters() AwsPipe_SagemakerPipelineParametersPropertyOutputReference
	// Experimental.
	SagemakerPipelineParametersInput() *AwsPipe_SagemakerPipelineParametersProperty
	// Experimental.
	SqsQueueParameters() AwsPipe_TargetParametersSqsQueueParametersPropertyOutputReference
	// Experimental.
	SqsQueueParametersInput() *AwsPipe_TargetParametersSqsQueueParametersProperty
	// Experimental.
	StepFunctionStateMachineParameters() AwsPipe_StepFunctionStateMachineParametersPropertyOutputReference
	// Experimental.
	StepFunctionStateMachineParametersInput() *AwsPipe_StepFunctionStateMachineParametersProperty
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
	PutBatchJobParameters(value *AwsPipe_BatchJobParametersProperty)
	// Experimental.
	PutCloudwatchLogsParameters(value *AwsPipe_CloudwatchLogsParametersProperty)
	// Experimental.
	PutEcsTaskParameters(value *AwsPipe_EcsTaskParametersProperty)
	// Experimental.
	PutEventbridgeEventBusParameters(value *AwsPipe_EventbridgeEventBusParametersProperty)
	// Experimental.
	PutHttpParameters(value *AwsPipe_TargetParametersHttpParametersProperty)
	// Experimental.
	PutKinesisStreamParameters(value *AwsPipe_TargetParametersKinesisStreamParametersProperty)
	// Experimental.
	PutLambdaFunctionParameters(value *AwsPipe_LambdaFunctionParametersProperty)
	// Experimental.
	PutRedshiftDataParameters(value *AwsPipe_RedshiftDataParametersProperty)
	// Experimental.
	PutSagemakerPipelineParameters(value *AwsPipe_SagemakerPipelineParametersProperty)
	// Experimental.
	PutSqsQueueParameters(value *AwsPipe_TargetParametersSqsQueueParametersProperty)
	// Experimental.
	PutStepFunctionStateMachineParameters(value *AwsPipe_StepFunctionStateMachineParametersProperty)
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

// The jsii proxy struct for AwsPipe_TargetParametersPropertyOutputReference
type jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) BatchJobParameters() AwsPipe_BatchJobParametersPropertyOutputReference {
	var returns AwsPipe_BatchJobParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"batchJobParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) BatchJobParametersInput() *AwsPipe_BatchJobParametersProperty {
	var returns *AwsPipe_BatchJobParametersProperty
	_jsii_.Get(
		j,
		"batchJobParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) CloudwatchLogsParameters() AwsPipe_CloudwatchLogsParametersPropertyOutputReference {
	var returns AwsPipe_CloudwatchLogsParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) CloudwatchLogsParametersInput() *AwsPipe_CloudwatchLogsParametersProperty {
	var returns *AwsPipe_CloudwatchLogsParametersProperty
	_jsii_.Get(
		j,
		"cloudwatchLogsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) EcsTaskParameters() AwsPipe_EcsTaskParametersPropertyOutputReference {
	var returns AwsPipe_EcsTaskParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"ecsTaskParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) EcsTaskParametersInput() *AwsPipe_EcsTaskParametersProperty {
	var returns *AwsPipe_EcsTaskParametersProperty
	_jsii_.Get(
		j,
		"ecsTaskParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) EventbridgeEventBusParameters() AwsPipe_EventbridgeEventBusParametersPropertyOutputReference {
	var returns AwsPipe_EventbridgeEventBusParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"eventbridgeEventBusParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) EventbridgeEventBusParametersInput() *AwsPipe_EventbridgeEventBusParametersProperty {
	var returns *AwsPipe_EventbridgeEventBusParametersProperty
	_jsii_.Get(
		j,
		"eventbridgeEventBusParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) HttpParameters() AwsPipe_TargetParametersHttpParametersPropertyOutputReference {
	var returns AwsPipe_TargetParametersHttpParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"httpParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) HttpParametersInput() *AwsPipe_TargetParametersHttpParametersProperty {
	var returns *AwsPipe_TargetParametersHttpParametersProperty
	_jsii_.Get(
		j,
		"httpParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) InputTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) InputTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) InternalValue() *AwsPipe_TargetParametersProperty {
	var returns *AwsPipe_TargetParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) KinesisStreamParameters() AwsPipe_TargetParametersKinesisStreamParametersPropertyOutputReference {
	var returns AwsPipe_TargetParametersKinesisStreamParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) KinesisStreamParametersInput() *AwsPipe_TargetParametersKinesisStreamParametersProperty {
	var returns *AwsPipe_TargetParametersKinesisStreamParametersProperty
	_jsii_.Get(
		j,
		"kinesisStreamParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) LambdaFunctionParameters() AwsPipe_LambdaFunctionParametersPropertyOutputReference {
	var returns AwsPipe_LambdaFunctionParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"lambdaFunctionParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) LambdaFunctionParametersInput() *AwsPipe_LambdaFunctionParametersProperty {
	var returns *AwsPipe_LambdaFunctionParametersProperty
	_jsii_.Get(
		j,
		"lambdaFunctionParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) RedshiftDataParameters() AwsPipe_RedshiftDataParametersPropertyOutputReference {
	var returns AwsPipe_RedshiftDataParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"redshiftDataParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) RedshiftDataParametersInput() *AwsPipe_RedshiftDataParametersProperty {
	var returns *AwsPipe_RedshiftDataParametersProperty
	_jsii_.Get(
		j,
		"redshiftDataParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) SagemakerPipelineParameters() AwsPipe_SagemakerPipelineParametersPropertyOutputReference {
	var returns AwsPipe_SagemakerPipelineParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"sagemakerPipelineParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) SagemakerPipelineParametersInput() *AwsPipe_SagemakerPipelineParametersProperty {
	var returns *AwsPipe_SagemakerPipelineParametersProperty
	_jsii_.Get(
		j,
		"sagemakerPipelineParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) SqsQueueParameters() AwsPipe_TargetParametersSqsQueueParametersPropertyOutputReference {
	var returns AwsPipe_TargetParametersSqsQueueParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"sqsQueueParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) SqsQueueParametersInput() *AwsPipe_TargetParametersSqsQueueParametersProperty {
	var returns *AwsPipe_TargetParametersSqsQueueParametersProperty
	_jsii_.Get(
		j,
		"sqsQueueParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) StepFunctionStateMachineParameters() AwsPipe_StepFunctionStateMachineParametersPropertyOutputReference {
	var returns AwsPipe_StepFunctionStateMachineParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"stepFunctionStateMachineParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) StepFunctionStateMachineParametersInput() *AwsPipe_StepFunctionStateMachineParametersProperty {
	var returns *AwsPipe_StepFunctionStateMachineParametersProperty
	_jsii_.Get(
		j,
		"stepFunctionStateMachineParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPipe_TargetParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPipe_TargetParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPipe_TargetParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipe.TargetParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPipe_TargetParametersPropertyOutputReference_Override(a AwsPipe_TargetParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipe.TargetParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference)SetInputTemplate(val *string) {
	if err := j.validateSetInputTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputTemplate",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference)SetInternalValue(val *AwsPipe_TargetParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) PutBatchJobParameters(value *AwsPipe_BatchJobParametersProperty) {
	if err := a.validatePutBatchJobParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBatchJobParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) PutCloudwatchLogsParameters(value *AwsPipe_CloudwatchLogsParametersProperty) {
	if err := a.validatePutCloudwatchLogsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchLogsParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) PutEcsTaskParameters(value *AwsPipe_EcsTaskParametersProperty) {
	if err := a.validatePutEcsTaskParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEcsTaskParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) PutEventbridgeEventBusParameters(value *AwsPipe_EventbridgeEventBusParametersProperty) {
	if err := a.validatePutEventbridgeEventBusParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEventbridgeEventBusParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) PutHttpParameters(value *AwsPipe_TargetParametersHttpParametersProperty) {
	if err := a.validatePutHttpParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) PutKinesisStreamParameters(value *AwsPipe_TargetParametersKinesisStreamParametersProperty) {
	if err := a.validatePutKinesisStreamParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesisStreamParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) PutLambdaFunctionParameters(value *AwsPipe_LambdaFunctionParametersProperty) {
	if err := a.validatePutLambdaFunctionParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambdaFunctionParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) PutRedshiftDataParameters(value *AwsPipe_RedshiftDataParametersProperty) {
	if err := a.validatePutRedshiftDataParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshiftDataParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) PutSagemakerPipelineParameters(value *AwsPipe_SagemakerPipelineParametersProperty) {
	if err := a.validatePutSagemakerPipelineParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSagemakerPipelineParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) PutSqsQueueParameters(value *AwsPipe_TargetParametersSqsQueueParametersProperty) {
	if err := a.validatePutSqsQueueParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqsQueueParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) PutStepFunctionStateMachineParameters(value *AwsPipe_StepFunctionStateMachineParametersProperty) {
	if err := a.validatePutStepFunctionStateMachineParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStepFunctionStateMachineParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) ResetBatchJobParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetBatchJobParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) ResetCloudwatchLogsParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLogsParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) ResetEcsTaskParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetEcsTaskParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) ResetEventbridgeEventBusParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetEventbridgeEventBusParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) ResetHttpParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) ResetInputTemplate() {
	_jsii_.InvokeVoid(
		a,
		"resetInputTemplate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) ResetKinesisStreamParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesisStreamParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) ResetLambdaFunctionParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFunctionParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) ResetRedshiftDataParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshiftDataParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) ResetSagemakerPipelineParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetSagemakerPipelineParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) ResetSqsQueueParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetSqsQueueParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) ResetStepFunctionStateMachineParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetStepFunctionStateMachineParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPipe_TargetParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

