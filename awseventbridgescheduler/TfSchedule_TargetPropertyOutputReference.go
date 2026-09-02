package awseventbridgescheduler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgescheduler/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgescheduler/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSchedule_TargetPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Arn() *string
	// Experimental.
	SetArn(val *string)
	// Experimental.
	ArnInput() *string
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
	DeadLetterConfig() TfSchedule_DeadLetterConfigPropertyOutputReference
	// Experimental.
	DeadLetterConfigInput() *TfSchedule_DeadLetterConfigProperty
	// Experimental.
	EcsParameters() TfSchedule_EcsParametersPropertyOutputReference
	// Experimental.
	EcsParametersInput() *TfSchedule_EcsParametersProperty
	// Experimental.
	EventbridgeParameters() TfSchedule_EventbridgeParametersPropertyOutputReference
	// Experimental.
	EventbridgeParametersInput() *TfSchedule_EventbridgeParametersProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Input() *string
	// Experimental.
	SetInput(val *string)
	// Experimental.
	InputInput() *string
	// Experimental.
	InternalValue() *TfSchedule_TargetProperty
	// Experimental.
	SetInternalValue(val *TfSchedule_TargetProperty)
	// Experimental.
	KinesisParameters() TfSchedule_KinesisParametersPropertyOutputReference
	// Experimental.
	KinesisParametersInput() *TfSchedule_KinesisParametersProperty
	// Experimental.
	RetryPolicy() TfSchedule_RetryPolicyPropertyOutputReference
	// Experimental.
	RetryPolicyInput() *TfSchedule_RetryPolicyProperty
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	SagemakerPipelineParameters() TfSchedule_SagemakerPipelineParametersPropertyOutputReference
	// Experimental.
	SagemakerPipelineParametersInput() *TfSchedule_SagemakerPipelineParametersProperty
	// Experimental.
	SqsParameters() TfSchedule_SqsParametersPropertyOutputReference
	// Experimental.
	SqsParametersInput() *TfSchedule_SqsParametersProperty
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
	PutDeadLetterConfig(value *TfSchedule_DeadLetterConfigProperty)
	// Experimental.
	PutEcsParameters(value *TfSchedule_EcsParametersProperty)
	// Experimental.
	PutEventbridgeParameters(value *TfSchedule_EventbridgeParametersProperty)
	// Experimental.
	PutKinesisParameters(value *TfSchedule_KinesisParametersProperty)
	// Experimental.
	PutRetryPolicy(value *TfSchedule_RetryPolicyProperty)
	// Experimental.
	PutSagemakerPipelineParameters(value *TfSchedule_SagemakerPipelineParametersProperty)
	// Experimental.
	PutSqsParameters(value *TfSchedule_SqsParametersProperty)
	// Experimental.
	ResetDeadLetterConfig()
	// Experimental.
	ResetEcsParameters()
	// Experimental.
	ResetEventbridgeParameters()
	// Experimental.
	ResetInput()
	// Experimental.
	ResetKinesisParameters()
	// Experimental.
	ResetRetryPolicy()
	// Experimental.
	ResetSagemakerPipelineParameters()
	// Experimental.
	ResetSqsParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfSchedule_TargetPropertyOutputReference
type jsiiProxy_TfSchedule_TargetPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) ArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) DeadLetterConfig() TfSchedule_DeadLetterConfigPropertyOutputReference {
	var returns TfSchedule_DeadLetterConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) DeadLetterConfigInput() *TfSchedule_DeadLetterConfigProperty {
	var returns *TfSchedule_DeadLetterConfigProperty
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) EcsParameters() TfSchedule_EcsParametersPropertyOutputReference {
	var returns TfSchedule_EcsParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"ecsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) EcsParametersInput() *TfSchedule_EcsParametersProperty {
	var returns *TfSchedule_EcsParametersProperty
	_jsii_.Get(
		j,
		"ecsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) EventbridgeParameters() TfSchedule_EventbridgeParametersPropertyOutputReference {
	var returns TfSchedule_EventbridgeParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"eventbridgeParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) EventbridgeParametersInput() *TfSchedule_EventbridgeParametersProperty {
	var returns *TfSchedule_EventbridgeParametersProperty
	_jsii_.Get(
		j,
		"eventbridgeParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) InputInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) InternalValue() *TfSchedule_TargetProperty {
	var returns *TfSchedule_TargetProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) KinesisParameters() TfSchedule_KinesisParametersPropertyOutputReference {
	var returns TfSchedule_KinesisParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) KinesisParametersInput() *TfSchedule_KinesisParametersProperty {
	var returns *TfSchedule_KinesisParametersProperty
	_jsii_.Get(
		j,
		"kinesisParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) RetryPolicy() TfSchedule_RetryPolicyPropertyOutputReference {
	var returns TfSchedule_RetryPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) RetryPolicyInput() *TfSchedule_RetryPolicyProperty {
	var returns *TfSchedule_RetryPolicyProperty
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) SagemakerPipelineParameters() TfSchedule_SagemakerPipelineParametersPropertyOutputReference {
	var returns TfSchedule_SagemakerPipelineParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"sagemakerPipelineParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) SagemakerPipelineParametersInput() *TfSchedule_SagemakerPipelineParametersProperty {
	var returns *TfSchedule_SagemakerPipelineParametersProperty
	_jsii_.Get(
		j,
		"sagemakerPipelineParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) SqsParameters() TfSchedule_SqsParametersPropertyOutputReference {
	var returns TfSchedule_SqsParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"sqsParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) SqsParametersInput() *TfSchedule_SqsParametersProperty {
	var returns *TfSchedule_SqsParametersProperty
	_jsii_.Get(
		j,
		"sqsParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfSchedule_TargetPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfSchedule_TargetPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfSchedule_TargetPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfSchedule_TargetPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-scheduler.TfSchedule.TargetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfSchedule_TargetPropertyOutputReference_Override(t TfSchedule_TargetPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-scheduler.TfSchedule.TargetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference)SetInput(val *string) {
	if err := j.validateSetInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference)SetInternalValue(val *TfSchedule_TargetProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfSchedule_TargetPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) PutDeadLetterConfig(value *TfSchedule_DeadLetterConfigProperty) {
	if err := t.validatePutDeadLetterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) PutEcsParameters(value *TfSchedule_EcsParametersProperty) {
	if err := t.validatePutEcsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEcsParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) PutEventbridgeParameters(value *TfSchedule_EventbridgeParametersProperty) {
	if err := t.validatePutEventbridgeParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEventbridgeParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) PutKinesisParameters(value *TfSchedule_KinesisParametersProperty) {
	if err := t.validatePutKinesisParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesisParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) PutRetryPolicy(value *TfSchedule_RetryPolicyProperty) {
	if err := t.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) PutSagemakerPipelineParameters(value *TfSchedule_SagemakerPipelineParametersProperty) {
	if err := t.validatePutSagemakerPipelineParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSagemakerPipelineParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) PutSqsParameters(value *TfSchedule_SqsParametersProperty) {
	if err := t.validatePutSqsParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSqsParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) ResetEcsParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetEcsParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) ResetEventbridgeParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetEventbridgeParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) ResetInput() {
	_jsii_.InvokeVoid(
		t,
		"resetInput",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) ResetKinesisParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesisParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) ResetSagemakerPipelineParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetSagemakerPipelineParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) ResetSqsParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetSqsParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfSchedule_TargetPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

